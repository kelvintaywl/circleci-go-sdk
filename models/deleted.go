// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Deleted deleted
//
// swagger:model Deleted
type Deleted struct {

	// A human-readable message
	Message string `json:"message,omitempty"`
}

// Validate validates this deleted
func (m *Deleted) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this deleted based on context it is used
func (m *Deleted) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Deleted) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Deleted) UnmarshalBinary(b []byte) error {
	var res Deleted
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
