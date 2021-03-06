// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type MwsNetworksErrorMessage struct {
	ErrorMessage *string `pulumi:"errorMessage"`
	ErrorType    *string `pulumi:"errorType"`
}

// MwsNetworksErrorMessageInput is an input type that accepts MwsNetworksErrorMessageArgs and MwsNetworksErrorMessageOutput values.
// You can construct a concrete instance of `MwsNetworksErrorMessageInput` via:
//
// 		 MwsNetworksErrorMessageArgs{...}
//
type MwsNetworksErrorMessageInput interface {
	pulumi.Input

	ToMwsNetworksErrorMessageOutput() MwsNetworksErrorMessageOutput
	ToMwsNetworksErrorMessageOutputWithContext(context.Context) MwsNetworksErrorMessageOutput
}

type MwsNetworksErrorMessageArgs struct {
	ErrorMessage pulumi.StringPtrInput `pulumi:"errorMessage"`
	ErrorType    pulumi.StringPtrInput `pulumi:"errorType"`
}

func (MwsNetworksErrorMessageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MwsNetworksErrorMessage)(nil)).Elem()
}

func (i MwsNetworksErrorMessageArgs) ToMwsNetworksErrorMessageOutput() MwsNetworksErrorMessageOutput {
	return i.ToMwsNetworksErrorMessageOutputWithContext(context.Background())
}

func (i MwsNetworksErrorMessageArgs) ToMwsNetworksErrorMessageOutputWithContext(ctx context.Context) MwsNetworksErrorMessageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MwsNetworksErrorMessageOutput)
}

// MwsNetworksErrorMessageArrayInput is an input type that accepts MwsNetworksErrorMessageArray and MwsNetworksErrorMessageArrayOutput values.
// You can construct a concrete instance of `MwsNetworksErrorMessageArrayInput` via:
//
// 		 MwsNetworksErrorMessageArray{ MwsNetworksErrorMessageArgs{...} }
//
type MwsNetworksErrorMessageArrayInput interface {
	pulumi.Input

	ToMwsNetworksErrorMessageArrayOutput() MwsNetworksErrorMessageArrayOutput
	ToMwsNetworksErrorMessageArrayOutputWithContext(context.Context) MwsNetworksErrorMessageArrayOutput
}

type MwsNetworksErrorMessageArray []MwsNetworksErrorMessageInput

func (MwsNetworksErrorMessageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MwsNetworksErrorMessage)(nil)).Elem()
}

func (i MwsNetworksErrorMessageArray) ToMwsNetworksErrorMessageArrayOutput() MwsNetworksErrorMessageArrayOutput {
	return i.ToMwsNetworksErrorMessageArrayOutputWithContext(context.Background())
}

func (i MwsNetworksErrorMessageArray) ToMwsNetworksErrorMessageArrayOutputWithContext(ctx context.Context) MwsNetworksErrorMessageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MwsNetworksErrorMessageArrayOutput)
}

type MwsNetworksErrorMessageOutput struct{ *pulumi.OutputState }

func (MwsNetworksErrorMessageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MwsNetworksErrorMessage)(nil)).Elem()
}

func (o MwsNetworksErrorMessageOutput) ToMwsNetworksErrorMessageOutput() MwsNetworksErrorMessageOutput {
	return o
}

func (o MwsNetworksErrorMessageOutput) ToMwsNetworksErrorMessageOutputWithContext(ctx context.Context) MwsNetworksErrorMessageOutput {
	return o
}

func (o MwsNetworksErrorMessageOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MwsNetworksErrorMessage) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o MwsNetworksErrorMessageOutput) ErrorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MwsNetworksErrorMessage) *string { return v.ErrorType }).(pulumi.StringPtrOutput)
}

type MwsNetworksErrorMessageArrayOutput struct{ *pulumi.OutputState }

func (MwsNetworksErrorMessageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MwsNetworksErrorMessage)(nil)).Elem()
}

func (o MwsNetworksErrorMessageArrayOutput) ToMwsNetworksErrorMessageArrayOutput() MwsNetworksErrorMessageArrayOutput {
	return o
}

func (o MwsNetworksErrorMessageArrayOutput) ToMwsNetworksErrorMessageArrayOutputWithContext(ctx context.Context) MwsNetworksErrorMessageArrayOutput {
	return o
}

func (o MwsNetworksErrorMessageArrayOutput) Index(i pulumi.IntInput) MwsNetworksErrorMessageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MwsNetworksErrorMessage {
		return vs[0].([]MwsNetworksErrorMessage)[vs[1].(int)]
	}).(MwsNetworksErrorMessageOutput)
}

type MwsWorkspacesNetworkErrorMessage struct {
	ErrorMessage *string `pulumi:"errorMessage"`
	ErrorType    *string `pulumi:"errorType"`
}

// MwsWorkspacesNetworkErrorMessageInput is an input type that accepts MwsWorkspacesNetworkErrorMessageArgs and MwsWorkspacesNetworkErrorMessageOutput values.
// You can construct a concrete instance of `MwsWorkspacesNetworkErrorMessageInput` via:
//
// 		 MwsWorkspacesNetworkErrorMessageArgs{...}
//
type MwsWorkspacesNetworkErrorMessageInput interface {
	pulumi.Input

	ToMwsWorkspacesNetworkErrorMessageOutput() MwsWorkspacesNetworkErrorMessageOutput
	ToMwsWorkspacesNetworkErrorMessageOutputWithContext(context.Context) MwsWorkspacesNetworkErrorMessageOutput
}

type MwsWorkspacesNetworkErrorMessageArgs struct {
	ErrorMessage pulumi.StringPtrInput `pulumi:"errorMessage"`
	ErrorType    pulumi.StringPtrInput `pulumi:"errorType"`
}

func (MwsWorkspacesNetworkErrorMessageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MwsWorkspacesNetworkErrorMessage)(nil)).Elem()
}

func (i MwsWorkspacesNetworkErrorMessageArgs) ToMwsWorkspacesNetworkErrorMessageOutput() MwsWorkspacesNetworkErrorMessageOutput {
	return i.ToMwsWorkspacesNetworkErrorMessageOutputWithContext(context.Background())
}

func (i MwsWorkspacesNetworkErrorMessageArgs) ToMwsWorkspacesNetworkErrorMessageOutputWithContext(ctx context.Context) MwsWorkspacesNetworkErrorMessageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MwsWorkspacesNetworkErrorMessageOutput)
}

// MwsWorkspacesNetworkErrorMessageArrayInput is an input type that accepts MwsWorkspacesNetworkErrorMessageArray and MwsWorkspacesNetworkErrorMessageArrayOutput values.
// You can construct a concrete instance of `MwsWorkspacesNetworkErrorMessageArrayInput` via:
//
// 		 MwsWorkspacesNetworkErrorMessageArray{ MwsWorkspacesNetworkErrorMessageArgs{...} }
//
type MwsWorkspacesNetworkErrorMessageArrayInput interface {
	pulumi.Input

	ToMwsWorkspacesNetworkErrorMessageArrayOutput() MwsWorkspacesNetworkErrorMessageArrayOutput
	ToMwsWorkspacesNetworkErrorMessageArrayOutputWithContext(context.Context) MwsWorkspacesNetworkErrorMessageArrayOutput
}

type MwsWorkspacesNetworkErrorMessageArray []MwsWorkspacesNetworkErrorMessageInput

func (MwsWorkspacesNetworkErrorMessageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MwsWorkspacesNetworkErrorMessage)(nil)).Elem()
}

func (i MwsWorkspacesNetworkErrorMessageArray) ToMwsWorkspacesNetworkErrorMessageArrayOutput() MwsWorkspacesNetworkErrorMessageArrayOutput {
	return i.ToMwsWorkspacesNetworkErrorMessageArrayOutputWithContext(context.Background())
}

func (i MwsWorkspacesNetworkErrorMessageArray) ToMwsWorkspacesNetworkErrorMessageArrayOutputWithContext(ctx context.Context) MwsWorkspacesNetworkErrorMessageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MwsWorkspacesNetworkErrorMessageArrayOutput)
}

type MwsWorkspacesNetworkErrorMessageOutput struct{ *pulumi.OutputState }

func (MwsWorkspacesNetworkErrorMessageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MwsWorkspacesNetworkErrorMessage)(nil)).Elem()
}

func (o MwsWorkspacesNetworkErrorMessageOutput) ToMwsWorkspacesNetworkErrorMessageOutput() MwsWorkspacesNetworkErrorMessageOutput {
	return o
}

func (o MwsWorkspacesNetworkErrorMessageOutput) ToMwsWorkspacesNetworkErrorMessageOutputWithContext(ctx context.Context) MwsWorkspacesNetworkErrorMessageOutput {
	return o
}

func (o MwsWorkspacesNetworkErrorMessageOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MwsWorkspacesNetworkErrorMessage) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o MwsWorkspacesNetworkErrorMessageOutput) ErrorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MwsWorkspacesNetworkErrorMessage) *string { return v.ErrorType }).(pulumi.StringPtrOutput)
}

type MwsWorkspacesNetworkErrorMessageArrayOutput struct{ *pulumi.OutputState }

func (MwsWorkspacesNetworkErrorMessageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MwsWorkspacesNetworkErrorMessage)(nil)).Elem()
}

func (o MwsWorkspacesNetworkErrorMessageArrayOutput) ToMwsWorkspacesNetworkErrorMessageArrayOutput() MwsWorkspacesNetworkErrorMessageArrayOutput {
	return o
}

func (o MwsWorkspacesNetworkErrorMessageArrayOutput) ToMwsWorkspacesNetworkErrorMessageArrayOutputWithContext(ctx context.Context) MwsWorkspacesNetworkErrorMessageArrayOutput {
	return o
}

func (o MwsWorkspacesNetworkErrorMessageArrayOutput) Index(i pulumi.IntInput) MwsWorkspacesNetworkErrorMessageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MwsWorkspacesNetworkErrorMessage {
		return vs[0].([]MwsWorkspacesNetworkErrorMessage)[vs[1].(int)]
	}).(MwsWorkspacesNetworkErrorMessageOutput)
}

func init() {
	pulumi.RegisterOutputType(MwsNetworksErrorMessageOutput{})
	pulumi.RegisterOutputType(MwsNetworksErrorMessageArrayOutput{})
	pulumi.RegisterOutputType(MwsWorkspacesNetworkErrorMessageOutput{})
	pulumi.RegisterOutputType(MwsWorkspacesNetworkErrorMessageArrayOutput{})
}
