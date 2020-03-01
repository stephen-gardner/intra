package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	Location struct {
		req      ftapi.RequestData
		ID       int        `json:"id"`
		BeginAt  ftapi.Time `json:"begin_at"`
		EndAt    ftapi.Time `json:"end_at"`
		Host     string     `json:"host"`
		CampusID int        `json:"campus_id"`
		Primary  bool       `json:"primary"`
		User     struct {
			ID    int    `json:"id"`
			Login string `json:"login"`
			URL   string `json:"url"`
		} `json:"user"`
	}
	Locations struct {
		req        ftapi.RequestData
		Collection []Location
	}
	LocationCUParams struct {
		Location struct {
			UserID   int        `json:"user_id,omitempty"`
			BeginAt  ftapi.Time `json:"begin_at,omitempty"`
			EndAt    ftapi.Time `json:"end_at,omitempty"`
			Primary  bool       `json:"primary,omitempty"`
			Host     string     `json:"host,omitempty"`
			CampusID int        `json:"campus_id,omitempty"`
		} `json:"location,omitempty"`
	}
)

func (ps *Location) Create(ctx context.Context, params LocationCUParams) ftapi.CachedRequest {
	ps.req.Endpoint = ftapi.GetEndpoint("locations", nil)
	ps.req.ExecuteMethod = func() {
		ps.req.Create(ftapi.GetClient(ctx, "public"), ps, params)
	}
	return &ps.req
}

func (ps *Location) Delete(ctx context.Context) ftapi.Request {
	ps.req.Endpoint = ftapi.GetEndpoint("locations/"+strconv.Itoa(ps.ID), nil)
	ps.req.ExecuteMethod = func() {
		ps.req.Delete(ftapi.GetClient(ctx, "public"), ps)
	}
	return &ps.req
}

func (ps *Location) Patch(ctx context.Context, params LocationCUParams) ftapi.Request {
	ps.req.Endpoint = ftapi.GetEndpoint("locations/"+strconv.Itoa(ps.ID), nil)
	ps.req.ExecuteMethod = func() {
		ps.req.Patch(ftapi.GetClient(ctx, "public"), ps, params)
	}
	return &ps.req
}

func (ps *Location) Get(ctx context.Context) ftapi.CachedRequest {
	ps.req.Endpoint = ftapi.GetEndpoint("locations/"+strconv.Itoa(ps.ID), nil)
	ps.req.ExecuteMethod = func() {
		ps.req.Get(ftapi.GetClient(ctx, "public"), ps)
	}
	return &ps.req
}

func (locs *Locations) GetAll(ctx context.Context) ftapi.CollectionRequest {
	locs.req.Endpoint = ftapi.GetEndpoint("locations", nil)
	locs.req.ExecuteMethod = func() {
		locs.req.GetAll(ftapi.GetClient(ctx, "public"), &locs.Collection)
	}
	return &locs.req
}
