package improvmx

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceEmailForward() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailForwardCreate,
		// Read:   resourceEmailForwardRead,
		// Update: resourceEmailForwardUpdate,
		// Delete: resourceEmailForwardDelete,
		// Importer: &schema.ResourceImporter{
		// 	State: resourceEmailForwardImport,
		// },

		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
			},

			"alias_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"destination_email": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceEmailForwardCreate(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(*Client)
	provider.client.CreateEmailForward(d.Get("domain").(string), d.Get("alias_name").(string), d.Get("destination_email").(string))

	return nil
}
