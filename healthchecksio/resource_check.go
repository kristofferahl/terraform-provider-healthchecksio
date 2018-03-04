package healthchecksio

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/kristofferahl/go-healthchecksio"
)

func resourceCheck() *schema.Resource {
	return &schema.Resource{
		Create: resourceCheckCreate,
		Read:   resourceCheckRead,
		Update: resourceCheckUpdate,
		Delete: resourceCheckDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceCheckCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*healthchecksio.Client)

	name := d.Get("name").(string)
	log.Printf("[INFO] creating healthcheck named: %s", name)

	healthcheck, err := createHealthcheckFromResourceData(d)
	if err != nil {
		return fmt.Errorf("Failed to prepare healthcheck from resource data: %s", err)
	}

	log.Printf("[DEBUG] healthcheck create: %#v", healthcheck)

	resp, err := client.Create(*healthcheck)
	if err != nil {
		return fmt.Errorf("Failed to create healthcheck: %s", err)
	}

	d.SetId(resp.ID())

	return resourceCheckRead(d, m)
}

func resourceCheckRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*healthchecksio.Client)

	key := d.Id()
	log.Printf("[INFO] Reading healthcheck with key: %s", key)

	healthchecks, err := client.GetAll()
	if err != nil {
		return fmt.Errorf("Error reading healthchecks: %s", err)
	}

	var healthcheck *healthchecksio.HealthcheckResponse
	for _, hc := range healthchecks {
		if hc.ID() == d.Id() {
			healthcheck = hc
			break
		}
	}

	if healthcheck == nil {
		d.SetId("")
		return nil
	}

	d.Set("name", healthcheck.Name)
	d.Set("tags", strings.Split(healthcheck.Tags, " "))

	return nil
}

func resourceCheckUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*healthchecksio.Client)

	key := d.Id()
	log.Printf("[INFO] Updating healthcheck with key: %s", key)

	healthcheck, err := createHealthcheckFromResourceData(d)
	if err != nil {
		return fmt.Errorf("Failed to prepare healthcheck from resource data: %s", err)
	}

	log.Printf("[DEBUG] healthcheck update: %#v", healthcheck)

	if d.HasChange("tags") {
		_, err = client.Update(key, *healthcheck)
		if err != nil {
			return fmt.Errorf("Failed to update healthcheck: %s", err)
		}
	}

	return nil
}

func resourceCheckDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*healthchecksio.Client)

	key := d.Id()
	log.Printf("[INFO] Deleting healthcheck with key: %s", key)

	if _, err := client.Delete(key); err != nil {
		return fmt.Errorf("Error deleting healthcheck: %s", err)
	}

	return nil
}

func createHealthcheckFromResourceData(d *schema.ResourceData) (*healthchecksio.Healthcheck, error) {
	healthcheck := healthchecksio.Healthcheck{}

	if attr, ok := d.GetOk("name"); ok {
		healthcheck.Name = attr.(string)
	}

	if attr, ok := d.GetOk("tags"); ok {
		tags := toSliceOfString(attr.([]interface{}))
		healthcheck.Tags = strings.Join(tags, " ")
	}

	return &healthcheck, nil
}

func toSliceOfString(a []interface{}) []string {
	vs := make([]string, 0, len(a))
	for _, v := range a {
		val, ok := v.(string)
		if ok && val != "" {
			vs = append(vs, val)
		}
	}
	return vs
}
