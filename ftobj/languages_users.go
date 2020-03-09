package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	LanguagesUser struct {
		req        ftapi.RequestData
		ID         int        `json:"id,omitempty"`
		LanguageID int        `json:"language_id,omitempty"`
		UserID     int        `json:"user_id,omitempty"`
		Position   int        `json:"position,omitempty"`
		CreatedAt  ftapi.Time `json:"created_at,omitempty"`
	}
	LanguagesUsers struct {
		req        ftapi.RequestData
		Collection []LanguagesUser
	}
	LanguagesUserCUParams struct {
		UserID     int `json:"user_id,omitempty"`
		LanguageID int `json:"language_id,omitempty"`
		Position   int `json:"position,omitempty"`
	}
)

func (lu *LanguagesUser) Create(ctx context.Context, params LanguagesUserCUParams) ftapi.CachedRequest {
	lu.req.Endpoint = ftapi.GetEndpoint("languages_users", nil)
	lu.req.ExecuteMethod = func() {
		lu.req.Create(ctx, lu, ftapi.EncapsulatedMarshal("languages_user", params))
	}
	return &lu.req
}

func (lu *LanguagesUser) Delete(ctx context.Context) ftapi.Request {
	lu.req.Endpoint = ftapi.GetEndpoint("languages_users/"+strconv.Itoa(lu.ID), nil)
	lu.req.ExecuteMethod = func() {
		lu.req.Delete(ctx, lu)
	}
	return &lu.req
}

func (lu *LanguagesUser) Patch(ctx context.Context, params LanguagesUserCUParams) ftapi.Request {
	lu.req.Endpoint = ftapi.GetEndpoint("languages_users/"+strconv.Itoa(lu.ID), nil)
	lu.req.ExecuteMethod = func() {
		lu.req.Patch(ctx, ftapi.EncapsulatedMarshal("languages_user", params))
	}
	return &lu.req
}

func (lu *LanguagesUser) Get(ctx context.Context) ftapi.CachedRequest {
	lu.req.Endpoint = ftapi.GetEndpoint("languages_users/"+strconv.Itoa(lu.ID), nil)
	lu.req.ExecuteMethod = func() {
		lu.req.Get(ctx, lu)
	}
	return &lu.req
}

func (lus *LanguagesUsers) GetAll(ctx context.Context) ftapi.CollectionRequest {
	lus.req.Endpoint = ftapi.GetEndpoint("languages_users", nil)
	lus.req.ExecuteMethod = func() {
		lus.req.GetAll(ctx, &lus.Collection)
	}
	return &lus.req
}
