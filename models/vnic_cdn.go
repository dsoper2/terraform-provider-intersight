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

// VnicCdn Consistent Device Naming
//
// Consistent Device Naming configuration for the virtual NIC.
//
// swagger:model vnicCdn
type VnicCdn struct {
	VnicCdnAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VnicCdn) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 VnicCdnAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.VnicCdnAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VnicCdn) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.VnicCdnAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vnic cdn
func (m *VnicCdn) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with VnicCdnAO0P0
	if err := m.VnicCdnAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VnicCdn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VnicCdn) UnmarshalBinary(b []byte) error {
	var res VnicCdn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VnicCdnAO0P0 vnic cdn a o0 p0
// swagger:model VnicCdnAO0P0
type VnicCdnAO0P0 struct {

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// Source of the CDN. It can either be user specified or be the same as the vNIC name.
	//
	// Enum: [vnic user]
	Source *string `json:"Source,omitempty"`

	// The CDN value entered in case of user defined mode.
	//
	Value string `json:"Value,omitempty"`

	// vnic cdn a o0 p0
	VnicCdnAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *VnicCdnAO0P0) UnmarshalJSON(data []byte) error {
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

		// Source of the CDN. It can either be user specified or be the same as the vNIC name.
		//
		// Enum: [vnic user]
		Source *string `json:"Source,omitempty"`

		// The CDN value entered in case of user defined mode.
		//
		Value string `json:"Value,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv VnicCdnAO0P0

	rcv.ObjectType = stage1.ObjectType

	rcv.Source = stage1.Source

	rcv.Value = stage1.Value

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "ObjectType")

	delete(stage2, "Source")

	delete(stage2, "Value")

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
		m.VnicCdnAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m VnicCdnAO0P0) MarshalJSON() ([]byte, error) {
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

		// Source of the CDN. It can either be user specified or be the same as the vNIC name.
		//
		// Enum: [vnic user]
		Source *string `json:"Source,omitempty"`

		// The CDN value entered in case of user defined mode.
		//
		Value string `json:"Value,omitempty"`
	}

	stage1.ObjectType = m.ObjectType

	stage1.Source = m.Source

	stage1.Value = m.Value

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.VnicCdnAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.VnicCdnAO0P0)
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

// Validate validates this vnic cdn a o0 p0
func (m *VnicCdnAO0P0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var vnicCdnAO0P0TypeSourcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["vnic","user"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vnicCdnAO0P0TypeSourcePropEnum = append(vnicCdnAO0P0TypeSourcePropEnum, v)
	}
}

const (

	// VnicCdnAO0P0SourceVnic captures enum value "vnic"
	VnicCdnAO0P0SourceVnic string = "vnic"

	// VnicCdnAO0P0SourceUser captures enum value "user"
	VnicCdnAO0P0SourceUser string = "user"
)

// prop value enum
func (m *VnicCdnAO0P0) validateSourceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, vnicCdnAO0P0TypeSourcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VnicCdnAO0P0) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(m.Source) { // not required
		return nil
	}

	// value enum
	if err := m.validateSourceEnum("Source", "body", *m.Source); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VnicCdnAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VnicCdnAO0P0) UnmarshalBinary(b []byte) error {
	var res VnicCdnAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
