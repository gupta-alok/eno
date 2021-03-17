// +build !ignore_autogenerated

/*


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
func (in *L2BridgeDomain) DeepCopyInto(out *L2BridgeDomain) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L2BridgeDomain.
func (in *L2BridgeDomain) DeepCopy() *L2BridgeDomain {
	if in == nil {
		return nil
	}
	out := new(L2BridgeDomain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *L2BridgeDomain) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L2BridgeDomainList) DeepCopyInto(out *L2BridgeDomainList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]L2BridgeDomain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L2BridgeDomainList.
func (in *L2BridgeDomainList) DeepCopy() *L2BridgeDomainList {
	if in == nil {
		return nil
	}
	out := new(L2BridgeDomainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *L2BridgeDomainList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L2BridgeDomainSpec) DeepCopyInto(out *L2BridgeDomainSpec) {
	*out = *in
	if in.ConnectionPoints != nil {
		in, out := &in.ConnectionPoints, &out.ConnectionPoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L2BridgeDomainSpec.
func (in *L2BridgeDomainSpec) DeepCopy() *L2BridgeDomainSpec {
	if in == nil {
		return nil
	}
	out := new(L2BridgeDomainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L2BridgeDomainStatus) DeepCopyInto(out *L2BridgeDomainStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L2BridgeDomainStatus.
func (in *L2BridgeDomainStatus) DeepCopy() *L2BridgeDomainStatus {
	if in == nil {
		return nil
	}
	out := new(L2BridgeDomainStatus)
	in.DeepCopyInto(out)
	return out
}