//go:build !ignore_autogenerated

/*
Copyright 2025 Akos Orban.

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HotReloadProject) DeepCopyInto(out *HotReloadProject) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HotReloadProject.
func (in *HotReloadProject) DeepCopy() *HotReloadProject {
	if in == nil {
		return nil
	}
	out := new(HotReloadProject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HotReloadProject) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HotReloadProjectList) DeepCopyInto(out *HotReloadProjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HotReloadProject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HotReloadProjectList.
func (in *HotReloadProjectList) DeepCopy() *HotReloadProjectList {
	if in == nil {
		return nil
	}
	out := new(HotReloadProjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HotReloadProjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HotReloadProjectSpec) DeepCopyInto(out *HotReloadProjectSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HotReloadProjectSpec.
func (in *HotReloadProjectSpec) DeepCopy() *HotReloadProjectSpec {
	if in == nil {
		return nil
	}
	out := new(HotReloadProjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HotReloadProjectStatus) DeepCopyInto(out *HotReloadProjectStatus) {
	*out = *in
	if in.LastDeploymentTime != nil {
		in, out := &in.LastDeploymentTime, &out.LastDeploymentTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HotReloadProjectStatus.
func (in *HotReloadProjectStatus) DeepCopy() *HotReloadProjectStatus {
	if in == nil {
		return nil
	}
	out := new(HotReloadProjectStatus)
	in.DeepCopyInto(out)
	return out
}
