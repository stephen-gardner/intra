package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	AchievementsUser struct {
		req           ftapi.RequestData
		ID            int        `json:"id"`
		AchievementID int        `json:"achievement_id"`
		UserID        int        `json:"user_id"`
		Login         string     `json:"login"`
		NbrOfSuccess  int        `json:"nbr_of_success"`
		URL           string     `json:"url"`
		CreatedAt     ftapi.Time `json:"created_at"`
	}
	AchievementsUsers struct {
		req        ftapi.RequestData
		Collection []AchievementsUser
	}
	AchievementsUserCUParams struct {
		AchievementsUser struct {
			UserID        int `json:"user_id,omitempty"`
			AchievementID int `json:"achievement_id,omitempty"`
			NbrOfSuccess  int `json:"nbr_of_success,omitempty"`
		} `json:"achievements_user,omitempty"`
	}
)

func (au *AchievementsUser) Create(ctx context.Context, params AchievementsUserCUParams) ftapi.CachedRequest {
	au.req.Endpoint = ftapi.GetEndpoint("achievements_users", nil)
	au.req.ExecuteMethod = func() {
		au.req.Create(ftapi.GetClient(ctx, "public"), au, params)
	}
	return &au.req
}

func (au *AchievementsUser) Delete(ctx context.Context) ftapi.Request {
	au.req.Endpoint = ftapi.GetEndpoint("achievements_users/"+strconv.Itoa(au.ID), nil)
	au.req.ExecuteMethod = func() {
		au.req.Delete(ftapi.GetClient(ctx, "public"), au)
	}
	return &au.req
}

func (au *AchievementsUser) Patch(ctx context.Context, params AchievementsUserCUParams) ftapi.Request {
	au.req.Endpoint = ftapi.GetEndpoint("achievements_users/"+strconv.Itoa(au.ID), nil)
	au.req.ExecuteMethod = func() {
		au.req.Patch(ftapi.GetClient(ctx, "public"), au, params)
	}
	return &au.req
}

func (au *AchievementsUser) Get(ctx context.Context) ftapi.CachedRequest {
	au.req.Endpoint = ftapi.GetEndpoint("achievements_users/"+strconv.Itoa(au.ID), nil)
	au.req.ExecuteMethod = func() {
		au.req.Get(ftapi.GetClient(ctx, "public"), au)
	}
	return &au.req
}

func (aus *AchievementsUsers) GetAll(ctx context.Context) ftapi.CollectionRequest {
	aus.req.Endpoint = ftapi.GetEndpoint("achievements_users", nil)
	aus.req.ExecuteMethod = func() {
		aus.req.GetAll(ftapi.GetClient(ctx, "public"), &aus.Collection)
	}
	return &aus.req
}
