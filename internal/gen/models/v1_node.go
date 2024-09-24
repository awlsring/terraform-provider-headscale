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

// V1Node v1 node
//
// swagger:model v1Node
type V1Node struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// disco key
	DiscoKey string `json:"discoKey,omitempty"`

	// expiry
	// Format: date-time
	Expiry strfmt.DateTime `json:"expiry,omitempty"`

	// forced tags
	ForcedTags []string `json:"forcedTags"`

	// given name
	GivenName string `json:"givenName,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// invalid tags
	InvalidTags []string `json:"invalidTags"`

	// ip addresses
	IPAddresses []string `json:"ipAddresses"`

	// last seen
	// Format: date-time
	LastSeen strfmt.DateTime `json:"lastSeen,omitempty"`

	// machine key
	MachineKey string `json:"machineKey,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// node key
	NodeKey string `json:"nodeKey,omitempty"`

	// online
	Online bool `json:"online,omitempty"`

	// pre auth key
	PreAuthKey *V1PreAuthKey `json:"preAuthKey,omitempty"`

	// register method
	RegisterMethod *V1RegisterMethod `json:"registerMethod,omitempty"`

	// user
	User *V1User `json:"user,omitempty"`

	// valid tags
	ValidTags []string `json:"validTags"`
}

// Validate validates this v1 node
func (m *V1Node) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastSeen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreAuthKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegisterMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Node) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1Node) validateExpiry(formats strfmt.Registry) error {
	if swag.IsZero(m.Expiry) { // not required
		return nil
	}

	if err := validate.FormatOf("expiry", "body", "date-time", m.Expiry.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1Node) validateLastSeen(formats strfmt.Registry) error {
	if swag.IsZero(m.LastSeen) { // not required
		return nil
	}

	if err := validate.FormatOf("lastSeen", "body", "date-time", m.LastSeen.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1Node) validatePreAuthKey(formats strfmt.Registry) error {
	if swag.IsZero(m.PreAuthKey) { // not required
		return nil
	}

	if m.PreAuthKey != nil {
		if err := m.PreAuthKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preAuthKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preAuthKey")
			}
			return err
		}
	}

	return nil
}

func (m *V1Node) validateRegisterMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.RegisterMethod) { // not required
		return nil
	}

	if m.RegisterMethod != nil {
		if err := m.RegisterMethod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registerMethod")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registerMethod")
			}
			return err
		}
	}

	return nil
}

func (m *V1Node) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 node based on the context it is used
func (m *V1Node) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePreAuthKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegisterMethod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Node) contextValidatePreAuthKey(ctx context.Context, formats strfmt.Registry) error {

	if m.PreAuthKey != nil {

		if swag.IsZero(m.PreAuthKey) { // not required
			return nil
		}

		if err := m.PreAuthKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preAuthKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preAuthKey")
			}
			return err
		}
	}

	return nil
}

func (m *V1Node) contextValidateRegisterMethod(ctx context.Context, formats strfmt.Registry) error {

	if m.RegisterMethod != nil {

		if swag.IsZero(m.RegisterMethod) { // not required
			return nil
		}

		if err := m.RegisterMethod.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registerMethod")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registerMethod")
			}
			return err
		}
	}

	return nil
}

func (m *V1Node) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {

		if swag.IsZero(m.User) { // not required
			return nil
		}

		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Node) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Node) UnmarshalBinary(b []byte) error {
	var res V1Node
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}