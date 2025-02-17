//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragepool

import "time"

// ACL - Access Control List (ACL) for an iSCSI Target; defines LUN masking policy
type ACL struct {
	// REQUIRED; iSCSI initiator IQN (iSCSI Qualified Name); example: "iqn.2005-03.org.iscsi:client".
	InitiatorIqn *string `json:"initiatorIqn,omitempty"`

	// REQUIRED; List of LUN names mapped to the ACL.
	MappedLuns []*string `json:"mappedLuns,omitempty"`
}

// Disk - Azure Managed Disk to attach to the Disk Pool.
type Disk struct {
	// REQUIRED; Unique Azure Resource ID of the Managed Disk.
	ID *string `json:"id,omitempty"`
}

// DiskPool - Response for Disk Pool request.
type DiskPool struct {
	// REQUIRED; The geo-location where the resource lives.
	Location *string `json:"location,omitempty"`

	// REQUIRED; Properties of Disk Pool.
	Properties *DiskPoolProperties `json:"properties,omitempty"`

	// Determines the SKU of the Disk pool
	SKU *SKU `json:"sku,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Azure resource id. Indicates if this resource is managed by another Azure resource.
	ManagedBy *string `json:"managedBy,omitempty" azure:"ro"`

	// READ-ONLY; List of Azure resource ids that manage this resource.
	ManagedByExtended []*string `json:"managedByExtended,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource metadata required by ARM RPC
	SystemData *SystemMetadata `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// DiskPoolCreate - Request payload for create or update Disk Pool request.
type DiskPoolCreate struct {
	// REQUIRED; The geo-location where the resource lives.
	Location *string `json:"location,omitempty"`

	// REQUIRED; Properties for Disk Pool create request.
	Properties *DiskPoolCreateProperties `json:"properties,omitempty"`

	// REQUIRED; Determines the SKU of the Disk Pool
	SKU *SKU `json:"sku,omitempty"`

	// Azure resource id. Indicates if this resource is managed by another Azure resource.
	ManagedBy *string `json:"managedBy,omitempty"`

	// List of Azure resource ids that manage this resource.
	ManagedByExtended []*string `json:"managedByExtended,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// DiskPoolCreateProperties - Properties for Disk Pool create or update request.
type DiskPoolCreateProperties struct {
	// REQUIRED; Azure Resource ID of a Subnet for the Disk Pool.
	SubnetID *string `json:"subnetId,omitempty"`

	// List of additional capabilities for a Disk Pool.
	AdditionalCapabilities []*string `json:"additionalCapabilities,omitempty"`

	// Logical zone for Disk Pool resource; example: ["1"].
	AvailabilityZones []*string `json:"availabilityZones,omitempty"`

	// List of Azure Managed Disks to attach to a Disk Pool.
	Disks []*Disk `json:"disks,omitempty"`
}

// DiskPoolListResult - List of Disk Pools
type DiskPoolListResult struct {
	// REQUIRED; An array of Disk pool objects.
	Value []*DiskPool `json:"value,omitempty"`

	// READ-ONLY; URI to fetch the next section of the paginated response.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// DiskPoolProperties - Disk Pool response properties.
type DiskPoolProperties struct {
	// REQUIRED; Logical zone for Disk Pool resource; example: ["1"].
	AvailabilityZones []*string `json:"availabilityZones,omitempty"`

	// REQUIRED; Operational status of the Disk Pool.
	Status *OperationalStatus `json:"status,omitempty"`

	// REQUIRED; Azure Resource ID of a Subnet for the Disk Pool.
	SubnetID *string `json:"subnetId,omitempty"`

	// READ-ONLY; State of the operation on the resource.
	ProvisioningState *ProvisioningStates `json:"provisioningState,omitempty" azure:"ro"`

	// List of additional capabilities for Disk Pool.
	AdditionalCapabilities []*string `json:"additionalCapabilities,omitempty"`

	// List of Azure Managed Disks to attach to a Disk Pool.
	Disks []*Disk `json:"disks,omitempty"`
}

// DiskPoolUpdate - Request payload for Update Disk Pool request.
type DiskPoolUpdate struct {
	// REQUIRED; Properties for Disk Pool update request.
	Properties *DiskPoolUpdateProperties `json:"properties,omitempty"`

	// Azure resource id. Indicates if this resource is managed by another Azure resource.
	ManagedBy *string `json:"managedBy,omitempty"`

	// List of Azure resource ids that manage this resource.
	ManagedByExtended []*string `json:"managedByExtended,omitempty"`

	// Determines the SKU of the Disk Pool
	SKU *SKU `json:"sku,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// DiskPoolUpdateProperties - Properties for Disk Pool update request.
type DiskPoolUpdateProperties struct {
	// List of Azure Managed Disks to attach to a Disk Pool.
	Disks []*Disk `json:"disks,omitempty"`
}

// DiskPoolZoneInfo - Disk Pool SKU Details
type DiskPoolZoneInfo struct {
	// READ-ONLY; List of additional capabilities for Disk Pool.
	AdditionalCapabilities []*string `json:"additionalCapabilities,omitempty" azure:"ro"`

	// READ-ONLY; Logical zone for Disk Pool resource; example: ["1"].
	AvailabilityZones []*string `json:"availabilityZones,omitempty" azure:"ro"`

	// READ-ONLY; Determines the SKU of VM deployed for Disk Pool
	SKU *SKU `json:"sku,omitempty" azure:"ro"`
}

// DiskPoolZoneListResult - List Disk Pool skus operation response.
type DiskPoolZoneListResult struct {
	// READ-ONLY; URI to fetch the next section of the paginated response.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; The list of Disk Pool Skus.
	Value []*DiskPoolZoneInfo `json:"value,omitempty" azure:"ro"`
}

// DiskPoolZonesClientListOptions contains the optional parameters for the DiskPoolZonesClient.List method.
type DiskPoolZonesClientListOptions struct {
	// placeholder for future optional parameters
}

// DiskPoolsClientBeginCreateOrUpdateOptions contains the optional parameters for the DiskPoolsClient.BeginCreateOrUpdate
// method.
type DiskPoolsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DiskPoolsClientBeginDeallocateOptions contains the optional parameters for the DiskPoolsClient.BeginDeallocate method.
type DiskPoolsClientBeginDeallocateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DiskPoolsClientBeginDeleteOptions contains the optional parameters for the DiskPoolsClient.BeginDelete method.
type DiskPoolsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DiskPoolsClientBeginStartOptions contains the optional parameters for the DiskPoolsClient.BeginStart method.
type DiskPoolsClientBeginStartOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DiskPoolsClientBeginUpdateOptions contains the optional parameters for the DiskPoolsClient.BeginUpdate method.
type DiskPoolsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DiskPoolsClientBeginUpgradeOptions contains the optional parameters for the DiskPoolsClient.BeginUpgrade method.
type DiskPoolsClientBeginUpgradeOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DiskPoolsClientGetOptions contains the optional parameters for the DiskPoolsClient.Get method.
type DiskPoolsClientGetOptions struct {
	// placeholder for future optional parameters
}

// DiskPoolsClientListByResourceGroupOptions contains the optional parameters for the DiskPoolsClient.ListByResourceGroup
// method.
type DiskPoolsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// DiskPoolsClientListBySubscriptionOptions contains the optional parameters for the DiskPoolsClient.ListBySubscription method.
type DiskPoolsClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// DiskPoolsClientListOutboundNetworkDependenciesEndpointsOptions contains the optional parameters for the DiskPoolsClient.ListOutboundNetworkDependenciesEndpoints
// method.
type DiskPoolsClientListOutboundNetworkDependenciesEndpointsOptions struct {
	// placeholder for future optional parameters
}

// EndpointDependency - A domain name that a service is reached at, including details of the current connection status.
type EndpointDependency struct {
	// The domain name of the dependency.
	DomainName *string `json:"domainName,omitempty"`

	// The IP Addresses and Ports used when connecting to DomainName.
	EndpointDetails []*EndpointDetail `json:"endpointDetails,omitempty"`
}

// EndpointDetail - Current TCP connectivity information from the App Service Environment to a single endpoint.
type EndpointDetail struct {
	// An IP Address that Domain Name currently resolves to.
	IPAddress *string `json:"ipAddress,omitempty"`

	// Whether it is possible to create a TCP connection from the App Service Environment to this IpAddress at this Port.
	IsAccessible *bool `json:"isAccessible,omitempty"`

	// The time in milliseconds it takes for a TCP connection to be created from the App Service Environment to this IpAddress
	// at this Port.
	Latency *float64 `json:"latency,omitempty"`

	// The port an endpoint is connected to.
	Port *int32 `json:"port,omitempty"`
}

// Error - The resource management error response.
type Error struct {
	// RP error response.
	Error *ErrorResponse `json:"error,omitempty"`
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info interface{} `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorResponse - The resource management error response.
type ErrorResponse struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error details.
	Details []*ErrorResponse `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The error target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// IscsiLun - LUN to expose the Azure Managed Disk.
type IscsiLun struct {
	// REQUIRED; Azure Resource ID of the Managed Disk.
	ManagedDiskAzureResourceID *string `json:"managedDiskAzureResourceId,omitempty"`

	// REQUIRED; User defined name for iSCSI LUN; example: "lun0"
	Name *string `json:"name,omitempty"`

	// READ-ONLY; Specifies the Logical Unit Number of the iSCSI LUN.
	Lun *int32 `json:"lun,omitempty" azure:"ro"`
}

// IscsiTarget - Response for iSCSI Target requests.
type IscsiTarget struct {
	// REQUIRED; Properties for iSCSI Target operations.
	Properties *IscsiTargetProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Azure resource id. Indicates if this resource is managed by another Azure resource.
	ManagedBy *string `json:"managedBy,omitempty" azure:"ro"`

	// READ-ONLY; List of Azure resource ids that manage this resource.
	ManagedByExtended []*string `json:"managedByExtended,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource metadata required by ARM RPC
	SystemData *SystemMetadata `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// IscsiTargetCreate - Payload for iSCSI Target create or update requests.
type IscsiTargetCreate struct {
	// REQUIRED; Properties for iSCSI Target create request.
	Properties *IscsiTargetCreateProperties `json:"properties,omitempty"`

	// Azure resource id. Indicates if this resource is managed by another Azure resource.
	ManagedBy *string `json:"managedBy,omitempty"`

	// List of Azure resource ids that manage this resource.
	ManagedByExtended []*string `json:"managedByExtended,omitempty"`

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// IscsiTargetCreateProperties - Properties for iSCSI Target create or update request.
type IscsiTargetCreateProperties struct {
	// REQUIRED; Mode for Target connectivity.
	ACLMode *IscsiTargetACLMode `json:"aclMode,omitempty"`

	// List of LUNs to be exposed through iSCSI Target.
	Luns []*IscsiLun `json:"luns,omitempty"`

	// Access Control List (ACL) for an iSCSI Target; defines LUN masking policy
	StaticACLs []*ACL `json:"staticAcls,omitempty"`

	// iSCSI Target IQN (iSCSI Qualified Name); example: "iqn.2005-03.org.iscsi:server".
	TargetIqn *string `json:"targetIqn,omitempty"`
}

// IscsiTargetList - List of iSCSI Targets.
type IscsiTargetList struct {
	// REQUIRED; An array of iSCSI Targets in a Disk Pool.
	Value []*IscsiTarget `json:"value,omitempty"`

	// READ-ONLY; URI to fetch the next section of the paginated response.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// IscsiTargetProperties - Response properties for iSCSI Target operations.
type IscsiTargetProperties struct {
	// REQUIRED; Mode for Target connectivity.
	ACLMode *IscsiTargetACLMode `json:"aclMode,omitempty"`

	// REQUIRED; Operational status of the iSCSI Target.
	Status *OperationalStatus `json:"status,omitempty"`

	// REQUIRED; iSCSI Target IQN (iSCSI Qualified Name); example: "iqn.2005-03.org.iscsi:server".
	TargetIqn *string `json:"targetIqn,omitempty"`

	// READ-ONLY; State of the operation on the resource.
	ProvisioningState *ProvisioningStates `json:"provisioningState,omitempty" azure:"ro"`

	// List of private IPv4 addresses to connect to the iSCSI Target.
	Endpoints []*string `json:"endpoints,omitempty"`

	// List of LUNs to be exposed through iSCSI Target.
	Luns []*IscsiLun `json:"luns,omitempty"`

	// The port used by iSCSI Target portal group.
	Port *int32 `json:"port,omitempty"`

	// Access Control List (ACL) for an iSCSI Target; defines LUN masking policy
	StaticACLs []*ACL `json:"staticAcls,omitempty"`

	// READ-ONLY; List of identifiers for active sessions on the iSCSI target
	Sessions []*string `json:"sessions,omitempty" azure:"ro"`
}

// IscsiTargetUpdate - Payload for iSCSI Target update requests.
type IscsiTargetUpdate struct {
	// REQUIRED; Properties for iSCSI Target update request.
	Properties *IscsiTargetUpdateProperties `json:"properties,omitempty"`

	// Azure resource id. Indicates if this resource is managed by another Azure resource.
	ManagedBy *string `json:"managedBy,omitempty"`

	// List of Azure resource ids that manage this resource.
	ManagedByExtended []*string `json:"managedByExtended,omitempty"`

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// IscsiTargetUpdateProperties - Properties for iSCSI Target update request.
type IscsiTargetUpdateProperties struct {
	// List of LUNs to be exposed through iSCSI Target.
	Luns []*IscsiLun `json:"luns,omitempty"`

	// Access Control List (ACL) for an iSCSI Target; defines LUN masking policy
	StaticACLs []*ACL `json:"staticAcls,omitempty"`
}

// IscsiTargetsClientBeginCreateOrUpdateOptions contains the optional parameters for the IscsiTargetsClient.BeginCreateOrUpdate
// method.
type IscsiTargetsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// IscsiTargetsClientBeginDeleteOptions contains the optional parameters for the IscsiTargetsClient.BeginDelete method.
type IscsiTargetsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// IscsiTargetsClientBeginUpdateOptions contains the optional parameters for the IscsiTargetsClient.BeginUpdate method.
type IscsiTargetsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// IscsiTargetsClientGetOptions contains the optional parameters for the IscsiTargetsClient.Get method.
type IscsiTargetsClientGetOptions struct {
	// placeholder for future optional parameters
}

// IscsiTargetsClientListByDiskPoolOptions contains the optional parameters for the IscsiTargetsClient.ListByDiskPool method.
type IscsiTargetsClientListByDiskPoolOptions struct {
	// placeholder for future optional parameters
}

// OperationDisplay - Metadata about an operation.
type OperationDisplay struct {
	// REQUIRED; Localized friendly description for the operation, as it should be shown to the user.
	Description *string `json:"description,omitempty"`

	// REQUIRED; Localized friendly name for the operation, as it should be shown to the user.
	Operation *string `json:"operation,omitempty"`

	// REQUIRED; Localized friendly form of the resource provider name.
	Provider *string `json:"provider,omitempty"`

	// REQUIRED; Localized friendly form of the resource type related to this action/operation.
	Resource *string `json:"resource,omitempty"`
}

// OperationListResult - List of operations supported by the RP.
type OperationListResult struct {
	// REQUIRED; An array of operations supported by the StoragePool RP.
	Value []*RPOperation `json:"value,omitempty"`

	// URI to fetch the next section of the paginated response.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// OutboundEnvironmentEndpoint - Endpoints accessed for a common purpose that the App Service Environment requires outbound
// network access to.
type OutboundEnvironmentEndpoint struct {
	// The type of service accessed by the App Service Environment, e.g., Azure Storage, Azure SQL Database, and Azure Active
	// Directory.
	Category *string `json:"category,omitempty"`

	// The endpoints that the App Service Environment reaches the service at.
	Endpoints []*EndpointDependency `json:"endpoints,omitempty"`
}

// OutboundEnvironmentEndpointList - Collection of Outbound Environment Endpoints
type OutboundEnvironmentEndpointList struct {
	// REQUIRED; Collection of resources.
	Value []*OutboundEnvironmentEndpoint `json:"value,omitempty"`

	// READ-ONLY; Link to next page of resources.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// ProxyResource - The resource model definition for a ARM proxy resource. It will have everything other than required location
// and tags
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// RPOperation - Description of a StoragePool RP Operation
type RPOperation struct {
	// REQUIRED; Additional metadata about RP operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// REQUIRED; Indicates whether the operation applies to data-plane.
	IsDataAction *bool `json:"isDataAction,omitempty"`

	// REQUIRED; The name of the operation being performed on this particular object
	Name *string `json:"name,omitempty"`

	// Indicates the action type.
	ActionType *string `json:"actionType,omitempty"`

	// The intended executor of the operation; governs the display of the operation in the RBAC UX and the audit logs UX.
	Origin *string `json:"origin,omitempty"`
}

// Resource - ARM resource model definition.
type Resource struct {
	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ResourceSKUCapability - Capability a resource SKU has.
type ResourceSKUCapability struct {
	// READ-ONLY; Capability name
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Capability value
	Value *string `json:"value,omitempty" azure:"ro"`
}

// ResourceSKUInfo - Resource SKU Details
type ResourceSKUInfo struct {
	// READ-ONLY; StoragePool RP API version
	APIVersion *string `json:"apiVersion,omitempty" azure:"ro"`

	// READ-ONLY; List of additional capabilities for StoragePool resource.
	Capabilities []*ResourceSKUCapability `json:"capabilities,omitempty" azure:"ro"`

	// READ-ONLY; Zones and zone capabilities in those locations where the SKU is available.
	LocationInfo *ResourceSKULocationInfo `json:"locationInfo,omitempty" azure:"ro"`

	// READ-ONLY; Sku name
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; StoragePool resource type
	ResourceType *string `json:"resourceType,omitempty" azure:"ro"`

	// READ-ONLY; The restrictions because of which SKU cannot be used. This is empty if there are no restrictions.
	Restrictions []*ResourceSKURestrictions `json:"restrictions,omitempty" azure:"ro"`

	// READ-ONLY; Sku tier
	Tier *string `json:"tier,omitempty" azure:"ro"`
}

// ResourceSKUListResult - List Disk Pool skus operation response.
type ResourceSKUListResult struct {
	// URI to fetch the next section of the paginated response.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of StoragePool resource skus.
	Value []*ResourceSKUInfo `json:"value,omitempty"`
}

// ResourceSKULocationInfo - Zone and capability info for resource sku
type ResourceSKULocationInfo struct {
	// READ-ONLY; Location of the SKU
	Location *string `json:"location,omitempty" azure:"ro"`

	// READ-ONLY; Details of capabilities available to a SKU in specific zones.
	ZoneDetails []*ResourceSKUZoneDetails `json:"zoneDetails,omitempty" azure:"ro"`

	// READ-ONLY; List of availability zones where the SKU is supported.
	Zones []*string `json:"zones,omitempty" azure:"ro"`
}

// ResourceSKURestrictionInfo - Describes an available Compute SKU Restriction Information.
type ResourceSKURestrictionInfo struct {
	// READ-ONLY; Locations where the SKU is restricted
	Locations []*string `json:"locations,omitempty" azure:"ro"`

	// READ-ONLY; List of availability zones where the SKU is restricted.
	Zones []*string `json:"zones,omitempty" azure:"ro"`
}

// ResourceSKURestrictions - Describes scaling information of a SKU.
type ResourceSKURestrictions struct {
	// READ-ONLY; The reason for restriction.
	ReasonCode *ResourceSKURestrictionsReasonCode `json:"reasonCode,omitempty" azure:"ro"`

	// READ-ONLY; The information about the restriction where the SKU cannot be used.
	RestrictionInfo *ResourceSKURestrictionInfo `json:"restrictionInfo,omitempty" azure:"ro"`

	// READ-ONLY; The type of restrictions.
	Type *ResourceSKURestrictionsType `json:"type,omitempty" azure:"ro"`

	// READ-ONLY; The value of restrictions. If the restriction type is set to location. This would be different locations where
	// the SKU is restricted.
	Values []*string `json:"values,omitempty" azure:"ro"`
}

// ResourceSKUZoneDetails - Describes The zonal capabilities of a SKU.
type ResourceSKUZoneDetails struct {
	// READ-ONLY; A list of capabilities that are available for the SKU in the specified list of zones.
	Capabilities []*ResourceSKUCapability `json:"capabilities,omitempty" azure:"ro"`

	// READ-ONLY; The set of zones that the SKU is available in with the specified capabilities.
	Name []*string `json:"name,omitempty" azure:"ro"`
}

// ResourceSKUsClientListOptions contains the optional parameters for the ResourceSKUsClient.List method.
type ResourceSKUsClientListOptions struct {
	// placeholder for future optional parameters
}

// SKU - Sku for ARM resource
type SKU struct {
	// REQUIRED; Sku name
	Name *string `json:"name,omitempty"`

	// Sku tier
	Tier *string `json:"tier,omitempty"`
}

// SystemMetadata - Metadata pertaining to creation and last modification of the resource.
type SystemMetadata struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// TrackedResource - The resource model definition for a ARM tracked top level resource.
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives.
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty" azure:"ro"`
}
