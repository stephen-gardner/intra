package ftobj

import (
	"context"
	"fmt"
	"intra/ftapi"
	"strconv"
)

type (
	ProjectsUser struct {
		req           ftapi.RequestData
		ID            int    `json:"id,omitempty"`
		Occurrence    int    `json:"occurrence,omitempty"`
		FinalMark     int    `json:"final_mark,omitempty"`
		Status        string `json:"status,omitempty"`
		Validated     bool   `json:"validated?,omitempty"`
		CurrentTeamID int    `json:"current_team_id,omitempty"`
		Project       struct {
			ID       int    `json:"id,omitempty"`
			Name     string `json:"name,omitempty"`
			Slug     string `json:"slug,omitempty"`
			ParentID int    `json:"parent_id,omitempty"`
		} `json:"project,omitempty"`
		CursusIds []int `json:"cursus_ids,omitempty"`
		User      struct {
			ID    int    `json:"id,omitempty"`
			Login string `json:"login,omitempty"`
			URL   string `json:"url,omitempty"`
		} `json:"user,omitempty"`
		Teams []struct {
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
		} `json:"teams,omitempty"`
	}
	ProjectsUsers struct {
		req        ftapi.RequestData
		Collection []ProjectsUser
	}
	ProjectsUserCUParams struct {
		ProjectID           int        `json:"project_id,omitempty"`
		UserID              int        `json:"user_id,omitempty"`
		CreatedAt           ftapi.Time `json:"created_at,omitempty"`
		UpdatedAt           ftapi.Time `json:"updated_at,omitempty"`
		Occurrence          int        `json:"occurrence,omitempty"`
		FinalMark           int        `json:"final_mark,omitempty"`
		RetriableAt         ftapi.Time `json:"retriable_at,omitempty"`
		MarkedAt            ftapi.Time `json:"marked_at,omitempty"`
		SkipCheckPermission string     `json:"skip_check_permission,omitempty"`
	}
)

func (pu *ProjectsUser) Create(ctx context.Context, params ProjectsUserCUParams) ftapi.CachedRequest {
	pu.req.Endpoint = ftapi.GetEndpoint("projects_users", nil)
	pu.req.ExecuteMethod = func() {
		pu.req.Create(ctx, pu, ftapi.EncapsulatedMarshal("projects_user", params))
	}
	return &pu.req
}

func (pu *ProjectsUser) Delete(ctx context.Context) ftapi.Request {
	pu.req.Endpoint = ftapi.GetEndpoint("projects_users/"+strconv.Itoa(pu.ID), nil)
	pu.req.ExecuteMethod = func() {
		pu.req.Delete(ctx, pu)
	}
	return &pu.req
}

func (pu *ProjectsUser) Patch(ctx context.Context, params ProjectsUserCUParams) ftapi.Request {
	pu.req.Endpoint = ftapi.GetEndpoint("projects_users/"+strconv.Itoa(pu.ID), nil)
	pu.req.ExecuteMethod = func() {
		pu.req.Patch(ctx, ftapi.EncapsulatedMarshal("projects_user", params))
	}
	return &pu.req
}

func (pu *ProjectsUser) Get(ctx context.Context) ftapi.CachedRequest {
	pu.req.Endpoint = ftapi.GetEndpoint("projects_users/"+strconv.Itoa(pu.ID), nil)
	pu.req.ExecuteMethod = func() {
		pu.req.Get(ctx, pu)
	}
	return &pu.req
}

func (pus *ProjectsUsers) GetAll(ctx context.Context) ftapi.CollectionRequest {
	pus.req.Endpoint = ftapi.GetEndpoint("projects_users", nil)
	pus.req.ExecuteMethod = func() {
		pus.req.GetAll(ctx, &pus.Collection)
	}
	return &pus.req
}

func (pu *ProjectsUser) Retry(ctx context.Context) ftapi.Request {
	pu.req.Endpoint = ftapi.GetEndpoint(fmt.Sprintf("projects_users/%d/retry", pu.ID), nil)
	pu.req.ExecuteMethod = func() {
		pu.req.Patch(ctx, nil)
	}
	return &pu.req
}
