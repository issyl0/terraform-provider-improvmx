# `improvmx_email_forward` Resource

A resource to create ImprovMX email forwards.

## Example Usage

### Single recipient

```hcl
resource "improvmx_email_forward" "example" {
  domain            = "example.com"
  alias_name        = "reception"
  destination_email = "joe@realdomain.com"
}
```

### Multiple recipients

```hcl
resource "improvmx_email_forward" "sales" {
  domain            = "example.com"
  alias_name        = "sales"
  destination_email = "alice@realdomain.com,bob@realdomain.com"
}
```

## Argument Reference

* `domain` - (Required) Name of the domain.
* `alias_name` - (Required) Alias to be used in front of your domain, like "contact", "info", etc.
* `destination_email` - (Required) Email address to forward to.

## Import

Email forwards can be imported using their name, for example:

```shell
$ terraform import improvmx_email_forward.example example.com_hello
```