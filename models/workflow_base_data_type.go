// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// WorkflowBaseDataType Workflow:Base Data Type
//
// The base data type that is used to derive all the other data types for inputs and outputs.
//
// swagger:model workflowBaseDataType
type WorkflowBaseDataType struct {
	WorkflowBaseDataTypeAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowBaseDataType) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WorkflowBaseDataTypeAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WorkflowBaseDataTypeAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowBaseDataType) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.WorkflowBaseDataTypeAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow base data type
func (m *WorkflowBaseDataType) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WorkflowBaseDataTypeAO0P0
	if err := m.WorkflowBaseDataTypeAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowBaseDataType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowBaseDataType) UnmarshalBinary(b []byte) error {
	var res WorkflowBaseDataType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WorkflowBaseDataTypeAO0P0 workflow base data type a o0 p0
// swagger:model WorkflowBaseDataTypeAO0P0
type WorkflowBaseDataTypeAO0P0 struct {

	// Provide a detailed description of the data type.
	//
	Description string `json:"Description,omitempty"`

	// Descriptive name for the data type.
	//
	Label string `json:"Label,omitempty"`

	// Pick a descriptive name for the data type.
	//
	Name string `json:"Name,omitempty"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// Specifies whether this parameter is required. The field is applicable for task and workflow.
	//
	//
	Required *bool `json:"Required,omitempty"`

	// workflow base data type a o0 p0
	WorkflowBaseDataTypeAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *WorkflowBaseDataTypeAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Provide a detailed description of the data type.
		//
		Description string `json:"Description,omitempty"`

		// Descriptive name for the data type.
		//
		Label string `json:"Label,omitempty"`

		// Pick a descriptive name for the data type.
		//
		Name string `json:"Name,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// Specifies whether this parameter is required. The field is applicable for task and workflow.
		//
		//
		Required *bool `json:"Required,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv WorkflowBaseDataTypeAO0P0

	rcv.Description = stage1.Description

	rcv.Label = stage1.Label

	rcv.Name = stage1.Name

	rcv.ObjectType = stage1.ObjectType

	rcv.Required = stage1.Required

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Description")

	delete(stage2, "Label")

	delete(stage2, "Name")

	delete(stage2, "ObjectType")

	delete(stage2, "Required")

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
		m.WorkflowBaseDataTypeAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m WorkflowBaseDataTypeAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Provide a detailed description of the data type.
		//
		Description string `json:"Description,omitempty"`

		// Descriptive name for the data type.
		//
		Label string `json:"Label,omitempty"`

		// Pick a descriptive name for the data type.
		//
		Name string `json:"Name,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// Specifies whether this parameter is required. The field is applicable for task and workflow.
		//
		//
		Required *bool `json:"Required,omitempty"`
	}

	stage1.Description = m.Description

	stage1.Label = m.Label

	stage1.Name = m.Name

	stage1.ObjectType = m.ObjectType

	stage1.Required = m.Required

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.WorkflowBaseDataTypeAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.WorkflowBaseDataTypeAO0P0)
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

// Validate validates this workflow base data type a o0 p0
func (m *WorkflowBaseDataTypeAO0P0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowBaseDataTypeAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowBaseDataTypeAO0P0) UnmarshalBinary(b []byte) error {
	var res WorkflowBaseDataTypeAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
