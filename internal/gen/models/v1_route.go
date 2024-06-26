// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1Route v1 route
//
// swagger:model v1Route
type V1Route struct {

	// advertised
	Advertised bool `json:"advertised,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// deleted at
	// Format: date-time
	DeletedAt strfmt.DateTime `json:"deletedAt,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// is primary
	IsPrimary bool `json:"isPrimary,omitempty"`

	// machine
	Machine *V1Machine `json:"machine,omitempty"`

	// prefix
	Prefix string `json:"prefix,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this v1 route
func (m *V1Route) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Route) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1Route) validateDeletedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.DeletedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("deletedAt", "body", "date-time", m.DeletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1Route) validateMachine(formats strfmt.Registry) error {
	if swag.IsZero(m.Machine) { // not required
		return nil
	}

	if m.Machine != nil {
		if err := m.Machine.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machine")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("machine")
			}
			return err
		}
	}

	return nil
}

func (m *V1Route) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 route based on the context it is used
func (m *V1Route) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMachine(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Route) contextValidateMachine(ctx context.Context, formats strfmt.Registry) error {

	if m.Machine != nil {

		if swag.IsZero(m.Machine) { // not required
			return nil
		}

		if err := m.Machine.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machine")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("machine")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Route) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Route) UnmarshalBinary(b []byte) error {
	var res V1Route
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
