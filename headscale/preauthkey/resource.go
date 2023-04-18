package preauthkey

import (
	"context"
	"fmt"
	"regexp"

	"github.com/awlsring/terraform-provider-headscale/internal/service"
	"github.com/awlsring/terraform-provider-headscale/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &preAuthKeyResource{}
	_ resource.ResourceWithConfigure = &preAuthKeyResource{}
)

func Resource() resource.Resource {
	return &preAuthKeyResource{}
}

type preAuthKeyResource struct {
	client service.Headscale
}

func (d *preAuthKeyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pre_auth_key"
}

func (d *preAuthKeyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(service.Headscale)
}

func (d *preAuthKeyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The pre auth key resource allows you to create a pre auth key that can be used to register a new device on the Headscale instance. By default keys that are created with this resource will be not reusable, not ephemeral, and expire in 1 hour. Keys cannot be modified, so any change to the input on this resource will cause the key to be expired and a new key to be created.",
		Attributes: map[string]schema.Attribute{
			"time_to_expire": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The time until the key expires. This is a string in the format of `30m`, `3h`, `90d`, etc. Defaults to `1h`",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^-?\d+(\.\d+)?(ns|us|µs|ms|s|m|h|d|w|M|y)$`), "must be a valid duration string. (e.g. 30m, 3h, 90d, 4M, 1y, etc)"),
				},
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The id of the pre auth key.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"user": schema.StringAttribute{
				Required:    true,
				Description: "The user that owns the pre auth key.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"key": schema.StringAttribute{
				Computed:    true,
				Description: "The pre auth key.",
				Sensitive:   true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"reusable": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: "If the key is reusable.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
					boolplanmodifier.RequiresReplace(),
				},
			},
			"ephemeral": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: "If the key is ephemeral.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
					boolplanmodifier.RequiresReplace(),
				},
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
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: "The time the key was created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"acl_tags": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: "ACL tags on the pre auth key.",
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Validators: []validator.List{
					listvalidator.UniqueValues(),
					listvalidator.ValueStringsAre(
						stringvalidator.RegexMatches(regexp.MustCompile("^tag:[\\w\\d]+$"), "tag must follow scheme of `tag:<value>`"),
					),
				},
			},
		},
	}
}

func (r *preAuthKeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan preAuthKeyResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	user := plan.User.ValueString()
	reusable := plan.Reusable.ValueBool()
	ephemeral := plan.Ephemeral.ValueBool()
	aclTags := []string{}
	for _, r := range plan.ACLTags.Elements() {
		conv := r.(types.String)
		aclTags = append(aclTags, conv.ValueString())
	}

	expireDuration := "1d"
	if plan.TimeToExpire.ValueString() != "" {
		expireDuration = plan.TimeToExpire.ValueString()
	}

	expireAt, err := utils.MakeExpiraryTime(expireDuration)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating pre auth key",
			"Could not create pre auth key, unexpected error: "+err.Error(),
		)
		return
	}

	input := service.CreatePreAuthKeyInput{
		User:       user,
		Reusable:   reusable,
		Ephemeral:  ephemeral,
		Expiration: expireAt,
		ACLTags:    aclTags,
	}
	tflog.Debug(ctx, fmt.Sprintf("Creating pre auth key with input: %v", input))
	key, err := r.client.CreatePreAuthKey(ctx, input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating api key",
			"Could not create api key, unexpected error: "+err.Error(),
		)
		return
	}

	expiresAt := key.Expiration.DeepCopy().String()
	isExpired, err := utils.IsExpired(expiresAt)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating api key",
			"Could not create api key, unexpected error: "+err.Error(),
		)
		return
	}

	m := preAuthKeyResourceModel{
		TimeToExpire: plan.TimeToExpire,
		Id:           types.StringValue(key.ID),
		User:         types.StringValue(key.User),
		Key:          types.StringValue(key.Key),
		Reusable:     types.BoolValue(key.Reusable),
		Ephemeral:    types.BoolValue(key.Ephemeral),
		Used:         types.BoolValue(key.Used),
		Expired:      types.BoolValue(isExpired),
		Expiration:   types.StringValue(expiresAt),
		CreatedAt:    types.StringValue(key.CreatedAt.DeepCopy().String()),
	}

	// API method doesn't return the list of tags so we need to set it from the plan if any were defined
	if plan.ACLTags.IsNull() || plan.ACLTags.IsUnknown() {
		tags, diags := types.ListValueFrom(ctx, types.StringType, key.ACLTags)
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}

		m.ACLTags = tags
	} else {
		m.ACLTags = plan.ACLTags
	}

	diags = resp.State.Set(ctx, &m)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *preAuthKeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError(
		"Error updating pre auth key",
		"Api keys cannot be updated",
	)
}

func (r *preAuthKeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state preAuthKeyResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.ExpirePreAuthKey(ctx, state.User.ValueString(), state.Key.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting pre auth key",
			"Could not expire key, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *preAuthKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state preAuthKeyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := state.Id.ValueString()
	user := state.User.ValueString()
	keys, err := r.client.ListPreAuthKeys(ctx, user)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to get api key",
			"An error was encountered retrieving the api key.\n"+
				err.Error(),
		)
		return
	}

	var m preAuthKeyResourceModel
	for _, key := range keys {
		if key.ID == id {
			expiresAt := key.Expiration.DeepCopy().String()
			isExpired, err := utils.IsExpired(expiresAt)
			if err != nil {
				resp.Diagnostics.AddError(
					"Error reading the pre auth key",
					"Could not read pre auth key, unexpected error: "+err.Error(),
				)
				return
			}

			m = preAuthKeyResourceModel{
				TimeToExpire: state.TimeToExpire,
				Id:           types.StringValue(key.ID),
				User:         types.StringValue(key.User),
				Key:          types.StringValue(key.Key),
				Reusable:     types.BoolValue(key.Reusable),
				Ephemeral:    types.BoolValue(key.Ephemeral),
				Used:         types.BoolValue(key.Used),
				Expired:      types.BoolValue(isExpired),
				Expiration:   types.StringValue(expiresAt),
				CreatedAt:    types.StringValue(key.CreatedAt.DeepCopy().String()),
			}

			tags, diags := types.ListValueFrom(ctx, types.StringType, key.ACLTags)
			if diags.HasError() {
				resp.Diagnostics.Append(diags...)
				return
			}

			m.ACLTags = tags
		}
	}

	diags := resp.State.Set(ctx, m)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
