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

package platform_configuration_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// CreateDeploymentTemplateReader is a Reader for the CreateDeploymentTemplate structure.
type CreateDeploymentTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDeploymentTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDeploymentTemplateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDeploymentTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewCreateDeploymentTemplateRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateDeploymentTemplateCreated creates a CreateDeploymentTemplateCreated with default headers values
func NewCreateDeploymentTemplateCreated() *CreateDeploymentTemplateCreated {
	return &CreateDeploymentTemplateCreated{}
}

/*CreateDeploymentTemplateCreated handles this case with default header values.

The deployment definition was valid and the template has been created.
*/
type CreateDeploymentTemplateCreated struct {
	Payload *models.IDResponse
}

func (o *CreateDeploymentTemplateCreated) Error() string {
	return fmt.Sprintf("[POST /platform/configuration/templates/deployments][%d] createDeploymentTemplateCreated  %+v", 201, o.Payload)
}

func (o *CreateDeploymentTemplateCreated) GetPayload() *models.IDResponse {
	return o.Payload
}

func (o *CreateDeploymentTemplateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDeploymentTemplateBadRequest creates a CreateDeploymentTemplateBadRequest with default headers values
func NewCreateDeploymentTemplateBadRequest() *CreateDeploymentTemplateBadRequest {
	return &CreateDeploymentTemplateBadRequest{}
}

/*CreateDeploymentTemplateBadRequest handles this case with default header values.

The template definition contained errors. (code: `templates.invalid_template`)
*/
type CreateDeploymentTemplateBadRequest struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *CreateDeploymentTemplateBadRequest) Error() string {
	return fmt.Sprintf("[POST /platform/configuration/templates/deployments][%d] createDeploymentTemplateBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDeploymentTemplateBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateDeploymentTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDeploymentTemplateRetryWith creates a CreateDeploymentTemplateRetryWith with default headers values
func NewCreateDeploymentTemplateRetryWith() *CreateDeploymentTemplateRetryWith {
	return &CreateDeploymentTemplateRetryWith{}
}

/*CreateDeploymentTemplateRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type CreateDeploymentTemplateRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *CreateDeploymentTemplateRetryWith) Error() string {
	return fmt.Sprintf("[POST /platform/configuration/templates/deployments][%d] createDeploymentTemplateRetryWith  %+v", 449, o.Payload)
}

func (o *CreateDeploymentTemplateRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateDeploymentTemplateRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
