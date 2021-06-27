package improvmx

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	improvmxApi "github.com/issyl0/go-improvmx"
)

func resourceDomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceDomainCreate,
		Read:   resourceDomainRead,
		// Update: resourceDomainUpdate,
		Delete: resourceDomainDelete,
		// Importer: &schema.ResourceImporter{
		// 	State: resourceDomainImport,
		// },

		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDomainCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*improvmxApi.Client)
	client.CreateDomain(d.Get("domain").(string))

	return resourceDomainRead(d, meta)
}

func resourceDomainRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*improvmxApi.Client)
	resp := client.GetDomain(d.Get("domain").(string))

	d.SetId(strconv.FormatInt(resp.Domain.Aliases[0].Id, 10))
	d.Set("domain", resp.Domain.Domain)

	return nil
}

func resourceDomainDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*improvmxApi.Client)
	client.DeleteDomain(d.Get("domain").(string))

	return nil
}
