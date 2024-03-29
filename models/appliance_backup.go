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

// ApplianceBackup Appliance:Backup
//
// Backup tracks all backup requests to create a full system backup of the Intersight
// Appliance. There will be only one Backup managed object with a 'Started' state at
// any time. All other Backup managed objects will be in terminal states.
//
// swagger:model applianceBackup
type ApplianceBackup struct {
	ApplianceBackupBase

	// Backup managed object to Account relationship.
	//
	Account *IamAccountRef `json:"Account,omitempty"`

	// Elapsed time in seconds since the backup process has started.
	//
	// Read Only: true
	ElapsedTime int64 `json:"ElapsedTime,omitempty"`

	// End date and time of the backup process.
	//
	// Read Only: true
	// Format: date-time
	EndTime strfmt.DateTime `json:"EndTime,omitempty"`

	// is password set
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`

	// Messages generated during the backup process.
	//
	// Read Only: true
	Messages []string `json:"Messages"`

	// Password to authenticate the fileserver.
	//
	Password string `json:"Password,omitempty"`

	// Start date and time of the backup process.
	//
	// Read Only: true
	// Format: date-time
	StartTime strfmt.DateTime `json:"StartTime,omitempty"`

	// Status of the backup managed object.
	//
	// Read Only: true
	// Enum: [Started Created Failed Completed Copied]
	Status string `json:"Status,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ApplianceBackup) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ApplianceBackupBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ApplianceBackupBase = aO0

	// AO1
	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		ElapsedTime int64 `json:"ElapsedTime,omitempty"`

		EndTime strfmt.DateTime `json:"EndTime,omitempty"`

		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`

		Messages []string `json:"Messages"`

		Password string `json:"Password,omitempty"`

		StartTime strfmt.DateTime `json:"StartTime,omitempty"`

		Status string `json:"Status,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Account = dataAO1.Account

	m.ElapsedTime = dataAO1.ElapsedTime

	m.EndTime = dataAO1.EndTime

	m.IsPasswordSet = dataAO1.IsPasswordSet

	m.Messages = dataAO1.Messages

	m.Password = dataAO1.Password

	m.StartTime = dataAO1.StartTime

	m.Status = dataAO1.Status

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ApplianceBackup) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ApplianceBackupBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		ElapsedTime int64 `json:"ElapsedTime,omitempty"`

		EndTime strfmt.DateTime `json:"EndTime,omitempty"`

		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`

		Messages []string `json:"Messages"`

		Password string `json:"Password,omitempty"`

		StartTime strfmt.DateTime `json:"StartTime,omitempty"`

		Status string `json:"Status,omitempty"`
	}

	dataAO1.Account = m.Account

	dataAO1.ElapsedTime = m.ElapsedTime

	dataAO1.EndTime = m.EndTime

	dataAO1.IsPasswordSet = m.IsPasswordSet

	dataAO1.Messages = m.Messages

	dataAO1.Password = m.Password

	dataAO1.StartTime = m.StartTime

	dataAO1.Status = m.Status

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this appliance backup
func (m *ApplianceBackup) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ApplianceBackupBase
	if err := m.ApplianceBackupBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
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

func (m *ApplianceBackup) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Account")
			}
			return err
		}
	}

	return nil
}

func (m *ApplianceBackup) validateEndTime(formats strfmt.Registry) error {

	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("EndTime", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ApplianceBackup) validateStartTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("StartTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var applianceBackupTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Started","Created","Failed","Completed","Copied"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applianceBackupTypeStatusPropEnum = append(applianceBackupTypeStatusPropEnum, v)
	}
}

// property enum
func (m *ApplianceBackup) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, applianceBackupTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ApplianceBackup) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("Status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplianceBackup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplianceBackup) UnmarshalBinary(b []byte) error {
	var res ApplianceBackup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
