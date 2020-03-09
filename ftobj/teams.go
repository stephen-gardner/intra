package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	Team struct {
		req           ftapi.RequestData
		ID            int        `json:"id,omitempty"`
		Name          string     `json:"name,omitempty"`
		URL           string     `json:"url,omitempty"`
		FinalMark     int        `json:"final_mark,omitempty"`
		ProjectID     int        `json:"project_id,omitempty"`
		CreatedAt     ftapi.Time `json:"created_at,omitempty"`
		UpdatedAt     ftapi.Time `json:"updated_at,omitempty"`
		Status        string     `json:"status,omitempty"`
		TerminatingAt ftapi.Time `json:"terminating_at,omitempty"`
		Users         []struct {
			ID             int    `json:"id,omitempty"`
			Login          string `json:"login,omitempty"`
			URL            string `json:"url,omitempty"`
			Leader         bool   `json:"leader,omitempty"`
			Occurrence     int    `json:"occurrence,omitempty"`
			Validated      bool   `json:"validated,omitempty"`
			ProjectsUserID int    `json:"projects_user_id,omitempty"`
		} `json:"users,omitempty"`
		Locked           bool       `json:"locked?,omitempty"`
		Validated        bool       `json:"validated?,omitempty"`
		Closed           bool       `json:"closed?,omitempty"`
		RepoURL          string     `json:"repo_url,omitempty"`
		RepoUUID         string     `json:"repo_uuid,omitempty"`
		LockedAt         ftapi.Time `json:"locked_at,omitempty"`
		ClosedAt         ftapi.Time `json:"closed_at,omitempty"`
		ProjectSessionID int        `json:"project_session_id,omitempty"`
		ScaleTeams       []struct {
			ID        int        `json:"id,omitempty"`
			ScaleID   int        `json:"scale_id,omitempty"`
			Comment   string     `json:"comment,omitempty"`
			CreatedAt ftapi.Time `json:"created_at,omitempty"`
			UpdatedAt ftapi.Time `json:"updated_at,omitempty"`
			Feedback  string     `json:"feedback,omitempty"`
			FinalMark int        `json:"final_mark,omitempty"`
			Flag      struct {
				ID        int        `json:"id,omitempty"`
				Name      string     `json:"name,omitempty"`
				Positive  bool       `json:"positive,omitempty"`
				Icon      string     `json:"icon,omitempty"`
				CreatedAt ftapi.Time `json:"created_at,omitempty"`
				UpdatedAt ftapi.Time `json:"updated_at,omitempty"`
			} `json:"flag,omitempty"`
			BeginAt  ftapi.Time `json:"begin_at,omitempty"`
			FilledAt ftapi.Time `json:"filled_at,omitempty"`
		} `json:"scale_teams,omitempty"`
		TeamUploads []struct {
			ID        int        `json:"id,omitempty"`
			FinalMark int        `json:"final_mark,omitempty"`
			Comment   string     `json:"comment,omitempty"`
			CreatedAt ftapi.Time `json:"created_at,omitempty"`
			UploadID  int        `json:"upload_id,omitempty"`
		} `json:"team_uploads,omitempty"`
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
		team.req.Create(ctx, team, ftapi.EncapsulatedMarshal("team", params))
	}
	return &team.req
}

func (team *Team) Delete(ctx context.Context) ftapi.Request {
	team.req.Endpoint = ftapi.GetEndpoint("teams/"+strconv.Itoa(team.ID), nil)
	team.req.ExecuteMethod = func() {
		team.req.Delete(ctx, team)
	}
	return &team.req
}

func (team *Team) Patch(ctx context.Context, params TeamCUParams) ftapi.Request {
	team.req.Endpoint = ftapi.GetEndpoint("teams/"+strconv.Itoa(team.ID), nil)
	team.req.ExecuteMethod = func() {
		team.req.Patch(ctx, ftapi.EncapsulatedMarshal("team", params))
	}
	return &team.req
}

func (team *Team) Get(ctx context.Context) ftapi.CachedRequest {
	team.req.Endpoint = ftapi.GetEndpoint("teams/"+strconv.Itoa(team.ID), nil)
	team.req.ExecuteMethod = func() {
		team.req.Get(ctx, team)
	}
	return &team.req
}

func (teams *Teams) GetAll(ctx context.Context) ftapi.CollectionRequest {
	teams.req.Endpoint = ftapi.GetEndpoint("teams", nil)
	teams.req.ExecuteMethod = func() {
		teams.req.GetAll(ctx, &teams.Collection)
	}
	return &teams.req
}
