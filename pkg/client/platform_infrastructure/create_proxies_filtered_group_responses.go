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

// CreateProxiesFilteredGroupReader is a Reader for the CreateProxiesFilteredGroup structure.
type CreateProxiesFilteredGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProxiesFilteredGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateProxiesFilteredGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateProxiesFilteredGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewCreateProxiesFilteredGroupRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateProxiesFilteredGroupOK creates a CreateProxiesFilteredGroupOK with default headers values
func NewCreateProxiesFilteredGroupOK() *CreateProxiesFilteredGroupOK {
	return &CreateProxiesFilteredGroupOK{}
}

/*CreateProxiesFilteredGroupOK handles this case with default header values.

Returns the created or updated filtered group of proxies
*/
type CreateProxiesFilteredGroupOK struct {
	/*The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string
	/*The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string
	/*The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload *models.ProxiesFilteredGroup
}

func (o *CreateProxiesFilteredGroupOK) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/proxies/filtered-groups][%d] createProxiesFilteredGroupOK  %+v", 200, o.Payload)
}

func (o *CreateProxiesFilteredGroupOK) GetPayload() *models.ProxiesFilteredGroup {
	return o.Payload
}

func (o *CreateProxiesFilteredGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-resource-created
	o.XCloudResourceCreated = response.GetHeader("x-cloud-resource-created")

	// response header x-cloud-resource-last-modified
	o.XCloudResourceLastModified = response.GetHeader("x-cloud-resource-last-modified")

	// response header x-cloud-resource-version
	o.XCloudResourceVersion = response.GetHeader("x-cloud-resource-version")

	o.Payload = new(models.ProxiesFilteredGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProxiesFilteredGroupBadRequest creates a CreateProxiesFilteredGroupBadRequest with default headers values
func NewCreateProxiesFilteredGroupBadRequest() *CreateProxiesFilteredGroupBadRequest {
	return &CreateProxiesFilteredGroupBadRequest{}
}

/*CreateProxiesFilteredGroupBadRequest handles this case with default header values.

* The filtered group of proxies has empty id. (code: `proxies.proxies_filtered_group_empty_id`)
* A filtered group of proxies with the same identifier already exists. (code: `proxies.proxies_filtered_group_already_exists`)
 */
type CreateProxiesFilteredGroupBadRequest struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *CreateProxiesFilteredGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/proxies/filtered-groups][%d] createProxiesFilteredGroupBadRequest  %+v", 400, o.Payload)
}

func (o *CreateProxiesFilteredGroupBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateProxiesFilteredGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProxiesFilteredGroupRetryWith creates a CreateProxiesFilteredGroupRetryWith with default headers values
func NewCreateProxiesFilteredGroupRetryWith() *CreateProxiesFilteredGroupRetryWith {
	return &CreateProxiesFilteredGroupRetryWith{}
}

/*CreateProxiesFilteredGroupRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type CreateProxiesFilteredGroupRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *CreateProxiesFilteredGroupRetryWith) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/proxies/filtered-groups][%d] createProxiesFilteredGroupRetryWith  %+v", 449, o.Payload)
}

func (o *CreateProxiesFilteredGroupRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateProxiesFilteredGroupRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
