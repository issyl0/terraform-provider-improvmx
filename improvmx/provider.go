package improvmx

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Provider returns a terraform.Provider.
func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("IMPROVMX_API_TOKEN", nil),
				Description: "The API token for API operations.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			// "improvmx_domain":        resourceDomain(),
			"improvmx_email_forward": resourceEmailForward(),
		},
	}
	p.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		return providerConfigure(d)
	}

	return p
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Token: d.Get("token").(string),
	}

	return config.Client()
}
