package user

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type userModel struct {
	ID                types.String `tfsdk:"id"`
	Name              types.String `tfsdk:"name"`
	DisplayName       types.String `tfsdk:"display_name"`
	Email             types.String `tfsdk:"email"`
	ProfilePictureURL types.String `tfsdk:"profile_picture_url"`
	CreatedAt         types.String `tfsdk:"created_at"`
	ForceDelete       types.Bool   `tfsdk:"force_delete"`
}

type userModelList struct {
	ID                types.String `tfsdk:"id"`
	Name              types.String `tfsdk:"name"`
	DisplayName       types.String `tfsdk:"display_name"`
	Email             types.String `tfsdk:"email"`
	ProfilePictureURL types.String `tfsdk:"profile_picture_url"`
	CreatedAt         types.String `tfsdk:"created_at"`
}

type dataSourceUsersModel struct {
	ID    types.String    `tfsdk:"id"`
	Users []userModelList `tfsdk:"users"`
}
