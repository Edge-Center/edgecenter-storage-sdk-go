// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EventsMetaResponse EventsMeta describe how to form POST/PATCH request
//
// swagger:model EventsMetaResponse
type EventsMetaResponse struct {

	// example
	Example *EventsResponse `json:"example,omitempty"`

	// Event structure
	Structure []*EventDescription `json:"structure"`
}

// Validate validates this events meta response
func (m *EventsMetaResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExample(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStructure(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventsMetaResponse) validateExample(formats strfmt.Registry) error {
	if swag.IsZero(m.Example) { // not required
		return nil
	}

	if m.Example != nil {
		if err := m.Example.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("example")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("example")
			}
			return err
		}
	}

	return nil
}

func (m *EventsMetaResponse) validateStructure(formats strfmt.Registry) error {
	if swag.IsZero(m.Structure) { // not required
		return nil
	}

	for i := 0; i < len(m.Structure); i++ {
		if swag.IsZero(m.Structure[i]) { // not required
			continue
		}

		if m.Structure[i] != nil {
			if err := m.Structure[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("structure" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("structure" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this events meta response based on the context it is used
func (m *EventsMetaResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExample(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStructure(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventsMetaResponse) contextValidateExample(ctx context.Context, formats strfmt.Registry) error {

	if m.Example != nil {

		if swag.IsZero(m.Example) { // not required
			return nil
		}

		if err := m.Example.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("example")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("example")
			}
			return err
		}
	}

	return nil
}

func (m *EventsMetaResponse) contextValidateStructure(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Structure); i++ {

		if m.Structure[i] != nil {

			if swag.IsZero(m.Structure[i]) { // not required
				return nil
			}

			if err := m.Structure[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("structure" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("structure" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EventsMetaResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventsMetaResponse) UnmarshalBinary(b []byte) error {
	var res EventsMetaResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
