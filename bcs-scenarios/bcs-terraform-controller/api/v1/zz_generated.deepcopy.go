//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendConfigsReference) DeepCopyInto(out *BackendConfigsReference) {
	*out = *in
	if in.Keys != nil {
		in, out := &in.Keys, &out.Keys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendConfigsReference.
func (in *BackendConfigsReference) DeepCopy() *BackendConfigsReference {
	if in == nil {
		return nil
	}
	out := new(BackendConfigsReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitRepository) DeepCopyInto(out *GitRepository) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitRepository.
func (in *GitRepository) DeepCopy() *GitRepository {
	if in == nil {
		return nil
	}
	out := new(GitRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanStatus) DeepCopyInto(out *PlanStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanStatus.
func (in *PlanStatus) DeepCopy() *PlanStatus {
	if in == nil {
		return nil
	}
	out := new(PlanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Terraform) DeepCopyInto(out *Terraform) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Terraform.
func (in *Terraform) DeepCopy() *Terraform {
	if in == nil {
		return nil
	}
	out := new(Terraform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Terraform) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformList) DeepCopyInto(out *TerraformList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Terraform, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformList.
func (in *TerraformList) DeepCopy() *TerraformList {
	if in == nil {
		return nil
	}
	out := new(TerraformList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TerraformList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformSpec) DeepCopyInto(out *TerraformSpec) {
	*out = *in
	if in.BackendConfigsFrom != nil {
		in, out := &in.BackendConfigsFrom, &out.BackendConfigsFrom
		*out = make([]BackendConfigsReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Repository = in.Repository
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformSpec.
func (in *TerraformSpec) DeepCopy() *TerraformSpec {
	if in == nil {
		return nil
	}
	out := new(TerraformSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformStatus) DeepCopyInto(out *TerraformStatus) {
	*out = *in
	if in.LastPlanAt != nil {
		in, out := &in.LastPlanAt, &out.LastPlanAt
		*out = (*in).DeepCopy()
	}
	if in.LastAppliedAt != nil {
		in, out := &in.LastAppliedAt, &out.LastAppliedAt
		*out = (*in).DeepCopy()
	}
	out.Plan = in.Plan
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformStatus.
func (in *TerraformStatus) DeepCopy() *TerraformStatus {
	if in == nil {
		return nil
	}
	out := new(TerraformStatus)
	in.DeepCopyInto(out)
	return out
}
