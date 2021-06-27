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