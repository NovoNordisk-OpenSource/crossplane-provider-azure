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

type APIDiagnosticObservation struct {

	// The ID (name) of the Diagnostics Logger.
	APIManagementLoggerID *string `json:"apiManagementLoggerId,omitempty" tf:"api_management_logger_id,omitempty"`

	// The name of the API Management Service instance. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// The name of the API on which to configure the Diagnostics Logs. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	APIName *string `json:"apiName,omitempty" tf:"api_name,omitempty"`

	// Always log errors. Send telemetry if there is an erroneous condition, regardless of sampling settings.
	AlwaysLogErrors *bool `json:"alwaysLogErrors,omitempty" tf:"always_log_errors,omitempty"`

	// A backend_request block as defined below.
	BackendRequest []BackendRequestObservation `json:"backendRequest,omitempty" tf:"backend_request,omitempty"`

	// A backend_response block as defined below.
	BackendResponse []BackendResponseObservation `json:"backendResponse,omitempty" tf:"backend_response,omitempty"`

	// A frontend_request block as defined below.
	FrontendRequest []FrontendRequestObservation `json:"frontendRequest,omitempty" tf:"frontend_request,omitempty"`

	// A frontend_response block as defined below.
	FrontendResponse []FrontendResponseObservation `json:"frontendResponse,omitempty" tf:"frontend_response,omitempty"`

	// The HTTP Correlation Protocol to use. Possible values are None, Legacy or W3C.
	HTTPCorrelationProtocol *string `json:"httpCorrelationProtocol,omitempty" tf:"http_correlation_protocol,omitempty"`

	// The ID of the API Management Service API Diagnostics Logs.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Log client IP address.
	LogClientIP *bool `json:"logClientIp,omitempty" tf:"log_client_ip,omitempty"`

	// The format of the Operation Name for Application Insights telemetries. Possible values are Name, and Url. Defaults to Name.
	OperationNameFormat *string `json:"operationNameFormat,omitempty" tf:"operation_name_format,omitempty"`

	// The name of the Resource Group where the API Management Service API Diagnostics Logs should exist. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Sampling (%). For high traffic APIs, please read this documentation to understand performance implications and log sampling. Valid values are between 0.0 and 100.0.
	SamplingPercentage *float64 `json:"samplingPercentage,omitempty" tf:"sampling_percentage,omitempty"`

	// Logging verbosity. Possible values are verbose, information or error.
	Verbosity *string `json:"verbosity,omitempty" tf:"verbosity,omitempty"`
}

type APIDiagnosticParameters struct {

	// The ID (name) of the Diagnostics Logger.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Logger
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	APIManagementLoggerID *string `json:"apiManagementLoggerId,omitempty" tf:"api_management_logger_id,omitempty"`

	// Reference to a Logger in apimanagement to populate apiManagementLoggerId.
	// +kubebuilder:validation:Optional
	APIManagementLoggerIDRef *v1.Reference `json:"apiManagementLoggerIdRef,omitempty" tf:"-"`

	// Selector for a Logger in apimanagement to populate apiManagementLoggerId.
	// +kubebuilder:validation:Optional
	APIManagementLoggerIDSelector *v1.Selector `json:"apiManagementLoggerIdSelector,omitempty" tf:"-"`

	// The name of the API Management Service instance. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Management
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// The name of the API on which to configure the Diagnostics Logs. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.API
	// +kubebuilder:validation:Optional
	APIName *string `json:"apiName,omitempty" tf:"api_name,omitempty"`

	// Reference to a API in apimanagement to populate apiName.
	// +kubebuilder:validation:Optional
	APINameRef *v1.Reference `json:"apiNameRef,omitempty" tf:"-"`

	// Selector for a API in apimanagement to populate apiName.
	// +kubebuilder:validation:Optional
	APINameSelector *v1.Selector `json:"apiNameSelector,omitempty" tf:"-"`

	// Always log errors. Send telemetry if there is an erroneous condition, regardless of sampling settings.
	// +kubebuilder:validation:Optional
	AlwaysLogErrors *bool `json:"alwaysLogErrors,omitempty" tf:"always_log_errors,omitempty"`

	// A backend_request block as defined below.
	// +kubebuilder:validation:Optional
	BackendRequest []BackendRequestParameters `json:"backendRequest,omitempty" tf:"backend_request,omitempty"`

	// A backend_response block as defined below.
	// +kubebuilder:validation:Optional
	BackendResponse []BackendResponseParameters `json:"backendResponse,omitempty" tf:"backend_response,omitempty"`

	// A frontend_request block as defined below.
	// +kubebuilder:validation:Optional
	FrontendRequest []FrontendRequestParameters `json:"frontendRequest,omitempty" tf:"frontend_request,omitempty"`

	// A frontend_response block as defined below.
	// +kubebuilder:validation:Optional
	FrontendResponse []FrontendResponseParameters `json:"frontendResponse,omitempty" tf:"frontend_response,omitempty"`

	// The HTTP Correlation Protocol to use. Possible values are None, Legacy or W3C.
	// +kubebuilder:validation:Optional
	HTTPCorrelationProtocol *string `json:"httpCorrelationProtocol,omitempty" tf:"http_correlation_protocol,omitempty"`

	// Log client IP address.
	// +kubebuilder:validation:Optional
	LogClientIP *bool `json:"logClientIp,omitempty" tf:"log_client_ip,omitempty"`

	// The format of the Operation Name for Application Insights telemetries. Possible values are Name, and Url. Defaults to Name.
	// +kubebuilder:validation:Optional
	OperationNameFormat *string `json:"operationNameFormat,omitempty" tf:"operation_name_format,omitempty"`

	// The name of the Resource Group where the API Management Service API Diagnostics Logs should exist. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Sampling (%). For high traffic APIs, please read this documentation to understand performance implications and log sampling. Valid values are between 0.0 and 100.0.
	// +kubebuilder:validation:Optional
	SamplingPercentage *float64 `json:"samplingPercentage,omitempty" tf:"sampling_percentage,omitempty"`

	// Logging verbosity. Possible values are verbose, information or error.
	// +kubebuilder:validation:Optional
	Verbosity *string `json:"verbosity,omitempty" tf:"verbosity,omitempty"`
}

type BackendRequestObservation struct {

	// Number of payload bytes to log (up to 8192).
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// A data_masking block as defined below.
	DataMasking []DataMaskingObservation `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies a list of headers to log.
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type BackendRequestParameters struct {

	// Number of payload bytes to log (up to 8192).
	// +kubebuilder:validation:Optional
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// A data_masking block as defined below.
	// +kubebuilder:validation:Optional
	DataMasking []DataMaskingParameters `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies a list of headers to log.
	// +kubebuilder:validation:Optional
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type BackendResponseDataMaskingObservation struct {

	// A headers block as defined below.
	Headers []DataMaskingHeadersObservation `json:"headers,omitempty" tf:"headers,omitempty"`

	// A query_params block as defined below.
	QueryParams []DataMaskingQueryParamsObservation `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type BackendResponseDataMaskingParameters struct {

	// A headers block as defined below.
	// +kubebuilder:validation:Optional
	Headers []DataMaskingHeadersParameters `json:"headers,omitempty" tf:"headers,omitempty"`

	// A query_params block as defined below.
	// +kubebuilder:validation:Optional
	QueryParams []DataMaskingQueryParamsParameters `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type BackendResponseObservation struct {

	// Number of payload bytes to log (up to 8192).
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// A data_masking block as defined below.
	DataMasking []BackendResponseDataMaskingObservation `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies a list of headers to log.
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type BackendResponseParameters struct {

	// Number of payload bytes to log (up to 8192).
	// +kubebuilder:validation:Optional
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// A data_masking block as defined below.
	// +kubebuilder:validation:Optional
	DataMasking []BackendResponseDataMaskingParameters `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies a list of headers to log.
	// +kubebuilder:validation:Optional
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type DataMaskingHeadersObservation struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DataMaskingHeadersParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type DataMaskingObservation struct {

	// A headers block as defined below.
	Headers []HeadersObservation `json:"headers,omitempty" tf:"headers,omitempty"`

	// A query_params block as defined below.
	QueryParams []QueryParamsObservation `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type DataMaskingParameters struct {

	// A headers block as defined below.
	// +kubebuilder:validation:Optional
	Headers []HeadersParameters `json:"headers,omitempty" tf:"headers,omitempty"`

	// A query_params block as defined below.
	// +kubebuilder:validation:Optional
	QueryParams []QueryParamsParameters `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type DataMaskingQueryParamsObservation struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DataMaskingQueryParamsParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type FrontendRequestDataMaskingHeadersObservation struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type FrontendRequestDataMaskingHeadersParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type FrontendRequestDataMaskingObservation struct {

	// A headers block as defined below.
	Headers []FrontendRequestDataMaskingHeadersObservation `json:"headers,omitempty" tf:"headers,omitempty"`

	// A query_params block as defined below.
	QueryParams []FrontendRequestDataMaskingQueryParamsObservation `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type FrontendRequestDataMaskingParameters struct {

	// A headers block as defined below.
	// +kubebuilder:validation:Optional
	Headers []FrontendRequestDataMaskingHeadersParameters `json:"headers,omitempty" tf:"headers,omitempty"`

	// A query_params block as defined below.
	// +kubebuilder:validation:Optional
	QueryParams []FrontendRequestDataMaskingQueryParamsParameters `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type FrontendRequestDataMaskingQueryParamsObservation struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type FrontendRequestDataMaskingQueryParamsParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type FrontendRequestObservation struct {

	// Number of payload bytes to log (up to 8192).
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// A data_masking block as defined below.
	DataMasking []FrontendRequestDataMaskingObservation `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies a list of headers to log.
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type FrontendRequestParameters struct {

	// Number of payload bytes to log (up to 8192).
	// +kubebuilder:validation:Optional
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// A data_masking block as defined below.
	// +kubebuilder:validation:Optional
	DataMasking []FrontendRequestDataMaskingParameters `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies a list of headers to log.
	// +kubebuilder:validation:Optional
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type FrontendResponseDataMaskingHeadersObservation struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type FrontendResponseDataMaskingHeadersParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type FrontendResponseDataMaskingObservation struct {

	// A headers block as defined below.
	Headers []FrontendResponseDataMaskingHeadersObservation `json:"headers,omitempty" tf:"headers,omitempty"`

	// A query_params block as defined below.
	QueryParams []FrontendResponseDataMaskingQueryParamsObservation `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type FrontendResponseDataMaskingParameters struct {

	// A headers block as defined below.
	// +kubebuilder:validation:Optional
	Headers []FrontendResponseDataMaskingHeadersParameters `json:"headers,omitempty" tf:"headers,omitempty"`

	// A query_params block as defined below.
	// +kubebuilder:validation:Optional
	QueryParams []FrontendResponseDataMaskingQueryParamsParameters `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type FrontendResponseDataMaskingQueryParamsObservation struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type FrontendResponseDataMaskingQueryParamsParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type FrontendResponseObservation struct {

	// Number of payload bytes to log (up to 8192).
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// A data_masking block as defined below.
	DataMasking []FrontendResponseDataMaskingObservation `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies a list of headers to log.
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type FrontendResponseParameters struct {

	// Number of payload bytes to log (up to 8192).
	// +kubebuilder:validation:Optional
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// A data_masking block as defined below.
	// +kubebuilder:validation:Optional
	DataMasking []FrontendResponseDataMaskingParameters `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies a list of headers to log.
	// +kubebuilder:validation:Optional
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type HeadersObservation struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type HeadersParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type QueryParamsObservation struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type QueryParamsParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// APIDiagnosticSpec defines the desired state of APIDiagnostic
type APIDiagnosticSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APIDiagnosticParameters `json:"forProvider"`
}

// APIDiagnosticStatus defines the observed state of APIDiagnostic.
type APIDiagnosticStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APIDiagnosticObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// APIDiagnostic is the Schema for the APIDiagnostics API. Manages a API Management Service API Diagnostics Logs.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type APIDiagnostic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APIDiagnosticSpec   `json:"spec"`
	Status            APIDiagnosticStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIDiagnosticList contains a list of APIDiagnostics
type APIDiagnosticList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIDiagnostic `json:"items"`
}

// Repository type metadata.
var (
	APIDiagnostic_Kind             = "APIDiagnostic"
	APIDiagnostic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: APIDiagnostic_Kind}.String()
	APIDiagnostic_KindAPIVersion   = APIDiagnostic_Kind + "." + CRDGroupVersion.String()
	APIDiagnostic_GroupVersionKind = CRDGroupVersion.WithKind(APIDiagnostic_Kind)
)

func init() {
	SchemeBuilder.Register(&APIDiagnostic{}, &APIDiagnosticList{})
}
