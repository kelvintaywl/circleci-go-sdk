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

// OrgLevelClaimsInfo org level claims info
//
// swagger:model OrgLevelClaimsInfo
type OrgLevelClaimsInfo struct {

	// List of audience.
	Audience []string `json:"audience"`

	// The date and time the audience list was updated.
	// Format: date-time
	AudienceUpdatedAt strfmt.DateTime `json:"audience_updated_at,omitempty"`

	// The unique ID of the organization.
	// Format: uuid
	OrgID strfmt.UUID `json:"org_id,omitempty"`

	// Time to live (JSON duration).
	// Example: 20d
	TTL string `json:"ttl,omitempty"`

	// The date and time the TTL was updated.
	// Format: date-time
	TTLUpdatedAt strfmt.DateTime `json:"ttl_updated_at,omitempty"`
}

// Validate validates this org level claims info
func (m *OrgLevelClaimsInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAudienceUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrgID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTTLUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrgLevelClaimsInfo) validateAudienceUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.AudienceUpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("audience_updated_at", "body", "date-time", m.AudienceUpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OrgLevelClaimsInfo) validateOrgID(formats strfmt.Registry) error {
	if swag.IsZero(m.OrgID) { // not required
		return nil
	}

	if err := validate.FormatOf("org_id", "body", "uuid", m.OrgID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OrgLevelClaimsInfo) validateTTLUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.TTLUpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("ttl_updated_at", "body", "date-time", m.TTLUpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this org level claims info based on context it is used
func (m *OrgLevelClaimsInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrgLevelClaimsInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrgLevelClaimsInfo) UnmarshalBinary(b []byte) error {
	var res OrgLevelClaimsInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
