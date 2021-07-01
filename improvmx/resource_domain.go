package improvmx

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	improvmxApi "github.com/issyl0/go-improvmx"
)

func resourceDomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceDomainCreate,
		Read:   resourceDomainRead,
		Update: resourceDomainUpdate,
		Delete: resourceDomainDelete,
		Importer: &schema.ResourceImporter{
			State: resourceDomainImport,
		},

		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"notification_email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"whitelabel": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceDomainCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*improvmxApi.Client)
	client.CreateDomain(d.Get("domain").(string), d.Get("notification_email").(string), d.Get("whitelabel").(string))

	return resourceDomainRead(d, meta)
}

func resourceDomainRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*improvmxApi.Client)
	resp := client.GetDomain(d.Get("domain").(string))

	d.SetId(resp.Domain.Domain)
	d.Set("domain", resp.Domain.Domain)
	d.Set("notification_email", resp.Domain.NotificationEmail)
	d.Set("whitelabel", resp.Domain.Whitelabel)

	return nil
}

func resourceDomainUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*improvmxApi.Client)
	client.UpdateDomain(d.Get("domain").(string), d.Get("notification_email").(string), d.Get("whitelabel").(string))

	return resourceDomainRead(d, meta)
}

func resourceDomainDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*improvmxApi.Client)
	client.DeleteDomain(d.Get("domain").(string))

	return nil
}

func resourceDomainImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("domain", d.Id())
	resourceDomainRead(d, meta)

	return []*schema.ResourceData{d}, nil
}
