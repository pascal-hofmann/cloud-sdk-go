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

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewReAuthenticateParams creates a new ReAuthenticateParams object
// with the default values initialized.
func NewReAuthenticateParams() *ReAuthenticateParams {
	var ()
	return &ReAuthenticateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReAuthenticateParamsWithTimeout creates a new ReAuthenticateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReAuthenticateParamsWithTimeout(timeout time.Duration) *ReAuthenticateParams {
	var ()
	return &ReAuthenticateParams{

		timeout: timeout,
	}
}

// NewReAuthenticateParamsWithContext creates a new ReAuthenticateParams object
// with the default values initialized, and the ability to set a context for a request
func NewReAuthenticateParamsWithContext(ctx context.Context) *ReAuthenticateParams {
	var ()
	return &ReAuthenticateParams{

		Context: ctx,
	}
}

// NewReAuthenticateParamsWithHTTPClient creates a new ReAuthenticateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReAuthenticateParamsWithHTTPClient(client *http.Client) *ReAuthenticateParams {
	var ()
	return &ReAuthenticateParams{
		HTTPClient: client,
	}
}

/*ReAuthenticateParams contains all the parameters to send to the API endpoint
for the re authenticate operation typically these are written to a http.Request
*/
type ReAuthenticateParams struct {

	/*Body
	  The request to reauthenticate

	*/
	Body *models.ReAuthenticationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the re authenticate params
func (o *ReAuthenticateParams) WithTimeout(timeout time.Duration) *ReAuthenticateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the re authenticate params
func (o *ReAuthenticateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the re authenticate params
func (o *ReAuthenticateParams) WithContext(ctx context.Context) *ReAuthenticateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the re authenticate params
func (o *ReAuthenticateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the re authenticate params
func (o *ReAuthenticateParams) WithHTTPClient(client *http.Client) *ReAuthenticateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the re authenticate params
func (o *ReAuthenticateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the re authenticate params
func (o *ReAuthenticateParams) WithBody(body *models.ReAuthenticationRequest) *ReAuthenticateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the re authenticate params
func (o *ReAuthenticateParams) SetBody(body *models.ReAuthenticationRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ReAuthenticateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
