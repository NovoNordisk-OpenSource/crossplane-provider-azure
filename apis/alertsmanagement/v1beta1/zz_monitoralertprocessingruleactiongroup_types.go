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

type AlertRuleNameObservation struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type AlertRuleNameParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type DailyObservation struct {

	// Specifies the recurrence end time (H:M:S).
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// Specifies the recurrence start time (H:M:S).
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type DailyParameters struct {

	// Specifies the recurrence end time (H:M:S).
	// +kubebuilder:validation:Required
	EndTime *string `json:"endTime" tf:"end_time,omitempty"`

	// Specifies the recurrence start time (H:M:S).
	// +kubebuilder:validation:Required
	StartTime *string `json:"startTime" tf:"start_time,omitempty"`
}

type MonitorAlertProcessingRuleActionGroupConditionAlertContextObservation struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleActionGroupConditionAlertContextParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleActionGroupConditionAlertRuleIDObservation struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleActionGroupConditionAlertRuleIDParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleActionGroupConditionDescriptionObservation struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleActionGroupConditionDescriptionParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleActionGroupConditionMonitorServiceObservation struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleActionGroupConditionMonitorServiceParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleActionGroupConditionObservation struct {

	// A alert_context block as defined above.
	AlertContext []MonitorAlertProcessingRuleActionGroupConditionAlertContextObservation `json:"alertContext,omitempty" tf:"alert_context,omitempty"`

	// A alert_rule_id block as defined above.
	AlertRuleID []MonitorAlertProcessingRuleActionGroupConditionAlertRuleIDObservation `json:"alertRuleId,omitempty" tf:"alert_rule_id,omitempty"`

	// A alert_rule_name block as defined above.
	AlertRuleName []AlertRuleNameObservation `json:"alertRuleName,omitempty" tf:"alert_rule_name,omitempty"`

	// A description block as defined below.
	Description []MonitorAlertProcessingRuleActionGroupConditionDescriptionObservation `json:"description,omitempty" tf:"description,omitempty"`

	// A monitor_condition block as defined below.
	MonitorCondition []MonitorConditionObservation `json:"monitorCondition,omitempty" tf:"monitor_condition,omitempty"`

	// A monitor_service block as defined below.
	MonitorService []MonitorAlertProcessingRuleActionGroupConditionMonitorServiceObservation `json:"monitorService,omitempty" tf:"monitor_service,omitempty"`

	// A severity block as defined below.
	Severity []MonitorAlertProcessingRuleActionGroupConditionSeverityObservation `json:"severity,omitempty" tf:"severity,omitempty"`

	// A signal_type block as defined below.
	SignalType []SignalTypeObservation `json:"signalType,omitempty" tf:"signal_type,omitempty"`

	// A target_resource block as defined below.
	TargetResource []TargetResourceObservation `json:"targetResource,omitempty" tf:"target_resource,omitempty"`

	// A target_resource_group block as defined below.
	TargetResourceGroup []TargetResourceGroupObservation `json:"targetResourceGroup,omitempty" tf:"target_resource_group,omitempty"`

	// A target_resource_type block as defined below.
	TargetResourceType []MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeObservation `json:"targetResourceType,omitempty" tf:"target_resource_type,omitempty"`
}

type MonitorAlertProcessingRuleActionGroupConditionParameters struct {

	// A alert_context block as defined above.
	// +kubebuilder:validation:Optional
	AlertContext []MonitorAlertProcessingRuleActionGroupConditionAlertContextParameters `json:"alertContext,omitempty" tf:"alert_context,omitempty"`

	// A alert_rule_id block as defined above.
	// +kubebuilder:validation:Optional
	AlertRuleID []MonitorAlertProcessingRuleActionGroupConditionAlertRuleIDParameters `json:"alertRuleId,omitempty" tf:"alert_rule_id,omitempty"`

	// A alert_rule_name block as defined above.
	// +kubebuilder:validation:Optional
	AlertRuleName []AlertRuleNameParameters `json:"alertRuleName,omitempty" tf:"alert_rule_name,omitempty"`

	// A description block as defined below.
	// +kubebuilder:validation:Optional
	Description []MonitorAlertProcessingRuleActionGroupConditionDescriptionParameters `json:"description,omitempty" tf:"description,omitempty"`

	// A monitor_condition block as defined below.
	// +kubebuilder:validation:Optional
	MonitorCondition []MonitorConditionParameters `json:"monitorCondition,omitempty" tf:"monitor_condition,omitempty"`

	// A monitor_service block as defined below.
	// +kubebuilder:validation:Optional
	MonitorService []MonitorAlertProcessingRuleActionGroupConditionMonitorServiceParameters `json:"monitorService,omitempty" tf:"monitor_service,omitempty"`

	// A severity block as defined below.
	// +kubebuilder:validation:Optional
	Severity []MonitorAlertProcessingRuleActionGroupConditionSeverityParameters `json:"severity,omitempty" tf:"severity,omitempty"`

	// A signal_type block as defined below.
	// +kubebuilder:validation:Optional
	SignalType []SignalTypeParameters `json:"signalType,omitempty" tf:"signal_type,omitempty"`

	// A target_resource block as defined below.
	// +kubebuilder:validation:Optional
	TargetResource []TargetResourceParameters `json:"targetResource,omitempty" tf:"target_resource,omitempty"`

	// A target_resource_group block as defined below.
	// +kubebuilder:validation:Optional
	TargetResourceGroup []TargetResourceGroupParameters `json:"targetResourceGroup,omitempty" tf:"target_resource_group,omitempty"`

	// A target_resource_type block as defined below.
	// +kubebuilder:validation:Optional
	TargetResourceType []MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeParameters `json:"targetResourceType,omitempty" tf:"target_resource_type,omitempty"`
}

type MonitorAlertProcessingRuleActionGroupConditionSeverityObservation struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleActionGroupConditionSeverityParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeObservation struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleActionGroupConditionTargetResourceTypeParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleActionGroupObservation struct {

	// Specifies a list of Action Group IDs.
	AddActionGroupIds []*string `json:"addActionGroupIds,omitempty" tf:"add_action_group_ids,omitempty"`

	// A condition block as defined below.
	Condition []MonitorAlertProcessingRuleActionGroupConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	// Specifies a description for the Alert Processing Rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Should the Alert Processing Rule be enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The ID of the Alert Processing Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Resource Group where the Alert Processing Rule should exist. Changing this forces a new Alert Processing Rule to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A schedule block as defined below.
	Schedule []MonitorAlertProcessingRuleActionGroupScheduleObservation `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// A list of resource IDs which will be the target of alert processing rule.
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// A mapping of tags which should be assigned to the Alert Processing Rule.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MonitorAlertProcessingRuleActionGroupParameters struct {

	// Specifies a list of Action Group IDs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/insights/v1beta1.MonitorActionGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	AddActionGroupIds []*string `json:"addActionGroupIds,omitempty" tf:"add_action_group_ids,omitempty"`

	// References to MonitorActionGroup in insights to populate addActionGroupIds.
	// +kubebuilder:validation:Optional
	AddActionGroupIdsRefs []v1.Reference `json:"addActionGroupIdsRefs,omitempty" tf:"-"`

	// Selector for a list of MonitorActionGroup in insights to populate addActionGroupIds.
	// +kubebuilder:validation:Optional
	AddActionGroupIdsSelector *v1.Selector `json:"addActionGroupIdsSelector,omitempty" tf:"-"`

	// A condition block as defined below.
	// +kubebuilder:validation:Optional
	Condition []MonitorAlertProcessingRuleActionGroupConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// Specifies a description for the Alert Processing Rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Should the Alert Processing Rule be enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The name of the Resource Group where the Alert Processing Rule should exist. Changing this forces a new Alert Processing Rule to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A schedule block as defined below.
	// +kubebuilder:validation:Optional
	Schedule []MonitorAlertProcessingRuleActionGroupScheduleParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// A list of resource IDs which will be the target of alert processing rule.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// References to ResourceGroup in azure to populate scopes.
	// +kubebuilder:validation:Optional
	ScopesRefs []v1.Reference `json:"scopesRefs,omitempty" tf:"-"`

	// Selector for a list of ResourceGroup in azure to populate scopes.
	// +kubebuilder:validation:Optional
	ScopesSelector *v1.Selector `json:"scopesSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Alert Processing Rule.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MonitorAlertProcessingRuleActionGroupScheduleObservation struct {

	// Specifies the Alert Processing Rule effective start time (Y-m-d'T'H:M:S).
	EffectiveFrom *string `json:"effectiveFrom,omitempty" tf:"effective_from,omitempty"`

	// Specifies the Alert Processing Rule effective end time (Y-m-d'T'H:M:S).
	EffectiveUntil *string `json:"effectiveUntil,omitempty" tf:"effective_until,omitempty"`

	// A recurrence block as defined above.
	Recurrence []RecurrenceObservation `json:"recurrence,omitempty" tf:"recurrence,omitempty"`

	// The time zone (e.g. Pacific Standard time, Eastern Standard Time). Defaults to UTC. possible values are defined here.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type MonitorAlertProcessingRuleActionGroupScheduleParameters struct {

	// Specifies the Alert Processing Rule effective start time (Y-m-d'T'H:M:S).
	// +kubebuilder:validation:Optional
	EffectiveFrom *string `json:"effectiveFrom,omitempty" tf:"effective_from,omitempty"`

	// Specifies the Alert Processing Rule effective end time (Y-m-d'T'H:M:S).
	// +kubebuilder:validation:Optional
	EffectiveUntil *string `json:"effectiveUntil,omitempty" tf:"effective_until,omitempty"`

	// A recurrence block as defined above.
	// +kubebuilder:validation:Optional
	Recurrence []RecurrenceParameters `json:"recurrence,omitempty" tf:"recurrence,omitempty"`

	// The time zone (e.g. Pacific Standard time, Eastern Standard Time). Defaults to UTC. possible values are defined here.
	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type MonitorConditionObservation struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MonitorConditionParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type MonthlyObservation struct {

	// Specifies a list of dayOfMonth to recurrence. Possible values are integers between 1 - 31.
	DaysOfMonth []*float64 `json:"daysOfMonth,omitempty" tf:"days_of_month,omitempty"`

	// Specifies the recurrence end time (H:M:S).
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// Specifies the recurrence start time (H:M:S).
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type MonthlyParameters struct {

	// Specifies a list of dayOfMonth to recurrence. Possible values are integers between 1 - 31.
	// +kubebuilder:validation:Required
	DaysOfMonth []*float64 `json:"daysOfMonth" tf:"days_of_month,omitempty"`

	// Specifies the recurrence end time (H:M:S).
	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// Specifies the recurrence start time (H:M:S).
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type RecurrenceObservation struct {

	// One or more daily blocks as defined above.
	Daily []DailyObservation `json:"daily,omitempty" tf:"daily,omitempty"`

	// One or more monthly blocks as defined above.
	Monthly []MonthlyObservation `json:"monthly,omitempty" tf:"monthly,omitempty"`

	// One or more weekly blocks as defined below.
	Weekly []WeeklyObservation `json:"weekly,omitempty" tf:"weekly,omitempty"`
}

type RecurrenceParameters struct {

	// One or more daily blocks as defined above.
	// +kubebuilder:validation:Optional
	Daily []DailyParameters `json:"daily,omitempty" tf:"daily,omitempty"`

	// One or more monthly blocks as defined above.
	// +kubebuilder:validation:Optional
	Monthly []MonthlyParameters `json:"monthly,omitempty" tf:"monthly,omitempty"`

	// One or more weekly blocks as defined below.
	// +kubebuilder:validation:Optional
	Weekly []WeeklyParameters `json:"weekly,omitempty" tf:"weekly,omitempty"`
}

type SignalTypeObservation struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type SignalTypeParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type TargetResourceGroupObservation struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type TargetResourceGroupParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type TargetResourceObservation struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type TargetResourceParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type WeeklyObservation struct {

	// Specifies a list of dayOfWeek to recurrence. Possible values are Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, and Saturday.
	DaysOfWeek []*string `json:"daysOfWeek,omitempty" tf:"days_of_week,omitempty"`

	// Specifies the recurrence end time (H:M:S).
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// Specifies the recurrence start time (H:M:S).
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type WeeklyParameters struct {

	// Specifies a list of dayOfWeek to recurrence. Possible values are Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, and Saturday.
	// +kubebuilder:validation:Required
	DaysOfWeek []*string `json:"daysOfWeek" tf:"days_of_week,omitempty"`

	// Specifies the recurrence end time (H:M:S).
	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// Specifies the recurrence start time (H:M:S).
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

// MonitorAlertProcessingRuleActionGroupSpec defines the desired state of MonitorAlertProcessingRuleActionGroup
type MonitorAlertProcessingRuleActionGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorAlertProcessingRuleActionGroupParameters `json:"forProvider"`
}

// MonitorAlertProcessingRuleActionGroupStatus defines the observed state of MonitorAlertProcessingRuleActionGroup.
type MonitorAlertProcessingRuleActionGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorAlertProcessingRuleActionGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorAlertProcessingRuleActionGroup is the Schema for the MonitorAlertProcessingRuleActionGroups API. Manages an Alert Processing Rule which apply action group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MonitorAlertProcessingRuleActionGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorAlertProcessingRuleActionGroupSpec   `json:"spec"`
	Status            MonitorAlertProcessingRuleActionGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorAlertProcessingRuleActionGroupList contains a list of MonitorAlertProcessingRuleActionGroups
type MonitorAlertProcessingRuleActionGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorAlertProcessingRuleActionGroup `json:"items"`
}

// Repository type metadata.
var (
	MonitorAlertProcessingRuleActionGroup_Kind             = "MonitorAlertProcessingRuleActionGroup"
	MonitorAlertProcessingRuleActionGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitorAlertProcessingRuleActionGroup_Kind}.String()
	MonitorAlertProcessingRuleActionGroup_KindAPIVersion   = MonitorAlertProcessingRuleActionGroup_Kind + "." + CRDGroupVersion.String()
	MonitorAlertProcessingRuleActionGroup_GroupVersionKind = CRDGroupVersion.WithKind(MonitorAlertProcessingRuleActionGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitorAlertProcessingRuleActionGroup{}, &MonitorAlertProcessingRuleActionGroupList{})
}
