# terraform-provider-improvmx

A very nascent Terraform provider for configuring [ImprovMX](https://improvmx.com) email forwards. Uses my [ImprovMX Golang API client](https://github.com/issyl0/go-improvmx). Download from the [Terraform Registry](https://registry.terraform.io/providers/issyl0/improvmx/latest).

## Features

- Create a domain (ImprovMX creates a wildcard forward for a domain by default).
- Update a domain (to add/remove whitelabel (Enterprise plans only) and notification email settings).
- Delete a domain.
- Import a domain.
- Create an email forward.
- Delete an email forward.
- Import an email forward.
- Update an email forward (ImprovMX allows updating an email forward to send to more than one address, ie `alice@example.com,bob@example.com`).

## Usage

```hcl
terraform {
  required_providers {
    improvmx = {
      source  = "issyl0/improvmx"
      version = "0.1.0"
    }
  }
}

provider "improvmx" {
  // Set the `IMPROVMX_API_TOKEN` environment variable.
}

resource "improvmx_domain" "example" {
  domain = "example.com"
}

resource "improvmx_email_forward" "hello" {
  domain            = "example.com"
  alias_name        = "hello"
  destination_email = "me@realdomain.com"
}
```
