package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	ExpertisesUser struct {
		req         ftapi.RequestData
		ID          int        `json:"id"`
		ExpertiseID int        `json:"expertise_id"`
		Interested  bool       `json:"interested"`
		Value       int        `json:"value"`
		ContactMe   bool       `json:"contact_me"`
		CreatedAt   ftapi.Time `json:"created_at"`
		UserID      int        `json:"user_id"`
		Expertise   Expertise  `json:"expertise"`
		User        struct {
			ID    int    `json:"id"`
			Login string `json:"login"`
			URL   string `json:"url"`
		} `json:"user"`
	}
	ExpertisesUsers struct {
		req        ftapi.RequestData
		Collection []ExpertisesUser
	}
	ExpertisesUserCUParams struct {
		UserID      int  `json:"user_id,omitempty"`
		ExpertiseID int  `json:"expertise_id,omitempty"`
		Value       int  `json:"value,omitempty"`
		Interested  bool `json:"interested,omitempty"`
		ContactMe   bool `json:"contact_me,omitempty"`
	}
)

func (eu *ExpertisesUser) Create(ctx context.Context, params ExpertisesUserCUParams) ftapi.CachedRequest {
	eu.req.Endpoint = ftapi.GetEndpoint("expertises_users", nil)
	eu.req.ExecuteMethod = func() {
		eu.req.Create(ctx, eu, ftapi.EncapsulatedMarshal("expertises_user", params))
	}
	return &eu.req
}

func (eu *ExpertisesUser) Delete(ctx context.Context) ftapi.Request {
	eu.req.Endpoint = ftapi.GetEndpoint("expertises_users/"+strconv.Itoa(eu.ID), nil)
	eu.req.ExecuteMethod = func() {
		eu.req.Delete(ctx, eu)
	}
	return &eu.req
}

func (eu *ExpertisesUser) Patch(ctx context.Context, params ExpertisesUserCUParams) ftapi.Request {
	eu.req.Endpoint = ftapi.GetEndpoint("expertises_users/"+strconv.Itoa(eu.ID), nil)
	eu.req.ExecuteMethod = func() {
		eu.req.Patch(ctx, ftapi.EncapsulatedMarshal("expertises_user", params))
	}
	return &eu.req
}

func (eu *ExpertisesUser) Get(ctx context.Context) ftapi.CachedRequest {
	eu.req.Endpoint = ftapi.GetEndpoint("expertises_users/"+strconv.Itoa(eu.ID), nil)
	eu.req.ExecuteMethod = func() {
		eu.req.Get(ctx, eu)
	}
	return &eu.req
}

func (eus *ExpertisesUsers) GetAll(ctx context.Context) ftapi.CollectionRequest {
	eus.req.Endpoint = ftapi.GetEndpoint("expertises_users", nil)
	eus.req.ExecuteMethod = func() {
		eus.req.GetAll(ctx, &eus.Collection)
	}
	return &eus.req
}
