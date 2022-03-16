// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/G-Core/gcore-storage-sdk-go/swagger/models"
)

// StorageBucketLifecycleDeleteHTTPReader is a Reader for the StorageBucketLifecycleDeleteHTTP structure.
type StorageBucketLifecycleDeleteHTTPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageBucketLifecycleDeleteHTTPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewStorageBucketLifecycleDeleteHTTPNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStorageBucketLifecycleDeleteHTTPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStorageBucketLifecycleDeleteHTTPNoContent creates a StorageBucketLifecycleDeleteHTTPNoContent with default headers values
func NewStorageBucketLifecycleDeleteHTTPNoContent() *StorageBucketLifecycleDeleteHTTPNoContent {
	return &StorageBucketLifecycleDeleteHTTPNoContent{}
}

/* StorageBucketLifecycleDeleteHTTPNoContent describes a response with status code 204, with default header values.

A SuccessResponse is a response that shows that operations was completed successfully
*/
type StorageBucketLifecycleDeleteHTTPNoContent struct {
}

func (o *StorageBucketLifecycleDeleteHTTPNoContent) Error() string {
	return fmt.Sprintf("[DELETE /provisioning/v1/storage/{id}/s3/bucket/{name}/lifecycle][%d] storageBucketLifecycleDeleteHttpNoContent ", 204)
}

func (o *StorageBucketLifecycleDeleteHTTPNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageBucketLifecycleDeleteHTTPBadRequest creates a StorageBucketLifecycleDeleteHTTPBadRequest with default headers values
func NewStorageBucketLifecycleDeleteHTTPBadRequest() *StorageBucketLifecycleDeleteHTTPBadRequest {
	return &StorageBucketLifecycleDeleteHTTPBadRequest{}
}

/* StorageBucketLifecycleDeleteHTTPBadRequest describes a response with status code 400, with default header values.

ErrResponse
*/
type StorageBucketLifecycleDeleteHTTPBadRequest struct {
	Payload *models.ErrResponse
}

func (o *StorageBucketLifecycleDeleteHTTPBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /provisioning/v1/storage/{id}/s3/bucket/{name}/lifecycle][%d] storageBucketLifecycleDeleteHttpBadRequest  %+v", 400, o.Payload)
}
func (o *StorageBucketLifecycleDeleteHTTPBadRequest) GetPayload() *models.ErrResponse {
	return o.Payload
}

func (o *StorageBucketLifecycleDeleteHTTPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
