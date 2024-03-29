// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// HyperflexClusterStoragePolicy Storage Configuration
//
// A policy specifying HyperFlex cluster storage settings (optional).
//
// swagger:model hyperflexClusterStoragePolicy
type HyperflexClusterStoragePolicy struct {
	PolicyAbstractPolicy

	// List of cluster profiles using this policy.
	//
	ClusterProfiles []*HyperflexClusterProfileRef `json:"ClusterProfiles"`

	// If enabled, formats existing disk partitions (destroys all user data).
	//
	DiskPartitionCleanup *bool `json:"DiskPartitionCleanup,omitempty"`

	// Enable or disable Logical Availability Zones (LAZ). If enabled, HyperFlex Data Platform automatically selects and groups nodes into different availability zones. For HyperFlex Data Platform versions prior to 3.0 release, this setting does not apply. For HyperFlex Data Platform versions 3.0 or higher, this setting is only applicable to Fabric Interconnect attached HyperFlex systems with 8 or more converged nodes.
	//
	LogicalAvalabilityZoneConfig *HyperflexLogicalAvailabilityZone `json:"LogicalAvalabilityZoneConfig,omitempty"`

	// Relationship to the Organization that owns the Managed Object.
	//
	Organization *IamAccountRef `json:"Organization,omitempty"`

	// Enable or disable VDI optimization (hybrid HyperFlex systems only).
	//
	VdiOptimization *bool `json:"VdiOptimization,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexClusterStoragePolicy) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyAbstractPolicy
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyAbstractPolicy = aO0

	// AO1
	var dataAO1 struct {
		ClusterProfiles []*HyperflexClusterProfileRef `json:"ClusterProfiles"`

		DiskPartitionCleanup *bool `json:"DiskPartitionCleanup,omitempty"`

		LogicalAvalabilityZoneConfig *HyperflexLogicalAvailabilityZone `json:"LogicalAvalabilityZoneConfig,omitempty"`

		Organization *IamAccountRef `json:"Organization,omitempty"`

		VdiOptimization *bool `json:"VdiOptimization,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ClusterProfiles = dataAO1.ClusterProfiles

	m.DiskPartitionCleanup = dataAO1.DiskPartitionCleanup

	m.LogicalAvalabilityZoneConfig = dataAO1.LogicalAvalabilityZoneConfig

	m.Organization = dataAO1.Organization

	m.VdiOptimization = dataAO1.VdiOptimization

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexClusterStoragePolicy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyAbstractPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ClusterProfiles []*HyperflexClusterProfileRef `json:"ClusterProfiles"`

		DiskPartitionCleanup *bool `json:"DiskPartitionCleanup,omitempty"`

		LogicalAvalabilityZoneConfig *HyperflexLogicalAvailabilityZone `json:"LogicalAvalabilityZoneConfig,omitempty"`

		Organization *IamAccountRef `json:"Organization,omitempty"`

		VdiOptimization *bool `json:"VdiOptimization,omitempty"`
	}

	dataAO1.ClusterProfiles = m.ClusterProfiles

	dataAO1.DiskPartitionCleanup = m.DiskPartitionCleanup

	dataAO1.LogicalAvalabilityZoneConfig = m.LogicalAvalabilityZoneConfig

	dataAO1.Organization = m.Organization

	dataAO1.VdiOptimization = m.VdiOptimization

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex cluster storage policy
func (m *HyperflexClusterStoragePolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyAbstractPolicy
	if err := m.PolicyAbstractPolicy.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterProfiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogicalAvalabilityZoneConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HyperflexClusterStoragePolicy) validateClusterProfiles(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterProfiles) { // not required
		return nil
	}

	for i := 0; i < len(m.ClusterProfiles); i++ {
		if swag.IsZero(m.ClusterProfiles[i]) { // not required
			continue
		}

		if m.ClusterProfiles[i] != nil {
			if err := m.ClusterProfiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ClusterProfiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HyperflexClusterStoragePolicy) validateLogicalAvalabilityZoneConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.LogicalAvalabilityZoneConfig) { // not required
		return nil
	}

	if m.LogicalAvalabilityZoneConfig != nil {
		if err := m.LogicalAvalabilityZoneConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LogicalAvalabilityZoneConfig")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexClusterStoragePolicy) validateOrganization(formats strfmt.Registry) error {

	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Organization")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexClusterStoragePolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexClusterStoragePolicy) UnmarshalBinary(b []byte) error {
	var res HyperflexClusterStoragePolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
