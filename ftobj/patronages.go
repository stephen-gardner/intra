package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	Patronage struct {
		req         ftapi.RequestData
		ID          int        `json:"id,omitempty"`
		UserID      int        `json:"user_id,omitempty"`
		GodfatherID int        `json:"godfather_id,omitempty"`
		Ongoing     bool       `json:"ongoing,omitempty"`
		CreatedAt   ftapi.Time `json:"created_at,omitempty"`
		UpdatedAt   ftapi.Time `json:"updated_at,omitempty"`
		User        struct {
			ID    int    `json:"id,omitempty"`
			Login string `json:"login,omitempty"`
			URL   string `json:"url,omitempty"`
		} `json:"user,omitempty"`
		Godfather struct {
			ID    int    `json:"id,omitempty"`
			Login string `json:"login,omitempty"`
			URL   string `json:"url,omitempty"`
		} `json:"godfather,omitempty"`
	}
	Patronages struct {
		req        ftapi.RequestData
		Collection []Patronage
	}
	PatronageCUParams struct {
		UserID      int  `json:"user_id,omitempty"`
		GodfatherID int  `json:"godfather_id,omitempty"`
		Ongoing     bool `json:"ongoing,omitempty"`
	}
)

func (p *Patronage) Create(ctx context.Context, params PatronageCUParams) ftapi.CachedRequest {
	p.req.Endpoint = ftapi.GetEndpoint("patronages", nil)
	p.req.ExecuteMethod = func() {
		p.req.Create(ctx, p, ftapi.EncapsulatedMarshal("patronage", params))
	}
	return &p.req
}

func (p *Patronage) Delete(ctx context.Context) ftapi.Request {
	p.req.Endpoint = ftapi.GetEndpoint("patronages/"+strconv.Itoa(p.ID), nil)
	p.req.ExecuteMethod = func() {
		p.req.Delete(ctx, p)
	}
	return &p.req
}

func (p *Patronage) Patch(ctx context.Context, params PatronageCUParams) ftapi.Request {
	p.req.Endpoint = ftapi.GetEndpoint("patronages/"+strconv.Itoa(p.ID), nil)
	p.req.ExecuteMethod = func() {
		p.req.Patch(ctx, ftapi.EncapsulatedMarshal("patronage", params))
	}
	return &p.req
}

func (p *Patronage) Get(ctx context.Context) ftapi.CachedRequest {
	p.req.Endpoint = ftapi.GetEndpoint("patronages/"+strconv.Itoa(p.ID), nil)
	p.req.ExecuteMethod = func() {
		p.req.Get(ctx, p)
	}
	return &p.req
}

func (ps *Patronages) GetAll(ctx context.Context) ftapi.CollectionRequest {
	ps.req.Endpoint = ftapi.GetEndpoint("patronages", nil)
	ps.req.ExecuteMethod = func() {
		ps.req.GetAll(ctx, &ps.Collection)
	}
	return &ps.req
}
