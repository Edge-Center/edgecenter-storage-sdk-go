// Code generated by go-swagger; DO NOT EDIT.

package buckets

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

// NewStorageBucketCORSCreateHTTPParams creates a new StorageBucketCORSCreateHTTPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStorageBucketCORSCreateHTTPParams() *StorageBucketCORSCreateHTTPParams {
	return &StorageBucketCORSCreateHTTPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStorageBucketCORSCreateHTTPParamsWithTimeout creates a new StorageBucketCORSCreateHTTPParams object
// with the ability to set a timeout on a request.
func NewStorageBucketCORSCreateHTTPParamsWithTimeout(timeout time.Duration) *StorageBucketCORSCreateHTTPParams {
	return &StorageBucketCORSCreateHTTPParams{
		timeout: timeout,
	}
}

// NewStorageBucketCORSCreateHTTPParamsWithContext creates a new StorageBucketCORSCreateHTTPParams object
// with the ability to set a context for a request.
func NewStorageBucketCORSCreateHTTPParamsWithContext(ctx context.Context) *StorageBucketCORSCreateHTTPParams {
	return &StorageBucketCORSCreateHTTPParams{
		Context: ctx,
	}
}

// NewStorageBucketCORSCreateHTTPParamsWithHTTPClient creates a new StorageBucketCORSCreateHTTPParams object
// with the ability to set a custom HTTPClient for a request.
func NewStorageBucketCORSCreateHTTPParamsWithHTTPClient(client *http.Client) *StorageBucketCORSCreateHTTPParams {
	return &StorageBucketCORSCreateHTTPParams{
		HTTPClient: client,
	}
}

/*
StorageBucketCORSCreateHTTPParams contains all the parameters to send to the API endpoint

	for the storage bucket c o r s create Http operation.

	Typically these are written to a http.Request.
*/
type StorageBucketCORSCreateHTTPParams struct {

	// Body.
	Body StorageBucketCORSCreateHTTPBody

	/* ID.

	   Storage ID

	   Format: int64
	*/
	ID int64

	/* Name.

	   Bucket name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the storage bucket c o r s create Http params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageBucketCORSCreateHTTPParams) WithDefaults() *StorageBucketCORSCreateHTTPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the storage bucket c o r s create Http params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageBucketCORSCreateHTTPParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the storage bucket c o r s create Http params
func (o *StorageBucketCORSCreateHTTPParams) WithTimeout(timeout time.Duration) *StorageBucketCORSCreateHTTPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage bucket c o r s create Http params
func (o *StorageBucketCORSCreateHTTPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage bucket c o r s create Http params
func (o *StorageBucketCORSCreateHTTPParams) WithContext(ctx context.Context) *StorageBucketCORSCreateHTTPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage bucket c o r s create Http params
func (o *StorageBucketCORSCreateHTTPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage bucket c o r s create Http params
func (o *StorageBucketCORSCreateHTTPParams) WithHTTPClient(client *http.Client) *StorageBucketCORSCreateHTTPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage bucket c o r s create Http params
func (o *StorageBucketCORSCreateHTTPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the storage bucket c o r s create Http params
func (o *StorageBucketCORSCreateHTTPParams) WithBody(body StorageBucketCORSCreateHTTPBody) *StorageBucketCORSCreateHTTPParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the storage bucket c o r s create Http params
func (o *StorageBucketCORSCreateHTTPParams) SetBody(body StorageBucketCORSCreateHTTPBody) {
	o.Body = body
}

// WithID adds the id to the storage bucket c o r s create Http params
func (o *StorageBucketCORSCreateHTTPParams) WithID(id int64) *StorageBucketCORSCreateHTTPParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the storage bucket c o r s create Http params
func (o *StorageBucketCORSCreateHTTPParams) SetID(id int64) {
	o.ID = id
}

// WithName adds the name to the storage bucket c o r s create Http params
func (o *StorageBucketCORSCreateHTTPParams) WithName(name string) *StorageBucketCORSCreateHTTPParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the storage bucket c o r s create Http params
func (o *StorageBucketCORSCreateHTTPParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *StorageBucketCORSCreateHTTPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
