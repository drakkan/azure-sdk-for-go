package storageapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-09-01/storage"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result storage.OperationListResult, err error)
}

var _ OperationsClientAPI = (*storage.OperationsClient)(nil)

// SkusClientAPI contains the set of methods on the SkusClient type.
type SkusClientAPI interface {
	List(ctx context.Context) (result storage.SkuListResult, err error)
}

var _ SkusClientAPI = (*storage.SkusClient)(nil)

// AccountsClientAPI contains the set of methods on the AccountsClient type.
type AccountsClientAPI interface {
	AbortHierarchicalNamespaceMigration(ctx context.Context, resourceGroupName string, accountName string) (result storage.AccountsAbortHierarchicalNamespaceMigrationFuture, err error)
	CheckNameAvailability(ctx context.Context, accountName storage.AccountCheckNameAvailabilityParameters) (result storage.CheckNameAvailabilityResult, err error)
	Create(ctx context.Context, resourceGroupName string, accountName string, parameters storage.AccountCreateParameters) (result storage.AccountsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error)
	Failover(ctx context.Context, resourceGroupName string, accountName string) (result storage.AccountsFailoverFuture, err error)
	GetProperties(ctx context.Context, resourceGroupName string, accountName string, expand storage.AccountExpand) (result storage.Account, err error)
	HierarchicalNamespaceMigration(ctx context.Context, resourceGroupName string, accountName string, requestType string) (result storage.AccountsHierarchicalNamespaceMigrationFuture, err error)
	List(ctx context.Context) (result storage.AccountListResultPage, err error)
	ListComplete(ctx context.Context) (result storage.AccountListResultIterator, err error)
	ListAccountSAS(ctx context.Context, resourceGroupName string, accountName string, parameters storage.AccountSasParameters) (result storage.ListAccountSasResponse, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result storage.AccountListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result storage.AccountListResultIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, accountName string, expand storage.ListKeyExpand) (result storage.AccountListKeysResult, err error)
	ListServiceSAS(ctx context.Context, resourceGroupName string, accountName string, parameters storage.ServiceSasParameters) (result storage.ListServiceSasResponse, err error)
	RegenerateKey(ctx context.Context, resourceGroupName string, accountName string, regenerateKey storage.AccountRegenerateKeyParameters) (result storage.AccountListKeysResult, err error)
	RestoreBlobRanges(ctx context.Context, resourceGroupName string, accountName string, parameters storage.BlobRestoreParameters) (result storage.AccountsRestoreBlobRangesFuture, err error)
	RevokeUserDelegationKeys(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, parameters storage.AccountUpdateParameters) (result storage.Account, err error)
}

var _ AccountsClientAPI = (*storage.AccountsClient)(nil)

// DeletedAccountsClientAPI contains the set of methods on the DeletedAccountsClient type.
type DeletedAccountsClientAPI interface {
	Get(ctx context.Context, deletedAccountName string, location string) (result storage.DeletedAccount, err error)
	List(ctx context.Context) (result storage.DeletedAccountListResultPage, err error)
	ListComplete(ctx context.Context) (result storage.DeletedAccountListResultIterator, err error)
}

var _ DeletedAccountsClientAPI = (*storage.DeletedAccountsClient)(nil)

// UsagesClientAPI contains the set of methods on the UsagesClient type.
type UsagesClientAPI interface {
	ListByLocation(ctx context.Context, location string) (result storage.UsageListResult, err error)
}

var _ UsagesClientAPI = (*storage.UsagesClient)(nil)

// ManagementPoliciesClientAPI contains the set of methods on the ManagementPoliciesClient type.
type ManagementPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, properties storage.ManagementPolicy) (result storage.ManagementPolicy, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result storage.ManagementPolicy, err error)
}

var _ ManagementPoliciesClientAPI = (*storage.ManagementPoliciesClient)(nil)

// BlobInventoryPoliciesClientAPI contains the set of methods on the BlobInventoryPoliciesClient type.
type BlobInventoryPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, properties storage.BlobInventoryPolicy) (result storage.BlobInventoryPolicy, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result storage.BlobInventoryPolicy, err error)
	List(ctx context.Context, resourceGroupName string, accountName string) (result storage.ListBlobInventoryPolicy, err error)
}

var _ BlobInventoryPoliciesClientAPI = (*storage.BlobInventoryPoliciesClient)(nil)

// PrivateEndpointConnectionsClientAPI contains the set of methods on the PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string) (result storage.PrivateEndpointConnection, err error)
	List(ctx context.Context, resourceGroupName string, accountName string) (result storage.PrivateEndpointConnectionListResult, err error)
	Put(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string, properties storage.PrivateEndpointConnection) (result storage.PrivateEndpointConnection, err error)
}

var _ PrivateEndpointConnectionsClientAPI = (*storage.PrivateEndpointConnectionsClient)(nil)

// PrivateLinkResourcesClientAPI contains the set of methods on the PrivateLinkResourcesClient type.
type PrivateLinkResourcesClientAPI interface {
	ListByStorageAccount(ctx context.Context, resourceGroupName string, accountName string) (result storage.PrivateLinkResourceListResult, err error)
}

var _ PrivateLinkResourcesClientAPI = (*storage.PrivateLinkResourcesClient)(nil)

// ObjectReplicationPoliciesClientAPI contains the set of methods on the ObjectReplicationPoliciesClient type.
type ObjectReplicationPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, objectReplicationPolicyID string, properties storage.ObjectReplicationPolicy) (result storage.ObjectReplicationPolicy, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, objectReplicationPolicyID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, objectReplicationPolicyID string) (result storage.ObjectReplicationPolicy, err error)
	List(ctx context.Context, resourceGroupName string, accountName string) (result storage.ObjectReplicationPolicies, err error)
}

var _ ObjectReplicationPoliciesClientAPI = (*storage.ObjectReplicationPoliciesClient)(nil)

// LocalUsersClientAPI contains the set of methods on the LocalUsersClient type.
type LocalUsersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, username string, properties storage.LocalUser) (result storage.LocalUser, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, username string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, username string) (result storage.LocalUser, err error)
	List(ctx context.Context, resourceGroupName string, accountName string) (result storage.LocalUsers, err error)
	ListKeys(ctx context.Context, resourceGroupName string, accountName string, username string) (result storage.LocalUserKeys, err error)
	RegeneratePassword(ctx context.Context, resourceGroupName string, accountName string, username string) (result storage.LocalUserRegeneratePasswordResult, err error)
}

var _ LocalUsersClientAPI = (*storage.LocalUsersClient)(nil)

// EncryptionScopesClientAPI contains the set of methods on the EncryptionScopesClient type.
type EncryptionScopesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, accountName string, encryptionScopeName string) (result storage.EncryptionScope, err error)
	List(ctx context.Context, resourceGroupName string, accountName string) (result storage.EncryptionScopeListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, accountName string) (result storage.EncryptionScopeListResultIterator, err error)
	Patch(ctx context.Context, resourceGroupName string, accountName string, encryptionScopeName string, encryptionScope storage.EncryptionScope) (result storage.EncryptionScope, err error)
	Put(ctx context.Context, resourceGroupName string, accountName string, encryptionScopeName string, encryptionScope storage.EncryptionScope) (result storage.EncryptionScope, err error)
}

var _ EncryptionScopesClientAPI = (*storage.EncryptionScopesClient)(nil)

// BlobServicesClientAPI contains the set of methods on the BlobServicesClient type.
type BlobServicesClientAPI interface {
	GetServiceProperties(ctx context.Context, resourceGroupName string, accountName string) (result storage.BlobServiceProperties, err error)
	List(ctx context.Context, resourceGroupName string, accountName string) (result storage.BlobServiceItems, err error)
	SetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, parameters storage.BlobServiceProperties) (result storage.BlobServiceProperties, err error)
}

var _ BlobServicesClientAPI = (*storage.BlobServicesClient)(nil)

// BlobContainersClientAPI contains the set of methods on the BlobContainersClient type.
type BlobContainersClientAPI interface {
	ClearLegalHold(ctx context.Context, resourceGroupName string, accountName string, containerName string, legalHold storage.LegalHold) (result storage.LegalHold, err error)
	Create(ctx context.Context, resourceGroupName string, accountName string, containerName string, blobContainer storage.BlobContainer) (result storage.BlobContainer, err error)
	CreateOrUpdateImmutabilityPolicy(ctx context.Context, resourceGroupName string, accountName string, containerName string, parameters *storage.ImmutabilityPolicy, ifMatch string) (result storage.ImmutabilityPolicy, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, containerName string) (result autorest.Response, err error)
	DeleteImmutabilityPolicy(ctx context.Context, resourceGroupName string, accountName string, containerName string, ifMatch string) (result storage.ImmutabilityPolicy, err error)
	ExtendImmutabilityPolicy(ctx context.Context, resourceGroupName string, accountName string, containerName string, ifMatch string, parameters *storage.ImmutabilityPolicy) (result storage.ImmutabilityPolicy, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, containerName string) (result storage.BlobContainer, err error)
	GetImmutabilityPolicy(ctx context.Context, resourceGroupName string, accountName string, containerName string, ifMatch string) (result storage.ImmutabilityPolicy, err error)
	Lease(ctx context.Context, resourceGroupName string, accountName string, containerName string, parameters *storage.LeaseContainerRequest) (result storage.LeaseContainerResponse, err error)
	List(ctx context.Context, resourceGroupName string, accountName string, maxpagesize string, filter string, include storage.ListContainersInclude) (result storage.ListContainerItemsPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, accountName string, maxpagesize string, filter string, include storage.ListContainersInclude) (result storage.ListContainerItemsIterator, err error)
	LockImmutabilityPolicy(ctx context.Context, resourceGroupName string, accountName string, containerName string, ifMatch string) (result storage.ImmutabilityPolicy, err error)
	ObjectLevelWorm(ctx context.Context, resourceGroupName string, accountName string, containerName string) (result storage.BlobContainersObjectLevelWormFuture, err error)
	SetLegalHold(ctx context.Context, resourceGroupName string, accountName string, containerName string, legalHold storage.LegalHold) (result storage.LegalHold, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, containerName string, blobContainer storage.BlobContainer) (result storage.BlobContainer, err error)
}

var _ BlobContainersClientAPI = (*storage.BlobContainersClient)(nil)

// FileServicesClientAPI contains the set of methods on the FileServicesClient type.
type FileServicesClientAPI interface {
	GetServiceProperties(ctx context.Context, resourceGroupName string, accountName string) (result storage.FileServiceProperties, err error)
	List(ctx context.Context, resourceGroupName string, accountName string) (result storage.FileServiceItems, err error)
	SetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, parameters storage.FileServiceProperties) (result storage.FileServiceProperties, err error)
}

var _ FileServicesClientAPI = (*storage.FileServicesClient)(nil)

// FileSharesClientAPI contains the set of methods on the FileSharesClient type.
type FileSharesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, shareName string, fileShare storage.FileShare, expand string) (result storage.FileShare, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, shareName string, xMsSnapshot string, include string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, shareName string, expand string, xMsSnapshot string) (result storage.FileShare, err error)
	Lease(ctx context.Context, resourceGroupName string, accountName string, shareName string, parameters *storage.LeaseShareRequest, xMsSnapshot string) (result storage.LeaseShareResponse, err error)
	List(ctx context.Context, resourceGroupName string, accountName string, maxpagesize string, filter string, expand string) (result storage.FileShareItemsPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, accountName string, maxpagesize string, filter string, expand string) (result storage.FileShareItemsIterator, err error)
	Restore(ctx context.Context, resourceGroupName string, accountName string, shareName string, deletedShare storage.DeletedShare) (result autorest.Response, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, shareName string, fileShare storage.FileShare) (result storage.FileShare, err error)
}

var _ FileSharesClientAPI = (*storage.FileSharesClient)(nil)

// QueueServicesClientAPI contains the set of methods on the QueueServicesClient type.
type QueueServicesClientAPI interface {
	GetServiceProperties(ctx context.Context, resourceGroupName string, accountName string) (result storage.QueueServiceProperties, err error)
	List(ctx context.Context, resourceGroupName string, accountName string) (result storage.ListQueueServices, err error)
	SetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, parameters storage.QueueServiceProperties) (result storage.QueueServiceProperties, err error)
}

var _ QueueServicesClientAPI = (*storage.QueueServicesClient)(nil)

// QueueClientAPI contains the set of methods on the QueueClient type.
type QueueClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, queueName string, queue storage.Queue) (result storage.Queue, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, queueName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, queueName string) (result storage.Queue, err error)
	List(ctx context.Context, resourceGroupName string, accountName string, maxpagesize string, filter string) (result storage.ListQueueResourcePage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, accountName string, maxpagesize string, filter string) (result storage.ListQueueResourceIterator, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, queueName string, queue storage.Queue) (result storage.Queue, err error)
}

var _ QueueClientAPI = (*storage.QueueClient)(nil)

// TableServicesClientAPI contains the set of methods on the TableServicesClient type.
type TableServicesClientAPI interface {
	GetServiceProperties(ctx context.Context, resourceGroupName string, accountName string) (result storage.TableServiceProperties, err error)
	List(ctx context.Context, resourceGroupName string, accountName string) (result storage.ListTableServices, err error)
	SetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, parameters storage.TableServiceProperties) (result storage.TableServiceProperties, err error)
}

var _ TableServicesClientAPI = (*storage.TableServicesClient)(nil)

// TableClientAPI contains the set of methods on the TableClient type.
type TableClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, tableName string, parameters *storage.Table) (result storage.Table, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, tableName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, tableName string) (result storage.Table, err error)
	List(ctx context.Context, resourceGroupName string, accountName string) (result storage.ListTableResourcePage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, accountName string) (result storage.ListTableResourceIterator, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, tableName string, parameters *storage.Table) (result storage.Table, err error)
}

var _ TableClientAPI = (*storage.TableClient)(nil)
