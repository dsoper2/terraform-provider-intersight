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

// FeedbackFeedbackData Feedback:Feedback Data
//
// Feedback data that collected on the UI from user.
//
// swagger:model feedbackFeedbackData
type FeedbackFeedbackData struct {
	FeedbackFeedbackDataAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *FeedbackFeedbackData) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 FeedbackFeedbackDataAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.FeedbackFeedbackDataAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m FeedbackFeedbackData) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.FeedbackFeedbackDataAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this feedback feedback data
func (m *FeedbackFeedbackData) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with FeedbackFeedbackDataAO0P0
	if err := m.FeedbackFeedbackDataAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *FeedbackFeedbackData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeedbackFeedbackData) UnmarshalBinary(b []byte) error {
	var res FeedbackFeedbackData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FeedbackFeedbackDataAO0P0 feedback feedback data a o0 p0
// swagger:model FeedbackFeedbackDataAO0P0
type FeedbackFeedbackDataAO0P0 struct {

	// Text of the feedback as provided by the user, if it is a bug or a comment.
	//
	Comment string `json:"Comment,omitempty"`

	// User's email address details.
	//
	Email string `json:"Email,omitempty"`

	// Evalation rating as provided by the user to capture user sentiment regarding the issue.
	//
	// Enum: [Excellent Poor Fair Good]
	Evaluation *string `json:"Evaluation,omitempty"`

	// If a user is open for follow-up or not.
	//
	FollowUp *bool `json:"FollowUp,omitempty"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// Bunch of last traceId for reproducing user last activity.
	//
	TraceIds interface{} `json:"TraceIds,omitempty"`

	// Type of the feedback from user.
	//
	// Enum: [Evaluation Bug]
	Type *string `json:"Type,omitempty"`

	// feedback feedback data a o0 p0
	FeedbackFeedbackDataAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *FeedbackFeedbackDataAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Text of the feedback as provided by the user, if it is a bug or a comment.
		//
		Comment string `json:"Comment,omitempty"`

		// User's email address details.
		//
		Email string `json:"Email,omitempty"`

		// Evalation rating as provided by the user to capture user sentiment regarding the issue.
		//
		// Enum: [Excellent Poor Fair Good]
		Evaluation *string `json:"Evaluation,omitempty"`

		// If a user is open for follow-up or not.
		//
		FollowUp *bool `json:"FollowUp,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// Bunch of last traceId for reproducing user last activity.
		//
		TraceIds interface{} `json:"TraceIds,omitempty"`

		// Type of the feedback from user.
		//
		// Enum: [Evaluation Bug]
		Type *string `json:"Type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv FeedbackFeedbackDataAO0P0

	rcv.Comment = stage1.Comment

	rcv.Email = stage1.Email

	rcv.Evaluation = stage1.Evaluation

	rcv.FollowUp = stage1.FollowUp

	rcv.ObjectType = stage1.ObjectType

	rcv.TraceIds = stage1.TraceIds

	rcv.Type = stage1.Type

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Comment")

	delete(stage2, "Email")

	delete(stage2, "Evaluation")

	delete(stage2, "FollowUp")

	delete(stage2, "ObjectType")

	delete(stage2, "TraceIds")

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
		m.FeedbackFeedbackDataAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m FeedbackFeedbackDataAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Text of the feedback as provided by the user, if it is a bug or a comment.
		//
		Comment string `json:"Comment,omitempty"`

		// User's email address details.
		//
		Email string `json:"Email,omitempty"`

		// Evalation rating as provided by the user to capture user sentiment regarding the issue.
		//
		// Enum: [Excellent Poor Fair Good]
		Evaluation *string `json:"Evaluation,omitempty"`

		// If a user is open for follow-up or not.
		//
		FollowUp *bool `json:"FollowUp,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// Bunch of last traceId for reproducing user last activity.
		//
		TraceIds interface{} `json:"TraceIds,omitempty"`

		// Type of the feedback from user.
		//
		// Enum: [Evaluation Bug]
		Type *string `json:"Type,omitempty"`
	}

	stage1.Comment = m.Comment

	stage1.Email = m.Email

	stage1.Evaluation = m.Evaluation

	stage1.FollowUp = m.FollowUp

	stage1.ObjectType = m.ObjectType

	stage1.TraceIds = m.TraceIds

	stage1.Type = m.Type

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.FeedbackFeedbackDataAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.FeedbackFeedbackDataAO0P0)
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

// Validate validates this feedback feedback data a o0 p0
func (m *FeedbackFeedbackDataAO0P0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvaluation(formats); err != nil {
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

var feedbackFeedbackDataAO0P0TypeEvaluationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Excellent","Poor","Fair","Good"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		feedbackFeedbackDataAO0P0TypeEvaluationPropEnum = append(feedbackFeedbackDataAO0P0TypeEvaluationPropEnum, v)
	}
}

const (

	// FeedbackFeedbackDataAO0P0EvaluationExcellent captures enum value "Excellent"
	FeedbackFeedbackDataAO0P0EvaluationExcellent string = "Excellent"

	// FeedbackFeedbackDataAO0P0EvaluationPoor captures enum value "Poor"
	FeedbackFeedbackDataAO0P0EvaluationPoor string = "Poor"

	// FeedbackFeedbackDataAO0P0EvaluationFair captures enum value "Fair"
	FeedbackFeedbackDataAO0P0EvaluationFair string = "Fair"

	// FeedbackFeedbackDataAO0P0EvaluationGood captures enum value "Good"
	FeedbackFeedbackDataAO0P0EvaluationGood string = "Good"
)

// prop value enum
func (m *FeedbackFeedbackDataAO0P0) validateEvaluationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, feedbackFeedbackDataAO0P0TypeEvaluationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FeedbackFeedbackDataAO0P0) validateEvaluation(formats strfmt.Registry) error {

	if swag.IsZero(m.Evaluation) { // not required
		return nil
	}

	// value enum
	if err := m.validateEvaluationEnum("Evaluation", "body", *m.Evaluation); err != nil {
		return err
	}

	return nil
}

var feedbackFeedbackDataAO0P0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Evaluation","Bug"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		feedbackFeedbackDataAO0P0TypeTypePropEnum = append(feedbackFeedbackDataAO0P0TypeTypePropEnum, v)
	}
}

const (

	// FeedbackFeedbackDataAO0P0TypeEvaluation captures enum value "Evaluation"
	FeedbackFeedbackDataAO0P0TypeEvaluation string = "Evaluation"

	// FeedbackFeedbackDataAO0P0TypeBug captures enum value "Bug"
	FeedbackFeedbackDataAO0P0TypeBug string = "Bug"
)

// prop value enum
func (m *FeedbackFeedbackDataAO0P0) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, feedbackFeedbackDataAO0P0TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FeedbackFeedbackDataAO0P0) validateType(formats strfmt.Registry) error {

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
func (m *FeedbackFeedbackDataAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeedbackFeedbackDataAO0P0) UnmarshalBinary(b []byte) error {
	var res FeedbackFeedbackDataAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
