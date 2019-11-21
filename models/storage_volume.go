// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StorageVolume Storage:Volume
//
// Generic storage volume object. It is a provisioned storage identity which can be mapped to host to enable host access.
//
// swagger:model storageVolume
type StorageVolume struct {
	MoBaseMo

	// Short description about the volume.
	//
	// Read Only: true
	Description string `json:"Description,omitempty"`

	// NAA id of volume. It is a significant number to identify corresponding lun path in hypervisor.
	//
	// Read Only: true
	NaaID string `json:"NaaId,omitempty"`

	// Named entitiy of the volume.
	//
	// Read Only: true
	Name string `json:"Name,omitempty"`

	// User provisioned volume size. It is a size exposed to host.
	//
	// Read Only: true
	Size int64 `json:"Size,omitempty"`

	// Storage array managed object.
	//
	// Read Only: true
	StorageArray *StorageGenericArrayRef `json:"StorageArray,omitempty"`

	// Storage utilization of volume entity in storage array.
	//
	// Read Only: true
	StorageUtilization *StorageCapacity `json:"StorageUtilization,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StorageVolume) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Description string `json:"Description,omitempty"`

		NaaID string `json:"NaaId,omitempty"`

		Name string `json:"Name,omitempty"`

		Size int64 `json:"Size,omitempty"`

		StorageArray *StorageGenericArrayRef `json:"StorageArray,omitempty"`

		StorageUtilization *StorageCapacity `json:"StorageUtilization,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Description = dataAO1.Description

	m.NaaID = dataAO1.NaaID

	m.Name = dataAO1.Name

	m.Size = dataAO1.Size

	m.StorageArray = dataAO1.StorageArray

	m.StorageUtilization = dataAO1.StorageUtilization

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StorageVolume) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Description string `json:"Description,omitempty"`

		NaaID string `json:"NaaId,omitempty"`

		Name string `json:"Name,omitempty"`

		Size int64 `json:"Size,omitempty"`

		StorageArray *StorageGenericArrayRef `json:"StorageArray,omitempty"`

		StorageUtilization *StorageCapacity `json:"StorageUtilization,omitempty"`
	}

	dataAO1.Description = m.Description

	dataAO1.NaaID = m.NaaID

	dataAO1.Name = m.Name

	dataAO1.Size = m.Size

	dataAO1.StorageArray = m.StorageArray

	dataAO1.StorageUtilization = m.StorageUtilization

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage volume
func (m *StorageVolume) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageArray(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageUtilization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageVolume) validateStorageArray(formats strfmt.Registry) error {

	if swag.IsZero(m.StorageArray) { // not required
		return nil
	}

	if m.StorageArray != nil {
		if err := m.StorageArray.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StorageArray")
			}
			return err
		}
	}

	return nil
}

func (m *StorageVolume) validateStorageUtilization(formats strfmt.Registry) error {

	if swag.IsZero(m.StorageUtilization) { // not required
		return nil
	}

	if m.StorageUtilization != nil {
		if err := m.StorageUtilization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StorageUtilization")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageVolume) UnmarshalBinary(b []byte) error {
	var res StorageVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}