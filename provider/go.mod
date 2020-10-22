module github.com/matchlighter/pulumi-proxmoxve/provider

go 1.14

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible

replace github.com/Telmate/terraform-provider-proxmox => github.com/matchlighter/terraform-provider-proxmox v0.0.0-20201022044959-84553b1dce1f

require (
	github.com/Telmate/terraform-provider-proxmox v0.0.0-20201016141324-4d4bf30a4b0d
	github.com/hashicorp/terraform-plugin-sdk v1.13.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.11.0
	github.com/pulumi/pulumi/sdk/v2 v2.12.0
)
