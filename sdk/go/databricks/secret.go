// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databricks

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Secret struct {
	pulumi.CustomResourceState

	Key                  pulumi.StringOutput `pulumi:"key"`
	LastUpdatedTimestamp pulumi.IntOutput    `pulumi:"lastUpdatedTimestamp"`
	Scope                pulumi.StringOutput `pulumi:"scope"`
	StringValue          pulumi.StringOutput `pulumi:"stringValue"`
}

// NewSecret registers a new resource with the given unique name, arguments, and options.
func NewSecret(ctx *pulumi.Context,
	name string, args *SecretArgs, opts ...pulumi.ResourceOption) (*Secret, error) {
	if args == nil || args.Key == nil {
		return nil, errors.New("missing required argument 'Key'")
	}
	if args == nil || args.Scope == nil {
		return nil, errors.New("missing required argument 'Scope'")
	}
	if args == nil || args.StringValue == nil {
		return nil, errors.New("missing required argument 'StringValue'")
	}
	if args == nil {
		args = &SecretArgs{}
	}
	var resource Secret
	err := ctx.RegisterResource("databricks:index/secret:Secret", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecret gets an existing Secret resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecret(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretState, opts ...pulumi.ResourceOption) (*Secret, error) {
	var resource Secret
	err := ctx.ReadResource("databricks:index/secret:Secret", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Secret resources.
type secretState struct {
	Key                  *string `pulumi:"key"`
	LastUpdatedTimestamp *int    `pulumi:"lastUpdatedTimestamp"`
	Scope                *string `pulumi:"scope"`
	StringValue          *string `pulumi:"stringValue"`
}

type SecretState struct {
	Key                  pulumi.StringPtrInput
	LastUpdatedTimestamp pulumi.IntPtrInput
	Scope                pulumi.StringPtrInput
	StringValue          pulumi.StringPtrInput
}

func (SecretState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretState)(nil)).Elem()
}

type secretArgs struct {
	Key         string `pulumi:"key"`
	Scope       string `pulumi:"scope"`
	StringValue string `pulumi:"stringValue"`
}

// The set of arguments for constructing a Secret resource.
type SecretArgs struct {
	Key         pulumi.StringInput
	Scope       pulumi.StringInput
	StringValue pulumi.StringInput
}

func (SecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretArgs)(nil)).Elem()
}