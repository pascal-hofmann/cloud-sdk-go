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

package deployments

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

// NewUpdateDeploymentNoteParams creates a new UpdateDeploymentNoteParams object
// with the default values initialized.
func NewUpdateDeploymentNoteParams() *UpdateDeploymentNoteParams {
	var ()
	return &UpdateDeploymentNoteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeploymentNoteParamsWithTimeout creates a new UpdateDeploymentNoteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateDeploymentNoteParamsWithTimeout(timeout time.Duration) *UpdateDeploymentNoteParams {
	var ()
	return &UpdateDeploymentNoteParams{

		timeout: timeout,
	}
}

// NewUpdateDeploymentNoteParamsWithContext creates a new UpdateDeploymentNoteParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateDeploymentNoteParamsWithContext(ctx context.Context) *UpdateDeploymentNoteParams {
	var ()
	return &UpdateDeploymentNoteParams{

		Context: ctx,
	}
}

// NewUpdateDeploymentNoteParamsWithHTTPClient creates a new UpdateDeploymentNoteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateDeploymentNoteParamsWithHTTPClient(client *http.Client) *UpdateDeploymentNoteParams {
	var ()
	return &UpdateDeploymentNoteParams{
		HTTPClient: client,
	}
}

/*UpdateDeploymentNoteParams contains all the parameters to send to the API endpoint
for the update deployment note operation typically these are written to a http.Request
*/
type UpdateDeploymentNoteParams struct {

	/*Body
	  New content for deployment note

	*/
	Body *models.Note
	/*DeploymentID
	  Identifier for the deployment

	*/
	DeploymentID string
	/*NoteID
	  Identifier of the note to be updated

	*/
	NoteID string
	/*Version
	  If specified then checks for conflicts against the version of the deployment note

	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update deployment note params
func (o *UpdateDeploymentNoteParams) WithTimeout(timeout time.Duration) *UpdateDeploymentNoteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update deployment note params
func (o *UpdateDeploymentNoteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update deployment note params
func (o *UpdateDeploymentNoteParams) WithContext(ctx context.Context) *UpdateDeploymentNoteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update deployment note params
func (o *UpdateDeploymentNoteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update deployment note params
func (o *UpdateDeploymentNoteParams) WithHTTPClient(client *http.Client) *UpdateDeploymentNoteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update deployment note params
func (o *UpdateDeploymentNoteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update deployment note params
func (o *UpdateDeploymentNoteParams) WithBody(body *models.Note) *UpdateDeploymentNoteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update deployment note params
func (o *UpdateDeploymentNoteParams) SetBody(body *models.Note) {
	o.Body = body
}

// WithDeploymentID adds the deploymentID to the update deployment note params
func (o *UpdateDeploymentNoteParams) WithDeploymentID(deploymentID string) *UpdateDeploymentNoteParams {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the update deployment note params
func (o *UpdateDeploymentNoteParams) SetDeploymentID(deploymentID string) {
	o.DeploymentID = deploymentID
}

// WithNoteID adds the noteID to the update deployment note params
func (o *UpdateDeploymentNoteParams) WithNoteID(noteID string) *UpdateDeploymentNoteParams {
	o.SetNoteID(noteID)
	return o
}

// SetNoteID adds the noteId to the update deployment note params
func (o *UpdateDeploymentNoteParams) SetNoteID(noteID string) {
	o.NoteID = noteID
}

// WithVersion adds the version to the update deployment note params
func (o *UpdateDeploymentNoteParams) WithVersion(version *int64) *UpdateDeploymentNoteParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the update deployment note params
func (o *UpdateDeploymentNoteParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeploymentNoteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param deployment_id
	if err := r.SetPathParam("deployment_id", o.DeploymentID); err != nil {
		return err
	}

	// path param note_id
	if err := r.SetPathParam("note_id", o.NoteID); err != nil {
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
