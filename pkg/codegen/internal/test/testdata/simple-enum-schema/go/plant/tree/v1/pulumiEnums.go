// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Diameter pulumi.Float64

const (
	DiameterSixinch    = Diameter(6)
	DiameterTwelveinch = Diameter(12)
)

func (Diameter) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.Float64)(nil)).Elem()
}

func (e Diameter) ToFloat64Output() pulumi.Float64Output {
	return pulumi.ToOutput(pulumi.Float64(e)).(pulumi.Float64Output)
}

func (e Diameter) ToFloat64OutputWithContext(ctx context.Context) pulumi.Float64Output {
	return pulumi.ToOutputWithContext(ctx, pulumi.Float64(e)).(pulumi.Float64Output)
}

func (e Diameter) ToFloat64PtrOutput() pulumi.Float64PtrOutput {
	return pulumi.Float64(e).ToFloat64PtrOutputWithContext(context.Background())
}

func (e Diameter) ToFloat64PtrOutputWithContext(ctx context.Context) pulumi.Float64PtrOutput {
	return pulumi.Float64(e).ToFloat64OutputWithContext(ctx).ToFloat64PtrOutputWithContext(ctx)
}

type Farm pulumi.String

const (
	Farm_Pulumi_Planters_Inc_ = Farm("Pulumi Planters Inc.")
	Farm_Plants_R_Us          = Farm("Plants'R'Us")
)

func (Farm) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e Farm) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Farm) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Farm) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Farm) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// types of rubber trees
type RubberTreeVariety pulumi.String

const (
	// A burgundy rubber tree.
	RubberTreeVarietyBurgundy = RubberTreeVariety("Burgundy")
	// A ruby rubber tree.
	RubberTreeVarietyRuby = RubberTreeVariety("Ruby")
	// A tineke rubber tree.
	RubberTreeVarietyTineke = RubberTreeVariety("Tineke")
)

func (RubberTreeVariety) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e RubberTreeVariety) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RubberTreeVariety) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RubberTreeVariety) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RubberTreeVariety) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// RubberTreeVarietyArrayInput is an input type that accepts RubberTreeVarietyArray and RubberTreeVarietyArrayOutput values.
// You can construct a concrete instance of `RubberTreeVarietyArrayInput` via:
//
//          RubberTreeVarietyArray{ RubberTreeVarietyArgs{...} }
type RubberTreeVarietyArrayInput interface {
	pulumi.Input

	ToRubberTreeVarietyArrayOutput() RubberTreeVarietyArrayOutput
	ToRubberTreeVarietyArrayOutputWithContext(context.Context) RubberTreeVarietyArrayOutput
}

type RubberTreeVarietyArray []RubberTreeVariety

func (RubberTreeVarietyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RubberTreeVariety)(nil)).Elem()
}

func (i RubberTreeVarietyArray) ToRubberTreeVarietyArrayOutput() RubberTreeVarietyArrayOutput {
	return i.ToRubberTreeVarietyArrayOutputWithContext(context.Background())
}

func (i RubberTreeVarietyArray) ToRubberTreeVarietyArrayOutputWithContext(ctx context.Context) RubberTreeVarietyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RubberTreeVarietyArrayOutput)
}

type RubberTreeVarietyArrayOutput struct{ *pulumi.OutputState }

func (RubberTreeVarietyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RubberTreeVariety)(nil)).Elem()
}

func (o RubberTreeVarietyArrayOutput) ToRubberTreeVarietyArrayOutput() RubberTreeVarietyArrayOutput {
	return o
}

func (o RubberTreeVarietyArrayOutput) ToRubberTreeVarietyArrayOutputWithContext(ctx context.Context) RubberTreeVarietyArrayOutput {
	return o
}

func (o RubberTreeVarietyArrayOutput) Index(i pulumi.IntInput) pulumi.StringOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) pulumi.StringOutput {
		return vs[0].([]RubberTreeVariety)[vs[1].(int)].ToStringOutput()
	}).(pulumi.StringOutput)
}

type TreeSize pulumi.String

const (
	TreeSizeSmall  = TreeSize("small")
	TreeSizeMedium = TreeSize("medium")
	TreeSizeLarge  = TreeSize("large")
)

func (TreeSize) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e TreeSize) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TreeSize) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TreeSize) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TreeSize) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// TreeSizeMapInput is an input type that accepts TreeSizeMap and TreeSizeMapOutput values.
// You can construct a concrete instance of `TreeSizeMapInput` via:
//
//          TreeSizeMap{ "key": TreeSizeArgs{...} }
type TreeSizeMapInput interface {
	pulumi.Input

	ToTreeSizeMapOutput() TreeSizeMapOutput
	ToTreeSizeMapOutputWithContext(context.Context) TreeSizeMapOutput
}

type TreeSizeMap map[string]TreeSize

func (TreeSizeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TreeSize)(nil)).Elem()
}

func (i TreeSizeMap) ToTreeSizeMapOutput() TreeSizeMapOutput {
	return i.ToTreeSizeMapOutputWithContext(context.Background())
}

func (i TreeSizeMap) ToTreeSizeMapOutputWithContext(ctx context.Context) TreeSizeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TreeSizeMapOutput)
}

type TreeSizeMapOutput struct{ *pulumi.OutputState }

func (TreeSizeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TreeSize)(nil)).Elem()
}

func (o TreeSizeMapOutput) ToTreeSizeMapOutput() TreeSizeMapOutput {
	return o
}

func (o TreeSizeMapOutput) ToTreeSizeMapOutputWithContext(ctx context.Context) TreeSizeMapOutput {
	return o
}

func (o TreeSizeMapOutput) MapIndex(k pulumi.StringInput) pulumi.StringOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) pulumi.StringOutput {
		return vs[0].(map[string]TreeSize)[vs[1].(string)].ToStringOutput()
	}).(pulumi.StringOutput)
}