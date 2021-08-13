# `improvmx_smtp_credential` Resource

A resource to create ImprovMX domain SMTP credentials.

## Example Usage

```hcl
resource "improvmx_smtp_credential" "example" {
  domain   = "example.com"
  username = "example"
  password = var.password
}
```

## Argument Reference

* `domain` - (Required) Name of the domain.
* `username` - (Required) Username of the SMTP sender.
* `password` - (Required) Password for the SMTP sending account. (Will always be hidden in `plan`.)
