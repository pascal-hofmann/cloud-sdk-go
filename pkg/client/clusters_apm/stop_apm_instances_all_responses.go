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

// StopApmInstancesAllReader is a Reader for the StopApmInstancesAll structure.
type StopApmInstancesAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopApmInstancesAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewStopApmInstancesAllAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewStopApmInstancesAllForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopApmInstancesAllNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewStopApmInstancesAllRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStopApmInstancesAllAccepted creates a StopApmInstancesAllAccepted with default headers values
func NewStopApmInstancesAllAccepted() *StopApmInstancesAllAccepted {
	return &StopApmInstancesAllAccepted{}
}

/*StopApmInstancesAllAccepted handles this case with default header values.

The stop command was issued successfully, use the "GET" command on the /{cluster_id} resource to monitor progress
*/
type StopApmInstancesAllAccepted struct {
	Payload *models.ClusterCommandResponse
}

func (o *StopApmInstancesAllAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/instances/_stop][%d] stopApmInstancesAllAccepted  %+v", 202, o.Payload)
}

func (o *StopApmInstancesAllAccepted) GetPayload() *models.ClusterCommandResponse {
	return o.Payload
}

func (o *StopApmInstancesAllAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopApmInstancesAllForbidden creates a StopApmInstancesAllForbidden with default headers values
func NewStopApmInstancesAllForbidden() *StopApmInstancesAllForbidden {
	return &StopApmInstancesAllForbidden{}
}

/*StopApmInstancesAllForbidden handles this case with default header values.

The stop command was prohibited for the given cluster. (code: `clusters.command_prohibited`)
*/
type StopApmInstancesAllForbidden struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StopApmInstancesAllForbidden) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/instances/_stop][%d] stopApmInstancesAllForbidden  %+v", 403, o.Payload)
}

func (o *StopApmInstancesAllForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StopApmInstancesAllForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopApmInstancesAllNotFound creates a StopApmInstancesAllNotFound with default headers values
func NewStopApmInstancesAllNotFound() *StopApmInstancesAllNotFound {
	return &StopApmInstancesAllNotFound{}
}

/*StopApmInstancesAllNotFound handles this case with default header values.

The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
*/
type StopApmInstancesAllNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StopApmInstancesAllNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/instances/_stop][%d] stopApmInstancesAllNotFound  %+v", 404, o.Payload)
}

func (o *StopApmInstancesAllNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StopApmInstancesAllNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopApmInstancesAllRetryWith creates a StopApmInstancesAllRetryWith with default headers values
func NewStopApmInstancesAllRetryWith() *StopApmInstancesAllRetryWith {
	return &StopApmInstancesAllRetryWith{}
}

/*StopApmInstancesAllRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type StopApmInstancesAllRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StopApmInstancesAllRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/instances/_stop][%d] stopApmInstancesAllRetryWith  %+v", 449, o.Payload)
}

func (o *StopApmInstancesAllRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StopApmInstancesAllRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
