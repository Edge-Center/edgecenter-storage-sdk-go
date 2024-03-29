// Code generated by go-swagger; DO NOT EDIT.

package locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewLocationListHTTPParams creates a new LocationListHTTPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocationListHTTPParams() *LocationListHTTPParams {
	return &LocationListHTTPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocationListHTTPParamsWithTimeout creates a new LocationListHTTPParams object
// with the ability to set a timeout on a request.
func NewLocationListHTTPParamsWithTimeout(timeout time.Duration) *LocationListHTTPParams {
	return &LocationListHTTPParams{
		timeout: timeout,
	}
}

// NewLocationListHTTPParamsWithContext creates a new LocationListHTTPParams object
// with the ability to set a context for a request.
func NewLocationListHTTPParamsWithContext(ctx context.Context) *LocationListHTTPParams {
	return &LocationListHTTPParams{
		Context: ctx,
	}
}

// NewLocationListHTTPParamsWithHTTPClient creates a new LocationListHTTPParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocationListHTTPParamsWithHTTPClient(client *http.Client) *LocationListHTTPParams {
	return &LocationListHTTPParams{
		HTTPClient: client,
	}
}

/*
LocationListHTTPParams contains all the parameters to send to the API endpoint

	for the location list Http operation.

	Typically these are written to a http.Request.
*/
type LocationListHTTPParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the location list Http params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationListHTTPParams) WithDefaults() *LocationListHTTPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the location list Http params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationListHTTPParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the location list Http params
func (o *LocationListHTTPParams) WithTimeout(timeout time.Duration) *LocationListHTTPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the location list Http params
func (o *LocationListHTTPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the location list Http params
func (o *LocationListHTTPParams) WithContext(ctx context.Context) *LocationListHTTPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the location list Http params
func (o *LocationListHTTPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the location list Http params
func (o *LocationListHTTPParams) WithHTTPClient(client *http.Client) *LocationListHTTPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the location list Http params
func (o *LocationListHTTPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *LocationListHTTPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
