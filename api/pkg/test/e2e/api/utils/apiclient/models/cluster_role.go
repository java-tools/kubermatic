// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterRole ClusterRole defines cluster RBAC role for the user cluster
// swagger:model ClusterRole
type ClusterRole struct {

	// CreationTimestamp is a timestamp representing the server time when this object was created.
	// Format: date-time
	CreationTimestamp strfmt.DateTime `json:"creationTimestamp,omitempty"`

	// DeletionTimestamp is a timestamp representing the server time when this object was deleted.
	// Format: date-time
	DeletionTimestamp strfmt.DateTime `json:"deletionTimestamp,omitempty"`

	// ID unique value that identifies the resource generated by the server. Read-Only.
	ID string `json:"id,omitempty"`

	// Name represents human readable name for the resource
	Name string `json:"name,omitempty"`

	// Rules holds all the PolicyRules for this ClusterRole
	Rules []*PolicyRule `json:"rules"`
}

// Validate validates this cluster role
func (m *ClusterRole) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletionTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterRole) validateCreationTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("creationTimestamp", "body", "date-time", m.CreationTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ClusterRole) validateDeletionTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.DeletionTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("deletionTimestamp", "body", "date-time", m.DeletionTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ClusterRole) validateRules(formats strfmt.Registry) error {

	if swag.IsZero(m.Rules) { // not required
		return nil
	}

	for i := 0; i < len(m.Rules); i++ {
		if swag.IsZero(m.Rules[i]) { // not required
			continue
		}

		if m.Rules[i] != nil {
			if err := m.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterRole) UnmarshalBinary(b []byte) error {
	var res ClusterRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
