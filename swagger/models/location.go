// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Location location
//
// swagger:model Location
type Location struct {

	// address
	Address string `json:"address,omitempty"`

	// allow for new storage
	// Enum: [deny undefined allow]
	AllowForNewStorage string `json:"allow_for_new_storage,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// title
	Title map[string]string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this location
func (m *Location) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowForNewStorage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var locationTypeAllowForNewStoragePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["deny","undefined","allow"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		locationTypeAllowForNewStoragePropEnum = append(locationTypeAllowForNewStoragePropEnum, v)
	}
}

const (

	// LocationAllowForNewStorageDeny captures enum value "deny"
	LocationAllowForNewStorageDeny string = "deny"

	// LocationAllowForNewStorageUndefined captures enum value "undefined"
	LocationAllowForNewStorageUndefined string = "undefined"

	// LocationAllowForNewStorageAllow captures enum value "allow"
	LocationAllowForNewStorageAllow string = "allow"
)

// prop value enum
func (m *Location) validateAllowForNewStorageEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, locationTypeAllowForNewStoragePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Location) validateAllowForNewStorage(formats strfmt.Registry) error {
	if swag.IsZero(m.AllowForNewStorage) { // not required
		return nil
	}

	// value enum
	if err := m.validateAllowForNewStorageEnum("allow_for_new_storage", "body", m.AllowForNewStorage); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this location based on context it is used
func (m *Location) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Location) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Location) UnmarshalBinary(b []byte) error {
	var res Location
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
