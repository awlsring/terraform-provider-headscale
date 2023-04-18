package apikey

import (
	"context"
	"strings"
	"time"

	"github.com/awlsring/terraform-provider-headscale/internal/service"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &apiKeyResource{}
	_ resource.ResourceWithConfigure = &apiKeyResource{}
)

func Resource() resource.Resource {
	return &apiKeyResource{}
}

type apiKeyResource struct {
	client service.Headscale
}

func (d *apiKeyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_key"
}

func (d *apiKeyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(service.Headscale)
}

func (d *apiKeyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The API key resource allows you to create an API key that can be used to authenticate with the Headscale API. By default keys that are created with this resource will not expire. To create a key that expires, set the `days_to_expire` attribute to the number of days until the key expires. Keys cannot be modified, so any change to the input on this resource will cause the key to be expired and a new key to be created.",
		Attributes: map[string]schema.Attribute{
			"days_to_expire": schema.Int64Attribute{
				Optional:    true,
				Description: "The number of days until the api key expires. No value creates a key that won't expire.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The id of the api key.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"key": schema.StringAttribute{
				Computed:    true,
				Description: "The api key.",
				Sensitive:   true,
			},
			"prefix": schema.StringAttribute{
				Computed:    true,
				Description: "The api key's prefix.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"expiration": schema.StringAttribute{
				Computed:    true,
				Description: "Expiration date of the api key.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"expired": schema.BoolAttribute{
				Computed:    true,
				Description: "If the api key is expired.",
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: "The time the device entry was created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *apiKeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan apikeyResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	expireAt := makeExpiraryTime(int(plan.DaysToExpire.ValueInt64()))
	apiKey, err := r.client.CreateAPIKey(ctx, expireAt)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating api key",
			"Could not create api key, unexpected error: "+err.Error(),
		)
		return
	}

	apiKeys, err := r.client.ListAPIKeys(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating api key",
			"Could not create api key, unexpected error: "+err.Error(),
		)
		return
	}

	var m apikeyResourceModel
	for _, key := range apiKeys {
		if strings.Contains(apiKey, key.Prefix) {
			expiresAt := key.Expiration.DeepCopy().String()
			isExpired, err := isExpired(expiresAt)
			if err != nil {
				resp.Diagnostics.AddError(
					"Error creating api key",
					"Could not create api key, unexpected error: "+err.Error(),
				)
				return
			}

			m = apikeyResourceModel{
				DaysToExpire: plan.DaysToExpire,
				Id:           types.StringValue(key.ID),
				Prefix:       types.StringValue(key.Prefix),
				Key:          types.StringValue(apiKey),
				Expiration:   types.StringValue(expiresAt),
				Expired:      types.BoolValue(isExpired),
				CreatedAt:    types.StringValue(key.CreatedAt.DeepCopy().String()),
			}
		}
	}

	diags = resp.State.Set(ctx, &m)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *apiKeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError(
		"Error creating api key",
		"Api keys cannot be updated",
	)
}

func (r *apiKeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state apikeyResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.ExpireAPIKey(ctx, state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting api key",
			"Could not expire key, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *apiKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state apikeyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiKey := state.Prefix.ValueString()
	apiKeys, err := r.client.ListAPIKeys(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to get api key",
			"An error was encountered retrieving the api key.\n"+
				err.Error(),
		)
		return
	}

	var m apikeyResourceModel
	for _, key := range apiKeys {
		if key.Prefix == apiKey {
			expiresAt := key.Expiration.DeepCopy().String()
			isExpired, err := isExpired(expiresAt)
			if err != nil {
				resp.Diagnostics.AddError(
					"Error creating api key",
					"Could not create api key, unexpected error: "+err.Error(),
				)
				return
			}

			m = apikeyResourceModel{
				DaysToExpire: state.DaysToExpire,
				Id:           types.StringValue(key.ID),
				Key:          state.Key,
				Prefix:       types.StringValue(key.Prefix),
				Expiration:   types.StringValue(expiresAt),
				Expired:      types.BoolValue(isExpired),
				CreatedAt:    types.StringValue(key.CreatedAt.DeepCopy().String()),
			}
		}
	}

	diags := resp.State.Set(ctx, m)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func makeExpiraryTime(d int) *strfmt.DateTime {
	if d == 0 {
		return nil
	}
	now := time.Now()
	h := 24 * d
	then := now.Add(time.Duration(h) * time.Hour)

	time := strfmt.DateTime(then)
	return &time
}
