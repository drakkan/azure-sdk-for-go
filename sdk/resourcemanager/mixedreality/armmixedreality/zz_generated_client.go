//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmixedreality

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// Client contains the methods for the MixedRealityClient group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewClient creates a new instance of Client with the specified values.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*Client, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &Client{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CheckNameAvailabilityLocal - Check Name Availability for local uniqueness
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-03-01-preview
// location - The location in which uniqueness will be verified.
// checkNameAvailability - Check Name Availability Request.
// options - ClientCheckNameAvailabilityLocalOptions contains the optional parameters for the Client.CheckNameAvailabilityLocal
// method.
func (client *Client) CheckNameAvailabilityLocal(ctx context.Context, location string, checkNameAvailability CheckNameAvailabilityRequest, options *ClientCheckNameAvailabilityLocalOptions) (ClientCheckNameAvailabilityLocalResponse, error) {
	req, err := client.checkNameAvailabilityLocalCreateRequest(ctx, location, checkNameAvailability, options)
	if err != nil {
		return ClientCheckNameAvailabilityLocalResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientCheckNameAvailabilityLocalResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientCheckNameAvailabilityLocalResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityLocalHandleResponse(resp)
}

// checkNameAvailabilityLocalCreateRequest creates the CheckNameAvailabilityLocal request.
func (client *Client) checkNameAvailabilityLocalCreateRequest(ctx context.Context, location string, checkNameAvailability CheckNameAvailabilityRequest, options *ClientCheckNameAvailabilityLocalOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.MixedReality/locations/{location}/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, checkNameAvailability)
}

// checkNameAvailabilityLocalHandleResponse handles the CheckNameAvailabilityLocal response.
func (client *Client) checkNameAvailabilityLocalHandleResponse(resp *http.Response) (ClientCheckNameAvailabilityLocalResponse, error) {
	result := ClientCheckNameAvailabilityLocalResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResponse); err != nil {
		return ClientCheckNameAvailabilityLocalResponse{}, err
	}
	return result, nil
}
