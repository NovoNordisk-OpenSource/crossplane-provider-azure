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

type OutputMSSQLObservation struct {

	// The authentication mode for the Stream Output. Possible values are Msi and ConnectionString. Defaults to ConnectionString.
	AuthenticationMode *string `json:"authenticationMode,omitempty" tf:"authentication_mode,omitempty"`

	// The MS SQL database name where the reference table exists. Changing this forces a new resource to be created.
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// The ID of the Stream Analytics Output Microsoft SQL Server Database.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The max batch count to write to the SQL Database. Defaults to 10000. Possible values are between 1 and 1073741824.
	MaxBatchCount *float64 `json:"maxBatchCount,omitempty" tf:"max_batch_count,omitempty"`

	// The max writer count for the SQL Database. Defaults to 1. Possible values are 0 which bases the writer count on the query partition and 1 which corresponds to a single writer.
	MaxWriterCount *float64 `json:"maxWriterCount,omitempty" tf:"max_writer_count,omitempty"`

	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The SQL server url. Changing this forces a new resource to be created.
	Server *string `json:"server,omitempty" tf:"server,omitempty"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName,omitempty" tf:"stream_analytics_job_name,omitempty"`

	// Table in the database that the output points to. Changing this forces a new resource to be created.
	Table *string `json:"table,omitempty" tf:"table,omitempty"`

	// Username used to login to the Microsoft SQL Server. Changing this forces a new resource to be created. Required if authentication_mode is ConnectionString.
	User *string `json:"user,omitempty" tf:"user,omitempty"`
}

type OutputMSSQLParameters struct {

	// The authentication mode for the Stream Output. Possible values are Msi and ConnectionString. Defaults to ConnectionString.
	// +kubebuilder:validation:Optional
	AuthenticationMode *string `json:"authenticationMode,omitempty" tf:"authentication_mode,omitempty"`

	// The MS SQL database name where the reference table exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// The max batch count to write to the SQL Database. Defaults to 10000. Possible values are between 1 and 1073741824.
	// +kubebuilder:validation:Optional
	MaxBatchCount *float64 `json:"maxBatchCount,omitempty" tf:"max_batch_count,omitempty"`

	// The max writer count for the SQL Database. Defaults to 1. Possible values are 0 which bases the writer count on the query partition and 1 which corresponds to a single writer.
	// +kubebuilder:validation:Optional
	MaxWriterCount *float64 `json:"maxWriterCount,omitempty" tf:"max_writer_count,omitempty"`

	// The name of the Stream Output. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Password used together with username, to login to the Microsoft SQL Server. Required if authentication_mode is ConnectionString.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The SQL server url. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/sql/v1beta1.MSSQLServer
	// +kubebuilder:validation:Optional
	Server *string `json:"server,omitempty" tf:"server,omitempty"`

	// Reference to a MSSQLServer in sql to populate server.
	// +kubebuilder:validation:Optional
	ServerRef *v1.Reference `json:"serverRef,omitempty" tf:"-"`

	// Selector for a MSSQLServer in sql to populate server.
	// +kubebuilder:validation:Optional
	ServerSelector *v1.Selector `json:"serverSelector,omitempty" tf:"-"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Job
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName,omitempty" tf:"stream_analytics_job_name,omitempty"`

	// Reference to a Job to populate streamAnalyticsJobName.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobNameRef *v1.Reference `json:"streamAnalyticsJobNameRef,omitempty" tf:"-"`

	// Selector for a Job to populate streamAnalyticsJobName.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobNameSelector *v1.Selector `json:"streamAnalyticsJobNameSelector,omitempty" tf:"-"`

	// Table in the database that the output points to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Table
	// +kubebuilder:validation:Optional
	Table *string `json:"table,omitempty" tf:"table,omitempty"`

	// Reference to a Table in storage to populate table.
	// +kubebuilder:validation:Optional
	TableRef *v1.Reference `json:"tableRef,omitempty" tf:"-"`

	// Selector for a Table in storage to populate table.
	// +kubebuilder:validation:Optional
	TableSelector *v1.Selector `json:"tableSelector,omitempty" tf:"-"`

	// Username used to login to the Microsoft SQL Server. Changing this forces a new resource to be created. Required if authentication_mode is ConnectionString.
	// +kubebuilder:validation:Optional
	User *string `json:"user,omitempty" tf:"user,omitempty"`
}

// OutputMSSQLSpec defines the desired state of OutputMSSQL
type OutputMSSQLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OutputMSSQLParameters `json:"forProvider"`
}

// OutputMSSQLStatus defines the observed state of OutputMSSQL.
type OutputMSSQLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OutputMSSQLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OutputMSSQL is the Schema for the OutputMSSQLs API. Manages a Stream Analytics Output to Microsoft SQL Server Database.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type OutputMSSQL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.database)",message="database is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   OutputMSSQLSpec   `json:"spec"`
	Status OutputMSSQLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OutputMSSQLList contains a list of OutputMSSQLs
type OutputMSSQLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OutputMSSQL `json:"items"`
}

// Repository type metadata.
var (
	OutputMSSQL_Kind             = "OutputMSSQL"
	OutputMSSQL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OutputMSSQL_Kind}.String()
	OutputMSSQL_KindAPIVersion   = OutputMSSQL_Kind + "." + CRDGroupVersion.String()
	OutputMSSQL_GroupVersionKind = CRDGroupVersion.WithKind(OutputMSSQL_Kind)
)

func init() {
	SchemeBuilder.Register(&OutputMSSQL{}, &OutputMSSQLList{})
}
