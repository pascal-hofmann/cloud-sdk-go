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

package deployments_ip_filtering

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new deployments ip filtering API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for deployments ip filtering API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateIPFilterRuleset(params *CreateIPFilterRulesetParams, authInfo runtime.ClientAuthInfoWriter) (*CreateIPFilterRulesetCreated, error)

	CreateIPFilterRulesetAssociation(params *CreateIPFilterRulesetAssociationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateIPFilterRulesetAssociationCreated, error)

	DeleteIPFilterRuleset(params *DeleteIPFilterRulesetParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteIPFilterRulesetOK, error)

	DeleteIPFilterRulesetAssociation(params *DeleteIPFilterRulesetAssociationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteIPFilterRulesetAssociationOK, error)

	GetIPFilterDeploymentRulesetAssociations(params *GetIPFilterDeploymentRulesetAssociationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetIPFilterDeploymentRulesetAssociationsOK, error)

	GetIPFilterRuleset(params *GetIPFilterRulesetParams, authInfo runtime.ClientAuthInfoWriter) (*GetIPFilterRulesetOK, error)

	GetIPFilterRulesetDeploymentAssociations(params *GetIPFilterRulesetDeploymentAssociationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetIPFilterRulesetDeploymentAssociationsOK, error)

	GetIPFilterRulesets(params *GetIPFilterRulesetsParams, authInfo runtime.ClientAuthInfoWriter) (*GetIPFilterRulesetsOK, error)

	UpdateIPFilterRuleset(params *UpdateIPFilterRulesetParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateIPFilterRulesetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateIPFilterRuleset creates a ruleset

  Creates a ruleset that combines a set of rules.
*/
func (a *Client) CreateIPFilterRuleset(params *CreateIPFilterRulesetParams, authInfo runtime.ClientAuthInfoWriter) (*CreateIPFilterRulesetCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateIPFilterRulesetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "create-ip-filter-ruleset",
		Method:             "POST",
		PathPattern:        "/deployments/ip-filtering/rulesets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateIPFilterRulesetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateIPFilterRulesetCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create-ip-filter-ruleset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateIPFilterRulesetAssociation creates ruleset association

  Applies the ruleset to the specified deployment.
*/
func (a *Client) CreateIPFilterRulesetAssociation(params *CreateIPFilterRulesetAssociationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateIPFilterRulesetAssociationCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateIPFilterRulesetAssociationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "create-ip-filter-ruleset-association",
		Method:             "POST",
		PathPattern:        "/deployments/ip-filtering/rulesets/{ruleset_id}/associations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateIPFilterRulesetAssociationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateIPFilterRulesetAssociationCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create-ip-filter-ruleset-association: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteIPFilterRuleset deletes a ruleset

  Deletes the ruleset by ID.
*/
func (a *Client) DeleteIPFilterRuleset(params *DeleteIPFilterRulesetParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteIPFilterRulesetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteIPFilterRulesetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete-ip-filter-ruleset",
		Method:             "DELETE",
		PathPattern:        "/deployments/ip-filtering/rulesets/{ruleset_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteIPFilterRulesetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteIPFilterRulesetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-ip-filter-ruleset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteIPFilterRulesetAssociation deletes ruleset association

  Deletes the traffic rules in the ruleset from the deployment.
*/
func (a *Client) DeleteIPFilterRulesetAssociation(params *DeleteIPFilterRulesetAssociationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteIPFilterRulesetAssociationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteIPFilterRulesetAssociationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete-ip-filter-ruleset-association",
		Method:             "DELETE",
		PathPattern:        "/deployments/ip-filtering/rulesets/{ruleset_id}/associations/{association_type}/{associated_entity_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteIPFilterRulesetAssociationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteIPFilterRulesetAssociationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-ip-filter-ruleset-association: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetIPFilterDeploymentRulesetAssociations gets associated rulesets

  Retrieves the rulesets associated with a deployment.
*/
func (a *Client) GetIPFilterDeploymentRulesetAssociations(params *GetIPFilterDeploymentRulesetAssociationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetIPFilterDeploymentRulesetAssociationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIPFilterDeploymentRulesetAssociationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-ip-filter-deployment-ruleset-associations",
		Method:             "GET",
		PathPattern:        "/deployments/ip-filtering/associations/{association_type}/{associated_entity_id}/rulesets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIPFilterDeploymentRulesetAssociationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIPFilterDeploymentRulesetAssociationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-ip-filter-deployment-ruleset-associations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetIPFilterRuleset gets a ruleset

  Retrieves the ruleset by ID.
*/
func (a *Client) GetIPFilterRuleset(params *GetIPFilterRulesetParams, authInfo runtime.ClientAuthInfoWriter) (*GetIPFilterRulesetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIPFilterRulesetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-ip-filter-ruleset",
		Method:             "GET",
		PathPattern:        "/deployments/ip-filtering/rulesets/{ruleset_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIPFilterRulesetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIPFilterRulesetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-ip-filter-ruleset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetIPFilterRulesetDeploymentAssociations gets associated deployments

  Retrieves a list of deployments that are associated to the specified ruleset.
*/
func (a *Client) GetIPFilterRulesetDeploymentAssociations(params *GetIPFilterRulesetDeploymentAssociationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetIPFilterRulesetDeploymentAssociationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIPFilterRulesetDeploymentAssociationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-ip-filter-ruleset-deployment-associations",
		Method:             "GET",
		PathPattern:        "/deployments/ip-filtering/rulesets/{ruleset_id}/associations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIPFilterRulesetDeploymentAssociationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIPFilterRulesetDeploymentAssociationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-ip-filter-ruleset-deployment-associations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetIPFilterRulesets gets all rulesets

  Retrieves all of the user rulesets.
*/
func (a *Client) GetIPFilterRulesets(params *GetIPFilterRulesetsParams, authInfo runtime.ClientAuthInfoWriter) (*GetIPFilterRulesetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIPFilterRulesetsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-ip-filter-rulesets",
		Method:             "GET",
		PathPattern:        "/deployments/ip-filtering/rulesets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIPFilterRulesetsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIPFilterRulesetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-ip-filter-rulesets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateIPFilterRuleset updates a ruleset

  Updates the ruleset with the definition.
*/
func (a *Client) UpdateIPFilterRuleset(params *UpdateIPFilterRulesetParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateIPFilterRulesetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateIPFilterRulesetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "update-ip-filter-ruleset",
		Method:             "PUT",
		PathPattern:        "/deployments/ip-filtering/rulesets/{ruleset_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateIPFilterRulesetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateIPFilterRulesetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update-ip-filter-ruleset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
