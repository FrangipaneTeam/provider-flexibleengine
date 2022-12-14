//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessRule) DeepCopyInto(out *AccessRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessRule.
func (in *AccessRule) DeepCopy() *AccessRule {
	if in == nil {
		return nil
	}
	out := new(AccessRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessRuleList) DeepCopyInto(out *AccessRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccessRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessRuleList.
func (in *AccessRuleList) DeepCopy() *AccessRuleList {
	if in == nil {
		return nil
	}
	out := new(AccessRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessRuleObservation) DeepCopyInto(out *AccessRuleObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessRuleObservation.
func (in *AccessRuleObservation) DeepCopy() *AccessRuleObservation {
	if in == nil {
		return nil
	}
	out := new(AccessRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessRuleParameters) DeepCopyInto(out *AccessRuleParameters) {
	*out = *in
	if in.AccessLevel != nil {
		in, out := &in.AccessLevel, &out.AccessLevel
		*out = new(string)
		**out = **in
	}
	if in.AccessTo != nil {
		in, out := &in.AccessTo, &out.AccessTo
		*out = new(string)
		**out = **in
	}
	if in.AccessType != nil {
		in, out := &in.AccessType, &out.AccessType
		*out = new(string)
		**out = **in
	}
	if in.SfsID != nil {
		in, out := &in.SfsID, &out.SfsID
		*out = new(string)
		**out = **in
	}
	if in.SfsIDRef != nil {
		in, out := &in.SfsIDRef, &out.SfsIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SfsIDSelector != nil {
		in, out := &in.SfsIDSelector, &out.SfsIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessRuleParameters.
func (in *AccessRuleParameters) DeepCopy() *AccessRuleParameters {
	if in == nil {
		return nil
	}
	out := new(AccessRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessRuleSpec) DeepCopyInto(out *AccessRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessRuleSpec.
func (in *AccessRuleSpec) DeepCopy() *AccessRuleSpec {
	if in == nil {
		return nil
	}
	out := new(AccessRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessRuleStatus) DeepCopyInto(out *AccessRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessRuleStatus.
func (in *AccessRuleStatus) DeepCopy() *AccessRuleStatus {
	if in == nil {
		return nil
	}
	out := new(AccessRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessRulesObservation) DeepCopyInto(out *AccessRulesObservation) {
	*out = *in
	if in.AccessLevel != nil {
		in, out := &in.AccessLevel, &out.AccessLevel
		*out = new(string)
		**out = **in
	}
	if in.AccessRuleID != nil {
		in, out := &in.AccessRuleID, &out.AccessRuleID
		*out = new(string)
		**out = **in
	}
	if in.AccessTo != nil {
		in, out := &in.AccessTo, &out.AccessTo
		*out = new(string)
		**out = **in
	}
	if in.AccessType != nil {
		in, out := &in.AccessType, &out.AccessType
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessRulesObservation.
func (in *AccessRulesObservation) DeepCopy() *AccessRulesObservation {
	if in == nil {
		return nil
	}
	out := new(AccessRulesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessRulesParameters) DeepCopyInto(out *AccessRulesParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessRulesParameters.
func (in *AccessRulesParameters) DeepCopy() *AccessRulesParameters {
	if in == nil {
		return nil
	}
	out := new(AccessRulesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystem) DeepCopyInto(out *FileSystem) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystem.
func (in *FileSystem) DeepCopy() *FileSystem {
	if in == nil {
		return nil
	}
	out := new(FileSystem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FileSystem) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemList) DeepCopyInto(out *FileSystemList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FileSystem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemList.
func (in *FileSystemList) DeepCopy() *FileSystemList {
	if in == nil {
		return nil
	}
	out := new(FileSystemList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FileSystemList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemObservation) DeepCopyInto(out *FileSystemObservation) {
	*out = *in
	if in.AccessRuleStatus != nil {
		in, out := &in.AccessRuleStatus, &out.AccessRuleStatus
		*out = new(string)
		**out = **in
	}
	if in.AccessRules != nil {
		in, out := &in.AccessRules, &out.AccessRules
		*out = make([]AccessRulesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExportLocation != nil {
		in, out := &in.ExportLocation, &out.ExportLocation
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ShareAccessID != nil {
		in, out := &in.ShareAccessID, &out.ShareAccessID
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemObservation.
func (in *FileSystemObservation) DeepCopy() *FileSystemObservation {
	if in == nil {
		return nil
	}
	out := new(FileSystemObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemParameters) DeepCopyInto(out *FileSystemParameters) {
	*out = *in
	if in.AccessLevel != nil {
		in, out := &in.AccessLevel, &out.AccessLevel
		*out = new(string)
		**out = **in
	}
	if in.AccessTo != nil {
		in, out := &in.AccessTo, &out.AccessTo
		*out = new(string)
		**out = **in
	}
	if in.AccessType != nil {
		in, out := &in.AccessType, &out.AccessType
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.IsPublic != nil {
		in, out := &in.IsPublic, &out.IsPublic
		*out = new(bool)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ShareProto != nil {
		in, out := &in.ShareProto, &out.ShareProto
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemParameters.
func (in *FileSystemParameters) DeepCopy() *FileSystemParameters {
	if in == nil {
		return nil
	}
	out := new(FileSystemParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemSpec) DeepCopyInto(out *FileSystemSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemSpec.
func (in *FileSystemSpec) DeepCopy() *FileSystemSpec {
	if in == nil {
		return nil
	}
	out := new(FileSystemSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemStatus) DeepCopyInto(out *FileSystemStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemStatus.
func (in *FileSystemStatus) DeepCopy() *FileSystemStatus {
	if in == nil {
		return nil
	}
	out := new(FileSystemStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Turbo) DeepCopyInto(out *Turbo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Turbo.
func (in *Turbo) DeepCopy() *Turbo {
	if in == nil {
		return nil
	}
	out := new(Turbo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Turbo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TurboList) DeepCopyInto(out *TurboList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Turbo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TurboList.
func (in *TurboList) DeepCopy() *TurboList {
	if in == nil {
		return nil
	}
	out := new(TurboList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TurboList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TurboObservation) DeepCopyInto(out *TurboObservation) {
	*out = *in
	if in.AvailableCapacity != nil {
		in, out := &in.AvailableCapacity, &out.AvailableCapacity
		*out = new(string)
		**out = **in
	}
	if in.ExportLocation != nil {
		in, out := &in.ExportLocation, &out.ExportLocation
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TurboObservation.
func (in *TurboObservation) DeepCopy() *TurboObservation {
	if in == nil {
		return nil
	}
	out := new(TurboObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TurboParameters) DeepCopyInto(out *TurboParameters) {
	*out = *in
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.CryptKeyID != nil {
		in, out := &in.CryptKeyID, &out.CryptKeyID
		*out = new(string)
		**out = **in
	}
	if in.CryptKeyIDRef != nil {
		in, out := &in.CryptKeyIDRef, &out.CryptKeyIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.CryptKeyIDSelector != nil {
		in, out := &in.CryptKeyIDSelector, &out.CryptKeyIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupID != nil {
		in, out := &in.SecurityGroupID, &out.SecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupIDRef != nil {
		in, out := &in.SecurityGroupIDRef, &out.SecurityGroupIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityGroupIDSelector != nil {
		in, out := &in.SecurityGroupIDSelector, &out.SecurityGroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ShareProto != nil {
		in, out := &in.ShareProto, &out.ShareProto
		*out = new(string)
		**out = **in
	}
	if in.ShareType != nil {
		in, out := &in.ShareType, &out.ShareType
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDRef != nil {
		in, out := &in.SubnetIDRef, &out.SubnetIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
	if in.VPCIDRef != nil {
		in, out := &in.VPCIDRef, &out.VPCIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.VPCIDSelector != nil {
		in, out := &in.VPCIDSelector, &out.VPCIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TurboParameters.
func (in *TurboParameters) DeepCopy() *TurboParameters {
	if in == nil {
		return nil
	}
	out := new(TurboParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TurboSpec) DeepCopyInto(out *TurboSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TurboSpec.
func (in *TurboSpec) DeepCopy() *TurboSpec {
	if in == nil {
		return nil
	}
	out := new(TurboSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TurboStatus) DeepCopyInto(out *TurboStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TurboStatus.
func (in *TurboStatus) DeepCopy() *TurboStatus {
	if in == nil {
		return nil
	}
	out := new(TurboStatus)
	in.DeepCopyInto(out)
	return out
}
