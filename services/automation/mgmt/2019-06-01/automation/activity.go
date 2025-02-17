package automation

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

// ActivityClient is the automation Client
type ActivityClient struct {
	BaseClient
}

// NewActivityClient creates an instance of the ActivityClient client.
func NewActivityClient(subscriptionID string) ActivityClient {
	return NewActivityClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewActivityClientWithBaseURI creates an instance of the ActivityClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewActivityClientWithBaseURI(baseURI string, subscriptionID string) ActivityClient {
	return ActivityClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get retrieve the activity in the module identified by module name and activity name.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// moduleName - the name of module.
// activityName - the name of activity.
func (client ActivityClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, activityName string) (result Activity, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActivityClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.ActivityClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, automationAccountName, moduleName, activityName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ActivityClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.ActivityClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ActivityClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ActivityClient) GetPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, activityName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"activityName":          autorest.Encode("path", activityName),
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"moduleName":            autorest.Encode("path", moduleName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}/activities/{activityName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ActivityClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ActivityClient) GetResponder(resp *http.Response) (result Activity, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByModule retrieve a list of activities in the module identified by module name.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// moduleName - the name of module.
func (client ActivityClient) ListByModule(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string) (result ActivityListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActivityClient.ListByModule")
		defer func() {
			sc := -1
			if result.alr.Response.Response != nil {
				sc = result.alr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.ActivityClient", "ListByModule", err.Error())
	}

	result.fn = client.listByModuleNextResults
	req, err := client.ListByModulePreparer(ctx, resourceGroupName, automationAccountName, moduleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ActivityClient", "ListByModule", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByModuleSender(req)
	if err != nil {
		result.alr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.ActivityClient", "ListByModule", resp, "Failure sending request")
		return
	}

	result.alr, err = client.ListByModuleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ActivityClient", "ListByModule", resp, "Failure responding to request")
		return
	}
	if result.alr.hasNextLink() && result.alr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByModulePreparer prepares the ListByModule request.
func (client ActivityClient) ListByModulePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"moduleName":            autorest.Encode("path", moduleName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}/activities", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByModuleSender sends the ListByModule request. The method will close the
// http.Response Body if it receives an error.
func (client ActivityClient) ListByModuleSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByModuleResponder handles the response to the ListByModule request. The method always
// closes the http.Response Body.
func (client ActivityClient) ListByModuleResponder(resp *http.Response) (result ActivityListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByModuleNextResults retrieves the next set of results, if any.
func (client ActivityClient) listByModuleNextResults(ctx context.Context, lastResults ActivityListResult) (result ActivityListResult, err error) {
	req, err := lastResults.activityListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "automation.ActivityClient", "listByModuleNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByModuleSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "automation.ActivityClient", "listByModuleNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByModuleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ActivityClient", "listByModuleNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByModuleComplete enumerates all values, automatically crossing page boundaries as required.
func (client ActivityClient) ListByModuleComplete(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string) (result ActivityListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActivityClient.ListByModule")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByModule(ctx, resourceGroupName, automationAccountName, moduleName)
	return
}
