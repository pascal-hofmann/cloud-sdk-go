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

package comments

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

// NewCreateCommentParams creates a new CreateCommentParams object
// with the default values initialized.
func NewCreateCommentParams() *CreateCommentParams {
	var ()
	return &CreateCommentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCommentParamsWithTimeout creates a new CreateCommentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateCommentParamsWithTimeout(timeout time.Duration) *CreateCommentParams {
	var ()
	return &CreateCommentParams{

		timeout: timeout,
	}
}

// NewCreateCommentParamsWithContext creates a new CreateCommentParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateCommentParamsWithContext(ctx context.Context) *CreateCommentParams {
	var ()
	return &CreateCommentParams{

		Context: ctx,
	}
}

// NewCreateCommentParamsWithHTTPClient creates a new CreateCommentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateCommentParamsWithHTTPClient(client *http.Client) *CreateCommentParams {
	var ()
	return &CreateCommentParams{
		HTTPClient: client,
	}
}

/*CreateCommentParams contains all the parameters to send to the API endpoint
for the create comment operation typically these are written to a http.Request
*/
type CreateCommentParams struct {

	/*Body
	  Data for comment creation

	*/
	Body *models.CommentCreateRequest
	/*ResourceID
	  Id of the Resource that a Comment belongs to.

	*/
	ResourceID string
	/*ResourceType
	  The kind of Resource that a Comment belongs to. Should be one of [elasticsearch, kibana, apm, appsearch, enterprisesearch, sitesearch, allocator, constructor, runner, proxy].

	*/
	ResourceType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create comment params
func (o *CreateCommentParams) WithTimeout(timeout time.Duration) *CreateCommentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create comment params
func (o *CreateCommentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create comment params
func (o *CreateCommentParams) WithContext(ctx context.Context) *CreateCommentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create comment params
func (o *CreateCommentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create comment params
func (o *CreateCommentParams) WithHTTPClient(client *http.Client) *CreateCommentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create comment params
func (o *CreateCommentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create comment params
func (o *CreateCommentParams) WithBody(body *models.CommentCreateRequest) *CreateCommentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create comment params
func (o *CreateCommentParams) SetBody(body *models.CommentCreateRequest) {
	o.Body = body
}

// WithResourceID adds the resourceID to the create comment params
func (o *CreateCommentParams) WithResourceID(resourceID string) *CreateCommentParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the create comment params
func (o *CreateCommentParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WithResourceType adds the resourceType to the create comment params
func (o *CreateCommentParams) WithResourceType(resourceType string) *CreateCommentParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the create comment params
func (o *CreateCommentParams) SetResourceType(resourceType string) {
	o.ResourceType = resourceType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCommentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param resource_id
	if err := r.SetPathParam("resource_id", o.ResourceID); err != nil {
		return err
	}

	// path param resource_type
	if err := r.SetPathParam("resource_type", o.ResourceType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
