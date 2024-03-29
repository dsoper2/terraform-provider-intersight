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

// WorkflowWorkflowDefinition Workflow:Workflow Definition
//
// Workflow definition is a collection of tasks that are sequenced in a certain way using control tasks. The tasks in the workflow definition is represented as a directed acyclic graph where each node in the graph is a task and the edges in the graph are transitions from one task to another.
//
// swagger:model workflowWorkflowDefinition
type WorkflowWorkflowDefinition struct {
	MoBaseMo

	// The catalog under which the definition is present.
	//
	Catalog *WorkflowCatalogRef `json:"Catalog,omitempty"`

	// When true this will be the workflow version that is used when a specific workflow definition version is not specified. The default version is used when user executes a workflow without specifying a version or when workflow is included in another workflow without a specific version. The very first workflow definition created with a name will be set as the default version, after that user can explicitly set any version of the workflow definition as the default version.
	//
	DefaultVersion *bool `json:"DefaultVersion,omitempty"`

	// The description for this workflow.
	//
	Description string `json:"Description,omitempty"`

	// The schema expected for input parameters for this workflow.
	//
	InputDefinition []*WorkflowBaseDataType `json:"InputDefinition"`

	// A user friendly short name to identify the workflow.
	//
	Label string `json:"Label,omitempty"`

	// The name for this workflow. You can have multiple version of the workflow with the same name.
	//
	Name string `json:"Name,omitempty"`

	// The schema expected for output parameters for this workflow.
	//
	OutputDefinition []*WorkflowBaseDataType `json:"OutputDefinition"`

	// The output mappings for the workflow. The outputs for worflows will generally be task output variables that we want to export out at the end of the workflow. The format to specify the mapping is '${Source.output.JsonPath}'. 'Source' is the name of the task within the workflow. You can map any task output to a workflow output as long as the types are compatible. Following this is JSON path expression to extract JSON fragment from source's output.
	//
	OutputParameters interface{} `json:"OutputParameters,omitempty"`

	// The tasks contained inside of the workflow.
	//
	Tasks []*WorkflowWorkflowTask `json:"Tasks"`

	// This will hold the data needed for workflow to be rendered in the user interface.
	//
	UIRenderingData interface{} `json:"UiRenderingData,omitempty"`

	// The current validation state and associated information for this workflow.
	//
	// Read Only: true
	ValidationInformation *WorkflowValidationInformation `json:"ValidationInformation,omitempty"`

	// The version of the workflow to support multiple versions.
	//
	Version int64 `json:"Version,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowWorkflowDefinition) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Catalog *WorkflowCatalogRef `json:"Catalog,omitempty"`

		DefaultVersion *bool `json:"DefaultVersion,omitempty"`

		Description string `json:"Description,omitempty"`

		InputDefinition []*WorkflowBaseDataType `json:"InputDefinition"`

		Label string `json:"Label,omitempty"`

		Name string `json:"Name,omitempty"`

		OutputDefinition []*WorkflowBaseDataType `json:"OutputDefinition"`

		OutputParameters interface{} `json:"OutputParameters,omitempty"`

		Tasks []*WorkflowWorkflowTask `json:"Tasks"`

		UIRenderingData interface{} `json:"UiRenderingData,omitempty"`

		ValidationInformation *WorkflowValidationInformation `json:"ValidationInformation,omitempty"`

		Version int64 `json:"Version,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Catalog = dataAO1.Catalog

	m.DefaultVersion = dataAO1.DefaultVersion

	m.Description = dataAO1.Description

	m.InputDefinition = dataAO1.InputDefinition

	m.Label = dataAO1.Label

	m.Name = dataAO1.Name

	m.OutputDefinition = dataAO1.OutputDefinition

	m.OutputParameters = dataAO1.OutputParameters

	m.Tasks = dataAO1.Tasks

	m.UIRenderingData = dataAO1.UIRenderingData

	m.ValidationInformation = dataAO1.ValidationInformation

	m.Version = dataAO1.Version

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowWorkflowDefinition) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Catalog *WorkflowCatalogRef `json:"Catalog,omitempty"`

		DefaultVersion *bool `json:"DefaultVersion,omitempty"`

		Description string `json:"Description,omitempty"`

		InputDefinition []*WorkflowBaseDataType `json:"InputDefinition"`

		Label string `json:"Label,omitempty"`

		Name string `json:"Name,omitempty"`

		OutputDefinition []*WorkflowBaseDataType `json:"OutputDefinition"`

		OutputParameters interface{} `json:"OutputParameters,omitempty"`

		Tasks []*WorkflowWorkflowTask `json:"Tasks"`

		UIRenderingData interface{} `json:"UiRenderingData,omitempty"`

		ValidationInformation *WorkflowValidationInformation `json:"ValidationInformation,omitempty"`

		Version int64 `json:"Version,omitempty"`
	}

	dataAO1.Catalog = m.Catalog

	dataAO1.DefaultVersion = m.DefaultVersion

	dataAO1.Description = m.Description

	dataAO1.InputDefinition = m.InputDefinition

	dataAO1.Label = m.Label

	dataAO1.Name = m.Name

	dataAO1.OutputDefinition = m.OutputDefinition

	dataAO1.OutputParameters = m.OutputParameters

	dataAO1.Tasks = m.Tasks

	dataAO1.UIRenderingData = m.UIRenderingData

	dataAO1.ValidationInformation = m.ValidationInformation

	dataAO1.Version = m.Version

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow workflow definition
func (m *WorkflowWorkflowDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCatalog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInputDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutputDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTasks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationInformation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowWorkflowDefinition) validateCatalog(formats strfmt.Registry) error {

	if swag.IsZero(m.Catalog) { // not required
		return nil
	}

	if m.Catalog != nil {
		if err := m.Catalog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Catalog")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowWorkflowDefinition) validateInputDefinition(formats strfmt.Registry) error {

	if swag.IsZero(m.InputDefinition) { // not required
		return nil
	}

	for i := 0; i < len(m.InputDefinition); i++ {
		if swag.IsZero(m.InputDefinition[i]) { // not required
			continue
		}

		if m.InputDefinition[i] != nil {
			if err := m.InputDefinition[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("InputDefinition" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkflowWorkflowDefinition) validateOutputDefinition(formats strfmt.Registry) error {

	if swag.IsZero(m.OutputDefinition) { // not required
		return nil
	}

	for i := 0; i < len(m.OutputDefinition); i++ {
		if swag.IsZero(m.OutputDefinition[i]) { // not required
			continue
		}

		if m.OutputDefinition[i] != nil {
			if err := m.OutputDefinition[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OutputDefinition" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkflowWorkflowDefinition) validateTasks(formats strfmt.Registry) error {

	if swag.IsZero(m.Tasks) { // not required
		return nil
	}

	for i := 0; i < len(m.Tasks); i++ {
		if swag.IsZero(m.Tasks[i]) { // not required
			continue
		}

		if m.Tasks[i] != nil {
			if err := m.Tasks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Tasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkflowWorkflowDefinition) validateValidationInformation(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidationInformation) { // not required
		return nil
	}

	if m.ValidationInformation != nil {
		if err := m.ValidationInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ValidationInformation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowWorkflowDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowWorkflowDefinition) UnmarshalBinary(b []byte) error {
	var res WorkflowWorkflowDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
