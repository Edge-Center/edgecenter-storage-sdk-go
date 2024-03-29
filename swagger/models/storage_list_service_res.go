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

// StorageListServiceRes storage list service res
//
// swagger:model StorageListServiceRes
type StorageListServiceRes struct {

	// num storages
	NumStorages int64 `json:"num_storages,omitempty"`

	// storages
	Storages []*Storage `json:"storages"`
}

// Validate validates this storage list service res
func (m *StorageListServiceRes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStorages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageListServiceRes) validateStorages(formats strfmt.Registry) error {
	if swag.IsZero(m.Storages) { // not required
		return nil
	}

	for i := 0; i < len(m.Storages); i++ {
		if swag.IsZero(m.Storages[i]) { // not required
			continue
		}

		if m.Storages[i] != nil {
			if err := m.Storages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("storages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("storages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this storage list service res based on the context it is used
func (m *StorageListServiceRes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStorages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageListServiceRes) contextValidateStorages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Storages); i++ {

		if m.Storages[i] != nil {

			if swag.IsZero(m.Storages[i]) { // not required
				return nil
			}

			if err := m.Storages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("storages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("storages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageListServiceRes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageListServiceRes) UnmarshalBinary(b []byte) error {
	var res StorageListServiceRes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
