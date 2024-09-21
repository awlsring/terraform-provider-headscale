package user

import (
	"context"

	"github.com/awlsring/terraform-provider-headscale/internal/gen/models"
	"github.com/awlsring/terraform-provider-headscale/internal/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &userResource{}
	_ resource.ResourceWithConfigure   = &userResource{}
	_ resource.ResourceWithImportState = &userResource{}
)

func Resource() resource.Resource {
	return &userResource{}
}

type userResource struct {
	client service.Headscale
}

func (d *userResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (d *userResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(service.Headscale)
}

func (d *userResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The user resource allows you to register a user on the Headscale instance.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the user.",
			},
			"force_delete": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "If the user should be deleted even if it has nodes attached to it. Defaults to `false`.",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The user's id.",
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: "The time the user was created.",
			},
		},
	}
}

func (r *userResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan userModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	userName := plan.Name.ValueString()

	user, err := r.client.CreateUser(ctx, userName)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating user",
			"Could not create user, unexpected error: "+err.Error(),
		)
		return
	}

	m := userModel{
		Id:        types.StringValue(user.ID),
		Name:      types.StringValue(user.Name),
		CreatedAt: types.StringValue(user.CreatedAt.DeepCopy().String()),
	}

	diags = resp.State.Set(ctx, &m)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *userResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan userModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state userModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	oldName := state.Name.ValueString()
	newName := plan.Name.ValueString()

	var user *models.V1User
	var err error
	if oldName != newName {
		user, err = r.client.RenameUser(ctx, oldName, newName)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error updating user",
				"Could not update user, unexpected error: "+err.Error(),
			)
			return
		}
	} else {
		user, err = r.client.GetUser(ctx, oldName)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to get user.",
				"An error was encountered retrieving the user.\n"+
					err.Error(),
			)
			return
		}
	}

	m := userModel{
		Id:          types.StringValue(user.ID),
		Name:        types.StringValue(user.Name),
		ForceDelete: types.BoolValue(plan.ForceDelete.ValueBool()),
		CreatedAt:   types.StringValue(user.CreatedAt.DeepCopy().String()),
	}

	diags = resp.State.Set(ctx, &m)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *userResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state userModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	user := state.Name.ValueString()
	devices, err := r.client.ListDevices(ctx, &user)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting user",
			"Could not list devices, unexpected error: "+err.Error(),
		)
		return
	}

	if len(devices) > 0 && !state.ForceDelete.ValueBool() {
		resp.Diagnostics.AddError(
			"Error deleting user",
			"User has devices attached to it. Set `force_delete` to `true` to delete the user anyway.",
		)
		return
	}

	for _, device := range devices {
		err = r.client.DeleteDevice(ctx, device.ID)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error deleting user",
				"Could not remove device, unexpected error: "+err.Error(),
			)
			return
		}
	}

	err = r.client.DeleteUser(ctx, state.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting user",
			"Could not remove user, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *userResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state userModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	name := state.Name.ValueString()
	user, err := r.client.GetUser(ctx, name)
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

	diags = resp.State.Set(ctx, &m)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *userResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	user, err := r.client.GetUser(ctx, req.ID)
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
