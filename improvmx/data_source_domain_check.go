package improvmx

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceDomainCheck() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDomainCheckRead,
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"records_are_valid": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"record_mx_is_valid": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"record_mx_expected_values": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"record_mx_actual_values": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"record_spf_is_valid": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"record_spf_expected_value": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"record_spf_actual_value": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceDomainCheckRead(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*Meta)
	domainString := d.Get("domain").(string)

	response := m.Client.GetDomainCheck(domainString)

	log.Printf(
		"[DEBUG] Got status code %v from ImprovMX API on Read for domain check %s, success: %v, errors: %v.",
		response.Code,
		domainString,
		response.Success,
		response.Errors,
	)

	if response.Code == 429 {
		log.Printf("[DEBUG] Rate limit hit. Too many requests. Aborting.")
		return fmt.Errorf("Rate limit hit. Too many requests. Aborting.")
	}

	if response.Code == 404 {
		log.Printf("[DEBUG] Couldn't find the resource in ImprovMX. Aborting.")
		return fmt.Errorf("HTTP response code %d, error text: %s", response.Code, response.Error)
	}

	if response.Success == false {
		log.Printf("[DEBUG] Request was not successful. Aborting.")
		return fmt.Errorf("HTTP response code %d, error text: %s", response.Code, response.Error)
	}

	d.Set("records_are_valid", response.Records.Valid)

	d.Set("record_mx_is_valid", response.Records.Mx.Valid)
	d.Set("record_mx_expected_values", response.Records.Mx.Expected)
	d.Set("record_mx_actual_values", response.Records.Mx.Values)

	d.Set("record_spf_is_valid", response.Records.Spf.Valid)
	d.Set("record_spf_expected_value", response.Records.Spf.Expected)
	d.Set("record_spf_actual_value", response.Records.Spf.Values)

	d.SetId(domainString)

	return nil
}
