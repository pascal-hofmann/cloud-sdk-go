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

// RestartEsClusterReader is a Reader for the RestartEsCluster structure.
type RestartEsClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestartEsClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewRestartEsClusterAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRestartEsClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRestartEsClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewRestartEsClusterPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewRestartEsClusterRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRestartEsClusterAccepted creates a RestartEsClusterAccepted with default headers values
func NewRestartEsClusterAccepted() *RestartEsClusterAccepted {
	return &RestartEsClusterAccepted{}
}

/*RestartEsClusterAccepted handles this case with default header values.

The stop command was issued successfully, use the "GET" command on the /{cluster_id} resource to monitor progress
*/
type RestartEsClusterAccepted struct {
	Payload *models.ClusterCommandResponse
}

func (o *RestartEsClusterAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/_restart][%d] restartEsClusterAccepted  %+v", 202, o.Payload)
}

func (o *RestartEsClusterAccepted) GetPayload() *models.ClusterCommandResponse {
	return o.Payload
}

func (o *RestartEsClusterAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestartEsClusterBadRequest creates a RestartEsClusterBadRequest with default headers values
func NewRestartEsClusterBadRequest() *RestartEsClusterBadRequest {
	return &RestartEsClusterBadRequest{}
}

/*RestartEsClusterBadRequest handles this case with default header values.

The cluster specified by {cluster_id} is unable to restart. (code: `clusters.restart.failed`)
*/
type RestartEsClusterBadRequest struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *RestartEsClusterBadRequest) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/_restart][%d] restartEsClusterBadRequest  %+v", 400, o.Payload)
}

func (o *RestartEsClusterBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RestartEsClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestartEsClusterNotFound creates a RestartEsClusterNotFound with default headers values
func NewRestartEsClusterNotFound() *RestartEsClusterNotFound {
	return &RestartEsClusterNotFound{}
}

/*RestartEsClusterNotFound handles this case with default header values.

The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
*/
type RestartEsClusterNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *RestartEsClusterNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/_restart][%d] restartEsClusterNotFound  %+v", 404, o.Payload)
}

func (o *RestartEsClusterNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RestartEsClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestartEsClusterPreconditionFailed creates a RestartEsClusterPreconditionFailed with default headers values
func NewRestartEsClusterPreconditionFailed() *RestartEsClusterPreconditionFailed {
	return &RestartEsClusterPreconditionFailed{}
}

/*RestartEsClusterPreconditionFailed handles this case with default header values.

The command sent to a cluster found the cluster in an illegal state, the error message gives more details. (code: `clusters.cluster_plan_state_error`)
*/
type RestartEsClusterPreconditionFailed struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *RestartEsClusterPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/_restart][%d] restartEsClusterPreconditionFailed  %+v", 412, o.Payload)
}

func (o *RestartEsClusterPreconditionFailed) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RestartEsClusterPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestartEsClusterRetryWith creates a RestartEsClusterRetryWith with default headers values
func NewRestartEsClusterRetryWith() *RestartEsClusterRetryWith {
	return &RestartEsClusterRetryWith{}
}

/*RestartEsClusterRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type RestartEsClusterRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *RestartEsClusterRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/_restart][%d] restartEsClusterRetryWith  %+v", 449, o.Payload)
}

func (o *RestartEsClusterRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RestartEsClusterRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
