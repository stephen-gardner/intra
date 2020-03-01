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
		ID                int        `json:"id"`
		Reason            string     `json:"reason"`
		State             string     `json:"state"`
		PrimaryCampusID   int        `json:"primary_campus_id"`
		CreatedAt         ftapi.Time `json:"created_at"`
		UpdatedAt         ftapi.Time `json:"updated_at"`
		CommunityServices []struct {
			ID         int        `json:"id"`
			Duration   int        `json:"duration"`
			ScheduleAt ftapi.Time `json:"schedule_at"`
			Occupation string     `json:"occupation"`
			State      string     `json:"state"`
			CreatedAt  ftapi.Time `json:"created_at"`
			UpdatedAt  ftapi.Time `json:"updated_at"`
		} `json:"community_services"`
		User struct {
			ID    int    `json:"id"`
			Login string `json:"login"`
			URL   string `json:"url"`
		} `json:"user"`
		Closer struct {
			ID    int    `json:"id"`
			Login string `json:"login"`
			URL   string `json:"url"`
		} `json:"closer"`
	}
	UserCloses struct {
		req        ftapi.RequestData
		Collection []UserClose
	}
	UserCloseCUParams struct {
		Close struct {
			UserID   int    `json:"user_id,omitempty"`
			CloserID int    `json:"closer_id,omitempty"`
			Kind     string `json:"kind,omitempty"`
			Reason   string `json:"reason,omitempty"`
		} `json:"close,omitempty"`
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
		uc.req.Create(ftapi.GetClient(ctx, "public", "tig"), uc, params)
	}
	return &uc.req
}

func (uc *UserClose) Delete(ctx context.Context) ftapi.Request {
	uc.req.Endpoint = ftapi.GetEndpoint("closes/"+strconv.Itoa(uc.ID), nil)
	uc.req.ExecuteMethod = func() {
		uc.req.Delete(ftapi.GetClient(ctx, "public", "tig"), uc)
	}
	return &uc.req
}

func (uc *UserClose) Patch(ctx context.Context, params UserCloseCUParams) ftapi.Request {
	uc.req.Endpoint = ftapi.GetEndpoint("closes/"+strconv.Itoa(uc.ID), nil)
	uc.req.ExecuteMethod = func() {
		uc.req.Patch(ftapi.GetClient(ctx, "public", "tig"), uc, params)
	}
	return &uc.req
}

func (uc *UserClose) Get(ctx context.Context) ftapi.CachedRequest {
	uc.req.Endpoint = ftapi.GetEndpoint("closes/"+strconv.Itoa(uc.ID), nil)
	uc.req.ExecuteMethod = func() {
		uc.req.Get(ftapi.GetClient(ctx, "public", "tig"), uc)
	}
	return &uc.req
}

func (ucs *UserCloses) GetAll(ctx context.Context) ftapi.CollectionRequest {
	ucs.req.Endpoint = ftapi.GetEndpoint("closes", nil)
	ucs.req.ExecuteMethod = func() {
		ucs.req.GetAll(ftapi.GetClient(ctx, "public", "tig"), &ucs.Collection)
	}
	return &ucs.req
}

func (uc *UserClose) Reclose(ctx context.Context) ftapi.CachedRequest {
	uc.req.Endpoint = ftapi.GetEndpoint(fmt.Sprintf("closes/%d/close", uc.ID), nil)
	uc.req.ExecuteMethod = func() {
		err := uc.req.Patch(ftapi.GetClient(ctx, "public", "tig"), uc, nil).Error
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
		err := uc.req.Patch(ftapi.GetClient(ctx, "public", "tig"), uc, nil).Error
		if err != nil {
			return
		}
		uc.State = "unclose"
		ftapi.CacheObject(uc)
	}
	return &uc.req
}
