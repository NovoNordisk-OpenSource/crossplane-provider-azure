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

type MonitorDataCollectionRuleAssociationObservation struct {

	// The ID of the Data Collection Endpoint which will be associated to the target resource.
	DataCollectionEndpointID *string `json:"dataCollectionEndpointId,omitempty" tf:"data_collection_endpoint_id,omitempty"`

	// The ID of the Data Collection Rule which will be associated to the target resource.
	DataCollectionRuleID *string `json:"dataCollectionRuleId,omitempty" tf:"data_collection_rule_id,omitempty"`

	// The description of the Data Collection Rule Association.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Data Collection Rule Association.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Azure Resource which to associate to a Data Collection Rule or a Data Collection Endpoint. Changing this forces a new resource to be created.
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`
}

type MonitorDataCollectionRuleAssociationParameters struct {

	// The ID of the Data Collection Endpoint which will be associated to the target resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/insights/v1beta1.MonitorDataCollectionEndpoint
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataCollectionEndpointID *string `json:"dataCollectionEndpointId,omitempty" tf:"data_collection_endpoint_id,omitempty"`

	// Reference to a MonitorDataCollectionEndpoint in insights to populate dataCollectionEndpointId.
	// +kubebuilder:validation:Optional
	DataCollectionEndpointIDRef *v1.Reference `json:"dataCollectionEndpointIdRef,omitempty" tf:"-"`

	// Selector for a MonitorDataCollectionEndpoint in insights to populate dataCollectionEndpointId.
	// +kubebuilder:validation:Optional
	DataCollectionEndpointIDSelector *v1.Selector `json:"dataCollectionEndpointIdSelector,omitempty" tf:"-"`

	// The ID of the Data Collection Rule which will be associated to the target resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/insights/v1beta1.MonitorDataCollectionRule
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataCollectionRuleID *string `json:"dataCollectionRuleId,omitempty" tf:"data_collection_rule_id,omitempty"`

	// Reference to a MonitorDataCollectionRule in insights to populate dataCollectionRuleId.
	// +kubebuilder:validation:Optional
	DataCollectionRuleIDRef *v1.Reference `json:"dataCollectionRuleIdRef,omitempty" tf:"-"`

	// Selector for a MonitorDataCollectionRule in insights to populate dataCollectionRuleId.
	// +kubebuilder:validation:Optional
	DataCollectionRuleIDSelector *v1.Selector `json:"dataCollectionRuleIdSelector,omitempty" tf:"-"`

	// The description of the Data Collection Rule Association.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Azure Resource which to associate to a Data Collection Rule or a Data Collection Endpoint. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/compute/v1beta1.LinuxVirtualMachine
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`

	// Reference to a LinuxVirtualMachine in compute to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDRef *v1.Reference `json:"targetResourceIdRef,omitempty" tf:"-"`

	// Selector for a LinuxVirtualMachine in compute to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDSelector *v1.Selector `json:"targetResourceIdSelector,omitempty" tf:"-"`
}

// MonitorDataCollectionRuleAssociationSpec defines the desired state of MonitorDataCollectionRuleAssociation
type MonitorDataCollectionRuleAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorDataCollectionRuleAssociationParameters `json:"forProvider"`
}

// MonitorDataCollectionRuleAssociationStatus defines the observed state of MonitorDataCollectionRuleAssociation.
type MonitorDataCollectionRuleAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorDataCollectionRuleAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorDataCollectionRuleAssociation is the Schema for the MonitorDataCollectionRuleAssociations API. Manages a Data Collection Rule Association.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MonitorDataCollectionRuleAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorDataCollectionRuleAssociationSpec   `json:"spec"`
	Status            MonitorDataCollectionRuleAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorDataCollectionRuleAssociationList contains a list of MonitorDataCollectionRuleAssociations
type MonitorDataCollectionRuleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorDataCollectionRuleAssociation `json:"items"`
}

// Repository type metadata.
var (
	MonitorDataCollectionRuleAssociation_Kind             = "MonitorDataCollectionRuleAssociation"
	MonitorDataCollectionRuleAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitorDataCollectionRuleAssociation_Kind}.String()
	MonitorDataCollectionRuleAssociation_KindAPIVersion   = MonitorDataCollectionRuleAssociation_Kind + "." + CRDGroupVersion.String()
	MonitorDataCollectionRuleAssociation_GroupVersionKind = CRDGroupVersion.WithKind(MonitorDataCollectionRuleAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitorDataCollectionRuleAssociation{}, &MonitorDataCollectionRuleAssociationList{})
}
