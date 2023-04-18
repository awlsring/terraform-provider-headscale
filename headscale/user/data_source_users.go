package user

import (
	"context"

	"github.com/awlsring/terraform-provider-headscale/internal/service"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &usersDataSource{}
	_ datasource.DataSourceWithConfigure = &usersDataSource{}
)

func DataSourceMultiple() datasource.DataSource {
	return &usersDataSource{}
}

type usersDataSource struct {
	client service.Headscale
}

func (d *usersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_users"
}

func (d *usersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(service.Headscale)
}

func (d *usersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The users data source allows you to get information about users registered in the Headscale instance.",
		Attributes: map[string]schema.Attribute{
			"users": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Computed:    true,
							Description: "The name of the user.",
						},
						"id": schema.StringAttribute{
							Computed:    true,
							Description: "The user's id.",
						},
						"created_at": schema.StringAttribute{
							Computed:    true,
							Description: "The time the user entry was created.",
						},
					},
				},
			},
		},
	}
}

func (d *usersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state dataSourceUsersModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	users, err := d.client.ListUsers(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to get users.",
			"An error was encountered retrieving the users.\n"+
				err.Error(),
		)
		return
	}

	for _, user := range users {
		m := userModel{
			Id:        types.StringValue(user.ID),
			Name:      types.StringValue(user.Name),
			CreatedAt: types.StringValue(user.CreatedAt.DeepCopy().String()),
		}

		state.Users = append(state.Users, m)
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
