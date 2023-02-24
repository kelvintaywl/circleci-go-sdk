// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProjectCheckoutKeyPayload project checkout key payload
//
// swagger:model ProjectCheckoutKeyPayload
type ProjectCheckoutKeyPayload struct {

	// type
	// Example: deploy-key
	// Enum: [user-key deploy-key]
	Type string `json:"type,omitempty"`
}

// Validate validates this project checkout key payload
func (m *ProjectCheckoutKeyPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var projectCheckoutKeyPayloadTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["user-key","deploy-key"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		projectCheckoutKeyPayloadTypeTypePropEnum = append(projectCheckoutKeyPayloadTypeTypePropEnum, v)
	}
}

const (

	// ProjectCheckoutKeyPayloadTypeUserDashKey captures enum value "user-key"
	ProjectCheckoutKeyPayloadTypeUserDashKey string = "user-key"

	// ProjectCheckoutKeyPayloadTypeDeployDashKey captures enum value "deploy-key"
	ProjectCheckoutKeyPayloadTypeDeployDashKey string = "deploy-key"
)

// prop value enum
func (m *ProjectCheckoutKeyPayload) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, projectCheckoutKeyPayloadTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProjectCheckoutKeyPayload) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this project checkout key payload based on context it is used
func (m *ProjectCheckoutKeyPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectCheckoutKeyPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectCheckoutKeyPayload) UnmarshalBinary(b []byte) error {
	var res ProjectCheckoutKeyPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
