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

package clusters_kibana

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// StartKibanaClusterMaintenanceModeReader is a Reader for the StartKibanaClusterMaintenanceMode structure.
type StartKibanaClusterMaintenanceModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartKibanaClusterMaintenanceModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewStartKibanaClusterMaintenanceModeAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewStartKibanaClusterMaintenanceModeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStartKibanaClusterMaintenanceModeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewStartKibanaClusterMaintenanceModeRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStartKibanaClusterMaintenanceModeAccepted creates a StartKibanaClusterMaintenanceModeAccepted with default headers values
func NewStartKibanaClusterMaintenanceModeAccepted() *StartKibanaClusterMaintenanceModeAccepted {
	return &StartKibanaClusterMaintenanceModeAccepted{}
}

/*StartKibanaClusterMaintenanceModeAccepted handles this case with default header values.

The start maintenance mode command was issued successfully, use the "GET" command on the /{cluster_id} resource to monitor progress
*/
type StartKibanaClusterMaintenanceModeAccepted struct {
	Payload *models.ClusterCommandResponse
}

func (o *StartKibanaClusterMaintenanceModeAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/instances/{instance_ids}/maintenance-mode/_start][%d] startKibanaClusterMaintenanceModeAccepted  %+v", 202, o.Payload)
}

func (o *StartKibanaClusterMaintenanceModeAccepted) GetPayload() *models.ClusterCommandResponse {
	return o.Payload
}

func (o *StartKibanaClusterMaintenanceModeAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartKibanaClusterMaintenanceModeForbidden creates a StartKibanaClusterMaintenanceModeForbidden with default headers values
func NewStartKibanaClusterMaintenanceModeForbidden() *StartKibanaClusterMaintenanceModeForbidden {
	return &StartKibanaClusterMaintenanceModeForbidden{}
}

/*StartKibanaClusterMaintenanceModeForbidden handles this case with default header values.

The start maintenance mode command was prohibited for the given cluster. (code: `clusters.command_prohibited`)
*/
type StartKibanaClusterMaintenanceModeForbidden struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartKibanaClusterMaintenanceModeForbidden) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/instances/{instance_ids}/maintenance-mode/_start][%d] startKibanaClusterMaintenanceModeForbidden  %+v", 403, o.Payload)
}

func (o *StartKibanaClusterMaintenanceModeForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartKibanaClusterMaintenanceModeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartKibanaClusterMaintenanceModeNotFound creates a StartKibanaClusterMaintenanceModeNotFound with default headers values
func NewStartKibanaClusterMaintenanceModeNotFound() *StartKibanaClusterMaintenanceModeNotFound {
	return &StartKibanaClusterMaintenanceModeNotFound{}
}

/*StartKibanaClusterMaintenanceModeNotFound handles this case with default header values.

* The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
* One or more of the instances specified at {instance_ids} could not be found. (code: `clusters.instances_not_found`)
 */
type StartKibanaClusterMaintenanceModeNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartKibanaClusterMaintenanceModeNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/instances/{instance_ids}/maintenance-mode/_start][%d] startKibanaClusterMaintenanceModeNotFound  %+v", 404, o.Payload)
}

func (o *StartKibanaClusterMaintenanceModeNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartKibanaClusterMaintenanceModeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartKibanaClusterMaintenanceModeRetryWith creates a StartKibanaClusterMaintenanceModeRetryWith with default headers values
func NewStartKibanaClusterMaintenanceModeRetryWith() *StartKibanaClusterMaintenanceModeRetryWith {
	return &StartKibanaClusterMaintenanceModeRetryWith{}
}

/*StartKibanaClusterMaintenanceModeRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type StartKibanaClusterMaintenanceModeRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartKibanaClusterMaintenanceModeRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/instances/{instance_ids}/maintenance-mode/_start][%d] startKibanaClusterMaintenanceModeRetryWith  %+v", 449, o.Payload)
}

func (o *StartKibanaClusterMaintenanceModeRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartKibanaClusterMaintenanceModeRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
