// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Session session
//
// swagger:model Session
type Session struct {

	// host
	Host string `json:"host,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// last access date
	// Format: date-time
	LastAccessDate strfmt.DateTime `json:"lastAccessDate,omitempty"`

	// start date
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// timeout
	Timeout int64 `json:"timeout,omitempty"`
}

// Validate validates this session
func (m *Session) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastAccessDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Session) validateLastAccessDate(formats strfmt.Registry) error {
	if swag.IsZero(m.LastAccessDate) { // not required
		return nil
	}

	if err := validate.FormatOf("lastAccessDate", "body", "date-time", m.LastAccessDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Session) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this session based on context it is used
func (m *Session) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Session) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Session) UnmarshalBinary(b []byte) error {
	var res Session
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
