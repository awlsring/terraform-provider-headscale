package policy

import (
	"context"
	"strings"

	"github.com/awlsring/terraform-provider-headscale/internal/service"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	defaultMininalPolicy = `"{"acls":[{"action":"accept"}]}"`
)

var (
	_ datasource.DataSource              = &policyDataSource{}
	_ datasource.DataSourceWithConfigure = &policyDataSource{}
)

func DataSource() datasource.DataSource {
	return &policyDataSource{}
}

type policyDataSource struct {
	client service.Headscale
}

func (d *policyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy"
}

func (d *policyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(service.Headscale)
}

func (d *policyDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The policy data source allows reading the ACL policy used for the tailnet.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The Terraform Id of the resource.",
			},
			"policy": schema.StringAttribute{
				Computed:    true,
				Description: "The policy data in HuJSON format. See https://tailscale.com/kb/1337/policy-syntax.",
			},
			"updated": schema.StringAttribute{
				Computed:    true,
				Description: "The time when the policy was last updated.",
			},
		},
	}
}

func (d *policyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state policyModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	policy, err := d.client.GetPolicy(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "acl policy not found") {
			policy = &service.Policy{}
		} else {
			resp.Diagnostics.AddError(
				"Unable to get policy.",
				"An error was encountered retrieving the policy.\n"+
					err.Error(),
			)
			return
		}
	}

	var policyId types.String
	if state.ID.IsUnknown() || state.ID.IsNull() {
		uuid := uuid.New().String()
		policyId = types.StringValue(uuid)
	} else {
		policyId = state.ID
	}

	m := policyModel{
		ID:      policyId,
		Policy:  types.StringValue(policy.Policy),
		Updated: types.StringValue(policy.Updated.DeepCopy().String()),
	}

	diags := resp.State.Set(ctx, &m)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
