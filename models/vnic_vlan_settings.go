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

// VnicVlanSettings VLAN Settings
//
// VLAN configuration for the virtual interface.
//
// swagger:model vnicVlanSettings
type VnicVlanSettings struct {
	VnicVlanSettingsAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VnicVlanSettings) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 VnicVlanSettingsAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.VnicVlanSettingsAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VnicVlanSettings) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.VnicVlanSettingsAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vnic vlan settings
func (m *VnicVlanSettings) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with VnicVlanSettingsAO0P0
	if err := m.VnicVlanSettingsAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VnicVlanSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VnicVlanSettings) UnmarshalBinary(b []byte) error {
	var res VnicVlanSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VnicVlanSettingsAO0P0 vnic vlan settings a o0 p0
// swagger:model VnicVlanSettingsAO0P0
type VnicVlanSettingsAO0P0 struct {

	// Default VLAN ID of the virtual interface. Setting the ID to 0 will not associate any default VLAN to the traffic on the virtual interface.
	//
	DefaultVlan int64 `json:"DefaultVlan,omitempty"`

	// Option to determine if the port can carry single VLAN (Access) or multiple VLANs (Trunk) traffic.
	//
	// Enum: [ACCESS TRUNK]
	Mode *string `json:"Mode,omitempty"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// vnic vlan settings a o0 p0
	VnicVlanSettingsAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *VnicVlanSettingsAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Default VLAN ID of the virtual interface. Setting the ID to 0 will not associate any default VLAN to the traffic on the virtual interface.
		//
		DefaultVlan int64 `json:"DefaultVlan,omitempty"`

		// Option to determine if the port can carry single VLAN (Access) or multiple VLANs (Trunk) traffic.
		//
		// Enum: [ACCESS TRUNK]
		Mode *string `json:"Mode,omitempty"`

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
	var rcv VnicVlanSettingsAO0P0

	rcv.DefaultVlan = stage1.DefaultVlan

	rcv.Mode = stage1.Mode

	rcv.ObjectType = stage1.ObjectType

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "DefaultVlan")

	delete(stage2, "Mode")

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
		m.VnicVlanSettingsAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m VnicVlanSettingsAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Default VLAN ID of the virtual interface. Setting the ID to 0 will not associate any default VLAN to the traffic on the virtual interface.
		//
		DefaultVlan int64 `json:"DefaultVlan,omitempty"`

		// Option to determine if the port can carry single VLAN (Access) or multiple VLANs (Trunk) traffic.
		//
		// Enum: [ACCESS TRUNK]
		Mode *string `json:"Mode,omitempty"`

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

	stage1.DefaultVlan = m.DefaultVlan

	stage1.Mode = m.Mode

	stage1.ObjectType = m.ObjectType

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.VnicVlanSettingsAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.VnicVlanSettingsAO0P0)
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

// Validate validates this vnic vlan settings a o0 p0
func (m *VnicVlanSettingsAO0P0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var vnicVlanSettingsAO0P0TypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACCESS","TRUNK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vnicVlanSettingsAO0P0TypeModePropEnum = append(vnicVlanSettingsAO0P0TypeModePropEnum, v)
	}
}

const (

	// VnicVlanSettingsAO0P0ModeACCESS captures enum value "ACCESS"
	VnicVlanSettingsAO0P0ModeACCESS string = "ACCESS"

	// VnicVlanSettingsAO0P0ModeTRUNK captures enum value "TRUNK"
	VnicVlanSettingsAO0P0ModeTRUNK string = "TRUNK"
)

// prop value enum
func (m *VnicVlanSettingsAO0P0) validateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, vnicVlanSettingsAO0P0TypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VnicVlanSettingsAO0P0) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("Mode", "body", *m.Mode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VnicVlanSettingsAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VnicVlanSettingsAO0P0) UnmarshalBinary(b []byte) error {
	var res VnicVlanSettingsAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
