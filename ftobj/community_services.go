package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	CommunityService struct {
		req        ftapi.RequestData
		ID         int        `json:"id,omitempty"`
		Duration   int        `json:"duration,omitempty"`
		ScheduleAt ftapi.Time `json:"schedule_at,omitempty"`
		Occupation string     `json:"occupation,omitempty"`
		State      string     `json:"state,omitempty"`
		CreatedAt  ftapi.Time `json:"created_at,omitempty"`
		UpdatedAt  ftapi.Time `json:"updated_at,omitempty"`
		Close      UserClose  `json:"close,omitempty"`
	}
	CommunityServices struct {
		req        ftapi.RequestData
		Collection []CommunityService
	}
	CommunityServiceCUParams struct {
		Duration   int        `json:"duration,omitempty"`
		Occupation string     `json:"occupation,omitempty"`
		ScheduleAt ftapi.Time `json:"schedule_at,omitempty"`
		CloseID    int        `json:"close_id,omitempty"`
		TigerID    int        `json:"tiger_id,omitempty"`
	}
)

func (cs *CommunityService) Create(ctx context.Context, params CommunityServiceCUParams) ftapi.CachedRequest {
	cs.req.Endpoint = ftapi.GetEndpoint("community_services", nil)
	cs.req.ExecuteMethod = func() {
		cs.req.Create(ctx, cs, ftapi.EncapsulatedMarshal("community_service", params))
	}
	return &cs.req
}

func (cs *CommunityService) Delete(ctx context.Context) ftapi.Request {
	cs.req.Endpoint = ftapi.GetEndpoint("community_services/"+strconv.Itoa(cs.ID), nil)
	cs.req.ExecuteMethod = func() {
		cs.req.Delete(ctx, cs)
	}
	return &cs.req
}

func (cs *CommunityService) Patch(ctx context.Context, params CommunityServiceCUParams) ftapi.Request {
	cs.req.Endpoint = ftapi.GetEndpoint("community_services/"+strconv.Itoa(cs.ID), nil)
	cs.req.ExecuteMethod = func() {
		cs.req.Patch(ctx, ftapi.EncapsulatedMarshal("community_service", params))
	}
	return &cs.req
}

func (cs *CommunityService) Get(ctx context.Context) ftapi.CachedRequest {
	cs.req.Endpoint = ftapi.GetEndpoint("community_services/"+strconv.Itoa(cs.ID), nil)
	cs.req.ExecuteMethod = func() {
		cs.req.Get(ctx, cs)
	}
	return &cs.req
}

func (css *CommunityServices) GetAll(ctx context.Context) ftapi.CollectionRequest {
	css.req.Endpoint = ftapi.GetEndpoint("community_services", nil)
	css.req.ExecuteMethod = func() {
		css.req.GetAll(ctx, &css.Collection)
	}
	return &css.req
}
