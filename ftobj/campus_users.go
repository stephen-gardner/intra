package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	CampusUser struct {
		req       ftapi.RequestData
		ID        int  `json:"id"`
		CampusID  int  `json:"campus_id"`
		UserID    int  `json:"user_id"`
		IsPrimary bool `json:"is_primary"`
	}
	CampusUsers struct {
		req        ftapi.RequestData
		Collection []CampusUser
	}
	CampusUserCUParams struct {
		CampusID  int  `json:"campus_id,omitempty"`
		UserID    int  `json:"user_id,omitempty"`
		IsPrimary bool `json:"is_primary,omitempty"`
	}
)

func (cu *CampusUser) Create(ctx context.Context, params CursusCUParams) ftapi.CachedRequest {
	cu.req.Endpoint = ftapi.GetEndpoint("campus_users", nil)
	cu.req.ExecuteMethod = func() {
		data := ftapi.EncapsulatedMarshal("campus_user", params)
		cu.req.Create(ftapi.GetClient(ctx, "public"), cu, data)
	}
	return &cu.req
}

func (cu *CampusUser) Delete(ctx context.Context) ftapi.Request {
	cu.req.Endpoint = ftapi.GetEndpoint("campus_users/"+strconv.Itoa(cu.ID), nil)
	cu.req.ExecuteMethod = func() {
		cu.req.Delete(ftapi.GetClient(ctx, "public"), cu)
	}
	return &cu.req
}

func (cu *CampusUser) Patch(ctx context.Context, params CursusCUParams) ftapi.Request {
	cu.req.Endpoint = ftapi.GetEndpoint("campus_users/"+strconv.Itoa(cu.ID), nil)
	cu.req.ExecuteMethod = func() {
		data := ftapi.EncapsulatedMarshal("campus_user", params)
		cu.req.Patch(ftapi.GetClient(ctx, "public"), cu, data)
	}
	return &cu.req
}

func (cu *CampusUser) Get(ctx context.Context) ftapi.CachedRequest {
	cu.req.Endpoint = ftapi.GetEndpoint("campus_users/"+strconv.Itoa(cu.ID), nil)
	cu.req.ExecuteMethod = func() {
		cu.req.Get(ftapi.GetClient(ctx, "public"), cu)
	}
	return &cu.req
}

func (cus *CampusUsers) GetAll(ctx context.Context) ftapi.CollectionRequest {
	cus.req.Endpoint = ftapi.GetEndpoint("campus_users", nil)
	cus.req.ExecuteMethod = func() {
		cus.req.GetAll(ftapi.GetClient(ctx, "public"), &cus.Collection)
	}
	return &cus.req
}
