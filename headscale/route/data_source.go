package route

import (
	"context"
	"fmt"

	"github.com/awlsring/terraform-provider-headscale/internal/service"
	"github.com/awlsring/terraform-provider-headscale/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ datasource.DataSource              = &routeDataSource{}
	_ datasource.DataSourceWithConfigure = &routeDataSource{}
)

func DataSource() datasource.DataSource {
	return &routeDataSource{}
}

type routeDataSource struct {
	client service.Headscale
}

func (d *routeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_subnet_routes"
}

func (d *routeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(service.Headscale)
}

func (d *routeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The subnet routes data source allows you to get information on routes advertised by devices registered on the Headscale instance.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The Terraform resource ID.",
			},
			"device_id": schema.StringAttribute{
				Optional:    true,
				Description: "Filters the route list to elements belonging to the device with the provided ID.",
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
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "The status of the route.",
						},
						"device_id": schema.StringAttribute{
							Computed:    true,
							Description: "The device id the route is advertised by.",
						},
						"user_id": schema.StringAttribute{
							Computed:    true,
							Description: "The ID of the user who owns the device the route belong to.",
						},
						"created_at": schema.StringAttribute{
							Computed:    true,
							Description: "The time the route entry was created.",
						},
					},
				},
			},
		},
	}
}

func (d *routeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state dataSourceRouteModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.Id = types.StringValue(utils.CreateUUID())

	var device *string
	if state.DeviceId.ValueString() != "" {
		d := state.DeviceId.ValueString()
		device = &d
		tflog.Debug(ctx, fmt.Sprintf("Device ID: %v", *device))
	}

	var status *string
	if state.Status.ValueString() != "" {
		s := state.Status.ValueString()
		status = &s
		tflog.Debug(ctx, fmt.Sprintf("Status: %v", *status))
	}

	routes, err := d.client.ListRoutes(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to get routes",
			"An error was encountered retrieving the routes.\n"+
				err.Error(),
		)
		return
	}

	for _, route := range routes {
		if device != nil {
			if route.Node.ID != *device {
				continue
			}
		}

		stat := "disabled"
		if route.Enabled {
			stat = "enabled"
		}

		if status != nil {
			if stat != *status {
				continue
			}
		}

		r := routeModel{
			Id:        types.StringValue(route.ID),
			Route:     types.StringValue(route.Prefix),
			Status:    types.StringValue(stat),
			DeviceId:  types.StringValue(route.Node.ID),
			UserId:    types.StringValue(route.Node.User.ID),
			CreatedAt: types.StringValue(route.CreatedAt.DeepCopy().String()),
		}

		state.Routes = append(state.Routes, r)
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
