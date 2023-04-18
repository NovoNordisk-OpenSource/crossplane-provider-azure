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

type GatewayAPIObservation struct {

	// The Identifier of the API Management API within the API Management Service. Changing this forces a new API Management Gateway API to be created.
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// The Identifier for the API Management Gateway. Changing this forces a new API Management Gateway API to be created.
	GatewayID *string `json:"gatewayId,omitempty" tf:"gateway_id,omitempty"`

	// The ID of the API Management Gateway API.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GatewayAPIParameters struct {

	// The Identifier of the API Management API within the API Management Service. Changing this forces a new API Management Gateway API to be created.
	// +crossplane:generate:reference:type=API
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Reference to a API to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// Selector for a API to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// The Identifier for the API Management Gateway. Changing this forces a new API Management Gateway API to be created.
	// +crossplane:generate:reference:type=Gateway
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	GatewayID *string `json:"gatewayId,omitempty" tf:"gateway_id,omitempty"`

	// Reference to a Gateway to populate gatewayId.
	// +kubebuilder:validation:Optional
	GatewayIDRef *v1.Reference `json:"gatewayIdRef,omitempty" tf:"-"`

	// Selector for a Gateway to populate gatewayId.
	// +kubebuilder:validation:Optional
	GatewayIDSelector *v1.Selector `json:"gatewayIdSelector,omitempty" tf:"-"`
}

// GatewayAPISpec defines the desired state of GatewayAPI
type GatewayAPISpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayAPIParameters `json:"forProvider"`
}

// GatewayAPIStatus defines the observed state of GatewayAPI.
type GatewayAPIStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayAPIObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayAPI is the Schema for the GatewayAPIs API. Manages a API Management Gateway API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type GatewayAPI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewayAPISpec   `json:"spec"`
	Status            GatewayAPIStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayAPIList contains a list of GatewayAPIs
type GatewayAPIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatewayAPI `json:"items"`
}

// Repository type metadata.
var (
	GatewayAPI_Kind             = "GatewayAPI"
	GatewayAPI_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GatewayAPI_Kind}.String()
	GatewayAPI_KindAPIVersion   = GatewayAPI_Kind + "." + CRDGroupVersion.String()
	GatewayAPI_GroupVersionKind = CRDGroupVersion.WithKind(GatewayAPI_Kind)
)

func init() {
	SchemeBuilder.Register(&GatewayAPI{}, &GatewayAPIList{})
}
