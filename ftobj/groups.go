package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	Group struct {
		req  ftapi.RequestData
		ID   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	}
	Groups struct {
		req        ftapi.RequestData
		Collection []Group
	}
	GroupCUParams struct {
		Name      string `json:"name,omitempty"`
		Color     string `json:"color,omitempty"`
		Important bool   `json:"important,omitempty"`
	}
)

func (group *Group) Create(ctx context.Context, params GroupCUParams) ftapi.CachedRequest {
	group.req.Endpoint = ftapi.GetEndpoint("groups", nil)
	group.req.ExecuteMethod = func() {
		group.req.Create(ctx, group, ftapi.EncapsulatedMarshal("group", params))
	}
	return &group.req
}

func (group *Group) Delete(ctx context.Context) ftapi.Request {
	group.req.Endpoint = ftapi.GetEndpoint("groups/"+strconv.Itoa(group.ID), nil)
	group.req.ExecuteMethod = func() {
		group.req.Delete(ctx, group)
	}
	return &group.req
}

func (group *Group) Patch(ctx context.Context, params GroupCUParams) ftapi.Request {
	group.req.Endpoint = ftapi.GetEndpoint("groups/"+strconv.Itoa(group.ID), nil)
	group.req.ExecuteMethod = func() {
		group.req.Patch(ctx, ftapi.EncapsulatedMarshal("group", params))
	}
	return &group.req
}

func (group *Group) Get(ctx context.Context) ftapi.CachedRequest {
	group.req.Endpoint = ftapi.GetEndpoint("groups/"+strconv.Itoa(group.ID), nil)
	group.req.ExecuteMethod = func() {
		group.req.Get(ctx, group)
	}
	return &group.req
}

func (groups *Groups) GetAll(ctx context.Context) ftapi.CollectionRequest {
	groups.req.Endpoint = ftapi.GetEndpoint("groups", nil)
	groups.req.ExecuteMethod = func() {
		groups.req.GetAll(ctx, &groups.Collection)
	}
	return &groups.req
}
