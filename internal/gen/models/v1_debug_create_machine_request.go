// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1DebugCreateMachineRequest v1 debug create machine request
//
// swagger:model v1DebugCreateMachineRequest
type V1DebugCreateMachineRequest struct {

	// key
	Key string `json:"key,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// routes
	Routes []string `json:"routes"`

	// user
	User string `json:"user,omitempty"`
}

// Validate validates this v1 debug create machine request
func (m *V1DebugCreateMachineRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 debug create machine request based on context it is used
func (m *V1DebugCreateMachineRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1DebugCreateMachineRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1DebugCreateMachineRequest) UnmarshalBinary(b []byte) error {
	var res V1DebugCreateMachineRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
