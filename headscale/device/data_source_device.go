package device

import (
	"context"

	"github.com/awlsring/terraform-provider-headscale/internal/gen/models"
	"github.com/awlsring/terraform-provider-headscale/internal/service"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &deviceDataSource{}
	_ datasource.DataSourceWithConfigure = &deviceDataSource{}
)

func DataSource() datasource.DataSource {
	return &deviceDataSource{}
}

type deviceDataSource struct {
	client service.Headscale
}

func (d *deviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device"
}

func (d *deviceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(service.Headscale)
}

func (d *deviceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The device data source allows you to get information about a device registered on the Headscale instance.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The id of the device",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The device's name.",
			},
			"given_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The device's given name.",
			},
			"addresses": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "List of the device's ip addresses.",
			},
			"user_id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the user who owns the device.",
			},
			"user_name": schema.StringAttribute{
				Computed:    true,
				Description: "The name of the user who owns the device.",
			},
			"expiry": schema.StringAttribute{
				Computed:    true,
				Description: "The expiry date of the device.",
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: "The time the device entry was created.",
			},
			"register_method": schema.StringAttribute{
				Computed:    true,
				Description: "The method used to register the device.",
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "The tags applied to the device.",
			},
			"approved_routes": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "The routes that the device is allowed to advertise.",
			},
			"available_routes": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "The routes the device is advertising.",
			},
		},
	}
}

func (d *deviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state deviceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if state.Id.ValueString() == "" && state.Name.ValueString() == "" && state.GivenName.ValueString() == "" {
		resp.Diagnostics.AddError(
			"Invalid device data source configuration",
			"At least one of `id`, `name`, or `given_name` must be provided to read a device.",
		)
		return
	}

	var device *models.V1Node
	var err error
	if state.Id.IsNull() || state.Id.IsUnknown() {
		deviceName := state.Name.ValueString()
		deviceGivenName := state.GivenName.ValueString()

		devices, err := d.client.ListDevices(ctx, nil)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to get devices",
				"An error was encountered retrieving the devices.\n"+
					err.Error(),
			)
			return
		}

		matches := 0

		for _, d := range devices {
			if d.Name == deviceName || d.GivenName == deviceGivenName {
				device = d
				matches++
			}
		}
		if matches == 0 {
			resp.Diagnostics.AddError(
				"Device not found",
				"An error was encountered retrieving the device.\n"+
					"Please check the `name` or `given_name` provided in the data source configuration.",
			)
			return
		} else if matches > 1 {
			resp.Diagnostics.AddError(
				"Multiple devices found",
				"An error was encountered retrieving the device.\n"+
					"Please check the `name` or `given_name` provided in the data source configuration.\n"+
					"Multiple devices were found with the same name or given name.",
			)
			return
		}
	} else {
		// we have an ID, so we'll try to get the device by ID
		device, err = d.client.GetDevice(ctx, state.Id.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to get device",
				"An error was encountered retrieving the device.\n"+
					err.Error(),
			)
			return
		}
	}

	dm := deviceModel{
		Id:              types.StringValue(device.ID),
		Addresses:       []types.String{},
		Name:            types.StringValue(device.Name),
		UserID:          types.StringValue(device.User.ID),
		UserName:        types.StringValue(device.User.Name),
		Expiry:          types.StringValue(device.Expiry.DeepCopy().String()),
		CreatedAt:       types.StringValue(device.CreatedAt.DeepCopy().String()),
		Tags:            []types.String{},
		GivenName:       types.StringValue(device.GivenName),
		ApprovedRoutes:  []types.String{},
		AvailableRoutes: []types.String{},
	}

	for _, add := range device.IPAddresses {
		dm.Addresses = append(dm.Addresses, types.StringValue(add))
	}

	for _, t := range device.Tags {
		dm.Tags = append(dm.Tags, types.StringValue(t))
	}

	for _, route := range device.ApprovedRoutes {
		dm.ApprovedRoutes = append(dm.ApprovedRoutes, types.StringValue(route))
	}

	for _, route := range device.AvailableRoutes {
		dm.AvailableRoutes = append(dm.AvailableRoutes, types.StringValue(route))
	}

	if device.RegisterMethod != nil {
		dm.RegisterMethod = types.StringValue(string(*device.RegisterMethod))
	} else {
		dm.RegisterMethod = types.StringValue("unknown")
	}

	diags := resp.State.Set(ctx, &dm)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
