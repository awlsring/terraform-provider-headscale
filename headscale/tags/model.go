package tags

import "github.com/hashicorp/terraform-plugin-framework/types"

type deviceTagModel struct {
	DeviceId types.String `tfsdk:"device_id"`
	Tags     types.List   `tfsdk:"tags"`
}
