// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProjectEnvVarPayload project env var payload
//
// swagger:model ProjectEnvVarPayload
type ProjectEnvVarPayload struct {

	// The name of the environment variable.
	// Required: true
	Name *string `json:"name"`

	// The value of the environment variable.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this project env var payload
func (m *ProjectEnvVarPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectEnvVarPayload) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ProjectEnvVarPayload) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this project env var payload based on context it is used
func (m *ProjectEnvVarPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectEnvVarPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectEnvVarPayload) UnmarshalBinary(b []byte) error {
	var res ProjectEnvVarPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}