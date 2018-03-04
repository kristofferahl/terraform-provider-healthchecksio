package healthchecksio

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider creates a new healthchecksio provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(EnvironmentKey, nil),
				Description: "A healthchecks.io api key.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"healthchecksio_check": resourceHealthcheck(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		APIKey: d.Get("api_key").(string),
	}
	return config.Client()
}
