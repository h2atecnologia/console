// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2020 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// KesConfiguration kes configuration
//
// swagger:model kesConfiguration
type KesConfiguration struct {

	// client
	// Required: true
	Client *KesConfigurationClient `json:"client"`

	// server
	// Required: true
	Server *KesConfigurationServer `json:"server"`

	// server config yaml
	// Required: true
	ServerConfigYaml *string `json:"server-config.yaml"`
}

// Validate validates this kes configuration
func (m *KesConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerConfigYaml(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KesConfiguration) validateClient(formats strfmt.Registry) error {

	if err := validate.Required("client", "body", m.Client); err != nil {
		return err
	}

	if m.Client != nil {
		if err := m.Client.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client")
			}
			return err
		}
	}

	return nil
}

func (m *KesConfiguration) validateServer(formats strfmt.Registry) error {

	if err := validate.Required("server", "body", m.Server); err != nil {
		return err
	}

	if m.Server != nil {
		if err := m.Server.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server")
			}
			return err
		}
	}

	return nil
}

func (m *KesConfiguration) validateServerConfigYaml(formats strfmt.Registry) error {

	if err := validate.Required("server-config.yaml", "body", m.ServerConfigYaml); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KesConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KesConfiguration) UnmarshalBinary(b []byte) error {
	var res KesConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KesConfigurationClient kes configuration client
//
// swagger:model KesConfigurationClient
type KesConfigurationClient struct {

	// tls crt
	// Required: true
	TLSCrt *string `json:"tls.crt"`

	// tls key
	// Required: true
	TLSKey *string `json:"tls.key"`
}

// Validate validates this kes configuration client
func (m *KesConfigurationClient) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTLSCrt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTLSKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KesConfigurationClient) validateTLSCrt(formats strfmt.Registry) error {

	if err := validate.Required("client"+"."+"tls.crt", "body", m.TLSCrt); err != nil {
		return err
	}

	return nil
}

func (m *KesConfigurationClient) validateTLSKey(formats strfmt.Registry) error {

	if err := validate.Required("client"+"."+"tls.key", "body", m.TLSKey); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KesConfigurationClient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KesConfigurationClient) UnmarshalBinary(b []byte) error {
	var res KesConfigurationClient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KesConfigurationServer kes configuration server
//
// swagger:model KesConfigurationServer
type KesConfigurationServer struct {

	// tls crt
	// Required: true
	TLSCrt *string `json:"tls.crt"`

	// tls key
	// Required: true
	TLSKey *string `json:"tls.key"`
}

// Validate validates this kes configuration server
func (m *KesConfigurationServer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTLSCrt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTLSKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KesConfigurationServer) validateTLSCrt(formats strfmt.Registry) error {

	if err := validate.Required("server"+"."+"tls.crt", "body", m.TLSCrt); err != nil {
		return err
	}

	return nil
}

func (m *KesConfigurationServer) validateTLSKey(formats strfmt.Registry) error {

	if err := validate.Required("server"+"."+"tls.key", "body", m.TLSKey); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KesConfigurationServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KesConfigurationServer) UnmarshalBinary(b []byte) error {
	var res KesConfigurationServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
