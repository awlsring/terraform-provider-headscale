package policy

import "github.com/hashicorp/terraform-plugin-framework/types"

type policyModel struct {
	ID      types.String `tfsdk:"id"`
	Policy  types.String `tfsdk:"policy"`
	Updated types.String `tfsdk:"updated"`
}
