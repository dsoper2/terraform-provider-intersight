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

// InventoryInventoryMo Inventory:Inventory Mo
//
// Complex type representing the inventory MO.
//
// swagger:model inventoryInventoryMo
type InventoryInventoryMo struct {
	InventoryInventoryMoAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *InventoryInventoryMo) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 InventoryInventoryMoAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.InventoryInventoryMoAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m InventoryInventoryMo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.InventoryInventoryMoAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this inventory inventory mo
func (m *InventoryInventoryMo) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with InventoryInventoryMoAO0P0
	if err := m.InventoryInventoryMoAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *InventoryInventoryMo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryInventoryMo) UnmarshalBinary(b []byte) error {
	var res InventoryInventoryMo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InventoryInventoryMoAO0P0 inventory inventory mo a o0 p0
// swagger:model InventoryInventoryMoAO0P0
type InventoryInventoryMoAO0P0 struct {

	// The UCS DN of the MO for which the latest inventory to be fetched. If this property is empty and moId property has the Moid of the MO to be updated, the Moid will be used. If this property is empty and moId is also empty, all the MOs of the given moType will be updated.
	//
	MoDn string `json:"MoDn,omitempty"`

	// The MO id of an MO for which the latest inventory to be fetched. If this property is empty and moDn property has the UCS DN of the MO to be updated, the DN will be used. If this property is empty and moDn is also empty, all the MOs of the given moType will be updated.
	//
	MoID string `json:"MoId,omitempty"`

	// The type of the MO for which the latest inventory to be fetched.
	//
	MoType string `json:"MoType,omitempty"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// inventory inventory mo a o0 p0
	InventoryInventoryMoAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *InventoryInventoryMoAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The UCS DN of the MO for which the latest inventory to be fetched. If this property is empty and moId property has the Moid of the MO to be updated, the Moid will be used. If this property is empty and moId is also empty, all the MOs of the given moType will be updated.
		//
		MoDn string `json:"MoDn,omitempty"`

		// The MO id of an MO for which the latest inventory to be fetched. If this property is empty and moDn property has the UCS DN of the MO to be updated, the DN will be used. If this property is empty and moDn is also empty, all the MOs of the given moType will be updated.
		//
		MoID string `json:"MoId,omitempty"`

		// The type of the MO for which the latest inventory to be fetched.
		//
		MoType string `json:"MoType,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv InventoryInventoryMoAO0P0

	rcv.MoDn = stage1.MoDn

	rcv.MoID = stage1.MoID

	rcv.MoType = stage1.MoType

	rcv.ObjectType = stage1.ObjectType

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "MoDn")

	delete(stage2, "MoId")

	delete(stage2, "MoType")

	delete(stage2, "ObjectType")

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
		m.InventoryInventoryMoAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m InventoryInventoryMoAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The UCS DN of the MO for which the latest inventory to be fetched. If this property is empty and moId property has the Moid of the MO to be updated, the Moid will be used. If this property is empty and moId is also empty, all the MOs of the given moType will be updated.
		//
		MoDn string `json:"MoDn,omitempty"`

		// The MO id of an MO for which the latest inventory to be fetched. If this property is empty and moDn property has the UCS DN of the MO to be updated, the DN will be used. If this property is empty and moDn is also empty, all the MOs of the given moType will be updated.
		//
		MoID string `json:"MoId,omitempty"`

		// The type of the MO for which the latest inventory to be fetched.
		//
		MoType string `json:"MoType,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`
	}

	stage1.MoDn = m.MoDn

	stage1.MoID = m.MoID

	stage1.MoType = m.MoType

	stage1.ObjectType = m.ObjectType

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.InventoryInventoryMoAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.InventoryInventoryMoAO0P0)
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

// Validate validates this inventory inventory mo a o0 p0
func (m *InventoryInventoryMoAO0P0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InventoryInventoryMoAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryInventoryMoAO0P0) UnmarshalBinary(b []byte) error {
	var res InventoryInventoryMoAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}