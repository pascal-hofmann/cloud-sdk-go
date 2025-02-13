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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// UpdateProxiesSettingsReader is a Reader for the UpdateProxiesSettings structure.
type UpdateProxiesSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProxiesSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProxiesSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewUpdateProxiesSettingsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewUpdateProxiesSettingsRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProxiesSettingsOK creates a UpdateProxiesSettingsOK with default headers values
func NewUpdateProxiesSettingsOK() *UpdateProxiesSettingsOK {
	return &UpdateProxiesSettingsOK{}
}

/* UpdateProxiesSettingsOK describes a response with status code 200, with default header values.

Returns the updated settings
*/
type UpdateProxiesSettingsOK struct {
	Payload *models.ProxiesSettings
}

func (o *UpdateProxiesSettingsOK) Error() string {
	return fmt.Sprintf("[PATCH /platform/infrastructure/proxies/settings][%d] updateProxiesSettingsOK  %+v", 200, o.Payload)
}
func (o *UpdateProxiesSettingsOK) GetPayload() *models.ProxiesSettings {
	return o.Payload
}

func (o *UpdateProxiesSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxiesSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProxiesSettingsConflict creates a UpdateProxiesSettingsConflict with default headers values
func NewUpdateProxiesSettingsConflict() *UpdateProxiesSettingsConflict {
	return &UpdateProxiesSettingsConflict{}
}

/* UpdateProxiesSettingsConflict describes a response with status code 409, with default header values.

There is a version conflict. (code: `proxies.version_conflict`)
*/
type UpdateProxiesSettingsConflict struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateProxiesSettingsConflict) Error() string {
	return fmt.Sprintf("[PATCH /platform/infrastructure/proxies/settings][%d] updateProxiesSettingsConflict  %+v", 409, o.Payload)
}
func (o *UpdateProxiesSettingsConflict) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateProxiesSettingsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-error-codes
	hdrXCloudErrorCodes := response.GetHeader("x-cloud-error-codes")

	if hdrXCloudErrorCodes != "" {
		o.XCloudErrorCodes = hdrXCloudErrorCodes
	}

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProxiesSettingsRetryWith creates a UpdateProxiesSettingsRetryWith with default headers values
func NewUpdateProxiesSettingsRetryWith() *UpdateProxiesSettingsRetryWith {
	return &UpdateProxiesSettingsRetryWith{}
}

/* UpdateProxiesSettingsRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type UpdateProxiesSettingsRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateProxiesSettingsRetryWith) Error() string {
	return fmt.Sprintf("[PATCH /platform/infrastructure/proxies/settings][%d] updateProxiesSettingsRetryWith  %+v", 449, o.Payload)
}
func (o *UpdateProxiesSettingsRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateProxiesSettingsRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-error-codes
	hdrXCloudErrorCodes := response.GetHeader("x-cloud-error-codes")

	if hdrXCloudErrorCodes != "" {
		o.XCloudErrorCodes = hdrXCloudErrorCodes
	}

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
