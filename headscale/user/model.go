package user

import "github.com/hashicorp/terraform-plugin-framework/types"

type userModel struct {
	Id          types.String `tfsdk:"id"`
	ForceDelete types.Bool   `tfsdk:"force_delete"`
	Name        types.String `tfsdk:"name"`
	CreatedAt   types.String `tfsdk:"created_at"`
}

type dataSourceUsersModel struct {
	Id    types.String `tfsdk:"id"`
	Users []userModel  `tfsdk:"users"`
}
