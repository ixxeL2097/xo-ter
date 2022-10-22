//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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
func (in *ACL) DeepCopyInto(out *ACL) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACL.
func (in *ACL) DeepCopy() *ACL {
	if in == nil {
		return nil
	}
	out := new(ACL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ACL) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACLList) DeepCopyInto(out *ACLList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ACL, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACLList.
func (in *ACLList) DeepCopy() *ACLList {
	if in == nil {
		return nil
	}
	out := new(ACLList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ACLList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACLObservation) DeepCopyInto(out *ACLObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACLObservation.
func (in *ACLObservation) DeepCopy() *ACLObservation {
	if in == nil {
		return nil
	}
	out := new(ACLObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACLParameters) DeepCopyInto(out *ACLParameters) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = new(string)
		**out = **in
	}
	if in.Subject != nil {
		in, out := &in.Subject, &out.Subject
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACLParameters.
func (in *ACLParameters) DeepCopy() *ACLParameters {
	if in == nil {
		return nil
	}
	out := new(ACLParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACLSpec) DeepCopyInto(out *ACLSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACLSpec.
func (in *ACLSpec) DeepCopy() *ACLSpec {
	if in == nil {
		return nil
	}
	out := new(ACLSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACLStatus) DeepCopyInto(out *ACLStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACLStatus.
func (in *ACLStatus) DeepCopy() *ACLStatus {
	if in == nil {
		return nil
	}
	out := new(ACLStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CdromObservation) DeepCopyInto(out *CdromObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CdromObservation.
func (in *CdromObservation) DeepCopy() *CdromObservation {
	if in == nil {
		return nil
	}
	out := new(CdromObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CdromParameters) DeepCopyInto(out *CdromParameters) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CdromParameters.
func (in *CdromParameters) DeepCopy() *CdromParameters {
	if in == nil {
		return nil
	}
	out := new(CdromParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskObservation) DeepCopyInto(out *DiskObservation) {
	*out = *in
	if in.Position != nil {
		in, out := &in.Position, &out.Position
		*out = new(string)
		**out = **in
	}
	if in.VbdID != nil {
		in, out := &in.VbdID, &out.VbdID
		*out = new(string)
		**out = **in
	}
	if in.VdiID != nil {
		in, out := &in.VdiID, &out.VdiID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskObservation.
func (in *DiskObservation) DeepCopy() *DiskObservation {
	if in == nil {
		return nil
	}
	out := new(DiskObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskParameters) DeepCopyInto(out *DiskParameters) {
	*out = *in
	if in.Attached != nil {
		in, out := &in.Attached, &out.Attached
		*out = new(bool)
		**out = **in
	}
	if in.NameDescription != nil {
		in, out := &in.NameDescription, &out.NameDescription
		*out = new(string)
		**out = **in
	}
	if in.NameLabel != nil {
		in, out := &in.NameLabel, &out.NameLabel
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.SrID != nil {
		in, out := &in.SrID, &out.SrID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskParameters.
func (in *DiskParameters) DeepCopy() *DiskParameters {
	if in == nil {
		return nil
	}
	out := new(DiskParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkObservation) DeepCopyInto(out *NetworkObservation) {
	*out = *in
	if in.Device != nil {
		in, out := &in.Device, &out.Device
		*out = new(string)
		**out = **in
	}
	if in.IPv4Addresses != nil {
		in, out := &in.IPv4Addresses, &out.IPv4Addresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IPv6Addresses != nil {
		in, out := &in.IPv6Addresses, &out.IPv6Addresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkObservation.
func (in *NetworkObservation) DeepCopy() *NetworkObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkParameters) DeepCopyInto(out *NetworkParameters) {
	*out = *in
	if in.Attached != nil {
		in, out := &in.Attached, &out.Attached
		*out = new(bool)
		**out = **in
	}
	if in.MacAddress != nil {
		in, out := &in.MacAddress, &out.MacAddress
		*out = new(string)
		**out = **in
	}
	if in.NetworkID != nil {
		in, out := &in.NetworkID, &out.NetworkID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkParameters.
func (in *NetworkParameters) DeepCopy() *NetworkParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VM) DeepCopyInto(out *VM) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VM.
func (in *VM) DeepCopy() *VM {
	if in == nil {
		return nil
	}
	out := new(VM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VM) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMList) DeepCopyInto(out *VMList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMList.
func (in *VMList) DeepCopy() *VMList {
	if in == nil {
		return nil
	}
	out := new(VMList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VMList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMObservation) DeepCopyInto(out *VMObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IPv4Addresses != nil {
		in, out := &in.IPv4Addresses, &out.IPv4Addresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IPv6Addresses != nil {
		in, out := &in.IPv6Addresses, &out.IPv6Addresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PowerState != nil {
		in, out := &in.PowerState, &out.PowerState
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMObservation.
func (in *VMObservation) DeepCopy() *VMObservation {
	if in == nil {
		return nil
	}
	out := new(VMObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMParameters) DeepCopyInto(out *VMParameters) {
	*out = *in
	if in.AffinityHost != nil {
		in, out := &in.AffinityHost, &out.AffinityHost
		*out = new(string)
		**out = **in
	}
	if in.AutoPoweron != nil {
		in, out := &in.AutoPoweron, &out.AutoPoweron
		*out = new(bool)
		**out = **in
	}
	if in.BlockedOperations != nil {
		in, out := &in.BlockedOperations, &out.BlockedOperations
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CPUCap != nil {
		in, out := &in.CPUCap, &out.CPUCap
		*out = new(float64)
		**out = **in
	}
	if in.CPUWeight != nil {
		in, out := &in.CPUWeight, &out.CPUWeight
		*out = new(float64)
		**out = **in
	}
	if in.Cdrom != nil {
		in, out := &in.Cdrom, &out.Cdrom
		*out = make([]CdromParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CloudConfig != nil {
		in, out := &in.CloudConfig, &out.CloudConfig
		*out = new(string)
		**out = **in
	}
	if in.CloudNetworkConfig != nil {
		in, out := &in.CloudNetworkConfig, &out.CloudNetworkConfig
		*out = new(string)
		**out = **in
	}
	if in.CoreOs != nil {
		in, out := &in.CoreOs, &out.CoreOs
		*out = new(bool)
		**out = **in
	}
	if in.Cpus != nil {
		in, out := &in.Cpus, &out.Cpus
		*out = new(float64)
		**out = **in
	}
	if in.Disk != nil {
		in, out := &in.Disk, &out.Disk
		*out = make([]DiskParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExpNestedHvm != nil {
		in, out := &in.ExpNestedHvm, &out.ExpNestedHvm
		*out = new(bool)
		**out = **in
	}
	if in.HighAvailability != nil {
		in, out := &in.HighAvailability, &out.HighAvailability
		*out = new(string)
		**out = **in
	}
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.HvmBootFirmware != nil {
		in, out := &in.HvmBootFirmware, &out.HvmBootFirmware
		*out = new(string)
		**out = **in
	}
	if in.InstallationMethod != nil {
		in, out := &in.InstallationMethod, &out.InstallationMethod
		*out = new(string)
		**out = **in
	}
	if in.MemoryMax != nil {
		in, out := &in.MemoryMax, &out.MemoryMax
		*out = new(float64)
		**out = **in
	}
	if in.NameDescription != nil {
		in, out := &in.NameDescription, &out.NameDescription
		*out = new(string)
		**out = **in
	}
	if in.NameLabel != nil {
		in, out := &in.NameLabel, &out.NameLabel
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = make([]NetworkParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceSet != nil {
		in, out := &in.ResourceSet, &out.ResourceSet
		*out = new(string)
		**out = **in
	}
	if in.StartDelay != nil {
		in, out := &in.StartDelay, &out.StartDelay
		*out = new(float64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(string)
		**out = **in
	}
	if in.Vga != nil {
		in, out := &in.Vga, &out.Vga
		*out = new(string)
		**out = **in
	}
	if in.Videoram != nil {
		in, out := &in.Videoram, &out.Videoram
		*out = new(float64)
		**out = **in
	}
	if in.WaitForIP != nil {
		in, out := &in.WaitForIP, &out.WaitForIP
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMParameters.
func (in *VMParameters) DeepCopy() *VMParameters {
	if in == nil {
		return nil
	}
	out := new(VMParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMSpec) DeepCopyInto(out *VMSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMSpec.
func (in *VMSpec) DeepCopy() *VMSpec {
	if in == nil {
		return nil
	}
	out := new(VMSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMStatus) DeepCopyInto(out *VMStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMStatus.
func (in *VMStatus) DeepCopy() *VMStatus {
	if in == nil {
		return nil
	}
	out := new(VMStatus)
	in.DeepCopyInto(out)
	return out
}
