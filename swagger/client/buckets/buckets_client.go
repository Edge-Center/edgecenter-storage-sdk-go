// Code generated by go-swagger; DO NOT EDIT.

package buckets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new buckets API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for buckets API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetStorageBucketCORSHTTP(params *GetStorageBucketCORSHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStorageBucketCORSHTTPOK, error)

	StorageBucketCORSCreateHTTP(params *StorageBucketCORSCreateHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageBucketCORSCreateHTTPNoContent, error)

	StorageBucketCreateHTTP(params *StorageBucketCreateHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageBucketCreateHTTPNoContent, error)

	StorageBucketLifecycleCreateHTTP(params *StorageBucketLifecycleCreateHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageBucketLifecycleCreateHTTPNoContent, error)

	StorageBucketLifecycleDeleteHTTP(params *StorageBucketLifecycleDeleteHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageBucketLifecycleDeleteHTTPNoContent, error)

	StorageBucketPolicyCreateHTTP(params *StorageBucketPolicyCreateHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageBucketPolicyCreateHTTPNoContent, error)

	StorageBucketRemoveHTTP(params *StorageBucketRemoveHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageBucketRemoveHTTPNoContent, error)

	StorageListBucketsHTTP(params *StorageListBucketsHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageListBucketsHTTPOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetStorageBucketCORSHTTP gets a bucket c o r s

Returns the CORS configuration for your bucket.
*/
func (a *Client) GetStorageBucketCORSHTTP(params *GetStorageBucketCORSHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStorageBucketCORSHTTPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStorageBucketCORSHTTPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getStorageBucketCORSHttp",
		Method:             "GET",
		PathPattern:        "/resource/v3/storage/{id}/s3/bucket/{name}/cors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStorageBucketCORSHTTPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStorageBucketCORSHTTPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getStorageBucketCORSHttp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StorageBucketCORSCreateHTTP creates a bucket c o r s

Creates the CORS configuration for your bucket.
*/
func (a *Client) StorageBucketCORSCreateHTTP(params *StorageBucketCORSCreateHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageBucketCORSCreateHTTPNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStorageBucketCORSCreateHTTPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "storageBucketCORSCreateHttp",
		Method:             "POST",
		PathPattern:        "/resource/v3/storage/{id}/s3/bucket/{name}/cors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StorageBucketCORSCreateHTTPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StorageBucketCORSCreateHTTPNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for storageBucketCORSCreateHttp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StorageBucketCreateHTTP creates a bucket

Creates a bucket in your S3 storage.
*/
func (a *Client) StorageBucketCreateHTTP(params *StorageBucketCreateHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageBucketCreateHTTPNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStorageBucketCreateHTTPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "storageBucketCreateHttp",
		Method:             "POST",
		PathPattern:        "/resource/v3/storage/{id}/s3/bucket/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StorageBucketCreateHTTPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StorageBucketCreateHTTPNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for storageBucketCreateHttp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	StorageBucketLifecycleCreateHTTP creates a bucket lifecycle

	Creates a lifecycle for your bucket. Lifecycle determines the number of days after which the objects in the

bucket are considered expired.
*/
func (a *Client) StorageBucketLifecycleCreateHTTP(params *StorageBucketLifecycleCreateHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageBucketLifecycleCreateHTTPNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStorageBucketLifecycleCreateHTTPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "storageBucketLifecycleCreateHttp",
		Method:             "POST",
		PathPattern:        "/resource/v3/storage/{id}/s3/bucket/{name}/lifecycle",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StorageBucketLifecycleCreateHTTPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StorageBucketLifecycleCreateHTTPNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for storageBucketLifecycleCreateHttp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StorageBucketLifecycleDeleteHTTP deletes a bucket lifecycle

Delete a bucket lifecycle
*/
func (a *Client) StorageBucketLifecycleDeleteHTTP(params *StorageBucketLifecycleDeleteHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageBucketLifecycleDeleteHTTPNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStorageBucketLifecycleDeleteHTTPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "storageBucketLifecycleDeleteHttp",
		Method:             "DELETE",
		PathPattern:        "/resource/v3/storage/{id}/s3/bucket/{name}/lifecycle",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StorageBucketLifecycleDeleteHTTPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StorageBucketLifecycleDeleteHTTPNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for storageBucketLifecycleDeleteHttp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	StorageBucketPolicyCreateHTTP creates a bucket policy

	Creates a policy for your bucket. Policies are used to set permissions for actions on buckets, objects, and

groups of objects.
*/
func (a *Client) StorageBucketPolicyCreateHTTP(params *StorageBucketPolicyCreateHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageBucketPolicyCreateHTTPNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStorageBucketPolicyCreateHTTPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "storageBucketPolicyCreateHttp",
		Method:             "POST",
		PathPattern:        "/resource/v3/storage/{id}/s3/bucket/{name}/policy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StorageBucketPolicyCreateHTTPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StorageBucketPolicyCreateHTTPNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for storageBucketPolicyCreateHttp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StorageBucketRemoveHTTP deletes a bucket

Deletes a bucket from your S3 storage.
*/
func (a *Client) StorageBucketRemoveHTTP(params *StorageBucketRemoveHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageBucketRemoveHTTPNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStorageBucketRemoveHTTPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "storageBucketRemoveHttp",
		Method:             "DELETE",
		PathPattern:        "/resource/v3/storage/{id}/s3/bucket/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StorageBucketRemoveHTTPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StorageBucketRemoveHTTPNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for storageBucketRemoveHttp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StorageListBucketsHTTP gets all buckets

Returns information about all the buckets created in your S3 storage.
*/
func (a *Client) StorageListBucketsHTTP(params *StorageListBucketsHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageListBucketsHTTPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStorageListBucketsHTTPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "storageListBucketsHttp",
		Method:             "GET",
		PathPattern:        "/resource/v3/storage/{id}/s3/buckets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StorageListBucketsHTTPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StorageListBucketsHTTPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for storageListBucketsHttp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
