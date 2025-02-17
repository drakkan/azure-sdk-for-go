package media

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// OperationStatusesClient is the client for the OperationStatuses methods of the Media service.
type OperationStatusesClient struct {
	BaseClient
}

// NewOperationStatusesClient creates an instance of the OperationStatusesClient client.
func NewOperationStatusesClient(subscriptionID string) OperationStatusesClient {
	return NewOperationStatusesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewOperationStatusesClientWithBaseURI creates an instance of the OperationStatusesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewOperationStatusesClientWithBaseURI(baseURI string, subscriptionID string) OperationStatusesClient {
	return OperationStatusesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get asset track operation status.
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// assetName - the Asset name.
// trackName - the Asset Track name.
// operationID - operation Id.
func (client OperationStatusesClient) Get(ctx context.Context, resourceGroupName string, accountName string, assetName string, trackName string, operationID string) (result AssetTrackOperationStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationStatusesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, assetName, trackName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.OperationStatusesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.OperationStatusesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.OperationStatusesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client OperationStatusesClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, assetName string, trackName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"assetName":         autorest.Encode("path", assetName),
		"operationId":       autorest.Encode("path", operationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"trackName":         autorest.Encode("path", trackName),
	}

	const APIVersion = "2021-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/tracks/{trackName}/operationStatuses/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client OperationStatusesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client OperationStatusesClient) GetResponder(resp *http.Response) (result AssetTrackOperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
