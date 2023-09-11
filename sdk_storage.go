package ec_storage

import (
	"fmt"

	"github.com/Edge-Center/edgecenter-storage-sdk-go/swagger/client/storages"
	"github.com/Edge-Center/edgecenter-storage-sdk-go/swagger/models"
)

type sdkStorage struct {
	*apiCore
}

// StoragesList getter for EdgeCenter Storage API
// same result like on UI here https://api.edgecenter.ru/storage/list
func (sdk *sdkStorage) StoragesList(opts ...func(params *storages.StorageListHTTPV2Params)) ([]models.Storage, error) {
	params := &storages.StorageListHTTPV2Params{}
	for _, opt := range opts {
		opt(params)
	}
	res, err := sdk.client.Storages.StorageListHTTPV2(params, sdk.authWriter)
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}
	list := make([]models.Storage, len(res.Payload.Data))
	for i := range res.Payload.Data {
		list[i] = *res.Payload.Data[i]
	}
	return list, nil
}

// CreateStorage writer for EdgeCenter Storage API
func (sdk *sdkStorage) CreateStorage(opts ...func(params *storages.StorageCreateHTTPParams)) (*models.Storage, error) {
	params := &storages.StorageCreateHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	res, err := sdk.client.Storages.StorageCreateHTTP(params, sdk.authWriter)
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}
	return res.Payload, nil
}

// DeleteStorage writer for EdgeCenter Storage API
// be noticed that delete action is async in ec end service
func (sdk *sdkStorage) DeleteStorage(opts ...func(params *storages.StorageDeleteHTTPParams)) error {
	params := &storages.StorageDeleteHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	_, err := sdk.client.Storages.StorageDeleteHTTP(params, sdk.authWriter)
	if err != nil {
		return fmt.Errorf("request: %w", err)
	}
	return nil
}

// UpdateStorageCredentials writer for EdgeCenter Storage API
func (sdk *sdkStorage) UpdatestoragesCredentials(
	opts ...func(params *storages.StorageUpdateCredentialsHTTPParams)) (*models.Credentials, error) {

	params := &storages.StorageUpdateCredentialsHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	res, err := sdk.client.Storages.StorageUpdateCredentialsHTTP(params, sdk.authWriter)
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}
	if res.Payload == nil {
		// nolint:typecheck
		return nil, fmt.Errorf("empty: %w", EmptyResultErr("response payload"))
	}
	return res.Payload.Credentials, nil
}
