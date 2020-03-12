package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	Location struct {
		req      ftapi.RequestData
		ID       int         `json:"id,omitempty"`
		BeginAt  *ftapi.Time `json:"begin_at,omitempty"`
		EndAt    *ftapi.Time `json:"end_at,omitempty"`
		Host     string      `json:"host,omitempty"`
		CampusID int         `json:"campus_id,omitempty"`
		Primary  bool        `json:"primary,omitempty"`
		User     *User       `json:"user,omitempty"`
	}
	Locations struct {
		req        ftapi.RequestData
		Collection []Location
	}
	LocationCUParams struct {
		UserID   int         `json:"user_id,omitempty"`
		BeginAt  *ftapi.Time `json:"begin_at,omitempty"`
		EndAt    *ftapi.Time `json:"end_at,omitempty"`
		Primary  bool        `json:"primary,omitempty"`
		Host     string      `json:"host,omitempty"`
		CampusID int         `json:"campus_id,omitempty"`
	}
)

func (ps *Location) Create(ctx context.Context, params LocationCUParams) ftapi.CachedRequest {
	ps.req.Endpoint = ftapi.GetEndpoint("locations", nil)
	ps.req.ExecuteMethod = func() {
		ps.req.Create(ctx, ps, ftapi.EncapsulatedMarshal("location", params))
	}
	return &ps.req
}

func (ps *Location) Delete(ctx context.Context) ftapi.Request {
	ps.req.Endpoint = ftapi.GetEndpoint("locations/"+strconv.Itoa(ps.ID), nil)
	ps.req.ExecuteMethod = func() {
		ps.req.Delete(ctx, ps)
	}
	return &ps.req
}

func (ps *Location) Patch(ctx context.Context, params LocationCUParams) ftapi.Request {
	ps.req.Endpoint = ftapi.GetEndpoint("locations/"+strconv.Itoa(ps.ID), nil)
	ps.req.ExecuteMethod = func() {
		ps.req.Patch(ctx, ftapi.EncapsulatedMarshal("location", params))
	}
	return &ps.req
}

func (ps *Location) Get(ctx context.Context) ftapi.CachedRequest {
	ps.req.Endpoint = ftapi.GetEndpoint("locations/"+strconv.Itoa(ps.ID), nil)
	ps.req.ExecuteMethod = func() {
		ps.req.Get(ctx, ps)
	}
	return &ps.req
}

func (locs *Locations) GetAll(ctx context.Context) ftapi.CollectionRequest {
	locs.req.Endpoint = ftapi.GetEndpoint("locations", nil)
	locs.req.ExecuteMethod = func() {
		locs.req.GetAll(ctx, &locs.Collection)
	}
	return &locs.req
}
