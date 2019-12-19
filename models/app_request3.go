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

// AppRequest3 app request 3
// swagger:model app_request_3
type AppRequest3 struct {

	// handle
	// Required: true
	Handle *string `json:"handle"`
}

// Validate validates this app request 3
func (m *AppRequest3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHandle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppRequest3) validateHandle(formats strfmt.Registry) error {

	if err := validate.Required("handle", "body", m.Handle); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppRequest3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppRequest3) UnmarshalBinary(b []byte) error {
	var res AppRequest3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}