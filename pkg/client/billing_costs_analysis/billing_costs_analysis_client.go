// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package billing_costs_analysis

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new billing costs analysis API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for billing costs analysis API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetCostsCharts(params *GetCostsChartsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostsChartsOK, error)

	GetCostsChartsByDeployment(params *GetCostsChartsByDeploymentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostsChartsByDeploymentOK, error)

	GetCostsDeployments(params *GetCostsDeploymentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostsDeploymentsOK, error)

	GetCostsItems(params *GetCostsItemsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostsItemsOK, error)

	GetCostsItemsByDeployment(params *GetCostsItemsByDeploymentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostsItemsByDeploymentOK, error)

	GetCostsOverview(params *GetCostsOverviewParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostsOverviewOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetCostsCharts gets charts for the organization currently unavailable in self hosted e c e

  EXPERIMENTAL (it may change in future versions): Retrieves the usage charts for the organization.
*/
func (a *Client) GetCostsCharts(params *GetCostsChartsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostsChartsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCostsChartsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-costs-charts",
		Method:             "GET",
		PathPattern:        "/billing/costs/{organization_id}/charts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCostsChartsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCostsChartsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-costs-charts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCostsChartsByDeployment gets charts by deployment currently unavailable in self hosted e c e

  EXPERIMENTAL (it may change in future versions): Retrieves the usage charts for the given  deployment.
*/
func (a *Client) GetCostsChartsByDeployment(params *GetCostsChartsByDeploymentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostsChartsByDeploymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCostsChartsByDeploymentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-costs-charts-by-deployment",
		Method:             "GET",
		PathPattern:        "/billing/costs/{organization_id}/deployments/{deployment_id}/charts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCostsChartsByDeploymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCostsChartsByDeploymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-costs-charts-by-deployment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCostsDeployments gets deployments costs for the organization currently unavailable in self hosted e c e

  EXPERIMENTAL (it may change in future versions): Retrieves the costs associated with all deployments for the organization.
*/
func (a *Client) GetCostsDeployments(params *GetCostsDeploymentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostsDeploymentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCostsDeploymentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-costs-deployments",
		Method:             "GET",
		PathPattern:        "/billing/costs/{organization_id}/deployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCostsDeploymentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCostsDeploymentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-costs-deployments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCostsItems gets itemized costs for the organization currently unavailable in self hosted e c e

  EXPERIMENTAL (it may change in future versions): Retrieves the itemized costs for the organization.
*/
func (a *Client) GetCostsItems(params *GetCostsItemsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostsItemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCostsItemsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-costs-items",
		Method:             "GET",
		PathPattern:        "/billing/costs/{organization_id}/items",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCostsItemsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCostsItemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-costs-items: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCostsItemsByDeployment gets itemized costs by deployments currently unavailable in self hosted e c e

  EXPERIMENTAL (it may change in future versions): Retrieves the itemized costs for the given deployment.
*/
func (a *Client) GetCostsItemsByDeployment(params *GetCostsItemsByDeploymentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostsItemsByDeploymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCostsItemsByDeploymentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-costs-items-by-deployment",
		Method:             "GET",
		PathPattern:        "/billing/costs/{organization_id}/deployments/{deployment_id}/items",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCostsItemsByDeploymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCostsItemsByDeploymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-costs-items-by-deployment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCostsOverview gets costs overview for the organization currently unavailable in self hosted e c e

  EXPERIMENTAL (it may change in future versions): Retrieves an overview of the costs by organization ID.
*/
func (a *Client) GetCostsOverview(params *GetCostsOverviewParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostsOverviewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCostsOverviewParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-costs-overview",
		Method:             "GET",
		PathPattern:        "/billing/costs/{organization_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCostsOverviewReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCostsOverviewOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-costs-overview: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
