// Code generated by go-swagger; DO NOT EDIT.

package storages

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

// NewStorageListHTTPV2Params creates a new StorageListHTTPV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStorageListHTTPV2Params() *StorageListHTTPV2Params {
	return &StorageListHTTPV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewStorageListHTTPV2ParamsWithTimeout creates a new StorageListHTTPV2Params object
// with the ability to set a timeout on a request.
func NewStorageListHTTPV2ParamsWithTimeout(timeout time.Duration) *StorageListHTTPV2Params {
	return &StorageListHTTPV2Params{
		timeout: timeout,
	}
}

// NewStorageListHTTPV2ParamsWithContext creates a new StorageListHTTPV2Params object
// with the ability to set a context for a request.
func NewStorageListHTTPV2ParamsWithContext(ctx context.Context) *StorageListHTTPV2Params {
	return &StorageListHTTPV2Params{
		Context: ctx,
	}
}

// NewStorageListHTTPV2ParamsWithHTTPClient creates a new StorageListHTTPV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewStorageListHTTPV2ParamsWithHTTPClient(client *http.Client) *StorageListHTTPV2Params {
	return &StorageListHTTPV2Params{
		HTTPClient: client,
	}
}

/*
StorageListHTTPV2Params contains all the parameters to send to the API endpoint

	for the storage list Http v2 operation.

	Typically these are written to a http.Request.
*/
type StorageListHTTPV2Params struct {

	/* ID.

	     Specify a storage ID that should be displayed in the output.<br>
	Note: <ul>
	<li>If you type any number, all storage IDs that contain that number will be displayed</li>
	</ul>
	*/
	ID *string

	/* Limit.

	   Max number of records in response

	   Format: uint64
	*/
	Limit *uint64

	/* Location.

	   Specify a storage location name that should be displayed in the output
	*/
	Location *string

	/* Name.

	     Specify a storage name that should be displayed in the output.<br>
	Note:<ul>
	<li>If you type any letter, all storage names that contain that letter will be displayed</li>
	</ul>
	*/
	Name *string

	/* Offset.

	   Amount of records to skip before beginning to write in response

	   Format: uint64
	*/
	Offset *uint64

	/* OrderBy.

	   Field name to sort by
	*/
	OrderBy *string

	/* OrderDirection.

	     Ascending or descending order <br>
	Choose one of the values: <ul>
	<li><b>asc</b> — to set the ascending order;</li>
	<li><b>desc</b> — to set the descending order</li>
	</ul>
	*/
	OrderDirection *string

	/* ShowDeleted.

	     Specify whether you want to see deleted storages in the output.<br>
	Choose one of the values: <ul>
	<li><b>true</b> — to get details of deleted storages;</li>
	<li><b>false</b> — to get details about existing storages only</li>
	</ul>
	*/
	ShowDeleted *bool

	/* Status.

	   Specify a storage status that should be displayed in the output
	*/
	Status *string

	/* Type.

	   Specify a storage type that should be displayed in the output
	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the storage list Http v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageListHTTPV2Params) WithDefaults() *StorageListHTTPV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the storage list Http v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageListHTTPV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the storage list Http v2 params
func (o *StorageListHTTPV2Params) WithTimeout(timeout time.Duration) *StorageListHTTPV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage list Http v2 params
func (o *StorageListHTTPV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage list Http v2 params
func (o *StorageListHTTPV2Params) WithContext(ctx context.Context) *StorageListHTTPV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage list Http v2 params
func (o *StorageListHTTPV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage list Http v2 params
func (o *StorageListHTTPV2Params) WithHTTPClient(client *http.Client) *StorageListHTTPV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage list Http v2 params
func (o *StorageListHTTPV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the storage list Http v2 params
func (o *StorageListHTTPV2Params) WithID(id *string) *StorageListHTTPV2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the storage list Http v2 params
func (o *StorageListHTTPV2Params) SetID(id *string) {
	o.ID = id
}

// WithLimit adds the limit to the storage list Http v2 params
func (o *StorageListHTTPV2Params) WithLimit(limit *uint64) *StorageListHTTPV2Params {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the storage list Http v2 params
func (o *StorageListHTTPV2Params) SetLimit(limit *uint64) {
	o.Limit = limit
}

// WithLocation adds the location to the storage list Http v2 params
func (o *StorageListHTTPV2Params) WithLocation(location *string) *StorageListHTTPV2Params {
	o.SetLocation(location)
	return o
}

// SetLocation adds the location to the storage list Http v2 params
func (o *StorageListHTTPV2Params) SetLocation(location *string) {
	o.Location = location
}

// WithName adds the name to the storage list Http v2 params
func (o *StorageListHTTPV2Params) WithName(name *string) *StorageListHTTPV2Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the storage list Http v2 params
func (o *StorageListHTTPV2Params) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the storage list Http v2 params
func (o *StorageListHTTPV2Params) WithOffset(offset *uint64) *StorageListHTTPV2Params {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the storage list Http v2 params
func (o *StorageListHTTPV2Params) SetOffset(offset *uint64) {
	o.Offset = offset
}

// WithOrderBy adds the orderBy to the storage list Http v2 params
func (o *StorageListHTTPV2Params) WithOrderBy(orderBy *string) *StorageListHTTPV2Params {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the storage list Http v2 params
func (o *StorageListHTTPV2Params) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WithOrderDirection adds the orderDirection to the storage list Http v2 params
func (o *StorageListHTTPV2Params) WithOrderDirection(orderDirection *string) *StorageListHTTPV2Params {
	o.SetOrderDirection(orderDirection)
	return o
}

// SetOrderDirection adds the orderDirection to the storage list Http v2 params
func (o *StorageListHTTPV2Params) SetOrderDirection(orderDirection *string) {
	o.OrderDirection = orderDirection
}

// WithShowDeleted adds the showDeleted to the storage list Http v2 params
func (o *StorageListHTTPV2Params) WithShowDeleted(showDeleted *bool) *StorageListHTTPV2Params {
	o.SetShowDeleted(showDeleted)
	return o
}

// SetShowDeleted adds the showDeleted to the storage list Http v2 params
func (o *StorageListHTTPV2Params) SetShowDeleted(showDeleted *bool) {
	o.ShowDeleted = showDeleted
}

// WithStatus adds the status to the storage list Http v2 params
func (o *StorageListHTTPV2Params) WithStatus(status *string) *StorageListHTTPV2Params {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the storage list Http v2 params
func (o *StorageListHTTPV2Params) SetStatus(status *string) {
	o.Status = status
}

// WithType adds the typeVar to the storage list Http v2 params
func (o *StorageListHTTPV2Params) WithType(typeVar *string) *StorageListHTTPV2Params {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the storage list Http v2 params
func (o *StorageListHTTPV2Params) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *StorageListHTTPV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
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

	if o.Location != nil {

		// query param location
		var qrLocation string

		if o.Location != nil {
			qrLocation = *o.Location
		}
		qLocation := qrLocation
		if qLocation != "" {

			if err := r.SetQueryParam("location", qLocation); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
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

	if o.OrderBy != nil {

		// query param order_by
		var qrOrderBy string

		if o.OrderBy != nil {
			qrOrderBy = *o.OrderBy
		}
		qOrderBy := qrOrderBy
		if qOrderBy != "" {

			if err := r.SetQueryParam("order_by", qOrderBy); err != nil {
				return err
			}
		}
	}

	if o.OrderDirection != nil {

		// query param order_direction
		var qrOrderDirection string

		if o.OrderDirection != nil {
			qrOrderDirection = *o.OrderDirection
		}
		qOrderDirection := qrOrderDirection
		if qOrderDirection != "" {

			if err := r.SetQueryParam("order_direction", qOrderDirection); err != nil {
				return err
			}
		}
	}

	if o.ShowDeleted != nil {

		// query param show_deleted
		var qrShowDeleted bool

		if o.ShowDeleted != nil {
			qrShowDeleted = *o.ShowDeleted
		}
		qShowDeleted := swag.FormatBool(qrShowDeleted)
		if qShowDeleted != "" {

			if err := r.SetQueryParam("show_deleted", qShowDeleted); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
