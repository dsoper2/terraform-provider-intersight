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

// WorkflowMoReferenceProperty Workflow:Mo Reference Property
//
// Capture all the properties for an Intersight managed object reference.
//
// swagger:model workflowMoReferenceProperty
type WorkflowMoReferenceProperty struct {
	WorkflowMoReferencePropertyAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowMoReferenceProperty) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WorkflowMoReferencePropertyAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WorkflowMoReferencePropertyAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowMoReferenceProperty) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.WorkflowMoReferencePropertyAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow mo reference property
func (m *WorkflowMoReferenceProperty) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WorkflowMoReferencePropertyAO0P0
	if err := m.WorkflowMoReferencePropertyAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowMoReferenceProperty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowMoReferenceProperty) UnmarshalBinary(b []byte) error {
	var res WorkflowMoReferenceProperty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WorkflowMoReferencePropertyAO0P0 workflow mo reference property a o0 p0
// swagger:model WorkflowMoReferencePropertyAO0P0
type WorkflowMoReferencePropertyAO0P0 struct {

	// List of properties from the Intersight object which can help to identify the object. Typically the set of identity constraints on the object can be listed here to help the user identity the managed object.
	//
	DisplayAttributes []string `json:"DisplayAttributes"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// Field to hold an Intersight API along with an optional filter to narrow down the search options.
	//
	Selector string `json:"Selector,omitempty"`

	// workflow mo reference property a o0 p0
	WorkflowMoReferencePropertyAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *WorkflowMoReferencePropertyAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// List of properties from the Intersight object which can help to identify the object. Typically the set of identity constraints on the object can be listed here to help the user identity the managed object.
		//
		DisplayAttributes []string `json:"DisplayAttributes"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// Field to hold an Intersight API along with an optional filter to narrow down the search options.
		//
		Selector string `json:"Selector,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv WorkflowMoReferencePropertyAO0P0

	rcv.DisplayAttributes = stage1.DisplayAttributes

	rcv.ObjectType = stage1.ObjectType

	rcv.Selector = stage1.Selector

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "DisplayAttributes")

	delete(stage2, "ObjectType")

	delete(stage2, "Selector")

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
		m.WorkflowMoReferencePropertyAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m WorkflowMoReferencePropertyAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// List of properties from the Intersight object which can help to identify the object. Typically the set of identity constraints on the object can be listed here to help the user identity the managed object.
		//
		DisplayAttributes []string `json:"DisplayAttributes"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// Field to hold an Intersight API along with an optional filter to narrow down the search options.
		//
		Selector string `json:"Selector,omitempty"`
	}

	stage1.DisplayAttributes = m.DisplayAttributes

	stage1.ObjectType = m.ObjectType

	stage1.Selector = m.Selector

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.WorkflowMoReferencePropertyAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.WorkflowMoReferencePropertyAO0P0)
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

// Validate validates this workflow mo reference property a o0 p0
func (m *WorkflowMoReferencePropertyAO0P0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowMoReferencePropertyAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowMoReferencePropertyAO0P0) UnmarshalBinary(b []byte) error {
	var res WorkflowMoReferencePropertyAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
