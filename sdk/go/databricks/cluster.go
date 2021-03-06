// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databricks

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Cluster struct {
	pulumi.CustomResourceState

	Autoscale              ClusterAutoscalePtrOutput      `pulumi:"autoscale"`
	AutoterminationMinutes pulumi.IntPtrOutput            `pulumi:"autoterminationMinutes"`
	AwsAttributes          ClusterAwsAttributesPtrOutput  `pulumi:"awsAttributes"`
	ClusterId              pulumi.StringOutput            `pulumi:"clusterId"`
	ClusterLogConf         ClusterClusterLogConfPtrOutput `pulumi:"clusterLogConf"`
	ClusterName            pulumi.StringOutput            `pulumi:"clusterName"`
	CustomTags             pulumi.MapOutput               `pulumi:"customTags"`
	DefaultTags            pulumi.MapOutput               `pulumi:"defaultTags"`
	DockerImage            ClusterDockerImagePtrOutput    `pulumi:"dockerImage"`
	DriverNodeTypeId       pulumi.StringOutput            `pulumi:"driverNodeTypeId"`
	EnableElasticDisk      pulumi.BoolOutput              `pulumi:"enableElasticDisk"`
	IdempotencyToken       pulumi.IntPtrOutput            `pulumi:"idempotencyToken"`
	InitScripts            ClusterInitScriptArrayOutput   `pulumi:"initScripts"`
	InstancePoolId         pulumi.StringPtrOutput         `pulumi:"instancePoolId"`
	LibraryCrans           ClusterLibraryCranArrayOutput  `pulumi:"libraryCrans"`
	LibraryEggs            ClusterLibraryEggArrayOutput   `pulumi:"libraryEggs"`
	LibraryJars            ClusterLibraryJarArrayOutput   `pulumi:"libraryJars"`
	LibraryMavens          ClusterLibraryMavenArrayOutput `pulumi:"libraryMavens"`
	LibraryPypis           ClusterLibraryPypiArrayOutput  `pulumi:"libraryPypis"`
	LibraryWhls            ClusterLibraryWhlArrayOutput   `pulumi:"libraryWhls"`
	NodeTypeId             pulumi.StringPtrOutput         `pulumi:"nodeTypeId"`
	NumWorkers             pulumi.IntPtrOutput            `pulumi:"numWorkers"`
	SingleUserName         pulumi.StringPtrOutput         `pulumi:"singleUserName"`
	SparkConf              pulumi.MapOutput               `pulumi:"sparkConf"`
	SparkEnvVars           pulumi.MapOutput               `pulumi:"sparkEnvVars"`
	SparkVersion           pulumi.StringOutput            `pulumi:"sparkVersion"`
	SshPublicKeys          pulumi.StringArrayOutput       `pulumi:"sshPublicKeys"`
	State                  pulumi.StringOutput            `pulumi:"state"`
	StateMessage           pulumi.StringOutput            `pulumi:"stateMessage"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil || args.SparkVersion == nil {
		return nil, errors.New("missing required argument 'SparkVersion'")
	}
	if args == nil {
		args = &ClusterArgs{}
	}
	var resource Cluster
	err := ctx.RegisterResource("databricks:index/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("databricks:index/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	Autoscale              *ClusterAutoscale      `pulumi:"autoscale"`
	AutoterminationMinutes *int                   `pulumi:"autoterminationMinutes"`
	AwsAttributes          *ClusterAwsAttributes  `pulumi:"awsAttributes"`
	ClusterId              *string                `pulumi:"clusterId"`
	ClusterLogConf         *ClusterClusterLogConf `pulumi:"clusterLogConf"`
	ClusterName            *string                `pulumi:"clusterName"`
	CustomTags             map[string]interface{} `pulumi:"customTags"`
	DefaultTags            map[string]interface{} `pulumi:"defaultTags"`
	DockerImage            *ClusterDockerImage    `pulumi:"dockerImage"`
	DriverNodeTypeId       *string                `pulumi:"driverNodeTypeId"`
	EnableElasticDisk      *bool                  `pulumi:"enableElasticDisk"`
	IdempotencyToken       *int                   `pulumi:"idempotencyToken"`
	InitScripts            []ClusterInitScript    `pulumi:"initScripts"`
	InstancePoolId         *string                `pulumi:"instancePoolId"`
	LibraryCrans           []ClusterLibraryCran   `pulumi:"libraryCrans"`
	LibraryEggs            []ClusterLibraryEgg    `pulumi:"libraryEggs"`
	LibraryJars            []ClusterLibraryJar    `pulumi:"libraryJars"`
	LibraryMavens          []ClusterLibraryMaven  `pulumi:"libraryMavens"`
	LibraryPypis           []ClusterLibraryPypi   `pulumi:"libraryPypis"`
	LibraryWhls            []ClusterLibraryWhl    `pulumi:"libraryWhls"`
	NodeTypeId             *string                `pulumi:"nodeTypeId"`
	NumWorkers             *int                   `pulumi:"numWorkers"`
	SingleUserName         *string                `pulumi:"singleUserName"`
	SparkConf              map[string]interface{} `pulumi:"sparkConf"`
	SparkEnvVars           map[string]interface{} `pulumi:"sparkEnvVars"`
	SparkVersion           *string                `pulumi:"sparkVersion"`
	SshPublicKeys          []string               `pulumi:"sshPublicKeys"`
	State                  *string                `pulumi:"state"`
	StateMessage           *string                `pulumi:"stateMessage"`
}

type ClusterState struct {
	Autoscale              ClusterAutoscalePtrInput
	AutoterminationMinutes pulumi.IntPtrInput
	AwsAttributes          ClusterAwsAttributesPtrInput
	ClusterId              pulumi.StringPtrInput
	ClusterLogConf         ClusterClusterLogConfPtrInput
	ClusterName            pulumi.StringPtrInput
	CustomTags             pulumi.MapInput
	DefaultTags            pulumi.MapInput
	DockerImage            ClusterDockerImagePtrInput
	DriverNodeTypeId       pulumi.StringPtrInput
	EnableElasticDisk      pulumi.BoolPtrInput
	IdempotencyToken       pulumi.IntPtrInput
	InitScripts            ClusterInitScriptArrayInput
	InstancePoolId         pulumi.StringPtrInput
	LibraryCrans           ClusterLibraryCranArrayInput
	LibraryEggs            ClusterLibraryEggArrayInput
	LibraryJars            ClusterLibraryJarArrayInput
	LibraryMavens          ClusterLibraryMavenArrayInput
	LibraryPypis           ClusterLibraryPypiArrayInput
	LibraryWhls            ClusterLibraryWhlArrayInput
	NodeTypeId             pulumi.StringPtrInput
	NumWorkers             pulumi.IntPtrInput
	SingleUserName         pulumi.StringPtrInput
	SparkConf              pulumi.MapInput
	SparkEnvVars           pulumi.MapInput
	SparkVersion           pulumi.StringPtrInput
	SshPublicKeys          pulumi.StringArrayInput
	State                  pulumi.StringPtrInput
	StateMessage           pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	Autoscale              *ClusterAutoscale      `pulumi:"autoscale"`
	AutoterminationMinutes *int                   `pulumi:"autoterminationMinutes"`
	AwsAttributes          *ClusterAwsAttributes  `pulumi:"awsAttributes"`
	ClusterLogConf         *ClusterClusterLogConf `pulumi:"clusterLogConf"`
	ClusterName            *string                `pulumi:"clusterName"`
	CustomTags             map[string]interface{} `pulumi:"customTags"`
	DockerImage            *ClusterDockerImage    `pulumi:"dockerImage"`
	DriverNodeTypeId       *string                `pulumi:"driverNodeTypeId"`
	EnableElasticDisk      *bool                  `pulumi:"enableElasticDisk"`
	IdempotencyToken       *int                   `pulumi:"idempotencyToken"`
	InitScripts            []ClusterInitScript    `pulumi:"initScripts"`
	InstancePoolId         *string                `pulumi:"instancePoolId"`
	LibraryCrans           []ClusterLibraryCran   `pulumi:"libraryCrans"`
	LibraryEggs            []ClusterLibraryEgg    `pulumi:"libraryEggs"`
	LibraryJars            []ClusterLibraryJar    `pulumi:"libraryJars"`
	LibraryMavens          []ClusterLibraryMaven  `pulumi:"libraryMavens"`
	LibraryPypis           []ClusterLibraryPypi   `pulumi:"libraryPypis"`
	LibraryWhls            []ClusterLibraryWhl    `pulumi:"libraryWhls"`
	NodeTypeId             *string                `pulumi:"nodeTypeId"`
	NumWorkers             *int                   `pulumi:"numWorkers"`
	SingleUserName         *string                `pulumi:"singleUserName"`
	SparkConf              map[string]interface{} `pulumi:"sparkConf"`
	SparkEnvVars           map[string]interface{} `pulumi:"sparkEnvVars"`
	SparkVersion           string                 `pulumi:"sparkVersion"`
	SshPublicKeys          []string               `pulumi:"sshPublicKeys"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	Autoscale              ClusterAutoscalePtrInput
	AutoterminationMinutes pulumi.IntPtrInput
	AwsAttributes          ClusterAwsAttributesPtrInput
	ClusterLogConf         ClusterClusterLogConfPtrInput
	ClusterName            pulumi.StringPtrInput
	CustomTags             pulumi.MapInput
	DockerImage            ClusterDockerImagePtrInput
	DriverNodeTypeId       pulumi.StringPtrInput
	EnableElasticDisk      pulumi.BoolPtrInput
	IdempotencyToken       pulumi.IntPtrInput
	InitScripts            ClusterInitScriptArrayInput
	InstancePoolId         pulumi.StringPtrInput
	LibraryCrans           ClusterLibraryCranArrayInput
	LibraryEggs            ClusterLibraryEggArrayInput
	LibraryJars            ClusterLibraryJarArrayInput
	LibraryMavens          ClusterLibraryMavenArrayInput
	LibraryPypis           ClusterLibraryPypiArrayInput
	LibraryWhls            ClusterLibraryWhlArrayInput
	NodeTypeId             pulumi.StringPtrInput
	NumWorkers             pulumi.IntPtrInput
	SingleUserName         pulumi.StringPtrInput
	SparkConf              pulumi.MapInput
	SparkEnvVars           pulumi.MapInput
	SparkVersion           pulumi.StringInput
	SshPublicKeys          pulumi.StringArrayInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}
