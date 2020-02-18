package intra

import (
	"context"
	"fmt"
	"net/url"
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

func (location *Location) Get(ctx context.Context, bypassCache bool) error {
	if !bypassCache {
		if loc, present := intraCache.get(catLocations, location.ID); present {
			*location = *loc.(*Location)
			return nil
		}
	}
	locations := &Locations{}
	if err := locations.GetAll(ctx, bypassCache, getSingleParams(location.ID)); err != nil {
		return err
	}
	if len(*locations) == 0 {
		return fmt.Errorf("location %d does not exist", location.ID)
	}
	*location = (*locations)[0]
	return nil
}

func (locations *Locations) GetAll(ctx context.Context, bypassCache bool, params url.Values) error {
	if err := GetAll(GetClient(ctx, "public"), "locations", params, locations); err != nil {
		return err
	}
	if !bypassCache {
		for _, loc := range *locations {
			cached := loc
			intraCache.put(catLocations, cached.ID, &cached)
		}
	}
	return nil
}
