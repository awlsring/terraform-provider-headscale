package device

import (
	"context"

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
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:    true,
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
	}
}

func (d *deviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state deviceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	deviceId := state.Id.ValueString()

	device, err := d.client.GetDevice(ctx, deviceId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to get devices",
			"An error was encountered retrieving the device.\n"+
				err.Error(),
		)
		return
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

	state = dm

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
