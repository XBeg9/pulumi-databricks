// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databricks

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type AzureAdlsGen2Mount struct {
	pulumi.CustomResourceState

	ClientId             pulumi.StringOutput `pulumi:"clientId"`
	ClientSecretKey      pulumi.StringOutput `pulumi:"clientSecretKey"`
	ClientSecretScope    pulumi.StringOutput `pulumi:"clientSecretScope"`
	ClusterId            pulumi.StringOutput `pulumi:"clusterId"`
	ContainerName        pulumi.StringOutput `pulumi:"containerName"`
	Directory            pulumi.StringOutput `pulumi:"directory"`
	InitializeFileSystem pulumi.BoolOutput   `pulumi:"initializeFileSystem"`
	MountName            pulumi.StringOutput `pulumi:"mountName"`
	StorageAccountName   pulumi.StringOutput `pulumi:"storageAccountName"`
	TenantId             pulumi.StringOutput `pulumi:"tenantId"`
}

// NewAzureAdlsGen2Mount registers a new resource with the given unique name, arguments, and options.
func NewAzureAdlsGen2Mount(ctx *pulumi.Context,
	name string, args *AzureAdlsGen2MountArgs, opts ...pulumi.ResourceOption) (*AzureAdlsGen2Mount, error) {
	if args == nil || args.ClientId == nil {
		return nil, errors.New("missing required argument 'ClientId'")
	}
	if args == nil || args.ClientSecretKey == nil {
		return nil, errors.New("missing required argument 'ClientSecretKey'")
	}
	if args == nil || args.ClientSecretScope == nil {
		return nil, errors.New("missing required argument 'ClientSecretScope'")
	}
	if args == nil || args.ClusterId == nil {
		return nil, errors.New("missing required argument 'ClusterId'")
	}
	if args == nil || args.ContainerName == nil {
		return nil, errors.New("missing required argument 'ContainerName'")
	}
	if args == nil || args.InitializeFileSystem == nil {
		return nil, errors.New("missing required argument 'InitializeFileSystem'")
	}
	if args == nil || args.MountName == nil {
		return nil, errors.New("missing required argument 'MountName'")
	}
	if args == nil || args.StorageAccountName == nil {
		return nil, errors.New("missing required argument 'StorageAccountName'")
	}
	if args == nil || args.TenantId == nil {
		return nil, errors.New("missing required argument 'TenantId'")
	}
	if args == nil {
		args = &AzureAdlsGen2MountArgs{}
	}
	var resource AzureAdlsGen2Mount
	err := ctx.RegisterResource("databricks:index/azureAdlsGen2Mount:AzureAdlsGen2Mount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAzureAdlsGen2Mount gets an existing AzureAdlsGen2Mount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAzureAdlsGen2Mount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureAdlsGen2MountState, opts ...pulumi.ResourceOption) (*AzureAdlsGen2Mount, error) {
	var resource AzureAdlsGen2Mount
	err := ctx.ReadResource("databricks:index/azureAdlsGen2Mount:AzureAdlsGen2Mount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AzureAdlsGen2Mount resources.
type azureAdlsGen2MountState struct {
	ClientId             *string `pulumi:"clientId"`
	ClientSecretKey      *string `pulumi:"clientSecretKey"`
	ClientSecretScope    *string `pulumi:"clientSecretScope"`
	ClusterId            *string `pulumi:"clusterId"`
	ContainerName        *string `pulumi:"containerName"`
	Directory            *string `pulumi:"directory"`
	InitializeFileSystem *bool   `pulumi:"initializeFileSystem"`
	MountName            *string `pulumi:"mountName"`
	StorageAccountName   *string `pulumi:"storageAccountName"`
	TenantId             *string `pulumi:"tenantId"`
}

type AzureAdlsGen2MountState struct {
	ClientId             pulumi.StringPtrInput
	ClientSecretKey      pulumi.StringPtrInput
	ClientSecretScope    pulumi.StringPtrInput
	ClusterId            pulumi.StringPtrInput
	ContainerName        pulumi.StringPtrInput
	Directory            pulumi.StringPtrInput
	InitializeFileSystem pulumi.BoolPtrInput
	MountName            pulumi.StringPtrInput
	StorageAccountName   pulumi.StringPtrInput
	TenantId             pulumi.StringPtrInput
}

func (AzureAdlsGen2MountState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureAdlsGen2MountState)(nil)).Elem()
}

type azureAdlsGen2MountArgs struct {
	ClientId             string  `pulumi:"clientId"`
	ClientSecretKey      string  `pulumi:"clientSecretKey"`
	ClientSecretScope    string  `pulumi:"clientSecretScope"`
	ClusterId            string  `pulumi:"clusterId"`
	ContainerName        string  `pulumi:"containerName"`
	Directory            *string `pulumi:"directory"`
	InitializeFileSystem bool    `pulumi:"initializeFileSystem"`
	MountName            string  `pulumi:"mountName"`
	StorageAccountName   string  `pulumi:"storageAccountName"`
	TenantId             string  `pulumi:"tenantId"`
}

// The set of arguments for constructing a AzureAdlsGen2Mount resource.
type AzureAdlsGen2MountArgs struct {
	ClientId             pulumi.StringInput
	ClientSecretKey      pulumi.StringInput
	ClientSecretScope    pulumi.StringInput
	ClusterId            pulumi.StringInput
	ContainerName        pulumi.StringInput
	Directory            pulumi.StringPtrInput
	InitializeFileSystem pulumi.BoolInput
	MountName            pulumi.StringInput
	StorageAccountName   pulumi.StringInput
	TenantId             pulumi.StringInput
}

func (AzureAdlsGen2MountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureAdlsGen2MountArgs)(nil)).Elem()
}
