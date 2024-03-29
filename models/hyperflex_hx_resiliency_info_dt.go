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

// HyperflexHxResiliencyInfoDt Hyperflex:Hx Resiliency Info Dt
// swagger:model hyperflexHxResiliencyInfoDt
type HyperflexHxResiliencyInfoDt struct {
	HyperflexHxResiliencyInfoDtAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexHxResiliencyInfoDt) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 HyperflexHxResiliencyInfoDtAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.HyperflexHxResiliencyInfoDtAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexHxResiliencyInfoDt) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.HyperflexHxResiliencyInfoDtAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex hx resiliency info dt
func (m *HyperflexHxResiliencyInfoDt) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with HyperflexHxResiliencyInfoDtAO0P0
	if err := m.HyperflexHxResiliencyInfoDtAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexHxResiliencyInfoDt) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexHxResiliencyInfoDt) UnmarshalBinary(b []byte) error {
	var res HyperflexHxResiliencyInfoDt
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HyperflexHxResiliencyInfoDtAO0P0 hyperflex hx resiliency info dt a o0 p0
// swagger:model HyperflexHxResiliencyInfoDtAO0P0
type HyperflexHxResiliencyInfoDtAO0P0 struct {

	// data replication factor
	// Read Only: true
	// Enum: [ONE_COPY TWO_COPIES THREE_COPIES FOUR_COPIES SIX_COPIES]
	DataReplicationFactor string `json:"DataReplicationFactor,omitempty"`

	// hdd failures tolerable
	// Read Only: true
	HddFailuresTolerable int64 `json:"HddFailuresTolerable,omitempty"`

	// messages
	// Read Only: true
	Messages []string `json:"Messages"`

	// node failures tolerable
	// Read Only: true
	NodeFailuresTolerable int64 `json:"NodeFailuresTolerable,omitempty"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// policy compliance
	// Read Only: true
	// Enum: [UNKNOWN COMPLIANT NON_COMPLIANT]
	PolicyCompliance string `json:"PolicyCompliance,omitempty"`

	// resiliency state
	// Read Only: true
	// Enum: [UNKNOWN HEALTHY WARNING OFFLINE]
	ResiliencyState string `json:"ResiliencyState,omitempty"`

	// ssd failures tolerable
	// Read Only: true
	SsdFailuresTolerable int64 `json:"SsdFailuresTolerable,omitempty"`

	// hyperflex hx resiliency info dt a o0 p0
	HyperflexHxResiliencyInfoDtAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *HyperflexHxResiliencyInfoDtAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// data replication factor
		// Read Only: true
		// Enum: [ONE_COPY TWO_COPIES THREE_COPIES FOUR_COPIES SIX_COPIES]
		DataReplicationFactor string `json:"DataReplicationFactor,omitempty"`

		// hdd failures tolerable
		// Read Only: true
		HddFailuresTolerable int64 `json:"HddFailuresTolerable,omitempty"`

		// messages
		// Read Only: true
		Messages []string `json:"Messages"`

		// node failures tolerable
		// Read Only: true
		NodeFailuresTolerable int64 `json:"NodeFailuresTolerable,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// policy compliance
		// Read Only: true
		// Enum: [UNKNOWN COMPLIANT NON_COMPLIANT]
		PolicyCompliance string `json:"PolicyCompliance,omitempty"`

		// resiliency state
		// Read Only: true
		// Enum: [UNKNOWN HEALTHY WARNING OFFLINE]
		ResiliencyState string `json:"ResiliencyState,omitempty"`

		// ssd failures tolerable
		// Read Only: true
		SsdFailuresTolerable int64 `json:"SsdFailuresTolerable,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv HyperflexHxResiliencyInfoDtAO0P0

	rcv.DataReplicationFactor = stage1.DataReplicationFactor

	rcv.HddFailuresTolerable = stage1.HddFailuresTolerable

	rcv.Messages = stage1.Messages

	rcv.NodeFailuresTolerable = stage1.NodeFailuresTolerable

	rcv.ObjectType = stage1.ObjectType

	rcv.PolicyCompliance = stage1.PolicyCompliance

	rcv.ResiliencyState = stage1.ResiliencyState

	rcv.SsdFailuresTolerable = stage1.SsdFailuresTolerable

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "DataReplicationFactor")

	delete(stage2, "HddFailuresTolerable")

	delete(stage2, "Messages")

	delete(stage2, "NodeFailuresTolerable")

	delete(stage2, "ObjectType")

	delete(stage2, "PolicyCompliance")

	delete(stage2, "ResiliencyState")

	delete(stage2, "SsdFailuresTolerable")

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
		m.HyperflexHxResiliencyInfoDtAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m HyperflexHxResiliencyInfoDtAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// data replication factor
		// Read Only: true
		// Enum: [ONE_COPY TWO_COPIES THREE_COPIES FOUR_COPIES SIX_COPIES]
		DataReplicationFactor string `json:"DataReplicationFactor,omitempty"`

		// hdd failures tolerable
		// Read Only: true
		HddFailuresTolerable int64 `json:"HddFailuresTolerable,omitempty"`

		// messages
		// Read Only: true
		Messages []string `json:"Messages"`

		// node failures tolerable
		// Read Only: true
		NodeFailuresTolerable int64 `json:"NodeFailuresTolerable,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// policy compliance
		// Read Only: true
		// Enum: [UNKNOWN COMPLIANT NON_COMPLIANT]
		PolicyCompliance string `json:"PolicyCompliance,omitempty"`

		// resiliency state
		// Read Only: true
		// Enum: [UNKNOWN HEALTHY WARNING OFFLINE]
		ResiliencyState string `json:"ResiliencyState,omitempty"`

		// ssd failures tolerable
		// Read Only: true
		SsdFailuresTolerable int64 `json:"SsdFailuresTolerable,omitempty"`
	}

	stage1.DataReplicationFactor = m.DataReplicationFactor

	stage1.HddFailuresTolerable = m.HddFailuresTolerable

	stage1.Messages = m.Messages

	stage1.NodeFailuresTolerable = m.NodeFailuresTolerable

	stage1.ObjectType = m.ObjectType

	stage1.PolicyCompliance = m.PolicyCompliance

	stage1.ResiliencyState = m.ResiliencyState

	stage1.SsdFailuresTolerable = m.SsdFailuresTolerable

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.HyperflexHxResiliencyInfoDtAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.HyperflexHxResiliencyInfoDtAO0P0)
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

// Validate validates this hyperflex hx resiliency info dt a o0 p0
func (m *HyperflexHxResiliencyInfoDtAO0P0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataReplicationFactor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyCompliance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResiliencyState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var hyperflexHxResiliencyInfoDtAO0P0TypeDataReplicationFactorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ONE_COPY","TWO_COPIES","THREE_COPIES","FOUR_COPIES","SIX_COPIES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hyperflexHxResiliencyInfoDtAO0P0TypeDataReplicationFactorPropEnum = append(hyperflexHxResiliencyInfoDtAO0P0TypeDataReplicationFactorPropEnum, v)
	}
}

const (

	// HyperflexHxResiliencyInfoDtAO0P0DataReplicationFactorONECOPY captures enum value "ONE_COPY"
	HyperflexHxResiliencyInfoDtAO0P0DataReplicationFactorONECOPY string = "ONE_COPY"

	// HyperflexHxResiliencyInfoDtAO0P0DataReplicationFactorTWOCOPIES captures enum value "TWO_COPIES"
	HyperflexHxResiliencyInfoDtAO0P0DataReplicationFactorTWOCOPIES string = "TWO_COPIES"

	// HyperflexHxResiliencyInfoDtAO0P0DataReplicationFactorTHREECOPIES captures enum value "THREE_COPIES"
	HyperflexHxResiliencyInfoDtAO0P0DataReplicationFactorTHREECOPIES string = "THREE_COPIES"

	// HyperflexHxResiliencyInfoDtAO0P0DataReplicationFactorFOURCOPIES captures enum value "FOUR_COPIES"
	HyperflexHxResiliencyInfoDtAO0P0DataReplicationFactorFOURCOPIES string = "FOUR_COPIES"

	// HyperflexHxResiliencyInfoDtAO0P0DataReplicationFactorSIXCOPIES captures enum value "SIX_COPIES"
	HyperflexHxResiliencyInfoDtAO0P0DataReplicationFactorSIXCOPIES string = "SIX_COPIES"
)

// prop value enum
func (m *HyperflexHxResiliencyInfoDtAO0P0) validateDataReplicationFactorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hyperflexHxResiliencyInfoDtAO0P0TypeDataReplicationFactorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HyperflexHxResiliencyInfoDtAO0P0) validateDataReplicationFactor(formats strfmt.Registry) error {

	if swag.IsZero(m.DataReplicationFactor) { // not required
		return nil
	}

	// value enum
	if err := m.validateDataReplicationFactorEnum("DataReplicationFactor", "body", m.DataReplicationFactor); err != nil {
		return err
	}

	return nil
}

var hyperflexHxResiliencyInfoDtAO0P0TypePolicyCompliancePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNKNOWN","COMPLIANT","NON_COMPLIANT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hyperflexHxResiliencyInfoDtAO0P0TypePolicyCompliancePropEnum = append(hyperflexHxResiliencyInfoDtAO0P0TypePolicyCompliancePropEnum, v)
	}
}

const (

	// HyperflexHxResiliencyInfoDtAO0P0PolicyComplianceUNKNOWN captures enum value "UNKNOWN"
	HyperflexHxResiliencyInfoDtAO0P0PolicyComplianceUNKNOWN string = "UNKNOWN"

	// HyperflexHxResiliencyInfoDtAO0P0PolicyComplianceCOMPLIANT captures enum value "COMPLIANT"
	HyperflexHxResiliencyInfoDtAO0P0PolicyComplianceCOMPLIANT string = "COMPLIANT"

	// HyperflexHxResiliencyInfoDtAO0P0PolicyComplianceNONCOMPLIANT captures enum value "NON_COMPLIANT"
	HyperflexHxResiliencyInfoDtAO0P0PolicyComplianceNONCOMPLIANT string = "NON_COMPLIANT"
)

// prop value enum
func (m *HyperflexHxResiliencyInfoDtAO0P0) validatePolicyComplianceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hyperflexHxResiliencyInfoDtAO0P0TypePolicyCompliancePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HyperflexHxResiliencyInfoDtAO0P0) validatePolicyCompliance(formats strfmt.Registry) error {

	if swag.IsZero(m.PolicyCompliance) { // not required
		return nil
	}

	// value enum
	if err := m.validatePolicyComplianceEnum("PolicyCompliance", "body", m.PolicyCompliance); err != nil {
		return err
	}

	return nil
}

var hyperflexHxResiliencyInfoDtAO0P0TypeResiliencyStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNKNOWN","HEALTHY","WARNING","OFFLINE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hyperflexHxResiliencyInfoDtAO0P0TypeResiliencyStatePropEnum = append(hyperflexHxResiliencyInfoDtAO0P0TypeResiliencyStatePropEnum, v)
	}
}

const (

	// HyperflexHxResiliencyInfoDtAO0P0ResiliencyStateUNKNOWN captures enum value "UNKNOWN"
	HyperflexHxResiliencyInfoDtAO0P0ResiliencyStateUNKNOWN string = "UNKNOWN"

	// HyperflexHxResiliencyInfoDtAO0P0ResiliencyStateHEALTHY captures enum value "HEALTHY"
	HyperflexHxResiliencyInfoDtAO0P0ResiliencyStateHEALTHY string = "HEALTHY"

	// HyperflexHxResiliencyInfoDtAO0P0ResiliencyStateWARNING captures enum value "WARNING"
	HyperflexHxResiliencyInfoDtAO0P0ResiliencyStateWARNING string = "WARNING"

	// HyperflexHxResiliencyInfoDtAO0P0ResiliencyStateOFFLINE captures enum value "OFFLINE"
	HyperflexHxResiliencyInfoDtAO0P0ResiliencyStateOFFLINE string = "OFFLINE"
)

// prop value enum
func (m *HyperflexHxResiliencyInfoDtAO0P0) validateResiliencyStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hyperflexHxResiliencyInfoDtAO0P0TypeResiliencyStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HyperflexHxResiliencyInfoDtAO0P0) validateResiliencyState(formats strfmt.Registry) error {

	if swag.IsZero(m.ResiliencyState) { // not required
		return nil
	}

	// value enum
	if err := m.validateResiliencyStateEnum("ResiliencyState", "body", m.ResiliencyState); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexHxResiliencyInfoDtAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexHxResiliencyInfoDtAO0P0) UnmarshalBinary(b []byte) error {
	var res HyperflexHxResiliencyInfoDtAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
