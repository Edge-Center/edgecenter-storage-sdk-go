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

// ClientStats client stats
//
// swagger:model ClientStats
type ClientStats struct {

	// Maximum number of files within the specified time period
	FileQuantitySumMax uint64 `json:"file_quantity_sum_max,omitempty"`

	// Your account ID
	ID int64 `json:"id,omitempty"`

	// Statistics grouped by storage locations
	Locations map[string]LocationStats `json:"locations,omitempty"`

	// Total number of incoming requests within the specified time period
	RequestsInSum uint64 `json:"requests_in_sum,omitempty"`

	// Total number of requests from the edges within the specified time period
	RequestsOutEdgesSum uint64 `json:"requests_out_edges_sum,omitempty"`

	// Total number of outсoming requests without requests from the edges within the specified time period
	RequestsOutWoEdgesSum uint64 `json:"requests_out_wo_edges_sum,omitempty"`

	// Total number of requests within the specified time period
	RequestsSum uint64 `json:"requests_sum,omitempty"`

	// RequestsWoEdgesSum is sum of requests without edges out requests for grouped period
	RequestsWoEdgesSum uint64 `json:"requests_wo_edges_sum,omitempty"`

	// Maximum amount of all file sizes within the specified time period
	SizeSumMax uint64 `json:"size_sum_max,omitempty"`

	// Mean amount of all file sizes within the specified time period
	SizeSumMean uint64 `json:"size_sum_mean,omitempty"`

	// Total amount of incoming traffic within the specified time period
	TrafficInSum uint64 `json:"traffic_in_sum,omitempty"`

	// Total number of traffic from the edges within the specified time period
	TrafficOutEdgesSum uint64 `json:"traffic_out_edges_sum,omitempty"`

	// Total number of outсoming traffic without requests from the edges within the specified time period
	TrafficOutWoEdgesSum uint64 `json:"traffic_out_wo_edges_sum,omitempty"`

	// Total amount of traffic within the specified time period
	TrafficSum uint64 `json:"traffic_sum,omitempty"`
}

// Validate validates this client stats
func (m *ClientStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientStats) validateLocations(formats strfmt.Registry) error {
	if swag.IsZero(m.Locations) { // not required
		return nil
	}

	for k := range m.Locations {

		if err := validate.Required("locations"+"."+k, "body", m.Locations[k]); err != nil {
			return err
		}
		if val, ok := m.Locations[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("locations" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("locations" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this client stats based on the context it is used
func (m *ClientStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientStats) contextValidateLocations(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Locations {

		if val, ok := m.Locations[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClientStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientStats) UnmarshalBinary(b []byte) error {
	var res ClientStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
