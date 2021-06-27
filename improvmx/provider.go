package improvmx

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	improvmxApi "github.com/issyl0/go-improvmx"
)

func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema:             map[string]*schema.Schema{"token": {Type: schema.TypeString, Required: true, DefaultFunc: schema.EnvDefaultFunc("IMPROVMX_API_TOKEN", nil), Description: "The API token for API operations."}},
		ResourcesMap:       map[string]*schema.Resource{"improvmx_email_forward": resourceEmailForward()},
		DataSourcesMap:     map[string]*schema.Resource{},
		ProviderMetaSchema: map[string]*schema.Schema{},
	}

	p.ConfigureFunc = providerConfigure

	return p
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	token := d.Get("token").(string)
	return improvmxApi.NewClient(token), nil
}
