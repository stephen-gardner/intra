package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	AchievementsUser struct {
		req           ftapi.RequestData
		ID            int        `json:"id,omitempty"`
		AchievementID int        `json:"achievement_id,omitempty"`
		UserID        int        `json:"user_id,omitempty"`
		Login         string     `json:"login,omitempty"`
		NbrOfSuccess  int        `json:"nbr_of_success,omitempty"`
		URL           string     `json:"url,omitempty"`
		CreatedAt     ftapi.Time `json:"created_at,omitempty"`
	}
	AchievementsUsers struct {
		req        ftapi.RequestData
		Collection []AchievementsUser
	}
	AchievementsUserCUParams struct {
		UserID        int `json:"user_id,omitempty"`
		AchievementID int `json:"achievement_id,omitempty"`
		NbrOfSuccess  int `json:"nbr_of_success,omitempty"`
	}
)

func (au *AchievementsUser) Create(ctx context.Context, params AchievementsUserCUParams) ftapi.CachedRequest {
	au.req.Endpoint = ftapi.GetEndpoint("achievements_users", nil)
	au.req.ExecuteMethod = func() {
		au.req.Create(ctx, au, ftapi.EncapsulatedMarshal("achievements_user", params))
	}
	return &au.req
}

func (au *AchievementsUser) Delete(ctx context.Context) ftapi.Request {
	au.req.Endpoint = ftapi.GetEndpoint("achievements_users/"+strconv.Itoa(au.ID), nil)
	au.req.ExecuteMethod = func() {
		au.req.Delete(ctx, au)
	}
	return &au.req
}

func (au *AchievementsUser) Patch(ctx context.Context, params AchievementsUserCUParams) ftapi.Request {
	au.req.Endpoint = ftapi.GetEndpoint("achievements_users/"+strconv.Itoa(au.ID), nil)
	au.req.ExecuteMethod = func() {
		au.req.Patch(ctx, ftapi.EncapsulatedMarshal("achievements_user", params))
	}
	return &au.req
}

func (au *AchievementsUser) Get(ctx context.Context) ftapi.CachedRequest {
	au.req.Endpoint = ftapi.GetEndpoint("achievements_users/"+strconv.Itoa(au.ID), nil)
	au.req.ExecuteMethod = func() {
		au.req.Get(ctx, au)
	}
	return &au.req
}

func (aus *AchievementsUsers) GetAll(ctx context.Context) ftapi.CollectionRequest {
	aus.req.Endpoint = ftapi.GetEndpoint("achievements_users", nil)
	aus.req.ExecuteMethod = func() {
		aus.req.GetAll(ctx, &aus.Collection)
	}
	return &aus.req
}
