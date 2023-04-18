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

type FirewallIPConfigurationObservation struct {

	// Specifies the name of the IP Configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Private IP address of the Azure Firewall.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The ID of the Public IP Address associated with the firewall.
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// Reference to the subnet associated with the IP Configuration. Changing this forces a new resource to be created.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type FirewallIPConfigurationParameters struct {

	// Specifies the name of the IP Configuration.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID of the Public IP Address associated with the firewall.
	// +crossplane:generate:reference:type=PublicIP
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// Reference to a PublicIP to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDRef *v1.Reference `json:"publicIpAddressIdRef,omitempty" tf:"-"`

	// Selector for a PublicIP to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDSelector *v1.Selector `json:"publicIpAddressIdSelector,omitempty" tf:"-"`

	// Reference to the subnet associated with the IP Configuration. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type FirewallObservation struct {

	// A list of DNS servers that the Azure Firewall will direct DNS traffic to the for name resolution.
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// The ID of the Firewall Policy applied to this Firewall.
	FirewallPolicyID *string `json:"firewallPolicyId,omitempty" tf:"firewall_policy_id,omitempty"`

	// The ID of the Azure Firewall.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An ip_configuration block as documented below.
	IPConfiguration []FirewallIPConfigurationObservation `json:"ipConfiguration,omitempty" tf:"ip_configuration,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A management_ip_configuration block as documented below, which allows force-tunnelling of traffic to be performed by the firewall. Adding or removing this block or changing the subnet_id in an existing block forces a new resource to be created. Changing this forces a new resource to be created.
	ManagementIPConfiguration []ManagementIPConfigurationObservation `json:"managementIpConfiguration,omitempty" tf:"management_ip_configuration,omitempty"`

	// A list of SNAT private CIDR IP ranges, or the special string IANAPrivateRanges, which indicates Azure Firewall does not SNAT when the destination IP address is a private range per IANA RFC 1918.
	PrivateIPRanges []*string `json:"privateIpRanges,omitempty" tf:"private_ip_ranges,omitempty"`

	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// SKU name of the Firewall. Possible values are AZFW_Hub and AZFW_VNet. Changing this forces a new resource to be created.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// SKU tier of the Firewall. Possible values are Premium, Standard and Basic.
	SkuTier *string `json:"skuTier,omitempty" tf:"sku_tier,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The operation mode for threat intelligence-based filtering. Possible values are: Off, Alert and Deny. Defaults to Alert.
	ThreatIntelMode *string `json:"threatIntelMode,omitempty" tf:"threat_intel_mode,omitempty"`

	// A virtual_hub block as documented below.
	VirtualHub []VirtualHubObservation `json:"virtualHub,omitempty" tf:"virtual_hub,omitempty"`

	// Specifies a list of Availability Zones in which this Azure Firewall should be located. Changing this forces a new Azure Firewall to be created.
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type FirewallParameters struct {

	// A list of DNS servers that the Azure Firewall will direct DNS traffic to the for name resolution.
	// +kubebuilder:validation:Optional
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// The ID of the Firewall Policy applied to this Firewall.
	// +kubebuilder:validation:Optional
	FirewallPolicyID *string `json:"firewallPolicyId,omitempty" tf:"firewall_policy_id,omitempty"`

	// An ip_configuration block as documented below.
	// +kubebuilder:validation:Optional
	IPConfiguration []FirewallIPConfigurationParameters `json:"ipConfiguration,omitempty" tf:"ip_configuration,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A management_ip_configuration block as documented below, which allows force-tunnelling of traffic to be performed by the firewall. Adding or removing this block or changing the subnet_id in an existing block forces a new resource to be created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ManagementIPConfiguration []ManagementIPConfigurationParameters `json:"managementIpConfiguration,omitempty" tf:"management_ip_configuration,omitempty"`

	// A list of SNAT private CIDR IP ranges, or the special string IANAPrivateRanges, which indicates Azure Firewall does not SNAT when the destination IP address is a private range per IANA RFC 1918.
	// +kubebuilder:validation:Optional
	PrivateIPRanges []*string `json:"privateIpRanges,omitempty" tf:"private_ip_ranges,omitempty"`

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

	// SKU name of the Firewall. Possible values are AZFW_Hub and AZFW_VNet. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// SKU tier of the Firewall. Possible values are Premium, Standard and Basic.
	// +kubebuilder:validation:Optional
	SkuTier *string `json:"skuTier,omitempty" tf:"sku_tier,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The operation mode for threat intelligence-based filtering. Possible values are: Off, Alert and Deny. Defaults to Alert.
	// +kubebuilder:validation:Optional
	ThreatIntelMode *string `json:"threatIntelMode,omitempty" tf:"threat_intel_mode,omitempty"`

	// A virtual_hub block as documented below.
	// +kubebuilder:validation:Optional
	VirtualHub []VirtualHubParameters `json:"virtualHub,omitempty" tf:"virtual_hub,omitempty"`

	// Specifies a list of Availability Zones in which this Azure Firewall should be located. Changing this forces a new Azure Firewall to be created.
	// +kubebuilder:validation:Optional
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type ManagementIPConfigurationObservation struct {

	// Specifies the name of the IP Configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The private IP address associated with the Firewall.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The ID of the Public IP Address associated with the firewall.
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// Reference to the subnet associated with the IP Configuration. Changing this forces a new resource to be created.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type ManagementIPConfigurationParameters struct {

	// Specifies the name of the IP Configuration.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID of the Public IP Address associated with the firewall.
	// +kubebuilder:validation:Required
	PublicIPAddressID *string `json:"publicIpAddressId" tf:"public_ip_address_id,omitempty"`

	// Reference to the subnet associated with the IP Configuration. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type VirtualHubObservation struct {

	// The private IP address associated with the Firewall.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The list of public IP addresses associated with the Firewall.
	PublicIPAddresses []*string `json:"publicIpAddresses,omitempty" tf:"public_ip_addresses,omitempty"`

	// Specifies the number of public IPs to assign to the Firewall. Defaults to 1.
	PublicIPCount *float64 `json:"publicIpCount,omitempty" tf:"public_ip_count,omitempty"`

	// Specifies the ID of the Virtual Hub where the Firewall resides in.
	VirtualHubID *string `json:"virtualHubId,omitempty" tf:"virtual_hub_id,omitempty"`
}

type VirtualHubParameters struct {

	// Specifies the number of public IPs to assign to the Firewall. Defaults to 1.
	// +kubebuilder:validation:Optional
	PublicIPCount *float64 `json:"publicIpCount,omitempty" tf:"public_ip_count,omitempty"`

	// Specifies the ID of the Virtual Hub where the Firewall resides in.
	// +kubebuilder:validation:Required
	VirtualHubID *string `json:"virtualHubId" tf:"virtual_hub_id,omitempty"`
}

// FirewallSpec defines the desired state of Firewall
type FirewallSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallParameters `json:"forProvider"`
}

// FirewallStatus defines the observed state of Firewall.
type FirewallStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Firewall is the Schema for the Firewalls API. Manages an Azure Firewall.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Firewall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.skuName)",message="skuName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.skuTier)",message="skuTier is a required parameter"
	Spec   FirewallSpec   `json:"spec"`
	Status FirewallStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallList contains a list of Firewalls
type FirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Firewall `json:"items"`
}

// Repository type metadata.
var (
	Firewall_Kind             = "Firewall"
	Firewall_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Firewall_Kind}.String()
	Firewall_KindAPIVersion   = Firewall_Kind + "." + CRDGroupVersion.String()
	Firewall_GroupVersionKind = CRDGroupVersion.WithKind(Firewall_Kind)
)

func init() {
	SchemeBuilder.Register(&Firewall{}, &FirewallList{})
}
