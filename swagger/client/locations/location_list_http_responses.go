// Code generated by go-swagger; DO NOT EDIT.

package locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Edge-Center/edgecenter-storage-sdk-go/swagger/models"
)

// LocationListHTTPReader is a Reader for the LocationListHTTP structure.
type LocationListHTTPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocationListHTTPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocationListHTTPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLocationListHTTPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /resource/v3/location] locationListHttp", response, response.Code())
	}
}

// NewLocationListHTTPOK creates a LocationListHTTPOK with default headers values
func NewLocationListHTTPOK() *LocationListHTTPOK {
	return &LocationListHTTPOK{}
}

/*
LocationListHTTPOK describes a response with status code 200, with default header values.

clientLocationRes
*/
type LocationListHTTPOK struct {
	Payload []*models.ClientLocationRes
}

// IsSuccess returns true when this location list Http o k response has a 2xx status code
func (o *LocationListHTTPOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this location list Http o k response has a 3xx status code
func (o *LocationListHTTPOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this location list Http o k response has a 4xx status code
func (o *LocationListHTTPOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this location list Http o k response has a 5xx status code
func (o *LocationListHTTPOK) IsServerError() bool {
	return false
}

// IsCode returns true when this location list Http o k response a status code equal to that given
func (o *LocationListHTTPOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the location list Http o k response
func (o *LocationListHTTPOK) Code() int {
	return 200
}

func (o *LocationListHTTPOK) Error() string {
	return fmt.Sprintf("[GET /resource/v3/location][%d] locationListHttpOK  %+v", 200, o.Payload)
}

func (o *LocationListHTTPOK) String() string {
	return fmt.Sprintf("[GET /resource/v3/location][%d] locationListHttpOK  %+v", 200, o.Payload)
}

func (o *LocationListHTTPOK) GetPayload() []*models.ClientLocationRes {
	return o.Payload
}

func (o *LocationListHTTPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocationListHTTPBadRequest creates a LocationListHTTPBadRequest with default headers values
func NewLocationListHTTPBadRequest() *LocationListHTTPBadRequest {
	return &LocationListHTTPBadRequest{}
}

/*
LocationListHTTPBadRequest describes a response with status code 400, with default header values.

ErrResponse
*/
type LocationListHTTPBadRequest struct {
	Payload *models.ErrResponse
}

// IsSuccess returns true when this location list Http bad request response has a 2xx status code
func (o *LocationListHTTPBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this location list Http bad request response has a 3xx status code
func (o *LocationListHTTPBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this location list Http bad request response has a 4xx status code
func (o *LocationListHTTPBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this location list Http bad request response has a 5xx status code
func (o *LocationListHTTPBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this location list Http bad request response a status code equal to that given
func (o *LocationListHTTPBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the location list Http bad request response
func (o *LocationListHTTPBadRequest) Code() int {
	return 400
}

func (o *LocationListHTTPBadRequest) Error() string {
	return fmt.Sprintf("[GET /resource/v3/location][%d] locationListHttpBadRequest  %+v", 400, o.Payload)
}

func (o *LocationListHTTPBadRequest) String() string {
	return fmt.Sprintf("[GET /resource/v3/location][%d] locationListHttpBadRequest  %+v", 400, o.Payload)
}

func (o *LocationListHTTPBadRequest) GetPayload() *models.ErrResponse {
	return o.Payload
}

func (o *LocationListHTTPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
