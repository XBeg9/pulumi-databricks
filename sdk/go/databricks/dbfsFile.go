// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databricks

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type DbfsFile struct {
	pulumi.CustomResourceState

	Content            pulumi.StringOutput  `pulumi:"content"`
	FileSize           pulumi.IntOutput     `pulumi:"fileSize"`
	Mkdirs             pulumi.BoolPtrOutput `pulumi:"mkdirs"`
	Overwrite          pulumi.BoolPtrOutput `pulumi:"overwrite"`
	Path               pulumi.StringOutput  `pulumi:"path"`
	ValidateRemoteFile pulumi.BoolPtrOutput `pulumi:"validateRemoteFile"`
}

// NewDbfsFile registers a new resource with the given unique name, arguments, and options.
func NewDbfsFile(ctx *pulumi.Context,
	name string, args *DbfsFileArgs, opts ...pulumi.ResourceOption) (*DbfsFile, error) {
	if args == nil || args.Content == nil {
		return nil, errors.New("missing required argument 'Content'")
	}
	if args == nil || args.Path == nil {
		return nil, errors.New("missing required argument 'Path'")
	}
	if args == nil {
		args = &DbfsFileArgs{}
	}
	var resource DbfsFile
	err := ctx.RegisterResource("databricks:index/dbfsFile:DbfsFile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDbfsFile gets an existing DbfsFile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDbfsFile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DbfsFileState, opts ...pulumi.ResourceOption) (*DbfsFile, error) {
	var resource DbfsFile
	err := ctx.ReadResource("databricks:index/dbfsFile:DbfsFile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DbfsFile resources.
type dbfsFileState struct {
	Content            *string `pulumi:"content"`
	FileSize           *int    `pulumi:"fileSize"`
	Mkdirs             *bool   `pulumi:"mkdirs"`
	Overwrite          *bool   `pulumi:"overwrite"`
	Path               *string `pulumi:"path"`
	ValidateRemoteFile *bool   `pulumi:"validateRemoteFile"`
}

type DbfsFileState struct {
	Content            pulumi.StringPtrInput
	FileSize           pulumi.IntPtrInput
	Mkdirs             pulumi.BoolPtrInput
	Overwrite          pulumi.BoolPtrInput
	Path               pulumi.StringPtrInput
	ValidateRemoteFile pulumi.BoolPtrInput
}

func (DbfsFileState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbfsFileState)(nil)).Elem()
}

type dbfsFileArgs struct {
	Content            string `pulumi:"content"`
	Mkdirs             *bool  `pulumi:"mkdirs"`
	Overwrite          *bool  `pulumi:"overwrite"`
	Path               string `pulumi:"path"`
	ValidateRemoteFile *bool  `pulumi:"validateRemoteFile"`
}

// The set of arguments for constructing a DbfsFile resource.
type DbfsFileArgs struct {
	Content            pulumi.StringInput
	Mkdirs             pulumi.BoolPtrInput
	Overwrite          pulumi.BoolPtrInput
	Path               pulumi.StringInput
	ValidateRemoteFile pulumi.BoolPtrInput
}

func (DbfsFileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbfsFileArgs)(nil)).Elem()
}
