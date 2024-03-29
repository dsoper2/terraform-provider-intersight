// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SecurityUnit Security:Unit
//
// The crypto card present on a server.
//
// swagger:model securityUnit
type SecurityUnit struct {
	EquipmentBase

	// A collection of references to the [compute.Board](mo://compute.Board) Managed Object.
	//
	// When this managed object is deleted, the referenced [compute.Board](mo://compute.Board) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	ComputeBoard *ComputeBoardRef `json:"ComputeBoard,omitempty"`

	// oper state
	// Read Only: true
	OperState string `json:"OperState,omitempty"`

	// operability
	// Read Only: true
	Operability string `json:"Operability,omitempty"`

	// part number
	// Read Only: true
	PartNumber string `json:"PartNumber,omitempty"`

	// pci slot
	// Read Only: true
	PciSlot string `json:"PciSlot,omitempty"`

	// power
	// Read Only: true
	Power string `json:"Power,omitempty"`

	// presence
	// Read Only: true
	Presence string `json:"Presence,omitempty"`

	// The Device to which this Managed Object is associated.
	//
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

	// thermal
	// Read Only: true
	Thermal string `json:"Thermal,omitempty"`

	// unit Id
	// Read Only: true
	UnitID int64 `json:"UnitId,omitempty"`

	// vid
	// Read Only: true
	Vid string `json:"Vid,omitempty"`

	// voltage
	// Read Only: true
	Voltage string `json:"Voltage,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SecurityUnit) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 EquipmentBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.EquipmentBase = aO0

	// AO1
	var dataAO1 struct {
		ComputeBoard *ComputeBoardRef `json:"ComputeBoard,omitempty"`

		OperState string `json:"OperState,omitempty"`

		Operability string `json:"Operability,omitempty"`

		PartNumber string `json:"PartNumber,omitempty"`

		PciSlot string `json:"PciSlot,omitempty"`

		Power string `json:"Power,omitempty"`

		Presence string `json:"Presence,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		Thermal string `json:"Thermal,omitempty"`

		UnitID int64 `json:"UnitId,omitempty"`

		Vid string `json:"Vid,omitempty"`

		Voltage string `json:"Voltage,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ComputeBoard = dataAO1.ComputeBoard

	m.OperState = dataAO1.OperState

	m.Operability = dataAO1.Operability

	m.PartNumber = dataAO1.PartNumber

	m.PciSlot = dataAO1.PciSlot

	m.Power = dataAO1.Power

	m.Presence = dataAO1.Presence

	m.RegisteredDevice = dataAO1.RegisteredDevice

	m.Thermal = dataAO1.Thermal

	m.UnitID = dataAO1.UnitID

	m.Vid = dataAO1.Vid

	m.Voltage = dataAO1.Voltage

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SecurityUnit) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.EquipmentBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ComputeBoard *ComputeBoardRef `json:"ComputeBoard,omitempty"`

		OperState string `json:"OperState,omitempty"`

		Operability string `json:"Operability,omitempty"`

		PartNumber string `json:"PartNumber,omitempty"`

		PciSlot string `json:"PciSlot,omitempty"`

		Power string `json:"Power,omitempty"`

		Presence string `json:"Presence,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		Thermal string `json:"Thermal,omitempty"`

		UnitID int64 `json:"UnitId,omitempty"`

		Vid string `json:"Vid,omitempty"`

		Voltage string `json:"Voltage,omitempty"`
	}

	dataAO1.ComputeBoard = m.ComputeBoard

	dataAO1.OperState = m.OperState

	dataAO1.Operability = m.Operability

	dataAO1.PartNumber = m.PartNumber

	dataAO1.PciSlot = m.PciSlot

	dataAO1.Power = m.Power

	dataAO1.Presence = m.Presence

	dataAO1.RegisteredDevice = m.RegisteredDevice

	dataAO1.Thermal = m.Thermal

	dataAO1.UnitID = m.UnitID

	dataAO1.Vid = m.Vid

	dataAO1.Voltage = m.Voltage

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this security unit
func (m *SecurityUnit) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with EquipmentBase
	if err := m.EquipmentBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComputeBoard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegisteredDevice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityUnit) validateComputeBoard(formats strfmt.Registry) error {

	if swag.IsZero(m.ComputeBoard) { // not required
		return nil
	}

	if m.ComputeBoard != nil {
		if err := m.ComputeBoard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ComputeBoard")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityUnit) validateRegisteredDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.RegisteredDevice) { // not required
		return nil
	}

	if m.RegisteredDevice != nil {
		if err := m.RegisteredDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RegisteredDevice")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityUnit) UnmarshalBinary(b []byte) error {
	var res SecurityUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
