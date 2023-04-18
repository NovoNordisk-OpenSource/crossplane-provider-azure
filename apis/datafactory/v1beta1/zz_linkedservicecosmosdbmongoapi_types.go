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

type LinkedServiceCosmosDBMongoapiObservation struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// The name of the database.
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// The description for the Data Factory Linked Service.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Data Factory Linked Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Whether API server version is 3.2 or higher. Defaults to false.
	ServerVersionIs32OrHigher *bool `json:"serverVersionIs32OrHigher,omitempty" tf:"server_version_is_32_or_higher,omitempty"`
}

type LinkedServiceCosmosDBMongoapiParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The connection string.
	// +kubebuilder:validation:Optional
	ConnectionStringSecretRef *v1.SecretKeySelector `json:"connectionStringSecretRef,omitempty" tf:"-"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// The name of the database.
	// +kubebuilder:validation:Optional
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// The description for the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Whether API server version is 3.2 or higher. Defaults to false.
	// +kubebuilder:validation:Optional
	ServerVersionIs32OrHigher *bool `json:"serverVersionIs32OrHigher,omitempty" tf:"server_version_is_32_or_higher,omitempty"`
}

// LinkedServiceCosmosDBMongoapiSpec defines the desired state of LinkedServiceCosmosDBMongoapi
type LinkedServiceCosmosDBMongoapiSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedServiceCosmosDBMongoapiParameters `json:"forProvider"`
}

// LinkedServiceCosmosDBMongoapiStatus defines the observed state of LinkedServiceCosmosDBMongoapi.
type LinkedServiceCosmosDBMongoapiStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedServiceCosmosDBMongoapiObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceCosmosDBMongoapi is the Schema for the LinkedServiceCosmosDBMongoapis API. Manages a Linked Service (connection) between a CosmosDB and Azure Data Factory using Mongo API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LinkedServiceCosmosDBMongoapi struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinkedServiceCosmosDBMongoapiSpec   `json:"spec"`
	Status            LinkedServiceCosmosDBMongoapiStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceCosmosDBMongoapiList contains a list of LinkedServiceCosmosDBMongoapis
type LinkedServiceCosmosDBMongoapiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedServiceCosmosDBMongoapi `json:"items"`
}

// Repository type metadata.
var (
	LinkedServiceCosmosDBMongoapi_Kind             = "LinkedServiceCosmosDBMongoapi"
	LinkedServiceCosmosDBMongoapi_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedServiceCosmosDBMongoapi_Kind}.String()
	LinkedServiceCosmosDBMongoapi_KindAPIVersion   = LinkedServiceCosmosDBMongoapi_Kind + "." + CRDGroupVersion.String()
	LinkedServiceCosmosDBMongoapi_GroupVersionKind = CRDGroupVersion.WithKind(LinkedServiceCosmosDBMongoapi_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedServiceCosmosDBMongoapi{}, &LinkedServiceCosmosDBMongoapiList{})
}
