package healthchecksio

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/kristofferahl/go-healthchecksio"
)

func resourceHealthcheck() *schema.Resource {
	return &schema.Resource{
		Create: resourceHealthcheckCreate,
		Read:   resourceHealthcheckRead,
		Update: resourceHealthcheckUpdate,
		Delete: resourceHealthcheckDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the healthcheck",
				Required:    true,
			},
			"tags": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Tags associated with the healthcheck",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"schedule": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Schedule defining the healthcheck",
				Optional:    true,
			},
			"timezone": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Timezone used for the schedule",
				Optional:    true,
			},
		},
	}
}

func resourceHealthcheckCreate(d *schema.ResourceData, m interface{}) error {
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

	return resourceHealthcheckRead(d, m)
}

func resourceHealthcheckRead(d *schema.ResourceData, m interface{}) error {
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
	d.Set("schedule", healthcheck.Schedule)
	d.Set("timezone", healthcheck.Timezone)

	return nil
}

func resourceHealthcheckUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*healthchecksio.Client)

	key := d.Id()
	log.Printf("[INFO] Updating healthcheck with key: %s", key)

	healthcheck, err := createHealthcheckFromResourceData(d)
	if err != nil {
		return fmt.Errorf("Failed to prepare healthcheck from resource data: %s", err)
	}

	log.Printf("[DEBUG] healthcheck update: %#v", healthcheck)

	if d.HasChange("tags") || d.HasChange("schedule") || d.HasChange("timezone") {
		_, err = client.Update(key, *healthcheck)
		if err != nil {
			return fmt.Errorf("Failed to update healthcheck: %s", err)
		}
	}

	return nil
}

func resourceHealthcheckDelete(d *schema.ResourceData, m interface{}) error {
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

	if attr, ok := d.GetOk("schedule"); ok {
		healthcheck.Schedule = attr.(string)
	}

	if attr, ok := d.GetOk("timezone"); ok {
		healthcheck.Timezone = attr.(string)
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
