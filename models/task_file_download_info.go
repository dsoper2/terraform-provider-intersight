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

// TaskFileDownloadInfo Task:File Download Info
//
// Specifies a download path or location currently supports S3 currently.
//
// swagger:model taskFileDownloadInfo
type TaskFileDownloadInfo struct {
	TaskFileDownloadInfoAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TaskFileDownloadInfo) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TaskFileDownloadInfoAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TaskFileDownloadInfoAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TaskFileDownloadInfo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.TaskFileDownloadInfoAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this task file download info
func (m *TaskFileDownloadInfo) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TaskFileDownloadInfoAO0P0
	if err := m.TaskFileDownloadInfoAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TaskFileDownloadInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskFileDownloadInfo) UnmarshalBinary(b []byte) error {
	var res TaskFileDownloadInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TaskFileDownloadInfoAO0P0 task file download info a o0 p0
// swagger:model TaskFileDownloadInfoAO0P0
type TaskFileDownloadInfoAO0P0 struct {

	// When the type of the download is inline, it will read the file from the contents here.
	//
	Contents string `json:"Contents,omitempty"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// The path of the download from the specified source location and if type is S3, then this is the object key.
	//
	Path string `json:"Path,omitempty"`

	// The source of the download location and if type is S3, then this is the bucket name.
	//
	Source string `json:"Source,omitempty"`

	// The type of download location is captured in type.
	//
	// Enum: [S3 Local Inline]
	Type *string `json:"Type,omitempty"`

	// task file download info a o0 p0
	TaskFileDownloadInfoAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *TaskFileDownloadInfoAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// When the type of the download is inline, it will read the file from the contents here.
		//
		Contents string `json:"Contents,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// The path of the download from the specified source location and if type is S3, then this is the object key.
		//
		Path string `json:"Path,omitempty"`

		// The source of the download location and if type is S3, then this is the bucket name.
		//
		Source string `json:"Source,omitempty"`

		// The type of download location is captured in type.
		//
		// Enum: [S3 Local Inline]
		Type *string `json:"Type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv TaskFileDownloadInfoAO0P0

	rcv.Contents = stage1.Contents

	rcv.ObjectType = stage1.ObjectType

	rcv.Path = stage1.Path

	rcv.Source = stage1.Source

	rcv.Type = stage1.Type

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Contents")

	delete(stage2, "ObjectType")

	delete(stage2, "Path")

	delete(stage2, "Source")

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
		m.TaskFileDownloadInfoAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m TaskFileDownloadInfoAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// When the type of the download is inline, it will read the file from the contents here.
		//
		Contents string `json:"Contents,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// The path of the download from the specified source location and if type is S3, then this is the object key.
		//
		Path string `json:"Path,omitempty"`

		// The source of the download location and if type is S3, then this is the bucket name.
		//
		Source string `json:"Source,omitempty"`

		// The type of download location is captured in type.
		//
		// Enum: [S3 Local Inline]
		Type *string `json:"Type,omitempty"`
	}

	stage1.Contents = m.Contents

	stage1.ObjectType = m.ObjectType

	stage1.Path = m.Path

	stage1.Source = m.Source

	stage1.Type = m.Type

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.TaskFileDownloadInfoAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.TaskFileDownloadInfoAO0P0)
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

// Validate validates this task file download info a o0 p0
func (m *TaskFileDownloadInfoAO0P0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var taskFileDownloadInfoAO0P0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["S3","Local","Inline"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		taskFileDownloadInfoAO0P0TypeTypePropEnum = append(taskFileDownloadInfoAO0P0TypeTypePropEnum, v)
	}
}

const (

	// TaskFileDownloadInfoAO0P0TypeS3 captures enum value "S3"
	TaskFileDownloadInfoAO0P0TypeS3 string = "S3"

	// TaskFileDownloadInfoAO0P0TypeLocal captures enum value "Local"
	TaskFileDownloadInfoAO0P0TypeLocal string = "Local"

	// TaskFileDownloadInfoAO0P0TypeInline captures enum value "Inline"
	TaskFileDownloadInfoAO0P0TypeInline string = "Inline"
)

// prop value enum
func (m *TaskFileDownloadInfoAO0P0) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, taskFileDownloadInfoAO0P0TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TaskFileDownloadInfoAO0P0) validateType(formats strfmt.Registry) error {

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
func (m *TaskFileDownloadInfoAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskFileDownloadInfoAO0P0) UnmarshalBinary(b []byte) error {
	var res TaskFileDownloadInfoAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
