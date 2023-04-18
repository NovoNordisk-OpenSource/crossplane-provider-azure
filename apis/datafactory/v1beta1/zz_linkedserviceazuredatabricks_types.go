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

type InstancePoolObservation struct {

	// Spark version of a the cluster.
	ClusterVersion *string `json:"clusterVersion,omitempty" tf:"cluster_version,omitempty"`

	// Identifier of the instance pool within the linked ADB instance.
	InstancePoolID *string `json:"instancePoolId,omitempty" tf:"instance_pool_id,omitempty"`

	// The max number of worker nodes. Set this value if you want to enable autoscaling between the min_number_of_workers and this value. Omit this value to use a fixed number of workers defined in the min_number_of_workers property.
	MaxNumberOfWorkers *float64 `json:"maxNumberOfWorkers,omitempty" tf:"max_number_of_workers,omitempty"`

	// The minimum number of worker nodes. Defaults to 1.
	MinNumberOfWorkers *float64 `json:"minNumberOfWorkers,omitempty" tf:"min_number_of_workers,omitempty"`
}

type InstancePoolParameters struct {

	// Spark version of a the cluster.
	// +kubebuilder:validation:Required
	ClusterVersion *string `json:"clusterVersion" tf:"cluster_version,omitempty"`

	// Identifier of the instance pool within the linked ADB instance.
	// +kubebuilder:validation:Required
	InstancePoolID *string `json:"instancePoolId" tf:"instance_pool_id,omitempty"`

	// The max number of worker nodes. Set this value if you want to enable autoscaling between the min_number_of_workers and this value. Omit this value to use a fixed number of workers defined in the min_number_of_workers property.
	// +kubebuilder:validation:Optional
	MaxNumberOfWorkers *float64 `json:"maxNumberOfWorkers,omitempty" tf:"max_number_of_workers,omitempty"`

	// The minimum number of worker nodes. Defaults to 1.
	// +kubebuilder:validation:Optional
	MinNumberOfWorkers *float64 `json:"minNumberOfWorkers,omitempty" tf:"min_number_of_workers,omitempty"`
}

type LinkedServiceAzureDatabricksKeyVaultPasswordObservation struct {

	// Specifies the name of an existing Key Vault Data Factory Linked Service.
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Specifies the secret name in Azure Key Vault that stores ADB access token.
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`
}

type LinkedServiceAzureDatabricksKeyVaultPasswordParameters struct {

	// Specifies the name of an existing Key Vault Data Factory Linked Service.
	// +kubebuilder:validation:Required
	LinkedServiceName *string `json:"linkedServiceName" tf:"linked_service_name,omitempty"`

	// Specifies the secret name in Azure Key Vault that stores ADB access token.
	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type LinkedServiceAzureDatabricksObservation struct {

	// The domain URL of the databricks instance.
	AdbDomain *string `json:"adbDomain,omitempty" tf:"adb_domain,omitempty"`

	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// The description for the Data Factory Linked Service.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The cluster_id of an existing cluster within the linked ADB instance.
	ExistingClusterID *string `json:"existingClusterId,omitempty" tf:"existing_cluster_id,omitempty"`

	// The ID of the Data Factory Linked Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Leverages an instance pool within the linked ADB instance as defined by instance_pool block below.
	InstancePool []InstancePoolObservation `json:"instancePool,omitempty" tf:"instance_pool,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// Authenticate to ADB via Azure Key Vault Linked Service as defined in the key_vault_password block below.
	KeyVaultPassword []LinkedServiceAzureDatabricksKeyVaultPasswordObservation `json:"keyVaultPassword,omitempty" tf:"key_vault_password,omitempty"`

	// Authenticate to ADB via managed service identity.
	MsiWorkSpaceResourceID *string `json:"msiWorkSpaceResourceId,omitempty" tf:"msi_work_space_resource_id,omitempty"`

	// Creates new clusters within the linked ADB instance as defined in the new_cluster_config block below.
	NewClusterConfig []NewClusterConfigObservation `json:"newClusterConfig,omitempty" tf:"new_cluster_config,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type LinkedServiceAzureDatabricksParameters struct {

	// Authenticate to ADB via an access token.
	// +kubebuilder:validation:Optional
	AccessTokenSecretRef *v1.SecretKeySelector `json:"accessTokenSecretRef,omitempty" tf:"-"`

	// The domain URL of the databricks instance.
	// +kubebuilder:validation:Optional
	AdbDomain *string `json:"adbDomain,omitempty" tf:"adb_domain,omitempty"`

	// A map of additional properties to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

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

	// The description for the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The cluster_id of an existing cluster within the linked ADB instance.
	// +kubebuilder:validation:Optional
	ExistingClusterID *string `json:"existingClusterId,omitempty" tf:"existing_cluster_id,omitempty"`

	// Leverages an instance pool within the linked ADB instance as defined by instance_pool block below.
	// +kubebuilder:validation:Optional
	InstancePool []InstancePoolParameters `json:"instancePool,omitempty" tf:"instance_pool,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// Authenticate to ADB via Azure Key Vault Linked Service as defined in the key_vault_password block below.
	// +kubebuilder:validation:Optional
	KeyVaultPassword []LinkedServiceAzureDatabricksKeyVaultPasswordParameters `json:"keyVaultPassword,omitempty" tf:"key_vault_password,omitempty"`

	// Authenticate to ADB via managed service identity.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/databricks/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	MsiWorkSpaceResourceID *string `json:"msiWorkSpaceResourceId,omitempty" tf:"msi_work_space_resource_id,omitempty"`

	// Reference to a Workspace in databricks to populate msiWorkSpaceResourceId.
	// +kubebuilder:validation:Optional
	MsiWorkSpaceResourceIDRef *v1.Reference `json:"msiWorkSpaceResourceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in databricks to populate msiWorkSpaceResourceId.
	// +kubebuilder:validation:Optional
	MsiWorkSpaceResourceIDSelector *v1.Selector `json:"msiWorkSpaceResourceIdSelector,omitempty" tf:"-"`

	// Creates new clusters within the linked ADB instance as defined in the new_cluster_config block below.
	// +kubebuilder:validation:Optional
	NewClusterConfig []NewClusterConfigParameters `json:"newClusterConfig,omitempty" tf:"new_cluster_config,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type NewClusterConfigObservation struct {

	// Spark version of a the cluster.
	ClusterVersion *string `json:"clusterVersion,omitempty" tf:"cluster_version,omitempty"`

	// Tags for the cluster resource.
	CustomTags map[string]*string `json:"customTags,omitempty" tf:"custom_tags,omitempty"`

	// Driver node type for the cluster.
	DriverNodeType *string `json:"driverNodeType,omitempty" tf:"driver_node_type,omitempty"`

	// User defined initialization scripts for the cluster.
	InitScripts []*string `json:"initScripts,omitempty" tf:"init_scripts,omitempty"`

	// Location to deliver Spark driver, worker, and event logs.
	LogDestination *string `json:"logDestination,omitempty" tf:"log_destination,omitempty"`

	// Specifies the maximum number of worker nodes. It should be between 1 and 25000.
	MaxNumberOfWorkers *float64 `json:"maxNumberOfWorkers,omitempty" tf:"max_number_of_workers,omitempty"`

	// Specifies the minimum number of worker nodes. It should be between 1 and 25000. It defaults to 1.
	MinNumberOfWorkers *float64 `json:"minNumberOfWorkers,omitempty" tf:"min_number_of_workers,omitempty"`

	// Node type for the new cluster.
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// User-specified Spark configuration variables key-value pairs.
	SparkConfig map[string]*string `json:"sparkConfig,omitempty" tf:"spark_config,omitempty"`

	// User-specified Spark environment variables key-value pairs.
	SparkEnvironmentVariables map[string]*string `json:"sparkEnvironmentVariables,omitempty" tf:"spark_environment_variables,omitempty"`
}

type NewClusterConfigParameters struct {

	// Spark version of a the cluster.
	// +kubebuilder:validation:Required
	ClusterVersion *string `json:"clusterVersion" tf:"cluster_version,omitempty"`

	// Tags for the cluster resource.
	// +kubebuilder:validation:Optional
	CustomTags map[string]*string `json:"customTags,omitempty" tf:"custom_tags,omitempty"`

	// Driver node type for the cluster.
	// +kubebuilder:validation:Optional
	DriverNodeType *string `json:"driverNodeType,omitempty" tf:"driver_node_type,omitempty"`

	// User defined initialization scripts for the cluster.
	// +kubebuilder:validation:Optional
	InitScripts []*string `json:"initScripts,omitempty" tf:"init_scripts,omitempty"`

	// Location to deliver Spark driver, worker, and event logs.
	// +kubebuilder:validation:Optional
	LogDestination *string `json:"logDestination,omitempty" tf:"log_destination,omitempty"`

	// Specifies the maximum number of worker nodes. It should be between 1 and 25000.
	// +kubebuilder:validation:Optional
	MaxNumberOfWorkers *float64 `json:"maxNumberOfWorkers,omitempty" tf:"max_number_of_workers,omitempty"`

	// Specifies the minimum number of worker nodes. It should be between 1 and 25000. It defaults to 1.
	// +kubebuilder:validation:Optional
	MinNumberOfWorkers *float64 `json:"minNumberOfWorkers,omitempty" tf:"min_number_of_workers,omitempty"`

	// Node type for the new cluster.
	// +kubebuilder:validation:Required
	NodeType *string `json:"nodeType" tf:"node_type,omitempty"`

	// User-specified Spark configuration variables key-value pairs.
	// +kubebuilder:validation:Optional
	SparkConfig map[string]*string `json:"sparkConfig,omitempty" tf:"spark_config,omitempty"`

	// User-specified Spark environment variables key-value pairs.
	// +kubebuilder:validation:Optional
	SparkEnvironmentVariables map[string]*string `json:"sparkEnvironmentVariables,omitempty" tf:"spark_environment_variables,omitempty"`
}

// LinkedServiceAzureDatabricksSpec defines the desired state of LinkedServiceAzureDatabricks
type LinkedServiceAzureDatabricksSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedServiceAzureDatabricksParameters `json:"forProvider"`
}

// LinkedServiceAzureDatabricksStatus defines the observed state of LinkedServiceAzureDatabricks.
type LinkedServiceAzureDatabricksStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedServiceAzureDatabricksObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceAzureDatabricks is the Schema for the LinkedServiceAzureDatabrickss API. Manages a Linked Service (connection) between Azure Databricks and Azure Data Factory.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LinkedServiceAzureDatabricks struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.adbDomain)",message="adbDomain is a required parameter"
	Spec   LinkedServiceAzureDatabricksSpec   `json:"spec"`
	Status LinkedServiceAzureDatabricksStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceAzureDatabricksList contains a list of LinkedServiceAzureDatabrickss
type LinkedServiceAzureDatabricksList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedServiceAzureDatabricks `json:"items"`
}

// Repository type metadata.
var (
	LinkedServiceAzureDatabricks_Kind             = "LinkedServiceAzureDatabricks"
	LinkedServiceAzureDatabricks_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedServiceAzureDatabricks_Kind}.String()
	LinkedServiceAzureDatabricks_KindAPIVersion   = LinkedServiceAzureDatabricks_Kind + "." + CRDGroupVersion.String()
	LinkedServiceAzureDatabricks_GroupVersionKind = CRDGroupVersion.WithKind(LinkedServiceAzureDatabricks_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedServiceAzureDatabricks{}, &LinkedServiceAzureDatabricksList{})
}
