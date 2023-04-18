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

type AppIntegrationAccountBatchConfigurationObservation struct {

	// The batch group name of the Logic App Integration Batch Configuration. Changing this forces a new resource to be created.
	BatchGroupName *string `json:"batchGroupName,omitempty" tf:"batch_group_name,omitempty"`

	// The ID of the Logic App Integration Account Batch Configuration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Logic App Integration Account. Changing this forces a new resource to be created.
	IntegrationAccountName *string `json:"integrationAccountName,omitempty" tf:"integration_account_name,omitempty"`

	// A JSON mapping of any Metadata for this Logic App Integration Account Batch Configuration.
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The name which should be used for this Logic App Integration Account Batch Configuration. Only Alphanumeric characters allowed. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A release_criteria block as documented below, which is used to select the criteria to meet before processing each batch.
	ReleaseCriteria []ReleaseCriteriaObservation `json:"releaseCriteria,omitempty" tf:"release_criteria,omitempty"`

	// The name of the Resource Group where the Logic App Integration Account Batch Configuration should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
}

type AppIntegrationAccountBatchConfigurationParameters struct {

	// The batch group name of the Logic App Integration Batch Configuration. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	BatchGroupName *string `json:"batchGroupName,omitempty" tf:"batch_group_name,omitempty"`

	// The name of the Logic App Integration Account. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/logic/v1beta1.AppIntegrationAccount
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	IntegrationAccountName *string `json:"integrationAccountName,omitempty" tf:"integration_account_name,omitempty"`

	// Reference to a AppIntegrationAccount in logic to populate integrationAccountName.
	// +kubebuilder:validation:Optional
	IntegrationAccountNameRef *v1.Reference `json:"integrationAccountNameRef,omitempty" tf:"-"`

	// Selector for a AppIntegrationAccount in logic to populate integrationAccountName.
	// +kubebuilder:validation:Optional
	IntegrationAccountNameSelector *v1.Selector `json:"integrationAccountNameSelector,omitempty" tf:"-"`

	// A JSON mapping of any Metadata for this Logic App Integration Account Batch Configuration.
	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The name which should be used for this Logic App Integration Account Batch Configuration. Only Alphanumeric characters allowed. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A release_criteria block as documented below, which is used to select the criteria to meet before processing each batch.
	// +kubebuilder:validation:Optional
	ReleaseCriteria []ReleaseCriteriaParameters `json:"releaseCriteria,omitempty" tf:"release_criteria,omitempty"`

	// The name of the Resource Group where the Logic App Integration Account Batch Configuration should exist. Changing this forces a new resource to be created.
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

type MonthlyObservation struct {

	// The occurrence of the week within the month.
	Week *float64 `json:"week,omitempty" tf:"week,omitempty"`

	// The day of the occurrence. Possible values are Sunday, Monday, Tuesday, Wednesday, Thursday, Friday and Saturday.
	Weekday *string `json:"weekday,omitempty" tf:"weekday,omitempty"`
}

type MonthlyParameters struct {

	// The occurrence of the week within the month.
	// +kubebuilder:validation:Required
	Week *float64 `json:"week" tf:"week,omitempty"`

	// The day of the occurrence. Possible values are Sunday, Monday, Tuesday, Wednesday, Thursday, Friday and Saturday.
	// +kubebuilder:validation:Required
	Weekday *string `json:"weekday" tf:"weekday,omitempty"`
}

type RecurrenceObservation struct {

	// The end time of the schedule, formatted as an RFC3339 string.
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// The frequency of the schedule. Possible values are Day, Hour, Minute, Month, NotSpecified, Second, Week and Year.
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// The number of frequencys between runs.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// A schedule block as documented below.
	Schedule []ScheduleObservation `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// The start time of the schedule, formatted as an RFC3339 string.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// The timezone of the start/end time.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type RecurrenceParameters struct {

	// The end time of the schedule, formatted as an RFC3339 string.
	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// The frequency of the schedule. Possible values are Day, Hour, Minute, Month, NotSpecified, Second, Week and Year.
	// +kubebuilder:validation:Required
	Frequency *string `json:"frequency" tf:"frequency,omitempty"`

	// The number of frequencys between runs.
	// +kubebuilder:validation:Required
	Interval *float64 `json:"interval" tf:"interval,omitempty"`

	// A schedule block as documented below.
	// +kubebuilder:validation:Optional
	Schedule []ScheduleParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// The start time of the schedule, formatted as an RFC3339 string.
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// The timezone of the start/end time.
	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type ReleaseCriteriaObservation struct {

	// The batch size in bytes for the Logic App Integration Batch Configuration.
	BatchSize *float64 `json:"batchSize,omitempty" tf:"batch_size,omitempty"`

	// The message count for the Logic App Integration Batch Configuration.
	MessageCount *float64 `json:"messageCount,omitempty" tf:"message_count,omitempty"`

	// A recurrence block as documented below.
	Recurrence []RecurrenceObservation `json:"recurrence,omitempty" tf:"recurrence,omitempty"`
}

type ReleaseCriteriaParameters struct {

	// The batch size in bytes for the Logic App Integration Batch Configuration.
	// +kubebuilder:validation:Optional
	BatchSize *float64 `json:"batchSize,omitempty" tf:"batch_size,omitempty"`

	// The message count for the Logic App Integration Batch Configuration.
	// +kubebuilder:validation:Optional
	MessageCount *float64 `json:"messageCount,omitempty" tf:"message_count,omitempty"`

	// A recurrence block as documented below.
	// +kubebuilder:validation:Optional
	Recurrence []RecurrenceParameters `json:"recurrence,omitempty" tf:"recurrence,omitempty"`
}

type ScheduleObservation struct {

	// A list containing a single item, which specifies the Hour interval at which this recurrence should be triggered.
	Hours []*float64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// A list containing a single item which specifies the Minute interval at which this recurrence should be triggered.
	Minutes []*float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// A list of days of the month that the job should execute on.
	MonthDays []*float64 `json:"monthDays,omitempty" tf:"month_days,omitempty"`

	// A monthly block as documented below.
	Monthly []MonthlyObservation `json:"monthly,omitempty" tf:"monthly,omitempty"`

	// A list of days of the week that the job should execute on. Possible values are Sunday, Monday, Tuesday, Wednesday, Thursday, Friday and Saturday.
	WeekDays []*string `json:"weekDays,omitempty" tf:"week_days,omitempty"`
}

type ScheduleParameters struct {

	// A list containing a single item, which specifies the Hour interval at which this recurrence should be triggered.
	// +kubebuilder:validation:Optional
	Hours []*float64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// A list containing a single item which specifies the Minute interval at which this recurrence should be triggered.
	// +kubebuilder:validation:Optional
	Minutes []*float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// A list of days of the month that the job should execute on.
	// +kubebuilder:validation:Optional
	MonthDays []*float64 `json:"monthDays,omitempty" tf:"month_days,omitempty"`

	// A monthly block as documented below.
	// +kubebuilder:validation:Optional
	Monthly []MonthlyParameters `json:"monthly,omitempty" tf:"monthly,omitempty"`

	// A list of days of the week that the job should execute on. Possible values are Sunday, Monday, Tuesday, Wednesday, Thursday, Friday and Saturday.
	// +kubebuilder:validation:Optional
	WeekDays []*string `json:"weekDays,omitempty" tf:"week_days,omitempty"`
}

// AppIntegrationAccountBatchConfigurationSpec defines the desired state of AppIntegrationAccountBatchConfiguration
type AppIntegrationAccountBatchConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppIntegrationAccountBatchConfigurationParameters `json:"forProvider"`
}

// AppIntegrationAccountBatchConfigurationStatus defines the observed state of AppIntegrationAccountBatchConfiguration.
type AppIntegrationAccountBatchConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppIntegrationAccountBatchConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppIntegrationAccountBatchConfiguration is the Schema for the AppIntegrationAccountBatchConfigurations API. Manages a Logic App Integration Account Batch Configuration.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AppIntegrationAccountBatchConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.batchGroupName)",message="batchGroupName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.releaseCriteria)",message="releaseCriteria is a required parameter"
	Spec   AppIntegrationAccountBatchConfigurationSpec   `json:"spec"`
	Status AppIntegrationAccountBatchConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppIntegrationAccountBatchConfigurationList contains a list of AppIntegrationAccountBatchConfigurations
type AppIntegrationAccountBatchConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppIntegrationAccountBatchConfiguration `json:"items"`
}

// Repository type metadata.
var (
	AppIntegrationAccountBatchConfiguration_Kind             = "AppIntegrationAccountBatchConfiguration"
	AppIntegrationAccountBatchConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppIntegrationAccountBatchConfiguration_Kind}.String()
	AppIntegrationAccountBatchConfiguration_KindAPIVersion   = AppIntegrationAccountBatchConfiguration_Kind + "." + CRDGroupVersion.String()
	AppIntegrationAccountBatchConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(AppIntegrationAccountBatchConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&AppIntegrationAccountBatchConfiguration{}, &AppIntegrationAccountBatchConfigurationList{})
}
