# ImprovMX Terraform Provider

A community Terraform provider for configuring [ImprovMX](https://improvmx.com) domains and their email forwarding rules.

## Example Usage

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
  token = "YOUR API TOKEN HERE"
}
```

## Argument Reference

* `token` - ImprovMX API token, alternatively set the `IMPROVMX_API_TOKEN` environment variable.
