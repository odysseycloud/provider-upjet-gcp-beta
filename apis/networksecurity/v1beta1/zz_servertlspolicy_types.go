// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CertificateProviderInstanceInitParameters struct {

	// Plugin instance name, used to locate and load CertificateProvider instance configuration. Set to "google_cloud_private_spiffe" to use Certificate Authority Service certificate provider instance.
	PluginInstance *string `json:"pluginInstance,omitempty" tf:"plugin_instance,omitempty"`
}

type CertificateProviderInstanceObservation struct {

	// Plugin instance name, used to locate and load CertificateProvider instance configuration. Set to "google_cloud_private_spiffe" to use Certificate Authority Service certificate provider instance.
	PluginInstance *string `json:"pluginInstance,omitempty" tf:"plugin_instance,omitempty"`
}

type CertificateProviderInstanceParameters struct {

	// Plugin instance name, used to locate and load CertificateProvider instance configuration. Set to "google_cloud_private_spiffe" to use Certificate Authority Service certificate provider instance.
	// +kubebuilder:validation:Optional
	PluginInstance *string `json:"pluginInstance" tf:"plugin_instance,omitempty"`
}

type ClientValidationCAInitParameters struct {

	// Optional if policy is to be used with Traffic Director. For external HTTPS load balancer must be empty.
	// Defines a mechanism to provision server identity (public and private keys). Cannot be combined with allowOpen as a permissive mode that allows both plain text and TLS is not supported.
	// Structure is documented below.
	CertificateProviderInstance *CertificateProviderInstanceInitParameters `json:"certificateProviderInstance,omitempty" tf:"certificate_provider_instance,omitempty"`

	// gRPC specific configuration to access the gRPC server to obtain the cert and private key.
	// Structure is documented below.
	GRPCEndpoint *GRPCEndpointInitParameters `json:"grpcEndpoint,omitempty" tf:"grpc_endpoint,omitempty"`
}

type ClientValidationCAObservation struct {

	// Optional if policy is to be used with Traffic Director. For external HTTPS load balancer must be empty.
	// Defines a mechanism to provision server identity (public and private keys). Cannot be combined with allowOpen as a permissive mode that allows both plain text and TLS is not supported.
	// Structure is documented below.
	CertificateProviderInstance *CertificateProviderInstanceObservation `json:"certificateProviderInstance,omitempty" tf:"certificate_provider_instance,omitempty"`

	// gRPC specific configuration to access the gRPC server to obtain the cert and private key.
	// Structure is documented below.
	GRPCEndpoint *GRPCEndpointObservation `json:"grpcEndpoint,omitempty" tf:"grpc_endpoint,omitempty"`
}

type ClientValidationCAParameters struct {

	// Optional if policy is to be used with Traffic Director. For external HTTPS load balancer must be empty.
	// Defines a mechanism to provision server identity (public and private keys). Cannot be combined with allowOpen as a permissive mode that allows both plain text and TLS is not supported.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	CertificateProviderInstance *CertificateProviderInstanceParameters `json:"certificateProviderInstance,omitempty" tf:"certificate_provider_instance,omitempty"`

	// gRPC specific configuration to access the gRPC server to obtain the cert and private key.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	GRPCEndpoint *GRPCEndpointParameters `json:"grpcEndpoint,omitempty" tf:"grpc_endpoint,omitempty"`
}

type GRPCEndpointInitParameters struct {

	// The target URI of the gRPC endpoint. Only UDS path is supported, and should start with "unix:".
	TargetURI *string `json:"targetUri,omitempty" tf:"target_uri,omitempty"`
}

type GRPCEndpointObservation struct {

	// The target URI of the gRPC endpoint. Only UDS path is supported, and should start with "unix:".
	TargetURI *string `json:"targetUri,omitempty" tf:"target_uri,omitempty"`
}

type GRPCEndpointParameters struct {

	// The target URI of the gRPC endpoint. Only UDS path is supported, and should start with "unix:".
	// +kubebuilder:validation:Optional
	TargetURI *string `json:"targetUri" tf:"target_uri,omitempty"`
}

type MtlsPolicyInitParameters struct {

	// Required if the policy is to be used with Traffic Director. For external HTTPS load balancers it must be empty.
	// Defines the mechanism to obtain the Certificate Authority certificate to validate the client certificate.
	// Structure is documented below.
	ClientValidationCA []ClientValidationCAInitParameters `json:"clientValidationCa,omitempty" tf:"client_validation_ca,omitempty"`

	// When the client presents an invalid certificate or no certificate to the load balancer, the clientValidationMode specifies how the client connection is handled.
	// Required if the policy is to be used with the external HTTPS load balancing. For Traffic Director it must be empty.
	// Possible values are: CLIENT_VALIDATION_MODE_UNSPECIFIED, ALLOW_INVALID_OR_MISSING_CLIENT_CERT, REJECT_INVALID.
	ClientValidationMode *string `json:"clientValidationMode,omitempty" tf:"client_validation_mode,omitempty"`

	// Reference to the TrustConfig from certificatemanager.googleapis.com namespace.
	// If specified, the chain validation will be performed against certificates configured in the given TrustConfig.
	// Allowed only if the policy is to be used with external HTTPS load balancers.
	ClientValidationTrustConfig *string `json:"clientValidationTrustConfig,omitempty" tf:"client_validation_trust_config,omitempty"`
}

type MtlsPolicyObservation struct {

	// Required if the policy is to be used with Traffic Director. For external HTTPS load balancers it must be empty.
	// Defines the mechanism to obtain the Certificate Authority certificate to validate the client certificate.
	// Structure is documented below.
	ClientValidationCA []ClientValidationCAObservation `json:"clientValidationCa,omitempty" tf:"client_validation_ca,omitempty"`

	// When the client presents an invalid certificate or no certificate to the load balancer, the clientValidationMode specifies how the client connection is handled.
	// Required if the policy is to be used with the external HTTPS load balancing. For Traffic Director it must be empty.
	// Possible values are: CLIENT_VALIDATION_MODE_UNSPECIFIED, ALLOW_INVALID_OR_MISSING_CLIENT_CERT, REJECT_INVALID.
	ClientValidationMode *string `json:"clientValidationMode,omitempty" tf:"client_validation_mode,omitempty"`

	// Reference to the TrustConfig from certificatemanager.googleapis.com namespace.
	// If specified, the chain validation will be performed against certificates configured in the given TrustConfig.
	// Allowed only if the policy is to be used with external HTTPS load balancers.
	ClientValidationTrustConfig *string `json:"clientValidationTrustConfig,omitempty" tf:"client_validation_trust_config,omitempty"`
}

type MtlsPolicyParameters struct {

	// Required if the policy is to be used with Traffic Director. For external HTTPS load balancers it must be empty.
	// Defines the mechanism to obtain the Certificate Authority certificate to validate the client certificate.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ClientValidationCA []ClientValidationCAParameters `json:"clientValidationCa,omitempty" tf:"client_validation_ca,omitempty"`

	// When the client presents an invalid certificate or no certificate to the load balancer, the clientValidationMode specifies how the client connection is handled.
	// Required if the policy is to be used with the external HTTPS load balancing. For Traffic Director it must be empty.
	// Possible values are: CLIENT_VALIDATION_MODE_UNSPECIFIED, ALLOW_INVALID_OR_MISSING_CLIENT_CERT, REJECT_INVALID.
	// +kubebuilder:validation:Optional
	ClientValidationMode *string `json:"clientValidationMode,omitempty" tf:"client_validation_mode,omitempty"`

	// Reference to the TrustConfig from certificatemanager.googleapis.com namespace.
	// If specified, the chain validation will be performed against certificates configured in the given TrustConfig.
	// Allowed only if the policy is to be used with external HTTPS load balancers.
	// +kubebuilder:validation:Optional
	ClientValidationTrustConfig *string `json:"clientValidationTrustConfig,omitempty" tf:"client_validation_trust_config,omitempty"`
}

type ServerCertificateCertificateProviderInstanceInitParameters struct {

	// Plugin instance name, used to locate and load CertificateProvider instance configuration. Set to "google_cloud_private_spiffe" to use Certificate Authority Service certificate provider instance.
	PluginInstance *string `json:"pluginInstance,omitempty" tf:"plugin_instance,omitempty"`
}

type ServerCertificateCertificateProviderInstanceObservation struct {

	// Plugin instance name, used to locate and load CertificateProvider instance configuration. Set to "google_cloud_private_spiffe" to use Certificate Authority Service certificate provider instance.
	PluginInstance *string `json:"pluginInstance,omitempty" tf:"plugin_instance,omitempty"`
}

type ServerCertificateCertificateProviderInstanceParameters struct {

	// Plugin instance name, used to locate and load CertificateProvider instance configuration. Set to "google_cloud_private_spiffe" to use Certificate Authority Service certificate provider instance.
	// +kubebuilder:validation:Optional
	PluginInstance *string `json:"pluginInstance" tf:"plugin_instance,omitempty"`
}

type ServerCertificateGRPCEndpointInitParameters struct {

	// The target URI of the gRPC endpoint. Only UDS path is supported, and should start with "unix:".
	TargetURI *string `json:"targetUri,omitempty" tf:"target_uri,omitempty"`
}

type ServerCertificateGRPCEndpointObservation struct {

	// The target URI of the gRPC endpoint. Only UDS path is supported, and should start with "unix:".
	TargetURI *string `json:"targetUri,omitempty" tf:"target_uri,omitempty"`
}

type ServerCertificateGRPCEndpointParameters struct {

	// The target URI of the gRPC endpoint. Only UDS path is supported, and should start with "unix:".
	// +kubebuilder:validation:Optional
	TargetURI *string `json:"targetUri" tf:"target_uri,omitempty"`
}

type ServerCertificateInitParameters struct {

	// Optional if policy is to be used with Traffic Director. For external HTTPS load balancer must be empty.
	// Defines a mechanism to provision server identity (public and private keys). Cannot be combined with allowOpen as a permissive mode that allows both plain text and TLS is not supported.
	// Structure is documented below.
	CertificateProviderInstance *ServerCertificateCertificateProviderInstanceInitParameters `json:"certificateProviderInstance,omitempty" tf:"certificate_provider_instance,omitempty"`

	// gRPC specific configuration to access the gRPC server to obtain the cert and private key.
	// Structure is documented below.
	GRPCEndpoint *ServerCertificateGRPCEndpointInitParameters `json:"grpcEndpoint,omitempty" tf:"grpc_endpoint,omitempty"`
}

type ServerCertificateObservation struct {

	// Optional if policy is to be used with Traffic Director. For external HTTPS load balancer must be empty.
	// Defines a mechanism to provision server identity (public and private keys). Cannot be combined with allowOpen as a permissive mode that allows both plain text and TLS is not supported.
	// Structure is documented below.
	CertificateProviderInstance *ServerCertificateCertificateProviderInstanceObservation `json:"certificateProviderInstance,omitempty" tf:"certificate_provider_instance,omitempty"`

	// gRPC specific configuration to access the gRPC server to obtain the cert and private key.
	// Structure is documented below.
	GRPCEndpoint *ServerCertificateGRPCEndpointObservation `json:"grpcEndpoint,omitempty" tf:"grpc_endpoint,omitempty"`
}

type ServerCertificateParameters struct {

	// Optional if policy is to be used with Traffic Director. For external HTTPS load balancer must be empty.
	// Defines a mechanism to provision server identity (public and private keys). Cannot be combined with allowOpen as a permissive mode that allows both plain text and TLS is not supported.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	CertificateProviderInstance *ServerCertificateCertificateProviderInstanceParameters `json:"certificateProviderInstance,omitempty" tf:"certificate_provider_instance,omitempty"`

	// gRPC specific configuration to access the gRPC server to obtain the cert and private key.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	GRPCEndpoint *ServerCertificateGRPCEndpointParameters `json:"grpcEndpoint,omitempty" tf:"grpc_endpoint,omitempty"`
}

type ServerTLSPolicyInitParameters struct {

	// This field applies only for Traffic Director policies. It is must be set to false for external HTTPS load balancer policies.
	// Determines if server allows plaintext connections. If set to true, server allows plain text connections. By default, it is set to false. This setting is not exclusive of other encryption modes. For example, if allowOpen and mtlsPolicy are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to confirm compatibility.
	// Consider using it if you wish to upgrade in place your deployment to TLS while having mixed TLS and non-TLS traffic reaching port :80.
	AllowOpen *bool `json:"allowOpen,omitempty" tf:"allow_open,omitempty"`

	// A free-text description of the resource. Max length 1024 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Set of label tags associated with the ServerTlsPolicy resource.
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// This field is required if the policy is used with external HTTPS load balancers. This field can be empty for Traffic Director.
	// Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS - mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If allowOpen and mtlsPolicy are set, server allows both plain text and mTLS connections.
	// Structure is documented below.
	MtlsPolicy *MtlsPolicyInitParameters `json:"mtlsPolicy,omitempty" tf:"mtls_policy,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence of this dictates mTLS.
	// Structure is documented below.
	ServerCertificate *ServerCertificateInitParameters `json:"serverCertificate,omitempty" tf:"server_certificate,omitempty"`
}

type ServerTLSPolicyObservation struct {

	// This field applies only for Traffic Director policies. It is must be set to false for external HTTPS load balancer policies.
	// Determines if server allows plaintext connections. If set to true, server allows plain text connections. By default, it is set to false. This setting is not exclusive of other encryption modes. For example, if allowOpen and mtlsPolicy are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to confirm compatibility.
	// Consider using it if you wish to upgrade in place your deployment to TLS while having mixed TLS and non-TLS traffic reaching port :80.
	AllowOpen *bool `json:"allowOpen,omitempty" tf:"allow_open,omitempty"`

	// Time the ServerTlsPolicy was created in UTC.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// A free-text description of the resource. Max length 1024 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +mapType=granular
	EffectiveLabels map[string]*string `json:"effectiveLabels,omitempty" tf:"effective_labels,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/serverTlsPolicies/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Set of label tags associated with the ServerTlsPolicy resource.
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The location of the server tls policy.
	// The default value is global.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// This field is required if the policy is used with external HTTPS load balancers. This field can be empty for Traffic Director.
	// Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS - mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If allowOpen and mtlsPolicy are set, server allows both plain text and mTLS connections.
	// Structure is documented below.
	MtlsPolicy *MtlsPolicyObservation `json:"mtlsPolicy,omitempty" tf:"mtls_policy,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence of this dictates mTLS.
	// Structure is documented below.
	ServerCertificate *ServerCertificateObservation `json:"serverCertificate,omitempty" tf:"server_certificate,omitempty"`

	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	// +mapType=granular
	TerraformLabels map[string]*string `json:"terraformLabels,omitempty" tf:"terraform_labels,omitempty"`

	// Time the ServerTlsPolicy was updated in UTC.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type ServerTLSPolicyParameters struct {

	// This field applies only for Traffic Director policies. It is must be set to false for external HTTPS load balancer policies.
	// Determines if server allows plaintext connections. If set to true, server allows plain text connections. By default, it is set to false. This setting is not exclusive of other encryption modes. For example, if allowOpen and mtlsPolicy are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to confirm compatibility.
	// Consider using it if you wish to upgrade in place your deployment to TLS while having mixed TLS and non-TLS traffic reaching port :80.
	// +kubebuilder:validation:Optional
	AllowOpen *bool `json:"allowOpen,omitempty" tf:"allow_open,omitempty"`

	// A free-text description of the resource. Max length 1024 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Set of label tags associated with the ServerTlsPolicy resource.
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The location of the server tls policy.
	// The default value is global.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// This field is required if the policy is used with external HTTPS load balancers. This field can be empty for Traffic Director.
	// Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS - mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If allowOpen and mtlsPolicy are set, server allows both plain text and mTLS connections.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	MtlsPolicy *MtlsPolicyParameters `json:"mtlsPolicy,omitempty" tf:"mtls_policy,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence of this dictates mTLS.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ServerCertificate *ServerCertificateParameters `json:"serverCertificate,omitempty" tf:"server_certificate,omitempty"`
}

// ServerTLSPolicySpec defines the desired state of ServerTLSPolicy
type ServerTLSPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServerTLSPolicyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ServerTLSPolicyInitParameters `json:"initProvider,omitempty"`
}

// ServerTLSPolicyStatus defines the observed state of ServerTLSPolicy.
type ServerTLSPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServerTLSPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ServerTLSPolicy is the Schema for the ServerTLSPolicys API. ClientTlsPolicy is a resource that specifies how a client should authenticate connections to backends of a service.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp-beta}
type ServerTLSPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServerTLSPolicySpec   `json:"spec"`
	Status            ServerTLSPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServerTLSPolicyList contains a list of ServerTLSPolicys
type ServerTLSPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServerTLSPolicy `json:"items"`
}

// Repository type metadata.
var (
	ServerTLSPolicy_Kind             = "ServerTLSPolicy"
	ServerTLSPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServerTLSPolicy_Kind}.String()
	ServerTLSPolicy_KindAPIVersion   = ServerTLSPolicy_Kind + "." + CRDGroupVersion.String()
	ServerTLSPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ServerTLSPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ServerTLSPolicy{}, &ServerTLSPolicyList{})
}
