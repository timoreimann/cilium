// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceSpec Configuration of a service
// swagger:model ServiceSpec
type ServiceSpec struct {

	// List of backend addresses
	BackendAddresses []*BackendAddress `json:"backend-addresses"`

	// flags
	Flags *ServiceSpecFlags `json:"flags,omitempty"`

	// Frontend address
	// Required: true
	FrontendAddress *FrontendAddress `json:"frontend-address"`

	// Unique identification
	ID int64 `json:"id,omitempty"`
}

// Validate validates this service spec
func (m *ServiceSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackendAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrontendAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceSpec) validateBackendAddresses(formats strfmt.Registry) error {

	if swag.IsZero(m.BackendAddresses) { // not required
		return nil
	}

	for i := 0; i < len(m.BackendAddresses); i++ {
		if swag.IsZero(m.BackendAddresses[i]) { // not required
			continue
		}

		if m.BackendAddresses[i] != nil {
			if err := m.BackendAddresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backend-addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceSpec) validateFlags(formats strfmt.Registry) error {

	if swag.IsZero(m.Flags) { // not required
		return nil
	}

	if m.Flags != nil {
		if err := m.Flags.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flags")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceSpec) validateFrontendAddress(formats strfmt.Registry) error {

	if err := validate.Required("frontend-address", "body", m.FrontendAddress); err != nil {
		return err
	}

	if m.FrontendAddress != nil {
		if err := m.FrontendAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("frontend-address")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceSpec) UnmarshalBinary(b []byte) error {
	var res ServiceSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceSpecFlags Optional service configuration flags
// swagger:model ServiceSpecFlags
type ServiceSpecFlags struct {

	// Frontend to backend translation activated
	ActiveFrontend bool `json:"active-frontend,omitempty"`

	// Perform direct server return
	DirectServerReturn bool `json:"direct-server-return,omitempty"`

	// Service is of ExternalIPs type
	ExternalIP bool `json:"external-ip,omitempty"`

	// Service is of Nodeport type
	NodePort bool `json:"node-port,omitempty"`
}

// Validate validates this service spec flags
func (m *ServiceSpecFlags) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceSpecFlags) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceSpecFlags) UnmarshalBinary(b []byte) error {
	var res ServiceSpecFlags
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
