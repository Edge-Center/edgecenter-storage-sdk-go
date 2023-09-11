package ec_storage

import (
	"fmt"
	"github.com/Edge-Center/edgecenter-storage-sdk-go/swagger/client/locations"

	"github.com/Edge-Center/edgecenter-storage-sdk-go/swagger/models"
)

type sdkLocation struct {
	*apiCore
}

// LocationsList getter for EdgeCenter Storage API
func (sdk *sdkLocation) LocationsList(opts ...func(params *locations.LocationListHTTPParams)) ([]models.ClientLocationRes, error) {
	params := &locations.LocationListHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	res, err := sdk.client.Locations.LocationListHTTP(params, sdk.authWriter)
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}
	list := make([]models.ClientLocationRes, len(res.Payload))
	for i := range res.Payload {
		list[i] = *res.Payload[i]
	}
	return list, nil
}
