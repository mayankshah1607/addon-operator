// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Addon) DeepCopyInto(out *Addon) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Addon.
func (in *Addon) DeepCopy() *Addon {
	if in == nil {
		return nil
	}
	out := new(Addon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Addon) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonInstallOlmAllNamespaces) DeepCopyInto(out *AddonInstallOlmAllNamespaces) {
	*out = *in
	out.AddonInstallOlmCommon = in.AddonInstallOlmCommon
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonInstallOlmAllNamespaces.
func (in *AddonInstallOlmAllNamespaces) DeepCopy() *AddonInstallOlmAllNamespaces {
	if in == nil {
		return nil
	}
	out := new(AddonInstallOlmAllNamespaces)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonInstallOlmCommon) DeepCopyInto(out *AddonInstallOlmCommon) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonInstallOlmCommon.
func (in *AddonInstallOlmCommon) DeepCopy() *AddonInstallOlmCommon {
	if in == nil {
		return nil
	}
	out := new(AddonInstallOlmCommon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonInstallOlmOwnNamespace) DeepCopyInto(out *AddonInstallOlmOwnNamespace) {
	*out = *in
	out.AddonInstallOlmCommon = in.AddonInstallOlmCommon
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonInstallOlmOwnNamespace.
func (in *AddonInstallOlmOwnNamespace) DeepCopy() *AddonInstallOlmOwnNamespace {
	if in == nil {
		return nil
	}
	out := new(AddonInstallOlmOwnNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonInstallSpec) DeepCopyInto(out *AddonInstallSpec) {
	*out = *in
	if in.OlmAllNamespaces != nil {
		in, out := &in.OlmAllNamespaces, &out.OlmAllNamespaces
		*out = new(AddonInstallOlmAllNamespaces)
		**out = **in
	}
	if in.OlmOwnNamespace != nil {
		in, out := &in.OlmOwnNamespace, &out.OlmOwnNamespace
		*out = new(AddonInstallOlmOwnNamespace)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonInstallSpec.
func (in *AddonInstallSpec) DeepCopy() *AddonInstallSpec {
	if in == nil {
		return nil
	}
	out := new(AddonInstallSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonList) DeepCopyInto(out *AddonList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Addon, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonList.
func (in *AddonList) DeepCopy() *AddonList {
	if in == nil {
		return nil
	}
	out := new(AddonList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddonList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonNamespace) DeepCopyInto(out *AddonNamespace) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonNamespace.
func (in *AddonNamespace) DeepCopy() *AddonNamespace {
	if in == nil {
		return nil
	}
	out := new(AddonNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonSpec) DeepCopyInto(out *AddonSpec) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]AddonNamespace, len(*in))
		copy(*out, *in)
	}
	in.Install.DeepCopyInto(&out.Install)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonSpec.
func (in *AddonSpec) DeepCopy() *AddonSpec {
	if in == nil {
		return nil
	}
	out := new(AddonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonStatus) DeepCopyInto(out *AddonStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonStatus.
func (in *AddonStatus) DeepCopy() *AddonStatus {
	if in == nil {
		return nil
	}
	out := new(AddonStatus)
	in.DeepCopyInto(out)
	return out
}
