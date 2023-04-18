package device

import (
	"context"
	"fmt"
	"strings"

	"github.com/awlsring/terraform-provider-headscale/internal/service"
	"github.com/awlsring/terraform-provider-headscale/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ datasource.DataSource              = &devicesDataSource{}
	_ datasource.DataSourceWithConfigure = &devicesDataSource{}
)

func DataSourceMultiple() datasource.DataSource {
	return &devicesDataSource{}
}

type devicesDataSource struct {
	client service.Headscale
}

func (d *devicesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices"
}

func (d *devicesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(service.Headscale)
}

func (d *devicesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the Terraform resource.",
			},
			"user": schema.StringAttribute{
				Optional:    true,
				Description: "Filters the device list to elements belonging to the user with the provided ID.",
			},
			"name_prefix": schema.StringAttribute{
				Optional:    true,
				Description: "Filters the device list to elements whose name has the provided prefix.",
			},
			"devices": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Description: "The id of the device",
						},
						"addresses": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							Description: "List of the device's ip addresses.",
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: "The device's name.",
						},
						"user": schema.StringAttribute{
							Computed:    true,
							Description: "The ID of the user who owns the device.",
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
							Optional:    true,
							Description: "The method used to register the device.",
						},
						"tags": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							Description: "The tags applied to the device.",
						},
						"given_name": schema.StringAttribute{
							Computed:    true,
							Description: "The device's given name.",
						},
					},
				},
			},
		},
	}
}

func (d *devicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state deviceListModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.Id = types.StringValue(utils.CreateUUID())

	var namePrefix *string
	if state.NamePrefix.ValueString() != "" {
		n := state.NamePrefix.ValueString()
		namePrefix = &n
		tflog.Debug(ctx, fmt.Sprintf("namePrefix: %v", *namePrefix))
	}

	var user *string
	if state.User.ValueString() != "" {
		u := state.User.ValueString()
		user = &u
		tflog.Debug(ctx, fmt.Sprintf("user: %v", *user))
	}

	deviceList, err := d.client.ListDevices(ctx, user)
	tflog.Debug(ctx, fmt.Sprintf("Devices: %v", len(deviceList)))
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to list devices",
			"An error was encountered retrieving devices.\n"+
				err.Error(),
		)
		return
	}

	for _, device := range deviceList {
		tflog.Debug(ctx, fmt.Sprintf("Device: %s", device.Name))

		if namePrefix != nil {
			if !strings.HasPrefix(device.Name, *namePrefix) {
				continue
			}
		}

		dm := deviceModel{
			Id:        types.StringValue(device.ID),
			Addresses: []types.String{},
			Name:      types.StringValue(device.Name),
			User:      types.StringValue(device.User.ID),
			Expiry:    types.StringValue(device.Expiry.DeepCopy().String()),
			CreatedAt: types.StringValue(device.CreatedAt.DeepCopy().String()),
			Tags:      []types.String{},
			GivenName: types.StringValue(device.GivenName),
		}

		for _, add := range device.IPAddresses {
			dm.Addresses = append(dm.Addresses, types.StringValue(add))
		}

		for _, t := range device.ValidTags {
			dm.Tags = append(dm.Tags, types.StringValue(t))
		}

		for _, t := range device.ForcedTags {
			dm.Tags = append(dm.Tags, types.StringValue(t))
		}

		if device.RegisterMethod != nil {
			dm.RegisterMethod = types.StringValue(string(*device.RegisterMethod))
		}

		state.Devices = append(state.Devices, dm)
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
