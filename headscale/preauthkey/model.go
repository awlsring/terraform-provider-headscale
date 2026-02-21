package preauthkey

import "github.com/hashicorp/terraform-plugin-framework/types"

type apikeyListModel struct {
	Id          types.String      `tfsdk:"id"`
	User        types.String      `tfsdk:"user"`
	All         types.Bool        `tfsdk:"all"`
	PreAuthKeys []preAuthKeyModel `tfsdk:"pre_auth_keys"`
}

type preAuthKeyModel struct {
	Id         types.String `tfsdk:"id"`
	User       types.String `tfsdk:"user"`
	Key        types.String `tfsdk:"key"`
	Reusable   types.Bool   `tfsdk:"reusable"`
	Ephemeral  types.Bool   `tfsdk:"ephemeral"`
	Used       types.Bool   `tfsdk:"used"`
	Expired    types.Bool   `tfsdk:"expired"`
	Expiration types.String `tfsdk:"expiration"`
	CreatedAt  types.String `tfsdk:"created_at"`
	ACLTags    types.List   `tfsdk:"acl_tags"`
}

type preAuthKeyResourceModel struct {
	TimeToExpire types.String `tfsdk:"time_to_expire"`
	Id           types.String `tfsdk:"id"`
	User         types.String `tfsdk:"user"`
	Key          types.String `tfsdk:"key"`
	Reusable     types.Bool   `tfsdk:"reusable"`
	Ephemeral    types.Bool   `tfsdk:"ephemeral"`
	Used         types.Bool   `tfsdk:"used"`
	Expired      types.Bool   `tfsdk:"expired"`
	Expiration   types.String `tfsdk:"expiration"`
	CreatedAt    types.String `tfsdk:"created_at"`
	ACLTags      types.Set    `tfsdk:"acl_tags"`
}
