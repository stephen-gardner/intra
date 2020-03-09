package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	ProjectSession struct {
		req              ftapi.RequestData
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
		Subscriptable    bool       `json:"is_subscriptable,omitempty"`
		Scales           []struct {
			ID               int  `json:"id,omitempty"`
			CorrectionNumber int  `json:"correction_number,omitempty"`
			Primary          bool `json:"is_primary,omitempty"`
		} `json:"scales,omitempty"`
		Uploads []struct {
			ID   int    `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"uploads,omitempty"`
		TeamBehavior string `json:"team_behaviour,omitempty"`
		Project      struct {
			ID     int    `json:"id,omitempty"`
			Name   string `json:"name,omitempty"`
			Slug   string `json:"slug,omitempty"`
			Parent struct {
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
			CreatedAt ftapi.Time `json:"created_at,omitempty"`
			UpdatedAt ftapi.Time `json:"updated_at,omitempty"`
			Exam      bool       `json:"exam,omitempty"`
		} `json:"project,omitempty"`
		Campus struct {
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
			UsersCount  int    `json:"users_count,omitempty"`
			VogsphereID int    `json:"vogsphere_id,omitempty"`
			Country     string `json:"country,omitempty"`
			Address     string `json:"address,omitempty"`
			Zip         string `json:"zip,omitempty"`
			City        string `json:"city,omitempty"`
			Website     string `json:"website,omitempty"`
			Facebook    string `json:"facebook,omitempty"`
			Twitter     string `json:"twitter,omitempty"`
		} `json:"campus,omitempty"`
		Cursus struct {
			ID        int        `json:"id,omitempty"`
			CreatedAt ftapi.Time `json:"created_at,omitempty"`
			Name      string     `json:"name,omitempty"`
			Slug      string     `json:"slug,omitempty"`
		} `json:"cursus,omitempty"`
		Evaluations []struct {
			ID   int    `json:"id,omitempty"`
			Kind string `json:"kind,omitempty"`
		} `json:"evaluations,omitempty"`
	}
	ProjectSessions struct {
		req        ftapi.RequestData
		Collection []ProjectSession
	}
)

func (ps *ProjectSession) Get(ctx context.Context) ftapi.CachedRequest {
	ps.req.Endpoint = ftapi.GetEndpoint("project_sessions/"+strconv.Itoa(ps.ID), nil)
	ps.req.ExecuteMethod = func() {
		ps.req.Get(ctx, ps)
	}
	return &ps.req
}

func (pss *ProjectSessions) GetAll(ctx context.Context) ftapi.CollectionRequest {
	pss.req.Endpoint = ftapi.GetEndpoint("project_sessions", nil)
	pss.req.ExecuteMethod = func() {
		pss.req.GetAll(ctx, &pss.Collection)
	}
	return &pss.req
}
