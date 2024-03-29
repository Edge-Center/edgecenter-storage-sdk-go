// Code generated by go-swagger; DO NOT EDIT.

package buckets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/Edge-Center/edgecenter-storage-sdk-go/swagger/models"
)

// StorageBucketCORSCreateHTTPReader is a Reader for the StorageBucketCORSCreateHTTP structure.
type StorageBucketCORSCreateHTTPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageBucketCORSCreateHTTPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewStorageBucketCORSCreateHTTPNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStorageBucketCORSCreateHTTPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /resource/v3/storage/{id}/s3/bucket/{name}/cors] storageBucketCORSCreateHttp", response, response.Code())
	}
}

// NewStorageBucketCORSCreateHTTPNoContent creates a StorageBucketCORSCreateHTTPNoContent with default headers values
func NewStorageBucketCORSCreateHTTPNoContent() *StorageBucketCORSCreateHTTPNoContent {
	return &StorageBucketCORSCreateHTTPNoContent{}
}

/*
StorageBucketCORSCreateHTTPNoContent describes a response with status code 204, with default header values.

A SuccessResponse is a response that shows that operations was completed successfully
*/
type StorageBucketCORSCreateHTTPNoContent struct {
}

// IsSuccess returns true when this storage bucket c o r s create Http no content response has a 2xx status code
func (o *StorageBucketCORSCreateHTTPNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this storage bucket c o r s create Http no content response has a 3xx status code
func (o *StorageBucketCORSCreateHTTPNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this storage bucket c o r s create Http no content response has a 4xx status code
func (o *StorageBucketCORSCreateHTTPNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this storage bucket c o r s create Http no content response has a 5xx status code
func (o *StorageBucketCORSCreateHTTPNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this storage bucket c o r s create Http no content response a status code equal to that given
func (o *StorageBucketCORSCreateHTTPNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the storage bucket c o r s create Http no content response
func (o *StorageBucketCORSCreateHTTPNoContent) Code() int {
	return 204
}

func (o *StorageBucketCORSCreateHTTPNoContent) Error() string {
	return fmt.Sprintf("[POST /resource/v3/storage/{id}/s3/bucket/{name}/cors][%d] storageBucketCORSCreateHttpNoContent ", 204)
}

func (o *StorageBucketCORSCreateHTTPNoContent) String() string {
	return fmt.Sprintf("[POST /resource/v3/storage/{id}/s3/bucket/{name}/cors][%d] storageBucketCORSCreateHttpNoContent ", 204)
}

func (o *StorageBucketCORSCreateHTTPNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageBucketCORSCreateHTTPBadRequest creates a StorageBucketCORSCreateHTTPBadRequest with default headers values
func NewStorageBucketCORSCreateHTTPBadRequest() *StorageBucketCORSCreateHTTPBadRequest {
	return &StorageBucketCORSCreateHTTPBadRequest{}
}

/*
StorageBucketCORSCreateHTTPBadRequest describes a response with status code 400, with default header values.

ErrResponse
*/
type StorageBucketCORSCreateHTTPBadRequest struct {
	Payload *models.ErrResponse
}

// IsSuccess returns true when this storage bucket c o r s create Http bad request response has a 2xx status code
func (o *StorageBucketCORSCreateHTTPBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this storage bucket c o r s create Http bad request response has a 3xx status code
func (o *StorageBucketCORSCreateHTTPBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this storage bucket c o r s create Http bad request response has a 4xx status code
func (o *StorageBucketCORSCreateHTTPBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this storage bucket c o r s create Http bad request response has a 5xx status code
func (o *StorageBucketCORSCreateHTTPBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this storage bucket c o r s create Http bad request response a status code equal to that given
func (o *StorageBucketCORSCreateHTTPBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the storage bucket c o r s create Http bad request response
func (o *StorageBucketCORSCreateHTTPBadRequest) Code() int {
	return 400
}

func (o *StorageBucketCORSCreateHTTPBadRequest) Error() string {
	return fmt.Sprintf("[POST /resource/v3/storage/{id}/s3/bucket/{name}/cors][%d] storageBucketCORSCreateHttpBadRequest  %+v", 400, o.Payload)
}

func (o *StorageBucketCORSCreateHTTPBadRequest) String() string {
	return fmt.Sprintf("[POST /resource/v3/storage/{id}/s3/bucket/{name}/cors][%d] storageBucketCORSCreateHttpBadRequest  %+v", 400, o.Payload)
}

func (o *StorageBucketCORSCreateHTTPBadRequest) GetPayload() *models.ErrResponse {
	return o.Payload
}

func (o *StorageBucketCORSCreateHTTPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
StorageBucketCORSCreateHTTPBody storage bucket c o r s create HTTP body
swagger:model StorageBucketCORSCreateHTTPBody
*/
type StorageBucketCORSCreateHTTPBody struct {

	// Specify the origins for your CORS configuration
	AllowedOrigins []string `json:"allowedOrigins"`
}

// Validate validates this storage bucket c o r s create HTTP body
func (o *StorageBucketCORSCreateHTTPBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage bucket c o r s create HTTP body based on context it is used
func (o *StorageBucketCORSCreateHTTPBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StorageBucketCORSCreateHTTPBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StorageBucketCORSCreateHTTPBody) UnmarshalBinary(b []byte) error {
	var res StorageBucketCORSCreateHTTPBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
