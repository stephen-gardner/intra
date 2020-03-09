package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	Title struct {
		req  ftapi.RequestData
		ID   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	}
	Titles struct {
		req        ftapi.RequestData
		Collection []Title
	}
	TitleCUParams struct {
		Name string `json:"name, omitempty"`
	}
)

func (title *Title) Create(ctx context.Context, params TitleCUParams) ftapi.CachedRequest {
	title.req.Endpoint = ftapi.GetEndpoint("titles", nil)
	title.req.ExecuteMethod = func() {
		title.req.Create(ctx, title, ftapi.EncapsulatedMarshal("title", params))
	}
	return &title.req
}

func (title *Title) Delete(ctx context.Context) ftapi.Request {
	title.req.Endpoint = ftapi.GetEndpoint("titles/"+strconv.Itoa(title.ID), nil)
	title.req.ExecuteMethod = func() {
		title.req.Delete(ctx, title)
	}
	return &title.req
}

func (title *Title) Patch(ctx context.Context, params TitleCUParams) ftapi.Request {
	title.req.Endpoint = ftapi.GetEndpoint("titles/"+strconv.Itoa(title.ID), nil)
	title.req.ExecuteMethod = func() {
		title.req.Patch(ctx, ftapi.EncapsulatedMarshal("title", params))
	}
	return &title.req
}

func (title *Title) Get(ctx context.Context) ftapi.CachedRequest {
	title.req.Endpoint = ftapi.GetEndpoint("titles/"+strconv.Itoa(title.ID), nil)
	title.req.ExecuteMethod = func() {
		title.req.Get(ctx, title)
	}
	return &title.req
}

func (titles *Titles) GetAll(ctx context.Context) ftapi.CollectionRequest {
	titles.req.Endpoint = ftapi.GetEndpoint("titles", nil)
	titles.req.ExecuteMethod = func() {
		titles.req.GetAll(ctx, &titles.Collection)
	}
	return &titles.req
}
