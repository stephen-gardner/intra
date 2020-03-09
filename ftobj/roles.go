package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	Role struct {
		req         ftapi.RequestData
		ID          int    `json:"id,omitempty"`
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	}
	Roles struct {
		req        ftapi.RequestData
		Collection []Role
	}
	RoleCUParams struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	}
)

func (role *Role) Create(ctx context.Context, params RoleCUParams) ftapi.CachedRequest {
	role.req.Endpoint = ftapi.GetEndpoint("roles", nil)
	role.req.ExecuteMethod = func() {
		role.req.Create(ctx, role, ftapi.EncapsulatedMarshal("role", params))
	}
	return &role.req
}

func (role *Role) Delete(ctx context.Context) ftapi.Request {
	role.req.Endpoint = ftapi.GetEndpoint("roles/"+strconv.Itoa(role.ID), nil)
	role.req.ExecuteMethod = func() {
		role.req.Delete(ctx, role)
	}
	return &role.req
}

func (role *Role) Patch(ctx context.Context, params RoleCUParams) ftapi.Request {
	role.req.Endpoint = ftapi.GetEndpoint("roles/"+strconv.Itoa(role.ID), nil)
	role.req.ExecuteMethod = func() {
		role.req.Patch(ctx, ftapi.EncapsulatedMarshal("role", params))
	}
	return &role.req
}

func (role *Role) Get(ctx context.Context) ftapi.CachedRequest {
	role.req.Endpoint = ftapi.GetEndpoint("roles/"+strconv.Itoa(role.ID), nil)
	role.req.ExecuteMethod = func() {
		role.req.Get(ctx, role)
	}
	return &role.req
}

func (roles *Roles) GetAll(ctx context.Context) ftapi.CollectionRequest {
	roles.req.Endpoint = ftapi.GetEndpoint("roles", nil)
	roles.req.ExecuteMethod = func() {
		roles.req.GetAll(ctx, &roles.Collection)
	}
	return &roles.req
}
