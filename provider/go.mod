module github.com/matchlighter/pulumi-proxmoxve/provider

go 1.14

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20200910230100-328eb4ff41df
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)

replace (
	github.com/Telmate/terraform-provider-proxmox => github.com/matchlighter/terraform-provider-proxmox v0.0.0-20201216064002-1547dfbe6e24
	// github.com/Telmate/proxmox-api-go => github.com/matchlighter/proxmox-api-go v0.0.0-20201102231845-1bf8a10baa24
)

require (
	github.com/Telmate/terraform-provider-proxmox v0.0.0-20201016141324-4d4bf30a4b0d
	github.com/hashicorp/terraform-plugin-sdk v1.14.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.1.0 // indirect
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.12.0
	github.com/pulumi/pulumi/pkg/v2 v2.12.1-0.20201019222817-89c956d18942 // indirect
	github.com/pulumi/pulumi/sdk/v2 v2.12.0
)
