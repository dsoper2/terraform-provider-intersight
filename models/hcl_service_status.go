// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HclServiceStatus Hcl:Service Status
//
// Status of the service indicatating if the service is up or under maintenance due to data update. Service will not be able serve any requests when the data is being updated. Collection will have only one document.
//
// swagger:model hclServiceStatus
type HclServiceStatus struct {
	MoBaseMo

	// Version of the last modified exemption file.
	//
	ExemptionFileVersion string `json:"ExemptionFileVersion,omitempty"`

	// A field to uniquely identify the document with the satus.
	//
	Identity string `json:"Identity,omitempty"`

	// The timestamp of the last modified record in the HCL tool database. Used to query and get updated records.
	//
	// Format: date-time
	LastHclDataModifiedTime strfmt.DateTime `json:"LastHclDataModifiedTime,omitempty"`

	// Checksum of the data dump used as the base for delta updates.
	//
	LastImportedDataChecksum string `json:"LastImportedDataChecksum,omitempty"`

	// Status of the service indicatating if the service is up or under maintenance due to data update.
	//
	// Enum: [Unknown Initializing DataRefreshing Active]
	Status *string `json:"Status,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HclServiceStatus) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		ExemptionFileVersion string `json:"ExemptionFileVersion,omitempty"`

		Identity string `json:"Identity,omitempty"`

		LastHclDataModifiedTime strfmt.DateTime `json:"LastHclDataModifiedTime,omitempty"`

		LastImportedDataChecksum string `json:"LastImportedDataChecksum,omitempty"`

		Status *string `json:"Status,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ExemptionFileVersion = dataAO1.ExemptionFileVersion

	m.Identity = dataAO1.Identity

	m.LastHclDataModifiedTime = dataAO1.LastHclDataModifiedTime

	m.LastImportedDataChecksum = dataAO1.LastImportedDataChecksum

	m.Status = dataAO1.Status

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HclServiceStatus) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ExemptionFileVersion string `json:"ExemptionFileVersion,omitempty"`

		Identity string `json:"Identity,omitempty"`

		LastHclDataModifiedTime strfmt.DateTime `json:"LastHclDataModifiedTime,omitempty"`

		LastImportedDataChecksum string `json:"LastImportedDataChecksum,omitempty"`

		Status *string `json:"Status,omitempty"`
	}

	dataAO1.ExemptionFileVersion = m.ExemptionFileVersion

	dataAO1.Identity = m.Identity

	dataAO1.LastHclDataModifiedTime = m.LastHclDataModifiedTime

	dataAO1.LastImportedDataChecksum = m.LastImportedDataChecksum

	dataAO1.Status = m.Status

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hcl service status
func (m *HclServiceStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastHclDataModifiedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HclServiceStatus) validateLastHclDataModifiedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastHclDataModifiedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("LastHclDataModifiedTime", "body", "date-time", m.LastHclDataModifiedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var hclServiceStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","Initializing","DataRefreshing","Active"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hclServiceStatusTypeStatusPropEnum = append(hclServiceStatusTypeStatusPropEnum, v)
	}
}

// property enum
func (m *HclServiceStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hclServiceStatusTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HclServiceStatus) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("Status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HclServiceStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HclServiceStatus) UnmarshalBinary(b []byte) error {
	var res HclServiceStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
