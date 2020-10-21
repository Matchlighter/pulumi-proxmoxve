# Pulumi Proxmox VE Provider

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @matchlighter/pulumi-proxmoxve

or `yarn`:

    $ yarn add @matchlighter/pulumi-proxmoxve

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_proxmoxve

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/matchlighter/pulumi-proxmoxve/sdk/go/...

## Configuration

The following configuration points are available for the `proxmoxve` provider:

- `proxmoxve:apiKey` (environment: `PROXMOXVE_API_KEY`) - the API key for `proxmoxve`
- `proxmoxve:region` (environment: `PROXMOXVE_REGION`) - the region in which to deploy resources

## Reference

For detailed reference documentation, please visit [the API docs][1].


[1]: https://www.pulumi.com/docs/reference/pkg/x/
