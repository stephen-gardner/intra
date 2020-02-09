package intra

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
	"time"
)

type (
	ProjectSession struct {
		ID               int       `json:"id"`
		Solo             bool      `json:"solo"`
		BeginAt          time.Time `json:"begin_at"`
		EndAt            time.Time `json:"end_at"`
		EstimateTime     int       `json:"estimate_time"`
		DurationDays     int       `json:"duration_days"`
		TerminatingAfter int       `json:"terminating_after"`
		ProjectID        int       `json:"project_id"`
		CampusID         int       `json:"campus_id"`
		CursusID         int       `json:"cursus_id"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		MaxPeople        int       `json:"max_people"`
		Subscriptable    bool      `json:"is_subscriptable"`
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
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
			Exam      bool      `json:"exam"`
		} `json:"project"`
		Campus struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			TimeZone string `json:"time_zone"`
			Language struct {
				ID         int       `json:"id"`
				Name       string    `json:"name"`
				Identifier string    `json:"identifier"`
				CreatedAt  time.Time `json:"created_at"`
				UpdatedAt  time.Time `json:"updated_at"`
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
			ID        int       `json:"id"`
			CreatedAt time.Time `json:"created_at"`
			Name      string    `json:"name"`
			Slug      string    `json:"slug"`
		} `json:"cursus"`
		Evaluations []struct {
			ID   int    `json:"id"`
			Kind string `json:"kind"`
		} `json:"evaluations"`
	}
	ProjectSessions []ProjectSession
)

func (pSession *ProjectSession) GetProjectSession(ctx context.Context, bypassCache bool, ID int) error {
	IDStr := strconv.Itoa(ID)
	endpoint := GetEndpoint("project_sessions/"+IDStr, nil)
	if !bypassCache {
		if ps, present := intraCache.get(endpoint); present {
			*pSession = ps.(ProjectSession)
			return nil
		}
	}
	params := url.Values{}
	params.Set("filter[id]", IDStr)
	params.Set("page[number]", "1")
	pSessions := &ProjectSessions{}
	err := pSessions.GetAllProjectSessions(ctx, bypassCache, params)
	if err == nil && len(*pSessions) > 0 {
		*pSession = (*pSessions)[0]
	}
	return err
}

func (pSessions *ProjectSessions) GetAllProjectSessions(ctx context.Context, bypassCache bool, params url.Values) error {
	data, err := GetAll(GetClient(ctx, "public"), "project_sessions", params)
	if err != nil {
		return err
	}
	for _, dataPage := range data {
		var page ProjectSessions
		if err := json.Unmarshal(dataPage, &page); err != nil {
			return err
		}
		if !bypassCache {
			for _, pSession := range page {
				endpoint := GetEndpoint("project_sessions/"+strconv.Itoa(pSession.ID), nil)
				intraCache.put(endpoint, pSession)
			}
		}
		*pSessions = append(*pSessions, page...)
	}
	return nil
}
