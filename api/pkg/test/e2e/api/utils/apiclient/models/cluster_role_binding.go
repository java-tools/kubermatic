// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterRoleBinding ClusterRoleBinding references a cluster role, but does not contain it.
// swagger:model ClusterRoleBinding
type ClusterRoleBinding struct {

	// role ref name
	RoleRefName string `json:"roleRefName,omitempty"`

	// Subjects holds references to the objects the role applies to.
	Subjects []*Subject `json:"subjects"`
}

// Validate validates this cluster role binding
func (m *ClusterRoleBinding) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubjects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterRoleBinding) validateSubjects(formats strfmt.Registry) error {

	if swag.IsZero(m.Subjects) { // not required
		return nil
	}

	for i := 0; i < len(m.Subjects); i++ {
		if swag.IsZero(m.Subjects[i]) { // not required
			continue
		}

		if m.Subjects[i] != nil {
			if err := m.Subjects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subjects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterRoleBinding) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterRoleBinding) UnmarshalBinary(b []byte) error {
	var res ClusterRoleBinding
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
