package route

import "github.com/hashicorp/terraform-plugin-framework/types"

type dataSourceRouteModel struct {
	Id       types.String `tfsdk:"id"`
	DeviceId types.String `tfsdk:"device_id"`
	Status   types.String `tfsdk:"status"`
	Routes   []routeModel `tfsdk:"routes"`
}

type routeModel struct {
	Id        types.String `tfsdk:"id"`
	Route     types.String `tfsdk:"route"`
	Status    types.String `tfsdk:"status"`
	DeviceId  types.String `tfsdk:"device_id"`
	UserId    types.String `tfsdk:"user_id"`
	CreatedAt types.String `tfsdk:"created_at"`
}
