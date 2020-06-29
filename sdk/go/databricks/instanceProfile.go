// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databricks

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type InstanceProfile struct {
	pulumi.CustomResourceState

	InstanceProfileArn pulumi.StringOutput `pulumi:"instanceProfileArn"`
	SkipValidation     pulumi.BoolOutput   `pulumi:"skipValidation"`
}

// NewInstanceProfile registers a new resource with the given unique name, arguments, and options.
func NewInstanceProfile(ctx *pulumi.Context,
	name string, args *InstanceProfileArgs, opts ...pulumi.ResourceOption) (*InstanceProfile, error) {
	if args == nil || args.InstanceProfileArn == nil {
		return nil, errors.New("missing required argument 'InstanceProfileArn'")
	}
	if args == nil || args.SkipValidation == nil {
		return nil, errors.New("missing required argument 'SkipValidation'")
	}
	if args == nil {
		args = &InstanceProfileArgs{}
	}
	var resource InstanceProfile
	err := ctx.RegisterResource("databricks:index/instanceProfile:InstanceProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceProfile gets an existing InstanceProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceProfileState, opts ...pulumi.ResourceOption) (*InstanceProfile, error) {
	var resource InstanceProfile
	err := ctx.ReadResource("databricks:index/instanceProfile:InstanceProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceProfile resources.
type instanceProfileState struct {
	InstanceProfileArn *string `pulumi:"instanceProfileArn"`
	SkipValidation     *bool   `pulumi:"skipValidation"`
}

type InstanceProfileState struct {
	InstanceProfileArn pulumi.StringPtrInput
	SkipValidation     pulumi.BoolPtrInput
}

func (InstanceProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceProfileState)(nil)).Elem()
}

type instanceProfileArgs struct {
	InstanceProfileArn string `pulumi:"instanceProfileArn"`
	SkipValidation     bool   `pulumi:"skipValidation"`
}

// The set of arguments for constructing a InstanceProfile resource.
type InstanceProfileArgs struct {
	InstanceProfileArn pulumi.StringInput
	SkipValidation     pulumi.BoolInput
}

func (InstanceProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceProfileArgs)(nil)).Elem()
}