package device

import (
	"context"
	"fmt"
	"strings"

	"github.com/awlsring/terraform-provider-headscale/internal/service"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ datasource.DataSource              = &devicesDataSource{}
	_ datasource.DataSourceWithConfigure = &devicesDataSource{}
)

func DataSource() datasource.DataSource {
	return &devicesDataSource{}
}

type devicesDataSource struct {
	client service.Headscale
}

type dataSourceModel struct {
	User       types.String  `tfsdk:"user"`
	NamePrefix types.String  `tfsdk:"name_prefix"`
	Devices    []deviceModel `tfsdk:"devices"`
}

type deviceModel struct {
	Id             types.String   `tfsdk:"id"`
	Addresses      []types.String `tfsdk:"addresses"`
	Name           types.String   `tfsdk:"name"`
	User           types.String   `tfsdk:"user"`
	Expiry         types.String   `tfsdk:"expiry"`
	CreatedAt      types.String   `tfsdk:"created_at"`
	RegisterMethod types.String   `tfsdk:"register_method"`
	GivenName      types.String   `tfsdk:"given_name"`
	Tags           []types.String `tfsdk:"tags"`
}

func (d *devicesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices"
}

func (d *devicesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(service.Headscale)
}

func (d *devicesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"user": schema.StringAttribute{
				Optional:    true,
				Description: "The id of the bond. Formatted as `{node}/{name}`.",
			},
			"name_prefix": schema.StringAttribute{
				Optional:    true,
				Description: "The id of the bond. Formatted as `{node}/{name}`.",
			},
			"devices": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Description: "The id of the bridge. Formatted as /{node}/{name}.",
						},
						"addresses": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							Description: "List of physical network interfaces on the machine.",
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: "The SSL fingerprint of the node",
						},
						"user": schema.StringAttribute{
							Computed:    true,
							Description: "The SSL fingerprint of the node",
						},
						"expiry": schema.StringAttribute{
							Computed:    true,
							Description: "The position of the network interface in the VM as an int. Used to determine the interface name (net0, net1, etc).",
						},
						"created_at": schema.StringAttribute{
							Computed: true,
						},
						"register_method": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"tags": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							Description: "List of physical network interfaces on the machine.",
						},
						"given_name": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *devicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state dataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var namePrefix *string
	if state.NamePrefix.ValueString() != "" {
		n := state.NamePrefix.ValueString()
		namePrefix = &n
		tflog.Debug(ctx, fmt.Sprintf("namePrefix: %v", *namePrefix))
	}

	var user *string
	if state.User.ValueString() != "" {
		u := state.User.ValueString()
		user = &u
		tflog.Debug(ctx, fmt.Sprintf("user: %v", *user))
	}

	deviceList, err := d.client.ListDevices(ctx, user)
	tflog.Debug(ctx, fmt.Sprintf("Devices: %v", len(deviceList)))
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to list devices",
			"An error was encountered retrieving devices.\n"+
				err.Error(),
		)
		return
	}

	for _, device := range deviceList {
		tflog.Debug(ctx, fmt.Sprintf("Device: %s", device.Name))

		if namePrefix != nil {
			if !strings.HasPrefix(device.Name, *namePrefix) {
				continue
			}
		}

		dm := deviceModel{
			Id:        types.StringValue(device.ID),
			Addresses: []types.String{},
			Name:      types.StringValue(device.Name),
			User:      types.StringValue(device.User.ID),
			Expiry:    types.StringValue(device.Expiry.DeepCopy().String()),
			CreatedAt: types.StringValue(device.CreatedAt.DeepCopy().String()),
			Tags:      []types.String{},
			GivenName: types.StringValue(device.GivenName),
		}

		for _, add := range device.IPAddresses {
			dm.Addresses = append(dm.Addresses, types.StringValue(add))
		}

		for _, t := range device.ValidTags {
			dm.Tags = append(dm.Tags, types.StringValue(t))
		}

		for _, t := range device.ForcedTags {
			dm.Tags = append(dm.Tags, types.StringValue(t))
		}

		if device.RegisterMethod != nil {
			dm.RegisterMethod = types.StringValue(string(*device.RegisterMethod))
		}

		state.Devices = append(state.Devices, dm)
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
