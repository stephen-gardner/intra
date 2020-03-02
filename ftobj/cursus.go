package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	Cursus struct {
		req       ftapi.RequestData
		ID        int        `json:"id"`
		CreatedAt ftapi.Time `json:"created_at"`
		Name      string     `json:"name"`
		Slug      string     `json:"slug"`
	}
	Cursuses struct {
		req        ftapi.RequestData
		Collection []Cursus
	}
	CursusCUParams struct {
		Cursus struct {
			Name     string `json:"name,omitempty"`
			Kind     string `json:"kind,omitempty"`
			SkillIDs []int  `json:"skill_ids,omitempty"`
		} `json:"cursus,omitempty"`
	}
)

const (
	CursusKindNormal   = "normal"
	CursusKindPiscine  = "piscine"
	CursusKindExternal = "external"
)

func (exp *Cursus) Create(ctx context.Context, params CursusCUParams) ftapi.CachedRequest {
	exp.req.Endpoint = ftapi.GetEndpoint("cursus", nil)
	exp.req.ExecuteMethod = func() {
		exp.req.Create(ftapi.GetClient(ctx, "public"), exp, params)
	}
	return &exp.req
}

func (exp *Cursus) Delete(ctx context.Context) ftapi.Request {
	exp.req.Endpoint = ftapi.GetEndpoint("cursus/"+strconv.Itoa(exp.ID), nil)
	exp.req.ExecuteMethod = func() {
		exp.req.Delete(ftapi.GetClient(ctx, "public"), exp)
	}
	return &exp.req
}

func (exp *Cursus) Patch(ctx context.Context, params CursusCUParams) ftapi.Request {
	exp.req.Endpoint = ftapi.GetEndpoint("cursus/"+strconv.Itoa(exp.ID), nil)
	exp.req.ExecuteMethod = func() {
		exp.req.Patch(ftapi.GetClient(ctx, "public"), exp, params)
	}
	return &exp.req
}

func (exp *Cursus) Get(ctx context.Context) ftapi.CachedRequest {
	exp.req.Endpoint = ftapi.GetEndpoint("cursus/"+strconv.Itoa(exp.ID), nil)
	exp.req.ExecuteMethod = func() {
		exp.req.Get(ftapi.GetClient(ctx, "public"), exp)
	}
	return &exp.req
}

func (exps *Cursuses) GetAll(ctx context.Context) ftapi.CollectionRequest {
	exps.req.Endpoint = ftapi.GetEndpoint("cursus", nil)
	exps.req.ExecuteMethod = func() {
		exps.req.GetAll(ftapi.GetClient(ctx, "public"), &exps.Collection)
	}
	return &exps.req
}