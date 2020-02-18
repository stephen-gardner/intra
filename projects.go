package intra

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

type (
	Project struct {
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
		Objectives []string  `json:"objectives"`
		Tier       int       `json:"tier"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		Exam       bool      `json:"exam"`
		Cursus     []struct {
			ID        int       `json:"id"`
			CreatedAt time.Time `json:"created_at"`
			Name      string    `json:"name"`
			Slug      string    `json:"slug"`
		} `json:"cursus"`
		Campus []struct {
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
			UsersCount  int `json:"users_count"`
			VogsphereID int `json:"vogsphere_id"`
		} `json:"campus"`
		Skills []struct {
			ID        int       `json:"id"`
			Name      string    `json:"name"`
			CreatedAt time.Time `json:"created_at"`
		} `json:"skills"`
		Tags []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Kind string `json:"kind"`
		} `json:"tags"`
		ProjectSessions []struct {
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
			IsSubscriptable  bool      `json:"is_subscriptable"`
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
	Projects []Project
)

func (project *Project) Get(ctx context.Context, bypassCache bool) error {
	if !bypassCache {
		if proj, present := intraCache.get(catProjects, project.ID); present {
			*project = *proj.(*Project)
			return nil
		}
	}
	projects := &Projects{}
	if err := projects.GetAll(ctx, bypassCache, getSingleParams(project.ID)); err != nil {
		return err
	}
	if len(*projects) == 0 {
		return fmt.Errorf("project %d does not exist", project.ID)
	}
	*project = (*projects)[0]
	return nil
}

func (projects *Projects) GetAll(ctx context.Context, bypassCache bool, params url.Values) error {
	if err := GetAll(GetClient(ctx, "public"), "projects", params, projects); err != nil {
		return err
	}
	if !bypassCache {
		for _, proj := range *projects {
			cached := proj
			intraCache.put(catProjects, cached.ID, &cached)
		}
	}
	return nil
}
