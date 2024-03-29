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

// StorageVirtualDriveConfig Virtual Drive
//
// This type models a single virtual drive that needs to be created through this policy.
//
// swagger:model storageVirtualDriveConfig
type StorageVirtualDriveConfig struct {
	StorageVirtualDriveConfigAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StorageVirtualDriveConfig) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 StorageVirtualDriveConfigAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.StorageVirtualDriveConfigAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StorageVirtualDriveConfig) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.StorageVirtualDriveConfigAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage virtual drive config
func (m *StorageVirtualDriveConfig) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with StorageVirtualDriveConfigAO0P0
	if err := m.StorageVirtualDriveConfigAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *StorageVirtualDriveConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageVirtualDriveConfig) UnmarshalBinary(b []byte) error {
	var res StorageVirtualDriveConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StorageVirtualDriveConfigAO0P0 storage virtual drive config a o0 p0
// swagger:model StorageVirtualDriveConfigAO0P0
type StorageVirtualDriveConfigAO0P0 struct {

	// This property holds the access policy that host has on this virtual drive.
	//
	// Enum: [Default ReadWrite ReadOnly Blocked]
	AccessPolicy *string `json:"AccessPolicy,omitempty"`

	// This flag enables the use of this virtual drive as a boot drive.
	//
	BootDrive *bool `json:"BootDrive,omitempty"`

	// Disk group policy that has the disk group in which this virtual drive needs to be created.
	//
	// Read Only: true
	DiskGroupName string `json:"DiskGroupName,omitempty"`

	// Disk group policy that has the disk group in which this virtual drive needs to be created.
	//
	DiskGroupPolicy string `json:"DiskGroupPolicy,omitempty"`

	// This property expect disk cache policy.
	//
	// Enum: [Default NoChange Enable Disable]
	DriveCache *string `json:"DriveCache,omitempty"`

	// This flag enables this virtual drive to use all the available space in the disk group. When this flag is configured, the size property is ignored.
	//
	ExpandToAvailable *bool `json:"ExpandToAvailable,omitempty"`

	// This property expects the desired IO mode - direct IO or cached IO.
	//
	// Enum: [Default Direct Cached]
	IoPolicy *string `json:"IoPolicy,omitempty"`

	// The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen), _ (underscore), : (colon), and . (period) are not allowed.
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

	// This property holds the read ahead mode to be used.
	//
	// Enum: [Default ReadAhead NoReadAhead]
	ReadPolicy *string `json:"ReadPolicy,omitempty"`

	// Virtual drive size in MB. This is a required field unless the 'Expand to Available' option is enabled.
	//
	Size int64 `json:"Size,omitempty"`

	// This property holds the write mode used to write the data in this virtual drive.
	//
	// Enum: [Default WriteThrough WriteBackGoodBbu AlwaysWriteBack]
	WritePolicy *string `json:"WritePolicy,omitempty"`

	// storage virtual drive config a o0 p0
	StorageVirtualDriveConfigAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *StorageVirtualDriveConfigAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// This property holds the access policy that host has on this virtual drive.
		//
		// Enum: [Default ReadWrite ReadOnly Blocked]
		AccessPolicy *string `json:"AccessPolicy,omitempty"`

		// This flag enables the use of this virtual drive as a boot drive.
		//
		BootDrive *bool `json:"BootDrive,omitempty"`

		// Disk group policy that has the disk group in which this virtual drive needs to be created.
		//
		// Read Only: true
		DiskGroupName string `json:"DiskGroupName,omitempty"`

		// Disk group policy that has the disk group in which this virtual drive needs to be created.
		//
		DiskGroupPolicy string `json:"DiskGroupPolicy,omitempty"`

		// This property expect disk cache policy.
		//
		// Enum: [Default NoChange Enable Disable]
		DriveCache *string `json:"DriveCache,omitempty"`

		// This flag enables this virtual drive to use all the available space in the disk group. When this flag is configured, the size property is ignored.
		//
		ExpandToAvailable *bool `json:"ExpandToAvailable,omitempty"`

		// This property expects the desired IO mode - direct IO or cached IO.
		//
		// Enum: [Default Direct Cached]
		IoPolicy *string `json:"IoPolicy,omitempty"`

		// The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen), _ (underscore), : (colon), and . (period) are not allowed.
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

		// This property holds the read ahead mode to be used.
		//
		// Enum: [Default ReadAhead NoReadAhead]
		ReadPolicy *string `json:"ReadPolicy,omitempty"`

		// Virtual drive size in MB. This is a required field unless the 'Expand to Available' option is enabled.
		//
		Size int64 `json:"Size,omitempty"`

		// This property holds the write mode used to write the data in this virtual drive.
		//
		// Enum: [Default WriteThrough WriteBackGoodBbu AlwaysWriteBack]
		WritePolicy *string `json:"WritePolicy,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv StorageVirtualDriveConfigAO0P0

	rcv.AccessPolicy = stage1.AccessPolicy

	rcv.BootDrive = stage1.BootDrive

	rcv.DiskGroupName = stage1.DiskGroupName

	rcv.DiskGroupPolicy = stage1.DiskGroupPolicy

	rcv.DriveCache = stage1.DriveCache

	rcv.ExpandToAvailable = stage1.ExpandToAvailable

	rcv.IoPolicy = stage1.IoPolicy

	rcv.Name = stage1.Name

	rcv.ObjectType = stage1.ObjectType

	rcv.ReadPolicy = stage1.ReadPolicy

	rcv.Size = stage1.Size

	rcv.WritePolicy = stage1.WritePolicy

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "AccessPolicy")

	delete(stage2, "BootDrive")

	delete(stage2, "DiskGroupName")

	delete(stage2, "DiskGroupPolicy")

	delete(stage2, "DriveCache")

	delete(stage2, "ExpandToAvailable")

	delete(stage2, "IoPolicy")

	delete(stage2, "Name")

	delete(stage2, "ObjectType")

	delete(stage2, "ReadPolicy")

	delete(stage2, "Size")

	delete(stage2, "WritePolicy")

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
		m.StorageVirtualDriveConfigAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m StorageVirtualDriveConfigAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// This property holds the access policy that host has on this virtual drive.
		//
		// Enum: [Default ReadWrite ReadOnly Blocked]
		AccessPolicy *string `json:"AccessPolicy,omitempty"`

		// This flag enables the use of this virtual drive as a boot drive.
		//
		BootDrive *bool `json:"BootDrive,omitempty"`

		// Disk group policy that has the disk group in which this virtual drive needs to be created.
		//
		// Read Only: true
		DiskGroupName string `json:"DiskGroupName,omitempty"`

		// Disk group policy that has the disk group in which this virtual drive needs to be created.
		//
		DiskGroupPolicy string `json:"DiskGroupPolicy,omitempty"`

		// This property expect disk cache policy.
		//
		// Enum: [Default NoChange Enable Disable]
		DriveCache *string `json:"DriveCache,omitempty"`

		// This flag enables this virtual drive to use all the available space in the disk group. When this flag is configured, the size property is ignored.
		//
		ExpandToAvailable *bool `json:"ExpandToAvailable,omitempty"`

		// This property expects the desired IO mode - direct IO or cached IO.
		//
		// Enum: [Default Direct Cached]
		IoPolicy *string `json:"IoPolicy,omitempty"`

		// The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen), _ (underscore), : (colon), and . (period) are not allowed.
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

		// This property holds the read ahead mode to be used.
		//
		// Enum: [Default ReadAhead NoReadAhead]
		ReadPolicy *string `json:"ReadPolicy,omitempty"`

		// Virtual drive size in MB. This is a required field unless the 'Expand to Available' option is enabled.
		//
		Size int64 `json:"Size,omitempty"`

		// This property holds the write mode used to write the data in this virtual drive.
		//
		// Enum: [Default WriteThrough WriteBackGoodBbu AlwaysWriteBack]
		WritePolicy *string `json:"WritePolicy,omitempty"`
	}

	stage1.AccessPolicy = m.AccessPolicy

	stage1.BootDrive = m.BootDrive

	stage1.DiskGroupName = m.DiskGroupName

	stage1.DiskGroupPolicy = m.DiskGroupPolicy

	stage1.DriveCache = m.DriveCache

	stage1.ExpandToAvailable = m.ExpandToAvailable

	stage1.IoPolicy = m.IoPolicy

	stage1.Name = m.Name

	stage1.ObjectType = m.ObjectType

	stage1.ReadPolicy = m.ReadPolicy

	stage1.Size = m.Size

	stage1.WritePolicy = m.WritePolicy

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.StorageVirtualDriveConfigAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.StorageVirtualDriveConfigAO0P0)
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

// Validate validates this storage virtual drive config a o0 p0
func (m *StorageVirtualDriveConfigAO0P0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDriveCache(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIoPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReadPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWritePolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var storageVirtualDriveConfigAO0P0TypeAccessPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Default","ReadWrite","ReadOnly","Blocked"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storageVirtualDriveConfigAO0P0TypeAccessPolicyPropEnum = append(storageVirtualDriveConfigAO0P0TypeAccessPolicyPropEnum, v)
	}
}

const (

	// StorageVirtualDriveConfigAO0P0AccessPolicyDefault captures enum value "Default"
	StorageVirtualDriveConfigAO0P0AccessPolicyDefault string = "Default"

	// StorageVirtualDriveConfigAO0P0AccessPolicyReadWrite captures enum value "ReadWrite"
	StorageVirtualDriveConfigAO0P0AccessPolicyReadWrite string = "ReadWrite"

	// StorageVirtualDriveConfigAO0P0AccessPolicyReadOnly captures enum value "ReadOnly"
	StorageVirtualDriveConfigAO0P0AccessPolicyReadOnly string = "ReadOnly"

	// StorageVirtualDriveConfigAO0P0AccessPolicyBlocked captures enum value "Blocked"
	StorageVirtualDriveConfigAO0P0AccessPolicyBlocked string = "Blocked"
)

// prop value enum
func (m *StorageVirtualDriveConfigAO0P0) validateAccessPolicyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, storageVirtualDriveConfigAO0P0TypeAccessPolicyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StorageVirtualDriveConfigAO0P0) validateAccessPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.AccessPolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccessPolicyEnum("AccessPolicy", "body", *m.AccessPolicy); err != nil {
		return err
	}

	return nil
}

var storageVirtualDriveConfigAO0P0TypeDriveCachePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Default","NoChange","Enable","Disable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storageVirtualDriveConfigAO0P0TypeDriveCachePropEnum = append(storageVirtualDriveConfigAO0P0TypeDriveCachePropEnum, v)
	}
}

const (

	// StorageVirtualDriveConfigAO0P0DriveCacheDefault captures enum value "Default"
	StorageVirtualDriveConfigAO0P0DriveCacheDefault string = "Default"

	// StorageVirtualDriveConfigAO0P0DriveCacheNoChange captures enum value "NoChange"
	StorageVirtualDriveConfigAO0P0DriveCacheNoChange string = "NoChange"

	// StorageVirtualDriveConfigAO0P0DriveCacheEnable captures enum value "Enable"
	StorageVirtualDriveConfigAO0P0DriveCacheEnable string = "Enable"

	// StorageVirtualDriveConfigAO0P0DriveCacheDisable captures enum value "Disable"
	StorageVirtualDriveConfigAO0P0DriveCacheDisable string = "Disable"
)

// prop value enum
func (m *StorageVirtualDriveConfigAO0P0) validateDriveCacheEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, storageVirtualDriveConfigAO0P0TypeDriveCachePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StorageVirtualDriveConfigAO0P0) validateDriveCache(formats strfmt.Registry) error {

	if swag.IsZero(m.DriveCache) { // not required
		return nil
	}

	// value enum
	if err := m.validateDriveCacheEnum("DriveCache", "body", *m.DriveCache); err != nil {
		return err
	}

	return nil
}

var storageVirtualDriveConfigAO0P0TypeIoPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Default","Direct","Cached"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storageVirtualDriveConfigAO0P0TypeIoPolicyPropEnum = append(storageVirtualDriveConfigAO0P0TypeIoPolicyPropEnum, v)
	}
}

const (

	// StorageVirtualDriveConfigAO0P0IoPolicyDefault captures enum value "Default"
	StorageVirtualDriveConfigAO0P0IoPolicyDefault string = "Default"

	// StorageVirtualDriveConfigAO0P0IoPolicyDirect captures enum value "Direct"
	StorageVirtualDriveConfigAO0P0IoPolicyDirect string = "Direct"

	// StorageVirtualDriveConfigAO0P0IoPolicyCached captures enum value "Cached"
	StorageVirtualDriveConfigAO0P0IoPolicyCached string = "Cached"
)

// prop value enum
func (m *StorageVirtualDriveConfigAO0P0) validateIoPolicyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, storageVirtualDriveConfigAO0P0TypeIoPolicyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StorageVirtualDriveConfigAO0P0) validateIoPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.IoPolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateIoPolicyEnum("IoPolicy", "body", *m.IoPolicy); err != nil {
		return err
	}

	return nil
}

var storageVirtualDriveConfigAO0P0TypeReadPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Default","ReadAhead","NoReadAhead"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storageVirtualDriveConfigAO0P0TypeReadPolicyPropEnum = append(storageVirtualDriveConfigAO0P0TypeReadPolicyPropEnum, v)
	}
}

const (

	// StorageVirtualDriveConfigAO0P0ReadPolicyDefault captures enum value "Default"
	StorageVirtualDriveConfigAO0P0ReadPolicyDefault string = "Default"

	// StorageVirtualDriveConfigAO0P0ReadPolicyReadAhead captures enum value "ReadAhead"
	StorageVirtualDriveConfigAO0P0ReadPolicyReadAhead string = "ReadAhead"

	// StorageVirtualDriveConfigAO0P0ReadPolicyNoReadAhead captures enum value "NoReadAhead"
	StorageVirtualDriveConfigAO0P0ReadPolicyNoReadAhead string = "NoReadAhead"
)

// prop value enum
func (m *StorageVirtualDriveConfigAO0P0) validateReadPolicyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, storageVirtualDriveConfigAO0P0TypeReadPolicyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StorageVirtualDriveConfigAO0P0) validateReadPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.ReadPolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateReadPolicyEnum("ReadPolicy", "body", *m.ReadPolicy); err != nil {
		return err
	}

	return nil
}

var storageVirtualDriveConfigAO0P0TypeWritePolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Default","WriteThrough","WriteBackGoodBbu","AlwaysWriteBack"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storageVirtualDriveConfigAO0P0TypeWritePolicyPropEnum = append(storageVirtualDriveConfigAO0P0TypeWritePolicyPropEnum, v)
	}
}

const (

	// StorageVirtualDriveConfigAO0P0WritePolicyDefault captures enum value "Default"
	StorageVirtualDriveConfigAO0P0WritePolicyDefault string = "Default"

	// StorageVirtualDriveConfigAO0P0WritePolicyWriteThrough captures enum value "WriteThrough"
	StorageVirtualDriveConfigAO0P0WritePolicyWriteThrough string = "WriteThrough"

	// StorageVirtualDriveConfigAO0P0WritePolicyWriteBackGoodBbu captures enum value "WriteBackGoodBbu"
	StorageVirtualDriveConfigAO0P0WritePolicyWriteBackGoodBbu string = "WriteBackGoodBbu"

	// StorageVirtualDriveConfigAO0P0WritePolicyAlwaysWriteBack captures enum value "AlwaysWriteBack"
	StorageVirtualDriveConfigAO0P0WritePolicyAlwaysWriteBack string = "AlwaysWriteBack"
)

// prop value enum
func (m *StorageVirtualDriveConfigAO0P0) validateWritePolicyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, storageVirtualDriveConfigAO0P0TypeWritePolicyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StorageVirtualDriveConfigAO0P0) validateWritePolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.WritePolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateWritePolicyEnum("WritePolicy", "body", *m.WritePolicy); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageVirtualDriveConfigAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageVirtualDriveConfigAO0P0) UnmarshalBinary(b []byte) error {
	var res StorageVirtualDriveConfigAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
