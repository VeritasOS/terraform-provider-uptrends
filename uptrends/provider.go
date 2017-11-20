package uptrends

import (
	"github.com/VeritasOS/terraform-provider-uptrends.git/uptapi"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a schema.Provider for Uptrends
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("UPTRENDS_URL", nil),
				Description: "URL of Uptrends API server.",
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("UPTRENDS_USER", nil),
				Description: "Credential user for Uptrends API.",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("UPTRENDS_PASS", nil),
				Description: "Credential password for Uptrends API.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"uptrends_probe":       resourceProbe(),
			"uptrends_probe_group": resourceProbeGroup(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"uptrends_checkpoint_server_ips": dataSourceCheckpointServerIPs(),
		},

		ConfigureFunc: getClient,
	}
}

func getClient(d *schema.ResourceData) (interface{}, error) {
	return uptapi.New(
		d.Get("username").(string),
		d.Get("password").(string),
		d.Get("url").(string),
	)
}
