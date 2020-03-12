package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	CursusUser struct {
		req          ftapi.RequestData
		ID           int         `json:"id,omitempty"`
		CursusID     int         `json:"cursus_id,omitempty"`
		BeginAt      *ftapi.Time `json:"begin_at,omitempty"`
		EndAt        *ftapi.Time `json:"end_at,omitempty"`
		Grade        string      `json:"grade,omitempty"`
		Level        float64     `json:"level,omitempty"`
		HasCoalition bool        `json:"has_coalition,omitempty"`
		Skills       []struct {
			Skill
			Level float64 `json:"level,omitempty"`
		} `json:"skills"`
		User   *User   `json:"user,omitempty"`
		Cursus *Cursus `json:"cursus,omitempty"`
	}
	CursusUsers struct {
		req        ftapi.RequestData
		Collection []CursusUser
	}
	CursusUserCUParams struct {
		CursusID            int         `json:"cursus_id,omitempty"`
		UserID              int         `json:"user_id,omitempty"`
		BeginAt             *ftapi.Time `json:"begin_at,omitempty"`
		EndAt               *ftapi.Time `json:"end_at,omitempty"`
		HasCoalition        bool        `json:"has_coalition,omitempty"`
		SkipBeginValidation string      `json:"skip_begin_validation,omitempty"`
	}
)

func (cu *CursusUser) Create(ctx context.Context, params CursusUserCUParams) ftapi.CachedRequest {
	cu.req.Endpoint = ftapi.GetEndpoint("cursus_users", nil)
	cu.req.ExecuteMethod = func() {
		cu.req.Create(ctx, cu, ftapi.EncapsulatedMarshal("cursus_user", params))
	}
	return &cu.req
}

func (cu *CursusUser) Delete(ctx context.Context) ftapi.Request {
	cu.req.Endpoint = ftapi.GetEndpoint("cursus_users/"+strconv.Itoa(cu.ID), nil)
	cu.req.ExecuteMethod = func() {
		cu.req.Delete(ctx, cu)
	}
	return &cu.req
}

func (cu *CursusUser) Patch(ctx context.Context, params CursusUserCUParams) ftapi.Request {
	cu.req.Endpoint = ftapi.GetEndpoint("cursus_users/"+strconv.Itoa(cu.ID), nil)
	cu.req.ExecuteMethod = func() {
		cu.req.Patch(ctx, ftapi.EncapsulatedMarshal("cursus_user", params))
	}
	return &cu.req
}

func (cu *CursusUser) Get(ctx context.Context) ftapi.CachedRequest {
	cu.req.Endpoint = ftapi.GetEndpoint("cursus_users/"+strconv.Itoa(cu.ID), nil)
	cu.req.ExecuteMethod = func() {
		cu.req.Get(ctx, cu)
	}
	return &cu.req
}

func (cus *CursusUsers) GetAll(ctx context.Context) ftapi.CollectionRequest {
	cus.req.Endpoint = ftapi.GetEndpoint("cursus_users", nil)
	cus.req.ExecuteMethod = func() {
		cus.req.GetAll(ctx, &cus.Collection)
	}
	return &cus.req
}
