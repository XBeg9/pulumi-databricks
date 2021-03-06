// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databricks

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Job struct {
	pulumi.CustomResourceState

	CreatedTime            pulumi.IntOutput               `pulumi:"createdTime"`
	CreatorUserName        pulumi.StringOutput            `pulumi:"creatorUserName"`
	EmailNotifications     JobEmailNotificationsPtrOutput `pulumi:"emailNotifications"`
	ExistingClusterId      pulumi.StringPtrOutput         `pulumi:"existingClusterId"`
	JarMainClassName       pulumi.StringPtrOutput         `pulumi:"jarMainClassName"`
	JarParameters          pulumi.StringArrayOutput       `pulumi:"jarParameters"`
	JarUri                 pulumi.StringPtrOutput         `pulumi:"jarUri"`
	JobId                  pulumi.IntOutput               `pulumi:"jobId"`
	LibraryCrans           JobLibraryCranArrayOutput      `pulumi:"libraryCrans"`
	LibraryEggs            pulumi.StringArrayOutput       `pulumi:"libraryEggs"`
	LibraryJars            pulumi.StringArrayOutput       `pulumi:"libraryJars"`
	LibraryMavens          JobLibraryMavenArrayOutput     `pulumi:"libraryMavens"`
	LibraryPypis           JobLibraryPypiArrayOutput      `pulumi:"libraryPypis"`
	LibraryWhls            pulumi.StringArrayOutput       `pulumi:"libraryWhls"`
	MaxConcurrentRuns      pulumi.IntPtrOutput            `pulumi:"maxConcurrentRuns"`
	MaxRetries             pulumi.IntPtrOutput            `pulumi:"maxRetries"`
	MinRetryIntervalMillis pulumi.IntPtrOutput            `pulumi:"minRetryIntervalMillis"`
	Name                   pulumi.StringOutput            `pulumi:"name"`
	NewCluster             JobNewClusterPtrOutput         `pulumi:"newCluster"`
	NotebookBaseParameters pulumi.StringMapOutput         `pulumi:"notebookBaseParameters"`
	NotebookPath           pulumi.StringPtrOutput         `pulumi:"notebookPath"`
	PythonFile             pulumi.StringPtrOutput         `pulumi:"pythonFile"`
	PythonParameters       pulumi.StringArrayOutput       `pulumi:"pythonParameters"`
	RetryOnTimeout         pulumi.BoolPtrOutput           `pulumi:"retryOnTimeout"`
	Schedule               JobSchedulePtrOutput           `pulumi:"schedule"`
	SparkSubmitParameters  pulumi.StringArrayOutput       `pulumi:"sparkSubmitParameters"`
	TimeoutSeconds         pulumi.IntPtrOutput            `pulumi:"timeoutSeconds"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		args = &JobArgs{}
	}
	var resource Job
	err := ctx.RegisterResource("databricks:index/job:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("databricks:index/job:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
	CreatedTime            *int                   `pulumi:"createdTime"`
	CreatorUserName        *string                `pulumi:"creatorUserName"`
	EmailNotifications     *JobEmailNotifications `pulumi:"emailNotifications"`
	ExistingClusterId      *string                `pulumi:"existingClusterId"`
	JarMainClassName       *string                `pulumi:"jarMainClassName"`
	JarParameters          []string               `pulumi:"jarParameters"`
	JarUri                 *string                `pulumi:"jarUri"`
	JobId                  *int                   `pulumi:"jobId"`
	LibraryCrans           []JobLibraryCran       `pulumi:"libraryCrans"`
	LibraryEggs            []string               `pulumi:"libraryEggs"`
	LibraryJars            []string               `pulumi:"libraryJars"`
	LibraryMavens          []JobLibraryMaven      `pulumi:"libraryMavens"`
	LibraryPypis           []JobLibraryPypi       `pulumi:"libraryPypis"`
	LibraryWhls            []string               `pulumi:"libraryWhls"`
	MaxConcurrentRuns      *int                   `pulumi:"maxConcurrentRuns"`
	MaxRetries             *int                   `pulumi:"maxRetries"`
	MinRetryIntervalMillis *int                   `pulumi:"minRetryIntervalMillis"`
	Name                   *string                `pulumi:"name"`
	NewCluster             *JobNewCluster         `pulumi:"newCluster"`
	NotebookBaseParameters map[string]string      `pulumi:"notebookBaseParameters"`
	NotebookPath           *string                `pulumi:"notebookPath"`
	PythonFile             *string                `pulumi:"pythonFile"`
	PythonParameters       []string               `pulumi:"pythonParameters"`
	RetryOnTimeout         *bool                  `pulumi:"retryOnTimeout"`
	Schedule               *JobSchedule           `pulumi:"schedule"`
	SparkSubmitParameters  []string               `pulumi:"sparkSubmitParameters"`
	TimeoutSeconds         *int                   `pulumi:"timeoutSeconds"`
}

type JobState struct {
	CreatedTime            pulumi.IntPtrInput
	CreatorUserName        pulumi.StringPtrInput
	EmailNotifications     JobEmailNotificationsPtrInput
	ExistingClusterId      pulumi.StringPtrInput
	JarMainClassName       pulumi.StringPtrInput
	JarParameters          pulumi.StringArrayInput
	JarUri                 pulumi.StringPtrInput
	JobId                  pulumi.IntPtrInput
	LibraryCrans           JobLibraryCranArrayInput
	LibraryEggs            pulumi.StringArrayInput
	LibraryJars            pulumi.StringArrayInput
	LibraryMavens          JobLibraryMavenArrayInput
	LibraryPypis           JobLibraryPypiArrayInput
	LibraryWhls            pulumi.StringArrayInput
	MaxConcurrentRuns      pulumi.IntPtrInput
	MaxRetries             pulumi.IntPtrInput
	MinRetryIntervalMillis pulumi.IntPtrInput
	Name                   pulumi.StringPtrInput
	NewCluster             JobNewClusterPtrInput
	NotebookBaseParameters pulumi.StringMapInput
	NotebookPath           pulumi.StringPtrInput
	PythonFile             pulumi.StringPtrInput
	PythonParameters       pulumi.StringArrayInput
	RetryOnTimeout         pulumi.BoolPtrInput
	Schedule               JobSchedulePtrInput
	SparkSubmitParameters  pulumi.StringArrayInput
	TimeoutSeconds         pulumi.IntPtrInput
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	EmailNotifications     *JobEmailNotifications `pulumi:"emailNotifications"`
	ExistingClusterId      *string                `pulumi:"existingClusterId"`
	JarMainClassName       *string                `pulumi:"jarMainClassName"`
	JarParameters          []string               `pulumi:"jarParameters"`
	JarUri                 *string                `pulumi:"jarUri"`
	LibraryCrans           []JobLibraryCran       `pulumi:"libraryCrans"`
	LibraryEggs            []string               `pulumi:"libraryEggs"`
	LibraryJars            []string               `pulumi:"libraryJars"`
	LibraryMavens          []JobLibraryMaven      `pulumi:"libraryMavens"`
	LibraryPypis           []JobLibraryPypi       `pulumi:"libraryPypis"`
	LibraryWhls            []string               `pulumi:"libraryWhls"`
	MaxConcurrentRuns      *int                   `pulumi:"maxConcurrentRuns"`
	MaxRetries             *int                   `pulumi:"maxRetries"`
	MinRetryIntervalMillis *int                   `pulumi:"minRetryIntervalMillis"`
	Name                   *string                `pulumi:"name"`
	NewCluster             *JobNewCluster         `pulumi:"newCluster"`
	NotebookBaseParameters map[string]string      `pulumi:"notebookBaseParameters"`
	NotebookPath           *string                `pulumi:"notebookPath"`
	PythonFile             *string                `pulumi:"pythonFile"`
	PythonParameters       []string               `pulumi:"pythonParameters"`
	RetryOnTimeout         *bool                  `pulumi:"retryOnTimeout"`
	Schedule               *JobSchedule           `pulumi:"schedule"`
	SparkSubmitParameters  []string               `pulumi:"sparkSubmitParameters"`
	TimeoutSeconds         *int                   `pulumi:"timeoutSeconds"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	EmailNotifications     JobEmailNotificationsPtrInput
	ExistingClusterId      pulumi.StringPtrInput
	JarMainClassName       pulumi.StringPtrInput
	JarParameters          pulumi.StringArrayInput
	JarUri                 pulumi.StringPtrInput
	LibraryCrans           JobLibraryCranArrayInput
	LibraryEggs            pulumi.StringArrayInput
	LibraryJars            pulumi.StringArrayInput
	LibraryMavens          JobLibraryMavenArrayInput
	LibraryPypis           JobLibraryPypiArrayInput
	LibraryWhls            pulumi.StringArrayInput
	MaxConcurrentRuns      pulumi.IntPtrInput
	MaxRetries             pulumi.IntPtrInput
	MinRetryIntervalMillis pulumi.IntPtrInput
	Name                   pulumi.StringPtrInput
	NewCluster             JobNewClusterPtrInput
	NotebookBaseParameters pulumi.StringMapInput
	NotebookPath           pulumi.StringPtrInput
	PythonFile             pulumi.StringPtrInput
	PythonParameters       pulumi.StringArrayInput
	RetryOnTimeout         pulumi.BoolPtrInput
	Schedule               JobSchedulePtrInput
	SparkSubmitParameters  pulumi.StringArrayInput
	TimeoutSeconds         pulumi.IntPtrInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}
