package route

import (
	"github.com/awlsring/terraform-provider-headscale/internal/service"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	RouteStatusDisabled = "disabled"
	RouteStatusEnabled  = "enabled"
)

type dataSourceRouteModel struct {
	Id       types.String `tfsdk:"id"`
	DeviceId types.String `tfsdk:"device_id"`
	Status   types.String `tfsdk:"status"`
	Routes   []routeModel `tfsdk:"routes"`
}

func ParseStatusFromModel(m *service.Route) string {
	if m.Enabled {
		return RouteStatusEnabled
	}
	return RouteStatusDisabled
}

type routeModel struct {
	Id       types.String `tfsdk:"id"`
	Route    types.String `tfsdk:"route"`
	Status   types.String `tfsdk:"status"`
	DeviceId types.String `tfsdk:"device_id"`
	UserId   types.String `tfsdk:"user_id"`
	Enabled  types.Bool   `tfsdk:"enabled"`
}
