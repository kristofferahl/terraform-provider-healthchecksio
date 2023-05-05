package healthchecksio

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kristofferahl/go-healthchecksio"
)

func dataSourceHealthcheckChannel() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHealthcheckChannelRead,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Description: "ID of the channel",
				Computed:    true,
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the channel to search for",
				Optional:    true,
			},
			"kind": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Kind of channel",
				Required:    true,
			},
		},
	}
}

func dataSourceHealthcheckChannelRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*healthchecksio.Client)

	name := d.Get("name").(string)
	kind := d.Get("kind").(string)
	log.Printf("[INFO] Reading healthcheck with keys: %s, %s", name, kind)

	channels, err := client.GetAllChannels()
	if err != nil {
		return fmt.Errorf("Error reading healthchecks: %s", err)
	}

	var channel *healthchecksio.HealthcheckChannelResponse
	for _, c := range channels {
		if isTarget(c, name, kind) {
			channel = c
			break
		}
	}

	if channel == nil {
		d.SetId("")
		return nil
	}

	d.SetId(channel.ID)
	d.Set("name", channel.Name)
	d.Set("kind", channel.Kind)

	return nil
}

func isTarget(c *healthchecksio.HealthcheckChannelResponse, name, kind string) bool {
	if name == "" {
		return c.Kind == kind
	}

	return c.Name == name && c.Kind == kind
}
