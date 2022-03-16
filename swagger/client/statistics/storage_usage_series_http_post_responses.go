// Code generated by go-swagger; DO NOT EDIT.

package statistics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/G-Core/gcore-storage-sdk-go/swagger/models"
)

// StorageUsageSeriesHTTPPostReader is a Reader for the StorageUsageSeriesHTTPPost structure.
type StorageUsageSeriesHTTPPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageUsageSeriesHTTPPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageUsageSeriesHTTPPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStorageUsageSeriesHTTPPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStorageUsageSeriesHTTPPostOK creates a StorageUsageSeriesHTTPPostOK with default headers values
func NewStorageUsageSeriesHTTPPostOK() *StorageUsageSeriesHTTPPostOK {
	return &StorageUsageSeriesHTTPPostOK{}
}

/* StorageUsageSeriesHTTPPostOK describes a response with status code 200, with default header values.

StorageUsageSeriesEndpointRes
*/
type StorageUsageSeriesHTTPPostOK struct {
	Payload *models.StorageUsageSeriesEndpointRes
}

func (o *StorageUsageSeriesHTTPPostOK) Error() string {
	return fmt.Sprintf("[POST /stats/v1/storage/usage/series][%d] storageUsageSeriesHttpPostOK  %+v", 200, o.Payload)
}
func (o *StorageUsageSeriesHTTPPostOK) GetPayload() *models.StorageUsageSeriesEndpointRes {
	return o.Payload
}

func (o *StorageUsageSeriesHTTPPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageUsageSeriesEndpointRes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageUsageSeriesHTTPPostBadRequest creates a StorageUsageSeriesHTTPPostBadRequest with default headers values
func NewStorageUsageSeriesHTTPPostBadRequest() *StorageUsageSeriesHTTPPostBadRequest {
	return &StorageUsageSeriesHTTPPostBadRequest{}
}

/* StorageUsageSeriesHTTPPostBadRequest describes a response with status code 400, with default header values.

ErrResponse
*/
type StorageUsageSeriesHTTPPostBadRequest struct {
	Payload *models.ErrResponse
}

func (o *StorageUsageSeriesHTTPPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /stats/v1/storage/usage/series][%d] storageUsageSeriesHttpPostBadRequest  %+v", 400, o.Payload)
}
func (o *StorageUsageSeriesHTTPPostBadRequest) GetPayload() *models.ErrResponse {
	return o.Payload
}

func (o *StorageUsageSeriesHTTPPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StorageUsageSeriesHTTPPostBody storage usage series HTTP post body
swagger:model StorageUsageSeriesHTTPPostBody
*/
type StorageUsageSeriesHTTPPostBody struct {

	// a From date filter
	// Example: 2006-01-02
	From string `json:"from,omitempty"`

	// a Granularity is period of time for grouping data
	// Valid time units are "s", "m", "h".
	// Example: 12h
	Granularity string `json:"granularity,omitempty"`

	// a Locations list of filter
	// Example: ["fra","mia","sin","ams","s-ed1","s-darz1","s-dt2","lux","drf"]
	Locations []string `json:"locations"`

	// a Source is deprecated parameter
	Source uint8 `json:"source,omitempty"`

	// a Storages list of filter
	// Example: ["123-myStorage"]
	Storages []string `json:"storages"`

	// a To date filter
	// Example: 2006-01-02
	To string `json:"to,omitempty"`

	// a TsString is configurator of response time format
	// switch response from unix time format to RFC3339 (2006-01-02T15:04:05Z07:00)
	TsString bool `json:"ts_string,omitempty"`
}

// Validate validates this storage usage series HTTP post body
func (o *StorageUsageSeriesHTTPPostBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage usage series HTTP post body based on context it is used
func (o *StorageUsageSeriesHTTPPostBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StorageUsageSeriesHTTPPostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StorageUsageSeriesHTTPPostBody) UnmarshalBinary(b []byte) error {
	var res StorageUsageSeriesHTTPPostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
