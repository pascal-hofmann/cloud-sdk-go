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

package clusters_elasticsearch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// StartEsClusterInstancesAllReader is a Reader for the StartEsClusterInstancesAll structure.
type StartEsClusterInstancesAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartEsClusterInstancesAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewStartEsClusterInstancesAllAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewStartEsClusterInstancesAllForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStartEsClusterInstancesAllNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewStartEsClusterInstancesAllRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStartEsClusterInstancesAllAccepted creates a StartEsClusterInstancesAllAccepted with default headers values
func NewStartEsClusterInstancesAllAccepted() *StartEsClusterInstancesAllAccepted {
	return &StartEsClusterInstancesAllAccepted{}
}

/*StartEsClusterInstancesAllAccepted handles this case with default header values.

The start command was issued successfully, use the "GET" command on the /{cluster_id} resource to monitor progress
*/
type StartEsClusterInstancesAllAccepted struct {
	Payload *models.ClusterCommandResponse
}

func (o *StartEsClusterInstancesAllAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/instances/_start][%d] startEsClusterInstancesAllAccepted  %+v", 202, o.Payload)
}

func (o *StartEsClusterInstancesAllAccepted) GetPayload() *models.ClusterCommandResponse {
	return o.Payload
}

func (o *StartEsClusterInstancesAllAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartEsClusterInstancesAllForbidden creates a StartEsClusterInstancesAllForbidden with default headers values
func NewStartEsClusterInstancesAllForbidden() *StartEsClusterInstancesAllForbidden {
	return &StartEsClusterInstancesAllForbidden{}
}

/*StartEsClusterInstancesAllForbidden handles this case with default header values.

The start command was prohibited for the given cluster. (code: `clusters.command_prohibited`)
*/
type StartEsClusterInstancesAllForbidden struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartEsClusterInstancesAllForbidden) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/instances/_start][%d] startEsClusterInstancesAllForbidden  %+v", 403, o.Payload)
}

func (o *StartEsClusterInstancesAllForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartEsClusterInstancesAllForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartEsClusterInstancesAllNotFound creates a StartEsClusterInstancesAllNotFound with default headers values
func NewStartEsClusterInstancesAllNotFound() *StartEsClusterInstancesAllNotFound {
	return &StartEsClusterInstancesAllNotFound{}
}

/*StartEsClusterInstancesAllNotFound handles this case with default header values.

The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
*/
type StartEsClusterInstancesAllNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartEsClusterInstancesAllNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/instances/_start][%d] startEsClusterInstancesAllNotFound  %+v", 404, o.Payload)
}

func (o *StartEsClusterInstancesAllNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartEsClusterInstancesAllNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartEsClusterInstancesAllRetryWith creates a StartEsClusterInstancesAllRetryWith with default headers values
func NewStartEsClusterInstancesAllRetryWith() *StartEsClusterInstancesAllRetryWith {
	return &StartEsClusterInstancesAllRetryWith{}
}

/*StartEsClusterInstancesAllRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type StartEsClusterInstancesAllRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartEsClusterInstancesAllRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/instances/_start][%d] startEsClusterInstancesAllRetryWith  %+v", 449, o.Payload)
}

func (o *StartEsClusterInstancesAllRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartEsClusterInstancesAllRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
