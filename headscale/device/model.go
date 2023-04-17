package device

import "github.com/hashicorp/terraform-plugin-framework/types"

type deviceListModel struct {
	Id         types.String  `tfsdk:"id"`
	User       types.String  `tfsdk:"user"`
	NamePrefix types.String  `tfsdk:"name_prefix"`
	Devices    []deviceModel `tfsdk:"devices"`
}

type deviceModel struct {
	Id             types.String   `tfsdk:"id"`
	Addresses      []types.String `tfsdk:"addresses"`
	Name           types.String   `tfsdk:"name"`
	User           types.String   `tfsdk:"user"`
	Expiry         types.String   `tfsdk:"expiry"`
	CreatedAt      types.String   `tfsdk:"created_at"`
	RegisterMethod types.String   `tfsdk:"register_method"`
	GivenName      types.String   `tfsdk:"given_name"`
	Tags           []types.String `tfsdk:"tags"`
}
