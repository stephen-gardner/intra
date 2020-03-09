package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	TitlesUser struct {
		req      ftapi.RequestData
		ID       int  `json:"id,omitempty"`
		UserID   int  `json:"user_id,omitempty"`
		TitleID  int  `json:"title_id,omitempty"`
		Selected bool `json:"selected,omitempty"`
	}
	TitlesUsers struct {
		req        ftapi.RequestData
		Collection []TitlesUser
	}
	TitlesUserCUParams struct {
		UserID  int `json:"user_id,omitempty"`
		TitleID int `json:"title_id,omitempty"`
	}
)

func (tu *TitlesUser) Create(ctx context.Context, params TitleCUParams) ftapi.CachedRequest {
	tu.req.Endpoint = ftapi.GetEndpoint("titles_users", nil)
	tu.req.ExecuteMethod = func() {
		tu.req.Create(ctx, tu, ftapi.EncapsulatedMarshal("titles_user", params))
	}
	return &tu.req
}

func (tu *TitlesUser) Delete(ctx context.Context) ftapi.Request {
	tu.req.Endpoint = ftapi.GetEndpoint("titles_users/"+strconv.Itoa(tu.ID), nil)
	tu.req.ExecuteMethod = func() {
		tu.req.Delete(ctx, tu)
	}
	return &tu.req
}

func (tu *TitlesUser) Patch(ctx context.Context, params TitleCUParams) ftapi.Request {
	tu.req.Endpoint = ftapi.GetEndpoint("titles_users/"+strconv.Itoa(tu.ID), nil)
	tu.req.ExecuteMethod = func() {
		tu.req.Patch(ctx, ftapi.EncapsulatedMarshal("titles_user", params))
	}
	return &tu.req
}

func (tu *TitlesUser) Get(ctx context.Context) ftapi.CachedRequest {
	tu.req.Endpoint = ftapi.GetEndpoint("titles_users/"+strconv.Itoa(tu.ID), nil)
	tu.req.ExecuteMethod = func() {
		tu.req.Get(ctx, tu)
	}
	return &tu.req
}

func (tus *TitlesUsers) GetAll(ctx context.Context) ftapi.CollectionRequest {
	tus.req.Endpoint = ftapi.GetEndpoint("titles_users", nil)
	tus.req.ExecuteMethod = func() {
		tus.req.GetAll(ctx, &tus.Collection)
	}
	return &tus.req
}
