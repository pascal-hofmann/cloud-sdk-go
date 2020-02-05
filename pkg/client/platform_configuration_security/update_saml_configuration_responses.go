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

package platform_configuration_security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// UpdateSamlConfigurationReader is a Reader for the UpdateSamlConfiguration structure.
type UpdateSamlConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSamlConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSamlConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateSamlConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateSamlConfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateSamlConfigurationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewUpdateSamlConfigurationRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateSamlConfigurationOK creates a UpdateSamlConfigurationOK with default headers values
func NewUpdateSamlConfigurationOK() *UpdateSamlConfigurationOK {
	return &UpdateSamlConfigurationOK{}
}

/*UpdateSamlConfigurationOK handles this case with default header values.

The SAML configuration was successfully updated
*/
type UpdateSamlConfigurationOK struct {
	/*The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string
	/*The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string
	/*The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload models.EmptyResponse
}

func (o *UpdateSamlConfigurationOK) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/realms/saml/{realm_id}][%d] updateSamlConfigurationOK  %+v", 200, o.Payload)
}

func (o *UpdateSamlConfigurationOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *UpdateSamlConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-resource-created
	o.XCloudResourceCreated = response.GetHeader("x-cloud-resource-created")

	// response header x-cloud-resource-last-modified
	o.XCloudResourceLastModified = response.GetHeader("x-cloud-resource-last-modified")

	// response header x-cloud-resource-version
	o.XCloudResourceVersion = response.GetHeader("x-cloud-resource-version")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSamlConfigurationBadRequest creates a UpdateSamlConfigurationBadRequest with default headers values
func NewUpdateSamlConfigurationBadRequest() *UpdateSamlConfigurationBadRequest {
	return &UpdateSamlConfigurationBadRequest{}
}

/*UpdateSamlConfigurationBadRequest handles this case with default header values.

* The realm id is already in use. (code: `security_realm.id_conflict`)
* The selected id is not valid. (code: `security_realm.invalid_id`)
* Order must be greater than zero. (code: `security_realm.invalid_order`)
* Invalid Elasticsearch Security realm type. (code: `security_realm.invalid_type`)
* The realm order is already in use. (code: `security_realm.order_conflict`)
* Advanced YAML format is invalid. (code: `security_realm.invalid_yaml`)
* The SAML IDP metadata endpoint returned an error response code 200 OK. (code: `security_realm.saml.invalid_idp_metadata_url`)
* Invalid certificate bundle URL. (code: `security_realm.invalid_bundle_url`)
 */
type UpdateSamlConfigurationBadRequest struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateSamlConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/realms/saml/{realm_id}][%d] updateSamlConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSamlConfigurationBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateSamlConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSamlConfigurationNotFound creates a UpdateSamlConfigurationNotFound with default headers values
func NewUpdateSamlConfigurationNotFound() *UpdateSamlConfigurationNotFound {
	return &UpdateSamlConfigurationNotFound{}
}

/*UpdateSamlConfigurationNotFound handles this case with default header values.

The realm specified by {realm_id} cannot be found. (code: `security_realm.not_found`)
*/
type UpdateSamlConfigurationNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateSamlConfigurationNotFound) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/realms/saml/{realm_id}][%d] updateSamlConfigurationNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSamlConfigurationNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateSamlConfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSamlConfigurationConflict creates a UpdateSamlConfigurationConflict with default headers values
func NewUpdateSamlConfigurationConflict() *UpdateSamlConfigurationConflict {
	return &UpdateSamlConfigurationConflict{}
}

/*UpdateSamlConfigurationConflict handles this case with default header values.

There is a version conflict. (code: `security_realm.version_conflict`)
*/
type UpdateSamlConfigurationConflict struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateSamlConfigurationConflict) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/realms/saml/{realm_id}][%d] updateSamlConfigurationConflict  %+v", 409, o.Payload)
}

func (o *UpdateSamlConfigurationConflict) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateSamlConfigurationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSamlConfigurationRetryWith creates a UpdateSamlConfigurationRetryWith with default headers values
func NewUpdateSamlConfigurationRetryWith() *UpdateSamlConfigurationRetryWith {
	return &UpdateSamlConfigurationRetryWith{}
}

/*UpdateSamlConfigurationRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type UpdateSamlConfigurationRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateSamlConfigurationRetryWith) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/realms/saml/{realm_id}][%d] updateSamlConfigurationRetryWith  %+v", 449, o.Payload)
}

func (o *UpdateSamlConfigurationRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateSamlConfigurationRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
