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

package platform_infrastructure

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

// NewUpdateProxiesFilteredGroupParams creates a new UpdateProxiesFilteredGroupParams object
// with the default values initialized.
func NewUpdateProxiesFilteredGroupParams() *UpdateProxiesFilteredGroupParams {
	var ()
	return &UpdateProxiesFilteredGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProxiesFilteredGroupParamsWithTimeout creates a new UpdateProxiesFilteredGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateProxiesFilteredGroupParamsWithTimeout(timeout time.Duration) *UpdateProxiesFilteredGroupParams {
	var ()
	return &UpdateProxiesFilteredGroupParams{

		timeout: timeout,
	}
}

// NewUpdateProxiesFilteredGroupParamsWithContext creates a new UpdateProxiesFilteredGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateProxiesFilteredGroupParamsWithContext(ctx context.Context) *UpdateProxiesFilteredGroupParams {
	var ()
	return &UpdateProxiesFilteredGroupParams{

		Context: ctx,
	}
}

// NewUpdateProxiesFilteredGroupParamsWithHTTPClient creates a new UpdateProxiesFilteredGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateProxiesFilteredGroupParamsWithHTTPClient(client *http.Client) *UpdateProxiesFilteredGroupParams {
	var ()
	return &UpdateProxiesFilteredGroupParams{
		HTTPClient: client,
	}
}

/*UpdateProxiesFilteredGroupParams contains all the parameters to send to the API endpoint
for the update proxies filtered group operation typically these are written to a http.Request
*/
type UpdateProxiesFilteredGroupParams struct {

	/*Body
	  Data for the filtered group of proxies to update

	*/
	Body *models.ProxiesFilteredGroup
	/*ProxiesFilteredGroupID
	  "The identifier for the filtered group of proxies

	*/
	ProxiesFilteredGroupID string
	/*Version
	  Checks for conflicts against the metadata version, then returns the value in the `x-cloud-resource-version` header.

	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update proxies filtered group params
func (o *UpdateProxiesFilteredGroupParams) WithTimeout(timeout time.Duration) *UpdateProxiesFilteredGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update proxies filtered group params
func (o *UpdateProxiesFilteredGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update proxies filtered group params
func (o *UpdateProxiesFilteredGroupParams) WithContext(ctx context.Context) *UpdateProxiesFilteredGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update proxies filtered group params
func (o *UpdateProxiesFilteredGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update proxies filtered group params
func (o *UpdateProxiesFilteredGroupParams) WithHTTPClient(client *http.Client) *UpdateProxiesFilteredGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update proxies filtered group params
func (o *UpdateProxiesFilteredGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update proxies filtered group params
func (o *UpdateProxiesFilteredGroupParams) WithBody(body *models.ProxiesFilteredGroup) *UpdateProxiesFilteredGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update proxies filtered group params
func (o *UpdateProxiesFilteredGroupParams) SetBody(body *models.ProxiesFilteredGroup) {
	o.Body = body
}

// WithProxiesFilteredGroupID adds the proxiesFilteredGroupID to the update proxies filtered group params
func (o *UpdateProxiesFilteredGroupParams) WithProxiesFilteredGroupID(proxiesFilteredGroupID string) *UpdateProxiesFilteredGroupParams {
	o.SetProxiesFilteredGroupID(proxiesFilteredGroupID)
	return o
}

// SetProxiesFilteredGroupID adds the proxiesFilteredGroupId to the update proxies filtered group params
func (o *UpdateProxiesFilteredGroupParams) SetProxiesFilteredGroupID(proxiesFilteredGroupID string) {
	o.ProxiesFilteredGroupID = proxiesFilteredGroupID
}

// WithVersion adds the version to the update proxies filtered group params
func (o *UpdateProxiesFilteredGroupParams) WithVersion(version *int64) *UpdateProxiesFilteredGroupParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the update proxies filtered group params
func (o *UpdateProxiesFilteredGroupParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProxiesFilteredGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param proxies_filtered_group_id
	if err := r.SetPathParam("proxies_filtered_group_id", o.ProxiesFilteredGroupID); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion int64
		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatInt64(qrVersion)
		if qVersion != "" {
			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
