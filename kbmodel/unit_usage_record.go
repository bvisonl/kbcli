// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UnitUsageRecord unit usage record
//
// swagger:model UnitUsageRecord
type UnitUsageRecord struct {

	// unit type
	UnitType string `json:"unitType,omitempty"`

	// usage records
	UsageRecords []*UsageRecord `json:"usageRecords"`
}

// Validate validates this unit usage record
func (m *UnitUsageRecord) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUsageRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UnitUsageRecord) validateUsageRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.UsageRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.UsageRecords); i++ {
		if swag.IsZero(m.UsageRecords[i]) { // not required
			continue
		}

		if m.UsageRecords[i] != nil {
			if err := m.UsageRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usageRecords" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usageRecords" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this unit usage record based on the context it is used
func (m *UnitUsageRecord) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUsageRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UnitUsageRecord) contextValidateUsageRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UsageRecords); i++ {

		if m.UsageRecords[i] != nil {
			if err := m.UsageRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usageRecords" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usageRecords" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UnitUsageRecord) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UnitUsageRecord) UnmarshalBinary(b []byte) error {
	var res UnitUsageRecord
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
