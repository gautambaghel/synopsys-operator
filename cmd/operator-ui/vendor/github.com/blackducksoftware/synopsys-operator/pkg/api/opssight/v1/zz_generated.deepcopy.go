// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	v2 "github.com/blackducksoftware/synopsys-operator/pkg/api/hub/v2"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hub) DeepCopyInto(out *Hub) {
	*out = *in
	if in.HubSpec != nil {
		in, out := &in.HubSpec, &out.HubSpec
		*out = new(v2.HubSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hub.
func (in *Hub) DeepCopy() *Hub {
	if in == nil {
		return nil
	}
	out := new(Hub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageFacade) DeepCopyInto(out *ImageFacade) {
	*out = *in
	if in.InternalRegistries != nil {
		in, out := &in.InternalRegistries, &out.InternalRegistries
		*out = make([]RegistryAuth, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageFacade.
func (in *ImageFacade) DeepCopy() *ImageFacade {
	if in == nil {
		return nil
	}
	out := new(ImageFacade)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePerceiver) DeepCopyInto(out *ImagePerceiver) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePerceiver.
func (in *ImagePerceiver) DeepCopy() *ImagePerceiver {
	if in == nil {
		return nil
	}
	out := new(ImagePerceiver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsSight) DeepCopyInto(out *OpsSight) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsSight.
func (in *OpsSight) DeepCopy() *OpsSight {
	if in == nil {
		return nil
	}
	out := new(OpsSight)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpsSight) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsSightList) DeepCopyInto(out *OpsSightList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpsSight, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsSightList.
func (in *OpsSightList) DeepCopy() *OpsSightList {
	if in == nil {
		return nil
	}
	out := new(OpsSightList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpsSightList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsSightSpec) DeepCopyInto(out *OpsSightSpec) {
	*out = *in
	if in.Perceptor != nil {
		in, out := &in.Perceptor, &out.Perceptor
		*out = new(Perceptor)
		**out = **in
	}
	if in.ScannerPod != nil {
		in, out := &in.ScannerPod, &out.ScannerPod
		*out = new(ScannerPod)
		(*in).DeepCopyInto(*out)
	}
	if in.Perceiver != nil {
		in, out := &in.Perceiver, &out.Perceiver
		*out = new(Perceiver)
		(*in).DeepCopyInto(*out)
	}
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(Prometheus)
		**out = **in
	}
	if in.Skyfire != nil {
		in, out := &in.Skyfire, &out.Skyfire
		*out = new(Skyfire)
		**out = **in
	}
	if in.Hub != nil {
		in, out := &in.Hub, &out.Hub
		*out = new(Hub)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsSightSpec.
func (in *OpsSightSpec) DeepCopy() *OpsSightSpec {
	if in == nil {
		return nil
	}
	out := new(OpsSightSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsSightStatus) DeepCopyInto(out *OpsSightStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsSightStatus.
func (in *OpsSightStatus) DeepCopy() *OpsSightStatus {
	if in == nil {
		return nil
	}
	out := new(OpsSightStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Perceiver) DeepCopyInto(out *Perceiver) {
	*out = *in
	if in.ImagePerceiver != nil {
		in, out := &in.ImagePerceiver, &out.ImagePerceiver
		*out = new(ImagePerceiver)
		**out = **in
	}
	if in.PodPerceiver != nil {
		in, out := &in.PodPerceiver, &out.PodPerceiver
		*out = new(PodPerceiver)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Perceiver.
func (in *Perceiver) DeepCopy() *Perceiver {
	if in == nil {
		return nil
	}
	out := new(Perceiver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Perceptor) DeepCopyInto(out *Perceptor) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Perceptor.
func (in *Perceptor) DeepCopy() *Perceptor {
	if in == nil {
		return nil
	}
	out := new(Perceptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodPerceiver) DeepCopyInto(out *PodPerceiver) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodPerceiver.
func (in *PodPerceiver) DeepCopy() *PodPerceiver {
	if in == nil {
		return nil
	}
	out := new(PodPerceiver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Prometheus) DeepCopyInto(out *Prometheus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Prometheus.
func (in *Prometheus) DeepCopy() *Prometheus {
	if in == nil {
		return nil
	}
	out := new(Prometheus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryAuth) DeepCopyInto(out *RegistryAuth) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryAuth.
func (in *RegistryAuth) DeepCopy() *RegistryAuth {
	if in == nil {
		return nil
	}
	out := new(RegistryAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Scanner) DeepCopyInto(out *Scanner) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scanner.
func (in *Scanner) DeepCopy() *Scanner {
	if in == nil {
		return nil
	}
	out := new(Scanner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScannerPod) DeepCopyInto(out *ScannerPod) {
	*out = *in
	if in.Scanner != nil {
		in, out := &in.Scanner, &out.Scanner
		*out = new(Scanner)
		**out = **in
	}
	if in.ImageFacade != nil {
		in, out := &in.ImageFacade, &out.ImageFacade
		*out = new(ImageFacade)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScannerPod.
func (in *ScannerPod) DeepCopy() *ScannerPod {
	if in == nil {
		return nil
	}
	out := new(ScannerPod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Skyfire) DeepCopyInto(out *Skyfire) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Skyfire.
func (in *Skyfire) DeepCopy() *Skyfire {
	if in == nil {
		return nil
	}
	out := new(Skyfire)
	in.DeepCopyInto(out)
	return out
}