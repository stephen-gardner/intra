package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	Team struct {
		req           ftapi.RequestData
		ID            int        `json:"id"`
		Name          string     `json:"name"`
		URL           string     `json:"url"`
		FinalMark     int        `json:"final_mark"`
		ProjectID     int        `json:"project_id"`
		CreatedAt     ftapi.Time `json:"created_at"`
		UpdatedAt     ftapi.Time `json:"updated_at"`
		Status        string     `json:"status"`
		TerminatingAt ftapi.Time `json:"terminating_at"`
		Users         []struct {
			ID             int    `json:"id"`
			Login          string `json:"login"`
			URL            string `json:"url"`
			Leader         bool   `json:"leader"`
			Occurrence     int    `json:"occurrence"`
			Validated      bool   `json:"validated"`
			ProjectsUserID int    `json:"projects_user_id"`
		} `json:"users"`
		Locked           bool       `json:"locked?"`
		Validated        bool       `json:"validated?"`
		Closed           bool       `json:"closed?"`
		RepoURL          string     `json:"repo_url"`
		RepoUUID         string     `json:"repo_uuid"`
		LockedAt         ftapi.Time `json:"locked_at"`
		ClosedAt         ftapi.Time `json:"closed_at"`
		ProjectSessionID int        `json:"project_session_id"`
		ScaleTeams       []struct {
			ID        int        `json:"id"`
			ScaleID   int        `json:"scale_id"`
			Comment   string     `json:"comment"`
			CreatedAt ftapi.Time `json:"created_at"`
			UpdatedAt ftapi.Time `json:"updated_at"`
			Feedback  string     `json:"feedback"`
			FinalMark int        `json:"final_mark"`
			Flag      struct {
				ID        int        `json:"id"`
				Name      string     `json:"name"`
				Positive  bool       `json:"positive"`
				Icon      string     `json:"icon"`
				CreatedAt ftapi.Time `json:"created_at"`
				UpdatedAt ftapi.Time `json:"updated_at"`
			} `json:"flag"`
			BeginAt  ftapi.Time `json:"begin_at"`
			FilledAt ftapi.Time `json:"filled_at"`
		} `json:"scale_teams"`
		TeamUploads []struct {
			ID        int        `json:"id"`
			FinalMark int        `json:"final_mark"`
			Comment   string     `json:"comment"`
			CreatedAt ftapi.Time `json:"created_at"`
			UploadID  int        `json:"upload_id"`
		} `json:"team_uploads"`
	}
	Teams struct {
		req        ftapi.RequestData
		Collection []Team
	}
	TeamCUParams struct {
		Name             string     `json:"name,omitempty"`
		CreatedAt        ftapi.Time `json:"created_at,omitempty"`
		UpdatedAt        ftapi.Time `json:"updated_at,omitempty"`
		LockedAt         ftapi.Time `json:"locked_at,omitempty"`
		ClosedAt         ftapi.Time `json:"closed_at,omitempty"`
		FinalMark        int        `json:"final_mark,omitempty"`
		RepoURL          string     `json:"repo_url,omitempty"`
		RepoUUID         string     `json:"repo_uuid,omitempty"`
		TerminatingAt    ftapi.Time `json:"terminating_at,omitempty"`
		ProjectSessionID int        `json:"project_session_id,omitempty"`
		UsersAttributes  struct {
			UserID     int  `json:"user_id,omitempty"`
			Leader     bool `json:"leader,omitempty"`
			Validated  bool `json:"validated,omitempty"`
			Occurrence int  `json:"occurrence,omitempty"`
		} `json:"teams_users_attributes,omitempty"`
	}
)

func (team *Team) Create(ctx context.Context, params TeamCUParams) ftapi.CachedRequest {
	team.req.Endpoint = ftapi.GetEndpoint("teams", nil)
	team.req.ExecuteMethod = func() {
		data := ftapi.EncapsulatedMarshal("team", params)
		team.req.Create(ftapi.GetClient(ctx, "public", "projects"), team, data)
	}
	return &team.req
}

func (team *Team) Delete(ctx context.Context) ftapi.Request {
	team.req.Endpoint = ftapi.GetEndpoint("teams/"+strconv.Itoa(team.ID), nil)
	team.req.ExecuteMethod = func() {
		team.req.Delete(ftapi.GetClient(ctx, "public", "projects"), team)
	}
	return &team.req
}

func (team *Team) Patch(ctx context.Context, params TeamCUParams) ftapi.Request {
	team.req.Endpoint = ftapi.GetEndpoint("teams/"+strconv.Itoa(team.ID), nil)
	team.req.ExecuteMethod = func() {
		data := ftapi.EncapsulatedMarshal("team", params)
		team.req.Patch(ftapi.GetClient(ctx, "public", "projects"), team, data)
	}
	return &team.req
}

func (team *Team) Get(ctx context.Context) ftapi.CachedRequest {
	team.req.Endpoint = ftapi.GetEndpoint("teams/"+strconv.Itoa(team.ID), nil)
	team.req.ExecuteMethod = func() {
		team.req.Get(ftapi.GetClient(ctx, "public"), team)
	}
	return &team.req
}

func (teams *Teams) GetAll(ctx context.Context) ftapi.CollectionRequest {
	teams.req.Endpoint = ftapi.GetEndpoint("teams", nil)
	teams.req.ExecuteMethod = func() {
		teams.req.GetAll(ftapi.GetClient(ctx, "public"), &teams.Collection)
	}
	return &teams.req
}
