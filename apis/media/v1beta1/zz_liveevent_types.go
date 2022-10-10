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

type CrossSiteAccessPolicyObservation struct {
}

type CrossSiteAccessPolicyParameters struct {

	// The content of clientaccesspolicy.xml used by Silverlight.
	// +kubebuilder:validation:Optional
	ClientAccessPolicy *string `json:"clientAccessPolicy,omitempty" tf:"client_access_policy,omitempty"`

	// The content of the Cross Domain Policy (crossdomain.xml).
	// +kubebuilder:validation:Optional
	CrossDomainPolicy *string `json:"crossDomainPolicy,omitempty" tf:"cross_domain_policy,omitempty"`
}

type EncodingObservation struct {
}

type EncodingParameters struct {

	// Use an ISO 8601 time value between 0.5 to 20 seconds to specify the output fragment length for the video and audio tracks of an encoding live event. For example, use PT2S to indicate 2 seconds. For the video track it also defines the key frame interval, or the length of a GoP (group of pictures). If this value is not set for an encoding live event, the fragment duration defaults to 2 seconds. The value cannot be set for pass-through live events.
	// +kubebuilder:validation:Optional
	KeyFrameInterval *string `json:"keyFrameInterval,omitempty" tf:"key_frame_interval,omitempty"`

	// The optional encoding preset name, used when type is not None. If the type is set to Standard, then the default preset name is Default720p. Else if the type is set to Premium1080p, the default preset is Default1080p. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PresetName *string `json:"presetName,omitempty" tf:"preset_name,omitempty"`

	// Specifies how the input video will be resized to fit the desired output resolution(s). Allowed values are None, AutoFit or AutoSize. Default is None.
	// +kubebuilder:validation:Optional
	StretchMode *string `json:"stretchMode,omitempty" tf:"stretch_mode,omitempty"`

	// Live event type. Allowed values are None, Premium1080p or Standard. When set to None, the service simply passes through the incoming video and audio layer(s) to the output. When type is set to Standard or Premium1080p, a live encoder transcodes the incoming stream into multiple bitrates or layers. Defaults to None. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type EndpointObservation struct {
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type EndpointParameters struct {
}

type IPAccessControlAllowObservation struct {
}

type IPAccessControlAllowParameters struct {

	// The IP address or CIDR range.
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The name which should be used for this Live Event. Changing this forces a new Live Event to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The subnet mask prefix length (see CIDR notation).
	// +kubebuilder:validation:Optional
	SubnetPrefixLength *float64 `json:"subnetPrefixLength,omitempty" tf:"subnet_prefix_length,omitempty"`
}

type InputObservation struct {
	Endpoint []EndpointObservation `json:"endpoint,omitempty" tf:"endpoint,omitempty"`
}

type InputParameters struct {

	// A UUID in string form to uniquely identify the stream. If omitted, the service will generate a unique value. Changing this forces a new value to be created.
	// +kubebuilder:validation:Optional
	AccessToken *string `json:"accessToken,omitempty" tf:"access_token,omitempty"`

	// One or more ip_access_control_allow blocks as defined below.
	// +kubebuilder:validation:Optional
	IPAccessControlAllow []IPAccessControlAllowParameters `json:"ipAccessControlAllow,omitempty" tf:"ip_access_control_allow,omitempty"`

	// ISO 8601 time duration of the key frame interval duration of the input. This value sets the EXT-X-TARGETDURATION property in the HLS output. For example, use PT2S to indicate 2 seconds. This field cannot be set when type is set to Encoding.
	// +kubebuilder:validation:Optional
	KeyFrameIntervalDuration *string `json:"keyFrameIntervalDuration,omitempty" tf:"key_frame_interval_duration,omitempty"`

	// The input protocol for the live event. Allowed values are FragmentedMP4 and RTMP. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	StreamingProtocol *string `json:"streamingProtocol,omitempty" tf:"streaming_protocol,omitempty"`
}

type LiveEventObservation struct {

	// The ID of the Live Event.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A input block as defined below.
	// +kubebuilder:validation:Required
	Input []InputObservation `json:"input,omitempty" tf:"input,omitempty"`

	// A preview block as defined below.
	// +kubebuilder:validation:Optional
	Preview []PreviewObservation `json:"preview,omitempty" tf:"preview,omitempty"`
}

type LiveEventParameters struct {

	// The flag indicates if the resource should be automatically started on creation. Default is false.
	// +kubebuilder:validation:Optional
	AutoStartEnabled *bool `json:"autoStartEnabled,omitempty" tf:"auto_start_enabled,omitempty"`

	// A cross_site_access_policy block as defined below.
	// +kubebuilder:validation:Optional
	CrossSiteAccessPolicy []CrossSiteAccessPolicyParameters `json:"crossSiteAccessPolicy,omitempty" tf:"cross_site_access_policy,omitempty"`

	// A description for the live event.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A encoding block as defined below.
	// +kubebuilder:validation:Optional
	Encoding []EncodingParameters `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// When use_static_hostname is set to true, the hostname_prefix specifies the first part of the hostname assigned to the live event preview and ingest endpoints. The final hostname would be a combination of this prefix, the media service account name and a short code for the Azure Media Services data center.
	// +kubebuilder:validation:Optional
	HostNamePrefix *string `json:"hostnamePrefix,omitempty" tf:"hostname_prefix,omitempty"`

	// A input block as defined below.
	// +kubebuilder:validation:Required
	Input []InputParameters `json:"input" tf:"input,omitempty"`

	// The Azure Region where the Live Event should exist. Changing this forces a new Live Event to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The Media Services account name. Changing this forces a new Live Event to be created.
	// +crossplane:generate:reference:type=ServicesAccount
	// +kubebuilder:validation:Optional
	MediaServicesAccountName *string `json:"mediaServicesAccountName,omitempty" tf:"media_services_account_name,omitempty"`

	// Reference to a ServicesAccount to populate mediaServicesAccountName.
	// +kubebuilder:validation:Optional
	MediaServicesAccountNameRef *v1.Reference `json:"mediaServicesAccountNameRef,omitempty" tf:"-"`

	// Selector for a ServicesAccount to populate mediaServicesAccountName.
	// +kubebuilder:validation:Optional
	MediaServicesAccountNameSelector *v1.Selector `json:"mediaServicesAccountNameSelector,omitempty" tf:"-"`

	// A preview block as defined below.
	// +kubebuilder:validation:Optional
	Preview []PreviewParameters `json:"preview,omitempty" tf:"preview,omitempty"`

	// The name of the Resource Group where the Live Event should exist. Changing this forces a new Live Event to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Live Event.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies a list of languages (locale) to be used for speech-to-text transcription – it should match the spoken language in the audio track. The value should be in BCP-47 format (e.g: en-US). See the Microsoft Documentation for more information about the live transcription feature and the list of supported languages.
	// +kubebuilder:validation:Optional
	TranscriptionLanguages []*string `json:"transcriptionLanguages,omitempty" tf:"transcription_languages,omitempty"`

	// Specifies whether a static hostname would be assigned to the live event preview and ingest endpoints. Changing this forces a new Live Event to be created.
	// +kubebuilder:validation:Optional
	UseStaticHostName *bool `json:"useStaticHostname,omitempty" tf:"use_static_hostname,omitempty"`
}

type PreviewEndpointObservation struct {
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type PreviewEndpointParameters struct {
}

type PreviewIPAccessControlAllowObservation struct {
}

type PreviewIPAccessControlAllowParameters struct {

	// The IP address or CIDR range.
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The name which should be used for this Live Event. Changing this forces a new Live Event to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The subnet mask prefix length (see CIDR notation).
	// +kubebuilder:validation:Optional
	SubnetPrefixLength *float64 `json:"subnetPrefixLength,omitempty" tf:"subnet_prefix_length,omitempty"`
}

type PreviewObservation struct {
	Endpoint []PreviewEndpointObservation `json:"endpoint,omitempty" tf:"endpoint,omitempty"`
}

type PreviewParameters struct {

	// An alternative media identifier associated with the streaming locator created for the preview. The identifier can be used in the CustomLicenseAcquisitionUrlTemplate or the CustomKeyAcquisitionUrlTemplate of the Streaming Policy specified in the streaming_policy_name field. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	AlternativeMediaID *string `json:"alternativeMediaId,omitempty" tf:"alternative_media_id,omitempty"`

	// One or more ip_access_control_allow blocks as defined above.
	// +kubebuilder:validation:Optional
	IPAccessControlAllow []PreviewIPAccessControlAllowParameters `json:"ipAccessControlAllow,omitempty" tf:"ip_access_control_allow,omitempty"`

	// The identifier of the preview locator in GUID format. Specifying this at creation time allows the caller to know the preview locator url before the event is created. If omitted, the service will generate a random identifier. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PreviewLocator *string `json:"previewLocator,omitempty" tf:"preview_locator,omitempty"`

	// The name of streaming policy used for the live event preview. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	StreamingPolicyName *string `json:"streamingPolicyName,omitempty" tf:"streaming_policy_name,omitempty"`
}

// LiveEventSpec defines the desired state of LiveEvent
type LiveEventSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LiveEventParameters `json:"forProvider"`
}

// LiveEventStatus defines the observed state of LiveEvent.
type LiveEventStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LiveEventObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LiveEvent is the Schema for the LiveEvents API. Manages a Live Event.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LiveEvent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LiveEventSpec   `json:"spec"`
	Status            LiveEventStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LiveEventList contains a list of LiveEvents
type LiveEventList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LiveEvent `json:"items"`
}

// Repository type metadata.
var (
	LiveEvent_Kind             = "LiveEvent"
	LiveEvent_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LiveEvent_Kind}.String()
	LiveEvent_KindAPIVersion   = LiveEvent_Kind + "." + CRDGroupVersion.String()
	LiveEvent_GroupVersionKind = CRDGroupVersion.WithKind(LiveEvent_Kind)
)

func init() {
	SchemeBuilder.Register(&LiveEvent{}, &LiveEventList{})
}