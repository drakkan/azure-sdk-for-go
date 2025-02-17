package documentdbapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2021-06-15/documentdb"
	"github.com/Azure/go-autorest/autorest"
)

// DatabaseAccountsClientAPI contains the set of methods on the DatabaseAccountsClient type.
type DatabaseAccountsClientAPI interface {
	CheckNameExists(ctx context.Context, accountName string) (result autorest.Response, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, createUpdateParameters documentdb.DatabaseAccountCreateUpdateParameters) (result documentdb.DatabaseAccountsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.DatabaseAccountsDeleteFuture, err error)
	FailoverPriorityChange(ctx context.Context, resourceGroupName string, accountName string, failoverParameters documentdb.FailoverPolicies) (result documentdb.DatabaseAccountsFailoverPriorityChangeFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.DatabaseAccountGetResults, err error)
	GetReadOnlyKeys(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.DatabaseAccountListReadOnlyKeysResult, err error)
	List(ctx context.Context) (result documentdb.DatabaseAccountsListResult, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result documentdb.DatabaseAccountsListResult, err error)
	ListConnectionStrings(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.DatabaseAccountListConnectionStringsResult, err error)
	ListKeys(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.DatabaseAccountListKeysResult, err error)
	ListMetricDefinitions(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.MetricDefinitionsListResult, err error)
	ListMetrics(ctx context.Context, resourceGroupName string, accountName string, filter string) (result documentdb.MetricListResult, err error)
	ListReadOnlyKeys(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.DatabaseAccountListReadOnlyKeysResult, err error)
	ListUsages(ctx context.Context, resourceGroupName string, accountName string, filter string) (result documentdb.UsagesResult, err error)
	OfflineRegion(ctx context.Context, resourceGroupName string, accountName string, regionParameterForOffline documentdb.RegionForOnlineOffline) (result documentdb.DatabaseAccountsOfflineRegionFuture, err error)
	OnlineRegion(ctx context.Context, resourceGroupName string, accountName string, regionParameterForOnline documentdb.RegionForOnlineOffline) (result documentdb.DatabaseAccountsOnlineRegionFuture, err error)
	RegenerateKey(ctx context.Context, resourceGroupName string, accountName string, keyToRegenerate documentdb.DatabaseAccountRegenerateKeyParameters) (result documentdb.DatabaseAccountsRegenerateKeyFuture, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, updateParameters documentdb.DatabaseAccountUpdateParameters) (result documentdb.DatabaseAccountsUpdateFuture, err error)
}

var _ DatabaseAccountsClientAPI = (*documentdb.DatabaseAccountsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result documentdb.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result documentdb.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*documentdb.OperationsClient)(nil)

// DatabaseClientAPI contains the set of methods on the DatabaseClient type.
type DatabaseClientAPI interface {
	ListMetricDefinitions(ctx context.Context, resourceGroupName string, accountName string, databaseRid string) (result documentdb.MetricDefinitionsListResult, err error)
	ListMetrics(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, filter string) (result documentdb.MetricListResult, err error)
	ListUsages(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, filter string) (result documentdb.UsagesResult, err error)
}

var _ DatabaseClientAPI = (*documentdb.DatabaseClient)(nil)

// CollectionClientAPI contains the set of methods on the CollectionClient type.
type CollectionClientAPI interface {
	ListMetricDefinitions(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string) (result documentdb.MetricDefinitionsListResult, err error)
	ListMetrics(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string) (result documentdb.MetricListResult, err error)
	ListUsages(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string) (result documentdb.UsagesResult, err error)
}

var _ CollectionClientAPI = (*documentdb.CollectionClient)(nil)

// CollectionRegionClientAPI contains the set of methods on the CollectionRegionClient type.
type CollectionRegionClientAPI interface {
	ListMetrics(ctx context.Context, resourceGroupName string, accountName string, region string, databaseRid string, collectionRid string, filter string) (result documentdb.MetricListResult, err error)
}

var _ CollectionRegionClientAPI = (*documentdb.CollectionRegionClient)(nil)

// DatabaseAccountRegionClientAPI contains the set of methods on the DatabaseAccountRegionClient type.
type DatabaseAccountRegionClientAPI interface {
	ListMetrics(ctx context.Context, resourceGroupName string, accountName string, region string, filter string) (result documentdb.MetricListResult, err error)
}

var _ DatabaseAccountRegionClientAPI = (*documentdb.DatabaseAccountRegionClient)(nil)

// PercentileSourceTargetClientAPI contains the set of methods on the PercentileSourceTargetClient type.
type PercentileSourceTargetClientAPI interface {
	ListMetrics(ctx context.Context, resourceGroupName string, accountName string, sourceRegion string, targetRegion string, filter string) (result documentdb.PercentileMetricListResult, err error)
}

var _ PercentileSourceTargetClientAPI = (*documentdb.PercentileSourceTargetClient)(nil)

// PercentileTargetClientAPI contains the set of methods on the PercentileTargetClient type.
type PercentileTargetClientAPI interface {
	ListMetrics(ctx context.Context, resourceGroupName string, accountName string, targetRegion string, filter string) (result documentdb.PercentileMetricListResult, err error)
}

var _ PercentileTargetClientAPI = (*documentdb.PercentileTargetClient)(nil)

// PercentileClientAPI contains the set of methods on the PercentileClient type.
type PercentileClientAPI interface {
	ListMetrics(ctx context.Context, resourceGroupName string, accountName string, filter string) (result documentdb.PercentileMetricListResult, err error)
}

var _ PercentileClientAPI = (*documentdb.PercentileClient)(nil)

// CollectionPartitionRegionClientAPI contains the set of methods on the CollectionPartitionRegionClient type.
type CollectionPartitionRegionClientAPI interface {
	ListMetrics(ctx context.Context, resourceGroupName string, accountName string, region string, databaseRid string, collectionRid string, filter string) (result documentdb.PartitionMetricListResult, err error)
}

var _ CollectionPartitionRegionClientAPI = (*documentdb.CollectionPartitionRegionClient)(nil)

// CollectionPartitionClientAPI contains the set of methods on the CollectionPartitionClient type.
type CollectionPartitionClientAPI interface {
	ListMetrics(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string) (result documentdb.PartitionMetricListResult, err error)
	ListUsages(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string) (result documentdb.PartitionUsagesResult, err error)
}

var _ CollectionPartitionClientAPI = (*documentdb.CollectionPartitionClient)(nil)

// PartitionKeyRangeIDClientAPI contains the set of methods on the PartitionKeyRangeIDClient type.
type PartitionKeyRangeIDClientAPI interface {
	ListMetrics(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, partitionKeyRangeID string, filter string) (result documentdb.PartitionMetricListResult, err error)
}

var _ PartitionKeyRangeIDClientAPI = (*documentdb.PartitionKeyRangeIDClient)(nil)

// PartitionKeyRangeIDRegionClientAPI contains the set of methods on the PartitionKeyRangeIDRegionClient type.
type PartitionKeyRangeIDRegionClientAPI interface {
	ListMetrics(ctx context.Context, resourceGroupName string, accountName string, region string, databaseRid string, collectionRid string, partitionKeyRangeID string, filter string) (result documentdb.PartitionMetricListResult, err error)
}

var _ PartitionKeyRangeIDRegionClientAPI = (*documentdb.PartitionKeyRangeIDRegionClient)(nil)

// SQLResourcesClientAPI contains the set of methods on the SQLResourcesClient type.
type SQLResourcesClientAPI interface {
	CreateUpdateSQLContainer(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string, createUpdateSQLContainerParameters documentdb.SQLContainerCreateUpdateParameters) (result documentdb.SQLResourcesCreateUpdateSQLContainerFuture, err error)
	CreateUpdateSQLDatabase(ctx context.Context, resourceGroupName string, accountName string, databaseName string, createUpdateSQLDatabaseParameters documentdb.SQLDatabaseCreateUpdateParameters) (result documentdb.SQLResourcesCreateUpdateSQLDatabaseFuture, err error)
	CreateUpdateSQLRoleAssignment(ctx context.Context, roleAssignmentID string, resourceGroupName string, accountName string, createUpdateSQLRoleAssignmentParameters documentdb.SQLRoleAssignmentCreateUpdateParameters) (result documentdb.SQLResourcesCreateUpdateSQLRoleAssignmentFuture, err error)
	CreateUpdateSQLRoleDefinition(ctx context.Context, roleDefinitionID string, resourceGroupName string, accountName string, createUpdateSQLRoleDefinitionParameters documentdb.SQLRoleDefinitionCreateUpdateParameters) (result documentdb.SQLResourcesCreateUpdateSQLRoleDefinitionFuture, err error)
	CreateUpdateSQLStoredProcedure(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string, storedProcedureName string, createUpdateSQLStoredProcedureParameters documentdb.SQLStoredProcedureCreateUpdateParameters) (result documentdb.SQLResourcesCreateUpdateSQLStoredProcedureFuture, err error)
	CreateUpdateSQLTrigger(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string, triggerName string, createUpdateSQLTriggerParameters documentdb.SQLTriggerCreateUpdateParameters) (result documentdb.SQLResourcesCreateUpdateSQLTriggerFuture, err error)
	CreateUpdateSQLUserDefinedFunction(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string, userDefinedFunctionName string, createUpdateSQLUserDefinedFunctionParameters documentdb.SQLUserDefinedFunctionCreateUpdateParameters) (result documentdb.SQLResourcesCreateUpdateSQLUserDefinedFunctionFuture, err error)
	DeleteSQLContainer(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string) (result documentdb.SQLResourcesDeleteSQLContainerFuture, err error)
	DeleteSQLDatabase(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.SQLResourcesDeleteSQLDatabaseFuture, err error)
	DeleteSQLRoleAssignment(ctx context.Context, roleAssignmentID string, resourceGroupName string, accountName string) (result documentdb.SQLResourcesDeleteSQLRoleAssignmentFuture, err error)
	DeleteSQLRoleDefinition(ctx context.Context, roleDefinitionID string, resourceGroupName string, accountName string) (result documentdb.SQLResourcesDeleteSQLRoleDefinitionFuture, err error)
	DeleteSQLStoredProcedure(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string, storedProcedureName string) (result documentdb.SQLResourcesDeleteSQLStoredProcedureFuture, err error)
	DeleteSQLTrigger(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string, triggerName string) (result documentdb.SQLResourcesDeleteSQLTriggerFuture, err error)
	DeleteSQLUserDefinedFunction(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string, userDefinedFunctionName string) (result documentdb.SQLResourcesDeleteSQLUserDefinedFunctionFuture, err error)
	GetSQLContainer(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string) (result documentdb.SQLContainerGetResults, err error)
	GetSQLContainerThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string) (result documentdb.ThroughputSettingsGetResults, err error)
	GetSQLDatabase(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.SQLDatabaseGetResults, err error)
	GetSQLDatabaseThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.ThroughputSettingsGetResults, err error)
	GetSQLRoleAssignment(ctx context.Context, roleAssignmentID string, resourceGroupName string, accountName string) (result documentdb.SQLRoleAssignmentGetResults, err error)
	GetSQLRoleDefinition(ctx context.Context, roleDefinitionID string, resourceGroupName string, accountName string) (result documentdb.SQLRoleDefinitionGetResults, err error)
	GetSQLStoredProcedure(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string, storedProcedureName string) (result documentdb.SQLStoredProcedureGetResults, err error)
	GetSQLTrigger(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string, triggerName string) (result documentdb.SQLTriggerGetResults, err error)
	GetSQLUserDefinedFunction(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string, userDefinedFunctionName string) (result documentdb.SQLUserDefinedFunctionGetResults, err error)
	ListSQLContainers(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.SQLContainerListResult, err error)
	ListSQLDatabases(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.SQLDatabaseListResult, err error)
	ListSQLRoleAssignments(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.SQLRoleAssignmentListResult, err error)
	ListSQLRoleDefinitions(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.SQLRoleDefinitionListResult, err error)
	ListSQLStoredProcedures(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string) (result documentdb.SQLStoredProcedureListResult, err error)
	ListSQLTriggers(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string) (result documentdb.SQLTriggerListResult, err error)
	ListSQLUserDefinedFunctions(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string) (result documentdb.SQLUserDefinedFunctionListResult, err error)
	MigrateSQLContainerToAutoscale(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string) (result documentdb.SQLResourcesMigrateSQLContainerToAutoscaleFuture, err error)
	MigrateSQLContainerToManualThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string) (result documentdb.SQLResourcesMigrateSQLContainerToManualThroughputFuture, err error)
	MigrateSQLDatabaseToAutoscale(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.SQLResourcesMigrateSQLDatabaseToAutoscaleFuture, err error)
	MigrateSQLDatabaseToManualThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.SQLResourcesMigrateSQLDatabaseToManualThroughputFuture, err error)
	RetrieveContinuousBackupInformation(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string, location documentdb.ContinuousBackupRestoreLocation) (result documentdb.SQLResourcesRetrieveContinuousBackupInformationFuture, err error)
	UpdateSQLContainerThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string, updateThroughputParameters documentdb.ThroughputSettingsUpdateParameters) (result documentdb.SQLResourcesUpdateSQLContainerThroughputFuture, err error)
	UpdateSQLDatabaseThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string, updateThroughputParameters documentdb.ThroughputSettingsUpdateParameters) (result documentdb.SQLResourcesUpdateSQLDatabaseThroughputFuture, err error)
}

var _ SQLResourcesClientAPI = (*documentdb.SQLResourcesClient)(nil)

// MongoDBResourcesClientAPI contains the set of methods on the MongoDBResourcesClient type.
type MongoDBResourcesClientAPI interface {
	CreateUpdateMongoDBCollection(ctx context.Context, resourceGroupName string, accountName string, databaseName string, collectionName string, createUpdateMongoDBCollectionParameters documentdb.MongoDBCollectionCreateUpdateParameters) (result documentdb.MongoDBResourcesCreateUpdateMongoDBCollectionFuture, err error)
	CreateUpdateMongoDBDatabase(ctx context.Context, resourceGroupName string, accountName string, databaseName string, createUpdateMongoDBDatabaseParameters documentdb.MongoDBDatabaseCreateUpdateParameters) (result documentdb.MongoDBResourcesCreateUpdateMongoDBDatabaseFuture, err error)
	DeleteMongoDBCollection(ctx context.Context, resourceGroupName string, accountName string, databaseName string, collectionName string) (result documentdb.MongoDBResourcesDeleteMongoDBCollectionFuture, err error)
	DeleteMongoDBDatabase(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.MongoDBResourcesDeleteMongoDBDatabaseFuture, err error)
	GetMongoDBCollection(ctx context.Context, resourceGroupName string, accountName string, databaseName string, collectionName string) (result documentdb.MongoDBCollectionGetResults, err error)
	GetMongoDBCollectionThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string, collectionName string) (result documentdb.ThroughputSettingsGetResults, err error)
	GetMongoDBDatabase(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.MongoDBDatabaseGetResults, err error)
	GetMongoDBDatabaseThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.ThroughputSettingsGetResults, err error)
	ListMongoDBCollections(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.MongoDBCollectionListResult, err error)
	ListMongoDBDatabases(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.MongoDBDatabaseListResult, err error)
	MigrateMongoDBCollectionToAutoscale(ctx context.Context, resourceGroupName string, accountName string, databaseName string, collectionName string) (result documentdb.MongoDBResourcesMigrateMongoDBCollectionToAutoscaleFuture, err error)
	MigrateMongoDBCollectionToManualThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string, collectionName string) (result documentdb.MongoDBResourcesMigrateMongoDBCollectionToManualThroughputFuture, err error)
	MigrateMongoDBDatabaseToAutoscale(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.MongoDBResourcesMigrateMongoDBDatabaseToAutoscaleFuture, err error)
	MigrateMongoDBDatabaseToManualThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.MongoDBResourcesMigrateMongoDBDatabaseToManualThroughputFuture, err error)
	UpdateMongoDBCollectionThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string, collectionName string, updateThroughputParameters documentdb.ThroughputSettingsUpdateParameters) (result documentdb.MongoDBResourcesUpdateMongoDBCollectionThroughputFuture, err error)
	UpdateMongoDBDatabaseThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string, updateThroughputParameters documentdb.ThroughputSettingsUpdateParameters) (result documentdb.MongoDBResourcesUpdateMongoDBDatabaseThroughputFuture, err error)
}

var _ MongoDBResourcesClientAPI = (*documentdb.MongoDBResourcesClient)(nil)

// TableResourcesClientAPI contains the set of methods on the TableResourcesClient type.
type TableResourcesClientAPI interface {
	CreateUpdateTable(ctx context.Context, resourceGroupName string, accountName string, tableName string, createUpdateTableParameters documentdb.TableCreateUpdateParameters) (result documentdb.TableResourcesCreateUpdateTableFuture, err error)
	DeleteTable(ctx context.Context, resourceGroupName string, accountName string, tableName string) (result documentdb.TableResourcesDeleteTableFuture, err error)
	GetTable(ctx context.Context, resourceGroupName string, accountName string, tableName string) (result documentdb.TableGetResults, err error)
	GetTableThroughput(ctx context.Context, resourceGroupName string, accountName string, tableName string) (result documentdb.ThroughputSettingsGetResults, err error)
	ListTables(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.TableListResult, err error)
	MigrateTableToAutoscale(ctx context.Context, resourceGroupName string, accountName string, tableName string) (result documentdb.TableResourcesMigrateTableToAutoscaleFuture, err error)
	MigrateTableToManualThroughput(ctx context.Context, resourceGroupName string, accountName string, tableName string) (result documentdb.TableResourcesMigrateTableToManualThroughputFuture, err error)
	UpdateTableThroughput(ctx context.Context, resourceGroupName string, accountName string, tableName string, updateThroughputParameters documentdb.ThroughputSettingsUpdateParameters) (result documentdb.TableResourcesUpdateTableThroughputFuture, err error)
}

var _ TableResourcesClientAPI = (*documentdb.TableResourcesClient)(nil)

// CassandraResourcesClientAPI contains the set of methods on the CassandraResourcesClient type.
type CassandraResourcesClientAPI interface {
	CreateUpdateCassandraKeyspace(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string, createUpdateCassandraKeyspaceParameters documentdb.CassandraKeyspaceCreateUpdateParameters) (result documentdb.CassandraResourcesCreateUpdateCassandraKeyspaceFuture, err error)
	CreateUpdateCassandraTable(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string, tableName string, createUpdateCassandraTableParameters documentdb.CassandraTableCreateUpdateParameters) (result documentdb.CassandraResourcesCreateUpdateCassandraTableFuture, err error)
	DeleteCassandraKeyspace(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string) (result documentdb.CassandraResourcesDeleteCassandraKeyspaceFuture, err error)
	DeleteCassandraTable(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string, tableName string) (result documentdb.CassandraResourcesDeleteCassandraTableFuture, err error)
	GetCassandraKeyspace(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string) (result documentdb.CassandraKeyspaceGetResults, err error)
	GetCassandraKeyspaceThroughput(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string) (result documentdb.ThroughputSettingsGetResults, err error)
	GetCassandraTable(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string, tableName string) (result documentdb.CassandraTableGetResults, err error)
	GetCassandraTableThroughput(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string, tableName string) (result documentdb.ThroughputSettingsGetResults, err error)
	ListCassandraKeyspaces(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.CassandraKeyspaceListResult, err error)
	ListCassandraTables(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string) (result documentdb.CassandraTableListResult, err error)
	MigrateCassandraKeyspaceToAutoscale(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string) (result documentdb.CassandraResourcesMigrateCassandraKeyspaceToAutoscaleFuture, err error)
	MigrateCassandraKeyspaceToManualThroughput(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string) (result documentdb.CassandraResourcesMigrateCassandraKeyspaceToManualThroughputFuture, err error)
	MigrateCassandraTableToAutoscale(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string, tableName string) (result documentdb.CassandraResourcesMigrateCassandraTableToAutoscaleFuture, err error)
	MigrateCassandraTableToManualThroughput(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string, tableName string) (result documentdb.CassandraResourcesMigrateCassandraTableToManualThroughputFuture, err error)
	UpdateCassandraKeyspaceThroughput(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string, updateThroughputParameters documentdb.ThroughputSettingsUpdateParameters) (result documentdb.CassandraResourcesUpdateCassandraKeyspaceThroughputFuture, err error)
	UpdateCassandraTableThroughput(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string, tableName string, updateThroughputParameters documentdb.ThroughputSettingsUpdateParameters) (result documentdb.CassandraResourcesUpdateCassandraTableThroughputFuture, err error)
}

var _ CassandraResourcesClientAPI = (*documentdb.CassandraResourcesClient)(nil)

// GremlinResourcesClientAPI contains the set of methods on the GremlinResourcesClient type.
type GremlinResourcesClientAPI interface {
	CreateUpdateGremlinDatabase(ctx context.Context, resourceGroupName string, accountName string, databaseName string, createUpdateGremlinDatabaseParameters documentdb.GremlinDatabaseCreateUpdateParameters) (result documentdb.GremlinResourcesCreateUpdateGremlinDatabaseFuture, err error)
	CreateUpdateGremlinGraph(ctx context.Context, resourceGroupName string, accountName string, databaseName string, graphName string, createUpdateGremlinGraphParameters documentdb.GremlinGraphCreateUpdateParameters) (result documentdb.GremlinResourcesCreateUpdateGremlinGraphFuture, err error)
	DeleteGremlinDatabase(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.GremlinResourcesDeleteGremlinDatabaseFuture, err error)
	DeleteGremlinGraph(ctx context.Context, resourceGroupName string, accountName string, databaseName string, graphName string) (result documentdb.GremlinResourcesDeleteGremlinGraphFuture, err error)
	GetGremlinDatabase(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.GremlinDatabaseGetResults, err error)
	GetGremlinDatabaseThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.ThroughputSettingsGetResults, err error)
	GetGremlinGraph(ctx context.Context, resourceGroupName string, accountName string, databaseName string, graphName string) (result documentdb.GremlinGraphGetResults, err error)
	GetGremlinGraphThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string, graphName string) (result documentdb.ThroughputSettingsGetResults, err error)
	ListGremlinDatabases(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.GremlinDatabaseListResult, err error)
	ListGremlinGraphs(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.GremlinGraphListResult, err error)
	MigrateGremlinDatabaseToAutoscale(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.GremlinResourcesMigrateGremlinDatabaseToAutoscaleFuture, err error)
	MigrateGremlinDatabaseToManualThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.GremlinResourcesMigrateGremlinDatabaseToManualThroughputFuture, err error)
	MigrateGremlinGraphToAutoscale(ctx context.Context, resourceGroupName string, accountName string, databaseName string, graphName string) (result documentdb.GremlinResourcesMigrateGremlinGraphToAutoscaleFuture, err error)
	MigrateGremlinGraphToManualThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string, graphName string) (result documentdb.GremlinResourcesMigrateGremlinGraphToManualThroughputFuture, err error)
	UpdateGremlinDatabaseThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string, updateThroughputParameters documentdb.ThroughputSettingsUpdateParameters) (result documentdb.GremlinResourcesUpdateGremlinDatabaseThroughputFuture, err error)
	UpdateGremlinGraphThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string, graphName string, updateThroughputParameters documentdb.ThroughputSettingsUpdateParameters) (result documentdb.GremlinResourcesUpdateGremlinGraphThroughputFuture, err error)
}

var _ GremlinResourcesClientAPI = (*documentdb.GremlinResourcesClient)(nil)

// NotebookWorkspacesClientAPI contains the set of methods on the NotebookWorkspacesClient type.
type NotebookWorkspacesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, notebookCreateUpdateParameters documentdb.NotebookWorkspaceCreateUpdateParameters) (result documentdb.NotebookWorkspacesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.NotebookWorkspacesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.NotebookWorkspace, err error)
	ListByDatabaseAccount(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.NotebookWorkspaceListResult, err error)
	ListConnectionInfo(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.NotebookWorkspaceConnectionInfoResult, err error)
	RegenerateAuthToken(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.NotebookWorkspacesRegenerateAuthTokenFuture, err error)
	Start(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.NotebookWorkspacesStartFuture, err error)
}

var _ NotebookWorkspacesClientAPI = (*documentdb.NotebookWorkspacesClient)(nil)

// PrivateEndpointConnectionsClientAPI contains the set of methods on the PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string, parameters documentdb.PrivateEndpointConnection) (result documentdb.PrivateEndpointConnectionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string) (result documentdb.PrivateEndpointConnectionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string) (result documentdb.PrivateEndpointConnection, err error)
	ListByDatabaseAccount(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.PrivateEndpointConnectionListResult, err error)
}

var _ PrivateEndpointConnectionsClientAPI = (*documentdb.PrivateEndpointConnectionsClient)(nil)

// PrivateLinkResourcesClientAPI contains the set of methods on the PrivateLinkResourcesClient type.
type PrivateLinkResourcesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, accountName string, groupName string) (result documentdb.PrivateLinkResource, err error)
	ListByDatabaseAccount(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.PrivateLinkResourceListResult, err error)
}

var _ PrivateLinkResourcesClientAPI = (*documentdb.PrivateLinkResourcesClient)(nil)

// RestorableDatabaseAccountsClientAPI contains the set of methods on the RestorableDatabaseAccountsClient type.
type RestorableDatabaseAccountsClientAPI interface {
	GetByLocation(ctx context.Context, location string, instanceID string) (result documentdb.RestorableDatabaseAccountGetResult, err error)
	List(ctx context.Context) (result documentdb.RestorableDatabaseAccountsListResult, err error)
	ListByLocation(ctx context.Context, location string) (result documentdb.RestorableDatabaseAccountsListResult, err error)
}

var _ RestorableDatabaseAccountsClientAPI = (*documentdb.RestorableDatabaseAccountsClient)(nil)

// RestorableSQLDatabasesClientAPI contains the set of methods on the RestorableSQLDatabasesClient type.
type RestorableSQLDatabasesClientAPI interface {
	List(ctx context.Context, location string, instanceID string) (result documentdb.RestorableSQLDatabasesListResult, err error)
}

var _ RestorableSQLDatabasesClientAPI = (*documentdb.RestorableSQLDatabasesClient)(nil)

// RestorableSQLContainersClientAPI contains the set of methods on the RestorableSQLContainersClient type.
type RestorableSQLContainersClientAPI interface {
	List(ctx context.Context, location string, instanceID string, restorableSQLDatabaseRid string, startTime string, endTime string) (result documentdb.RestorableSQLContainersListResult, err error)
}

var _ RestorableSQLContainersClientAPI = (*documentdb.RestorableSQLContainersClient)(nil)

// RestorableSQLResourcesClientAPI contains the set of methods on the RestorableSQLResourcesClient type.
type RestorableSQLResourcesClientAPI interface {
	List(ctx context.Context, location string, instanceID string, restoreLocation string, restoreTimestampInUtc string) (result documentdb.RestorableSQLResourcesListResult, err error)
}

var _ RestorableSQLResourcesClientAPI = (*documentdb.RestorableSQLResourcesClient)(nil)

// RestorableMongodbDatabasesClientAPI contains the set of methods on the RestorableMongodbDatabasesClient type.
type RestorableMongodbDatabasesClientAPI interface {
	List(ctx context.Context, location string, instanceID string) (result documentdb.RestorableMongodbDatabasesListResult, err error)
}

var _ RestorableMongodbDatabasesClientAPI = (*documentdb.RestorableMongodbDatabasesClient)(nil)

// RestorableMongodbCollectionsClientAPI contains the set of methods on the RestorableMongodbCollectionsClient type.
type RestorableMongodbCollectionsClientAPI interface {
	List(ctx context.Context, location string, instanceID string, restorableMongodbDatabaseRid string) (result documentdb.RestorableMongodbCollectionsListResult, err error)
}

var _ RestorableMongodbCollectionsClientAPI = (*documentdb.RestorableMongodbCollectionsClient)(nil)

// RestorableMongodbResourcesClientAPI contains the set of methods on the RestorableMongodbResourcesClient type.
type RestorableMongodbResourcesClientAPI interface {
	List(ctx context.Context, location string, instanceID string, restoreLocation string, restoreTimestampInUtc string) (result documentdb.RestorableMongodbResourcesListResult, err error)
}

var _ RestorableMongodbResourcesClientAPI = (*documentdb.RestorableMongodbResourcesClient)(nil)
