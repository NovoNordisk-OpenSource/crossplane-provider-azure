/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AdditionalLocationObservation struct {
	GatewayRegionalURL *string `json:"gatewayRegionalUrl,omitempty" tf:"gateway_regional_url,omitempty"`

	PrivateIPAddresses []*string `json:"privateIpAddresses,omitempty" tf:"private_ip_addresses,omitempty"`

	PublicIPAddresses []*string `json:"publicIpAddresses,omitempty" tf:"public_ip_addresses,omitempty"`
}

type AdditionalLocationParameters struct {

	// The number of compute units in this region. Defaults to the capacity of the main region.
	// +kubebuilder:validation:Optional
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkConfiguration []VirtualNetworkConfigurationParameters `json:"virtualNetworkConfiguration,omitempty" tf:"virtual_network_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type CertificateObservation struct {
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`

	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type CertificateParameters struct {

	// +kubebuilder:validation:Optional
	CertificatePasswordSecretRef *v1.SecretKeySelector `json:"certificatePasswordSecretRef,omitempty" tf:"-"`

	// The Base64 Encoded PFX or Base64 Encoded X.509 Certificate.
	// +kubebuilder:validation:Required
	EncodedCertificateSecretRef v1.SecretKeySelector `json:"encodedCertificateSecretRef" tf:"-"`

	// The name of the Certificate Store where this certificate should be stored. Possible values are CertificateAuthority and Root.
	// +kubebuilder:validation:Required
	StoreName *string `json:"storeName" tf:"store_name,omitempty"`
}

type DeveloperPortalObservation struct {
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`

	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	SSLKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`

	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type DeveloperPortalParameters struct {
}

type HostNameConfigurationManagementObservation struct {
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`

	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	SSLKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`

	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type HostNameConfigurationManagementParameters struct {
}

type HostNameConfigurationObservation struct {

	// One or more developer_portal blocks as documented below.
	DeveloperPortal []DeveloperPortalObservation `json:"developerPortal,omitempty" tf:"developer_portal,omitempty"`

	// One or more management blocks as documented below.
	Management []HostNameConfigurationManagementObservation `json:"management,omitempty" tf:"management,omitempty"`

	// One or more portal blocks as documented below.
	Portal []PortalObservation `json:"portal,omitempty" tf:"portal,omitempty"`

	// One or more proxy blocks as documented below.
	Proxy []ProxyObservation `json:"proxy,omitempty" tf:"proxy,omitempty"`

	// One or more scm blocks as documented below.
	Scm []ScmObservation `json:"scm,omitempty" tf:"scm,omitempty"`
}

type HostNameConfigurationParameters struct {
}

type IdentityObservation struct {

	// The Principal ID associated with this Managed Service Identity.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type IdentityParameters struct {

	// A list of User Assigned Managed Identity IDs to be assigned to this API Management Service.
	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this API Management Service. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned .
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ManagementObservation struct {

	// +kubebuilder:validation:Optional
	AdditionalLocation []AdditionalLocationObservation `json:"additionalLocation,omitempty" tf:"additional_location,omitempty"`

	// +kubebuilder:validation:Optional
	Certificate []CertificateObservation `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// The URL for the Developer Portal associated with this API Management service.
	DeveloperPortalURL *string `json:"developerPortalUrl,omitempty" tf:"developer_portal_url,omitempty"`

	GatewayRegionalURL *string `json:"gatewayRegionalUrl,omitempty" tf:"gateway_regional_url,omitempty"`

	// The URL of the Gateway for the API Management Service.
	GatewayURL *string `json:"gatewayUrl,omitempty" tf:"gateway_url,omitempty"`

	// A hostname_configuration block as defined below.
	HostNameConfiguration []HostNameConfigurationObservation `json:"hostnameConfiguration,omitempty" tf:"hostname_configuration,omitempty"`

	// The ID of the API Management Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// The URL for the Management API associated with this API Management service.
	ManagementAPIURL *string `json:"managementApiUrl,omitempty" tf:"management_api_url,omitempty"`

	// The URL for the Publisher Portal associated with this API Management service.
	PortalURL *string `json:"portalUrl,omitempty" tf:"portal_url,omitempty"`

	PrivateIPAddresses []*string `json:"privateIpAddresses,omitempty" tf:"private_ip_addresses,omitempty"`

	PublicIPAddresses []*string `json:"publicIpAddresses,omitempty" tf:"public_ip_addresses,omitempty"`

	// The URL for the SCM  Endpoint associated with this API Management service.
	ScmURL *string `json:"scmUrl,omitempty" tf:"scm_url,omitempty"`

	// +kubebuilder:validation:Optional
	TenantAccess []TenantAccessObservation `json:"tenantAccess,omitempty" tf:"tenant_access,omitempty"`
}

type ManagementParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalLocation []AdditionalLocationParameters `json:"additionalLocation,omitempty" tf:"additional_location,omitempty"`

	// +kubebuilder:validation:Optional
	Certificate []CertificateParameters `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// Enforce a client certificate to be presented on each request to the gateway? This is only supported when SKU type is Consumption.
	// +kubebuilder:validation:Optional
	ClientCertificateEnabled *bool `json:"clientCertificateEnabled,omitempty" tf:"client_certificate_enabled,omitempty"`

	// Disable the gateway in main region? This is only supported when additional_location is set.
	// +kubebuilder:validation:Optional
	GatewayDisabled *bool `json:"gatewayDisabled,omitempty" tf:"gateway_disabled,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The version which the control plane API calls to API Management service are limited with version equal to or newer than.
	// +kubebuilder:validation:Optional
	MinAPIVersion *string `json:"minApiVersion,omitempty" tf:"min_api_version,omitempty"`

	// Email address from which the notification will be sent.
	// +kubebuilder:validation:Optional
	NotificationSenderEmail *string `json:"notificationSenderEmail,omitempty" tf:"notification_sender_email,omitempty"`

	// A policy block as defined below.
	// +kubebuilder:validation:Optional
	Policy []PolicyParameters `json:"policy,omitempty" tf:"policy,omitempty"`

	// A protocols block as defined below.
	// +kubebuilder:validation:Optional
	Protocols []ProtocolsParameters `json:"protocols,omitempty" tf:"protocols,omitempty"`

	// +kubebuilder:validation:Optional
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// Is public access to the service allowed?. Defaults to true
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The email of publisher/company.
	// +kubebuilder:validation:Required
	PublisherEmail *string `json:"publisherEmail" tf:"publisher_email,omitempty"`

	// The name of publisher/company.
	// +kubebuilder:validation:Required
	PublisherName *string `json:"publisherName" tf:"publisher_name,omitempty"`

	// The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A security block as defined below.
	// +kubebuilder:validation:Optional
	Security []SecurityParameters `json:"security,omitempty" tf:"security,omitempty"`

	// A sign_in block as defined below.
	// +kubebuilder:validation:Optional
	SignIn []SignInParameters `json:"signIn,omitempty" tf:"sign_in,omitempty"`

	// A sign_up block as defined below.
	// +kubebuilder:validation:Optional
	SignUp []SignUpParameters `json:"signUp,omitempty" tf:"sign_up,omitempty"`

	// sku_name is a string consisting of two parts separated by an underscore. The first part is the name, valid values include: Consumption, Developer, Basic, Standard and Premium. The second part is the capacity , which must be a positive integer .
	// +kubebuilder:validation:Required
	SkuName *string `json:"skuName" tf:"sku_name,omitempty"`

	// A mapping of tags assigned to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TenantAccess []TenantAccessParameters `json:"tenantAccess,omitempty" tf:"tenant_access,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkConfiguration []ManagementVirtualNetworkConfigurationParameters `json:"virtualNetworkConfiguration,omitempty" tf:"virtual_network_configuration,omitempty"`

	// The type of virtual network you want to use, valid values include: None, External, Internal.
	// +kubebuilder:validation:Optional
	VirtualNetworkType *string `json:"virtualNetworkType,omitempty" tf:"virtual_network_type,omitempty"`

	// +kubebuilder:validation:Optional
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type ManagementVirtualNetworkConfigurationObservation struct {
}

type ManagementVirtualNetworkConfigurationParameters struct {

	// The id of the subnet that will be used for the API Management.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type PolicyObservation struct {
}

type PolicyParameters struct {

	// The XML Content for this Policy.
	// +kubebuilder:validation:Optional
	XMLContent *string `json:"xmlContent,omitempty" tf:"xml_content"`

	// A link to an API Management Policy XML Document, which must be publicly available.
	// +kubebuilder:validation:Optional
	XMLLink *string `json:"xmlLink,omitempty" tf:"xml_link"`
}

type PortalObservation struct {
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`

	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	SSLKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`

	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type PortalParameters struct {
}

type ProtocolsObservation struct {
}

type ProtocolsParameters struct {

	// Should HTTP/2 be supported by the API Management Service? Defaults to false.
	// +kubebuilder:validation:Optional
	EnableHttp2 *bool `json:"enableHttp2,omitempty" tf:"enable_http2,omitempty"`
}

type ProxyObservation struct {

	// Is the certificate associated with this Hostname the Default SSL Certificate? This is used when an SNI header isn't specified by a client. Defaults to false.
	DefaultSSLBinding *bool `json:"defaultSslBinding,omitempty" tf:"default_ssl_binding,omitempty"`

	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`

	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	SSLKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`

	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type ProxyParameters struct {
}

type ScmObservation struct {
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`

	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	SSLKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`

	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type ScmParameters struct {
}

type SecurityObservation struct {
}

type SecurityParameters struct {

	// Should SSL 3.0 be enabled on the backend of the gateway? Defaults to false.
	// +kubebuilder:validation:Optional
	EnableBackendSsl30 *bool `json:"enableBackendSsl30,omitempty" tf:"enable_backend_ssl30,omitempty"`

	// Should TLS 1.0 be enabled on the backend of the gateway? Defaults to false.
	// +kubebuilder:validation:Optional
	EnableBackendTls10 *bool `json:"enableBackendTls10,omitempty" tf:"enable_backend_tls10,omitempty"`

	// Should TLS 1.1 be enabled on the backend of the gateway? Defaults to false.
	// +kubebuilder:validation:Optional
	EnableBackendTls11 *bool `json:"enableBackendTls11,omitempty" tf:"enable_backend_tls11,omitempty"`

	// Should SSL 3.0 be enabled on the frontend of the gateway? Defaults to false.
	// +kubebuilder:validation:Optional
	EnableFrontendSsl30 *bool `json:"enableFrontendSsl30,omitempty" tf:"enable_frontend_ssl30,omitempty"`

	// Should TLS 1.0 be enabled on the frontend of the gateway? Defaults to false.
	// +kubebuilder:validation:Optional
	EnableFrontendTls10 *bool `json:"enableFrontendTls10,omitempty" tf:"enable_frontend_tls10,omitempty"`

	// Should TLS 1.1 be enabled on the frontend of the gateway? Defaults to false.
	// +kubebuilder:validation:Optional
	EnableFrontendTls11 *bool `json:"enableFrontendTls11,omitempty" tf:"enable_frontend_tls11,omitempty"`

	// Should the TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA cipher be enabled? Defaults to false.
	// +kubebuilder:validation:Optional
	TLSEcdheEcdsaWithAes128CbcShaCiphersEnabled *bool `json:"tlsEcdheEcdsaWithAes128CbcShaCiphersEnabled,omitempty" tf:"tls_ecdhe_ecdsa_with_aes128_cbc_sha_ciphers_enabled,omitempty"`

	// Should the TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA cipher be enabled? Defaults to false.
	// +kubebuilder:validation:Optional
	TLSEcdheEcdsaWithAes256CbcShaCiphersEnabled *bool `json:"tlsEcdheEcdsaWithAes256CbcShaCiphersEnabled,omitempty" tf:"tls_ecdhe_ecdsa_with_aes256_cbc_sha_ciphers_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	TLSEcdheRsaWithAes128CbcShaCiphersEnabled *bool `json:"tlsEcdheRsaWithAes128CbcShaCiphersEnabled,omitempty" tf:"tls_ecdhe_rsa_with_aes128_cbc_sha_ciphers_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	TLSEcdheRsaWithAes256CbcShaCiphersEnabled *bool `json:"tlsEcdheRsaWithAes256CbcShaCiphersEnabled,omitempty" tf:"tls_ecdhe_rsa_with_aes256_cbc_sha_ciphers_enabled,omitempty"`

	// Should the TLS_RSA_WITH_AES_128_CBC_SHA256 cipher be enabled? Defaults to false.
	// +kubebuilder:validation:Optional
	TLSRsaWithAes128CbcSha256CiphersEnabled *bool `json:"tlsRsaWithAes128CbcSha256CiphersEnabled,omitempty" tf:"tls_rsa_with_aes128_cbc_sha256_ciphers_enabled,omitempty"`

	// Should the TLS_RSA_WITH_AES_128_CBC_SHA cipher be enabled? Defaults to false.
	// +kubebuilder:validation:Optional
	TLSRsaWithAes128CbcShaCiphersEnabled *bool `json:"tlsRsaWithAes128CbcShaCiphersEnabled,omitempty" tf:"tls_rsa_with_aes128_cbc_sha_ciphers_enabled,omitempty"`

	// Should the TLS_RSA_WITH_AES_128_GCM_SHA256 cipher be enabled? Defaults to false.
	// +kubebuilder:validation:Optional
	TLSRsaWithAes128GCMSha256CiphersEnabled *bool `json:"tlsRsaWithAes128GcmSha256CiphersEnabled,omitempty" tf:"tls_rsa_with_aes128_gcm_sha256_ciphers_enabled,omitempty"`

	// Should the TLS_RSA_WITH_AES_256_CBC_SHA256 cipher be enabled? Defaults to false.
	// +kubebuilder:validation:Optional
	TLSRsaWithAes256CbcSha256CiphersEnabled *bool `json:"tlsRsaWithAes256CbcSha256CiphersEnabled,omitempty" tf:"tls_rsa_with_aes256_cbc_sha256_ciphers_enabled,omitempty"`

	// Should the TLS_RSA_WITH_AES_256_CBC_SHA cipher be enabled? Defaults to false.
	// +kubebuilder:validation:Optional
	TLSRsaWithAes256CbcShaCiphersEnabled *bool `json:"tlsRsaWithAes256CbcShaCiphersEnabled,omitempty" tf:"tls_rsa_with_aes256_cbc_sha_ciphers_enabled,omitempty"`

	// Should the TLS_RSA_WITH_3DES_EDE_CBC_SHA cipher be enabled for alL TLS versions ? Defaults to false.
	// +kubebuilder:validation:Optional
	TripleDesCiphersEnabled *bool `json:"tripleDesCiphersEnabled,omitempty" tf:"triple_des_ciphers_enabled,omitempty"`
}

type SignInObservation struct {
}

type SignInParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type SignUpObservation struct {
}

type SignUpParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// A terms_of_service block as defined below.
	// +kubebuilder:validation:Required
	TermsOfService []TermsOfServiceParameters `json:"termsOfService" tf:"terms_of_service,omitempty"`
}

type TenantAccessObservation struct {
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type TenantAccessParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type TermsOfServiceObservation struct {
}

type TermsOfServiceParameters struct {

	// Should the user be asked for consent during sign up?
	// +kubebuilder:validation:Required
	ConsentRequired *bool `json:"consentRequired" tf:"consent_required,omitempty"`

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// The Terms of Service which users are required to agree to in order to sign up.
	// +kubebuilder:validation:Optional
	Text *string `json:"text,omitempty" tf:"text,omitempty"`
}

type VirtualNetworkConfigurationObservation struct {
}

type VirtualNetworkConfigurationParameters struct {

	// The id of the subnet that will be used for the API Management.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// ManagementSpec defines the desired state of Management
type ManagementSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagementParameters `json:"forProvider"`
}

// ManagementStatus defines the observed state of Management.
type ManagementStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagementObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Management is the Schema for the Managements API. Manages an API Management Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Management struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagementSpec   `json:"spec"`
	Status            ManagementStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementList contains a list of Managements
type ManagementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Management `json:"items"`
}

// Repository type metadata.
var (
	Management_Kind             = "Management"
	Management_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Management_Kind}.String()
	Management_KindAPIVersion   = Management_Kind + "." + CRDGroupVersion.String()
	Management_GroupVersionKind = CRDGroupVersion.WithKind(Management_Kind)
)

func init() {
	SchemeBuilder.Register(&Management{}, &ManagementList{})
}
