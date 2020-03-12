package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	PartnershipsUser struct {
		req           ftapi.RequestData
		ID            int   `json:"id,omitempty"`
		PartnershipID int   `json:"partnership_id,omitempty"`
		FinalMark     int   `json:"final_mark,omitempty"`
		User          *User `json:"user,omitempty"`
	}
	PartnershipsUsers struct {
		req        ftapi.RequestData
		Collection []PartnershipsUser
	}
	PartnershipsUserCUParams struct {
		PartnershipID int `json:"partnership_id,omitempty"`
		UserID        int `json:"user_id,omitempty"`
		FinalMark     int `json:"final_mark,omitempty"`
	}
)

func (pu *PartnershipsUser) Create(ctx context.Context, params PartnershipsUserCUParams) ftapi.CachedRequest {
	pu.req.Endpoint = ftapi.GetEndpoint("partnerships_users", nil)
	pu.req.ExecuteMethod = func() {
		pu.req.Create(ctx, pu, ftapi.EncapsulatedMarshal("partnerships_user", params))
	}
	return &pu.req
}

func (pu *PartnershipsUser) Delete(ctx context.Context) ftapi.Request {
	pu.req.Endpoint = ftapi.GetEndpoint("partnerships_users/"+strconv.Itoa(pu.ID), nil)
	pu.req.ExecuteMethod = func() {
		pu.req.Delete(ctx, pu)
	}
	return &pu.req
}

func (pu *PartnershipsUser) Patch(ctx context.Context, params PartnershipsUserCUParams) ftapi.Request {
	pu.req.Endpoint = ftapi.GetEndpoint("partnerships_users/"+strconv.Itoa(pu.ID), nil)
	pu.req.ExecuteMethod = func() {
		pu.req.Patch(ctx, ftapi.EncapsulatedMarshal("partnerships_user", params))
	}
	return &pu.req
}

func (pu *PartnershipsUser) Get(ctx context.Context) ftapi.CachedRequest {
	pu.req.Endpoint = ftapi.GetEndpoint("partnerships_users/"+strconv.Itoa(pu.ID), nil)
	pu.req.ExecuteMethod = func() {
		pu.req.Get(ctx, pu)
	}
	return &pu.req
}

func (pus *PartnershipsUsers) GetAll(ctx context.Context) ftapi.CollectionRequest {
	pus.req.Endpoint = ftapi.GetEndpoint("partnerships_users", nil)
	pus.req.ExecuteMethod = func() {
		pus.req.GetAll(ctx, &pus.Collection)
	}
	return &pus.req
}
