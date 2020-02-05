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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// StartConstructorMaintenanceModeReader is a Reader for the StartConstructorMaintenanceMode structure.
type StartConstructorMaintenanceModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartConstructorMaintenanceModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewStartConstructorMaintenanceModeAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewStartConstructorMaintenanceModeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStartConstructorMaintenanceModeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewStartConstructorMaintenanceModeRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStartConstructorMaintenanceModeAccepted creates a StartConstructorMaintenanceModeAccepted with default headers values
func NewStartConstructorMaintenanceModeAccepted() *StartConstructorMaintenanceModeAccepted {
	return &StartConstructorMaintenanceModeAccepted{}
}

/*StartConstructorMaintenanceModeAccepted handles this case with default header values.

The start maintenance mode command was issued successfully
*/
type StartConstructorMaintenanceModeAccepted struct {
	Payload models.EmptyResponse
}

func (o *StartConstructorMaintenanceModeAccepted) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/constructors/{constructor_id}/maintenance-mode/_start][%d] startConstructorMaintenanceModeAccepted  %+v", 202, o.Payload)
}

func (o *StartConstructorMaintenanceModeAccepted) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *StartConstructorMaintenanceModeAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartConstructorMaintenanceModeForbidden creates a StartConstructorMaintenanceModeForbidden with default headers values
func NewStartConstructorMaintenanceModeForbidden() *StartConstructorMaintenanceModeForbidden {
	return &StartConstructorMaintenanceModeForbidden{}
}

/*StartConstructorMaintenanceModeForbidden handles this case with default header values.

The start maintenance mode command was prohibited for the given constructor. (code: `constructors.command_prohibited`)
*/
type StartConstructorMaintenanceModeForbidden struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartConstructorMaintenanceModeForbidden) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/constructors/{constructor_id}/maintenance-mode/_start][%d] startConstructorMaintenanceModeForbidden  %+v", 403, o.Payload)
}

func (o *StartConstructorMaintenanceModeForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartConstructorMaintenanceModeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartConstructorMaintenanceModeNotFound creates a StartConstructorMaintenanceModeNotFound with default headers values
func NewStartConstructorMaintenanceModeNotFound() *StartConstructorMaintenanceModeNotFound {
	return &StartConstructorMaintenanceModeNotFound{}
}

/*StartConstructorMaintenanceModeNotFound handles this case with default header values.

The constructor specified by {constructor_id} cannot be found. (code: `constructors.constructor_not_found`)
*/
type StartConstructorMaintenanceModeNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartConstructorMaintenanceModeNotFound) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/constructors/{constructor_id}/maintenance-mode/_start][%d] startConstructorMaintenanceModeNotFound  %+v", 404, o.Payload)
}

func (o *StartConstructorMaintenanceModeNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartConstructorMaintenanceModeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartConstructorMaintenanceModeRetryWith creates a StartConstructorMaintenanceModeRetryWith with default headers values
func NewStartConstructorMaintenanceModeRetryWith() *StartConstructorMaintenanceModeRetryWith {
	return &StartConstructorMaintenanceModeRetryWith{}
}

/*StartConstructorMaintenanceModeRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type StartConstructorMaintenanceModeRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartConstructorMaintenanceModeRetryWith) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/constructors/{constructor_id}/maintenance-mode/_start][%d] startConstructorMaintenanceModeRetryWith  %+v", 449, o.Payload)
}

func (o *StartConstructorMaintenanceModeRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartConstructorMaintenanceModeRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
