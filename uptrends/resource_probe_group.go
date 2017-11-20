package uptrends

import (
	"github.com/VeritasOS/terraform-provider-uptrends.git/uptapi"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceProbeGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceProbeGroupCreate,
		Read:   resourceProbeGroupRead,
		Update: resourceProbeGroupUpdate,
		Delete: resourceProbeGroupDelete,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceProbeGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*uptapi.Client)

	id, err := c.NewProbeGroup(
		&uptapi.ProbeGroup{
			Name: d.Get("name").(string),
		},
	)
	if err != nil {
		return err
	}

	d.SetId(id)

	return resourceProbeGroupRead(d, m)
}

func resourceProbeGroupRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*uptapi.Client)
	id := d.Id()

	g, err := c.ProbeGroup(id)
	if err != nil {
		return err
	}

	d.SetId(g.GUID)
	d.Set("name", g.Name)

	return nil
}

func resourceProbeGroupUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*uptapi.Client)
	id := d.Id()
	name := d.Get("name").(string)

	return c.UpdateProbeGroup(
		&uptapi.ProbeGroup{
			GUID: id,
			Name: name,
		},
	)
}

func resourceProbeGroupDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*uptapi.Client)
	id := d.Id()
	return c.DeleteProbeGroup(id)
}
