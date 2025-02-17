package documentdb

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RestorableSQLContainersClient is the client for the RestorableSQLContainers methods of the Documentdb service.
type RestorableSQLContainersClient struct {
	BaseClient
}

// NewRestorableSQLContainersClient creates an instance of the RestorableSQLContainersClient client.
func NewRestorableSQLContainersClient(subscriptionID string) RestorableSQLContainersClient {
	return NewRestorableSQLContainersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRestorableSQLContainersClientWithBaseURI creates an instance of the RestorableSQLContainersClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewRestorableSQLContainersClientWithBaseURI(baseURI string, subscriptionID string) RestorableSQLContainersClient {
	return RestorableSQLContainersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List show the event feed of all mutations done on all the Azure Cosmos DB SQL containers under a specific database.
// This helps in scenario where container was accidentally deleted.  This API requires
// 'Microsoft.DocumentDB/locations/restorableDatabaseAccounts/.../read' permission
// Parameters:
// location - cosmos DB region, with spaces between words and each word capitalized.
// instanceID - the instanceId GUID of a restorable database account.
// restorableSQLDatabaseRid - the resource ID of the SQL database.
// startTime - restorable Sql containers event feed start time.
// endTime - restorable Sql containers event feed end time.
func (client RestorableSQLContainersClient) List(ctx context.Context, location string, instanceID string, restorableSQLDatabaseRid string, startTime string, endTime string) (result RestorableSQLContainersListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RestorableSQLContainersClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("documentdb.RestorableSQLContainersClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, location, instanceID, restorableSQLDatabaseRid, startTime, endTime)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.RestorableSQLContainersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "documentdb.RestorableSQLContainersClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.RestorableSQLContainersClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client RestorableSQLContainersClient) ListPreparer(ctx context.Context, location string, instanceID string, restorableSQLDatabaseRid string, startTime string, endTime string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":     autorest.Encode("path", instanceID),
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(restorableSQLDatabaseRid) > 0 {
		queryParameters["restorableSqlDatabaseRid"] = autorest.Encode("query", restorableSQLDatabaseRid)
	}
	if len(startTime) > 0 {
		queryParameters["startTime"] = autorest.Encode("query", startTime)
	}
	if len(endTime) > 0 {
		queryParameters["endTime"] = autorest.Encode("query", endTime)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableSqlContainers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RestorableSQLContainersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RestorableSQLContainersClient) ListResponder(resp *http.Response) (result RestorableSQLContainersListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
