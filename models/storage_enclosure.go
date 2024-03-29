// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StorageEnclosure Storage:Enclosure
//
// Storage Enclosure for physical disks.
//
// swagger:model storageEnclosure
type StorageEnclosure struct {
	EquipmentBase

	// chassis Id
	// Read Only: true
	ChassisID int64 `json:"ChassisId,omitempty"`

	// A collection of references to the [compute.Blade](mo://compute.Blade) Managed Object.
	//
	// When this managed object is deleted, the referenced [compute.Blade](mo://compute.Blade) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	ComputeBlade *ComputeBladeRef `json:"ComputeBlade,omitempty"`

	// A collection of references to the [compute.RackUnit](mo://compute.RackUnit) Managed Object.
	//
	// When this managed object is deleted, the referenced [compute.RackUnit](mo://compute.RackUnit) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	ComputeRackUnit *ComputeRackUnitRef `json:"ComputeRackUnit,omitempty"`

	// description
	// Read Only: true
	Description string `json:"Description,omitempty"`

	// enclosure disk slots
	// Read Only: true
	EnclosureDiskSlots []*StorageEnclosureDiskSlotEpRef `json:"EnclosureDiskSlots"`

	// enclosure disks
	// Read Only: true
	EnclosureDisks []*StorageEnclosureDiskRef `json:"EnclosureDisks"`

	// enclosure Id
	// Read Only: true
	EnclosureID int64 `json:"EnclosureId,omitempty"`

	// A collection of references to the [equipment.Chassis](mo://equipment.Chassis) Managed Object.
	//
	// When this managed object is deleted, the referenced [equipment.Chassis](mo://equipment.Chassis) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	EquipmentChassis *EquipmentChassisRef `json:"EquipmentChassis,omitempty"`

	// num slots
	// Read Only: true
	NumSlots int64 `json:"NumSlots,omitempty"`

	// physical disks
	// Read Only: true
	PhysicalDisks []*StoragePhysicalDiskRef `json:"PhysicalDisks"`

	// presence
	// Read Only: true
	Presence string `json:"Presence,omitempty"`

	// The Device to which this Managed Object is associated.
	//
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

	// server Id
	// Read Only: true
	ServerID int64 `json:"ServerId,omitempty"`

	// type
	// Read Only: true
	Type string `json:"Type,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StorageEnclosure) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 EquipmentBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.EquipmentBase = aO0

	// AO1
	var dataAO1 struct {
		ChassisID int64 `json:"ChassisId,omitempty"`

		ComputeBlade *ComputeBladeRef `json:"ComputeBlade,omitempty"`

		ComputeRackUnit *ComputeRackUnitRef `json:"ComputeRackUnit,omitempty"`

		Description string `json:"Description,omitempty"`

		EnclosureDiskSlots []*StorageEnclosureDiskSlotEpRef `json:"EnclosureDiskSlots"`

		EnclosureDisks []*StorageEnclosureDiskRef `json:"EnclosureDisks"`

		EnclosureID int64 `json:"EnclosureId,omitempty"`

		EquipmentChassis *EquipmentChassisRef `json:"EquipmentChassis,omitempty"`

		NumSlots int64 `json:"NumSlots,omitempty"`

		PhysicalDisks []*StoragePhysicalDiskRef `json:"PhysicalDisks"`

		Presence string `json:"Presence,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		ServerID int64 `json:"ServerId,omitempty"`

		Type string `json:"Type,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ChassisID = dataAO1.ChassisID

	m.ComputeBlade = dataAO1.ComputeBlade

	m.ComputeRackUnit = dataAO1.ComputeRackUnit

	m.Description = dataAO1.Description

	m.EnclosureDiskSlots = dataAO1.EnclosureDiskSlots

	m.EnclosureDisks = dataAO1.EnclosureDisks

	m.EnclosureID = dataAO1.EnclosureID

	m.EquipmentChassis = dataAO1.EquipmentChassis

	m.NumSlots = dataAO1.NumSlots

	m.PhysicalDisks = dataAO1.PhysicalDisks

	m.Presence = dataAO1.Presence

	m.RegisteredDevice = dataAO1.RegisteredDevice

	m.ServerID = dataAO1.ServerID

	m.Type = dataAO1.Type

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StorageEnclosure) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.EquipmentBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ChassisID int64 `json:"ChassisId,omitempty"`

		ComputeBlade *ComputeBladeRef `json:"ComputeBlade,omitempty"`

		ComputeRackUnit *ComputeRackUnitRef `json:"ComputeRackUnit,omitempty"`

		Description string `json:"Description,omitempty"`

		EnclosureDiskSlots []*StorageEnclosureDiskSlotEpRef `json:"EnclosureDiskSlots"`

		EnclosureDisks []*StorageEnclosureDiskRef `json:"EnclosureDisks"`

		EnclosureID int64 `json:"EnclosureId,omitempty"`

		EquipmentChassis *EquipmentChassisRef `json:"EquipmentChassis,omitempty"`

		NumSlots int64 `json:"NumSlots,omitempty"`

		PhysicalDisks []*StoragePhysicalDiskRef `json:"PhysicalDisks"`

		Presence string `json:"Presence,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		ServerID int64 `json:"ServerId,omitempty"`

		Type string `json:"Type,omitempty"`
	}

	dataAO1.ChassisID = m.ChassisID

	dataAO1.ComputeBlade = m.ComputeBlade

	dataAO1.ComputeRackUnit = m.ComputeRackUnit

	dataAO1.Description = m.Description

	dataAO1.EnclosureDiskSlots = m.EnclosureDiskSlots

	dataAO1.EnclosureDisks = m.EnclosureDisks

	dataAO1.EnclosureID = m.EnclosureID

	dataAO1.EquipmentChassis = m.EquipmentChassis

	dataAO1.NumSlots = m.NumSlots

	dataAO1.PhysicalDisks = m.PhysicalDisks

	dataAO1.Presence = m.Presence

	dataAO1.RegisteredDevice = m.RegisteredDevice

	dataAO1.ServerID = m.ServerID

	dataAO1.Type = m.Type

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage enclosure
func (m *StorageEnclosure) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with EquipmentBase
	if err := m.EquipmentBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComputeBlade(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComputeRackUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnclosureDiskSlots(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnclosureDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEquipmentChassis(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhysicalDisks(formats); err != nil {
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

func (m *StorageEnclosure) validateComputeBlade(formats strfmt.Registry) error {

	if swag.IsZero(m.ComputeBlade) { // not required
		return nil
	}

	if m.ComputeBlade != nil {
		if err := m.ComputeBlade.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ComputeBlade")
			}
			return err
		}
	}

	return nil
}

func (m *StorageEnclosure) validateComputeRackUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.ComputeRackUnit) { // not required
		return nil
	}

	if m.ComputeRackUnit != nil {
		if err := m.ComputeRackUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ComputeRackUnit")
			}
			return err
		}
	}

	return nil
}

func (m *StorageEnclosure) validateEnclosureDiskSlots(formats strfmt.Registry) error {

	if swag.IsZero(m.EnclosureDiskSlots) { // not required
		return nil
	}

	for i := 0; i < len(m.EnclosureDiskSlots); i++ {
		if swag.IsZero(m.EnclosureDiskSlots[i]) { // not required
			continue
		}

		if m.EnclosureDiskSlots[i] != nil {
			if err := m.EnclosureDiskSlots[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("EnclosureDiskSlots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StorageEnclosure) validateEnclosureDisks(formats strfmt.Registry) error {

	if swag.IsZero(m.EnclosureDisks) { // not required
		return nil
	}

	for i := 0; i < len(m.EnclosureDisks); i++ {
		if swag.IsZero(m.EnclosureDisks[i]) { // not required
			continue
		}

		if m.EnclosureDisks[i] != nil {
			if err := m.EnclosureDisks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("EnclosureDisks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StorageEnclosure) validateEquipmentChassis(formats strfmt.Registry) error {

	if swag.IsZero(m.EquipmentChassis) { // not required
		return nil
	}

	if m.EquipmentChassis != nil {
		if err := m.EquipmentChassis.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EquipmentChassis")
			}
			return err
		}
	}

	return nil
}

func (m *StorageEnclosure) validatePhysicalDisks(formats strfmt.Registry) error {

	if swag.IsZero(m.PhysicalDisks) { // not required
		return nil
	}

	for i := 0; i < len(m.PhysicalDisks); i++ {
		if swag.IsZero(m.PhysicalDisks[i]) { // not required
			continue
		}

		if m.PhysicalDisks[i] != nil {
			if err := m.PhysicalDisks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PhysicalDisks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StorageEnclosure) validateRegisteredDevice(formats strfmt.Registry) error {

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
func (m *StorageEnclosure) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageEnclosure) UnmarshalBinary(b []byte) error {
	var res StorageEnclosure
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
