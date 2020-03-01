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
		ID            int    `json:"id"`
		Occurrence    int    `json:"occurrence"`
		FinalMark     int    `json:"final_mark"`
		Status        string `json:"status"`
		Validated     bool   `json:"validated?"`
		CurrentTeamID int    `json:"current_team_id"`
		Project       struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Slug     string `json:"slug"`
			ParentID int    `json:"parent_id"`
		} `json:"project"`
		CursusIds []int `json:"cursus_ids"`
		User      struct {
			ID    int    `json:"id"`
			Login string `json:"login"`
			URL   string `json:"url"`
		} `json:"user"`
		Teams []struct {
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
		} `json:"teams"`
	}
	ProjectsUsers struct {
		req        ftapi.RequestData
		Collection []ProjectsUser
	}
	ProjectsUserCUParams struct {
		ProjectsUser struct {
			ProjectID           int        `json:"project_id,omitempty"`
			UserID              int        `json:"user_id,omitempty"`
			CreatedAt           ftapi.Time `json:"created_at,omitempty"`
			UpdatedAt           ftapi.Time `json:"updated_at,omitempty"`
			Occurrence          int        `json:"occurrence,omitempty"`
			FinalMark           int        `json:"final_mark,omitempty"`
			RetriableAt         ftapi.Time `json:"retriable_at,omitempty"`
			MarkedAt            ftapi.Time `json:"marked_at,omitempty"`
			SkipCheckPermission string     `json:"skip_check_permission,omitempty"`
		} `json:"projects_user,omitempty"`
	}
)

func (pu *ProjectsUser) Create(ctx context.Context, params ProjectsUserCUParams) ftapi.CachedRequest {
	pu.req.Endpoint = ftapi.GetEndpoint("projects_users", nil)
	pu.req.ExecuteMethod = func() {
		pu.req.Create(ftapi.GetClient(ctx, "public", "projects"), pu, params)
	}
	return &pu.req
}

func (pu *ProjectsUser) Delete(ctx context.Context) ftapi.Request {
	pu.req.Endpoint = ftapi.GetEndpoint("projects_users/"+strconv.Itoa(pu.ID), nil)
	pu.req.ExecuteMethod = func() {
		pu.req.Delete(ftapi.GetClient(ctx, "public", "projects"), pu)
	}
	return &pu.req
}

func (pu *ProjectsUser) Patch(ctx context.Context, params ProjectsUserCUParams) ftapi.Request {
	pu.req.Endpoint = ftapi.GetEndpoint("projects_users/"+strconv.Itoa(pu.ID), nil)
	pu.req.ExecuteMethod = func() {
		pu.req.Patch(ftapi.GetClient(ctx, "public", "projects"), pu, params)
	}
	return &pu.req
}

func (pu *ProjectsUser) Get(ctx context.Context) ftapi.CachedRequest {
	pu.req.Endpoint = ftapi.GetEndpoint("projects_users/"+strconv.Itoa(pu.ID), nil)
	pu.req.ExecuteMethod = func() {
		pu.req.Get(ftapi.GetClient(ctx, "public"), pu)
	}
	return &pu.req
}

func (pus *ProjectsUsers) GetAll(ctx context.Context) ftapi.CollectionRequest {
	pus.req.Endpoint = ftapi.GetEndpoint("projects_users", nil)
	pus.req.ExecuteMethod = func() {
		pus.req.GetAll(ftapi.GetClient(ctx, "public"), &pus.Collection)
	}
	return &pus.req
}

func (pu *ProjectsUser) Retry(ctx context.Context) ftapi.Request {
	pu.req.Endpoint = ftapi.GetEndpoint(fmt.Sprintf("projects_users/%d/retry", pu.ID), nil)
	pu.req.ExecuteMethod = func() {
		pu.req.Patch(ftapi.GetClient(ctx, "public", "projects"), pu, nil)
	}
	return &pu.req
}
