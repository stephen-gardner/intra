package intra

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type (
	Location struct {
		ID       int       `json:"id"`
		BeginAt  time.Time `json:"begin_at"`
		EndAt    time.Time `json:"end_at"`
		Host     string    `json:"host"`
		CampusID int       `json:"campus_id"`
		Primary  bool      `json:"primary"`
		User     struct {
			ID    int    `json:"id"`
			Login string `json:"login"`
			URL   string `json:"url"`
		} `json:"user"`
	}
	Locations []Location
)

func (location *Location) GetLocation(ctx context.Context, bypassCache bool, ID int) error {
	IDStr := strconv.Itoa(ID)
	endpoint := GetEndpoint("locations/"+IDStr, nil)
	if !bypassCache {
		if loc, present := intraCache.get(endpoint); present {
			*location = loc.(Location)
			return nil
		}
	}
	locations := &Locations{}
	if err := locations.GetAllLocations(ctx, bypassCache, getSingleParams(IDStr)); err != nil {
		return err
	}
	if len(*locations) == 0 {
		return fmt.Errorf("location %d does not exist", ID)
	}
	*location = (*locations)[0]
	return nil
}

func (locations *Locations) GetAllLocations(ctx context.Context, bypassCache bool, params url.Values) error {
	if err := GetAll(GetClient(ctx, "public"), "locations", params, locations); err != nil {
		return err
	}
	if !bypassCache {
		for _, loc := range *locations {
			endpoint := GetEndpoint("locations/"+strconv.Itoa(loc.ID), nil)
			intraCache.put(endpoint, loc)
		}
	}
	return nil
}
