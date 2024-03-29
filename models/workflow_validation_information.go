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

// WorkflowValidationInformation Workflow:Validation Information
//
// Type used to capture all the validation information for the workflow definition.
//
// swagger:model workflowValidationInformation
type WorkflowValidationInformation struct {
	WorkflowValidationInformationAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowValidationInformation) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WorkflowValidationInformationAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WorkflowValidationInformationAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowValidationInformation) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.WorkflowValidationInformationAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow validation information
func (m *WorkflowValidationInformation) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WorkflowValidationInformationAO0P0
	if err := m.WorkflowValidationInformationAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowValidationInformation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowValidationInformation) UnmarshalBinary(b []byte) error {
	var res WorkflowValidationInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WorkflowValidationInformationAO0P0 workflow validation information a o0 p0
// swagger:model WorkflowValidationInformationAO0P0
type WorkflowValidationInformationAO0P0 struct {

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// The current validation state of this workflow. The possible states are Valid, Invalid, NotValidated (default).
	//
	// Read Only: true
	// Enum: [NotValidated Valid Invalid]
	State string `json:"State,omitempty"`

	// List of all workflow or task validation errors. The validation errors can be for worker task or for control tasks.
	//
	// Read Only: true
	ValidationError []*WorkflowValidationError `json:"ValidationError"`

	// workflow validation information a o0 p0
	WorkflowValidationInformationAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *WorkflowValidationInformationAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// The current validation state of this workflow. The possible states are Valid, Invalid, NotValidated (default).
		//
		// Read Only: true
		// Enum: [NotValidated Valid Invalid]
		State string `json:"State,omitempty"`

		// List of all workflow or task validation errors. The validation errors can be for worker task or for control tasks.
		//
		// Read Only: true
		ValidationError []*WorkflowValidationError `json:"ValidationError"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv WorkflowValidationInformationAO0P0

	rcv.ObjectType = stage1.ObjectType

	rcv.State = stage1.State

	rcv.ValidationError = stage1.ValidationError

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "ObjectType")

	delete(stage2, "State")

	delete(stage2, "ValidationError")

	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.WorkflowValidationInformationAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m WorkflowValidationInformationAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// The current validation state of this workflow. The possible states are Valid, Invalid, NotValidated (default).
		//
		// Read Only: true
		// Enum: [NotValidated Valid Invalid]
		State string `json:"State,omitempty"`

		// List of all workflow or task validation errors. The validation errors can be for worker task or for control tasks.
		//
		// Read Only: true
		ValidationError []*WorkflowValidationError `json:"ValidationError"`
	}

	stage1.ObjectType = m.ObjectType

	stage1.State = m.State

	stage1.ValidationError = m.ValidationError

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.WorkflowValidationInformationAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.WorkflowValidationInformationAO0P0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this workflow validation information a o0 p0
func (m *WorkflowValidationInformationAO0P0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var workflowValidationInformationAO0P0TypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NotValidated","Valid","Invalid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workflowValidationInformationAO0P0TypeStatePropEnum = append(workflowValidationInformationAO0P0TypeStatePropEnum, v)
	}
}

const (

	// WorkflowValidationInformationAO0P0StateNotValidated captures enum value "NotValidated"
	WorkflowValidationInformationAO0P0StateNotValidated string = "NotValidated"

	// WorkflowValidationInformationAO0P0StateValid captures enum value "Valid"
	WorkflowValidationInformationAO0P0StateValid string = "Valid"

	// WorkflowValidationInformationAO0P0StateInvalid captures enum value "Invalid"
	WorkflowValidationInformationAO0P0StateInvalid string = "Invalid"
)

// prop value enum
func (m *WorkflowValidationInformationAO0P0) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, workflowValidationInformationAO0P0TypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WorkflowValidationInformationAO0P0) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("State", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowValidationInformationAO0P0) validateValidationError(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidationError) { // not required
		return nil
	}

	for i := 0; i < len(m.ValidationError); i++ {
		if swag.IsZero(m.ValidationError[i]) { // not required
			continue
		}

		if m.ValidationError[i] != nil {
			if err := m.ValidationError[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ValidationError" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowValidationInformationAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowValidationInformationAO0P0) UnmarshalBinary(b []byte) error {
	var res WorkflowValidationInformationAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
