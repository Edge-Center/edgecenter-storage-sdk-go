// Code generated by go-swagger; DO NOT EDIT.

package storages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new storages API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for storages API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	StorageCreateHTTP(params *StorageCreateHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageCreateHTTPOK, error)

	StorageDeleteHTTP(params *StorageDeleteHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageDeleteHTTPNoContent, error)

	StorageGetHTTP(params *StorageGetHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageGetHTTPOK, error)

	StorageListHTTPV2(params *StorageListHTTPV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageListHTTPV2OK, error)

	StorageRestoreHTTP(params *StorageRestoreHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageRestoreHTTPNoContent, error)

	StorageUpdateCredentialsHTTP(params *StorageUpdateCredentialsHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageUpdateCredentialsHTTPOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
StorageCreateHTTP creates a storage

Creates a storage and returns it.
*/
func (a *Client) StorageCreateHTTP(params *StorageCreateHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageCreateHTTPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStorageCreateHTTPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "storageCreateHttp",
		Method:             "PUT",
		PathPattern:        "/resource/v3/storage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StorageCreateHTTPReader{formats: a.formats},
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
	success, ok := result.(*StorageCreateHTTPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for storageCreateHttp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StorageDeleteHTTP deletes a storage

Deletes a storage from your account.
*/
func (a *Client) StorageDeleteHTTP(params *StorageDeleteHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageDeleteHTTPNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStorageDeleteHTTPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "storageDeleteHttp",
		Method:             "DELETE",
		PathPattern:        "/resource/v3/storage/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StorageDeleteHTTPReader{formats: a.formats},
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
	success, ok := result.(*StorageDeleteHTTPNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for storageDeleteHttp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StorageGetHTTP gets a storage

Returns information about a specific storage in your account.
*/
func (a *Client) StorageGetHTTP(params *StorageGetHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageGetHTTPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStorageGetHTTPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "storageGetHttp",
		Method:             "GET",
		PathPattern:        "/resource/v3/storage/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StorageGetHTTPReader{formats: a.formats},
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
	success, ok := result.(*StorageGetHTTPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for storageGetHttp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StorageListHTTPV2 gets all storages

Returns information about all storages in your account
*/
func (a *Client) StorageListHTTPV2(params *StorageListHTTPV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageListHTTPV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStorageListHTTPV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "storageListHttpV2",
		Method:             "GET",
		PathPattern:        "/resource/v3/storage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StorageListHTTPV2Reader{formats: a.formats},
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
	success, ok := result.(*StorageListHTTPV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for storageListHttpV2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StorageRestoreHTTP recovers a deleted s3 storage

Recovers S3 storage that was deleted in the last 2 weeks. The storage will appear in your account.
*/
func (a *Client) StorageRestoreHTTP(params *StorageRestoreHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageRestoreHTTPNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStorageRestoreHTTPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "storageRestoreHttp",
		Method:             "POST",
		PathPattern:        "/resource/v3/storage/{id}/restore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StorageRestoreHTTPReader{formats: a.formats},
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
	success, ok := result.(*StorageRestoreHTTPNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for storageRestoreHttp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StorageUpdateCredentialsHTTP updates a storage credentials

Changes the credentials of your storage.
*/
func (a *Client) StorageUpdateCredentialsHTTP(params *StorageUpdateCredentialsHTTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StorageUpdateCredentialsHTTPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStorageUpdateCredentialsHTTPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "storageUpdateCredentialsHttp",
		Method:             "POST",
		PathPattern:        "/resource/v3/storage/{id}/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StorageUpdateCredentialsHTTPReader{formats: a.formats},
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
	success, ok := result.(*StorageUpdateCredentialsHTTPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for storageUpdateCredentialsHttp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}