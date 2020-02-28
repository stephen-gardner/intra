package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	CursusUser struct {
		req          ftapi.RequestData
		ID           int        `json:"id"`
		CursusID     int        `json:"cursus_id"`
		BeginAt      ftapi.Time `json:"begin_at"`
		EndAt        ftapi.Time `json:"end_at"`
		Grade        string     `json:"grade"`
		Level        float64    `json:"level"`
		HasCoalition bool       `json:"has_coalition"`
		Skills       []struct {
			ID    int     `json:"id"`
			Name  string  `json:"name"`
			Level float64 `json:"level"`
		} `json:"skills"`
		User struct {
			ID    int    `json:"id"`
			Login string `json:"login"`
			URL   string `json:"url"`
		} `json:"user"`
		Cursus struct {
			ID        int        `json:"id"`
			CreatedAt ftapi.Time `json:"created_at"`
			Name      string     `json:"name"`
			Slug      string     `json:"slug"`
		} `json:"cursus"`
	}
	CursusUsers struct {
		req        ftapi.RequestData
		Collection []CursusUser
	}
	CursusUserCUParams struct {
		CursusUser struct {
			CursusID            int        `json:"cursus_id,omitempty"`
			UserID              int        `json:"user_id,omitempty"`
			BeginAt             ftapi.Time `json:"begin_at,omitempty"`
			EndAt               ftapi.Time `json:"end_at,omitempty"`
			HasCoalition        bool       `json:"has_coalition,omitempty"`
			SkipBeginValidation string     `json:"skip_begin_validation,omitempty"`
		} `json:"cursus_user,omitempty"`
	}
)

func (cu *CursusUser) Create(ctx context.Context, params CursusUserCUParams) ftapi.BasicRequest {
	cu.req.Endpoint = ftapi.GetEndpoint("cursus_users", nil)
	cu.req.ExecuteMethod = func() {
		cu.req.Create(ftapi.GetClient(ctx, "public"), cu, params)
	}
	return &cu.req
}

func (cu *CursusUser) Delete(ctx context.Context) ftapi.Request {
	cu.req.Endpoint = ftapi.GetEndpoint("cursus_users/"+strconv.Itoa(cu.ID), nil)
	cu.req.ExecuteMethod = func() {
		cu.req.Delete(ftapi.GetClient(ctx, "public"), cu)
	}
	return &cu.req
}

func (cu *CursusUser) Patch(ctx context.Context, params CursusUserCUParams) ftapi.BasicRequest {
	cu.req.Endpoint = ftapi.GetEndpoint("cursus_users/"+strconv.Itoa(cu.ID), nil)
	cu.req.ExecuteMethod = func() {
		cu.req.Patch(ftapi.GetClient(ctx, "public"), cu, params)
	}
	return &cu.req
}

func (cu *CursusUser) Get(ctx context.Context) ftapi.BasicRequest {
	cu.req.Endpoint = ftapi.GetEndpoint("cursus_users/"+strconv.Itoa(cu.ID), nil)
	cu.req.ExecuteMethod = func() {
		cu.req.Get(ftapi.GetClient(ctx, "public"), cu)
	}
	return &cu.req
}

func (cus *CursusUsers) GetAll(ctx context.Context) ftapi.CollectionRequest {
	cus.req.Endpoint = ftapi.GetEndpoint("cursus_users", nil)
	cus.req.ExecuteMethod = func() {
		cus.req.GetAll(ftapi.GetClient(ctx, "public"), &cus.Collection)
	}
	return &cus.req
}
