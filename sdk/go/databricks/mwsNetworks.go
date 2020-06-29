// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databricks

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type MwsNetworks struct {
	pulumi.CustomResourceState

	AccountId        pulumi.StringOutput                `pulumi:"accountId"`
	CreationTime     pulumi.IntOutput                   `pulumi:"creationTime"`
	ErrorMessages    MwsNetworksErrorMessageArrayOutput `pulumi:"errorMessages"`
	NetworkId        pulumi.StringOutput                `pulumi:"networkId"`
	NetworkName      pulumi.StringOutput                `pulumi:"networkName"`
	SecurityGroupIds pulumi.StringArrayOutput           `pulumi:"securityGroupIds"`
	SubnetIds        pulumi.StringArrayOutput           `pulumi:"subnetIds"`
	VpcId            pulumi.StringOutput                `pulumi:"vpcId"`
	VpcStatus        pulumi.StringOutput                `pulumi:"vpcStatus"`
	WorkspaceId      pulumi.IntOutput                   `pulumi:"workspaceId"`
}

// NewMwsNetworks registers a new resource with the given unique name, arguments, and options.
func NewMwsNetworks(ctx *pulumi.Context,
	name string, args *MwsNetworksArgs, opts ...pulumi.ResourceOption) (*MwsNetworks, error) {
	if args == nil || args.AccountId == nil {
		return nil, errors.New("missing required argument 'AccountId'")
	}
	if args == nil || args.NetworkName == nil {
		return nil, errors.New("missing required argument 'NetworkName'")
	}
	if args == nil || args.SecurityGroupIds == nil {
		return nil, errors.New("missing required argument 'SecurityGroupIds'")
	}
	if args == nil || args.SubnetIds == nil {
		return nil, errors.New("missing required argument 'SubnetIds'")
	}
	if args == nil || args.VpcId == nil {
		return nil, errors.New("missing required argument 'VpcId'")
	}
	if args == nil {
		args = &MwsNetworksArgs{}
	}
	var resource MwsNetworks
	err := ctx.RegisterResource("databricks:index/mwsNetworks:MwsNetworks", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMwsNetworks gets an existing MwsNetworks resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMwsNetworks(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MwsNetworksState, opts ...pulumi.ResourceOption) (*MwsNetworks, error) {
	var resource MwsNetworks
	err := ctx.ReadResource("databricks:index/mwsNetworks:MwsNetworks", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MwsNetworks resources.
type mwsNetworksState struct {
	AccountId        *string                   `pulumi:"accountId"`
	CreationTime     *int                      `pulumi:"creationTime"`
	ErrorMessages    []MwsNetworksErrorMessage `pulumi:"errorMessages"`
	NetworkId        *string                   `pulumi:"networkId"`
	NetworkName      *string                   `pulumi:"networkName"`
	SecurityGroupIds []string                  `pulumi:"securityGroupIds"`
	SubnetIds        []string                  `pulumi:"subnetIds"`
	VpcId            *string                   `pulumi:"vpcId"`
	VpcStatus        *string                   `pulumi:"vpcStatus"`
	WorkspaceId      *int                      `pulumi:"workspaceId"`
}

type MwsNetworksState struct {
	AccountId        pulumi.StringPtrInput
	CreationTime     pulumi.IntPtrInput
	ErrorMessages    MwsNetworksErrorMessageArrayInput
	NetworkId        pulumi.StringPtrInput
	NetworkName      pulumi.StringPtrInput
	SecurityGroupIds pulumi.StringArrayInput
	SubnetIds        pulumi.StringArrayInput
	VpcId            pulumi.StringPtrInput
	VpcStatus        pulumi.StringPtrInput
	WorkspaceId      pulumi.IntPtrInput
}

func (MwsNetworksState) ElementType() reflect.Type {
	return reflect.TypeOf((*mwsNetworksState)(nil)).Elem()
}

type mwsNetworksArgs struct {
	AccountId        string                    `pulumi:"accountId"`
	ErrorMessages    []MwsNetworksErrorMessage `pulumi:"errorMessages"`
	NetworkName      string                    `pulumi:"networkName"`
	SecurityGroupIds []string                  `pulumi:"securityGroupIds"`
	SubnetIds        []string                  `pulumi:"subnetIds"`
	VpcId            string                    `pulumi:"vpcId"`
}

// The set of arguments for constructing a MwsNetworks resource.
type MwsNetworksArgs struct {
	AccountId        pulumi.StringInput
	ErrorMessages    MwsNetworksErrorMessageArrayInput
	NetworkName      pulumi.StringInput
	SecurityGroupIds pulumi.StringArrayInput
	SubnetIds        pulumi.StringArrayInput
	VpcId            pulumi.StringInput
}

func (MwsNetworksArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mwsNetworksArgs)(nil)).Elem()
}
