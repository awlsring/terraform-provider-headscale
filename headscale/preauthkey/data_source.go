package preauthkey

import (
	"context"

	"github.com/awlsring/terraform-provider-headscale/internal/service"
	"github.com/awlsring/terraform-provider-headscale/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &preAuthKeyDataSource{}
	_ datasource.DataSourceWithConfigure = &preAuthKeyDataSource{}
)

func DataSource() datasource.DataSource {
	return &preAuthKeyDataSource{}
}

type preAuthKeyDataSource struct {
	client service.Headscale
}

func (d *preAuthKeyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pre_auth_keys"
}

func (d *preAuthKeyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(service.Headscale)
}

func (d *preAuthKeyDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The pre auth key data source allows you to list all pre auth keys that belong to a specified user on the Headscale instance.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the Terraform resource.",
			},
			"user": schema.BoolAttribute{
				Required:            true,
				MarkdownDescription: "The user to get pre auth keys for.",
			},
			"all": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "If expired keys should be included in the list. Defaults to `false`.",
			},
			"pre_auth_keys": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Description: "The pre auth key's id'.",
						},
						"user": schema.StringAttribute{
							Computed:    true,
							Description: "The user who owns this key.",
						},
						"key": schema.StringAttribute{
							Computed:    true,
							Description: "The generated key.",
							Sensitive:   true,
						},
						"reusable": schema.BoolAttribute{
							Computed:    true,
							Description: "If the key is reusable.",
						},
						"ephemeral": schema.BoolAttribute{
							Computed:    true,
							Description: "If the key is ephemeral.",
						},
						"used": schema.BoolAttribute{
							Computed:    true,
							Description: "If the key has been used.",
						},
						"expired": schema.BoolAttribute{
							Computed:    true,
							Description: "If the key is expired.",
						},
						"expiration": schema.StringAttribute{
							Computed:    true,
							Description: "The time the pre auth key expires at.",
						},
						"created_at": schema.StringAttribute{
							Computed:    true,
							Description: "The time the key entry was created.",
						},
						"acl_tags": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							Description: "ACL tags on the pre auth key.",
						},
					},
				},
			},
		},
	}
}

func (d *preAuthKeyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state apikeyListModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.Id = types.StringValue(utils.CreateUUID())
	user := state.User.ValueString()
	returnAll := state.All.ValueBool()

	preAuthKeys, err := d.client.ListPreAuthKeys(ctx, user)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to get pre auth keys",
			"An error was encountered retrieving the pre auth key list.\n"+
				err.Error(),
		)
		return
	}

	for _, key := range preAuthKeys {
		expireString := key.Expiration.DeepCopy().String()
		isExpired, err := utils.IsExpired(expireString)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to parse expiration date",
				"An error was encountered parsing the expiration date.\n"+
					err.Error(),
			)
		}

		if !returnAll && isExpired {
			continue
		}

		m := preAuthKeyModel{
			Id:         types.StringValue(key.ID),
			User:       types.StringValue(key.User),
			Key:        types.StringValue(key.Key),
			Reusable:   types.BoolValue(key.Reusable),
			Ephemeral:  types.BoolValue(key.Ephemeral),
			Used:       types.BoolValue(key.Used),
			Expiration: types.StringValue(key.Expiration.DeepCopy().String()),
			Expired:    types.BoolValue(isExpired),
			CreatedAt:  types.StringValue(key.CreatedAt.DeepCopy().String()),
		}

		tags, diags := types.ListValueFrom(ctx, types.StringType, key.ACLTags)
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}
		m.ACLTags = tags

		state.PreAuthKeys = append(state.PreAuthKeys, m)
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
