// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GCPNetwork GCPNetwork represents a object of GCP networks.
// swagger:model GCPNetwork
type GCPNetwork struct {

	// auto create subnetworks
	AutoCreateSubnetworks bool `json:"autoCreateSubnetworks,omitempty"`

	// ID
	ID uint64 `json:"id,omitempty"`

	// kind
	Kind string `json:"kind,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// subnetworks
	Subnetworks []string `json:"subnetworks"`
}

// Validate validates this g c p network
func (m *GCPNetwork) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GCPNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GCPNetwork) UnmarshalBinary(b []byte) error {
	var res GCPNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
