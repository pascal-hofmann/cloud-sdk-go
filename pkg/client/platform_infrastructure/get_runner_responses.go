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

// GetRunnerReader is a Reader for the GetRunner structure.
type GetRunnerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunnerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunnerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRunnerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRunnerOK creates a GetRunnerOK with default headers values
func NewGetRunnerOK() *GetRunnerOK {
	return &GetRunnerOK{}
}

/*GetRunnerOK handles this case with default header values.

The information for the runner specified by {runner_id}.
*/
type GetRunnerOK struct {
	Payload *models.RunnerInfo
}

func (o *GetRunnerOK) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/runners/{runner_id}][%d] getRunnerOK  %+v", 200, o.Payload)
}

func (o *GetRunnerOK) GetPayload() *models.RunnerInfo {
	return o.Payload
}

func (o *GetRunnerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RunnerInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunnerNotFound creates a GetRunnerNotFound with default headers values
func NewGetRunnerNotFound() *GetRunnerNotFound {
	return &GetRunnerNotFound{}
}

/*GetRunnerNotFound handles this case with default header values.

Unable to find the {runner_id} specified runner. Edit your request, then try again. (code: `runners.runner_not_found`)
*/
type GetRunnerNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetRunnerNotFound) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/runners/{runner_id}][%d] getRunnerNotFound  %+v", 404, o.Payload)
}

func (o *GetRunnerNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetRunnerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
