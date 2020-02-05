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

// GetApmMetadataRawReader is a Reader for the GetApmMetadataRaw structure.
type GetApmMetadataRawReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApmMetadataRawReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApmMetadataRawOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetApmMetadataRawNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApmMetadataRawOK creates a GetApmMetadataRawOK with default headers values
func NewGetApmMetadataRawOK() *GetApmMetadataRawOK {
	return &GetApmMetadataRawOK{}
}

/*GetApmMetadataRawOK handles this case with default header values.

The cluster metadata was successfully returned
*/
type GetApmMetadataRawOK struct {
	Payload interface{}
}

func (o *GetApmMetadataRawOK) Error() string {
	return fmt.Sprintf("[GET /clusters/apm/{cluster_id}/metadata/raw][%d] getApmMetadataRawOK  %+v", 200, o.Payload)
}

func (o *GetApmMetadataRawOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetApmMetadataRawOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApmMetadataRawNotFound creates a GetApmMetadataRawNotFound with default headers values
func NewGetApmMetadataRawNotFound() *GetApmMetadataRawNotFound {
	return &GetApmMetadataRawNotFound{}
}

/*GetApmMetadataRawNotFound handles this case with default header values.

The cluster specified by {cluster_id} cannot be found (code: 'clusters.cluster_not_found')
*/
type GetApmMetadataRawNotFound struct {
	Payload *models.BasicFailedReply
}

func (o *GetApmMetadataRawNotFound) Error() string {
	return fmt.Sprintf("[GET /clusters/apm/{cluster_id}/metadata/raw][%d] getApmMetadataRawNotFound  %+v", 404, o.Payload)
}

func (o *GetApmMetadataRawNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetApmMetadataRawNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
