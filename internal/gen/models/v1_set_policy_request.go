// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SetPolicyRequest v1 set policy request
//
// swagger:model v1SetPolicyRequest
type V1SetPolicyRequest struct {

	// policy
	Policy string `json:"policy,omitempty"`
}

// Validate validates this v1 set policy request
func (m *V1SetPolicyRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 set policy request based on context it is used
func (m *V1SetPolicyRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SetPolicyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SetPolicyRequest) UnmarshalBinary(b []byte) error {
	var res V1SetPolicyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
