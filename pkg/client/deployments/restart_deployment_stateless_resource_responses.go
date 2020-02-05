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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// RestartDeploymentStatelessResourceReader is a Reader for the RestartDeploymentStatelessResource structure.
type RestartDeploymentStatelessResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestartDeploymentStatelessResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewRestartDeploymentStatelessResourceAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRestartDeploymentStatelessResourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewRestartDeploymentStatelessResourceUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewRestartDeploymentStatelessResourceRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRestartDeploymentStatelessResourceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRestartDeploymentStatelessResourceAccepted creates a RestartDeploymentStatelessResourceAccepted with default headers values
func NewRestartDeploymentStatelessResourceAccepted() *RestartDeploymentStatelessResourceAccepted {
	return &RestartDeploymentStatelessResourceAccepted{}
}

/*RestartDeploymentStatelessResourceAccepted handles this case with default header values.

The restart command was issued successfully
*/
type RestartDeploymentStatelessResourceAccepted struct {
	Payload models.DeploymentResourceCommandResponse
}

func (o *RestartDeploymentStatelessResourceAccepted) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{stateless_resource_kind}/{ref_id}/_restart][%d] restartDeploymentStatelessResourceAccepted  %+v", 202, o.Payload)
}

func (o *RestartDeploymentStatelessResourceAccepted) GetPayload() models.DeploymentResourceCommandResponse {
	return o.Payload
}

func (o *RestartDeploymentStatelessResourceAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestartDeploymentStatelessResourceNotFound creates a RestartDeploymentStatelessResourceNotFound with default headers values
func NewRestartDeploymentStatelessResourceNotFound() *RestartDeploymentStatelessResourceNotFound {
	return &RestartDeploymentStatelessResourceNotFound{}
}

/*RestartDeploymentStatelessResourceNotFound handles this case with default header values.

The Resource specified by {ref_id} cannot be found. (code: `deployments.deployment_resource_not_found`)
*/
type RestartDeploymentStatelessResourceNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *RestartDeploymentStatelessResourceNotFound) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{stateless_resource_kind}/{ref_id}/_restart][%d] restartDeploymentStatelessResourceNotFound  %+v", 404, o.Payload)
}

func (o *RestartDeploymentStatelessResourceNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RestartDeploymentStatelessResourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestartDeploymentStatelessResourceUnprocessableEntity creates a RestartDeploymentStatelessResourceUnprocessableEntity with default headers values
func NewRestartDeploymentStatelessResourceUnprocessableEntity() *RestartDeploymentStatelessResourceUnprocessableEntity {
	return &RestartDeploymentStatelessResourceUnprocessableEntity{}
}

/*RestartDeploymentStatelessResourceUnprocessableEntity handles this case with default header values.

The command sent to a Resource found the Resource in an illegal state, the error message gives more details. (code: `deployments.deployment_resource_plan_change_error`)
*/
type RestartDeploymentStatelessResourceUnprocessableEntity struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *RestartDeploymentStatelessResourceUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{stateless_resource_kind}/{ref_id}/_restart][%d] restartDeploymentStatelessResourceUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *RestartDeploymentStatelessResourceUnprocessableEntity) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RestartDeploymentStatelessResourceUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestartDeploymentStatelessResourceRetryWith creates a RestartDeploymentStatelessResourceRetryWith with default headers values
func NewRestartDeploymentStatelessResourceRetryWith() *RestartDeploymentStatelessResourceRetryWith {
	return &RestartDeploymentStatelessResourceRetryWith{}
}

/*RestartDeploymentStatelessResourceRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type RestartDeploymentStatelessResourceRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *RestartDeploymentStatelessResourceRetryWith) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{stateless_resource_kind}/{ref_id}/_restart][%d] restartDeploymentStatelessResourceRetryWith  %+v", 449, o.Payload)
}

func (o *RestartDeploymentStatelessResourceRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RestartDeploymentStatelessResourceRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestartDeploymentStatelessResourceInternalServerError creates a RestartDeploymentStatelessResourceInternalServerError with default headers values
func NewRestartDeploymentStatelessResourceInternalServerError() *RestartDeploymentStatelessResourceInternalServerError {
	return &RestartDeploymentStatelessResourceInternalServerError{}
}

/*RestartDeploymentStatelessResourceInternalServerError handles this case with default header values.

We have failed you. (code: `deployments.deployment_resource_no_longer_exists`)
*/
type RestartDeploymentStatelessResourceInternalServerError struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *RestartDeploymentStatelessResourceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{stateless_resource_kind}/{ref_id}/_restart][%d] restartDeploymentStatelessResourceInternalServerError  %+v", 500, o.Payload)
}

func (o *RestartDeploymentStatelessResourceInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RestartDeploymentStatelessResourceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
