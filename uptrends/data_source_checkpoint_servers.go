package uptrends

import (
	"time"

	"github.com/VeritasOS/terraform-provider-uptrends.git/uptapi"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceCheckpointServerIPs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCheckpointServerIPsRead,
		Schema: map[string]*schema.Schema{
			"ip_list": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
		},
	}
}

func dataSourceCheckpointServerIPsRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*uptapi.Client)

	l, err := c.CheckpointServerIPs()
	if err != nil {
		return err
	}

	d.SetId(time.Now().UTC().String())
	d.Set("ip_list", l)

	return nil
}
