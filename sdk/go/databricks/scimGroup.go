// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databricks

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ScimGroup struct {
	pulumi.CustomResourceState

	DisplayName    pulumi.StringOutput      `pulumi:"displayName"`
	Entitlements   pulumi.StringArrayOutput `pulumi:"entitlements"`
	InheritedRoles pulumi.StringArrayOutput `pulumi:"inheritedRoles"`
	Members        pulumi.StringArrayOutput `pulumi:"members"`
	Roles          pulumi.StringArrayOutput `pulumi:"roles"`
}

// NewScimGroup registers a new resource with the given unique name, arguments, and options.
func NewScimGroup(ctx *pulumi.Context,
	name string, args *ScimGroupArgs, opts ...pulumi.ResourceOption) (*ScimGroup, error) {
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil {
		args = &ScimGroupArgs{}
	}
	var resource ScimGroup
	err := ctx.RegisterResource("databricks:index/scimGroup:ScimGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScimGroup gets an existing ScimGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScimGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScimGroupState, opts ...pulumi.ResourceOption) (*ScimGroup, error) {
	var resource ScimGroup
	err := ctx.ReadResource("databricks:index/scimGroup:ScimGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScimGroup resources.
type scimGroupState struct {
	DisplayName    *string  `pulumi:"displayName"`
	Entitlements   []string `pulumi:"entitlements"`
	InheritedRoles []string `pulumi:"inheritedRoles"`
	Members        []string `pulumi:"members"`
	Roles          []string `pulumi:"roles"`
}

type ScimGroupState struct {
	DisplayName    pulumi.StringPtrInput
	Entitlements   pulumi.StringArrayInput
	InheritedRoles pulumi.StringArrayInput
	Members        pulumi.StringArrayInput
	Roles          pulumi.StringArrayInput
}

func (ScimGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*scimGroupState)(nil)).Elem()
}

type scimGroupArgs struct {
	DisplayName  string   `pulumi:"displayName"`
	Entitlements []string `pulumi:"entitlements"`
	Members      []string `pulumi:"members"`
	Roles        []string `pulumi:"roles"`
}

// The set of arguments for constructing a ScimGroup resource.
type ScimGroupArgs struct {
	DisplayName  pulumi.StringInput
	Entitlements pulumi.StringArrayInput
	Members      pulumi.StringArrayInput
	Roles        pulumi.StringArrayInput
}

func (ScimGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scimGroupArgs)(nil)).Elem()
}
