package device_route

import (
	"context"
	"fmt"
	"regexp"
	"slices"
	"time"

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
						stringvalidator.RegexMatches(regexp.MustCompile(`^(((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\/([0-9]|[1-2][0-9]|3[0-2])|([0-9a-fA-F]{0,4}:){2,7}[0-9a-fA-F]{0,4}\/([0-9]|[1-9][0-9]|1[0-1][0-9]|12[0-8]))$`), "route must be a valid IPv4 CIDR (e.g., 10.0.10.0/24, 0.0.0.0/0) or IPv6 CIDR (e.g., 2001:db8::/32, ::/0)"),
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

	routes, err := r.enableRoutes(ctx, plan.DeviceId.ValueString(), plan.Routes)
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

	routes, err := r.enableRoutes(ctx, plan.DeviceId.ValueString(), plan.Routes)
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

func (r *deviceRoutesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state deviceRouteModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	deviceId := state.DeviceId.ValueString()

	deviceRoutes, err := r.readDeviceRoutes(ctx, deviceId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to get device routes",
			"An error was encountered retrieving the device.\n"+
				err.Error(),
		)
		return
	}

	originalRoutes := tfStringListToStringList(state.Routes)
	enabledRoutesStrList := tfStringListToStringList(deviceRoutes.Routes)
	sanitizedRoutes := normalizeRoutesForState(originalRoutes, enabledRoutesStrList)
	deviceRoutes.Routes, err = stringListToTFStringList(ctx, sanitizedRoutes)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to get device routes",
			"An error was encountered when handling routes.\n"+
				err.Error(),
		)
		return
	}

	diags := resp.State.Set(ctx, deviceRoutes)
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

func stringListToTFStringList(ctx context.Context, list []string) (types.List, error) {
	c, diags := types.ListValueFrom(ctx, types.StringType, list)
	if diags.HasError() {
		return types.List{}, fmt.Errorf("error creating list of routes: %s", diags.Errors())
	}
	return c, nil
}

func tfStringListToStringList(list types.List) []string {
	if list.IsNull() || list.IsUnknown() {
		return []string{}
	}
	strList := []string{}
	for _, s := range list.Elements() {
		conv := s.(types.String)
		strList = append(strList, conv.ValueString())
	}

	return strList
}

func (r *deviceRoutesResource) enableRoutes(ctx context.Context, deviceId string, routes types.List) (*deviceRouteModel, error) {
	routesRequested := tfStringListToStringList(routes)

	if err := r.client.EnableDeviceRoutes(ctx, deviceId, routesRequested); err != nil {
		return nil, fmt.Errorf("error enabling routes on device %s: %w", deviceId, err)
	}

	expectedRoutes := normalizeRoutesForState(routesRequested, routesRequested)
	deadline := time.Now().Add(10 * time.Second)

	var lastRead *deviceRouteModel
	for {
		enabledRoutesModel, err := r.readDeviceRoutes(ctx, deviceId)
		if err != nil {
			return nil, fmt.Errorf("error reading enabled routes on device %s: %w", deviceId, err)
		}

		// normalize routes and keep requested order for deterministic list state.
		enabledRoutesStrList := tfStringListToStringList(enabledRoutesModel.Routes)
		sanitizedRoutes := normalizeRoutesForState(routesRequested, enabledRoutesStrList)
		enabledRoutesModel.Routes, err = stringListToTFStringList(ctx, sanitizedRoutes)
		if err != nil {
			return nil, fmt.Errorf("error converting enabled routes to TF string list: %w", err)
		}
		lastRead = enabledRoutesModel

		if slices.Equal(sanitizedRoutes, expectedRoutes) || time.Now().After(deadline) {
			return lastRead, nil
		}

		time.Sleep(250 * time.Millisecond)
	}
}

// Keep state deterministic by preserving requested order while filtering out implicit exit-node pair routes.
func normalizeRoutesForState(specifiedRoutes []string, detectedRoutes []string) []string {
	ipv4ExitNode := "0.0.0.0/0"
	ipv6ExitNode := "::/0"

	hasIPv4Specified := slices.Contains(specifiedRoutes, ipv4ExitNode)
	hasIPv6Specified := slices.Contains(specifiedRoutes, ipv6ExitNode)

	routeDetected := make(map[string]struct{}, len(detectedRoutes))
	for _, route := range detectedRoutes {
		if route == ipv4ExitNode && !hasIPv4Specified {
			continue
		}
		if route == ipv6ExitNode && !hasIPv6Specified {
			continue
		}
		routeDetected[route] = struct{}{}
	}

	// Return in user-specified order to avoid list-order drift when API ordering differs.
	filteredRoutes := make([]string, 0, len(specifiedRoutes))
	for _, route := range specifiedRoutes {
		if _, ok := routeDetected[route]; ok {
			filteredRoutes = append(filteredRoutes, route)
		}
	}

	return filteredRoutes
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
