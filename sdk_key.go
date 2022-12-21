package ec_storage

import (
	"fmt"

	"github.com/Edge-Center/edgecenter-storage-sdk-go/swagger/client/key"
	"github.com/Edge-Center/edgecenter-storage-sdk-go/swagger/models"
)

type sdkKey struct {
	*apiCore
}

// KeysList getter for EdgeCenter Storage API
// same result like on UI here https://api.edgecenter.ru/storage/ssh-key/list
func (sdk *sdkKey) KeysList(opts ...func(*key.KeyListHTTPV2Params)) ([]models.Key, error) {
	params := &key.KeyListHTTPV2Params{}
	for _, opt := range opts {
		opt(params)
	}
	res, err := sdk.client.Key.KeyListHTTPV2(params, sdk.authWriter)
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}
	list := make([]models.Key, len(res.Payload.Data))
	for i := range res.Payload.Data {
		list[i] = *res.Payload.Data[i]
	}
	return list, nil
}

// CreateKey writer for EdgeCenter Storage API
func (sdk *sdkKey) CreateKey(opts ...func(*key.KeyCreateHTTPParams)) (*models.Key, error) {
	params := &key.KeyCreateHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	res, err := sdk.client.Key.KeyCreateHTTP(params, sdk.authWriter)
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}
	return res.Payload, nil
}

// DeleteKey writer for EdgeCenter Storage API
func (sdk *sdkKey) DeleteKey(opts ...func(*key.KeyDeleteHTTPParams)) error {
	params := &key.KeyDeleteHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	_, err := sdk.client.Key.KeyDeleteHTTP(params, sdk.authWriter)
	if err != nil {
		return fmt.Errorf("request: %w", err)
	}
	return nil
}
