package improvmx

import (
	"sync"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	improvmxApi "github.com/issyl0/go-improvmx"
)

type Meta struct {
	Resource *schema.ResourceData
	Client   *improvmxApi.Client
	Mutex    sync.Mutex
}

func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("IMPROVMX_API_TOKEN", nil),
				Description: "The API token for API operations."},
		},
		ResourcesMap: map[string]*schema.Resource{
			"improvmx_domain":          resourceDomain(),
			"improvmx_email_forward":   resourceEmailForward(),
			"improvmx_smtp_credential": resourceSMTPCredential(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"improvmx_domain_check": DataSourceDomainCheck(),
		},
		ProviderMetaSchema: map[string]*schema.Schema{},
	}

	p.ConfigureFunc = providerConfigure

	return p
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	m := &Meta{
		Resource: d,
		Client:   improvmxApi.NewClient(d.Get("token").(string)),
		Mutex:    sync.Mutex{},
	}

	return m, nil
}
