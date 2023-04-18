package apikey

import (
	"context"
	"time"

	"github.com/awlsring/terraform-provider-headscale/internal/service"
	"github.com/awlsring/terraform-provider-headscale/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &apiKeyDataSource{}
	_ datasource.DataSourceWithConfigure = &apiKeyDataSource{}
)

func DataSource() datasource.DataSource {
	return &apiKeyDataSource{}
}

type apiKeyDataSource struct {
	client service.Headscale
}

func (d *apiKeyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_keys"
}

func (d *apiKeyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(service.Headscale)
}

func (d *apiKeyDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The API key data source allows you to list all API keys on Headscale instance. This will only return the API key metadata and not the actual key.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the Terraform resource.",
			},
			"all": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "If expired keys should be included in the list. Defaults to `false`.",
			},
			"api_keys": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Description: "The id of the api key.",
						},
						"prefix": schema.StringAttribute{
							Computed:    true,
							Description: "The api key's prefix.",
						},
						"expiration": schema.StringAttribute{
							Computed:    true,
							Description: "Expiration date of the api key.",
						},
						"expired": schema.BoolAttribute{
							Computed:    true,
							Description: "If the api key is expired.",
						},
						"created_at": schema.StringAttribute{
							Computed:    true,
							Description: "The time the device entry was created.",
						},
					},
				},
			},
		},
	}
}

func (d *apiKeyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state apikeyListModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.Id = types.StringValue(utils.CreateUUID())

	returnAll := state.All.ValueBool()

	apiKeys, err := d.client.ListAPIKeys(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to get api keys",
			"An error was encountered retrieving the api key list.\n"+
				err.Error(),
		)
		return
	}

	for _, key := range apiKeys {
		expireString := key.Expiration.DeepCopy().String()
		isExpired, err := isExpired(expireString)
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

		m := apikeyModel{
			Id:         types.StringValue(key.ID),
			Prefix:     types.StringValue(key.Prefix),
			Expiration: types.StringValue(key.Expiration.DeepCopy().String()),
			Expired:    types.BoolValue(isExpired),
			CreatedAt:  types.StringValue(key.CreatedAt.DeepCopy().String()),
		}

		state.ApiKeys = append(state.ApiKeys, m)
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func isExpired(t string) (bool, error) {
	if t == "0001-01-01T00:00:00.000Z" {
		return false, nil
	}

	expireTime, err := time.Parse(time.RFC3339Nano, t)
	if err != nil {
		return false, err
	}

	if expireTime.Before(time.Now()) {
		return true, nil
	}

	return false, nil
}
