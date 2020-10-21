// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package proxmoxve

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Lxc_containerFeature struct {
	Fuse    *bool   `pulumi:"fuse"`
	Keyctl  *bool   `pulumi:"keyctl"`
	Mount   *string `pulumi:"mount"`
	Nesting *bool   `pulumi:"nesting"`
}

// Lxc_containerFeatureInput is an input type that accepts Lxc_containerFeatureArgs and Lxc_containerFeatureOutput values.
// You can construct a concrete instance of `Lxc_containerFeatureInput` via:
//
//          Lxc_containerFeatureArgs{...}
type Lxc_containerFeatureInput interface {
	pulumi.Input

	ToLxc_containerFeatureOutput() Lxc_containerFeatureOutput
	ToLxc_containerFeatureOutputWithContext(context.Context) Lxc_containerFeatureOutput
}

type Lxc_containerFeatureArgs struct {
	Fuse    pulumi.BoolPtrInput   `pulumi:"fuse"`
	Keyctl  pulumi.BoolPtrInput   `pulumi:"keyctl"`
	Mount   pulumi.StringPtrInput `pulumi:"mount"`
	Nesting pulumi.BoolPtrInput   `pulumi:"nesting"`
}

func (Lxc_containerFeatureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Lxc_containerFeature)(nil)).Elem()
}

func (i Lxc_containerFeatureArgs) ToLxc_containerFeatureOutput() Lxc_containerFeatureOutput {
	return i.ToLxc_containerFeatureOutputWithContext(context.Background())
}

func (i Lxc_containerFeatureArgs) ToLxc_containerFeatureOutputWithContext(ctx context.Context) Lxc_containerFeatureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Lxc_containerFeatureOutput)
}

// Lxc_containerFeatureArrayInput is an input type that accepts Lxc_containerFeatureArray and Lxc_containerFeatureArrayOutput values.
// You can construct a concrete instance of `Lxc_containerFeatureArrayInput` via:
//
//          Lxc_containerFeatureArray{ Lxc_containerFeatureArgs{...} }
type Lxc_containerFeatureArrayInput interface {
	pulumi.Input

	ToLxc_containerFeatureArrayOutput() Lxc_containerFeatureArrayOutput
	ToLxc_containerFeatureArrayOutputWithContext(context.Context) Lxc_containerFeatureArrayOutput
}

type Lxc_containerFeatureArray []Lxc_containerFeatureInput

func (Lxc_containerFeatureArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Lxc_containerFeature)(nil)).Elem()
}

func (i Lxc_containerFeatureArray) ToLxc_containerFeatureArrayOutput() Lxc_containerFeatureArrayOutput {
	return i.ToLxc_containerFeatureArrayOutputWithContext(context.Background())
}

func (i Lxc_containerFeatureArray) ToLxc_containerFeatureArrayOutputWithContext(ctx context.Context) Lxc_containerFeatureArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Lxc_containerFeatureArrayOutput)
}

type Lxc_containerFeatureOutput struct{ *pulumi.OutputState }

func (Lxc_containerFeatureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Lxc_containerFeature)(nil)).Elem()
}

func (o Lxc_containerFeatureOutput) ToLxc_containerFeatureOutput() Lxc_containerFeatureOutput {
	return o
}

func (o Lxc_containerFeatureOutput) ToLxc_containerFeatureOutputWithContext(ctx context.Context) Lxc_containerFeatureOutput {
	return o
}

func (o Lxc_containerFeatureOutput) Fuse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Lxc_containerFeature) *bool { return v.Fuse }).(pulumi.BoolPtrOutput)
}

func (o Lxc_containerFeatureOutput) Keyctl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Lxc_containerFeature) *bool { return v.Keyctl }).(pulumi.BoolPtrOutput)
}

func (o Lxc_containerFeatureOutput) Mount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Lxc_containerFeature) *string { return v.Mount }).(pulumi.StringPtrOutput)
}

func (o Lxc_containerFeatureOutput) Nesting() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Lxc_containerFeature) *bool { return v.Nesting }).(pulumi.BoolPtrOutput)
}

type Lxc_containerFeatureArrayOutput struct{ *pulumi.OutputState }

func (Lxc_containerFeatureArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Lxc_containerFeature)(nil)).Elem()
}

func (o Lxc_containerFeatureArrayOutput) ToLxc_containerFeatureArrayOutput() Lxc_containerFeatureArrayOutput {
	return o
}

func (o Lxc_containerFeatureArrayOutput) ToLxc_containerFeatureArrayOutputWithContext(ctx context.Context) Lxc_containerFeatureArrayOutput {
	return o
}

func (o Lxc_containerFeatureArrayOutput) Index(i pulumi.IntInput) Lxc_containerFeatureOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Lxc_containerFeature {
		return vs[0].([]Lxc_containerFeature)[vs[1].(int)]
	}).(Lxc_containerFeatureOutput)
}

type Lxc_containerMountpoint struct {
	Acl       *bool  `pulumi:"acl"`
	Backup    *bool  `pulumi:"backup"`
	Mp        string `pulumi:"mp"`
	Quota     *bool  `pulumi:"quota"`
	Replicate *bool  `pulumi:"replicate"`
	Shared    *bool  `pulumi:"shared"`
	Size      *int   `pulumi:"size"`
	Volume    string `pulumi:"volume"`
}

// Lxc_containerMountpointInput is an input type that accepts Lxc_containerMountpointArgs and Lxc_containerMountpointOutput values.
// You can construct a concrete instance of `Lxc_containerMountpointInput` via:
//
//          Lxc_containerMountpointArgs{...}
type Lxc_containerMountpointInput interface {
	pulumi.Input

	ToLxc_containerMountpointOutput() Lxc_containerMountpointOutput
	ToLxc_containerMountpointOutputWithContext(context.Context) Lxc_containerMountpointOutput
}

type Lxc_containerMountpointArgs struct {
	Acl       pulumi.BoolPtrInput `pulumi:"acl"`
	Backup    pulumi.BoolPtrInput `pulumi:"backup"`
	Mp        pulumi.StringInput  `pulumi:"mp"`
	Quota     pulumi.BoolPtrInput `pulumi:"quota"`
	Replicate pulumi.BoolPtrInput `pulumi:"replicate"`
	Shared    pulumi.BoolPtrInput `pulumi:"shared"`
	Size      pulumi.IntPtrInput  `pulumi:"size"`
	Volume    pulumi.StringInput  `pulumi:"volume"`
}

func (Lxc_containerMountpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Lxc_containerMountpoint)(nil)).Elem()
}

func (i Lxc_containerMountpointArgs) ToLxc_containerMountpointOutput() Lxc_containerMountpointOutput {
	return i.ToLxc_containerMountpointOutputWithContext(context.Background())
}

func (i Lxc_containerMountpointArgs) ToLxc_containerMountpointOutputWithContext(ctx context.Context) Lxc_containerMountpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Lxc_containerMountpointOutput)
}

// Lxc_containerMountpointArrayInput is an input type that accepts Lxc_containerMountpointArray and Lxc_containerMountpointArrayOutput values.
// You can construct a concrete instance of `Lxc_containerMountpointArrayInput` via:
//
//          Lxc_containerMountpointArray{ Lxc_containerMountpointArgs{...} }
type Lxc_containerMountpointArrayInput interface {
	pulumi.Input

	ToLxc_containerMountpointArrayOutput() Lxc_containerMountpointArrayOutput
	ToLxc_containerMountpointArrayOutputWithContext(context.Context) Lxc_containerMountpointArrayOutput
}

type Lxc_containerMountpointArray []Lxc_containerMountpointInput

func (Lxc_containerMountpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Lxc_containerMountpoint)(nil)).Elem()
}

func (i Lxc_containerMountpointArray) ToLxc_containerMountpointArrayOutput() Lxc_containerMountpointArrayOutput {
	return i.ToLxc_containerMountpointArrayOutputWithContext(context.Background())
}

func (i Lxc_containerMountpointArray) ToLxc_containerMountpointArrayOutputWithContext(ctx context.Context) Lxc_containerMountpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Lxc_containerMountpointArrayOutput)
}

type Lxc_containerMountpointOutput struct{ *pulumi.OutputState }

func (Lxc_containerMountpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Lxc_containerMountpoint)(nil)).Elem()
}

func (o Lxc_containerMountpointOutput) ToLxc_containerMountpointOutput() Lxc_containerMountpointOutput {
	return o
}

func (o Lxc_containerMountpointOutput) ToLxc_containerMountpointOutputWithContext(ctx context.Context) Lxc_containerMountpointOutput {
	return o
}

func (o Lxc_containerMountpointOutput) Acl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Lxc_containerMountpoint) *bool { return v.Acl }).(pulumi.BoolPtrOutput)
}

func (o Lxc_containerMountpointOutput) Backup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Lxc_containerMountpoint) *bool { return v.Backup }).(pulumi.BoolPtrOutput)
}

func (o Lxc_containerMountpointOutput) Mp() pulumi.StringOutput {
	return o.ApplyT(func(v Lxc_containerMountpoint) string { return v.Mp }).(pulumi.StringOutput)
}

func (o Lxc_containerMountpointOutput) Quota() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Lxc_containerMountpoint) *bool { return v.Quota }).(pulumi.BoolPtrOutput)
}

func (o Lxc_containerMountpointOutput) Replicate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Lxc_containerMountpoint) *bool { return v.Replicate }).(pulumi.BoolPtrOutput)
}

func (o Lxc_containerMountpointOutput) Shared() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Lxc_containerMountpoint) *bool { return v.Shared }).(pulumi.BoolPtrOutput)
}

func (o Lxc_containerMountpointOutput) Size() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Lxc_containerMountpoint) *int { return v.Size }).(pulumi.IntPtrOutput)
}

func (o Lxc_containerMountpointOutput) Volume() pulumi.StringOutput {
	return o.ApplyT(func(v Lxc_containerMountpoint) string { return v.Volume }).(pulumi.StringOutput)
}

type Lxc_containerMountpointArrayOutput struct{ *pulumi.OutputState }

func (Lxc_containerMountpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Lxc_containerMountpoint)(nil)).Elem()
}

func (o Lxc_containerMountpointArrayOutput) ToLxc_containerMountpointArrayOutput() Lxc_containerMountpointArrayOutput {
	return o
}

func (o Lxc_containerMountpointArrayOutput) ToLxc_containerMountpointArrayOutputWithContext(ctx context.Context) Lxc_containerMountpointArrayOutput {
	return o
}

func (o Lxc_containerMountpointArrayOutput) Index(i pulumi.IntInput) Lxc_containerMountpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Lxc_containerMountpoint {
		return vs[0].([]Lxc_containerMountpoint)[vs[1].(int)]
	}).(Lxc_containerMountpointOutput)
}

type Lxc_containerNetwork struct {
	Bridge   *string `pulumi:"bridge"`
	Firewall *bool   `pulumi:"firewall"`
	Gw       *string `pulumi:"gw"`
	Gw6      *string `pulumi:"gw6"`
	Hwaddr   *string `pulumi:"hwaddr"`
	Ip       *string `pulumi:"ip"`
	Ip6      *string `pulumi:"ip6"`
	Mtu      *string `pulumi:"mtu"`
	Name     string  `pulumi:"name"`
	Rate     *int    `pulumi:"rate"`
	Tag      *int    `pulumi:"tag"`
	Trunks   *string `pulumi:"trunks"`
	Type     *string `pulumi:"type"`
}

// Lxc_containerNetworkInput is an input type that accepts Lxc_containerNetworkArgs and Lxc_containerNetworkOutput values.
// You can construct a concrete instance of `Lxc_containerNetworkInput` via:
//
//          Lxc_containerNetworkArgs{...}
type Lxc_containerNetworkInput interface {
	pulumi.Input

	ToLxc_containerNetworkOutput() Lxc_containerNetworkOutput
	ToLxc_containerNetworkOutputWithContext(context.Context) Lxc_containerNetworkOutput
}

type Lxc_containerNetworkArgs struct {
	Bridge   pulumi.StringPtrInput `pulumi:"bridge"`
	Firewall pulumi.BoolPtrInput   `pulumi:"firewall"`
	Gw       pulumi.StringPtrInput `pulumi:"gw"`
	Gw6      pulumi.StringPtrInput `pulumi:"gw6"`
	Hwaddr   pulumi.StringPtrInput `pulumi:"hwaddr"`
	Ip       pulumi.StringPtrInput `pulumi:"ip"`
	Ip6      pulumi.StringPtrInput `pulumi:"ip6"`
	Mtu      pulumi.StringPtrInput `pulumi:"mtu"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Rate     pulumi.IntPtrInput    `pulumi:"rate"`
	Tag      pulumi.IntPtrInput    `pulumi:"tag"`
	Trunks   pulumi.StringPtrInput `pulumi:"trunks"`
	Type     pulumi.StringPtrInput `pulumi:"type"`
}

func (Lxc_containerNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Lxc_containerNetwork)(nil)).Elem()
}

func (i Lxc_containerNetworkArgs) ToLxc_containerNetworkOutput() Lxc_containerNetworkOutput {
	return i.ToLxc_containerNetworkOutputWithContext(context.Background())
}

func (i Lxc_containerNetworkArgs) ToLxc_containerNetworkOutputWithContext(ctx context.Context) Lxc_containerNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Lxc_containerNetworkOutput)
}

// Lxc_containerNetworkArrayInput is an input type that accepts Lxc_containerNetworkArray and Lxc_containerNetworkArrayOutput values.
// You can construct a concrete instance of `Lxc_containerNetworkArrayInput` via:
//
//          Lxc_containerNetworkArray{ Lxc_containerNetworkArgs{...} }
type Lxc_containerNetworkArrayInput interface {
	pulumi.Input

	ToLxc_containerNetworkArrayOutput() Lxc_containerNetworkArrayOutput
	ToLxc_containerNetworkArrayOutputWithContext(context.Context) Lxc_containerNetworkArrayOutput
}

type Lxc_containerNetworkArray []Lxc_containerNetworkInput

func (Lxc_containerNetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Lxc_containerNetwork)(nil)).Elem()
}

func (i Lxc_containerNetworkArray) ToLxc_containerNetworkArrayOutput() Lxc_containerNetworkArrayOutput {
	return i.ToLxc_containerNetworkArrayOutputWithContext(context.Background())
}

func (i Lxc_containerNetworkArray) ToLxc_containerNetworkArrayOutputWithContext(ctx context.Context) Lxc_containerNetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Lxc_containerNetworkArrayOutput)
}

type Lxc_containerNetworkOutput struct{ *pulumi.OutputState }

func (Lxc_containerNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Lxc_containerNetwork)(nil)).Elem()
}

func (o Lxc_containerNetworkOutput) ToLxc_containerNetworkOutput() Lxc_containerNetworkOutput {
	return o
}

func (o Lxc_containerNetworkOutput) ToLxc_containerNetworkOutputWithContext(ctx context.Context) Lxc_containerNetworkOutput {
	return o
}

func (o Lxc_containerNetworkOutput) Bridge() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Lxc_containerNetwork) *string { return v.Bridge }).(pulumi.StringPtrOutput)
}

func (o Lxc_containerNetworkOutput) Firewall() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Lxc_containerNetwork) *bool { return v.Firewall }).(pulumi.BoolPtrOutput)
}

func (o Lxc_containerNetworkOutput) Gw() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Lxc_containerNetwork) *string { return v.Gw }).(pulumi.StringPtrOutput)
}

func (o Lxc_containerNetworkOutput) Gw6() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Lxc_containerNetwork) *string { return v.Gw6 }).(pulumi.StringPtrOutput)
}

func (o Lxc_containerNetworkOutput) Hwaddr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Lxc_containerNetwork) *string { return v.Hwaddr }).(pulumi.StringPtrOutput)
}

func (o Lxc_containerNetworkOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Lxc_containerNetwork) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

func (o Lxc_containerNetworkOutput) Ip6() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Lxc_containerNetwork) *string { return v.Ip6 }).(pulumi.StringPtrOutput)
}

func (o Lxc_containerNetworkOutput) Mtu() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Lxc_containerNetwork) *string { return v.Mtu }).(pulumi.StringPtrOutput)
}

func (o Lxc_containerNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Lxc_containerNetwork) string { return v.Name }).(pulumi.StringOutput)
}

func (o Lxc_containerNetworkOutput) Rate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Lxc_containerNetwork) *int { return v.Rate }).(pulumi.IntPtrOutput)
}

func (o Lxc_containerNetworkOutput) Tag() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Lxc_containerNetwork) *int { return v.Tag }).(pulumi.IntPtrOutput)
}

func (o Lxc_containerNetworkOutput) Trunks() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Lxc_containerNetwork) *string { return v.Trunks }).(pulumi.StringPtrOutput)
}

func (o Lxc_containerNetworkOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Lxc_containerNetwork) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type Lxc_containerNetworkArrayOutput struct{ *pulumi.OutputState }

func (Lxc_containerNetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Lxc_containerNetwork)(nil)).Elem()
}

func (o Lxc_containerNetworkArrayOutput) ToLxc_containerNetworkArrayOutput() Lxc_containerNetworkArrayOutput {
	return o
}

func (o Lxc_containerNetworkArrayOutput) ToLxc_containerNetworkArrayOutputWithContext(ctx context.Context) Lxc_containerNetworkArrayOutput {
	return o
}

func (o Lxc_containerNetworkArrayOutput) Index(i pulumi.IntInput) Lxc_containerNetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Lxc_containerNetwork {
		return vs[0].([]Lxc_containerNetwork)[vs[1].(int)]
	}).(Lxc_containerNetworkOutput)
}

type Qemu_vmDisk struct {
	Backup      *bool   `pulumi:"backup"`
	Cache       *string `pulumi:"cache"`
	Discard     *string `pulumi:"discard"`
	File        *string `pulumi:"file"`
	Format      *string `pulumi:"format"`
	Iothread    *bool   `pulumi:"iothread"`
	Mbps        *int    `pulumi:"mbps"`
	MbpsRd      *int    `pulumi:"mbpsRd"`
	MbpsRdMax   *int    `pulumi:"mbpsRdMax"`
	MbpsWr      *int    `pulumi:"mbpsWr"`
	MbpsWrMax   *int    `pulumi:"mbpsWrMax"`
	Media       *string `pulumi:"media"`
	Replicate   *bool   `pulumi:"replicate"`
	Size        string  `pulumi:"size"`
	Ssd         *bool   `pulumi:"ssd"`
	Storage     string  `pulumi:"storage"`
	StorageType *string `pulumi:"storageType"`
	Type        string  `pulumi:"type"`
}

// Qemu_vmDiskInput is an input type that accepts Qemu_vmDiskArgs and Qemu_vmDiskOutput values.
// You can construct a concrete instance of `Qemu_vmDiskInput` via:
//
//          Qemu_vmDiskArgs{...}
type Qemu_vmDiskInput interface {
	pulumi.Input

	ToQemu_vmDiskOutput() Qemu_vmDiskOutput
	ToQemu_vmDiskOutputWithContext(context.Context) Qemu_vmDiskOutput
}

type Qemu_vmDiskArgs struct {
	Backup      pulumi.BoolPtrInput   `pulumi:"backup"`
	Cache       pulumi.StringPtrInput `pulumi:"cache"`
	Discard     pulumi.StringPtrInput `pulumi:"discard"`
	File        pulumi.StringPtrInput `pulumi:"file"`
	Format      pulumi.StringPtrInput `pulumi:"format"`
	Iothread    pulumi.BoolPtrInput   `pulumi:"iothread"`
	Mbps        pulumi.IntPtrInput    `pulumi:"mbps"`
	MbpsRd      pulumi.IntPtrInput    `pulumi:"mbpsRd"`
	MbpsRdMax   pulumi.IntPtrInput    `pulumi:"mbpsRdMax"`
	MbpsWr      pulumi.IntPtrInput    `pulumi:"mbpsWr"`
	MbpsWrMax   pulumi.IntPtrInput    `pulumi:"mbpsWrMax"`
	Media       pulumi.StringPtrInput `pulumi:"media"`
	Replicate   pulumi.BoolPtrInput   `pulumi:"replicate"`
	Size        pulumi.StringInput    `pulumi:"size"`
	Ssd         pulumi.BoolPtrInput   `pulumi:"ssd"`
	Storage     pulumi.StringInput    `pulumi:"storage"`
	StorageType pulumi.StringPtrInput `pulumi:"storageType"`
	Type        pulumi.StringInput    `pulumi:"type"`
}

func (Qemu_vmDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Qemu_vmDisk)(nil)).Elem()
}

func (i Qemu_vmDiskArgs) ToQemu_vmDiskOutput() Qemu_vmDiskOutput {
	return i.ToQemu_vmDiskOutputWithContext(context.Background())
}

func (i Qemu_vmDiskArgs) ToQemu_vmDiskOutputWithContext(ctx context.Context) Qemu_vmDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Qemu_vmDiskOutput)
}

// Qemu_vmDiskArrayInput is an input type that accepts Qemu_vmDiskArray and Qemu_vmDiskArrayOutput values.
// You can construct a concrete instance of `Qemu_vmDiskArrayInput` via:
//
//          Qemu_vmDiskArray{ Qemu_vmDiskArgs{...} }
type Qemu_vmDiskArrayInput interface {
	pulumi.Input

	ToQemu_vmDiskArrayOutput() Qemu_vmDiskArrayOutput
	ToQemu_vmDiskArrayOutputWithContext(context.Context) Qemu_vmDiskArrayOutput
}

type Qemu_vmDiskArray []Qemu_vmDiskInput

func (Qemu_vmDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Qemu_vmDisk)(nil)).Elem()
}

func (i Qemu_vmDiskArray) ToQemu_vmDiskArrayOutput() Qemu_vmDiskArrayOutput {
	return i.ToQemu_vmDiskArrayOutputWithContext(context.Background())
}

func (i Qemu_vmDiskArray) ToQemu_vmDiskArrayOutputWithContext(ctx context.Context) Qemu_vmDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Qemu_vmDiskArrayOutput)
}

type Qemu_vmDiskOutput struct{ *pulumi.OutputState }

func (Qemu_vmDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Qemu_vmDisk)(nil)).Elem()
}

func (o Qemu_vmDiskOutput) ToQemu_vmDiskOutput() Qemu_vmDiskOutput {
	return o
}

func (o Qemu_vmDiskOutput) ToQemu_vmDiskOutputWithContext(ctx context.Context) Qemu_vmDiskOutput {
	return o
}

func (o Qemu_vmDiskOutput) Backup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Qemu_vmDisk) *bool { return v.Backup }).(pulumi.BoolPtrOutput)
}

func (o Qemu_vmDiskOutput) Cache() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Qemu_vmDisk) *string { return v.Cache }).(pulumi.StringPtrOutput)
}

func (o Qemu_vmDiskOutput) Discard() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Qemu_vmDisk) *string { return v.Discard }).(pulumi.StringPtrOutput)
}

func (o Qemu_vmDiskOutput) File() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Qemu_vmDisk) *string { return v.File }).(pulumi.StringPtrOutput)
}

func (o Qemu_vmDiskOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Qemu_vmDisk) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o Qemu_vmDiskOutput) Iothread() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Qemu_vmDisk) *bool { return v.Iothread }).(pulumi.BoolPtrOutput)
}

func (o Qemu_vmDiskOutput) Mbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Qemu_vmDisk) *int { return v.Mbps }).(pulumi.IntPtrOutput)
}

func (o Qemu_vmDiskOutput) MbpsRd() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Qemu_vmDisk) *int { return v.MbpsRd }).(pulumi.IntPtrOutput)
}

func (o Qemu_vmDiskOutput) MbpsRdMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Qemu_vmDisk) *int { return v.MbpsRdMax }).(pulumi.IntPtrOutput)
}

func (o Qemu_vmDiskOutput) MbpsWr() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Qemu_vmDisk) *int { return v.MbpsWr }).(pulumi.IntPtrOutput)
}

func (o Qemu_vmDiskOutput) MbpsWrMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Qemu_vmDisk) *int { return v.MbpsWrMax }).(pulumi.IntPtrOutput)
}

func (o Qemu_vmDiskOutput) Media() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Qemu_vmDisk) *string { return v.Media }).(pulumi.StringPtrOutput)
}

func (o Qemu_vmDiskOutput) Replicate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Qemu_vmDisk) *bool { return v.Replicate }).(pulumi.BoolPtrOutput)
}

func (o Qemu_vmDiskOutput) Size() pulumi.StringOutput {
	return o.ApplyT(func(v Qemu_vmDisk) string { return v.Size }).(pulumi.StringOutput)
}

func (o Qemu_vmDiskOutput) Ssd() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Qemu_vmDisk) *bool { return v.Ssd }).(pulumi.BoolPtrOutput)
}

func (o Qemu_vmDiskOutput) Storage() pulumi.StringOutput {
	return o.ApplyT(func(v Qemu_vmDisk) string { return v.Storage }).(pulumi.StringOutput)
}

func (o Qemu_vmDiskOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Qemu_vmDisk) *string { return v.StorageType }).(pulumi.StringPtrOutput)
}

func (o Qemu_vmDiskOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v Qemu_vmDisk) string { return v.Type }).(pulumi.StringOutput)
}

type Qemu_vmDiskArrayOutput struct{ *pulumi.OutputState }

func (Qemu_vmDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Qemu_vmDisk)(nil)).Elem()
}

func (o Qemu_vmDiskArrayOutput) ToQemu_vmDiskArrayOutput() Qemu_vmDiskArrayOutput {
	return o
}

func (o Qemu_vmDiskArrayOutput) ToQemu_vmDiskArrayOutputWithContext(ctx context.Context) Qemu_vmDiskArrayOutput {
	return o
}

func (o Qemu_vmDiskArrayOutput) Index(i pulumi.IntInput) Qemu_vmDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Qemu_vmDisk {
		return vs[0].([]Qemu_vmDisk)[vs[1].(int)]
	}).(Qemu_vmDiskOutput)
}

type Qemu_vmNetwork struct {
	Bridge   *string `pulumi:"bridge"`
	Firewall *bool   `pulumi:"firewall"`
	LinkDown *bool   `pulumi:"linkDown"`
	Macaddr  *string `pulumi:"macaddr"`
	Model    string  `pulumi:"model"`
	Queues   *int    `pulumi:"queues"`
	Rate     *int    `pulumi:"rate"`
	Tag      *int    `pulumi:"tag"`
}

// Qemu_vmNetworkInput is an input type that accepts Qemu_vmNetworkArgs and Qemu_vmNetworkOutput values.
// You can construct a concrete instance of `Qemu_vmNetworkInput` via:
//
//          Qemu_vmNetworkArgs{...}
type Qemu_vmNetworkInput interface {
	pulumi.Input

	ToQemu_vmNetworkOutput() Qemu_vmNetworkOutput
	ToQemu_vmNetworkOutputWithContext(context.Context) Qemu_vmNetworkOutput
}

type Qemu_vmNetworkArgs struct {
	Bridge   pulumi.StringPtrInput `pulumi:"bridge"`
	Firewall pulumi.BoolPtrInput   `pulumi:"firewall"`
	LinkDown pulumi.BoolPtrInput   `pulumi:"linkDown"`
	Macaddr  pulumi.StringPtrInput `pulumi:"macaddr"`
	Model    pulumi.StringInput    `pulumi:"model"`
	Queues   pulumi.IntPtrInput    `pulumi:"queues"`
	Rate     pulumi.IntPtrInput    `pulumi:"rate"`
	Tag      pulumi.IntPtrInput    `pulumi:"tag"`
}

func (Qemu_vmNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Qemu_vmNetwork)(nil)).Elem()
}

func (i Qemu_vmNetworkArgs) ToQemu_vmNetworkOutput() Qemu_vmNetworkOutput {
	return i.ToQemu_vmNetworkOutputWithContext(context.Background())
}

func (i Qemu_vmNetworkArgs) ToQemu_vmNetworkOutputWithContext(ctx context.Context) Qemu_vmNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Qemu_vmNetworkOutput)
}

// Qemu_vmNetworkArrayInput is an input type that accepts Qemu_vmNetworkArray and Qemu_vmNetworkArrayOutput values.
// You can construct a concrete instance of `Qemu_vmNetworkArrayInput` via:
//
//          Qemu_vmNetworkArray{ Qemu_vmNetworkArgs{...} }
type Qemu_vmNetworkArrayInput interface {
	pulumi.Input

	ToQemu_vmNetworkArrayOutput() Qemu_vmNetworkArrayOutput
	ToQemu_vmNetworkArrayOutputWithContext(context.Context) Qemu_vmNetworkArrayOutput
}

type Qemu_vmNetworkArray []Qemu_vmNetworkInput

func (Qemu_vmNetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Qemu_vmNetwork)(nil)).Elem()
}

func (i Qemu_vmNetworkArray) ToQemu_vmNetworkArrayOutput() Qemu_vmNetworkArrayOutput {
	return i.ToQemu_vmNetworkArrayOutputWithContext(context.Background())
}

func (i Qemu_vmNetworkArray) ToQemu_vmNetworkArrayOutputWithContext(ctx context.Context) Qemu_vmNetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Qemu_vmNetworkArrayOutput)
}

type Qemu_vmNetworkOutput struct{ *pulumi.OutputState }

func (Qemu_vmNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Qemu_vmNetwork)(nil)).Elem()
}

func (o Qemu_vmNetworkOutput) ToQemu_vmNetworkOutput() Qemu_vmNetworkOutput {
	return o
}

func (o Qemu_vmNetworkOutput) ToQemu_vmNetworkOutputWithContext(ctx context.Context) Qemu_vmNetworkOutput {
	return o
}

func (o Qemu_vmNetworkOutput) Bridge() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Qemu_vmNetwork) *string { return v.Bridge }).(pulumi.StringPtrOutput)
}

func (o Qemu_vmNetworkOutput) Firewall() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Qemu_vmNetwork) *bool { return v.Firewall }).(pulumi.BoolPtrOutput)
}

func (o Qemu_vmNetworkOutput) LinkDown() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Qemu_vmNetwork) *bool { return v.LinkDown }).(pulumi.BoolPtrOutput)
}

func (o Qemu_vmNetworkOutput) Macaddr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Qemu_vmNetwork) *string { return v.Macaddr }).(pulumi.StringPtrOutput)
}

func (o Qemu_vmNetworkOutput) Model() pulumi.StringOutput {
	return o.ApplyT(func(v Qemu_vmNetwork) string { return v.Model }).(pulumi.StringOutput)
}

func (o Qemu_vmNetworkOutput) Queues() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Qemu_vmNetwork) *int { return v.Queues }).(pulumi.IntPtrOutput)
}

func (o Qemu_vmNetworkOutput) Rate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Qemu_vmNetwork) *int { return v.Rate }).(pulumi.IntPtrOutput)
}

func (o Qemu_vmNetworkOutput) Tag() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Qemu_vmNetwork) *int { return v.Tag }).(pulumi.IntPtrOutput)
}

type Qemu_vmNetworkArrayOutput struct{ *pulumi.OutputState }

func (Qemu_vmNetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Qemu_vmNetwork)(nil)).Elem()
}

func (o Qemu_vmNetworkArrayOutput) ToQemu_vmNetworkArrayOutput() Qemu_vmNetworkArrayOutput {
	return o
}

func (o Qemu_vmNetworkArrayOutput) ToQemu_vmNetworkArrayOutputWithContext(ctx context.Context) Qemu_vmNetworkArrayOutput {
	return o
}

func (o Qemu_vmNetworkArrayOutput) Index(i pulumi.IntInput) Qemu_vmNetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Qemu_vmNetwork {
		return vs[0].([]Qemu_vmNetwork)[vs[1].(int)]
	}).(Qemu_vmNetworkOutput)
}

type Qemu_vmSerial struct {
	Id   int    `pulumi:"id"`
	Type string `pulumi:"type"`
}

// Qemu_vmSerialInput is an input type that accepts Qemu_vmSerialArgs and Qemu_vmSerialOutput values.
// You can construct a concrete instance of `Qemu_vmSerialInput` via:
//
//          Qemu_vmSerialArgs{...}
type Qemu_vmSerialInput interface {
	pulumi.Input

	ToQemu_vmSerialOutput() Qemu_vmSerialOutput
	ToQemu_vmSerialOutputWithContext(context.Context) Qemu_vmSerialOutput
}

type Qemu_vmSerialArgs struct {
	Id   pulumi.IntInput    `pulumi:"id"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (Qemu_vmSerialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Qemu_vmSerial)(nil)).Elem()
}

func (i Qemu_vmSerialArgs) ToQemu_vmSerialOutput() Qemu_vmSerialOutput {
	return i.ToQemu_vmSerialOutputWithContext(context.Background())
}

func (i Qemu_vmSerialArgs) ToQemu_vmSerialOutputWithContext(ctx context.Context) Qemu_vmSerialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Qemu_vmSerialOutput)
}

// Qemu_vmSerialArrayInput is an input type that accepts Qemu_vmSerialArray and Qemu_vmSerialArrayOutput values.
// You can construct a concrete instance of `Qemu_vmSerialArrayInput` via:
//
//          Qemu_vmSerialArray{ Qemu_vmSerialArgs{...} }
type Qemu_vmSerialArrayInput interface {
	pulumi.Input

	ToQemu_vmSerialArrayOutput() Qemu_vmSerialArrayOutput
	ToQemu_vmSerialArrayOutputWithContext(context.Context) Qemu_vmSerialArrayOutput
}

type Qemu_vmSerialArray []Qemu_vmSerialInput

func (Qemu_vmSerialArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Qemu_vmSerial)(nil)).Elem()
}

func (i Qemu_vmSerialArray) ToQemu_vmSerialArrayOutput() Qemu_vmSerialArrayOutput {
	return i.ToQemu_vmSerialArrayOutputWithContext(context.Background())
}

func (i Qemu_vmSerialArray) ToQemu_vmSerialArrayOutputWithContext(ctx context.Context) Qemu_vmSerialArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Qemu_vmSerialArrayOutput)
}

type Qemu_vmSerialOutput struct{ *pulumi.OutputState }

func (Qemu_vmSerialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Qemu_vmSerial)(nil)).Elem()
}

func (o Qemu_vmSerialOutput) ToQemu_vmSerialOutput() Qemu_vmSerialOutput {
	return o
}

func (o Qemu_vmSerialOutput) ToQemu_vmSerialOutputWithContext(ctx context.Context) Qemu_vmSerialOutput {
	return o
}

func (o Qemu_vmSerialOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v Qemu_vmSerial) int { return v.Id }).(pulumi.IntOutput)
}

func (o Qemu_vmSerialOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v Qemu_vmSerial) string { return v.Type }).(pulumi.StringOutput)
}

type Qemu_vmSerialArrayOutput struct{ *pulumi.OutputState }

func (Qemu_vmSerialArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Qemu_vmSerial)(nil)).Elem()
}

func (o Qemu_vmSerialArrayOutput) ToQemu_vmSerialArrayOutput() Qemu_vmSerialArrayOutput {
	return o
}

func (o Qemu_vmSerialArrayOutput) ToQemu_vmSerialArrayOutputWithContext(ctx context.Context) Qemu_vmSerialArrayOutput {
	return o
}

func (o Qemu_vmSerialArrayOutput) Index(i pulumi.IntInput) Qemu_vmSerialOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Qemu_vmSerial {
		return vs[0].([]Qemu_vmSerial)[vs[1].(int)]
	}).(Qemu_vmSerialOutput)
}

type Qemu_vmVga struct {
	Memory *int    `pulumi:"memory"`
	Type   *string `pulumi:"type"`
}

// Qemu_vmVgaInput is an input type that accepts Qemu_vmVgaArgs and Qemu_vmVgaOutput values.
// You can construct a concrete instance of `Qemu_vmVgaInput` via:
//
//          Qemu_vmVgaArgs{...}
type Qemu_vmVgaInput interface {
	pulumi.Input

	ToQemu_vmVgaOutput() Qemu_vmVgaOutput
	ToQemu_vmVgaOutputWithContext(context.Context) Qemu_vmVgaOutput
}

type Qemu_vmVgaArgs struct {
	Memory pulumi.IntPtrInput    `pulumi:"memory"`
	Type   pulumi.StringPtrInput `pulumi:"type"`
}

func (Qemu_vmVgaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Qemu_vmVga)(nil)).Elem()
}

func (i Qemu_vmVgaArgs) ToQemu_vmVgaOutput() Qemu_vmVgaOutput {
	return i.ToQemu_vmVgaOutputWithContext(context.Background())
}

func (i Qemu_vmVgaArgs) ToQemu_vmVgaOutputWithContext(ctx context.Context) Qemu_vmVgaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Qemu_vmVgaOutput)
}

// Qemu_vmVgaArrayInput is an input type that accepts Qemu_vmVgaArray and Qemu_vmVgaArrayOutput values.
// You can construct a concrete instance of `Qemu_vmVgaArrayInput` via:
//
//          Qemu_vmVgaArray{ Qemu_vmVgaArgs{...} }
type Qemu_vmVgaArrayInput interface {
	pulumi.Input

	ToQemu_vmVgaArrayOutput() Qemu_vmVgaArrayOutput
	ToQemu_vmVgaArrayOutputWithContext(context.Context) Qemu_vmVgaArrayOutput
}

type Qemu_vmVgaArray []Qemu_vmVgaInput

func (Qemu_vmVgaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Qemu_vmVga)(nil)).Elem()
}

func (i Qemu_vmVgaArray) ToQemu_vmVgaArrayOutput() Qemu_vmVgaArrayOutput {
	return i.ToQemu_vmVgaArrayOutputWithContext(context.Background())
}

func (i Qemu_vmVgaArray) ToQemu_vmVgaArrayOutputWithContext(ctx context.Context) Qemu_vmVgaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Qemu_vmVgaArrayOutput)
}

type Qemu_vmVgaOutput struct{ *pulumi.OutputState }

func (Qemu_vmVgaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Qemu_vmVga)(nil)).Elem()
}

func (o Qemu_vmVgaOutput) ToQemu_vmVgaOutput() Qemu_vmVgaOutput {
	return o
}

func (o Qemu_vmVgaOutput) ToQemu_vmVgaOutputWithContext(ctx context.Context) Qemu_vmVgaOutput {
	return o
}

func (o Qemu_vmVgaOutput) Memory() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Qemu_vmVga) *int { return v.Memory }).(pulumi.IntPtrOutput)
}

func (o Qemu_vmVgaOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Qemu_vmVga) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type Qemu_vmVgaArrayOutput struct{ *pulumi.OutputState }

func (Qemu_vmVgaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Qemu_vmVga)(nil)).Elem()
}

func (o Qemu_vmVgaArrayOutput) ToQemu_vmVgaArrayOutput() Qemu_vmVgaArrayOutput {
	return o
}

func (o Qemu_vmVgaArrayOutput) ToQemu_vmVgaArrayOutputWithContext(ctx context.Context) Qemu_vmVgaArrayOutput {
	return o
}

func (o Qemu_vmVgaArrayOutput) Index(i pulumi.IntInput) Qemu_vmVgaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Qemu_vmVga {
		return vs[0].([]Qemu_vmVga)[vs[1].(int)]
	}).(Qemu_vmVgaOutput)
}

func init() {
	pulumi.RegisterOutputType(Lxc_containerFeatureOutput{})
	pulumi.RegisterOutputType(Lxc_containerFeatureArrayOutput{})
	pulumi.RegisterOutputType(Lxc_containerMountpointOutput{})
	pulumi.RegisterOutputType(Lxc_containerMountpointArrayOutput{})
	pulumi.RegisterOutputType(Lxc_containerNetworkOutput{})
	pulumi.RegisterOutputType(Lxc_containerNetworkArrayOutput{})
	pulumi.RegisterOutputType(Qemu_vmDiskOutput{})
	pulumi.RegisterOutputType(Qemu_vmDiskArrayOutput{})
	pulumi.RegisterOutputType(Qemu_vmNetworkOutput{})
	pulumi.RegisterOutputType(Qemu_vmNetworkArrayOutput{})
	pulumi.RegisterOutputType(Qemu_vmSerialOutput{})
	pulumi.RegisterOutputType(Qemu_vmSerialArrayOutput{})
	pulumi.RegisterOutputType(Qemu_vmVgaOutput{})
	pulumi.RegisterOutputType(Qemu_vmVgaArrayOutput{})
}
