package improvmx

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSMTPCredential() *schema.Resource {
	return &schema.Resource{
		Create: resourceSMTPCredentialCreate,
		Read:   resourceSMTPCredentialRead,
		Update: resourceSMTPCredentialUpdate,
		Delete: resourceSMTPCredentialDelete,

		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceSMTPCredentialCreate(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*Meta)

	for {
		resp := m.Client.CreateSMTPCredential(d.Get("domain").(string), d.Get("username").(string), d.Get("password").(string))

		log.Printf("[DEBUG] Got status code %v from ImprovMX API on Create for SMTP %s@%s, success: %v, errors: %v.", resp.Code, d.Get("username").(string), d.Get("domain").(string), resp.Success, resp.Errors)

		if resp.Code == 429 {
			log.Printf("[DEBUG] Sleeping for 10 seconds to allow rate limit to recover.")
			time.Sleep(10 * time.Second)
		}

		if resp.Code == 404 {
			log.Printf("[DEBUG] Couldn't find the resource in ImprovMX. Aborting")
			return fmt.Errorf("HTTP response code %d, error text: %s", resp.Code, resp.Errors.Domain)
		}

		if resp.Success {
			return resourceSMTPCredentialRead(d, meta)
		}
	}
}

func resourceSMTPCredentialRead(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*Meta)
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	for {
		resp := m.Client.GetSMTPCredential(d.Get("domain").(string))

		log.Printf("[DEBUG] Got status code %v from ImprovMX API on Read for SMTP domain %s, success: %v, errors: %v.", resp.Code, d.Get("domain").(string), resp.Success, resp.Errors)

		if resp.Code == 429 {
			log.Printf("[DEBUG] Sleeping for 10 seconds to allow rate limit to recover.")
			time.Sleep(10 * time.Second)
		}

		if resp.Code == 404 {
			log.Printf("[DEBUG] Couldn't find the resource in ImprovMX. Aborting")
			return fmt.Errorf("HTTP response code %d, error text: %s", resp.Code, resp.Errors.Domain)
		}

		if resp.Success {
			d.SetId(d.Get("domain").(string))
			d.Set("domain", d.Get("domain").(string))
			return nil
		}
	}
}

func resourceSMTPCredentialUpdate(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*Meta)

	for {
		resp := m.Client.UpdateSMTPCredential(d.Get("domain").(string), d.Get("username").(string), d.Get("password").(string))

		log.Printf("[DEBUG] Got status code %v from ImprovMX API on Update for SMTP domain %s@%s, success: %v, errors: %v.", resp.Code, d.Get("username").(string), d.Get("domain").(string), resp.Success, resp.Errors)

		if resp.Code == 429 {
			log.Printf("[DEBUG] Sleeping for 10 seconds to allow rate limit to recover.")
			time.Sleep(10 * time.Second)
		}

		if resp.Code == 404 {
			log.Printf("[DEBUG] Couldn't find the resource in ImprovMX. Aborting")
			return fmt.Errorf("HTTP response code %d, error text: %s", resp.Code, resp.Errors.Domain)
		}

		if resp.Success {
			return resourceSMTPCredentialRead(d, meta)
		}
	}
}

func resourceSMTPCredentialDelete(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*Meta)
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	m.Client.DeleteSMTPCredential(d.Get("domain").(string), d.Get("username").(string))

	return nil
}
