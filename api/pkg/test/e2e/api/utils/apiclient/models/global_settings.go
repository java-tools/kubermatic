// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GlobalSettings GlobalSettings defines global settings
// swagger:model GlobalSettings
type GlobalSettings struct {
	SettingSpec
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GlobalSettings) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SettingSpec
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SettingSpec = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GlobalSettings) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.SettingSpec)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this global settings
func (m *GlobalSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GlobalSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GlobalSettings) UnmarshalBinary(b []byte) error {
	var res GlobalSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
