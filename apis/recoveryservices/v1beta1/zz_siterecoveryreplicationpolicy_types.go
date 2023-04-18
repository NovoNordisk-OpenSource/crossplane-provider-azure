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

type SiteRecoveryReplicationPolicyObservation struct {

	// Specifies the frequency(in minutes) at which to create application consistent recovery points.
	ApplicationConsistentSnapshotFrequencyInMinutes *float64 `json:"applicationConsistentSnapshotFrequencyInMinutes,omitempty" tf:"application_consistent_snapshot_frequency_in_minutes,omitempty"`

	// The ID of the Site Recovery Replication Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The duration in minutes for which the recovery points need to be stored.
	RecoveryPointRetentionInMinutes *float64 `json:"recoveryPointRetentionInMinutes,omitempty" tf:"recovery_point_retention_in_minutes,omitempty"`

	// The name of the vault that should be updated. Changing this forces a new resource to be created.
	RecoveryVaultName *string `json:"recoveryVaultName,omitempty" tf:"recovery_vault_name,omitempty"`

	// Name of the resource group where the vault that should be updated is located. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
}

type SiteRecoveryReplicationPolicyParameters struct {

	// Specifies the frequency(in minutes) at which to create application consistent recovery points.
	// +kubebuilder:validation:Optional
	ApplicationConsistentSnapshotFrequencyInMinutes *float64 `json:"applicationConsistentSnapshotFrequencyInMinutes,omitempty" tf:"application_consistent_snapshot_frequency_in_minutes,omitempty"`

	// The duration in minutes for which the recovery points need to be stored.
	// +kubebuilder:validation:Optional
	RecoveryPointRetentionInMinutes *float64 `json:"recoveryPointRetentionInMinutes,omitempty" tf:"recovery_point_retention_in_minutes,omitempty"`

	// The name of the vault that should be updated. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/recoveryservices/v1beta1.Vault
	// +kubebuilder:validation:Optional
	RecoveryVaultName *string `json:"recoveryVaultName,omitempty" tf:"recovery_vault_name,omitempty"`

	// Reference to a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameRef *v1.Reference `json:"recoveryVaultNameRef,omitempty" tf:"-"`

	// Selector for a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameSelector *v1.Selector `json:"recoveryVaultNameSelector,omitempty" tf:"-"`

	// Name of the resource group where the vault that should be updated is located. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// SiteRecoveryReplicationPolicySpec defines the desired state of SiteRecoveryReplicationPolicy
type SiteRecoveryReplicationPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SiteRecoveryReplicationPolicyParameters `json:"forProvider"`
}

// SiteRecoveryReplicationPolicyStatus defines the observed state of SiteRecoveryReplicationPolicy.
type SiteRecoveryReplicationPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SiteRecoveryReplicationPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SiteRecoveryReplicationPolicy is the Schema for the SiteRecoveryReplicationPolicys API. Manages an Azure Site Recovery replication policy on Azure.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SiteRecoveryReplicationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.applicationConsistentSnapshotFrequencyInMinutes)",message="applicationConsistentSnapshotFrequencyInMinutes is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.recoveryPointRetentionInMinutes)",message="recoveryPointRetentionInMinutes is a required parameter"
	Spec   SiteRecoveryReplicationPolicySpec   `json:"spec"`
	Status SiteRecoveryReplicationPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiteRecoveryReplicationPolicyList contains a list of SiteRecoveryReplicationPolicys
type SiteRecoveryReplicationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SiteRecoveryReplicationPolicy `json:"items"`
}

// Repository type metadata.
var (
	SiteRecoveryReplicationPolicy_Kind             = "SiteRecoveryReplicationPolicy"
	SiteRecoveryReplicationPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SiteRecoveryReplicationPolicy_Kind}.String()
	SiteRecoveryReplicationPolicy_KindAPIVersion   = SiteRecoveryReplicationPolicy_Kind + "." + CRDGroupVersion.String()
	SiteRecoveryReplicationPolicy_GroupVersionKind = CRDGroupVersion.WithKind(SiteRecoveryReplicationPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&SiteRecoveryReplicationPolicy{}, &SiteRecoveryReplicationPolicyList{})
}
