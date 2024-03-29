// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SolPolicy Serial Over LAN
//
// Policy for configuring Serial Over LAN settings on endpoint.
//
// swagger:model solPolicy
type SolPolicy struct {
	PolicyAbstractPolicy

	// Baud Rate used for Serial Over LAN communication.
	//
	// Enum: [9600 19200 38400 57600 115200]
	BaudRate *int64 `json:"BaudRate,omitempty"`

	// Serial port through which the system routes Serial Over LAN communication. This field is available only on some Cisco UCS C-Series servers. If it is unavailable, the server uses COM port 0 by default.
	//
	// Enum: [com0 com1]
	ComPort *string `json:"ComPort,omitempty"`

	// State of Serial Over LAN service on the endpoint.
	//
	Enabled *bool `json:"Enabled,omitempty"`

	// Relationship to the Organization that owns the Managed Object.
	//
	Organization *IamAccountRef `json:"Organization,omitempty"`

	// Relationship to the profile object.
	//
	Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`

	// SSH port used to access Serial Over LAN directly. Enables bypassing Cisco IMC shell to provide direct access to Serial Over LAN.
	//
	SSHPort int64 `json:"SshPort,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SolPolicy) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyAbstractPolicy
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyAbstractPolicy = aO0

	// AO1
	var dataAO1 struct {
		BaudRate *int64 `json:"BaudRate,omitempty"`

		ComPort *string `json:"ComPort,omitempty"`

		Enabled *bool `json:"Enabled,omitempty"`

		Organization *IamAccountRef `json:"Organization,omitempty"`

		Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`

		SSHPort int64 `json:"SshPort,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.BaudRate = dataAO1.BaudRate

	m.ComPort = dataAO1.ComPort

	m.Enabled = dataAO1.Enabled

	m.Organization = dataAO1.Organization

	m.Profiles = dataAO1.Profiles

	m.SSHPort = dataAO1.SSHPort

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SolPolicy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyAbstractPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		BaudRate *int64 `json:"BaudRate,omitempty"`

		ComPort *string `json:"ComPort,omitempty"`

		Enabled *bool `json:"Enabled,omitempty"`

		Organization *IamAccountRef `json:"Organization,omitempty"`

		Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`

		SSHPort int64 `json:"SshPort,omitempty"`
	}

	dataAO1.BaudRate = m.BaudRate

	dataAO1.ComPort = m.ComPort

	dataAO1.Enabled = m.Enabled

	dataAO1.Organization = m.Organization

	dataAO1.Profiles = m.Profiles

	dataAO1.SSHPort = m.SSHPort

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this sol policy
func (m *SolPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyAbstractPolicy
	if err := m.PolicyAbstractPolicy.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaudRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var solPolicyTypeBaudRatePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[9600,19200,38400,57600,115200]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		solPolicyTypeBaudRatePropEnum = append(solPolicyTypeBaudRatePropEnum, v)
	}
}

// property enum
func (m *SolPolicy) validateBaudRateEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, solPolicyTypeBaudRatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SolPolicy) validateBaudRate(formats strfmt.Registry) error {

	if swag.IsZero(m.BaudRate) { // not required
		return nil
	}

	// value enum
	if err := m.validateBaudRateEnum("BaudRate", "body", *m.BaudRate); err != nil {
		return err
	}

	return nil
}

var solPolicyTypeComPortPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["com0","com1"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		solPolicyTypeComPortPropEnum = append(solPolicyTypeComPortPropEnum, v)
	}
}

// property enum
func (m *SolPolicy) validateComPortEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, solPolicyTypeComPortPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SolPolicy) validateComPort(formats strfmt.Registry) error {

	if swag.IsZero(m.ComPort) { // not required
		return nil
	}

	// value enum
	if err := m.validateComPortEnum("ComPort", "body", *m.ComPort); err != nil {
		return err
	}

	return nil
}

func (m *SolPolicy) validateOrganization(formats strfmt.Registry) error {

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

func (m *SolPolicy) validateProfiles(formats strfmt.Registry) error {

	if swag.IsZero(m.Profiles) { // not required
		return nil
	}

	for i := 0; i < len(m.Profiles); i++ {
		if swag.IsZero(m.Profiles[i]) { // not required
			continue
		}

		if m.Profiles[i] != nil {
			if err := m.Profiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Profiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SolPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SolPolicy) UnmarshalBinary(b []byte) error {
	var res SolPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
