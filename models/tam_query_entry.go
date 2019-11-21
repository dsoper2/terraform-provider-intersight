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

// TamQueryEntry Tam:Query Entry
//
// Contains a set of queries, each with an integer priority. the queries are executed in the order of specified priority. The result of each query is used as an input to the query next in priority order.
//
// swagger:model tamQueryEntry
type TamQueryEntry struct {
	TamQueryEntryAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TamQueryEntry) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TamQueryEntryAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TamQueryEntryAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TamQueryEntry) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.TamQueryEntryAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tam query entry
func (m *TamQueryEntry) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TamQueryEntryAO0P0
	if err := m.TamQueryEntryAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TamQueryEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TamQueryEntry) UnmarshalBinary(b []byte) error {
	var res TamQueryEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TamQueryEntryAO0P0 tam query entry a o0 p0
// swagger:model TamQueryEntryAO0P0
type TamQueryEntryAO0P0 struct {

	// Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source.
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

	// An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection.
	//
	Priority int64 `json:"Priority,omitempty"`

	// A SparkSQL query to be used on a given data source.
	//
	Query string `json:"Query,omitempty"`

	// tam query entry a o0 p0
	TamQueryEntryAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *TamQueryEntryAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source.
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

		// An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection.
		//
		Priority int64 `json:"Priority,omitempty"`

		// A SparkSQL query to be used on a given data source.
		//
		Query string `json:"Query,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv TamQueryEntryAO0P0

	rcv.Name = stage1.Name

	rcv.ObjectType = stage1.ObjectType

	rcv.Priority = stage1.Priority

	rcv.Query = stage1.Query

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Name")

	delete(stage2, "ObjectType")

	delete(stage2, "Priority")

	delete(stage2, "Query")

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
		m.TamQueryEntryAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m TamQueryEntryAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source.
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

		// An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection.
		//
		Priority int64 `json:"Priority,omitempty"`

		// A SparkSQL query to be used on a given data source.
		//
		Query string `json:"Query,omitempty"`
	}

	stage1.Name = m.Name

	stage1.ObjectType = m.ObjectType

	stage1.Priority = m.Priority

	stage1.Query = m.Query

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.TamQueryEntryAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.TamQueryEntryAO0P0)
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

// Validate validates this tam query entry a o0 p0
func (m *TamQueryEntryAO0P0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TamQueryEntryAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TamQueryEntryAO0P0) UnmarshalBinary(b []byte) error {
	var res TamQueryEntryAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}