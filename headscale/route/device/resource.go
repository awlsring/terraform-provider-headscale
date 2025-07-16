package device_route

import (
	"context"
	"fmt"
	"regexp"

	"github.com/awlsring/terraform-provider-headscale/internal/service"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &deviceRoutesResource{}
	_ resource.ResourceWithConfigure   = &deviceRoutesResource{}
	_ resource.ResourceWithImportState = &deviceRoutesResource{}
)

func Resource() resource.Resource {
	return &deviceRoutesResource{}
}

type deviceRoutesResource struct {
	client service.Headscale
}

func (d *deviceRoutesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_subnet_routes"
}

func (d *deviceRoutesResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(service.Headscale)
}

func (d *deviceRoutesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The device subnet routes resource allows enabling routes advertised by a device registered on the Headscale instance. Utilizing this resource will reset any previous configuration for routes advertised by the device. If a route was previously enabled but is not present in the list of routes, it will be disabled.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The Terraform Id of the resource.",
			},
			"device_id": schema.StringAttribute{
				Required:    true,
				Description: "The id of the device to get subnet routes from.",
			},
			"routes": schema.ListAttribute{
				Required:    true,
				ElementType: types.StringType,
				Description: "The enabled routes of the device.",
				Validators: []validator.List{
					listvalidator.UniqueValues(),
					listvalidator.ValueStringsAre(
						stringvalidator.RegexMatches(regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\/([0-9]|[1-2][0-9]|3[0-2])$`), "tag must follow scheme like `10.0.10.0/24`"),
					),
				},
			},
		},
	}
}

func (r *deviceRoutesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan deviceRouteModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	routes, err := r.enableRoutes(ctx, &plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error enabling routes on device",
			"Could not enable routes, unexpected error: "+err.Error(),
		)
		return
	}

	diags = resp.State.Set(ctx, routes)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *deviceRoutesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan deviceRouteModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	routes, err := r.enableRoutes(ctx, &plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating routes on device",
			"Could not update routes, unexpected error: "+err.Error(),
		)
		return
	}

	diags = resp.State.Set(ctx, routes)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *deviceRoutesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning(
		"Deleting device routes!",
		"Deleting a device route resource will disable all routes on the node.",
	)

	var state deviceRouteModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DisableDeviceRoutes(ctx, state.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting device routes",
			"Could not remove routes, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *deviceRoutesResource) enableRoutes(ctx context.Context, m *deviceRouteModel) (*deviceRouteModel, error) {
	deviceId := m.DeviceId.ValueString()
	routesRequested := []string{}
	for _, r := range m.Routes.Elements() {
		conv := r.(types.String)
		routesRequested = append(routesRequested, conv.ValueString())
	}

	if err := r.client.EnableDeviceRoutes(ctx, deviceId, routesRequested); err != nil {
		return nil, fmt.Errorf("error enabling routes on device %s: %w", deviceId, err)
	}

	// routes, diags := types.ListValueFrom(ctx, types.StringType, routesRequested)
	// if diags.HasError() {
	// 	return nil, fmt.Errorf("error creating list of routes")
	// }

	return r.readDeviceRoutes(ctx, deviceId)
}

func (r *deviceRoutesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state deviceRouteModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	deviceId := state.DeviceId.ValueString()

	device, err := r.readDeviceRoutes(ctx, deviceId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to get device routes",
			"An error was encountered retrieving the device.\n"+
				err.Error(),
		)
		return
	}

	diags := resp.State.Set(ctx, device)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *deviceRoutesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	device, err := r.readDeviceRoutes(ctx, req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to get device routes",
			"An error was encountered retrieving the device.\n"+
				err.Error(),
		)
		return
	}

	diags := resp.State.Set(ctx, device)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *deviceRoutesResource) readDeviceRoutes(ctx context.Context, id string) (*deviceRouteModel, error) {
	routes, err := r.client.ListDeviceRoutes(ctx, id)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Error reading routes for device %s: %s", id, err.Error()))
		return nil, err
	}

	dm := deviceRouteModel{
		DeviceId: types.StringValue(id),
		Id:       types.StringValue(id),
	}

	enabledRoutes := []string{}
	for _, route := range routes {
		if route.Enabled {
			enabledRoutes = append(enabledRoutes, route.Prefix)
		}
	}

	c, diags := types.ListValueFrom(ctx, types.StringType, enabledRoutes)
	if diags.HasError() {
		return nil, fmt.Errorf("error creating list of routes")
	}

	dm.Routes = c

	return &dm, nil
}
