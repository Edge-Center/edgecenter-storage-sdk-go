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

// ClientLocationRes client location res
//
// swagger:model clientLocationRes
type ClientLocationRes struct {

	// Storage address
	Address string `json:"address,omitempty"`

	// Shows if you are allowed to create more storages.
	// Has one of the values: <br><ul>
	// <li><b>allow</b> — you can create more storages;</li>
	// <li><b>deny</b> — you cannot create any more storages</li>
	// </ul>
	// Enum: [deny allow]
	AllowForNewStorage string `json:"allow_for_new_storage,omitempty"`

	// Storage location ID
	ID int64 `json:"id,omitempty"`

	// Storage region name
	// Example: s-dt2
	// Enum: [s-dt2]
	Name string `json:"name,omitempty"`

	// Storage type
	// Enum: [s3]
	Type string `json:"type,omitempty"`
}

// Validate validates this client location res
func (m *ClientLocationRes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowForNewStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var clientLocationResTypeAllowForNewStoragePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["deny","allow"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clientLocationResTypeAllowForNewStoragePropEnum = append(clientLocationResTypeAllowForNewStoragePropEnum, v)
	}
}

const (

	// ClientLocationResAllowForNewStorageDeny captures enum value "deny"
	ClientLocationResAllowForNewStorageDeny string = "deny"

	// ClientLocationResAllowForNewStorageAllow captures enum value "allow"
	ClientLocationResAllowForNewStorageAllow string = "allow"
)

// prop value enum
func (m *ClientLocationRes) validateAllowForNewStorageEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clientLocationResTypeAllowForNewStoragePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClientLocationRes) validateAllowForNewStorage(formats strfmt.Registry) error {
	if swag.IsZero(m.AllowForNewStorage) { // not required
		return nil
	}

	// value enum
	if err := m.validateAllowForNewStorageEnum("allow_for_new_storage", "body", m.AllowForNewStorage); err != nil {
		return err
	}

	return nil
}

var clientLocationResTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["s-dt2"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clientLocationResTypeNamePropEnum = append(clientLocationResTypeNamePropEnum, v)
	}
}

const (

	// ClientLocationResNameSDashDt2 captures enum value "s-dt2"
	ClientLocationResNameSDashDt2 string = "s-dt2"
)

// prop value enum
func (m *ClientLocationRes) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clientLocationResTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClientLocationRes) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var clientLocationResTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["s3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clientLocationResTypeTypePropEnum = append(clientLocationResTypeTypePropEnum, v)
	}
}

const (

	// ClientLocationResTypeS3 captures enum value "s3"
	ClientLocationResTypeS3 string = "s3"
)

// prop value enum
func (m *ClientLocationRes) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clientLocationResTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClientLocationRes) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this client location res based on context it is used
func (m *ClientLocationRes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClientLocationRes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientLocationRes) UnmarshalBinary(b []byte) error {
	var res ClientLocationRes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
