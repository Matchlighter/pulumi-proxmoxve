# Pulumi Proxmox VE Provider

Work in Progress Pulumi Provider for use with Proxmox VE

1. Follow the steps above to verify the program runs successfully.

## Add End-to-end Testing

### Installing the Plugin
1. Download the appropriate archive file from the Releases page:
   `wget https://github.com/Matchlighter/pulumi-proxmoxve/releases/download/vX.Y.Z/pulumi-resource-proxmoxve-vX.Y.Z-OPERATING_SYSTEM-amd64.tar.gz`
2. Add the plugin to Pulumi:
   `pulumi plugin install resource proxmoxve X.Y.Z -f ./pulumi-resource-proxmoxve-vX.Y.Z-OPERATING_SYSTEM-amd64.tar.gz`

### Node.js (Java/TypeScript)

1. Add code to `examples_nodejs_test.go` to call the example you created, e.g.:

    $ npm install @matchlighter/pulumi-proxmoxve

1. Add a similar function for each example that you want to run in an integration test.  For examples written in other languages, create similar files for `examples_${LANGUAGE}_test.go`.

    $ yarn add @matchlighter/pulumi-proxmoxve

### Python
_*(Not published. I don't have a present need for this. If you do please open an issue.)*_

    You can also run each test file separately via test tags:

    $ pip install pulumi_proxmoxve

### .NET
_*(Not published. I don't have a present need for this. If you do please open an issue with some instructions for how to publish a .NET package)*_

## Configuring CI with GitHub Actions

In this section, we'll add the necessary configuration to work with GitHub Actions for Pulumi's standard CI/CD workflows for providers.

    $ go get github.com/matchlighter/pulumi-proxmoxve/sdk/go/...

1. Ensure that any required secrets are present as repository-level [secrets](https://docs.github.com/en/actions/security-guides/encrypted-secrets) in GitHub.  These will be used by the integration tests during the CI/CD process.

Provider configuration is as documented on https://github.com/Telmate/terraform-provider-proxmox
