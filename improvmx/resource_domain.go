package improvmx

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
	m := meta.(*Meta)

	for {
		resp := m.Client.CreateDomain(d.Get("domain").(string), d.Get("notification_email").(string), d.Get("whitelabel").(string))

		log.Printf("[DEBUG] Got status code %v from ImprovMX API on Create for domain %s, success: %v, errors: %v.", resp.Code, d.Get("domain").(string), resp.Success, resp.Errors)

		if resp.Code == 429 {
			log.Printf("[DEBUG] Sleeping for 10 seconds to allow rate limit to recover.")
			time.Sleep(10 * time.Second)
		}

		if resp.Code == 404 {
			log.Printf("[DEBUG] Couldn't find the resource in ImprovMX. Aborting")
			return fmt.Errorf("HTTP response code %d, error text: %s", resp.Code, resp.Errors.Domain)
		}

		if resp.Success {
			return resourceDomainRead(d, meta)
		}
	}
}

func resourceDomainRead(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*Meta)
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	for {
		resp := m.Client.GetDomain(d.Get("domain").(string))

		log.Printf("[DEBUG] Got status code %v from ImprovMX API on Read for domain %s, success: %v, errors: %v.", resp.Code, d.Get("domain").(string), resp.Success, resp.Errors)

		if resp.Code == 429 {
			log.Printf("[DEBUG] Sleeping for 10 seconds to allow rate limit to recover.")
			time.Sleep(10 * time.Second)
		}

		if resp.Code == 404 {
			log.Printf("[DEBUG] Couldn't find the resource in ImprovMX. Aborting")
			return fmt.Errorf("HTTP response code %d, error text: %s", resp.Code, resp.Errors.Domain)
		}

		if resp.Success {
			d.SetId(resp.Domain.Domain)
			d.Set("domain", resp.Domain.Domain)
			d.Set("notification_email", resp.Domain.NotificationEmail)
			d.Set("whitelabel", resp.Domain.Whitelabel)

			return nil
		}
	}
}

func resourceDomainUpdate(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*Meta)

	for {
		resp := m.Client.UpdateDomain(d.Get("domain").(string), d.Get("notification_email").(string), d.Get("whitelabel").(string))

		log.Printf("[DEBUG] Got status code %v from ImprovMX API on Update for domain %s, success: %v, errors: %v.", resp.Code, d.Get("domain").(string), resp.Success, resp.Errors)

		if resp.Code == 429 {
			log.Printf("[DEBUG] Sleeping for 10 seconds to allow rate limit to recover.")
			time.Sleep(10 * time.Second)
		}

		if resp.Code == 404 {
			log.Printf("[DEBUG] Couldn't find the resource in ImprovMX. Aborting")
			return fmt.Errorf("HTTP response code %d, error text: %s", resp.Code, resp.Errors.Domain)
		}

		if resp.Success {
			return resourceDomainRead(d, meta)
		}
	}
}

func resourceDomainDelete(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*Meta)
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	m.Client.DeleteDomain(d.Get("domain").(string))

	return nil
}

func resourceDomainImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("domain", d.Id())
	resourceDomainRead(d, meta)

	return []*schema.ResourceData{d}, nil
}
