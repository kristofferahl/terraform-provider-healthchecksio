package healthchecksio

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/kristofferahl/go-healthchecksio/v2"
)

func resourceHealthcheck() *schema.Resource {
	return &schema.Resource{
		Create: resourceHealthcheckCreate,
		Read:   resourceHealthcheckRead,
		Update: resourceHealthcheckUpdate,
		Delete: resourceHealthcheckDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the check",
				Required:    true,
			},
			"tags": {
				Type:        schema.TypeList,
				Description: "Tags associated with the check",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"timeout": {
				Type:         schema.TypeInt,
				Description:  "Timeout period of the check",
				Optional:     true,
				Default:      86400,
				ValidateFunc: validation.IntBetween(60, 31536000),
			},
			"grace": {
				Type:         schema.TypeInt,
				Description:  "Grace period for the check",
				Optional:     true,
				Default:      3600,
				ValidateFunc: validation.IntBetween(60, 31536000),
			},
			"schedule": {
				Type:        schema.TypeString,
				Description: "A cron expression defining the check's schedule",
				Optional:    true,
			},
			"timezone": {
				Type:        schema.TypeString,
				Description: "Timezone used for the schedule",
				Optional:    true,
			},
			"channels": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.IsUUID,
				},
				Description: "Channels integrated with the check",
				Optional:    true,
			},
			"ping_url": {
				Type:        schema.TypeString,
				Description: "Ping URL associated with this check",
				Computed:    true,
			},
			"pause_url": {
				Type:        schema.TypeString,
				Description: "Pause URL associated with this check",
				Computed:    true,
			},
			"desc": {
				Type:        schema.TypeString,
				Description: "Description of the check",
				Optional:    true,
			},
			"methods": {
				Type:         schema.TypeString,
				Description:  "Allowed HTTP methods for making ping requests",
				Optional:     true,
				Default:      "",
				ValidateFunc: validation.StringInSlice([]string{"", "POST"}, false),
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
		return fmt.Errorf("failed to prepare healthcheck from resource data: %s", err)
	}

	log.Printf("[DEBUG] healthcheck create: %#v", healthcheck)

	resp, err := client.Create(*healthcheck)
	if err != nil {
		return fmt.Errorf("failed to create healthcheck: %s", err)
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
		return fmt.Errorf("error reading healthchecks: %s", err)
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

	tags := make([]string, 0)
	if len(healthcheck.Tags) > 0 {
		tags = strings.Split(healthcheck.Tags, " ")
	}

	channels := make([]string, 0)
	if len(healthcheck.Channels) > 0 {
		channels = strings.Split(healthcheck.Channels, ",")

		if attr, ok := d.GetOk("channels"); ok {
			state := toSliceOfString(attr.([]interface{}))
			channels = sortByLeft(state, channels)
		}
	}

	values := map[string]interface{}{
		"name":      healthcheck.Name,
		"tags":      tags,
		"timeout":   healthcheck.Timeout,
		"grace":     healthcheck.Grace,
		"schedule":  healthcheck.Schedule,
		"timezone":  healthcheck.Timezone,
		"channels":  channels,
		"ping_url":  healthcheck.PingURL,
		"pause_url": healthcheck.PauseURL,
		"desc":      healthcheck.Description,
		"methods":   healthcheck.Methods,
	}

	for k, v := range values {
		if err := d.Set(k, v); err != nil {
			return err
		}
	}

	return nil
}

func resourceHealthcheckUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*healthchecksio.Client)

	key := d.Id()
	log.Printf("[INFO] Updating healthcheck with key: %s", key)

	healthcheck, err := createHealthcheckFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to prepare healthcheck from resource data: %s", err)
	}

	log.Printf("[DEBUG] healthcheck update: %#v", healthcheck)

	if hasChange(d) {
		_, err = client.Update(key, *healthcheck)
		if err != nil {
			return fmt.Errorf("failed to update healthcheck: %s", err)
		}
	}

	return resourceHealthcheckRead(d, m)
}

func resourceHealthcheckDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*healthchecksio.Client)

	key := d.Id()
	log.Printf("[INFO] Deleting healthcheck with key: %s", key)

	if _, err := client.Delete(key); err != nil {
		return fmt.Errorf("error deleting healthcheck: %s", err)
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

	if attr, ok := d.GetOk("timeout"); ok {
		healthcheck.Timeout = attr.(int)
	}

	if attr, ok := d.GetOk("grace"); ok {
		healthcheck.Grace = attr.(int)
	}

	if attr, ok := d.GetOk("schedule"); ok {
		healthcheck.Schedule = attr.(string)
	}

	if attr, ok := d.GetOk("timezone"); ok {
		healthcheck.Timezone = attr.(string)
	}

	if attr, ok := d.GetOk("channels"); ok {
		channels := toSliceOfString(attr.([]interface{}))
		healthcheck.Channels = strings.Join(channels, ",")
	}

	if attr, ok := d.GetOk("desc"); ok {
		healthcheck.Description = attr.(string)
	}

	if attr, ok := d.GetOk("methods"); ok {
		healthcheck.Methods = attr.(string)
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

func hasChange(d *schema.ResourceData) bool {
	return d.HasChange("desc") || d.HasChange("tags") || d.HasChange("timeout") ||
		d.HasChange("grace") || d.HasChange("schedule") || d.HasChange("methods") ||
		d.HasChange("timezone") || d.HasChange("channels") || d.HasChange("name")
}

func sortByLeft(left, right []string) []string {
	var sorted []string
	var diff []string

	for _, i := range left {
		if contains(right, i) && !contains(sorted, i) {
			sorted = append(sorted, i)
		}
	}

	for _, i := range right {
		if !contains(sorted, i) && !contains(diff, i) {
			diff = append(diff, i)
		}
	}

	sorted = append(sorted, sort.StringSlice(diff)...)

	return sorted
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
