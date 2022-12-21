package ec_storage

import (
	"fmt"

	"github.com/Edge-Center/edgecenter-storage-sdk-go/swagger/client/storage"
	"github.com/Edge-Center/edgecenter-storage-sdk-go/swagger/models"
)

type sdkBucket struct {
	*apiCore
}

// BucketsList getter for EdgeCenter Storage API
// same result like on UI here https://https://api.edgecenter.ru/storage/bucket/{storageID}
func (sdk *sdkBucket) BucketsList(opts ...func(params *storage.StorageListBucketsHTTPParams)) ([]models.BucketDto, error) {
	params := &storage.StorageListBucketsHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	res, err := sdk.client.Storage.StorageListBucketsHTTP(params, sdk.authWriter)
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}
	list := make([]models.BucketDto, len(res.Payload.Data))
	for i := range res.Payload.Data {
		list[i] = *res.Payload.Data[i]
	}
	return list, nil
}

// CreateBucket writer for EdgeCenter Storage API
func (sdk *sdkBucket) CreateBucket(opts ...func(params *storage.StorageBucketCreateHTTPParams)) error {
	params := &storage.StorageBucketCreateHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	_, err := sdk.client.Storage.StorageBucketCreateHTTP(params, sdk.authWriter)
	if err != nil {
		return fmt.Errorf("request: %w", err)
	}
	return nil
}

// DeleteBucket writer for EdgeCenter Storage API
func (sdk *sdkBucket) DeleteBucket(opts ...func(params *storage.StorageBucketRemoveHTTPParams)) error {
	params := &storage.StorageBucketRemoveHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	_, err := sdk.client.Storage.StorageBucketRemoveHTTP(params, sdk.authWriter)
	if err != nil {
		return fmt.Errorf("request: %w", err)
	}
	return nil
}

// BucketCORS getter for EdgeCenter Storage API
func (sdk *sdkBucket) BucketCORS(opts ...func(params *storage.GetStorageBucketCORSHTTPParams)) (string, error) {
	params := &storage.GetStorageBucketCORSHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	res, err := sdk.client.Storage.GetStorageBucketCORSHTTP(params, sdk.authWriter)
	if err != nil {
		return "", fmt.Errorf("request: %w", err)
	}
	return res.Payload.Data, nil
}

// CreateBucketCORS writer for EdgeCenter Storage API
func (sdk *sdkBucket) CreateBucketCORS(opts ...func(params *storage.StorageBucketCORSCreateHTTPParams)) error {
	params := &storage.StorageBucketCORSCreateHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	_, err := sdk.client.Storage.StorageBucketCORSCreateHTTP(params, sdk.authWriter)
	if err != nil {
		return fmt.Errorf("request: %w", err)
	}
	return nil
}

// CreateBucketLifecycle writer for EdgeCenter Storage API
func (sdk *sdkBucket) CreateBucketLifecycle(opts ...func(params *storage.StorageBucketLifecycleCreateHTTPParams)) error {
	params := &storage.StorageBucketLifecycleCreateHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	_, err := sdk.client.Storage.StorageBucketLifecycleCreateHTTP(params, sdk.authWriter)
	if err != nil {
		return fmt.Errorf("request: %w", err)
	}
	return nil
}

// DeleteBucketLifecycle writer for EdgeCenter Storage API
func (sdk *sdkBucket) DeleteBucketLifecycle(opts ...func(params *storage.StorageBucketLifecycleDeleteHTTPParams)) error {
	params := &storage.StorageBucketLifecycleDeleteHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	_, err := sdk.client.Storage.StorageBucketLifecycleDeleteHTTP(params, sdk.authWriter)
	if err != nil {
		return fmt.Errorf("request: %w", err)
	}
	return nil
}

// CreateBucketPolicy writer for EdgeCenter Storage API
func (sdk *sdkBucket) CreateBucketPolicy(opts ...func(params *storage.StorageBucketPolicyCreateHTTPParams)) error {
	params := &storage.StorageBucketPolicyCreateHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	_, err := sdk.client.Storage.StorageBucketPolicyCreateHTTP(params, sdk.authWriter)
	if err != nil {
		return fmt.Errorf("request: %w", err)
	}
	return nil
}
