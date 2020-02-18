package intra

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type (
	Team struct {
		ID            int       `json:"id"`
		Name          string    `json:"name"`
		URL           string    `json:"url"`
		FinalMark     int       `json:"final_mark"`
		ProjectID     int       `json:"project_id"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		Status        string    `json:"status"`
		TerminatingAt time.Time `json:"terminating_at"`
		Users         []struct {
			ID             int    `json:"id"`
			Login          string `json:"login"`
			URL            string `json:"url"`
			Leader         bool   `json:"leader"`
			Occurrence     int    `json:"occurrence"`
			Validated      bool   `json:"validated"`
			ProjectsUserID int    `json:"projects_user_id"`
		} `json:"users"`
		Locked           bool      `json:"locked?"`
		Validated        bool      `json:"validated?"`
		Closed           bool      `json:"closed?"`
		RepoURL          string    `json:"repo_url"`
		RepoUUID         string    `json:"repo_uuid"`
		LockedAt         time.Time `json:"locked_at"`
		ClosedAt         time.Time `json:"closed_at"`
		ProjectSessionID int       `json:"project_session_id"`
		ScaleTeams       []struct {
			ID        int       `json:"id"`
			ScaleID   int       `json:"scale_id"`
			Comment   string    `json:"comment"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
			Feedback  string    `json:"feedback"`
			FinalMark int       `json:"final_mark"`
			Flag      struct {
				ID        int       `json:"id"`
				Name      string    `json:"name"`
				Positive  bool      `json:"positive"`
				Icon      string    `json:"icon"`
				CreatedAt time.Time `json:"created_at"`
				UpdatedAt time.Time `json:"updated_at"`
			} `json:"flag"`
			BeginAt  time.Time `json:"begin_at"`
			FilledAt time.Time `json:"filled_at"`
		} `json:"scale_teams"`
		TeamUploads []struct {
			ID        int       `json:"id"`
			FinalMark int       `json:"final_mark"`
			Comment   string    `json:"comment"`
			CreatedAt time.Time `json:"created_at"`
			UploadID  int       `json:"upload_id"`
		} `json:"team_uploads"`
	}
	Teams []Team
)

func (team *Team) Delete(ctx context.Context) error {
	endpoint := GetEndpoint("teams/"+strconv.Itoa(team.ID), nil)
	_, _, err := RunRequest(GetClient(ctx, "public", "projects"), http.MethodDelete, endpoint, nil)
	if err == nil {
		intraCache.delete(catTeams, team.ID)
	}
	return err
}

func (team *Team) Patch(ctx context.Context, bypassCache bool, params url.Values) (int, error) {
	endpoint := GetEndpoint("teams/"+strconv.Itoa(team.ID), nil)
	status, _, err := RunRequest(GetClient(ctx, "public", "projects"), http.MethodPatch, endpoint, params)
	if err == nil && !bypassCache {
		cached := *team
		intraCache.put(catTeams, cached.ID, &cached)
	}
	return status, err
}

func (team *Team) Get(ctx context.Context, bypassCache bool) error {
	if !bypassCache {
		if t, present := intraCache.get(catTeams, team.ID); present {
			*team = *t.(*Team)
			return nil
		}
	}
	teams := &Teams{}
	if err := teams.GetAll(ctx, bypassCache, getSingleParams(team.ID)); err != nil {
		return err
	}
	if len(*teams) == 0 {
		return fmt.Errorf("team %d does not exist", team.ID)
	}
	*team = (*teams)[0]
	return nil
}

func (teams *Teams) GetAll(ctx context.Context, bypassCache bool, params url.Values) error {
	if err := GetAll(GetClient(ctx, "public"), "teams", params, teams); err != nil {
		return err
	}
	if !bypassCache {
		for _, team := range *teams {
			cached := team
			intraCache.put(catTeams, cached.ID, &cached)
		}
	}
	return nil
}
