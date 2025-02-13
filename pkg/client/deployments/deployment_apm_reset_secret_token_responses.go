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
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// DeploymentApmResetSecretTokenReader is a Reader for the DeploymentApmResetSecretToken structure.
type DeploymentApmResetSecretTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeploymentApmResetSecretTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeploymentApmResetSecretTokenAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeploymentApmResetSecretTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeploymentApmResetSecretTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewDeploymentApmResetSecretTokenRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeploymentApmResetSecretTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeploymentApmResetSecretTokenAccepted creates a DeploymentApmResetSecretTokenAccepted with default headers values
func NewDeploymentApmResetSecretTokenAccepted() *DeploymentApmResetSecretTokenAccepted {
	return &DeploymentApmResetSecretTokenAccepted{}
}

/* DeploymentApmResetSecretTokenAccepted describes a response with status code 202, with default header values.

Response containing the new secret token, plan to apply it starts
*/
type DeploymentApmResetSecretTokenAccepted struct {
	Payload *models.ApmCrudResponse
}

func (o *DeploymentApmResetSecretTokenAccepted) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/apm/{ref_id}/_reset-token][%d] deploymentApmResetSecretTokenAccepted  %+v", 202, o.Payload)
}
func (o *DeploymentApmResetSecretTokenAccepted) GetPayload() *models.ApmCrudResponse {
	return o.Payload
}

func (o *DeploymentApmResetSecretTokenAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApmCrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeploymentApmResetSecretTokenBadRequest creates a DeploymentApmResetSecretTokenBadRequest with default headers values
func NewDeploymentApmResetSecretTokenBadRequest() *DeploymentApmResetSecretTokenBadRequest {
	return &DeploymentApmResetSecretTokenBadRequest{}
}

/* DeploymentApmResetSecretTokenBadRequest describes a response with status code 400, with default header values.

Reset token is not supported when APM is managed by Elastic Agent. (code: `clusters.cluster_plan_state_error`)
*/
type DeploymentApmResetSecretTokenBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeploymentApmResetSecretTokenBadRequest) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/apm/{ref_id}/_reset-token][%d] deploymentApmResetSecretTokenBadRequest  %+v", 400, o.Payload)
}
func (o *DeploymentApmResetSecretTokenBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeploymentApmResetSecretTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeploymentApmResetSecretTokenNotFound creates a DeploymentApmResetSecretTokenNotFound with default headers values
func NewDeploymentApmResetSecretTokenNotFound() *DeploymentApmResetSecretTokenNotFound {
	return &DeploymentApmResetSecretTokenNotFound{}
}

/* DeploymentApmResetSecretTokenNotFound describes a response with status code 404, with default header values.

 * The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)
* The Resource specified by {ref_id} cannot be found. (code: `deployments.deployment_resource_not_found`)
*/
type DeploymentApmResetSecretTokenNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeploymentApmResetSecretTokenNotFound) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/apm/{ref_id}/_reset-token][%d] deploymentApmResetSecretTokenNotFound  %+v", 404, o.Payload)
}
func (o *DeploymentApmResetSecretTokenNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeploymentApmResetSecretTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeploymentApmResetSecretTokenRetryWith creates a DeploymentApmResetSecretTokenRetryWith with default headers values
func NewDeploymentApmResetSecretTokenRetryWith() *DeploymentApmResetSecretTokenRetryWith {
	return &DeploymentApmResetSecretTokenRetryWith{}
}

/* DeploymentApmResetSecretTokenRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type DeploymentApmResetSecretTokenRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeploymentApmResetSecretTokenRetryWith) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/apm/{ref_id}/_reset-token][%d] deploymentApmResetSecretTokenRetryWith  %+v", 449, o.Payload)
}
func (o *DeploymentApmResetSecretTokenRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeploymentApmResetSecretTokenRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeploymentApmResetSecretTokenInternalServerError creates a DeploymentApmResetSecretTokenInternalServerError with default headers values
func NewDeploymentApmResetSecretTokenInternalServerError() *DeploymentApmResetSecretTokenInternalServerError {
	return &DeploymentApmResetSecretTokenInternalServerError{}
}

/* DeploymentApmResetSecretTokenInternalServerError describes a response with status code 500, with default header values.

We have failed you. (code: `deployments.deployment_resource_no_longer_exists`)
*/
type DeploymentApmResetSecretTokenInternalServerError struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeploymentApmResetSecretTokenInternalServerError) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/apm/{ref_id}/_reset-token][%d] deploymentApmResetSecretTokenInternalServerError  %+v", 500, o.Payload)
}
func (o *DeploymentApmResetSecretTokenInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeploymentApmResetSecretTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
