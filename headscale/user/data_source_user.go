package user

import (
	"context"

	"github.com/awlsring/terraform-provider-headscale/internal/service"
	"github.com/awlsring/terraform-provider-headscale/internal/utils"
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
				Optional:    true,
				Computed:    true,
				Description: "The name of the user.",
			},
			"display_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The display name of the user.",
			},
			"id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The user's id.",
			},
			"email": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The user's email address.",
			},
			"profile_picture_url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The URL of the user's profile picture.",
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

	user, err := d.client.GetUser(ctx, service.GetUserInput{
		Name:  utils.StrToPtr(state.Name.ValueString()),
		ID:    utils.StrToPtr(state.ID.ValueString()),
		Email: utils.StrToPtr(state.Email.ValueString()),
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to get user.",
			"An error was encountered retrieving the user.\n"+
				err.Error(),
		)
		return
	}

	m := userModel{
		ID:        types.StringValue(user.ID),
		Name:      types.StringValue(user.Name),
		CreatedAt: types.StringValue(user.CreatedAt.DeepCopy().String()),
	}

	if user.DisplayName != "" {
		m.DisplayName = types.StringValue(user.DisplayName)
	} else {
		m.DisplayName = types.StringNull()
	}
	if user.Email != "" {
		m.Email = types.StringValue(user.Email)
	} else {
		m.Email = types.StringNull()
	}
	if user.ProfilePicURL != "" {
		m.ProfilePictureURL = types.StringValue(user.ProfilePicURL)
	} else {
		m.ProfilePictureURL = types.StringNull()
	}

	diags := resp.State.Set(ctx, &m)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
