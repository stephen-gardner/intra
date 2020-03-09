package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	Project struct {
		req         ftapi.RequestData
		ID          int    `json:"id,omitempty"`
		Name        string `json:"name,omitempty"`
		Slug        string `json:"slug,omitempty"`
		Description string `json:"description,omitempty"`
		Parent      struct {
			Name string `json:"name,omitempty"`
			ID   int    `json:"id,omitempty"`
			Slug string `json:"slug,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"parent,omitempty"`
		Children []struct {
			Name string `json:"name,omitempty"`
			ID   int    `json:"id,omitempty"`
			Slug string `json:"slug,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"children,omitempty"`
		Objectives []string   `json:"objectives,omitempty"`
		Tier       int        `json:"tier,omitempty"`
		CreatedAt  ftapi.Time `json:"created_at,omitempty"`
		UpdatedAt  ftapi.Time `json:"updated_at,omitempty"`
		Exam       bool       `json:"exam,omitempty"`
		Cursus     []struct {
			ID        int        `json:"id,omitempty"`
			CreatedAt ftapi.Time `json:"created_at,omitempty"`
			Name      string     `json:"name,omitempty"`
			Slug      string     `json:"slug,omitempty"`
		} `json:"cursus,omitempty"`
		Campus []struct {
			ID       int    `json:"id,omitempty"`
			Name     string `json:"name,omitempty"`
			TimeZone string `json:"time_zone,omitempty"`
			Language struct {
				ID         int        `json:"id,omitempty"`
				Name       string     `json:"name,omitempty"`
				Identifier string     `json:"identifier,omitempty"`
				CreatedAt  ftapi.Time `json:"created_at,omitempty"`
				UpdatedAt  ftapi.Time `json:"updated_at,omitempty"`
			} `json:"language,omitempty"`
			UsersCount  int `json:"users_count,omitempty"`
			VogsphereID int `json:"vogsphere_id,omitempty"`
		} `json:"campus,omitempty"`
		Skills []struct {
			ID        int        `json:"id,omitempty"`
			Name      string     `json:"name,omitempty"`
			CreatedAt ftapi.Time `json:"created_at,omitempty"`
		} `json:"skills,omitempty"`
		Tags []struct {
			ID   int    `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
			Kind string `json:"kind,omitempty"`
		} `json:"tags,omitempty"`
		ProjectSessions []struct {
			ID               int        `json:"id,omitempty"`
			Solo             bool       `json:"solo,omitempty"`
			BeginAt          ftapi.Time `json:"begin_at,omitempty"`
			EndAt            ftapi.Time `json:"end_at,omitempty"`
			EstimateTime     int        `json:"estimate_time,omitempty"`
			DurationDays     int        `json:"duration_days,omitempty"`
			TerminatingAfter int        `json:"terminating_after,omitempty"`
			ProjectID        int        `json:"project_id,omitempty"`
			CampusID         int        `json:"campus_id,omitempty"`
			CursusID         int        `json:"cursus_id,omitempty"`
			CreatedAt        ftapi.Time `json:"created_at,omitempty"`
			UpdatedAt        ftapi.Time `json:"updated_at,omitempty"`
			MaxPeople        int        `json:"max_people,omitempty"`
			IsSubscriptable  bool       `json:"is_subscriptable,omitempty"`
			Scales           []struct {
				ID               int  `json:"id,omitempty"`
				CorrectionNumber int  `json:"correction_number,omitempty"`
				IsPrimary        bool `json:"is_primary,omitempty"`
			} `json:"scales,omitempty"`
			Uploads []struct {
				ID   int    `json:"id,omitempty"`
				Name string `json:"name,omitempty"`
			} `json:"uploads,omitempty"`
			TeamBehaviour string `json:"team_behaviour,omitempty"`
		} `json:"project_sessions,omitempty"`
	}
	Projects struct {
		req        ftapi.RequestData
		Collection []Project
	}
)

func (proj *Project) Get(ctx context.Context) ftapi.CachedRequest {
	proj.req.Endpoint = ftapi.GetEndpoint("projects/"+strconv.Itoa(proj.ID), nil)
	proj.req.ExecuteMethod = func() {
		proj.req.Get(ctx, proj)
	}
	return &proj.req
}

func (projs *Projects) GetAll(ctx context.Context) ftapi.CollectionRequest {
	projs.req.Endpoint = ftapi.GetEndpoint("projects", nil)
	projs.req.ExecuteMethod = func() {
		projs.req.GetAll(ctx, &projs.Collection)
	}
	return &projs.req
}
