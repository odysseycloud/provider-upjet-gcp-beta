//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2025 Upbound Inc. <https://upbound.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateProviderInstanceInitParameters) DeepCopyInto(out *CertificateProviderInstanceInitParameters) {
	*out = *in
	if in.PluginInstance != nil {
		in, out := &in.PluginInstance, &out.PluginInstance
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateProviderInstanceInitParameters.
func (in *CertificateProviderInstanceInitParameters) DeepCopy() *CertificateProviderInstanceInitParameters {
	if in == nil {
		return nil
	}
	out := new(CertificateProviderInstanceInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateProviderInstanceObservation) DeepCopyInto(out *CertificateProviderInstanceObservation) {
	*out = *in
	if in.PluginInstance != nil {
		in, out := &in.PluginInstance, &out.PluginInstance
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateProviderInstanceObservation.
func (in *CertificateProviderInstanceObservation) DeepCopy() *CertificateProviderInstanceObservation {
	if in == nil {
		return nil
	}
	out := new(CertificateProviderInstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateProviderInstanceParameters) DeepCopyInto(out *CertificateProviderInstanceParameters) {
	*out = *in
	if in.PluginInstance != nil {
		in, out := &in.PluginInstance, &out.PluginInstance
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateProviderInstanceParameters.
func (in *CertificateProviderInstanceParameters) DeepCopy() *CertificateProviderInstanceParameters {
	if in == nil {
		return nil
	}
	out := new(CertificateProviderInstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientValidationCAInitParameters) DeepCopyInto(out *ClientValidationCAInitParameters) {
	*out = *in
	if in.CertificateProviderInstance != nil {
		in, out := &in.CertificateProviderInstance, &out.CertificateProviderInstance
		*out = new(CertificateProviderInstanceInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.GRPCEndpoint != nil {
		in, out := &in.GRPCEndpoint, &out.GRPCEndpoint
		*out = new(GRPCEndpointInitParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientValidationCAInitParameters.
func (in *ClientValidationCAInitParameters) DeepCopy() *ClientValidationCAInitParameters {
	if in == nil {
		return nil
	}
	out := new(ClientValidationCAInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientValidationCAObservation) DeepCopyInto(out *ClientValidationCAObservation) {
	*out = *in
	if in.CertificateProviderInstance != nil {
		in, out := &in.CertificateProviderInstance, &out.CertificateProviderInstance
		*out = new(CertificateProviderInstanceObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.GRPCEndpoint != nil {
		in, out := &in.GRPCEndpoint, &out.GRPCEndpoint
		*out = new(GRPCEndpointObservation)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientValidationCAObservation.
func (in *ClientValidationCAObservation) DeepCopy() *ClientValidationCAObservation {
	if in == nil {
		return nil
	}
	out := new(ClientValidationCAObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientValidationCAParameters) DeepCopyInto(out *ClientValidationCAParameters) {
	*out = *in
	if in.CertificateProviderInstance != nil {
		in, out := &in.CertificateProviderInstance, &out.CertificateProviderInstance
		*out = new(CertificateProviderInstanceParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.GRPCEndpoint != nil {
		in, out := &in.GRPCEndpoint, &out.GRPCEndpoint
		*out = new(GRPCEndpointParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientValidationCAParameters.
func (in *ClientValidationCAParameters) DeepCopy() *ClientValidationCAParameters {
	if in == nil {
		return nil
	}
	out := new(ClientValidationCAParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GRPCEndpointInitParameters) DeepCopyInto(out *GRPCEndpointInitParameters) {
	*out = *in
	if in.TargetURI != nil {
		in, out := &in.TargetURI, &out.TargetURI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GRPCEndpointInitParameters.
func (in *GRPCEndpointInitParameters) DeepCopy() *GRPCEndpointInitParameters {
	if in == nil {
		return nil
	}
	out := new(GRPCEndpointInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GRPCEndpointObservation) DeepCopyInto(out *GRPCEndpointObservation) {
	*out = *in
	if in.TargetURI != nil {
		in, out := &in.TargetURI, &out.TargetURI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GRPCEndpointObservation.
func (in *GRPCEndpointObservation) DeepCopy() *GRPCEndpointObservation {
	if in == nil {
		return nil
	}
	out := new(GRPCEndpointObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GRPCEndpointParameters) DeepCopyInto(out *GRPCEndpointParameters) {
	*out = *in
	if in.TargetURI != nil {
		in, out := &in.TargetURI, &out.TargetURI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GRPCEndpointParameters.
func (in *GRPCEndpointParameters) DeepCopy() *GRPCEndpointParameters {
	if in == nil {
		return nil
	}
	out := new(GRPCEndpointParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MtlsPolicyInitParameters) DeepCopyInto(out *MtlsPolicyInitParameters) {
	*out = *in
	if in.ClientValidationCA != nil {
		in, out := &in.ClientValidationCA, &out.ClientValidationCA
		*out = make([]ClientValidationCAInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClientValidationMode != nil {
		in, out := &in.ClientValidationMode, &out.ClientValidationMode
		*out = new(string)
		**out = **in
	}
	if in.ClientValidationTrustConfig != nil {
		in, out := &in.ClientValidationTrustConfig, &out.ClientValidationTrustConfig
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MtlsPolicyInitParameters.
func (in *MtlsPolicyInitParameters) DeepCopy() *MtlsPolicyInitParameters {
	if in == nil {
		return nil
	}
	out := new(MtlsPolicyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MtlsPolicyObservation) DeepCopyInto(out *MtlsPolicyObservation) {
	*out = *in
	if in.ClientValidationCA != nil {
		in, out := &in.ClientValidationCA, &out.ClientValidationCA
		*out = make([]ClientValidationCAObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClientValidationMode != nil {
		in, out := &in.ClientValidationMode, &out.ClientValidationMode
		*out = new(string)
		**out = **in
	}
	if in.ClientValidationTrustConfig != nil {
		in, out := &in.ClientValidationTrustConfig, &out.ClientValidationTrustConfig
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MtlsPolicyObservation.
func (in *MtlsPolicyObservation) DeepCopy() *MtlsPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(MtlsPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MtlsPolicyParameters) DeepCopyInto(out *MtlsPolicyParameters) {
	*out = *in
	if in.ClientValidationCA != nil {
		in, out := &in.ClientValidationCA, &out.ClientValidationCA
		*out = make([]ClientValidationCAParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClientValidationMode != nil {
		in, out := &in.ClientValidationMode, &out.ClientValidationMode
		*out = new(string)
		**out = **in
	}
	if in.ClientValidationTrustConfig != nil {
		in, out := &in.ClientValidationTrustConfig, &out.ClientValidationTrustConfig
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MtlsPolicyParameters.
func (in *MtlsPolicyParameters) DeepCopy() *MtlsPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(MtlsPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerCertificateCertificateProviderInstanceInitParameters) DeepCopyInto(out *ServerCertificateCertificateProviderInstanceInitParameters) {
	*out = *in
	if in.PluginInstance != nil {
		in, out := &in.PluginInstance, &out.PluginInstance
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerCertificateCertificateProviderInstanceInitParameters.
func (in *ServerCertificateCertificateProviderInstanceInitParameters) DeepCopy() *ServerCertificateCertificateProviderInstanceInitParameters {
	if in == nil {
		return nil
	}
	out := new(ServerCertificateCertificateProviderInstanceInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerCertificateCertificateProviderInstanceObservation) DeepCopyInto(out *ServerCertificateCertificateProviderInstanceObservation) {
	*out = *in
	if in.PluginInstance != nil {
		in, out := &in.PluginInstance, &out.PluginInstance
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerCertificateCertificateProviderInstanceObservation.
func (in *ServerCertificateCertificateProviderInstanceObservation) DeepCopy() *ServerCertificateCertificateProviderInstanceObservation {
	if in == nil {
		return nil
	}
	out := new(ServerCertificateCertificateProviderInstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerCertificateCertificateProviderInstanceParameters) DeepCopyInto(out *ServerCertificateCertificateProviderInstanceParameters) {
	*out = *in
	if in.PluginInstance != nil {
		in, out := &in.PluginInstance, &out.PluginInstance
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerCertificateCertificateProviderInstanceParameters.
func (in *ServerCertificateCertificateProviderInstanceParameters) DeepCopy() *ServerCertificateCertificateProviderInstanceParameters {
	if in == nil {
		return nil
	}
	out := new(ServerCertificateCertificateProviderInstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerCertificateGRPCEndpointInitParameters) DeepCopyInto(out *ServerCertificateGRPCEndpointInitParameters) {
	*out = *in
	if in.TargetURI != nil {
		in, out := &in.TargetURI, &out.TargetURI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerCertificateGRPCEndpointInitParameters.
func (in *ServerCertificateGRPCEndpointInitParameters) DeepCopy() *ServerCertificateGRPCEndpointInitParameters {
	if in == nil {
		return nil
	}
	out := new(ServerCertificateGRPCEndpointInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerCertificateGRPCEndpointObservation) DeepCopyInto(out *ServerCertificateGRPCEndpointObservation) {
	*out = *in
	if in.TargetURI != nil {
		in, out := &in.TargetURI, &out.TargetURI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerCertificateGRPCEndpointObservation.
func (in *ServerCertificateGRPCEndpointObservation) DeepCopy() *ServerCertificateGRPCEndpointObservation {
	if in == nil {
		return nil
	}
	out := new(ServerCertificateGRPCEndpointObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerCertificateGRPCEndpointParameters) DeepCopyInto(out *ServerCertificateGRPCEndpointParameters) {
	*out = *in
	if in.TargetURI != nil {
		in, out := &in.TargetURI, &out.TargetURI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerCertificateGRPCEndpointParameters.
func (in *ServerCertificateGRPCEndpointParameters) DeepCopy() *ServerCertificateGRPCEndpointParameters {
	if in == nil {
		return nil
	}
	out := new(ServerCertificateGRPCEndpointParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerCertificateInitParameters) DeepCopyInto(out *ServerCertificateInitParameters) {
	*out = *in
	if in.CertificateProviderInstance != nil {
		in, out := &in.CertificateProviderInstance, &out.CertificateProviderInstance
		*out = new(ServerCertificateCertificateProviderInstanceInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.GRPCEndpoint != nil {
		in, out := &in.GRPCEndpoint, &out.GRPCEndpoint
		*out = new(ServerCertificateGRPCEndpointInitParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerCertificateInitParameters.
func (in *ServerCertificateInitParameters) DeepCopy() *ServerCertificateInitParameters {
	if in == nil {
		return nil
	}
	out := new(ServerCertificateInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerCertificateObservation) DeepCopyInto(out *ServerCertificateObservation) {
	*out = *in
	if in.CertificateProviderInstance != nil {
		in, out := &in.CertificateProviderInstance, &out.CertificateProviderInstance
		*out = new(ServerCertificateCertificateProviderInstanceObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.GRPCEndpoint != nil {
		in, out := &in.GRPCEndpoint, &out.GRPCEndpoint
		*out = new(ServerCertificateGRPCEndpointObservation)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerCertificateObservation.
func (in *ServerCertificateObservation) DeepCopy() *ServerCertificateObservation {
	if in == nil {
		return nil
	}
	out := new(ServerCertificateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerCertificateParameters) DeepCopyInto(out *ServerCertificateParameters) {
	*out = *in
	if in.CertificateProviderInstance != nil {
		in, out := &in.CertificateProviderInstance, &out.CertificateProviderInstance
		*out = new(ServerCertificateCertificateProviderInstanceParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.GRPCEndpoint != nil {
		in, out := &in.GRPCEndpoint, &out.GRPCEndpoint
		*out = new(ServerCertificateGRPCEndpointParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerCertificateParameters.
func (in *ServerCertificateParameters) DeepCopy() *ServerCertificateParameters {
	if in == nil {
		return nil
	}
	out := new(ServerCertificateParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerTLSPolicy) DeepCopyInto(out *ServerTLSPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerTLSPolicy.
func (in *ServerTLSPolicy) DeepCopy() *ServerTLSPolicy {
	if in == nil {
		return nil
	}
	out := new(ServerTLSPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServerTLSPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerTLSPolicyInitParameters) DeepCopyInto(out *ServerTLSPolicyInitParameters) {
	*out = *in
	if in.AllowOpen != nil {
		in, out := &in.AllowOpen, &out.AllowOpen
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.MtlsPolicy != nil {
		in, out := &in.MtlsPolicy, &out.MtlsPolicy
		*out = new(MtlsPolicyInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.ServerCertificate != nil {
		in, out := &in.ServerCertificate, &out.ServerCertificate
		*out = new(ServerCertificateInitParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerTLSPolicyInitParameters.
func (in *ServerTLSPolicyInitParameters) DeepCopy() *ServerTLSPolicyInitParameters {
	if in == nil {
		return nil
	}
	out := new(ServerTLSPolicyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerTLSPolicyList) DeepCopyInto(out *ServerTLSPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServerTLSPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerTLSPolicyList.
func (in *ServerTLSPolicyList) DeepCopy() *ServerTLSPolicyList {
	if in == nil {
		return nil
	}
	out := new(ServerTLSPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServerTLSPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerTLSPolicyObservation) DeepCopyInto(out *ServerTLSPolicyObservation) {
	*out = *in
	if in.AllowOpen != nil {
		in, out := &in.AllowOpen, &out.AllowOpen
		*out = new(bool)
		**out = **in
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EffectiveLabels != nil {
		in, out := &in.EffectiveLabels, &out.EffectiveLabels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MtlsPolicy != nil {
		in, out := &in.MtlsPolicy, &out.MtlsPolicy
		*out = new(MtlsPolicyObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.ServerCertificate != nil {
		in, out := &in.ServerCertificate, &out.ServerCertificate
		*out = new(ServerCertificateObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.TerraformLabels != nil {
		in, out := &in.TerraformLabels, &out.TerraformLabels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerTLSPolicyObservation.
func (in *ServerTLSPolicyObservation) DeepCopy() *ServerTLSPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(ServerTLSPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerTLSPolicyParameters) DeepCopyInto(out *ServerTLSPolicyParameters) {
	*out = *in
	if in.AllowOpen != nil {
		in, out := &in.AllowOpen, &out.AllowOpen
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MtlsPolicy != nil {
		in, out := &in.MtlsPolicy, &out.MtlsPolicy
		*out = new(MtlsPolicyParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.ServerCertificate != nil {
		in, out := &in.ServerCertificate, &out.ServerCertificate
		*out = new(ServerCertificateParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerTLSPolicyParameters.
func (in *ServerTLSPolicyParameters) DeepCopy() *ServerTLSPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(ServerTLSPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerTLSPolicySpec) DeepCopyInto(out *ServerTLSPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerTLSPolicySpec.
func (in *ServerTLSPolicySpec) DeepCopy() *ServerTLSPolicySpec {
	if in == nil {
		return nil
	}
	out := new(ServerTLSPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerTLSPolicyStatus) DeepCopyInto(out *ServerTLSPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerTLSPolicyStatus.
func (in *ServerTLSPolicyStatus) DeepCopy() *ServerTLSPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ServerTLSPolicyStatus)
	in.DeepCopyInto(out)
	return out
}
