package ec_storage

import (
	"context"
	"fmt"
	"github.com/Edge-Center/edgecenter-storage-sdk-go/swagger/client/locations"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/Edge-Center/edgecenter-storage-sdk-go/swagger/client/buckets"
	"github.com/Edge-Center/edgecenter-storage-sdk-go/swagger/client/storages"
)

func setupTest(t *testing.T) *SDK {
	t.Helper()
	apiUrl := strings.TrimSpace(os.Getenv("TESTS_API_URL"))
	apiPath := strings.TrimSpace(os.Getenv("TESTS_API_PATH"))
	apiToken := strings.TrimSpace(os.Getenv("TESTS_API_PERMANENT_TOKEN"))
	bearerToken := strings.TrimSpace(os.Getenv("TESTS_BEARER_TOKEN"))
	if apiUrl == "" {
		t.Skip("no defined TESTS_API_URL")
	}
	switch {
	case apiToken != "":
		return NewSDK(apiUrl, apiPath,
			WithPermanentTokenAuth(func() string { return apiToken }),
			WithUserAgent("sdk-test"),
		)
	case bearerToken != "":
		return NewSDK(apiUrl, apiPath,
			WithBearerAuth(func() string { return bearerToken }),
			WithUserAgent("sdk-test"),
		)
	default:
		t.Skip("No defined TESTS_API_PERMANENT_TOKEN or TESTS_BEARER_TOKEN")
	}
	return nil
}

func TestS3Storage(t *testing.T) {
	sdk := setupTest(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	sold := time.Now().Unix()

	// create storage

	stName := fmt.Sprintf("sdk-test-st-s3-%d", sold)
	csOpts := []func(p *storages.StorageCreateHTTPParams){
		func(p *storages.StorageCreateHTTPParams) { p.Context = ctx },
		func(p *storages.StorageCreateHTTPParams) {
			p.Body.Type = "s3"
			p.Body.Name = stName
			p.Body.Location = "s-dt4"
		},
	}
	csResp, err := sdk.CreateStorage(csOpts...)
	if err != nil {
		t.Fatal("create storage s3", err)
	}
	stID := csResp.ID

	// create bucket

	bckName := fmt.Sprintf("sdk-test-bucket-s3-%d", sold)
	cbOpts := []func(p *buckets.StorageBucketCreateHTTPParams){
		func(p *buckets.StorageBucketCreateHTTPParams) { p.Context = ctx },
		func(p *buckets.StorageBucketCreateHTTPParams) {
			p.ID = stID
			p.Name = bckName
		},
	}
	err = sdk.CreateBucket(cbOpts...)
	if err != nil {
		t.Fatal("create bucket", err)
	}

	// delete bucket

	dbOpts := []func(p *buckets.StorageBucketRemoveHTTPParams){
		func(p *buckets.StorageBucketRemoveHTTPParams) { p.Context = ctx },
		func(p *buckets.StorageBucketRemoveHTTPParams) {
			p.ID = stID
			p.Name = bckName
		},
	}
	err = sdk.DeleteBucket(dbOpts...)
	if err != nil {
		t.Fatal("delete bucket", err)
	}

	// delete storage

	dsOpts := []func(p *storages.StorageDeleteHTTPParams){
		func(p *storages.StorageDeleteHTTPParams) {
			p.Context = ctx
			p.ID = stID
		},
	}
	err = sdk.DeleteStorage(dsOpts...)
	if err != nil {
		t.Fatal("delete storage s3", err)
	}
}

func TestLocation(t *testing.T) {
	sdk := setupTest(t)
	t.Log()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	llOpts := []func(l *locations.LocationListHTTPParams){
		func(l *locations.LocationListHTTPParams) {
			l.Context = ctx
		},
	}
	locationsList, err := sdk.LocationsList(llOpts...)
	t.Log(locationsList)
	if err != nil {
		t.Fatal("locations list", err)
	}

}
