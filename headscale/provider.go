package headscale

import (
	"context"
	"os"

	"github.com/awlsring/terraform-provider-headscale/headscale/apikey"
	"github.com/awlsring/terraform-provider-headscale/headscale/device"
	"github.com/awlsring/terraform-provider-headscale/headscale/route"
	device_route "github.com/awlsring/terraform-provider-headscale/headscale/route/device"
	"github.com/awlsring/terraform-provider-headscale/headscale/tags"
	"github.com/awlsring/terraform-provider-headscale/headscale/user"
	"github.com/awlsring/terraform-provider-headscale/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ provider.Provider = &HeadscaleProvider{}
)

type HeadscaleProvider struct{}

type HeadscaleProviderConfig struct {
	ApiKey   types.String `tfsdk:"api_key"`
	Endpoint types.String `tfsdk:"endpoint"`
}

func New() provider.Provider {
	return &HeadscaleProvider{}
}

func (p *HeadscaleProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "headscale"
}

func (p *HeadscaleProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_key": schema.StringAttribute{
				Optional:    true,
				Description: "A Headscale api key.",
				Sensitive:   true,
			},
			"endpoint": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Headscale endpoint to connect to. (e.g. `https://headscale.example.com`)",
			},
		},
	}
}

func (p *HeadscaleProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Debug(ctx, "Configuring headscale client")

	var cfg HeadscaleProviderConfig
	resp.Diagnostics.Append(req.Config.Get(ctx, &cfg)...)
	if resp.Diagnostics.HasError() {
		return
	}

	endpoint := os.Getenv("HEADSCALE_ENDPOINT")
	apiKey := os.Getenv("HEADSCALE_API_KEY")

	if !cfg.Endpoint.IsNull() {
		endpoint = cfg.Endpoint.ValueString()
	}
	if !cfg.ApiKey.IsNull() {
		apiKey = cfg.ApiKey.ValueString()
	}
	if endpoint == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("endpoint"),
			"Headscale endpoint",
			"The provider cannot create the Headscale API client as there is a missing or empty value for the endpoint. "+
				"Set the endpoint value in the configuration or use the HEADSCALE_ENDPOINT environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}
	if apiKey == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("apiKey"),
			"Headscale API Key",
			"The provider cannot create the Headscale API client as there is a missing or empty value for the API key. "+
				"Set the API key value in the configuration or use the HEADSCALE_API_KEY environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	scfg := service.ClientConfig{
		Token:    apiKey,
		Endpoint: endpoint,
	}

	ctx = tflog.SetField(ctx, "headscale_endpoint", endpoint)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "headscale_api_key")

	tflog.Debug(ctx, "Creating Headscale client")
	client, err := service.New(scfg)
	if err != nil {
		resp.Diagnostics.AddError(
			"Uable to create Headscale client",
			"An error was encountered configuring the client.\n"+
				err.Error(),
		)
		return
	}

	resp.DataSourceData = client
	resp.ResourceData = client

	tflog.Debug(ctx, "Configured Headscale client", map[string]any{"success": true})
}

func (p *HeadscaleProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		tags.Resource,
		device_route.Resource,
		user.Resource,
		apikey.Resource,
	}
}

func (p *HeadscaleProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		device.DataSourceMultiple,
		device.DataSource,
		route.DataSource,
		device_route.DataSource,
		user.DataSource,
		user.DataSourceMultiple,
		apikey.DataSource,
	}
}
