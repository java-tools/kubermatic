// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProjectGroup ProjectGroup is a helper data structure that
// stores the information about a project and a group prefix that a user belongs to
// swagger:model ProjectGroup
type ProjectGroup struct {

	// group prefix
	GroupPrefix string `json:"group,omitempty"`

	// ID
	ID string `json:"id,omitempty"`
}

// Validate validates this project group
func (m *ProjectGroup) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectGroup) UnmarshalBinary(b []byte) error {
	var res ProjectGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
