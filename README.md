# terraform-provider-improvmx

A Terraform provider for configuring ImprovMX email forwards

## Features

- Create an email forward.
- Delete an email forward.

## Coming Soon

- Update an email forward (ImprovMX allows updating an email forward to send to more than one address, ie `alice@example.com,bob@example.com`). Needs support in [the API client](https://github.com/issyl0/go-improvmx) too.
- Import existing email forwards.

## Usage

```hcl
terraform {
  required_providers {
    improvmx = {
      source  = "issyl0/improvmx"
    }
  }
}

provider "improvmx" {
  // Set the `IMPROVMX_API_TOKEN` environment variable.
}

resource "improvmx_email_forward" "hello" {
  domain            = "example.com"
  alias_name        = "hello"
  destination_email = "me@realdomain.com"
}
```
