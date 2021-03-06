// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package proxmoxve

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type LXCDisk struct {
	pulumi.CustomResourceState

	Acl          pulumi.BoolPtrOutput         `pulumi:"acl"`
	Backup       pulumi.BoolPtrOutput         `pulumi:"backup"`
	Container    pulumi.StringOutput          `pulumi:"container"`
	Mountoptions LXCDiskMountoptionsPtrOutput `pulumi:"mountoptions"`
	Mp           pulumi.StringOutput          `pulumi:"mp"`
	Quota        pulumi.BoolPtrOutput         `pulumi:"quota"`
	Replicate    pulumi.BoolPtrOutput         `pulumi:"replicate"`
	Shared       pulumi.BoolPtrOutput         `pulumi:"shared"`
	Size         pulumi.StringOutput          `pulumi:"size"`
	Slot         pulumi.IntOutput             `pulumi:"slot"`
	Storage      pulumi.StringOutput          `pulumi:"storage"`
	Volume       pulumi.StringOutput          `pulumi:"volume"`
}

// NewLXCDisk registers a new resource with the given unique name, arguments, and options.
func NewLXCDisk(ctx *pulumi.Context,
	name string, args *LXCDiskArgs, opts ...pulumi.ResourceOption) (*LXCDisk, error) {
	if args == nil || args.Container == nil {
		return nil, errors.New("missing required argument 'Container'")
	}
	if args == nil || args.Mp == nil {
		return nil, errors.New("missing required argument 'Mp'")
	}
	if args == nil || args.Size == nil {
		return nil, errors.New("missing required argument 'Size'")
	}
	if args == nil || args.Slot == nil {
		return nil, errors.New("missing required argument 'Slot'")
	}
	if args == nil || args.Storage == nil {
		return nil, errors.New("missing required argument 'Storage'")
	}
	if args == nil {
		args = &LXCDiskArgs{}
	}
	var resource LXCDisk
	err := ctx.RegisterResource("proxmoxve:index/lXCDisk:LXCDisk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLXCDisk gets an existing LXCDisk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLXCDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LXCDiskState, opts ...pulumi.ResourceOption) (*LXCDisk, error) {
	var resource LXCDisk
	err := ctx.ReadResource("proxmoxve:index/lXCDisk:LXCDisk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LXCDisk resources.
type lxcdiskState struct {
	Acl          *bool                `pulumi:"acl"`
	Backup       *bool                `pulumi:"backup"`
	Container    *string              `pulumi:"container"`
	Mountoptions *LXCDiskMountoptions `pulumi:"mountoptions"`
	Mp           *string              `pulumi:"mp"`
	Quota        *bool                `pulumi:"quota"`
	Replicate    *bool                `pulumi:"replicate"`
	Shared       *bool                `pulumi:"shared"`
	Size         *string              `pulumi:"size"`
	Slot         *int                 `pulumi:"slot"`
	Storage      *string              `pulumi:"storage"`
	Volume       *string              `pulumi:"volume"`
}

type LXCDiskState struct {
	Acl          pulumi.BoolPtrInput
	Backup       pulumi.BoolPtrInput
	Container    pulumi.StringPtrInput
	Mountoptions LXCDiskMountoptionsPtrInput
	Mp           pulumi.StringPtrInput
	Quota        pulumi.BoolPtrInput
	Replicate    pulumi.BoolPtrInput
	Shared       pulumi.BoolPtrInput
	Size         pulumi.StringPtrInput
	Slot         pulumi.IntPtrInput
	Storage      pulumi.StringPtrInput
	Volume       pulumi.StringPtrInput
}

func (LXCDiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*lxcdiskState)(nil)).Elem()
}

type lxcdiskArgs struct {
	Acl          *bool                `pulumi:"acl"`
	Backup       *bool                `pulumi:"backup"`
	Container    string               `pulumi:"container"`
	Mountoptions *LXCDiskMountoptions `pulumi:"mountoptions"`
	Mp           string               `pulumi:"mp"`
	Quota        *bool                `pulumi:"quota"`
	Replicate    *bool                `pulumi:"replicate"`
	Shared       *bool                `pulumi:"shared"`
	Size         string               `pulumi:"size"`
	Slot         int                  `pulumi:"slot"`
	Storage      string               `pulumi:"storage"`
	Volume       *string              `pulumi:"volume"`
}

// The set of arguments for constructing a LXCDisk resource.
type LXCDiskArgs struct {
	Acl          pulumi.BoolPtrInput
	Backup       pulumi.BoolPtrInput
	Container    pulumi.StringInput
	Mountoptions LXCDiskMountoptionsPtrInput
	Mp           pulumi.StringInput
	Quota        pulumi.BoolPtrInput
	Replicate    pulumi.BoolPtrInput
	Shared       pulumi.BoolPtrInput
	Size         pulumi.StringInput
	Slot         pulumi.IntInput
	Storage      pulumi.StringInput
	Volume       pulumi.StringPtrInput
}

func (LXCDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lxcdiskArgs)(nil)).Elem()
}
