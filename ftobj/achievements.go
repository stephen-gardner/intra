package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	AchievementData struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Description  string `json:"description"`
		Tier         string `json:"tier"`
		Kind         string `json:"kind"`
		Visible      bool   `json:"visible"`
		Image        string `json:"image"`
		NbrOfSuccess int    `json:"nbr_of_success"`
		UsersURL     string `json:"users_url"`
	}
	Achievement struct {
		AchievementData
		req          ftapi.RequestData
		Achievements []AchievementData `json:"achievements"`
		Parent       AchievementData   `json:"parent"`
		Title        struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"title"`
	}
	Achievements struct {
		req        ftapi.RequestData
		Collection []Achievement
	}
	AchievementCUParams struct {
		Name         string `json:"name,omitempty"`
		InternalName string `json:"internal_name,omitempty"`
		Description  string `json:"description,omitempty"`
		Pedago       bool   `json:"pedago,omitempty"`
		Visible      bool   `json:"visible,omitempty"`
		NbrOfSuccess int    `json:"nbr_of_success,omitempty"`
		ParentID     int    `json:"parent_id,omitempty"`
		Image        string `json:"image,omitempty"`
		ImageCache   string `json:"image_cache,omitempty"`
		Kind         string `json:"kind,omitempty"`
		TitleID      int    `json:"title_id,omitempty"`
		Tier         string `json:"tier,omitempty"`
		Language     string `json:"lg,omitempty"`
		Position     int    `json:"position,omitempty"`
		Reward       string `json:"reward,omitempty"`
		CursusIDs    []int  `json:"cursus_ids,omitempty"`
		CampusIDs    []int  `json:"campus_ids,omitempty"`
	}
)

const (
	AchievementKindProject   = "project"
	AchievementKindSocial    = "social"
	AchievementKindScolarity = "scolarity"
	AchievementKindPedagogy  = "pedagogy"
	AchievementTierNone      = "none"
	AchievementTierEasy      = "easy"
	AchievementTierMedium    = "medium"
	AchievementTierHard      = "hard"
	AchievementTierChallenge = "challenge"
)

func (a *Achievement) Create(ctx context.Context, params AchievementCUParams) ftapi.CachedRequest {
	a.req.Endpoint = ftapi.GetEndpoint("achievements", nil)
	a.req.ExecuteMethod = func() {
		data := ftapi.EncapsulatedMarshal("achievement", params)
		a.req.Create(ftapi.GetClient(ctx, "public"), a, data)
	}
	return &a.req
}

func (a *Achievement) Delete(ctx context.Context) ftapi.Request {
	a.req.Endpoint = ftapi.GetEndpoint("achievements/"+strconv.Itoa(a.ID), nil)
	a.req.ExecuteMethod = func() {
		a.req.Delete(ftapi.GetClient(ctx, "public"), a)
	}
	return &a.req
}

func (a *Achievement) Patch(ctx context.Context, params AchievementCUParams) ftapi.Request {
	a.req.Endpoint = ftapi.GetEndpoint("achievements/"+strconv.Itoa(a.ID), nil)
	a.req.ExecuteMethod = func() {
		data := ftapi.EncapsulatedMarshal("achievement", params)
		a.req.Patch(ftapi.GetClient(ctx, "public"), a, data)
	}
	return &a.req
}

func (a *Achievement) Get(ctx context.Context) ftapi.CachedRequest {
	a.req.Endpoint = ftapi.GetEndpoint("achievements/"+strconv.Itoa(a.ID), nil)
	a.req.ExecuteMethod = func() {
		a.req.Get(ftapi.GetClient(ctx, "public"), a)
	}
	return &a.req
}

func (a *Achievements) GetAll(ctx context.Context) ftapi.CollectionRequest {
	a.req.Endpoint = ftapi.GetEndpoint("achievements", nil)
	a.req.ExecuteMethod = func() {
		a.req.GetAll(ftapi.GetClient(ctx, "public"), &a.Collection)
	}
	return &a.req
}
