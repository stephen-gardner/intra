package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	Endpoint struct {
		req         ftapi.RequestData
		ID          int        `json:"id"`
		URL         string     `json:"url"`
		Description string     `json:"description"`
		CreatedAt   ftapi.Time `json:"created_at"`
		UpdatedAt   ftapi.Time `json:"updated_at"`
		Campus      []Campus   `json:"campus"`
	}
	Endpoints struct {
		req        ftapi.RequestData
		Collection []Endpoint
	}
	EndpointCUParams struct {
		URL         string     `json:"url,omitempty"`
		Secret      string     `json:"secret,omitempty"`
		Description string     `json:"description,omitempty"`
		CreatedAt   ftapi.Time `json:"created_at,omitempty"`
		UpdatedAt   ftapi.Time `json:"updated_at,omitempty"`
	}
)

func (ep *Endpoint) Create(ctx context.Context, params EndpointCUParams) ftapi.CachedRequest {
	ep.req.Endpoint = ftapi.GetEndpoint("endpoints", nil)
	ep.req.ExecuteMethod = func() {
		ep.req.Create(ctx, ep, ftapi.EncapsulatedMarshal("endpoint", params))
	}
	return &ep.req
}

func (ep *Endpoint) Delete(ctx context.Context) ftapi.Request {
	ep.req.Endpoint = ftapi.GetEndpoint("endpoints/"+strconv.Itoa(ep.ID), nil)
	ep.req.ExecuteMethod = func() {
		ep.req.Delete(ctx, ep)
	}
	return &ep.req
}

func (ep *Endpoint) Patch(ctx context.Context, params EndpointCUParams) ftapi.Request {
	ep.req.Endpoint = ftapi.GetEndpoint("endpoints/"+strconv.Itoa(ep.ID), nil)
	ep.req.ExecuteMethod = func() {
		ep.req.Patch(ctx, ftapi.EncapsulatedMarshal("endpoint", params))
	}
	return &ep.req
}

func (ep *Endpoint) Get(ctx context.Context) ftapi.CachedRequest {
	ep.req.Endpoint = ftapi.GetEndpoint("endpoints/"+strconv.Itoa(ep.ID), nil)
	ep.req.ExecuteMethod = func() {
		ep.req.Get(ctx, ep)
	}
	return &ep.req
}

func (eps *Endpoints) GetAll(ctx context.Context) ftapi.CollectionRequest {
	eps.req.Endpoint = ftapi.GetEndpoint("endpoints", nil)
	eps.req.ExecuteMethod = func() {
		eps.req.GetAll(ctx, &eps.Collection)
	}
	return &eps.req
}
