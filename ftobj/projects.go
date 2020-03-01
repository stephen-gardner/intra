package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	Project struct {
		req         ftapi.RequestData
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Slug        string `json:"slug"`
		Description string `json:"description"`
		Parent      struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
			Slug string `json:"slug"`
			URL  string `json:"url"`
		} `json:"parent"`
		Children []struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
			Slug string `json:"slug"`
			URL  string `json:"url"`
		} `json:"children"`
		Objectives []string   `json:"objectives"`
		Tier       int        `json:"tier"`
		CreatedAt  ftapi.Time `json:"created_at"`
		UpdatedAt  ftapi.Time `json:"updated_at"`
		Exam       bool       `json:"exam"`
		Cursus     []struct {
			ID        int        `json:"id"`
			CreatedAt ftapi.Time `json:"created_at"`
			Name      string     `json:"name"`
			Slug      string     `json:"slug"`
		} `json:"cursus"`
		Campus []struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			TimeZone string `json:"time_zone"`
			Language struct {
				ID         int        `json:"id"`
				Name       string     `json:"name"`
				Identifier string     `json:"identifier"`
				CreatedAt  ftapi.Time `json:"created_at"`
				UpdatedAt  ftapi.Time `json:"updated_at"`
			} `json:"language"`
			UsersCount  int `json:"users_count"`
			VogsphereID int `json:"vogsphere_id"`
		} `json:"campus"`
		Skills []struct {
			ID        int        `json:"id"`
			Name      string     `json:"name"`
			CreatedAt ftapi.Time `json:"created_at"`
		} `json:"skills"`
		Tags []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Kind string `json:"kind"`
		} `json:"tags"`
		ProjectSessions []struct {
			ID               int        `json:"id"`
			Solo             bool       `json:"solo"`
			BeginAt          ftapi.Time `json:"begin_at"`
			EndAt            ftapi.Time `json:"end_at"`
			EstimateTime     int        `json:"estimate_time"`
			DurationDays     int        `json:"duration_days"`
			TerminatingAfter int        `json:"terminating_after"`
			ProjectID        int        `json:"project_id"`
			CampusID         int        `json:"campus_id"`
			CursusID         int        `json:"cursus_id"`
			CreatedAt        ftapi.Time `json:"created_at"`
			UpdatedAt        ftapi.Time `json:"updated_at"`
			MaxPeople        int        `json:"max_people"`
			IsSubscriptable  bool       `json:"is_subscriptable"`
			Scales           []struct {
				ID               int  `json:"id"`
				CorrectionNumber int  `json:"correction_number"`
				IsPrimary        bool `json:"is_primary"`
			} `json:"scales"`
			Uploads []struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"uploads"`
			TeamBehaviour string `json:"team_behaviour"`
		} `json:"project_sessions"`
	}
	Projects struct {
		req        ftapi.RequestData
		Collection []Project
	}
)

func (proj *Project) Get(ctx context.Context) ftapi.CachedRequest {
	proj.req.Endpoint = ftapi.GetEndpoint("projects/"+strconv.Itoa(proj.ID), nil)
	proj.req.ExecuteMethod = func() {
		proj.req.Get(ftapi.GetClient(ctx, "public"), proj)
	}
	return &proj.req
}

func (projs *Projects) GetAll(ctx context.Context) ftapi.CollectionRequest {
	projs.req.Endpoint = ftapi.GetEndpoint("projects", nil)
	projs.req.ExecuteMethod = func() {
		projs.req.GetAll(ftapi.GetClient(ctx, "public"), &projs.Collection)
	}
	return &projs.req
}
