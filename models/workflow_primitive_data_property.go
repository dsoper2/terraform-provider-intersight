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

// WorkflowPrimitiveDataProperty Workflow:Primitive Data Property
//
// Capture all the properties for primitive data type.
//
// swagger:model workflowPrimitiveDataProperty
type WorkflowPrimitiveDataProperty struct {
	WorkflowPrimitiveDataPropertyAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowPrimitiveDataProperty) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WorkflowPrimitiveDataPropertyAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WorkflowPrimitiveDataPropertyAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowPrimitiveDataProperty) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.WorkflowPrimitiveDataPropertyAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow primitive data property
func (m *WorkflowPrimitiveDataProperty) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WorkflowPrimitiveDataPropertyAO0P0
	if err := m.WorkflowPrimitiveDataPropertyAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowPrimitiveDataProperty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowPrimitiveDataProperty) UnmarshalBinary(b []byte) error {
	var res WorkflowPrimitiveDataProperty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WorkflowPrimitiveDataPropertyAO0P0 workflow primitive data property a o0 p0
// swagger:model WorkflowPrimitiveDataPropertyAO0P0
type WorkflowPrimitiveDataPropertyAO0P0 struct {

	// Constraints that must be applied to the parameter value supplied for this data type.
	//
	Constraints *WorkflowConstraints `json:"Constraints,omitempty"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// Intersight allows the secure properties to be used as task input/output. The values of
	// these properties are encrypted and stored in Intersight.
	//
	// This flag marks the property to be secure when it is set to true.
	//
	//
	Secure *bool `json:"Secure,omitempty"`

	// Specify the enum type for primitive data type.
	//
	// Enum: [string integer float boolean json enum]
	Type *string `json:"Type,omitempty"`

	// workflow primitive data property a o0 p0
	WorkflowPrimitiveDataPropertyAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *WorkflowPrimitiveDataPropertyAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Constraints that must be applied to the parameter value supplied for this data type.
		//
		Constraints *WorkflowConstraints `json:"Constraints,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// Intersight allows the secure properties to be used as task input/output. The values of
		// these properties are encrypted and stored in Intersight.
		//
		// This flag marks the property to be secure when it is set to true.
		//
		//
		Secure *bool `json:"Secure,omitempty"`

		// Specify the enum type for primitive data type.
		//
		// Enum: [string integer float boolean json enum]
		Type *string `json:"Type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv WorkflowPrimitiveDataPropertyAO0P0

	rcv.Constraints = stage1.Constraints

	rcv.ObjectType = stage1.ObjectType

	rcv.Secure = stage1.Secure

	rcv.Type = stage1.Type

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Constraints")

	delete(stage2, "ObjectType")

	delete(stage2, "Secure")

	delete(stage2, "Type")

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
		m.WorkflowPrimitiveDataPropertyAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m WorkflowPrimitiveDataPropertyAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Constraints that must be applied to the parameter value supplied for this data type.
		//
		Constraints *WorkflowConstraints `json:"Constraints,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// Intersight allows the secure properties to be used as task input/output. The values of
		// these properties are encrypted and stored in Intersight.
		//
		// This flag marks the property to be secure when it is set to true.
		//
		//
		Secure *bool `json:"Secure,omitempty"`

		// Specify the enum type for primitive data type.
		//
		// Enum: [string integer float boolean json enum]
		Type *string `json:"Type,omitempty"`
	}

	stage1.Constraints = m.Constraints

	stage1.ObjectType = m.ObjectType

	stage1.Secure = m.Secure

	stage1.Type = m.Type

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.WorkflowPrimitiveDataPropertyAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.WorkflowPrimitiveDataPropertyAO0P0)
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

// Validate validates this workflow primitive data property a o0 p0
func (m *WorkflowPrimitiveDataPropertyAO0P0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowPrimitiveDataPropertyAO0P0) validateConstraints(formats strfmt.Registry) error {

	if swag.IsZero(m.Constraints) { // not required
		return nil
	}

	if m.Constraints != nil {
		if err := m.Constraints.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Constraints")
			}
			return err
		}
	}

	return nil
}

var workflowPrimitiveDataPropertyAO0P0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["string","integer","float","boolean","json","enum"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workflowPrimitiveDataPropertyAO0P0TypeTypePropEnum = append(workflowPrimitiveDataPropertyAO0P0TypeTypePropEnum, v)
	}
}

const (

	// WorkflowPrimitiveDataPropertyAO0P0TypeString captures enum value "string"
	WorkflowPrimitiveDataPropertyAO0P0TypeString string = "string"

	// WorkflowPrimitiveDataPropertyAO0P0TypeInteger captures enum value "integer"
	WorkflowPrimitiveDataPropertyAO0P0TypeInteger string = "integer"

	// WorkflowPrimitiveDataPropertyAO0P0TypeFloat captures enum value "float"
	WorkflowPrimitiveDataPropertyAO0P0TypeFloat string = "float"

	// WorkflowPrimitiveDataPropertyAO0P0TypeBoolean captures enum value "boolean"
	WorkflowPrimitiveDataPropertyAO0P0TypeBoolean string = "boolean"

	// WorkflowPrimitiveDataPropertyAO0P0TypeJSON captures enum value "json"
	WorkflowPrimitiveDataPropertyAO0P0TypeJSON string = "json"

	// WorkflowPrimitiveDataPropertyAO0P0TypeEnum captures enum value "enum"
	WorkflowPrimitiveDataPropertyAO0P0TypeEnum string = "enum"
)

// prop value enum
func (m *WorkflowPrimitiveDataPropertyAO0P0) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, workflowPrimitiveDataPropertyAO0P0TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WorkflowPrimitiveDataPropertyAO0P0) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("Type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowPrimitiveDataPropertyAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowPrimitiveDataPropertyAO0P0) UnmarshalBinary(b []byte) error {
	var res WorkflowPrimitiveDataPropertyAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
