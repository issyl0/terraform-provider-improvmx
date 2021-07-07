# `improvmx_domain` Resource

A resource to create ImprovMX domains.

## Example Usage

```hcl
// ImprovMX creates a wildcard email forward on each domain by default.
resource "improvmx_domain" "example" {
  domain = "example.com"
}
```

## Argument Reference

* `domain` - (Required) Name of the domain.
* `notification_email` - (Optional) Email to send notifications to.
* `whitelabel` - (Optional) Parent domain that will be displayed for the DNS settings. Only available on the Enterprise plan.

## Import

Domains can be imported using their name, for example:

```shell
$ terraform import improvmx_domain.example example.com
```