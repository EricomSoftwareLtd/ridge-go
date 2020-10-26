// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NodePoolUpdate node pool update
//
// swagger:model node_pool_update
type NodePoolUpdate struct {

	// The desired number of worker nodes in the node pool
	DesiredNodeCount *int32 `json:"desired_node_count,omitempty"`

	// The display name of the node pool
	// Max Length: 512
	DisplayName *string `json:"display_name,omitempty"`
}

// Validate validates this node pool update
func (m *NodePoolUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodePoolUpdate) validateDisplayName(formats strfmt.Registry) error {

	if swag.IsZero(m.DisplayName) { // not required
		return nil
	}

	if err := validate.MaxLength("display_name", "body", string(*m.DisplayName), 512); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodePoolUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodePoolUpdate) UnmarshalBinary(b []byte) error {
	var res NodePoolUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
