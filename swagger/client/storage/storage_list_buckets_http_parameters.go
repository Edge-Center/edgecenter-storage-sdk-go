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

// NewStorageListBucketsHTTPParams creates a new StorageListBucketsHTTPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStorageListBucketsHTTPParams() *StorageListBucketsHTTPParams {
	return &StorageListBucketsHTTPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStorageListBucketsHTTPParamsWithTimeout creates a new StorageListBucketsHTTPParams object
// with the ability to set a timeout on a request.
func NewStorageListBucketsHTTPParamsWithTimeout(timeout time.Duration) *StorageListBucketsHTTPParams {
	return &StorageListBucketsHTTPParams{
		timeout: timeout,
	}
}

// NewStorageListBucketsHTTPParamsWithContext creates a new StorageListBucketsHTTPParams object
// with the ability to set a context for a request.
func NewStorageListBucketsHTTPParamsWithContext(ctx context.Context) *StorageListBucketsHTTPParams {
	return &StorageListBucketsHTTPParams{
		Context: ctx,
	}
}

// NewStorageListBucketsHTTPParamsWithHTTPClient creates a new StorageListBucketsHTTPParams object
// with the ability to set a custom HTTPClient for a request.
func NewStorageListBucketsHTTPParamsWithHTTPClient(client *http.Client) *StorageListBucketsHTTPParams {
	return &StorageListBucketsHTTPParams{
		HTTPClient: client,
	}
}

/*
StorageListBucketsHTTPParams contains all the parameters to send to the API endpoint

	for the storage list buckets Http operation.

	Typically these are written to a http.Request.
*/
type StorageListBucketsHTTPParams struct {

	// ID.
	//
	// Format: int64
	ID int64

	/* Limit.

	   Max number of records in response

	   Format: uint64
	*/
	Limit *uint64

	/* Offset.

	   Amount of records to skip before beginning to write in response.

	   Format: uint64
	*/
	Offset *uint64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the storage list buckets Http params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageListBucketsHTTPParams) WithDefaults() *StorageListBucketsHTTPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the storage list buckets Http params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageListBucketsHTTPParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the storage list buckets Http params
func (o *StorageListBucketsHTTPParams) WithTimeout(timeout time.Duration) *StorageListBucketsHTTPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage list buckets Http params
func (o *StorageListBucketsHTTPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage list buckets Http params
func (o *StorageListBucketsHTTPParams) WithContext(ctx context.Context) *StorageListBucketsHTTPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage list buckets Http params
func (o *StorageListBucketsHTTPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage list buckets Http params
func (o *StorageListBucketsHTTPParams) WithHTTPClient(client *http.Client) *StorageListBucketsHTTPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage list buckets Http params
func (o *StorageListBucketsHTTPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the storage list buckets Http params
func (o *StorageListBucketsHTTPParams) WithID(id int64) *StorageListBucketsHTTPParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the storage list buckets Http params
func (o *StorageListBucketsHTTPParams) SetID(id int64) {
	o.ID = id
}

// WithLimit adds the limit to the storage list buckets Http params
func (o *StorageListBucketsHTTPParams) WithLimit(limit *uint64) *StorageListBucketsHTTPParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the storage list buckets Http params
func (o *StorageListBucketsHTTPParams) SetLimit(limit *uint64) {
	o.Limit = limit
}

// WithOffset adds the offset to the storage list buckets Http params
func (o *StorageListBucketsHTTPParams) WithOffset(offset *uint64) *StorageListBucketsHTTPParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the storage list buckets Http params
func (o *StorageListBucketsHTTPParams) SetOffset(offset *uint64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *StorageListBucketsHTTPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit uint64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatUint64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset uint64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatUint64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
