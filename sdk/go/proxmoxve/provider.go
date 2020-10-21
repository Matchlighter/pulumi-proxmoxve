// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package proxmoxve

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The provider type for the proxmoxve package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil || args.PmApiUrl == nil {
		return nil, errors.New("missing required argument 'PmApiUrl'")
	}
	if args == nil || args.PmPassword == nil {
		return nil, errors.New("missing required argument 'PmPassword'")
	}
	if args == nil || args.PmUser == nil {
		return nil, errors.New("missing required argument 'PmUser'")
	}
	if args == nil {
		args = &ProviderArgs{}
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:proxmoxve", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// https://host.fqdn:8006/api2/json
	PmApiUrl    string                 `pulumi:"pmApiUrl"`
	PmLogEnable *bool                  `pulumi:"pmLogEnable"`
	PmLogFile   *string                `pulumi:"pmLogFile"`
	PmLogLevels map[string]interface{} `pulumi:"pmLogLevels"`
	// OTP 2FA code (if required)
	PmOtp      *string `pulumi:"pmOtp"`
	PmParallel *int    `pulumi:"pmParallel"`
	// secret
	PmPassword    string `pulumi:"pmPassword"`
	PmTimeout     *int   `pulumi:"pmTimeout"`
	PmTlsInsecure *bool  `pulumi:"pmTlsInsecure"`
	// username, maywith with @pam
	PmUser string `pulumi:"pmUser"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// https://host.fqdn:8006/api2/json
	PmApiUrl    pulumi.StringInput
	PmLogEnable pulumi.BoolPtrInput
	PmLogFile   pulumi.StringPtrInput
	PmLogLevels pulumi.MapInput
	// OTP 2FA code (if required)
	PmOtp      pulumi.StringPtrInput
	PmParallel pulumi.IntPtrInput
	// secret
	PmPassword    pulumi.StringInput
	PmTimeout     pulumi.IntPtrInput
	PmTlsInsecure pulumi.BoolPtrInput
	// username, maywith with @pam
	PmUser pulumi.StringInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}
