package route

import "github.com/hashicorp/terraform-plugin-framework/types"

type deviceRouteModel struct {
	Id       types.String `tfsdk:"id"`
	DeviceId types.String `tfsdk:"device_id"`
	Routes   types.List   `tfsdk:"routes"`
}
