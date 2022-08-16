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

type FunctionJavascriptUdaObservation struct {

	// The ID of the Stream Analytics JavaScript UDA Function.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FunctionJavascriptUdaParameters struct {

	// One or more input blocks as defined below.
	// +kubebuilder:validation:Required
	Input []InputParameters `json:"input" tf:"input,omitempty"`

	// An output block as defined below.
	// +kubebuilder:validation:Required
	Output []OutputParameters `json:"output" tf:"output,omitempty"`

	// The JavaScript of this UDA Function.
	// +kubebuilder:validation:Required
	Script *string `json:"script" tf:"script,omitempty"`

	// The resource ID of the Stream Analytics Job where this Function should be created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	StreamAnalyticsJobID *string `json:"streamAnalyticsJobId" tf:"stream_analytics_job_id,omitempty"`
}

type InputObservation struct {
}

type InputParameters struct {

	// Is this input parameter a configuration parameter? Defaults to false.
	// +kubebuilder:validation:Optional
	ConfigurationParameter *bool `json:"configurationParameter,omitempty" tf:"configuration_parameter,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type OutputObservation struct {
}

type OutputParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// FunctionJavascriptUdaSpec defines the desired state of FunctionJavascriptUda
type FunctionJavascriptUdaSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FunctionJavascriptUdaParameters `json:"forProvider"`
}

// FunctionJavascriptUdaStatus defines the observed state of FunctionJavascriptUda.
type FunctionJavascriptUdaStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FunctionJavascriptUdaObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionJavascriptUda is the Schema for the FunctionJavascriptUdas API. Manages a JavaScript UDA Function within a Stream Analytics Streaming Job.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FunctionJavascriptUda struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionJavascriptUdaSpec   `json:"spec"`
	Status            FunctionJavascriptUdaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionJavascriptUdaList contains a list of FunctionJavascriptUdas
type FunctionJavascriptUdaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FunctionJavascriptUda `json:"items"`
}

// Repository type metadata.
var (
	FunctionJavascriptUda_Kind             = "FunctionJavascriptUda"
	FunctionJavascriptUda_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FunctionJavascriptUda_Kind}.String()
	FunctionJavascriptUda_KindAPIVersion   = FunctionJavascriptUda_Kind + "." + CRDGroupVersion.String()
	FunctionJavascriptUda_GroupVersionKind = CRDGroupVersion.WithKind(FunctionJavascriptUda_Kind)
)

func init() {
	SchemeBuilder.Register(&FunctionJavascriptUda{}, &FunctionJavascriptUdaList{})
}
