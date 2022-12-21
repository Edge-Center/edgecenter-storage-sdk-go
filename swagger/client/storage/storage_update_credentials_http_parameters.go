// Code generated by go-swagger; DO NOT EDIT.

package storage

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
	"github.com/go-openapi/swag"
)

// NewStorageUpdateCredentialsHTTPParams creates a new StorageUpdateCredentialsHTTPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStorageUpdateCredentialsHTTPParams() *StorageUpdateCredentialsHTTPParams {
	return &StorageUpdateCredentialsHTTPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStorageUpdateCredentialsHTTPParamsWithTimeout creates a new StorageUpdateCredentialsHTTPParams object
// with the ability to set a timeout on a request.
func NewStorageUpdateCredentialsHTTPParamsWithTimeout(timeout time.Duration) *StorageUpdateCredentialsHTTPParams {
	return &StorageUpdateCredentialsHTTPParams{
		timeout: timeout,
	}
}

// NewStorageUpdateCredentialsHTTPParamsWithContext creates a new StorageUpdateCredentialsHTTPParams object
// with the ability to set a context for a request.
func NewStorageUpdateCredentialsHTTPParamsWithContext(ctx context.Context) *StorageUpdateCredentialsHTTPParams {
	return &StorageUpdateCredentialsHTTPParams{
		Context: ctx,
	}
}

// NewStorageUpdateCredentialsHTTPParamsWithHTTPClient creates a new StorageUpdateCredentialsHTTPParams object
// with the ability to set a custom HTTPClient for a request.
func NewStorageUpdateCredentialsHTTPParamsWithHTTPClient(client *http.Client) *StorageUpdateCredentialsHTTPParams {
	return &StorageUpdateCredentialsHTTPParams{
		HTTPClient: client,
	}
}

/*
StorageUpdateCredentialsHTTPParams contains all the parameters to send to the API endpoint

	for the storage update credentials Http operation.

	Typically these are written to a http.Request.
*/
type StorageUpdateCredentialsHTTPParams struct {

	// Body.
	Body StorageUpdateCredentialsHTTPBody

	// ID.
	//
	// Format: int64
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the storage update credentials Http params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageUpdateCredentialsHTTPParams) WithDefaults() *StorageUpdateCredentialsHTTPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the storage update credentials Http params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageUpdateCredentialsHTTPParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the storage update credentials Http params
func (o *StorageUpdateCredentialsHTTPParams) WithTimeout(timeout time.Duration) *StorageUpdateCredentialsHTTPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage update credentials Http params
func (o *StorageUpdateCredentialsHTTPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage update credentials Http params
func (o *StorageUpdateCredentialsHTTPParams) WithContext(ctx context.Context) *StorageUpdateCredentialsHTTPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage update credentials Http params
func (o *StorageUpdateCredentialsHTTPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage update credentials Http params
func (o *StorageUpdateCredentialsHTTPParams) WithHTTPClient(client *http.Client) *StorageUpdateCredentialsHTTPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage update credentials Http params
func (o *StorageUpdateCredentialsHTTPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the storage update credentials Http params
func (o *StorageUpdateCredentialsHTTPParams) WithBody(body StorageUpdateCredentialsHTTPBody) *StorageUpdateCredentialsHTTPParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the storage update credentials Http params
func (o *StorageUpdateCredentialsHTTPParams) SetBody(body StorageUpdateCredentialsHTTPBody) {
	o.Body = body
}

// WithID adds the id to the storage update credentials Http params
func (o *StorageUpdateCredentialsHTTPParams) WithID(id int64) *StorageUpdateCredentialsHTTPParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the storage update credentials Http params
func (o *StorageUpdateCredentialsHTTPParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *StorageUpdateCredentialsHTTPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
