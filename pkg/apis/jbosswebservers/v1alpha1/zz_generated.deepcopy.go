// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JbossWebImageSpec) DeepCopyInto(out *JbossWebImageSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JbossWebImageSpec.
func (in *JbossWebImageSpec) DeepCopy() *JbossWebImageSpec {
	if in == nil {
		return nil
	}
	out := new(JbossWebImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JbossWebImageStreamSpec) DeepCopyInto(out *JbossWebImageStreamSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JbossWebImageStreamSpec.
func (in *JbossWebImageStreamSpec) DeepCopy() *JbossWebImageStreamSpec {
	if in == nil {
		return nil
	}
	out := new(JbossWebImageStreamSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JbossWebServer) DeepCopyInto(out *JbossWebServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JbossWebServer.
func (in *JbossWebServer) DeepCopy() *JbossWebServer {
	if in == nil {
		return nil
	}
	out := new(JbossWebServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JbossWebServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JbossWebServer53HealthCheckSpec) DeepCopyInto(out *JbossWebServer53HealthCheckSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JbossWebServer53HealthCheckSpec.
func (in *JbossWebServer53HealthCheckSpec) DeepCopy() *JbossWebServer53HealthCheckSpec {
	if in == nil {
		return nil
	}
	out := new(JbossWebServer53HealthCheckSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JbossWebServerHealthCheckSpec) DeepCopyInto(out *JbossWebServerHealthCheckSpec) {
	*out = *in
	if in.JbossWebServer53HealthCheck != nil {
		in, out := &in.JbossWebServer53HealthCheck, &out.JbossWebServer53HealthCheck
		*out = new(JbossWebServer53HealthCheckSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JbossWebServerHealthCheckSpec.
func (in *JbossWebServerHealthCheckSpec) DeepCopy() *JbossWebServerHealthCheckSpec {
	if in == nil {
		return nil
	}
	out := new(JbossWebServerHealthCheckSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JbossWebServerList) DeepCopyInto(out *JbossWebServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JbossWebServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JbossWebServerList.
func (in *JbossWebServerList) DeepCopy() *JbossWebServerList {
	if in == nil {
		return nil
	}
	out := new(JbossWebServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JbossWebServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JbossWebServerSpec) DeepCopyInto(out *JbossWebServerSpec) {
	*out = *in
	if in.JbossWebImage != nil {
		in, out := &in.JbossWebImage, &out.JbossWebImage
		*out = new(JbossWebImageSpec)
		**out = **in
	}
	if in.JbossWebImageStream != nil {
		in, out := &in.JbossWebImageStream, &out.JbossWebImageStream
		*out = new(JbossWebImageStreamSpec)
		**out = **in
	}
	if in.JbossWebSources != nil {
		in, out := &in.JbossWebSources, &out.JbossWebSources
		*out = new(JbossWebSourcesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.JbossWebServerHealthCheck != nil {
		in, out := &in.JbossWebServerHealthCheck, &out.JbossWebServerHealthCheck
		*out = new(JbossWebServerHealthCheckSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JbossWebServerSpec.
func (in *JbossWebServerSpec) DeepCopy() *JbossWebServerSpec {
	if in == nil {
		return nil
	}
	out := new(JbossWebServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JbossWebServerStatus) DeepCopyInto(out *JbossWebServerStatus) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]PodStatus, len(*in))
		copy(*out, *in)
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JbossWebServerStatus.
func (in *JbossWebServerStatus) DeepCopy() *JbossWebServerStatus {
	if in == nil {
		return nil
	}
	out := new(JbossWebServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JbossWebSourcesParamsSpec) DeepCopyInto(out *JbossWebSourcesParamsSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JbossWebSourcesParamsSpec.
func (in *JbossWebSourcesParamsSpec) DeepCopy() *JbossWebSourcesParamsSpec {
	if in == nil {
		return nil
	}
	out := new(JbossWebSourcesParamsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JbossWebSourcesSpec) DeepCopyInto(out *JbossWebSourcesSpec) {
	*out = *in
	if in.JbossWebSourcesParams != nil {
		in, out := &in.JbossWebSourcesParams, &out.JbossWebSourcesParams
		*out = new(JbossWebSourcesParamsSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JbossWebSourcesSpec.
func (in *JbossWebSourcesSpec) DeepCopy() *JbossWebSourcesSpec {
	if in == nil {
		return nil
	}
	out := new(JbossWebSourcesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodStatus) DeepCopyInto(out *PodStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodStatus.
func (in *PodStatus) DeepCopy() *PodStatus {
	if in == nil {
		return nil
	}
	out := new(PodStatus)
	in.DeepCopyInto(out)
	return out
}
