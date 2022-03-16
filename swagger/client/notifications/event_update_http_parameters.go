// Code generated by go-swagger; DO NOT EDIT.

package notifications

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

	"github.com/G-Core/gcorelabs-storage-sdk-go/swagger/models"
)

// NewEventUpdateHTTPParams creates a new EventUpdateHTTPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEventUpdateHTTPParams() *EventUpdateHTTPParams {
	return &EventUpdateHTTPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEventUpdateHTTPParamsWithTimeout creates a new EventUpdateHTTPParams object
// with the ability to set a timeout on a request.
func NewEventUpdateHTTPParamsWithTimeout(timeout time.Duration) *EventUpdateHTTPParams {
	return &EventUpdateHTTPParams{
		timeout: timeout,
	}
}

// NewEventUpdateHTTPParamsWithContext creates a new EventUpdateHTTPParams object
// with the ability to set a context for a request.
func NewEventUpdateHTTPParamsWithContext(ctx context.Context) *EventUpdateHTTPParams {
	return &EventUpdateHTTPParams{
		Context: ctx,
	}
}

// NewEventUpdateHTTPParamsWithHTTPClient creates a new EventUpdateHTTPParams object
// with the ability to set a custom HTTPClient for a request.
func NewEventUpdateHTTPParamsWithHTTPClient(client *http.Client) *EventUpdateHTTPParams {
	return &EventUpdateHTTPParams{
		HTTPClient: client,
	}
}

/* EventUpdateHTTPParams contains all the parameters to send to the API endpoint
   for the event update Http operation.

   Typically these are written to a http.Request.
*/
type EventUpdateHTTPParams struct {

	// Body.
	Body *models.Event

	/* ClientID.

	   ClientID who events we will get

	   Format: int64
	*/
	ClientID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the event update Http params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EventUpdateHTTPParams) WithDefaults() *EventUpdateHTTPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the event update Http params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EventUpdateHTTPParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the event update Http params
func (o *EventUpdateHTTPParams) WithTimeout(timeout time.Duration) *EventUpdateHTTPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the event update Http params
func (o *EventUpdateHTTPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the event update Http params
func (o *EventUpdateHTTPParams) WithContext(ctx context.Context) *EventUpdateHTTPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the event update Http params
func (o *EventUpdateHTTPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the event update Http params
func (o *EventUpdateHTTPParams) WithHTTPClient(client *http.Client) *EventUpdateHTTPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the event update Http params
func (o *EventUpdateHTTPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the event update Http params
func (o *EventUpdateHTTPParams) WithBody(body *models.Event) *EventUpdateHTTPParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the event update Http params
func (o *EventUpdateHTTPParams) SetBody(body *models.Event) {
	o.Body = body
}

// WithClientID adds the clientID to the event update Http params
func (o *EventUpdateHTTPParams) WithClientID(clientID *int64) *EventUpdateHTTPParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the event update Http params
func (o *EventUpdateHTTPParams) SetClientID(clientID *int64) {
	o.ClientID = clientID
}

// WriteToRequest writes these params to a swagger request
func (o *EventUpdateHTTPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.ClientID != nil {

		// query param client_id
		var qrClientID int64

		if o.ClientID != nil {
			qrClientID = *o.ClientID
		}
		qClientID := swag.FormatInt64(qrClientID)
		if qClientID != "" {

			if err := r.SetQueryParam("client_id", qClientID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}