package ec_storage

import (
	"fmt"

	"github.com/Edge-Center/edgecenter-storage-sdk-go/swagger/client/buckets"
	"github.com/Edge-Center/edgecenter-storage-sdk-go/swagger/models"
)

type sdkBucket struct {
	*apiCore
}

// BucketsList getter for EdgeCenter Storage API
// same result like on UI here https://api.edgecenter.ru/storage/bucket/{storageID}
func (sdk *sdkBucket) BucketsList(opts ...func(params *buckets.StorageListBucketsHTTPParams)) ([]models.BucketDto, error) {
	params := &buckets.StorageListBucketsHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	res, err := sdk.client.Buckets.StorageListBucketsHTTP(params, sdk.authWriter)
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
func (sdk *sdkBucket) CreateBucket(opts ...func(params *buckets.StorageBucketCreateHTTPParams)) error {
	params := &buckets.StorageBucketCreateHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	_, err := sdk.client.Buckets.StorageBucketCreateHTTP(params, sdk.authWriter)
	if err != nil {
		return fmt.Errorf("request: %w", err)
	}
	return nil
}

// DeleteBucket writer for EdgeCenter Storage API
func (sdk *sdkBucket) DeleteBucket(opts ...func(params *buckets.StorageBucketRemoveHTTPParams)) error {
	params := &buckets.StorageBucketRemoveHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	_, err := sdk.client.Buckets.StorageBucketRemoveHTTP(params, sdk.authWriter)
	if err != nil {
		return fmt.Errorf("request: %w", err)
	}
	return nil
}

// BucketCORS getter for EdgeCenter Storage API
func (sdk *sdkBucket) BucketCORS(opts ...func(params *buckets.GetStorageBucketCORSHTTPParams)) (string, error) {
	params := &buckets.GetStorageBucketCORSHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	res, err := sdk.client.Buckets.GetStorageBucketCORSHTTP(params, sdk.authWriter)
	if err != nil {
		return "", fmt.Errorf("request: %w", err)
	}
	return res.Payload.Data, nil
}

// CreateBucketCORS writer for EdgeCenter Storage API
func (sdk *sdkBucket) CreateBucketCORS(opts ...func(params *buckets.StorageBucketCORSCreateHTTPParams)) error {
	params := &buckets.StorageBucketCORSCreateHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	_, err := sdk.client.Buckets.StorageBucketCORSCreateHTTP(params, sdk.authWriter)
	if err != nil {
		return fmt.Errorf("request: %w", err)
	}
	return nil
}

// CreateBucketLifecycle writer for EdgeCenter Storage API
func (sdk *sdkBucket) CreateBucketLifecycle(opts ...func(params *buckets.StorageBucketLifecycleCreateHTTPParams)) error {
	params := &buckets.StorageBucketLifecycleCreateHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	_, err := sdk.client.Buckets.StorageBucketLifecycleCreateHTTP(params, sdk.authWriter)
	if err != nil {
		return fmt.Errorf("request: %w", err)
	}
	return nil
}

// DeleteBucketLifecycle writer for EdgeCenter Storage API
func (sdk *sdkBucket) DeleteBucketLifecycle(opts ...func(params *buckets.StorageBucketLifecycleDeleteHTTPParams)) error {
	params := &buckets.StorageBucketLifecycleDeleteHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	_, err := sdk.client.Buckets.StorageBucketLifecycleDeleteHTTP(params, sdk.authWriter)
	if err != nil {
		return fmt.Errorf("request: %w", err)
	}
	return nil
}

// CreateBucketPolicy writer for EdgeCenter Storage API
func (sdk *sdkBucket) CreateBucketPolicy(opts ...func(params *buckets.StorageBucketPolicyCreateHTTPParams)) error {
	params := &buckets.StorageBucketPolicyCreateHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	_, err := sdk.client.Buckets.StorageBucketPolicyCreateHTTP(params, sdk.authWriter)
	if err != nil {
		return fmt.Errorf("request: %w", err)
	}
	return nil
}
