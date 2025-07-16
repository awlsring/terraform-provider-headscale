package policy

import (
	"context"

	"github.com/awlsring/terraform-provider-headscale/internal/service"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &policyResource{}
	_ resource.ResourceWithConfigure = &policyResource{}
)

func Resource() resource.Resource {
	return &policyResource{}
}

type policyResource struct {
	client service.Headscale
}

func (d *policyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy"
}

func (d *policyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(service.Headscale)
}

func (d *policyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The policy resource allows definition of the ACL policy used for the tailnet. To use this resource, the Headscale instance must have the `policy.mode` setting set to `database`.\n\nTo read more about ACL policies, see the [Tailscale documentation](https://tailscale.com/kb/1018/acls).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The Terraform Id of the resource.",
			},
			"policy": schema.StringAttribute{
				Required:    true,
				Description: "The policy data in HuJSON format. See https://tailscale.com/kb/1337/policy-syntax.",
			},
			"updated": schema.StringAttribute{
				Computed:    true,
				Description: "The time when the policy was last updated.",
			},
		},
	}
}

func (r *policyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan policyModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	policyData := plan.Policy.ValueString()

	policy, err := r.client.SetPolicy(ctx, policyData)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error setting policy",
			"Could not set policy, unexpected error: "+err.Error(),
		)
		return
	}

	policyId := uuid.New()

	m := policyModel{
		ID:      types.StringValue(policyId.String()),
		Policy:  types.StringValue(policyData),
		Updated: types.StringValue(policy.Updated.DeepCopy().String()),
	}

	diags = resp.State.Set(ctx, &m)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *policyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan policyModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state policyModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	policyData := plan.Policy.ValueString()

	policy, err := r.client.SetPolicy(ctx, policyData)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error setting policy",
			"Could not set policy, unexpected error: "+err.Error(),
		)
		return
	}

	m := policyModel{
		ID:      state.ID,
		Policy:  types.StringValue(policyData),
		Updated: types.StringValue(policy.Updated.DeepCopy().String()),
	}

	diags = resp.State.Set(ctx, &m)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *policyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state policyModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.SetPolicy(ctx, "")
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting policy",
			"Could not delete policy, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *policyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state policyModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	policy, err := r.client.GetPolicy(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading policy",
			"Could not read policy, unexpected error: "+err.Error(),
		)
		return
	}

	m := policyModel{
		ID:      state.ID,
		Policy:  types.StringValue(policy.Policy),
		Updated: types.StringValue(policy.Updated.DeepCopy().String()),
	}

	diags := resp.State.Set(ctx, m)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
