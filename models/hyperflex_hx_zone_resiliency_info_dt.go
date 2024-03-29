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

// HyperflexHxZoneResiliencyInfoDt Hyperflex:Hx Zone Resiliency Info Dt
// swagger:model hyperflexHxZoneResiliencyInfoDt
type HyperflexHxZoneResiliencyInfoDt struct {
	HyperflexHxZoneResiliencyInfoDtAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexHxZoneResiliencyInfoDt) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 HyperflexHxZoneResiliencyInfoDtAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.HyperflexHxZoneResiliencyInfoDtAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexHxZoneResiliencyInfoDt) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.HyperflexHxZoneResiliencyInfoDtAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex hx zone resiliency info dt
func (m *HyperflexHxZoneResiliencyInfoDt) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with HyperflexHxZoneResiliencyInfoDtAO0P0
	if err := m.HyperflexHxZoneResiliencyInfoDtAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexHxZoneResiliencyInfoDt) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexHxZoneResiliencyInfoDt) UnmarshalBinary(b []byte) error {
	var res HyperflexHxZoneResiliencyInfoDt
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HyperflexHxZoneResiliencyInfoDtAO0P0 hyperflex hx zone resiliency info dt a o0 p0
// swagger:model HyperflexHxZoneResiliencyInfoDtAO0P0
type HyperflexHxZoneResiliencyInfoDtAO0P0 struct {

	// name
	// Read Only: true
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

	// resiliency info
	// Read Only: true
	ResiliencyInfo *HyperflexHxResiliencyInfoDt `json:"ResiliencyInfo,omitempty"`

	// hyperflex hx zone resiliency info dt a o0 p0
	HyperflexHxZoneResiliencyInfoDtAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *HyperflexHxZoneResiliencyInfoDtAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// name
		// Read Only: true
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

		// resiliency info
		// Read Only: true
		ResiliencyInfo *HyperflexHxResiliencyInfoDt `json:"ResiliencyInfo,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv HyperflexHxZoneResiliencyInfoDtAO0P0

	rcv.Name = stage1.Name

	rcv.ObjectType = stage1.ObjectType

	rcv.ResiliencyInfo = stage1.ResiliencyInfo

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Name")

	delete(stage2, "ObjectType")

	delete(stage2, "ResiliencyInfo")

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
		m.HyperflexHxZoneResiliencyInfoDtAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m HyperflexHxZoneResiliencyInfoDtAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// name
		// Read Only: true
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

		// resiliency info
		// Read Only: true
		ResiliencyInfo *HyperflexHxResiliencyInfoDt `json:"ResiliencyInfo,omitempty"`
	}

	stage1.Name = m.Name

	stage1.ObjectType = m.ObjectType

	stage1.ResiliencyInfo = m.ResiliencyInfo

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.HyperflexHxZoneResiliencyInfoDtAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.HyperflexHxZoneResiliencyInfoDtAO0P0)
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

// Validate validates this hyperflex hx zone resiliency info dt a o0 p0
func (m *HyperflexHxZoneResiliencyInfoDtAO0P0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResiliencyInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HyperflexHxZoneResiliencyInfoDtAO0P0) validateResiliencyInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.ResiliencyInfo) { // not required
		return nil
	}

	if m.ResiliencyInfo != nil {
		if err := m.ResiliencyInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ResiliencyInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexHxZoneResiliencyInfoDtAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexHxZoneResiliencyInfoDtAO0P0) UnmarshalBinary(b []byte) error {
	var res HyperflexHxZoneResiliencyInfoDtAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
