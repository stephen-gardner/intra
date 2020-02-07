package intra

import (
	"context"
	"encoding/json"
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
	}
	Teams []Team
)

func (team *Team) PatchTeam(ctx context.Context, bypassCache bool, params url.Values) (int, []byte, error) {
	endpoint := GetEndpoint("teams/"+strconv.Itoa(team.ID), nil)
	status, respData, err := RunRequest(GetClient(ctx, "public", "projects"), http.MethodPatch, endpoint, params)
	if err == nil && !bypassCache {
		intraCache.put(team.URL, *team)
	}
	return status, respData, err
}

func (team *Team) GetTeam(ctx context.Context, bypassCache bool, ID int) error {
	IDStr := strconv.Itoa(ID)
	endpoint := GetEndpoint("teams/"+IDStr, nil)
	if !bypassCache {
		if t, present := intraCache.get(endpoint); present {
			*team = t.(Team)
			return nil
		}
	}
	params := url.Values{}
	params.Set("filter[id]", IDStr)
	params.Set("page[number]", "1")
	teams := &Teams{}
	if err := teams.GetAllTeams(ctx, bypassCache, params); err != nil {
		return err
	}
	if len(*teams) == 0 {
		return fmt.Errorf("team %d does not exist", ID)
	}
	*team = (*teams)[0]
	return nil
}

func (teams *Teams) GetAllTeams(ctx context.Context, bypassCache bool, params url.Values) error {
	data, err := GetAll(GetClient(ctx, "public"), "teams", params)
	if err != nil {
		return err
	}
	for _, dataPage := range data {
		var page Teams
		if err := json.Unmarshal(dataPage, &page); err != nil {
			return err
		}
		if !bypassCache {
			for _, team := range page {
				intraCache.put(team.URL, team)
			}
		}
		*teams = append(*teams, page...)
	}
	return nil
}
