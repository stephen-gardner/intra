package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	ProjectSession struct {
		req              ftapi.RequestData
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
		Subscriptable    bool       `json:"is_subscriptable"`
		Scales           []struct {
			ID               int  `json:"id"`
			CorrectionNumber int  `json:"correction_number"`
			Primary          bool `json:"is_primary"`
		} `json:"scales"`
		Uploads []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"uploads"`
		TeamBehavior string `json:"team_behaviour"`
		Project      struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Slug   string `json:"slug"`
			Parent struct {
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
			CreatedAt ftapi.Time `json:"created_at"`
			UpdatedAt ftapi.Time `json:"updated_at"`
			Exam      bool       `json:"exam"`
		} `json:"project"`
		Campus struct {
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
			UsersCount  int    `json:"users_count"`
			VogsphereID int    `json:"vogsphere_id"`
			Country     string `json:"country"`
			Address     string `json:"address"`
			Zip         string `json:"zip"`
			City        string `json:"city"`
			Website     string `json:"website"`
			Facebook    string `json:"facebook"`
			Twitter     string `json:"twitter"`
		} `json:"campus"`
		Cursus struct {
			ID        int        `json:"id"`
			CreatedAt ftapi.Time `json:"created_at"`
			Name      string     `json:"name"`
			Slug      string     `json:"slug"`
		} `json:"cursus"`
		Evaluations []struct {
			ID   int    `json:"id"`
			Kind string `json:"kind"`
		} `json:"evaluations"`
	}
	ProjectSessions struct {
		req        ftapi.RequestData
		Collection []ProjectSession
	}
)

func (ps *ProjectSession) Get(ctx context.Context) ftapi.BasicRequest {
	ps.req.Endpoint = ftapi.GetEndpoint("project_sessions/"+strconv.Itoa(ps.ID), nil)
	ps.req.ExecuteMethod = func() {
		ps.req.Get(ftapi.GetClient(ctx, "public"), ps)
	}
	return &ps.req
}

func (pss *ProjectSessions) GetAll(ctx context.Context) ftapi.CollectionRequest {
	pss.req.Endpoint = ftapi.GetEndpoint("project_sessions", nil)
	pss.req.ExecuteMethod = func() {
		pss.req.GetAll(ftapi.GetClient(ctx, "public"), &pss.Collection)
	}
	return &pss.req
}
