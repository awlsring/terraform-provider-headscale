package user

import (
	"context"

	"github.com/awlsring/terraform-provider-headscale/internal/service"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &userDataSource{}
	_ datasource.DataSourceWithConfigure = &userDataSource{}
)

func DataSource() datasource.DataSource {
	return &userDataSource{}
}

type userDataSource struct {
	client service.Headscale
}

func (d *userDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (d *userDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(service.Headscale)
}

func (d *userDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The user data source allows you to get information about a user registered on the Headscale instance.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the user.",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The user's id.",
			},
			"force_delete": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "If the user should be deleted even if it has nodes attached to it. Defaults to `false`.",
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: "The time the user was created.",
			},
		},
	}
}

func (d *userDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state userModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	name := state.Name.ValueString()
	user, err := d.client.GetUserByName(ctx, name)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to get user.",
			"An error was encountered retrieving the user.\n"+
				err.Error(),
		)
		return
	}

	m := userModel{
		Id:        types.StringValue(user.ID),
		Name:      types.StringValue(user.Name),
		CreatedAt: types.StringValue(user.CreatedAt.DeepCopy().String()),
	}

	diags := resp.State.Set(ctx, &m)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
