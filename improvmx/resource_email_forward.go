package improvmx

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEmailForward() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailForwardCreate,
		Read:   resourceEmailForwardRead,
		Update: resourceEmailForwardUpdate,
		Delete: resourceEmailForwardDelete,
		Importer: &schema.ResourceImporter{
			State: resourceEmailForwardImport,
		},

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
	m := meta.(*Meta)
	m.Client.CreateEmailForward(d.Get("domain").(string), d.Get("alias_name").(string), d.Get("destination_email").(string))

	return resourceEmailForwardRead(d, meta)
}

func resourceEmailForwardRead(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*Meta)
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	for {
		resp := m.Client.GetEmailForward(d.Get("domain").(string), d.Get("alias_name").(string))

		log.Printf("[DEBUG] Got status code %v from ImprovMX API on Read for email_forward %s@%s, success: %v, errors: %v.", resp.Code, d.Get("alias_name").(string), d.Get("domain").(string), resp.Success, resp.Errors)

		if resp.Code == 429 {
			log.Printf("[DEBUG] Sleeping for 10 seconds to allow rate limit to recover.")
			time.Sleep(10 * time.Second)
		}

		if resp.Code == 404 {
			log.Printf("[DEBUG] Couldn't find the resource in ImprovMX. Aborting")
			return fmt.Errorf("HTTP response code %d, error text: %s", resp.Code, resp.Errors.Domain)
		}

		if resp.Success {
			d.SetId(strconv.FormatInt(resp.Alias.Id, 10))
			d.Set("alias_name", resp.Alias.Alias)
			d.Set("destination_email", resp.Alias.Forward)

			return nil
		}
	}
}

func resourceEmailForwardUpdate(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*Meta)

	for {
		resp := m.Client.UpdateEmailForward(d.Get("domain").(string), d.Get("alias_name").(string), d.Get("destination_email").(string))

		log.Printf("[DEBUG] Got status code %v from ImprovMX API on Update for email_forward %s@%s, success: %v, errors: %v.", resp.Code, d.Get("domain").(string), d.Get("alias_name").(string), resp.Success, resp.Errors)

		if resp.Code == 429 {
			log.Printf("[DEBUG] Sleeping for 10 seconds to allow rate limit to recover.")
			time.Sleep(10 * time.Second)
		}

		if resp.Code == 404 {
			log.Printf("[DEBUG] Couldn't find the resource in ImprovMX. Aborting")
			return fmt.Errorf("HTTP response code %d, error text: %s", resp.Code, resp.Errors.Domain)
		}

		if resp.Success {
			return resourceEmailForwardRead(d, meta)
		}
	}
}

func resourceEmailForwardDelete(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*Meta)
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	m.Client.DeleteEmailForward(d.Get("domain").(string), d.Get("alias_name").(string))

	return nil
}

func resourceEmailForwardImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), "_")

	if len(parts) != 2 {
		return nil, fmt.Errorf("Error Importing email forward. Please make sure the email forward ID is in the form DOMAIN_EMAILFORWARDNAME (i.e. example.com_hi)")
	}

	d.SetId(parts[1])
	d.Set("domain", parts[0])
	d.Set("alias_name", parts[1])

	resourceEmailForwardRead(d, meta)

	return []*schema.ResourceData{d}, nil
}
