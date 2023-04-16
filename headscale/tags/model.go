package tags

import "github.com/hashicorp/terraform-plugin-framework/types"

type deviceTagModel struct {
	Id       types.String `tfsdk:"id"`
	DeviceId types.String `tfsdk:"device_id"`
	Tags     types.List   `tfsdk:"tags"`
	// Tags           types.List   `tfsdk:"tags"`
}
