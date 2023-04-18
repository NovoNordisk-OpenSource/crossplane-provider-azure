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

type LoadBalancerNatPoolObservation struct {

	// The port used for the internal endpoint. Possible values range between 1 and 65535, inclusive.
	BackendPort *float64 `json:"backendPort,omitempty" tf:"backend_port,omitempty"`

	// Are the floating IPs enabled for this Load Balancer Rule? A floating IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group.
	FloatingIPEnabled *bool `json:"floatingIpEnabled,omitempty" tf:"floating_ip_enabled,omitempty"`

	// The ID of the Load Balancer NAT pool.
	FrontendIPConfigurationID *string `json:"frontendIpConfigurationId,omitempty" tf:"frontend_ip_configuration_id,omitempty"`

	// The name of the frontend IP configuration exposing this rule.
	FrontendIPConfigurationName *string `json:"frontendIpConfigurationName,omitempty" tf:"frontend_ip_configuration_name,omitempty"`

	// The last port number in the range of external ports that will be used to provide Inbound NAT to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.
	FrontendPortEnd *float64 `json:"frontendPortEnd,omitempty" tf:"frontend_port_end,omitempty"`

	// The first port number in the range of external ports that will be used to provide Inbound NAT to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.
	FrontendPortStart *float64 `json:"frontendPortStart,omitempty" tf:"frontend_port_start,omitempty"`

	// The ID of the Load Balancer NAT pool.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the idle timeout in minutes for TCP connections. Valid values are between 4 and 30. Defaults to 4.
	IdleTimeoutInMinutes *float64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// The ID of the Load Balancer in which to create the NAT pool. Changing this forces a new resource to be created.
	LoadbalancerID *string `json:"loadbalancerId,omitempty" tf:"loadbalancer_id,omitempty"`

	// The transport protocol for the external endpoint. Possible values are All, Tcp and Udp.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Is TCP Reset enabled for this Load Balancer Rule?
	TCPResetEnabled *bool `json:"tcpResetEnabled,omitempty" tf:"tcp_reset_enabled,omitempty"`
}

type LoadBalancerNatPoolParameters struct {

	// The port used for the internal endpoint. Possible values range between 1 and 65535, inclusive.
	// +kubebuilder:validation:Optional
	BackendPort *float64 `json:"backendPort,omitempty" tf:"backend_port,omitempty"`

	// Are the floating IPs enabled for this Load Balancer Rule? A floating IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group.
	// +kubebuilder:validation:Optional
	FloatingIPEnabled *bool `json:"floatingIpEnabled,omitempty" tf:"floating_ip_enabled,omitempty"`

	// The name of the frontend IP configuration exposing this rule.
	// +kubebuilder:validation:Optional
	FrontendIPConfigurationName *string `json:"frontendIpConfigurationName,omitempty" tf:"frontend_ip_configuration_name,omitempty"`

	// The last port number in the range of external ports that will be used to provide Inbound NAT to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.
	// +kubebuilder:validation:Optional
	FrontendPortEnd *float64 `json:"frontendPortEnd,omitempty" tf:"frontend_port_end,omitempty"`

	// The first port number in the range of external ports that will be used to provide Inbound NAT to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.
	// +kubebuilder:validation:Optional
	FrontendPortStart *float64 `json:"frontendPortStart,omitempty" tf:"frontend_port_start,omitempty"`

	// Specifies the idle timeout in minutes for TCP connections. Valid values are between 4 and 30. Defaults to 4.
	// +kubebuilder:validation:Optional
	IdleTimeoutInMinutes *float64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// The ID of the Load Balancer in which to create the NAT pool. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=LoadBalancer
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LoadbalancerID *string `json:"loadbalancerId,omitempty" tf:"loadbalancer_id,omitempty"`

	// Reference to a LoadBalancer to populate loadbalancerId.
	// +kubebuilder:validation:Optional
	LoadbalancerIDRef *v1.Reference `json:"loadbalancerIdRef,omitempty" tf:"-"`

	// Selector for a LoadBalancer to populate loadbalancerId.
	// +kubebuilder:validation:Optional
	LoadbalancerIDSelector *v1.Selector `json:"loadbalancerIdSelector,omitempty" tf:"-"`

	// The transport protocol for the external endpoint. Possible values are All, Tcp and Udp.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Is TCP Reset enabled for this Load Balancer Rule?
	// +kubebuilder:validation:Optional
	TCPResetEnabled *bool `json:"tcpResetEnabled,omitempty" tf:"tcp_reset_enabled,omitempty"`
}

// LoadBalancerNatPoolSpec defines the desired state of LoadBalancerNatPool
type LoadBalancerNatPoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadBalancerNatPoolParameters `json:"forProvider"`
}

// LoadBalancerNatPoolStatus defines the observed state of LoadBalancerNatPool.
type LoadBalancerNatPoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadBalancerNatPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerNatPool is the Schema for the LoadBalancerNatPools API. Manages a Load Balancer NAT Pool.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LoadBalancerNatPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.backendPort)",message="backendPort is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.frontendIpConfigurationName)",message="frontendIpConfigurationName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.frontendPortEnd)",message="frontendPortEnd is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.frontendPortStart)",message="frontendPortStart is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.protocol)",message="protocol is a required parameter"
	Spec   LoadBalancerNatPoolSpec   `json:"spec"`
	Status LoadBalancerNatPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerNatPoolList contains a list of LoadBalancerNatPools
type LoadBalancerNatPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancerNatPool `json:"items"`
}

// Repository type metadata.
var (
	LoadBalancerNatPool_Kind             = "LoadBalancerNatPool"
	LoadBalancerNatPool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoadBalancerNatPool_Kind}.String()
	LoadBalancerNatPool_KindAPIVersion   = LoadBalancerNatPool_Kind + "." + CRDGroupVersion.String()
	LoadBalancerNatPool_GroupVersionKind = CRDGroupVersion.WithKind(LoadBalancerNatPool_Kind)
)

func init() {
	SchemeBuilder.Register(&LoadBalancerNatPool{}, &LoadBalancerNatPoolList{})
}
