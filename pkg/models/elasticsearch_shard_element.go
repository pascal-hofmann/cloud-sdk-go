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
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ElasticsearchShardElement Information about the shards for each Elasticsearch instance container that hosts an Elasticsearch node. TIP: When the shard is unavailable, the cluster is unable to serve all of the data.
// swagger:model ElasticsearchShardElement
type ElasticsearchShardElement struct {

	// The Elastic Cloud name/id of the instance (container)
	// Required: true
	InstanceName *string `json:"instance_name"`

	// The number of shards of the given type (available/unavailable) on this instance
	// Required: true
	ShardCount *int32 `json:"shard_count"`
}

// Validate validates this elasticsearch shard element
func (m *ElasticsearchShardElement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShardCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElasticsearchShardElement) validateInstanceName(formats strfmt.Registry) error {

	if err := validate.Required("instance_name", "body", m.InstanceName); err != nil {
		return err
	}

	return nil
}

func (m *ElasticsearchShardElement) validateShardCount(formats strfmt.Registry) error {

	if err := validate.Required("shard_count", "body", m.ShardCount); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ElasticsearchShardElement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElasticsearchShardElement) UnmarshalBinary(b []byte) error {
	var res ElasticsearchShardElement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
