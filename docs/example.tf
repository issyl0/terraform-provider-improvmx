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

// ImprovMX creates a wildcard email forward on each domain by default.
resource "improvmx_domain" "example" {
  domain = "example.com"
}

resource "improvmx_email_forward" "hello" {
  domain            = "example.com"
  alias_name        = "hello"
  destination_email = "me@realdomain.com,another@realdomain.com"
}