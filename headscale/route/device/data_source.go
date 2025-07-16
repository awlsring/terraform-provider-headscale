package device_route

import (
	"context"

	"github.com/awlsring/terraform-provider-headscale/headscale/route"
	"github.com/awlsring/terraform-provider-headscale/internal/service"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
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
	resp.TypeName = req.ProviderTypeName + "_device_subnet_routes"
}

func (d *deviceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(service.Headscale)
}

func (d *deviceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The device subnet routes data source allows you to get information on routes a device registered on the Headscale instance advertises.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The resolved ID of the device.",
			},
			"device_id": schema.StringAttribute{
				Required:    true,
				Description: "The device to get the routes of.",
			},
			"status": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Filters the route list to elements whose status matches what is provided. Can be `enabled` or `disabled`.",
				Validators: []validator.String{
					stringvalidator.OneOf("enabled", "disabled"),
				},
			},
			"routes": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Description: "The id of the route.",
						},
						"route": schema.StringAttribute{
							Computed:    true,
							Description: "The subnet route.",
						},
						"enabled": schema.BoolAttribute{
							Computed:    true,
							Description: "The status of the route.",
						},
					},
				},
			},
		},
	}
}

func (d *deviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state dataSourceRouteModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	device := state.DeviceId.ValueString()
	state.Id = types.StringValue(device)

	var status *string
	if state.Status.ValueString() != "" {
		s := state.Status.ValueString()
		status = &s
	}

	routes, err := d.client.ListDeviceRoutes(ctx, device)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to get routes",
			"An error was encountered retrieving the routes.\n"+
				err.Error(),
		)
		return
	}

	for _, r := range routes {
		stat := route.ParseStatusFromModel(r)

		if status != nil {
			if stat != *status {
				continue
			}
		}

		r := routeModel{
			Id:      types.StringValue(r.ID),
			Route:   types.StringValue(r.Prefix),
			Enabled: types.BoolValue(r.Enabled),
		}

		state.Routes = append(state.Routes, r)
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
