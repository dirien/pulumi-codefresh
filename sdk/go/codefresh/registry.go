// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package codefresh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Registry struct {
	pulumi.CustomResourceState

	// _(Optional, Default = false)_ default registry
	Default pulumi.BoolPtrOutput `pulumi:"default"`
	// _(Optional)_ fallback registry
	FallbackRegistry pulumi.StringPtrOutput `pulumi:"fallbackRegistry"`
	Kind             pulumi.StringOutput    `pulumi:"kind"`
	// _(Required)_ some unique name for registry
	Name pulumi.StringOutput `pulumi:"name"`
	// _(Optional, Default = true)_ primary registry
	Primary pulumi.BoolPtrOutput `pulumi:"primary"`
	// _(Required)_ A `spec` block as documented below.
	Spec RegistrySpecOutput `pulumi:"spec"`
}

// NewRegistry registers a new resource with the given unique name, arguments, and options.
func NewRegistry(ctx *pulumi.Context,
	name string, args *RegistryArgs, opts ...pulumi.ResourceOption) (*Registry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Registry
	err := ctx.RegisterResource("codefresh:index/registry:Registry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistry gets an existing Registry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryState, opts ...pulumi.ResourceOption) (*Registry, error) {
	var resource Registry
	err := ctx.ReadResource("codefresh:index/registry:Registry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Registry resources.
type registryState struct {
	// _(Optional, Default = false)_ default registry
	Default *bool `pulumi:"default"`
	// _(Optional)_ fallback registry
	FallbackRegistry *string `pulumi:"fallbackRegistry"`
	Kind             *string `pulumi:"kind"`
	// _(Required)_ some unique name for registry
	Name *string `pulumi:"name"`
	// _(Optional, Default = true)_ primary registry
	Primary *bool `pulumi:"primary"`
	// _(Required)_ A `spec` block as documented below.
	Spec *RegistrySpec `pulumi:"spec"`
}

type RegistryState struct {
	// _(Optional, Default = false)_ default registry
	Default pulumi.BoolPtrInput
	// _(Optional)_ fallback registry
	FallbackRegistry pulumi.StringPtrInput
	Kind             pulumi.StringPtrInput
	// _(Required)_ some unique name for registry
	Name pulumi.StringPtrInput
	// _(Optional, Default = true)_ primary registry
	Primary pulumi.BoolPtrInput
	// _(Required)_ A `spec` block as documented below.
	Spec RegistrySpecPtrInput
}

func (RegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryState)(nil)).Elem()
}

type registryArgs struct {
	// _(Optional, Default = false)_ default registry
	Default *bool `pulumi:"default"`
	// _(Optional)_ fallback registry
	FallbackRegistry *string `pulumi:"fallbackRegistry"`
	// _(Required)_ some unique name for registry
	Name *string `pulumi:"name"`
	// _(Optional, Default = true)_ primary registry
	Primary *bool `pulumi:"primary"`
	// _(Required)_ A `spec` block as documented below.
	Spec RegistrySpec `pulumi:"spec"`
}

// The set of arguments for constructing a Registry resource.
type RegistryArgs struct {
	// _(Optional, Default = false)_ default registry
	Default pulumi.BoolPtrInput
	// _(Optional)_ fallback registry
	FallbackRegistry pulumi.StringPtrInput
	// _(Required)_ some unique name for registry
	Name pulumi.StringPtrInput
	// _(Optional, Default = true)_ primary registry
	Primary pulumi.BoolPtrInput
	// _(Required)_ A `spec` block as documented below.
	Spec RegistrySpecInput
}

func (RegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryArgs)(nil)).Elem()
}

type RegistryInput interface {
	pulumi.Input

	ToRegistryOutput() RegistryOutput
	ToRegistryOutputWithContext(ctx context.Context) RegistryOutput
}

func (*Registry) ElementType() reflect.Type {
	return reflect.TypeOf((**Registry)(nil)).Elem()
}

func (i *Registry) ToRegistryOutput() RegistryOutput {
	return i.ToRegistryOutputWithContext(context.Background())
}

func (i *Registry) ToRegistryOutputWithContext(ctx context.Context) RegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryOutput)
}

// RegistryArrayInput is an input type that accepts RegistryArray and RegistryArrayOutput values.
// You can construct a concrete instance of `RegistryArrayInput` via:
//
//	RegistryArray{ RegistryArgs{...} }
type RegistryArrayInput interface {
	pulumi.Input

	ToRegistryArrayOutput() RegistryArrayOutput
	ToRegistryArrayOutputWithContext(context.Context) RegistryArrayOutput
}

type RegistryArray []RegistryInput

func (RegistryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Registry)(nil)).Elem()
}

func (i RegistryArray) ToRegistryArrayOutput() RegistryArrayOutput {
	return i.ToRegistryArrayOutputWithContext(context.Background())
}

func (i RegistryArray) ToRegistryArrayOutputWithContext(ctx context.Context) RegistryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryArrayOutput)
}

// RegistryMapInput is an input type that accepts RegistryMap and RegistryMapOutput values.
// You can construct a concrete instance of `RegistryMapInput` via:
//
//	RegistryMap{ "key": RegistryArgs{...} }
type RegistryMapInput interface {
	pulumi.Input

	ToRegistryMapOutput() RegistryMapOutput
	ToRegistryMapOutputWithContext(context.Context) RegistryMapOutput
}

type RegistryMap map[string]RegistryInput

func (RegistryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Registry)(nil)).Elem()
}

func (i RegistryMap) ToRegistryMapOutput() RegistryMapOutput {
	return i.ToRegistryMapOutputWithContext(context.Background())
}

func (i RegistryMap) ToRegistryMapOutputWithContext(ctx context.Context) RegistryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryMapOutput)
}

type RegistryOutput struct{ *pulumi.OutputState }

func (RegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Registry)(nil)).Elem()
}

func (o RegistryOutput) ToRegistryOutput() RegistryOutput {
	return o
}

func (o RegistryOutput) ToRegistryOutputWithContext(ctx context.Context) RegistryOutput {
	return o
}

// _(Optional, Default = false)_ default registry
func (o RegistryOutput) Default() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Registry) pulumi.BoolPtrOutput { return v.Default }).(pulumi.BoolPtrOutput)
}

// _(Optional)_ fallback registry
func (o RegistryOutput) FallbackRegistry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringPtrOutput { return v.FallbackRegistry }).(pulumi.StringPtrOutput)
}

func (o RegistryOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// _(Required)_ some unique name for registry
func (o RegistryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// _(Optional, Default = true)_ primary registry
func (o RegistryOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Registry) pulumi.BoolPtrOutput { return v.Primary }).(pulumi.BoolPtrOutput)
}

// _(Required)_ A `spec` block as documented below.
func (o RegistryOutput) Spec() RegistrySpecOutput {
	return o.ApplyT(func(v *Registry) RegistrySpecOutput { return v.Spec }).(RegistrySpecOutput)
}

type RegistryArrayOutput struct{ *pulumi.OutputState }

func (RegistryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Registry)(nil)).Elem()
}

func (o RegistryArrayOutput) ToRegistryArrayOutput() RegistryArrayOutput {
	return o
}

func (o RegistryArrayOutput) ToRegistryArrayOutputWithContext(ctx context.Context) RegistryArrayOutput {
	return o
}

func (o RegistryArrayOutput) Index(i pulumi.IntInput) RegistryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Registry {
		return vs[0].([]*Registry)[vs[1].(int)]
	}).(RegistryOutput)
}

type RegistryMapOutput struct{ *pulumi.OutputState }

func (RegistryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Registry)(nil)).Elem()
}

func (o RegistryMapOutput) ToRegistryMapOutput() RegistryMapOutput {
	return o
}

func (o RegistryMapOutput) ToRegistryMapOutputWithContext(ctx context.Context) RegistryMapOutput {
	return o
}

func (o RegistryMapOutput) MapIndex(k pulumi.StringInput) RegistryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Registry {
		return vs[0].(map[string]*Registry)[vs[1].(string)]
	}).(RegistryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryInput)(nil)).Elem(), &Registry{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryArrayInput)(nil)).Elem(), RegistryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryMapInput)(nil)).Elem(), RegistryMap{})
	pulumi.RegisterOutputType(RegistryOutput{})
	pulumi.RegisterOutputType(RegistryArrayOutput{})
	pulumi.RegisterOutputType(RegistryMapOutput{})
}
