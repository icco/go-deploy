// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PatchRequest1 patch request 1
// swagger:model patch_request_1
type PatchRequest1 struct {

	// current configuration
	// Format: uri
	CurrentConfiguration strfmt.URI `json:"current_configuration,omitempty"`

	// current image
	// Format: uri
	CurrentImage strfmt.URI `json:"current_image,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this patch request 1
func (m *PatchRequest1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentImage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchRequest1) validateCurrentConfiguration(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentConfiguration) { // not required
		return nil
	}

	if err := validate.FormatOf("current_configuration", "body", "uri", m.CurrentConfiguration.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PatchRequest1) validateCurrentImage(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentImage) { // not required
		return nil
	}

	if err := validate.FormatOf("current_image", "body", "uri", m.CurrentImage.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchRequest1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchRequest1) UnmarshalBinary(b []byte) error {
	var res PatchRequest1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}