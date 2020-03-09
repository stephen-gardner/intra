package ftobj

import (
	"context"
	"fmt"
	"intra/ftapi"
	"strconv"
)

type (
	UserClose struct {
		req               ftapi.RequestData
		ID                int        `json:"id,omitempty"`
		Reason            string     `json:"reason,omitempty"`
		State             string     `json:"state,omitempty"`
		PrimaryCampusID   int        `json:"primary_campus_id,omitempty"`
		CreatedAt         ftapi.Time `json:"created_at,omitempty"`
		UpdatedAt         ftapi.Time `json:"updated_at,omitempty"`
		CommunityServices []struct {
			ID         int        `json:"id,omitempty"`
			Duration   int        `json:"duration,omitempty"`
			ScheduleAt ftapi.Time `json:"schedule_at,omitempty"`
			Occupation string     `json:"occupation,omitempty"`
			State      string     `json:"state,omitempty"`
			CreatedAt  ftapi.Time `json:"created_at,omitempty"`
			UpdatedAt  ftapi.Time `json:"updated_at,omitempty"`
		} `json:"community_services,omitempty"`
		User struct {
			ID    int    `json:"id,omitempty"`
			Login string `json:"login,omitempty"`
			URL   string `json:"url,omitempty"`
		} `json:"user,omitempty"`
		Closer struct {
			ID    int    `json:"id,omitempty"`
			Login string `json:"login,omitempty"`
			URL   string `json:"url,omitempty"`
		} `json:"closer,omitempty"`
	}
	UserCloses struct {
		req        ftapi.RequestData
		Collection []UserClose
	}
	UserCloseCUParams struct {
		UserID   int    `json:"user_id,omitempty"`
		CloserID int    `json:"closer_id,omitempty"`
		Kind     string `json:"kind,omitempty"`
		Reason   string `json:"reason,omitempty"`
	}
)

const (
	CloseKindBlackHole         = "black_hole"
	CloseKindDeserter          = "deserter"
	CloseKindNonAdmitted       = "non_admitted"
	CloseKindOther             = "other"
	CloseKindSeriousMisconduct = "serious_misconduct"
)

func (uc *UserClose) Create(ctx context.Context, params UserCloseCUParams) ftapi.CachedRequest {
	uc.req.Endpoint = ftapi.GetEndpoint("closes", nil)
	uc.req.ExecuteMethod = func() {
		uc.req.Create(ctx, uc, ftapi.EncapsulatedMarshal("close", params))
	}
	return &uc.req
}

func (uc *UserClose) Delete(ctx context.Context) ftapi.Request {
	uc.req.Endpoint = ftapi.GetEndpoint("closes/"+strconv.Itoa(uc.ID), nil)
	uc.req.ExecuteMethod = func() {
		uc.req.Delete(ctx, uc)
	}
	return &uc.req
}

func (uc *UserClose) Patch(ctx context.Context, params UserCloseCUParams) ftapi.Request {
	uc.req.Endpoint = ftapi.GetEndpoint("closes/"+strconv.Itoa(uc.ID), nil)
	uc.req.ExecuteMethod = func() {
		uc.req.Patch(ctx, ftapi.EncapsulatedMarshal("close", params))
	}
	return &uc.req
}

func (uc *UserClose) Get(ctx context.Context) ftapi.CachedRequest {
	uc.req.Endpoint = ftapi.GetEndpoint("closes/"+strconv.Itoa(uc.ID), nil)
	uc.req.ExecuteMethod = func() {
		uc.req.Get(ctx, uc)
	}
	return &uc.req
}

func (ucs *UserCloses) GetAll(ctx context.Context) ftapi.CollectionRequest {
	ucs.req.Endpoint = ftapi.GetEndpoint("closes", nil)
	ucs.req.ExecuteMethod = func() {
		ucs.req.GetAll(ctx, &ucs.Collection)
	}
	return &ucs.req
}

func (uc *UserClose) Reclose(ctx context.Context) ftapi.CachedRequest {
	uc.req.Endpoint = ftapi.GetEndpoint(fmt.Sprintf("closes/%d/close", uc.ID), nil)
	uc.req.ExecuteMethod = func() {
		err := uc.req.Patch(ctx, nil).Error
		if err != nil {
			return
		}
		uc.State = "close"
		ftapi.CacheObject(uc)
	}
	return &uc.req
}

func (uc *UserClose) Unclose(ctx context.Context) ftapi.CachedRequest {
	uc.req.Endpoint = ftapi.GetEndpoint(fmt.Sprintf("closes/%d/unclose", uc.ID), nil)
	uc.req.ExecuteMethod = func() {
		err := uc.req.Patch(ctx, nil).Error
		if err != nil {
			return
		}
		uc.State = "unclose"
		ftapi.CacheObject(uc)
	}
	return &uc.req
}
