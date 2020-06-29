// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databricks

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Notebook struct {
	pulumi.CustomResourceState

	Content    pulumi.StringOutput    `pulumi:"content"`
	Format     pulumi.StringPtrOutput `pulumi:"format"`
	Language   pulumi.StringPtrOutput `pulumi:"language"`
	Mkdirs     pulumi.BoolPtrOutput   `pulumi:"mkdirs"`
	ObjectId   pulumi.IntOutput       `pulumi:"objectId"`
	ObjectType pulumi.StringOutput    `pulumi:"objectType"`
	Overwrite  pulumi.BoolPtrOutput   `pulumi:"overwrite"`
	Path       pulumi.StringOutput    `pulumi:"path"`
}

// NewNotebook registers a new resource with the given unique name, arguments, and options.
func NewNotebook(ctx *pulumi.Context,
	name string, args *NotebookArgs, opts ...pulumi.ResourceOption) (*Notebook, error) {
	if args == nil || args.Content == nil {
		return nil, errors.New("missing required argument 'Content'")
	}
	if args == nil || args.Path == nil {
		return nil, errors.New("missing required argument 'Path'")
	}
	if args == nil {
		args = &NotebookArgs{}
	}
	var resource Notebook
	err := ctx.RegisterResource("databricks:index/notebook:Notebook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotebook gets an existing Notebook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotebook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotebookState, opts ...pulumi.ResourceOption) (*Notebook, error) {
	var resource Notebook
	err := ctx.ReadResource("databricks:index/notebook:Notebook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Notebook resources.
type notebookState struct {
	Content    *string `pulumi:"content"`
	Format     *string `pulumi:"format"`
	Language   *string `pulumi:"language"`
	Mkdirs     *bool   `pulumi:"mkdirs"`
	ObjectId   *int    `pulumi:"objectId"`
	ObjectType *string `pulumi:"objectType"`
	Overwrite  *bool   `pulumi:"overwrite"`
	Path       *string `pulumi:"path"`
}

type NotebookState struct {
	Content    pulumi.StringPtrInput
	Format     pulumi.StringPtrInput
	Language   pulumi.StringPtrInput
	Mkdirs     pulumi.BoolPtrInput
	ObjectId   pulumi.IntPtrInput
	ObjectType pulumi.StringPtrInput
	Overwrite  pulumi.BoolPtrInput
	Path       pulumi.StringPtrInput
}

func (NotebookState) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookState)(nil)).Elem()
}

type notebookArgs struct {
	Content   string  `pulumi:"content"`
	Format    *string `pulumi:"format"`
	Language  *string `pulumi:"language"`
	Mkdirs    *bool   `pulumi:"mkdirs"`
	Overwrite *bool   `pulumi:"overwrite"`
	Path      string  `pulumi:"path"`
}

// The set of arguments for constructing a Notebook resource.
type NotebookArgs struct {
	Content   pulumi.StringInput
	Format    pulumi.StringPtrInput
	Language  pulumi.StringPtrInput
	Mkdirs    pulumi.BoolPtrInput
	Overwrite pulumi.BoolPtrInput
	Path      pulumi.StringInput
}

func (NotebookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookArgs)(nil)).Elem()
}