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

// AssetProductInformation Asset:Product Information
//
// Type for saving the product information.
//
// swagger:model assetProductInformation
type AssetProductInformation struct {
	AssetProductInformationAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AssetProductInformation) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 AssetProductInformationAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.AssetProductInformationAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AssetProductInformation) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.AssetProductInformationAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this asset product information
func (m *AssetProductInformation) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AssetProductInformationAO0P0
	if err := m.AssetProductInformationAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AssetProductInformation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetProductInformation) UnmarshalBinary(b []byte) error {
	var res AssetProductInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AssetProductInformationAO0P0 asset product information a o0 p0
// swagger:model AssetProductInformationAO0P0
type AssetProductInformationAO0P0 struct {

	// Billing address provided by customer while buying this Cisco product.
	//
	// Read Only: true
	BillTo *AssetAddressInformation `json:"BillTo,omitempty"`

	// Short description of the Cisco product that helps identify the product easily. example "DISTI:UCS 6248UP 1RU Fabric Int/No PSU/32 UP/ 12p LIC".
	//
	// Read Only: true
	Description string `json:"Description,omitempty"`

	// Family that the product belongs to. Example "UCSB".
	//
	// Read Only: true
	Family string `json:"Family,omitempty"`

	// Group that the product belongs to. It is one higher level categorization above family. example "Switch".
	//
	// Read Only: true
	Group string `json:"Group,omitempty"`

	// Product number that identifies the product. example PID. example "UCS-FI-6248UP-CH2".
	//
	// Read Only: true
	Number string `json:"Number,omitempty"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// Shipping address provided by customer while buying this Cisco product.
	//
	// Read Only: true
	ShipTo *AssetAddressInformation `json:"ShipTo,omitempty"`

	// Sub type of the product being specified. example "UCS 6200 SER".
	//
	// Read Only: true
	SubType string `json:"SubType,omitempty"`

	// asset product information a o0 p0
	AssetProductInformationAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *AssetProductInformationAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Billing address provided by customer while buying this Cisco product.
		//
		// Read Only: true
		BillTo *AssetAddressInformation `json:"BillTo,omitempty"`

		// Short description of the Cisco product that helps identify the product easily. example "DISTI:UCS 6248UP 1RU Fabric Int/No PSU/32 UP/ 12p LIC".
		//
		// Read Only: true
		Description string `json:"Description,omitempty"`

		// Family that the product belongs to. Example "UCSB".
		//
		// Read Only: true
		Family string `json:"Family,omitempty"`

		// Group that the product belongs to. It is one higher level categorization above family. example "Switch".
		//
		// Read Only: true
		Group string `json:"Group,omitempty"`

		// Product number that identifies the product. example PID. example "UCS-FI-6248UP-CH2".
		//
		// Read Only: true
		Number string `json:"Number,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// Shipping address provided by customer while buying this Cisco product.
		//
		// Read Only: true
		ShipTo *AssetAddressInformation `json:"ShipTo,omitempty"`

		// Sub type of the product being specified. example "UCS 6200 SER".
		//
		// Read Only: true
		SubType string `json:"SubType,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv AssetProductInformationAO0P0

	rcv.BillTo = stage1.BillTo

	rcv.Description = stage1.Description

	rcv.Family = stage1.Family

	rcv.Group = stage1.Group

	rcv.Number = stage1.Number

	rcv.ObjectType = stage1.ObjectType

	rcv.ShipTo = stage1.ShipTo

	rcv.SubType = stage1.SubType

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "BillTo")

	delete(stage2, "Description")

	delete(stage2, "Family")

	delete(stage2, "Group")

	delete(stage2, "Number")

	delete(stage2, "ObjectType")

	delete(stage2, "ShipTo")

	delete(stage2, "SubType")

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
		m.AssetProductInformationAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m AssetProductInformationAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Billing address provided by customer while buying this Cisco product.
		//
		// Read Only: true
		BillTo *AssetAddressInformation `json:"BillTo,omitempty"`

		// Short description of the Cisco product that helps identify the product easily. example "DISTI:UCS 6248UP 1RU Fabric Int/No PSU/32 UP/ 12p LIC".
		//
		// Read Only: true
		Description string `json:"Description,omitempty"`

		// Family that the product belongs to. Example "UCSB".
		//
		// Read Only: true
		Family string `json:"Family,omitempty"`

		// Group that the product belongs to. It is one higher level categorization above family. example "Switch".
		//
		// Read Only: true
		Group string `json:"Group,omitempty"`

		// Product number that identifies the product. example PID. example "UCS-FI-6248UP-CH2".
		//
		// Read Only: true
		Number string `json:"Number,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// Shipping address provided by customer while buying this Cisco product.
		//
		// Read Only: true
		ShipTo *AssetAddressInformation `json:"ShipTo,omitempty"`

		// Sub type of the product being specified. example "UCS 6200 SER".
		//
		// Read Only: true
		SubType string `json:"SubType,omitempty"`
	}

	stage1.BillTo = m.BillTo

	stage1.Description = m.Description

	stage1.Family = m.Family

	stage1.Group = m.Group

	stage1.Number = m.Number

	stage1.ObjectType = m.ObjectType

	stage1.ShipTo = m.ShipTo

	stage1.SubType = m.SubType

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.AssetProductInformationAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.AssetProductInformationAO0P0)
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

// Validate validates this asset product information a o0 p0
func (m *AssetProductInformationAO0P0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetProductInformationAO0P0) validateBillTo(formats strfmt.Registry) error {

	if swag.IsZero(m.BillTo) { // not required
		return nil
	}

	if m.BillTo != nil {
		if err := m.BillTo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BillTo")
			}
			return err
		}
	}

	return nil
}

func (m *AssetProductInformationAO0P0) validateShipTo(formats strfmt.Registry) error {

	if swag.IsZero(m.ShipTo) { // not required
		return nil
	}

	if m.ShipTo != nil {
		if err := m.ShipTo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShipTo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssetProductInformationAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetProductInformationAO0P0) UnmarshalBinary(b []byte) error {
	var res AssetProductInformationAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
