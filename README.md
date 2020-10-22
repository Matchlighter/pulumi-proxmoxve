# Pulumi Proxmox VE Provider

Work in Progress Pulumi Provider for use with Proxmox VE

## Installing

This package is available in many languages in the standard packaging formats.

### Installing the Plugin
1. Download the appropriate archive file from the Releases page:
   `wget https://github.com/Matchlighter/pulumi-proxmoxve/releases/download/vX.Y.Z/pulumi-resource-proxmoxve-vX.Y.Z-OPERATING_SYSTEM-amd64.tar.gz`
2. Add the plugin to Pulumi:
   `pulumi plugin install resource proxmoxve X.Y.Z -f ./pulumi-resource-proxmoxve-vX.Y.Z-OPERATING_SYSTEM-amd64.tar.gz`

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @matchlighter/pulumi-proxmoxve

or `yarn`:

    $ yarn add @matchlighter/pulumi-proxmoxve

### Python
_*(Not published. I don't have a present need for this. If you do please open an issue.)*_

To use from Python, install using `pip`:

    $ pip install pulumi_proxmoxve

### .NET
_*(Not published. I don't have a present need for this. If you do please open an issue with some instructions for how to publish a .NET package)*_

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/matchlighter/pulumi-proxmoxve/sdk/go/...

## Configuration

Provider configuration is as documented on https://github.com/Telmate/terraform-provider-proxmox
