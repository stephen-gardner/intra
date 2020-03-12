package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
	"time"
)

type (
	TeamsUser struct {
		req        ftapi.RequestData
		ID         int        `json:"id,omitempty"`
		TeamID     int        `json:"team_id,omitempty"`
		UserID     int        `json:"user_id,omitempty"`
		CreatedAt  *time.Time `json:"created_at,omitempty"`
		Validated  bool       `json:"validated,omitempty"`
		Leader     bool       `json:"leader,omitempty"`
		Occurrence int        `json:"occurrence,omitempty"`
		Team       *Team      `json:"team,omitempty"`
		User       *User      `json:"user,omitempty"`
	}
	TeamsUsers struct {
		req        ftapi.RequestData
		Collection []TeamsUser
	}
	TeamsUserCUParams struct {
		TeamID     int  `json:"team_id,omitempty"`
		UserID     int  `json:"user_id,omitempty"`
		Leader     bool `json:"leader,omitempty"`
		Validated  bool `json:"validated,omitempty"`
		Occurrence int  `json:"occurrence,omitempty"`
	}
)

func (tu *TeamsUser) Create(ctx context.Context, params TeamsUserCUParams) ftapi.CachedRequest {
	tu.req.Endpoint = ftapi.GetEndpoint("teams_users", nil)
	tu.req.ExecuteMethod = func() {
		tu.req.Create(ctx, tu, ftapi.EncapsulatedMarshal("teams_user", params))
	}
	return &tu.req
}

func (tu *TeamsUser) Delete(ctx context.Context) ftapi.Request {
	tu.req.Endpoint = ftapi.GetEndpoint("teams_users/"+strconv.Itoa(tu.ID), nil)
	tu.req.ExecuteMethod = func() {
		tu.req.Delete(ctx, tu)
	}
	return &tu.req
}

func (tu *TeamsUser) Patch(ctx context.Context, params TeamsUserCUParams) ftapi.Request {
	tu.req.Endpoint = ftapi.GetEndpoint("teams_users/"+strconv.Itoa(tu.ID), nil)
	tu.req.ExecuteMethod = func() {
		tu.req.Patch(ctx, ftapi.EncapsulatedMarshal("teams_user", params))
	}
	return &tu.req
}

func (tu *TeamsUser) Get(ctx context.Context) ftapi.CachedRequest {
	tu.req.Endpoint = ftapi.GetEndpoint("teams_users/"+strconv.Itoa(tu.ID), nil)
	tu.req.ExecuteMethod = func() {
		tu.req.Get(ctx, tu)
	}
	return &tu.req
}

func (tus *TeamsUsers) GetAll(ctx context.Context) ftapi.CollectionRequest {
	tus.req.Endpoint = ftapi.GetEndpoint("teams_users", nil)
	tus.req.ExecuteMethod = func() {
		tus.req.GetAll(ctx, &tus.Collection)
	}
	return &tus.req
}
