// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Edge-Center/edgecenter-storage-sdk-go/swagger/models"
)

// StorageDeleteHTTPReader is a Reader for the StorageDeleteHTTP structure.
type StorageDeleteHTTPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageDeleteHTTPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewStorageDeleteHTTPNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStorageDeleteHTTPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStorageDeleteHTTPNoContent creates a StorageDeleteHTTPNoContent with default headers values
func NewStorageDeleteHTTPNoContent() *StorageDeleteHTTPNoContent {
	return &StorageDeleteHTTPNoContent{}
}

/*
StorageDeleteHTTPNoContent describes a response with status code 204, with default header values.

A SuccessResponse is a response that shows that operations was completed successfully
*/
type StorageDeleteHTTPNoContent struct {
}

// IsSuccess returns true when this storage delete Http no content response has a 2xx status code
func (o *StorageDeleteHTTPNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this storage delete Http no content response has a 3xx status code
func (o *StorageDeleteHTTPNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this storage delete Http no content response has a 4xx status code
func (o *StorageDeleteHTTPNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this storage delete Http no content response has a 5xx status code
func (o *StorageDeleteHTTPNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this storage delete Http no content response a status code equal to that given
func (o *StorageDeleteHTTPNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *StorageDeleteHTTPNoContent) Error() string {
	return fmt.Sprintf("[DELETE /provisioning/v1/storage/{id}][%d] storageDeleteHttpNoContent ", 204)
}

func (o *StorageDeleteHTTPNoContent) String() string {
	return fmt.Sprintf("[DELETE /provisioning/v1/storage/{id}][%d] storageDeleteHttpNoContent ", 204)
}

func (o *StorageDeleteHTTPNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageDeleteHTTPBadRequest creates a StorageDeleteHTTPBadRequest with default headers values
func NewStorageDeleteHTTPBadRequest() *StorageDeleteHTTPBadRequest {
	return &StorageDeleteHTTPBadRequest{}
}

/*
StorageDeleteHTTPBadRequest describes a response with status code 400, with default header values.

ErrResponse
*/
type StorageDeleteHTTPBadRequest struct {
	Payload *models.ErrResponse
}

// IsSuccess returns true when this storage delete Http bad request response has a 2xx status code
func (o *StorageDeleteHTTPBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this storage delete Http bad request response has a 3xx status code
func (o *StorageDeleteHTTPBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this storage delete Http bad request response has a 4xx status code
func (o *StorageDeleteHTTPBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this storage delete Http bad request response has a 5xx status code
func (o *StorageDeleteHTTPBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this storage delete Http bad request response a status code equal to that given
func (o *StorageDeleteHTTPBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StorageDeleteHTTPBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /provisioning/v1/storage/{id}][%d] storageDeleteHttpBadRequest  %+v", 400, o.Payload)
}

func (o *StorageDeleteHTTPBadRequest) String() string {
	return fmt.Sprintf("[DELETE /provisioning/v1/storage/{id}][%d] storageDeleteHttpBadRequest  %+v", 400, o.Payload)
}

func (o *StorageDeleteHTTPBadRequest) GetPayload() *models.ErrResponse {
	return o.Payload
}

func (o *StorageDeleteHTTPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
