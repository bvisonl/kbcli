// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CatalogValidationError catalog validation error
//
// swagger:model CatalogValidationError
type CatalogValidationError struct {

	// error description
	ErrorDescription string `json:"errorDescription,omitempty"`
}

// Validate validates this catalog validation error
func (m *CatalogValidationError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this catalog validation error based on context it is used
func (m *CatalogValidationError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CatalogValidationError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogValidationError) UnmarshalBinary(b []byte) error {
	var res CatalogValidationError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
