# `improvmx_domain_check` Data Source

A data source to read domain check status.

## Example Usage

```hcl
data "improvmx_domain_check" "example_com" {
  domain = "example.com"
}
```

## Argument Reference

* `domain` - (Required) Name of the domain.

## Attribute Reference

* `id` - (int) Unique ID
* `domain` - (string) Name of the domain
* `records_are_valid` - (bool) Whether all records are valid or not
* `record_mx_is_valid` - (bool) Whether mx record is valid or not
* `record_mx_expected_values` - (list) Expected mx records
* `record_mx_actual_values` - (list) Actual mx records
* `record_spf_is_valid` - (bool) Whether spf record is valid or not
* `record_spf_expected_value` - (list) Expected spf record
* `record_spf_actual_value` - (list) Actual spf record
