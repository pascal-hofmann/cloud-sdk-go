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

package clusters_elasticsearch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewCreateEsClusterParams creates a new CreateEsClusterParams object
// with the default values initialized.
func NewCreateEsClusterParams() *CreateEsClusterParams {
	var (
		validateOnlyDefault = bool(false)
	)
	return &CreateEsClusterParams{
		ValidateOnly: &validateOnlyDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateEsClusterParamsWithTimeout creates a new CreateEsClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateEsClusterParamsWithTimeout(timeout time.Duration) *CreateEsClusterParams {
	var (
		validateOnlyDefault = bool(false)
	)
	return &CreateEsClusterParams{
		ValidateOnly: &validateOnlyDefault,

		timeout: timeout,
	}
}

// NewCreateEsClusterParamsWithContext creates a new CreateEsClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateEsClusterParamsWithContext(ctx context.Context) *CreateEsClusterParams {
	var (
		validateOnlyDefault = bool(false)
	)
	return &CreateEsClusterParams{
		ValidateOnly: &validateOnlyDefault,

		Context: ctx,
	}
}

// NewCreateEsClusterParamsWithHTTPClient creates a new CreateEsClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateEsClusterParamsWithHTTPClient(client *http.Client) *CreateEsClusterParams {
	var (
		validateOnlyDefault = bool(false)
	)
	return &CreateEsClusterParams{
		ValidateOnly: &validateOnlyDefault,
		HTTPClient:   client,
	}
}

/*CreateEsClusterParams contains all the parameters to send to the API endpoint
for the create es cluster operation typically these are written to a http.Request
*/
type CreateEsClusterParams struct {

	/*Body
	  The cluster definition

	*/
	Body *models.CreateElasticsearchClusterRequest
	/*RequestID
	  (Optional) The idempotency token. When two create requests share the same `request_id` (minimum size of 32 characters, maximum size of 128 characters), only one cluster is created. The second request returns the information for that cluster, but the password field is blank.

	*/
	RequestID *string
	/*ValidateOnly
	  When `true`, validates the cluster definition without creating the cluster.

	*/
	ValidateOnly *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create es cluster params
func (o *CreateEsClusterParams) WithTimeout(timeout time.Duration) *CreateEsClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create es cluster params
func (o *CreateEsClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create es cluster params
func (o *CreateEsClusterParams) WithContext(ctx context.Context) *CreateEsClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create es cluster params
func (o *CreateEsClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create es cluster params
func (o *CreateEsClusterParams) WithHTTPClient(client *http.Client) *CreateEsClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create es cluster params
func (o *CreateEsClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create es cluster params
func (o *CreateEsClusterParams) WithBody(body *models.CreateElasticsearchClusterRequest) *CreateEsClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create es cluster params
func (o *CreateEsClusterParams) SetBody(body *models.CreateElasticsearchClusterRequest) {
	o.Body = body
}

// WithRequestID adds the requestID to the create es cluster params
func (o *CreateEsClusterParams) WithRequestID(requestID *string) *CreateEsClusterParams {
	o.SetRequestID(requestID)
	return o
}

// SetRequestID adds the requestId to the create es cluster params
func (o *CreateEsClusterParams) SetRequestID(requestID *string) {
	o.RequestID = requestID
}

// WithValidateOnly adds the validateOnly to the create es cluster params
func (o *CreateEsClusterParams) WithValidateOnly(validateOnly *bool) *CreateEsClusterParams {
	o.SetValidateOnly(validateOnly)
	return o
}

// SetValidateOnly adds the validateOnly to the create es cluster params
func (o *CreateEsClusterParams) SetValidateOnly(validateOnly *bool) {
	o.ValidateOnly = validateOnly
}

// WriteToRequest writes these params to a swagger request
func (o *CreateEsClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.RequestID != nil {

		// query param request_id
		var qrRequestID string
		if o.RequestID != nil {
			qrRequestID = *o.RequestID
		}
		qRequestID := qrRequestID
		if qRequestID != "" {
			if err := r.SetQueryParam("request_id", qRequestID); err != nil {
				return err
			}
		}

	}

	if o.ValidateOnly != nil {

		// query param validate_only
		var qrValidateOnly bool
		if o.ValidateOnly != nil {
			qrValidateOnly = *o.ValidateOnly
		}
		qValidateOnly := swag.FormatBool(qrValidateOnly)
		if qValidateOnly != "" {
			if err := r.SetQueryParam("validate_only", qValidateOnly); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
