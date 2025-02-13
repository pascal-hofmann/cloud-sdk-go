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

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewCreateOrganizationInvitationsParams creates a new CreateOrganizationInvitationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateOrganizationInvitationsParams() *CreateOrganizationInvitationsParams {
	return &CreateOrganizationInvitationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrganizationInvitationsParamsWithTimeout creates a new CreateOrganizationInvitationsParams object
// with the ability to set a timeout on a request.
func NewCreateOrganizationInvitationsParamsWithTimeout(timeout time.Duration) *CreateOrganizationInvitationsParams {
	return &CreateOrganizationInvitationsParams{
		timeout: timeout,
	}
}

// NewCreateOrganizationInvitationsParamsWithContext creates a new CreateOrganizationInvitationsParams object
// with the ability to set a context for a request.
func NewCreateOrganizationInvitationsParamsWithContext(ctx context.Context) *CreateOrganizationInvitationsParams {
	return &CreateOrganizationInvitationsParams{
		Context: ctx,
	}
}

// NewCreateOrganizationInvitationsParamsWithHTTPClient creates a new CreateOrganizationInvitationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateOrganizationInvitationsParamsWithHTTPClient(client *http.Client) *CreateOrganizationInvitationsParams {
	return &CreateOrganizationInvitationsParams{
		HTTPClient: client,
	}
}

/* CreateOrganizationInvitationsParams contains all the parameters to send to the API endpoint
   for the create organization invitations operation.

   Typically these are written to a http.Request.
*/
type CreateOrganizationInvitationsParams struct {

	/* Body.

	   The organization invitations to create or refresh
	*/
	Body *models.OrganizationInvitationRequest

	/* OrganizationID.

	   Identifier for the Organization
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create organization invitations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrganizationInvitationsParams) WithDefaults() *CreateOrganizationInvitationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create organization invitations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrganizationInvitationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create organization invitations params
func (o *CreateOrganizationInvitationsParams) WithTimeout(timeout time.Duration) *CreateOrganizationInvitationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create organization invitations params
func (o *CreateOrganizationInvitationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create organization invitations params
func (o *CreateOrganizationInvitationsParams) WithContext(ctx context.Context) *CreateOrganizationInvitationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create organization invitations params
func (o *CreateOrganizationInvitationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create organization invitations params
func (o *CreateOrganizationInvitationsParams) WithHTTPClient(client *http.Client) *CreateOrganizationInvitationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create organization invitations params
func (o *CreateOrganizationInvitationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create organization invitations params
func (o *CreateOrganizationInvitationsParams) WithBody(body *models.OrganizationInvitationRequest) *CreateOrganizationInvitationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create organization invitations params
func (o *CreateOrganizationInvitationsParams) SetBody(body *models.OrganizationInvitationRequest) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the create organization invitations params
func (o *CreateOrganizationInvitationsParams) WithOrganizationID(organizationID string) *CreateOrganizationInvitationsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create organization invitations params
func (o *CreateOrganizationInvitationsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrganizationInvitationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
