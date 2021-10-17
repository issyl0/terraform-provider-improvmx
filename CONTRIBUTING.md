Thank you for wanting to contribute to this Terraform provider!

1. Raise an issue first to discuss your new feature, if it's a big one.
1. Fork and clone this repo.
1. Make your changes.
1. Run `go mod download`, then `go build .` in the root directory. This should install all dependencies and give a `terraform-provider-improvmx` binary.
1. Configure your local Terraform to use the development provider rather than the released version on the Registry. Make a `~/.terraformrc` file pointing to the filepath of your dev binary:
    ```hcl
    provider_installation {
      dev_overrides {
        "issyl0/improvmx" = "/full/path/to/terraform-provider-improvmx/directory"
      }

      direct {}
    }
    ```
1. Write some Terraform config for your new feature, then iterate against your ImprovMX account until it does what you intend.
1. Submit a PR. Thank you! 🙇
