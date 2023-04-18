package apikey

import "github.com/hashicorp/terraform-plugin-framework/types"

type apikeyListModel struct {
	Id      types.String  `tfsdk:"id"`
	All     types.Bool    `tfsdk:"all"`
	ApiKeys []apikeyModel `tfsdk:"api_keys"`
}

type apikeyModel struct {
	Id         types.String `tfsdk:"id"`
	Prefix     types.String `tfsdk:"prefix"`
	Expiration types.String `tfsdk:"expiration"`
	Expired    types.Bool   `tfsdk:"expired"`
	CreatedAt  types.String `tfsdk:"created_at"`
}

type apikeyResourceModel struct {
	DaysToExpire types.Int64  `tfsdk:"days_to_expire"`
	Id           types.String `tfsdk:"id"`
	Key          types.String `tfsdk:"key"`
	Prefix       types.String `tfsdk:"prefix"`
	Expiration   types.String `tfsdk:"expiration"`
	Expired      types.Bool   `tfsdk:"expired"`
	CreatedAt    types.String `tfsdk:"created_at"`
}
