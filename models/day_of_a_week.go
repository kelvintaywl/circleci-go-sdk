// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// DayOfAWeek day of a week
//
// swagger:model DayOfAWeek
type DayOfAWeek string

func NewDayOfAWeek(value DayOfAWeek) *DayOfAWeek {
	return &value
}

// Pointer returns a pointer to a freshly-allocated DayOfAWeek.
func (m DayOfAWeek) Pointer() *DayOfAWeek {
	return &m
}

const (

	// DayOfAWeekMON captures enum value "MON"
	DayOfAWeekMON DayOfAWeek = "MON"

	// DayOfAWeekTUE captures enum value "TUE"
	DayOfAWeekTUE DayOfAWeek = "TUE"

	// DayOfAWeekWED captures enum value "WED"
	DayOfAWeekWED DayOfAWeek = "WED"

	// DayOfAWeekTHU captures enum value "THU"
	DayOfAWeekTHU DayOfAWeek = "THU"

	// DayOfAWeekFRI captures enum value "FRI"
	DayOfAWeekFRI DayOfAWeek = "FRI"

	// DayOfAWeekSAT captures enum value "SAT"
	DayOfAWeekSAT DayOfAWeek = "SAT"

	// DayOfAWeekSUN captures enum value "SUN"
	DayOfAWeekSUN DayOfAWeek = "SUN"
)

// for schema
var dayOfAWeekEnum []interface{}

func init() {
	var res []DayOfAWeek
	if err := json.Unmarshal([]byte(`["MON","TUE","WED","THU","FRI","SAT","SUN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dayOfAWeekEnum = append(dayOfAWeekEnum, v)
	}
}

func (m DayOfAWeek) validateDayOfAWeekEnum(path, location string, value DayOfAWeek) error {
	if err := validate.EnumCase(path, location, value, dayOfAWeekEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this day of a week
func (m DayOfAWeek) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDayOfAWeekEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this day of a week based on context it is used
func (m DayOfAWeek) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}