package user

import "github.com/hashicorp/terraform-plugin-framework/types"

type userModel struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	CreatedAt types.String `tfsdk:"created_at"`
}

type dataSourceUsersModel struct {
	Users []userModel `tfsdk:"users"`
}
