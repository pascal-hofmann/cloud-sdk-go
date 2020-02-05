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

package clusters_apm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// StartApmInstancesAllReader is a Reader for the StartApmInstancesAll structure.
type StartApmInstancesAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartApmInstancesAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewStartApmInstancesAllAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewStartApmInstancesAllForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStartApmInstancesAllNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewStartApmInstancesAllRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStartApmInstancesAllAccepted creates a StartApmInstancesAllAccepted with default headers values
func NewStartApmInstancesAllAccepted() *StartApmInstancesAllAccepted {
	return &StartApmInstancesAllAccepted{}
}

/*StartApmInstancesAllAccepted handles this case with default header values.

The start command was issued successfully, use the "GET" command on the /{cluster_id} resource to monitor progress
*/
type StartApmInstancesAllAccepted struct {
	Payload *models.ClusterCommandResponse
}

func (o *StartApmInstancesAllAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/instances/_start][%d] startApmInstancesAllAccepted  %+v", 202, o.Payload)
}

func (o *StartApmInstancesAllAccepted) GetPayload() *models.ClusterCommandResponse {
	return o.Payload
}

func (o *StartApmInstancesAllAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartApmInstancesAllForbidden creates a StartApmInstancesAllForbidden with default headers values
func NewStartApmInstancesAllForbidden() *StartApmInstancesAllForbidden {
	return &StartApmInstancesAllForbidden{}
}

/*StartApmInstancesAllForbidden handles this case with default header values.

The start command was prohibited for the given cluster. (code: `clusters.command_prohibited`)
*/
type StartApmInstancesAllForbidden struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartApmInstancesAllForbidden) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/instances/_start][%d] startApmInstancesAllForbidden  %+v", 403, o.Payload)
}

func (o *StartApmInstancesAllForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartApmInstancesAllForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartApmInstancesAllNotFound creates a StartApmInstancesAllNotFound with default headers values
func NewStartApmInstancesAllNotFound() *StartApmInstancesAllNotFound {
	return &StartApmInstancesAllNotFound{}
}

/*StartApmInstancesAllNotFound handles this case with default header values.

The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
*/
type StartApmInstancesAllNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartApmInstancesAllNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/instances/_start][%d] startApmInstancesAllNotFound  %+v", 404, o.Payload)
}

func (o *StartApmInstancesAllNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartApmInstancesAllNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartApmInstancesAllRetryWith creates a StartApmInstancesAllRetryWith with default headers values
func NewStartApmInstancesAllRetryWith() *StartApmInstancesAllRetryWith {
	return &StartApmInstancesAllRetryWith{}
}

/*StartApmInstancesAllRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type StartApmInstancesAllRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartApmInstancesAllRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/instances/_start][%d] startApmInstancesAllRetryWith  %+v", 449, o.Payload)
}

func (o *StartApmInstancesAllRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartApmInstancesAllRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
