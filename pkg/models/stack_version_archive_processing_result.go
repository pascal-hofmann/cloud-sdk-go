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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StackVersionArchiveProcessingResult The result from processing an archive of an Elastic Stack version.
// swagger:model StackVersionArchiveProcessingResult
type StackVersionArchiveProcessingResult struct {

	// Errors occurred while processing the Elastic Stack pack. Key: stack version, Value: List of errors
	// Required: true
	Errors []*StackVersionArchiveProcessingError `json:"errors"`

	// stacks
	// Required: true
	Stacks []*StackVersionConfig `json:"stacks"`
}

// Validate validates this stack version archive processing result
func (m *StackVersionArchiveProcessingResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStacks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackVersionArchiveProcessingResult) validateErrors(formats strfmt.Registry) error {

	if err := validate.Required("errors", "body", m.Errors); err != nil {
		return err
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StackVersionArchiveProcessingResult) validateStacks(formats strfmt.Registry) error {

	if err := validate.Required("stacks", "body", m.Stacks); err != nil {
		return err
	}

	for i := 0; i < len(m.Stacks); i++ {
		if swag.IsZero(m.Stacks[i]) { // not required
			continue
		}

		if m.Stacks[i] != nil {
			if err := m.Stacks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackVersionArchiveProcessingResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackVersionArchiveProcessingResult) UnmarshalBinary(b []byte) error {
	var res StackVersionArchiveProcessingResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
