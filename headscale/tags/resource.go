package tags

import (
	"context"
	"fmt"
	"regexp"

	"github.com/awlsring/terraform-provider-headscale/internal/service"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &deviceTagsResource{}
	_ resource.ResourceWithConfigure   = &deviceTagsResource{}
	_ resource.ResourceWithImportState = &deviceTagsResource{}
)

func Resource() resource.Resource {
	return &deviceTagsResource{}
}

type deviceTagsResource struct {
	client service.Headscale
}

func (d *deviceTagsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_tags"
}

func (d *deviceTagsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(service.Headscale)
}

func (d *deviceTagsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The device tags resource allows setting tags on a device registered in Headscale instance. Utilizing this resource will reset any previous configuration for tags applied to the device. If a tag was previously applied, but is not present in the list of tags, it will be removed.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The resolved id of the device.",
			},
			"device_id": schema.StringAttribute{
				Required:    true,
				Description: "The id of the device to set tags on.",
			},
			"tags": schema.ListAttribute{
				Required:    true,
				ElementType: types.StringType,
				Description: "The tags applied to the device.",
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

func (r *deviceTagsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan deviceTagModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	device, err := r.tagDevice(ctx, &plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating tagging device",
			"Could not tag device, unexpected error: "+err.Error(),
		)
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("tagged device with id '%s'", device.DeviceId.ValueString()))

	diags = resp.State.Set(ctx, device)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *deviceTagsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan deviceTagModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	device, err := r.tagDevice(ctx, &plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating tags on device",
			"Could update tags, unexpected error: "+err.Error(),
		)
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("tagged device with id '%s'", device.DeviceId.ValueString()))

	diags = resp.State.Set(ctx, device)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *deviceTagsResource) tagDevice(ctx context.Context, m *deviceTagModel) (*deviceTagModel, error) {
	deviceId := m.DeviceId.ValueString()
	tags := []string{}
	for _, tag := range m.Tags.Elements() {
		conv := tag.(types.String)
		tags = append(tags, conv.ValueString())
	}
	tflog.Debug(ctx, fmt.Sprintf("tagging device with id '%s' with tags '%v'", deviceId, tags))

	device, err := r.client.TagDevice(ctx, deviceId, tags)
	if err != nil {
		return nil, err
	}

	dm := deviceTagModel{
		DeviceId: types.StringValue(device.ID),
		Id:       types.StringValue(device.ID),
	}

	allTags := []string{}
	allTags = append(allTags, device.ForcedTags...)

	c, diags := types.ListValueFrom(ctx, types.StringType, allTags)
	if diags.HasError() {
		return nil, fmt.Errorf("error creating list of tags")
	}

	dm.Tags = c

	return &dm, nil
}

func (r *deviceTagsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state deviceTagModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.TagDevice(ctx, state.DeviceId.ValueString(), []string{})
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting device tags",
			"Could not remove tags, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *deviceTagsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state deviceTagModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	deviceId := state.DeviceId.ValueString()

	device, err := r.readDevice(ctx, deviceId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to get device",
			"An error was encountered retrieving the device.\n"+
				err.Error(),
		)
		return
	}

	diags := resp.State.Set(ctx, device)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *deviceTagsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	device, err := r.readDevice(ctx, req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to get device",
			"An error was encountered retrieving the device.\n"+
				err.Error(),
		)
		return
	}

	diags := resp.State.Set(ctx, device)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *deviceTagsResource) readDevice(ctx context.Context, id string) (*deviceTagModel, error) {
	device, err := r.client.GetDevice(ctx, id)
	if err != nil {
		return nil, err
	}

	dm := deviceTagModel{
		DeviceId: types.StringValue(device.ID),
		Id:       types.StringValue(device.ID),
	}

	allTags := []string{}
	allTags = append(allTags, device.ForcedTags...)

	c, diags := types.ListValueFrom(ctx, types.StringType, allTags)
	if diags.HasError() {
		return nil, fmt.Errorf("error creating list of tags")
	}

	dm.Tags = c

	return &dm, nil
}
