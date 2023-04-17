package device_route

import "github.com/hashicorp/terraform-plugin-framework/types"

type deviceRouteModel struct {
	DeviceId types.String `tfsdk:"device_id"`
	Routes   types.List   `tfsdk:"routes"`
}

type dataSourceRouteModel struct {
	DeviceId types.String `tfsdk:"device_id"`
	Status   types.String `tfsdk:"status"`
	Routes   []routeModel `tfsdk:"routes"`
}

type routeModel struct {
	Id      types.String `tfsdk:"id"`
	Route   types.String `tfsdk:"route"`
	Enabled types.Bool   `tfsdk:"enabled"`
}
