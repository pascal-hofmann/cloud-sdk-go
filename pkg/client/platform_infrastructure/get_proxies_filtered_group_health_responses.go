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

// GetProxiesFilteredGroupHealthReader is a Reader for the GetProxiesFilteredGroupHealth structure.
type GetProxiesFilteredGroupHealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProxiesFilteredGroupHealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProxiesFilteredGroupHealthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 417:
		result := NewGetProxiesFilteredGroupHealthExpectationFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewGetProxiesFilteredGroupHealthRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProxiesFilteredGroupHealthOK creates a GetProxiesFilteredGroupHealthOK with default headers values
func NewGetProxiesFilteredGroupHealthOK() *GetProxiesFilteredGroupHealthOK {
	return &GetProxiesFilteredGroupHealthOK{}
}

/*GetProxiesFilteredGroupHealthOK handles this case with default header values.

Returns health information on a filtered group of proxies
*/
type GetProxiesFilteredGroupHealthOK struct {
	Payload *models.ProxiesFilteredGroupHealth
}

func (o *GetProxiesFilteredGroupHealthOK) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/proxies/filtered-groups/{proxies_filtered_group_id}/health][%d] getProxiesFilteredGroupHealthOK  %+v", 200, o.Payload)
}

func (o *GetProxiesFilteredGroupHealthOK) GetPayload() *models.ProxiesFilteredGroupHealth {
	return o.Payload
}

func (o *GetProxiesFilteredGroupHealthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxiesFilteredGroupHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProxiesFilteredGroupHealthExpectationFailed creates a GetProxiesFilteredGroupHealthExpectationFailed with default headers values
func NewGetProxiesFilteredGroupHealthExpectationFailed() *GetProxiesFilteredGroupHealthExpectationFailed {
	return &GetProxiesFilteredGroupHealthExpectationFailed{}
}

/*GetProxiesFilteredGroupHealthExpectationFailed handles this case with default header values.

The health status is worse than the expected one.
*/
type GetProxiesFilteredGroupHealthExpectationFailed struct {
	Payload *models.ProxiesHealth
}

func (o *GetProxiesFilteredGroupHealthExpectationFailed) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/proxies/filtered-groups/{proxies_filtered_group_id}/health][%d] getProxiesFilteredGroupHealthExpectationFailed  %+v", 417, o.Payload)
}

func (o *GetProxiesFilteredGroupHealthExpectationFailed) GetPayload() *models.ProxiesHealth {
	return o.Payload
}

func (o *GetProxiesFilteredGroupHealthExpectationFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxiesHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProxiesFilteredGroupHealthRetryWith creates a GetProxiesFilteredGroupHealthRetryWith with default headers values
func NewGetProxiesFilteredGroupHealthRetryWith() *GetProxiesFilteredGroupHealthRetryWith {
	return &GetProxiesFilteredGroupHealthRetryWith{}
}

/*GetProxiesFilteredGroupHealthRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type GetProxiesFilteredGroupHealthRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetProxiesFilteredGroupHealthRetryWith) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/proxies/filtered-groups/{proxies_filtered_group_id}/health][%d] getProxiesFilteredGroupHealthRetryWith  %+v", 449, o.Payload)
}

func (o *GetProxiesFilteredGroupHealthRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetProxiesFilteredGroupHealthRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
