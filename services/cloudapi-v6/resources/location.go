package resources

import (
	"context"

	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

type Location struct {
	ionoscloud.Location
}

type CpuArchitectureProperties struct {
	ionoscloud.CpuArchitectureProperties
}

type Locations struct {
	ionoscloud.Locations
}

// LocationsService is a wrapper around ionoscloud.Location
type LocationsService interface {
	List() (Locations, *Response, error)
	GetByRegionAndLocationId(regionId, locationId string) (*Location, *Response, error)
	GetByRegionId(regionId string) (Locations, *Response, error)
}

type locationsService struct {
	client  *Client
	context context.Context
}

var _ LocationsService = &locationsService{}

func NewLocationService(client *Client, ctx context.Context) LocationsService {
	return &locationsService{
		client:  client,
		context: ctx,
	}
}

func (s *locationsService) List() (Locations, *Response, error) {
	req := s.client.LocationsApi.LocationsGet(s.context)
	locations, resp, err := s.client.LocationsApi.LocationsGetExecute(req)
	return Locations{locations}, &Response{*resp}, err
}

func (s *locationsService) GetByRegionAndLocationId(regionId, locationId string) (*Location, *Response, error) {
	req := s.client.LocationsApi.LocationsFindByRegionIdAndId(s.context, regionId, locationId)
	loc, resp, err := s.client.LocationsApi.LocationsFindByRegionIdAndIdExecute(req)
	return &Location{loc}, &Response{*resp}, err
}

func (s *locationsService) GetByRegionId(regionId string) (Locations, *Response, error) {
	req := s.client.LocationsApi.LocationsFindByRegionId(s.context, regionId)
	locs, resp, err := s.client.LocationsApi.LocationsFindByRegionIdExecute(req)
	return Locations{locs}, &Response{*resp}, err
}