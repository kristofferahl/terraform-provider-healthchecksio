package healthchecksio

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Provider creates a new healthchecksio provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(APIKeyEnvName, nil),
				Description: "A healthchecks.io api key.",
			},
			"api_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(APIURLEnvName, ""),
				Description: "A healthchecks.io api base URL.",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"healthchecksio_channel": dataSourceHealthcheckChannel(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"healthchecksio_check": resourceHealthcheck(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		APIKey:  d.Get("api_key").(string),
		BaseURL: d.Get("api_url").(string),
	}
	return config.Client()
}
