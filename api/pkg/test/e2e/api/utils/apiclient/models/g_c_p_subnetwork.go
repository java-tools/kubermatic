// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GCPSubnetwork GCPSubnetwork represents a object of GCP subnetworks.
// swagger:model GCPSubnetwork
type GCPSubnetwork struct {

	// gateway address
	GatewayAddress string `json:"gatewayAddress,omitempty"`

	// ID
	ID uint64 `json:"id,omitempty"`

	// IP cidr range
	IPCidrRange string `json:"ipCidrRange,omitempty"`

	// kind
	Kind string `json:"kind,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// network
	Network string `json:"network,omitempty"`

	// private IP google access
	PrivateIPGoogleAccess bool `json:"privateIpGoogleAccess,omitempty"`

	// region
	Region string `json:"region,omitempty"`

	// self link
	SelfLink string `json:"selfLink,omitempty"`
}

// Validate validates this g c p subnetwork
func (m *GCPSubnetwork) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GCPSubnetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GCPSubnetwork) UnmarshalBinary(b []byte) error {
	var res GCPSubnetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
