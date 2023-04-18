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

type MSSQLManagedDatabaseObservation struct {

	// The Azure SQL Managed Database ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Azure SQL Managed Instance on which to create this Managed Database. Changing this forces a new resource to be created.
	ManagedInstanceID *string `json:"managedInstanceId,omitempty" tf:"managed_instance_id,omitempty"`
}

type MSSQLManagedDatabaseParameters struct {

	// The ID of the Azure SQL Managed Instance on which to create this Managed Database. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=MSSQLManagedInstance
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ManagedInstanceID *string `json:"managedInstanceId,omitempty" tf:"managed_instance_id,omitempty"`

	// Reference to a MSSQLManagedInstance to populate managedInstanceId.
	// +kubebuilder:validation:Optional
	ManagedInstanceIDRef *v1.Reference `json:"managedInstanceIdRef,omitempty" tf:"-"`

	// Selector for a MSSQLManagedInstance to populate managedInstanceId.
	// +kubebuilder:validation:Optional
	ManagedInstanceIDSelector *v1.Selector `json:"managedInstanceIdSelector,omitempty" tf:"-"`
}

// MSSQLManagedDatabaseSpec defines the desired state of MSSQLManagedDatabase
type MSSQLManagedDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MSSQLManagedDatabaseParameters `json:"forProvider"`
}

// MSSQLManagedDatabaseStatus defines the observed state of MSSQLManagedDatabase.
type MSSQLManagedDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MSSQLManagedDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLManagedDatabase is the Schema for the MSSQLManagedDatabases API. Manages an Azure SQL Azure Managed Database.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MSSQLManagedDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MSSQLManagedDatabaseSpec   `json:"spec"`
	Status            MSSQLManagedDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLManagedDatabaseList contains a list of MSSQLManagedDatabases
type MSSQLManagedDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSSQLManagedDatabase `json:"items"`
}

// Repository type metadata.
var (
	MSSQLManagedDatabase_Kind             = "MSSQLManagedDatabase"
	MSSQLManagedDatabase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MSSQLManagedDatabase_Kind}.String()
	MSSQLManagedDatabase_KindAPIVersion   = MSSQLManagedDatabase_Kind + "." + CRDGroupVersion.String()
	MSSQLManagedDatabase_GroupVersionKind = CRDGroupVersion.WithKind(MSSQLManagedDatabase_Kind)
)

func init() {
	SchemeBuilder.Register(&MSSQLManagedDatabase{}, &MSSQLManagedDatabaseList{})
}
