// Code generated by go-swagger; DO NOT EDIT.

package buckets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Edge-Center/edgecenter-storage-sdk-go/swagger/models"
)

// StorageBucketPolicyCreateHTTPReader is a Reader for the StorageBucketPolicyCreateHTTP structure.
type StorageBucketPolicyCreateHTTPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageBucketPolicyCreateHTTPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewStorageBucketPolicyCreateHTTPNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStorageBucketPolicyCreateHTTPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /resource/v3/storage/{id}/s3/bucket/{name}/policy] storageBucketPolicyCreateHttp", response, response.Code())
	}
}

// NewStorageBucketPolicyCreateHTTPNoContent creates a StorageBucketPolicyCreateHTTPNoContent with default headers values
func NewStorageBucketPolicyCreateHTTPNoContent() *StorageBucketPolicyCreateHTTPNoContent {
	return &StorageBucketPolicyCreateHTTPNoContent{}
}

/*
StorageBucketPolicyCreateHTTPNoContent describes a response with status code 204, with default header values.

A SuccessResponse is a response that shows that operations was completed successfully
*/
type StorageBucketPolicyCreateHTTPNoContent struct {
}

// IsSuccess returns true when this storage bucket policy create Http no content response has a 2xx status code
func (o *StorageBucketPolicyCreateHTTPNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this storage bucket policy create Http no content response has a 3xx status code
func (o *StorageBucketPolicyCreateHTTPNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this storage bucket policy create Http no content response has a 4xx status code
func (o *StorageBucketPolicyCreateHTTPNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this storage bucket policy create Http no content response has a 5xx status code
func (o *StorageBucketPolicyCreateHTTPNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this storage bucket policy create Http no content response a status code equal to that given
func (o *StorageBucketPolicyCreateHTTPNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the storage bucket policy create Http no content response
func (o *StorageBucketPolicyCreateHTTPNoContent) Code() int {
	return 204
}

func (o *StorageBucketPolicyCreateHTTPNoContent) Error() string {
	return fmt.Sprintf("[POST /resource/v3/storage/{id}/s3/bucket/{name}/policy][%d] storageBucketPolicyCreateHttpNoContent ", 204)
}

func (o *StorageBucketPolicyCreateHTTPNoContent) String() string {
	return fmt.Sprintf("[POST /resource/v3/storage/{id}/s3/bucket/{name}/policy][%d] storageBucketPolicyCreateHttpNoContent ", 204)
}

func (o *StorageBucketPolicyCreateHTTPNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageBucketPolicyCreateHTTPBadRequest creates a StorageBucketPolicyCreateHTTPBadRequest with default headers values
func NewStorageBucketPolicyCreateHTTPBadRequest() *StorageBucketPolicyCreateHTTPBadRequest {
	return &StorageBucketPolicyCreateHTTPBadRequest{}
}

/*
StorageBucketPolicyCreateHTTPBadRequest describes a response with status code 400, with default header values.

ErrResponse
*/
type StorageBucketPolicyCreateHTTPBadRequest struct {
	Payload *models.ErrResponse
}

// IsSuccess returns true when this storage bucket policy create Http bad request response has a 2xx status code
func (o *StorageBucketPolicyCreateHTTPBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this storage bucket policy create Http bad request response has a 3xx status code
func (o *StorageBucketPolicyCreateHTTPBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this storage bucket policy create Http bad request response has a 4xx status code
func (o *StorageBucketPolicyCreateHTTPBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this storage bucket policy create Http bad request response has a 5xx status code
func (o *StorageBucketPolicyCreateHTTPBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this storage bucket policy create Http bad request response a status code equal to that given
func (o *StorageBucketPolicyCreateHTTPBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the storage bucket policy create Http bad request response
func (o *StorageBucketPolicyCreateHTTPBadRequest) Code() int {
	return 400
}

func (o *StorageBucketPolicyCreateHTTPBadRequest) Error() string {
	return fmt.Sprintf("[POST /resource/v3/storage/{id}/s3/bucket/{name}/policy][%d] storageBucketPolicyCreateHttpBadRequest  %+v", 400, o.Payload)
}

func (o *StorageBucketPolicyCreateHTTPBadRequest) String() string {
	return fmt.Sprintf("[POST /resource/v3/storage/{id}/s3/bucket/{name}/policy][%d] storageBucketPolicyCreateHttpBadRequest  %+v", 400, o.Payload)
}

func (o *StorageBucketPolicyCreateHTTPBadRequest) GetPayload() *models.ErrResponse {
	return o.Payload
}

func (o *StorageBucketPolicyCreateHTTPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
