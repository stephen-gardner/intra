package ftobj

import (
	"context"
	"fmt"
	"intra/ftapi"
	"strconv"
)

type (
	GroupsUser struct {
		req    ftapi.RequestData
		ID     int    `json:"id,omitempty"`
		UserID int    `json:"user_id,omitempty"`
		Group  *Group `json:"group,omitempty"`
	}
	GroupsUsers struct {
		req        ftapi.RequestData
		Collection []GroupsUser
	}
	GroupsUserCUParams struct {
		UserID  int `json:"user_id,omitempty"`
		GroupID int `json:"group_id,omitempty"`
	}
)

func (gu *GroupsUser) Create(ctx context.Context, params GroupsUserCUParams) ftapi.CachedRequest {
	gu.req.Endpoint = ftapi.GetEndpoint("groups_users", nil)
	gu.req.ExecuteMethod = func() {
		gu.req.Create(ctx, gu, ftapi.EncapsulatedMarshal("groups_user", params))
	}
	return &gu.req
}

func (gu *GroupsUser) Delete(ctx context.Context) ftapi.Request {
	gu.req.Endpoint = ftapi.GetEndpoint("groups_users/"+strconv.Itoa(gu.ID), nil)
	gu.req.ExecuteMethod = func() {
		gu.req.Delete(ctx, gu)
	}
	return &gu.req
}

func (gu *GroupsUser) Patch(ctx context.Context, params GroupsUserCUParams) ftapi.Request {
	gu.req.Endpoint = ftapi.GetEndpoint("groups_users/"+strconv.Itoa(gu.ID), nil)
	gu.req.ExecuteMethod = func() {
		gu.req.Patch(ctx, ftapi.EncapsulatedMarshal("groups_user", params))
	}
	return &gu.req
}

func (gu *GroupsUser) Get(ctx context.Context) ftapi.CachedRequest {
	gu.req.Endpoint = ftapi.GetEndpoint("groups_users/"+strconv.Itoa(gu.ID), nil)
	gu.req.ExecuteMethod = func() {
		gu.req.Get(ctx, gu)
	}
	return &gu.req
}

func (gus *GroupsUsers) GetAll(ctx context.Context) ftapi.CollectionRequest {
	gus.req.Endpoint = ftapi.GetEndpoint("groups_users", nil)
	gus.req.ExecuteMethod = func() {
		gus.req.GetAll(ctx, &gus.Collection)
	}
	return &gus.req
}

func (gus *GroupsUsers) GetAllForGroup(ctx context.Context, groupID int) ftapi.CollectionRequest {
	gus.req.Endpoint = ftapi.GetEndpoint(fmt.Sprintf("groups/%d/groups_users", groupID), nil)
	gus.req.ExecuteMethod = func() {
		gus.req.GetAll(ctx, &gus.Collection)
	}
	return &gus.req
}

func (gus *GroupsUsers) GetAllForUser(ctx context.Context, userID int) ftapi.CollectionRequest {
	gus.req.Endpoint = ftapi.GetEndpoint(fmt.Sprintf("users/%d/groups_users", userID), nil)
	gus.req.ExecuteMethod = func() {
		gus.req.GetAll(ctx, &gus.Collection)
	}
	return &gus.req
}
