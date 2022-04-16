terraform {
  required_providers {
    improvmx = {
      source  = "issyl0/improvmx"
    }
  }
}

provider "improvmx" {}

resource "improvmx_domain" "lychee" {
  domain = "lychee.systems"
}
