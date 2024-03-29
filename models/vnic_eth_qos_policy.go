// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// VnicEthQosPolicy Ethernet QoS
//
// An Ethernet Quality of Service (QoS) policy assigns a system class to the outgoing traffic for a vNIC. This system class determines the quality of service for the outgoing traffic. For certain adapters you can also specify additional controls like burst and rate on the outgoing traffic.
//
// swagger:model vnicEthQosPolicy
type VnicEthQosPolicy struct {
	PolicyAbstractPolicy

	// Class of Service to be associated to the traffic on the virtual interface.
	//
	Cos int64 `json:"Cos,omitempty"`

	// The Maximum Transmission Unit (MTU) or packet size that the virtual interface accepts.
	//
	Mtu int64 `json:"Mtu,omitempty"`

	// Relationship to the Organization that owns the Managed Object.
	//
	Organization *IamAccountRef `json:"Organization,omitempty"`

	// The value in Mbps (0-40000) to use for limiting the data rate on the virtual interface. Setting this to zero will turn rate limiting off.
	//
	RateLimit int64 `json:"RateLimit,omitempty"`

	// Enables usage of the Class of Service provided by the operating system.
	//
	TrustHostCos *bool `json:"TrustHostCos,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VnicEthQosPolicy) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyAbstractPolicy
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyAbstractPolicy = aO0

	// AO1
	var dataAO1 struct {
		Cos int64 `json:"Cos,omitempty"`

		Mtu int64 `json:"Mtu,omitempty"`

		Organization *IamAccountRef `json:"Organization,omitempty"`

		RateLimit int64 `json:"RateLimit,omitempty"`

		TrustHostCos *bool `json:"TrustHostCos,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Cos = dataAO1.Cos

	m.Mtu = dataAO1.Mtu

	m.Organization = dataAO1.Organization

	m.RateLimit = dataAO1.RateLimit

	m.TrustHostCos = dataAO1.TrustHostCos

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VnicEthQosPolicy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyAbstractPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Cos int64 `json:"Cos,omitempty"`

		Mtu int64 `json:"Mtu,omitempty"`

		Organization *IamAccountRef `json:"Organization,omitempty"`

		RateLimit int64 `json:"RateLimit,omitempty"`

		TrustHostCos *bool `json:"TrustHostCos,omitempty"`
	}

	dataAO1.Cos = m.Cos

	dataAO1.Mtu = m.Mtu

	dataAO1.Organization = m.Organization

	dataAO1.RateLimit = m.RateLimit

	dataAO1.TrustHostCos = m.TrustHostCos

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vnic eth qos policy
func (m *VnicEthQosPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyAbstractPolicy
	if err := m.PolicyAbstractPolicy.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VnicEthQosPolicy) validateOrganization(formats strfmt.Registry) error {

	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Organization")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VnicEthQosPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VnicEthQosPolicy) UnmarshalBinary(b []byte) error {
	var res VnicEthQosPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
