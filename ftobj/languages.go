package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	Language struct {
		req        ftapi.RequestData
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Identifier string `json:"identifier"`
	}
	Languages struct {
		req        ftapi.RequestData
		Collection []Language
	}
	LanguageCUParams struct {
		Name       string `json:"name,omitempty"`
		Identifier string `json:"identifier,omitempty"`
	}
)

func (lang *Language) Create(ctx context.Context, params LanguageCUParams) ftapi.CachedRequest {
	lang.req.Endpoint = ftapi.GetEndpoint("languages", nil)
	lang.req.ExecuteMethod = func() {
		lang.req.Create(ctx, lang, ftapi.EncapsulatedMarshal("language", params))
	}
	return &lang.req
}

func (lang *Language) Delete(ctx context.Context) ftapi.Request {
	lang.req.Endpoint = ftapi.GetEndpoint("languages/"+strconv.Itoa(lang.ID), nil)
	lang.req.ExecuteMethod = func() {
		lang.req.Delete(ctx, lang)
	}
	return &lang.req
}

func (lang *Language) Patch(ctx context.Context, params LanguageCUParams) ftapi.Request {
	lang.req.Endpoint = ftapi.GetEndpoint("languages/"+strconv.Itoa(lang.ID), nil)
	lang.req.ExecuteMethod = func() {
		lang.req.Patch(ctx, ftapi.EncapsulatedMarshal("language", params))
	}
	return &lang.req
}

func (lang *Language) Get(ctx context.Context) ftapi.CachedRequest {
	lang.req.Endpoint = ftapi.GetEndpoint("languages/"+strconv.Itoa(lang.ID), nil)
	lang.req.ExecuteMethod = func() {
		lang.req.Get(ctx, lang)
	}
	return &lang.req
}

func (languages *Languages) GetAll(ctx context.Context) ftapi.CollectionRequest {
	languages.req.Endpoint = ftapi.GetEndpoint("languages", nil)
	languages.req.ExecuteMethod = func() {
		languages.req.GetAll(ctx, &languages.Collection)
	}
	return &languages.req
}
