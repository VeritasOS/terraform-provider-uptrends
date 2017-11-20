package uptrends

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceProbe() *schema.Resource {
	return &schema.Resource{
		Create: resourceProbeCreate,
		Read:   resourceProbeRead,
		Update: resourceProbeUpdate,
		Delete: resourceProbeDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceProbeCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceProbeRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceProbeUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceProbeDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
