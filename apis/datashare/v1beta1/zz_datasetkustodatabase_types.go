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

type DataSetKustoDatabaseObservation struct {

	// The name of the Data Share Dataset.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The resource ID of the Data Share Kusto Database Dataset.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The location of the Kusto Cluster.
	KustoClusterLocation *string `json:"kustoClusterLocation,omitempty" tf:"kusto_cluster_location,omitempty"`

	// The resource ID of the Kusto Cluster Database to be shared with the receiver. Changing this forces a new Data Share Kusto Database Dataset to be created.
	KustoDatabaseID *string `json:"kustoDatabaseId,omitempty" tf:"kusto_database_id,omitempty"`

	// The resource ID of the Data Share where this Data Share Kusto Database Dataset should be created. Changing this forces a new Data Share Kusto Database Dataset to be created.
	ShareID *string `json:"shareId,omitempty" tf:"share_id,omitempty"`
}

type DataSetKustoDatabaseParameters struct {

	// The resource ID of the Kusto Cluster Database to be shared with the receiver. Changing this forces a new Data Share Kusto Database Dataset to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/kusto/v1beta1.Database
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KustoDatabaseID *string `json:"kustoDatabaseId,omitempty" tf:"kusto_database_id,omitempty"`

	// Reference to a Database in kusto to populate kustoDatabaseId.
	// +kubebuilder:validation:Optional
	KustoDatabaseIDRef *v1.Reference `json:"kustoDatabaseIdRef,omitempty" tf:"-"`

	// Selector for a Database in kusto to populate kustoDatabaseId.
	// +kubebuilder:validation:Optional
	KustoDatabaseIDSelector *v1.Selector `json:"kustoDatabaseIdSelector,omitempty" tf:"-"`

	// The resource ID of the Data Share where this Data Share Kusto Database Dataset should be created. Changing this forces a new Data Share Kusto Database Dataset to be created.
	// +crossplane:generate:reference:type=DataShare
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ShareID *string `json:"shareId,omitempty" tf:"share_id,omitempty"`

	// Reference to a DataShare to populate shareId.
	// +kubebuilder:validation:Optional
	ShareIDRef *v1.Reference `json:"shareIdRef,omitempty" tf:"-"`

	// Selector for a DataShare to populate shareId.
	// +kubebuilder:validation:Optional
	ShareIDSelector *v1.Selector `json:"shareIdSelector,omitempty" tf:"-"`
}

// DataSetKustoDatabaseSpec defines the desired state of DataSetKustoDatabase
type DataSetKustoDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataSetKustoDatabaseParameters `json:"forProvider"`
}

// DataSetKustoDatabaseStatus defines the observed state of DataSetKustoDatabase.
type DataSetKustoDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataSetKustoDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetKustoDatabase is the Schema for the DataSetKustoDatabases API. Manages a Data Share Kusto Database Dataset.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataSetKustoDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataSetKustoDatabaseSpec   `json:"spec"`
	Status            DataSetKustoDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetKustoDatabaseList contains a list of DataSetKustoDatabases
type DataSetKustoDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataSetKustoDatabase `json:"items"`
}

// Repository type metadata.
var (
	DataSetKustoDatabase_Kind             = "DataSetKustoDatabase"
	DataSetKustoDatabase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataSetKustoDatabase_Kind}.String()
	DataSetKustoDatabase_KindAPIVersion   = DataSetKustoDatabase_Kind + "." + CRDGroupVersion.String()
	DataSetKustoDatabase_GroupVersionKind = CRDGroupVersion.WithKind(DataSetKustoDatabase_Kind)
)

func init() {
	SchemeBuilder.Register(&DataSetKustoDatabase{}, &DataSetKustoDatabaseList{})
}
