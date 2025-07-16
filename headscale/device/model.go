package device

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type deviceListModel struct {
	Id         types.String  `tfsdk:"id"`
	UserName   types.String  `tfsdk:"user_name"`
	NamePrefix types.String  `tfsdk:"name_prefix"`
	Devices    []deviceModel `tfsdk:"devices"`
}

type deviceModel struct {
	Id              types.String   `tfsdk:"id"`
	Addresses       []types.String `tfsdk:"addresses"`
	Name            types.String   `tfsdk:"name"`
	UserID          types.String   `tfsdk:"user_id"`
	UserName        types.String   `tfsdk:"user_name"`
	Expiry          types.String   `tfsdk:"expiry"`
	CreatedAt       types.String   `tfsdk:"created_at"`
	RegisterMethod  types.String   `tfsdk:"register_method"`
	GivenName       types.String   `tfsdk:"given_name"`
	Tags            []types.String `tfsdk:"tags"`
	ApprovedRoutes  []types.String `tfsdk:"approved_routes"`
	AvailableRoutes []types.String `tfsdk:"available_routes"`
}
