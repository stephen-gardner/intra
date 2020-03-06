package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	Expertise struct {
		req                ftapi.RequestData
		ID                 int        `json:"id"`
		Name               string     `json:"name"`
		Slug               string     `json:"slug"`
		URL                string     `json:"url"`
		Kind               string     `json:"kind"`
		CreatedAt          ftapi.Time `json:"created_at"`
		ExpertisesUsersURL string     `json:"expertises_users_url"`
	}
	Expertises struct {
		req        ftapi.RequestData
		Collection []Expertise
	}
	ExpertiseCUParams struct {
		Name string `json:"name,omitempty"`
		Slug string `json:"slug,omitempty"`
		Kind string `json:"kind,omitempty"`
	}
)

const (
	ExpertiseKindLanguages  = "languages"
	ExpertiseKindFrameworks = "frameworks"
	ExpertiseKindDatabases  = "databases"
	ExpertiseKindSysAdmin   = "admin_sys"
	ExpertiseKindOther      = "other"
)

func (exp *Expertise) Create(ctx context.Context, params ExpertiseCUParams) ftapi.CachedRequest {
	exp.req.Endpoint = ftapi.GetEndpoint("expertises", nil)
	exp.req.ExecuteMethod = func() {
		exp.req.Create(ctx, exp, ftapi.EncapsulatedMarshal("expertise", params))
	}
	return &exp.req
}

func (exp *Expertise) Delete(ctx context.Context) ftapi.Request {
	exp.req.Endpoint = ftapi.GetEndpoint("expertises/"+strconv.Itoa(exp.ID), nil)
	exp.req.ExecuteMethod = func() {
		exp.req.Delete(ctx, exp)
	}
	return &exp.req
}

func (exp *Expertise) Patch(ctx context.Context, params ExpertiseCUParams) ftapi.Request {
	exp.req.Endpoint = ftapi.GetEndpoint("expertises/"+strconv.Itoa(exp.ID), nil)
	exp.req.ExecuteMethod = func() {
		exp.req.Patch(ctx, ftapi.EncapsulatedMarshal("expertise", params))
	}
	return &exp.req
}

func (exp *Expertise) Get(ctx context.Context) ftapi.CachedRequest {
	exp.req.Endpoint = ftapi.GetEndpoint("expertises/"+strconv.Itoa(exp.ID), nil)
	exp.req.ExecuteMethod = func() {
		exp.req.Get(ctx, exp)
	}
	return &exp.req
}

func (exps *Expertises) GetAll(ctx context.Context) ftapi.CollectionRequest {
	exps.req.Endpoint = ftapi.GetEndpoint("expertises", nil)
	exps.req.ExecuteMethod = func() {
		exps.req.GetAll(ctx, &exps.Collection)
	}
	return &exps.req
}
