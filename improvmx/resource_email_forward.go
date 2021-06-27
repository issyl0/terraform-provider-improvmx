package improvmx

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	improvmxApi "github.com/issyl0/go-improvmx"
)

func resourceEmailForward() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailForwardCreate,
		Read:   resourceEmailForwardRead,
		// Update: resourceEmailForwardUpdate,
		Delete: resourceEmailForwardDelete,
		// Importer: &schema.ResourceImporter{
		// 	State: resourceEmailForwardImport,
		// },

		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"alias_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"destination_email": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEmailForwardCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*improvmxApi.Client)
	client.CreateEmailForward(d.Get("domain").(string), d.Get("alias_name").(string), d.Get("destination_email").(string))

	return resourceEmailForwardRead(d, meta)
}

func resourceEmailForwardRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*improvmxApi.Client)
	resp := client.GetEmailForward(d.Get("domain").(string), d.Get("alias_name").(string))

	d.SetId(strconv.FormatInt(resp.Alias.Id, 10))
	d.Set("alias_name", resp.Alias.Alias)
	d.Set("destination_email", resp.Alias.Forward)

	return nil
}

func resourceEmailForwardDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*improvmxApi.Client)
	client.DeleteEmailForward(d.Get("domain").(string), d.Get("alias_name").(string))

	return nil
}
