// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databricks

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type SecretAcl struct {
	pulumi.CustomResourceState

	Permission pulumi.StringOutput `pulumi:"permission"`
	Principal  pulumi.StringOutput `pulumi:"principal"`
	Scope      pulumi.StringOutput `pulumi:"scope"`
}

// NewSecretAcl registers a new resource with the given unique name, arguments, and options.
func NewSecretAcl(ctx *pulumi.Context,
	name string, args *SecretAclArgs, opts ...pulumi.ResourceOption) (*SecretAcl, error) {
	if args == nil || args.Permission == nil {
		return nil, errors.New("missing required argument 'Permission'")
	}
	if args == nil || args.Principal == nil {
		return nil, errors.New("missing required argument 'Principal'")
	}
	if args == nil || args.Scope == nil {
		return nil, errors.New("missing required argument 'Scope'")
	}
	if args == nil {
		args = &SecretAclArgs{}
	}
	var resource SecretAcl
	err := ctx.RegisterResource("databricks:index/secretAcl:SecretAcl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretAcl gets an existing SecretAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretAclState, opts ...pulumi.ResourceOption) (*SecretAcl, error) {
	var resource SecretAcl
	err := ctx.ReadResource("databricks:index/secretAcl:SecretAcl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretAcl resources.
type secretAclState struct {
	Permission *string `pulumi:"permission"`
	Principal  *string `pulumi:"principal"`
	Scope      *string `pulumi:"scope"`
}

type SecretAclState struct {
	Permission pulumi.StringPtrInput
	Principal  pulumi.StringPtrInput
	Scope      pulumi.StringPtrInput
}

func (SecretAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretAclState)(nil)).Elem()
}

type secretAclArgs struct {
	Permission string `pulumi:"permission"`
	Principal  string `pulumi:"principal"`
	Scope      string `pulumi:"scope"`
}

// The set of arguments for constructing a SecretAcl resource.
type SecretAclArgs struct {
	Permission pulumi.StringInput
	Principal  pulumi.StringInput
	Scope      pulumi.StringInput
}

func (SecretAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretAclArgs)(nil)).Elem()
}