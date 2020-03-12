package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	Project struct {
		req             ftapi.RequestData
		ID              int              `json:"id,omitempty"`
		Name            string           `json:"name,omitempty"`
		Slug            string           `json:"slug,omitempty"`
		Description     string           `json:"description,omitempty"`
		Parent          *Project         `json:"parent,omitempty"`
		Children        []Project        `json:"children,omitempty"`
		Objectives      []string         `json:"objectives,omitempty"`
		Tier            int              `json:"tier,omitempty"`
		CreatedAt       *ftapi.Time      `json:"created_at,omitempty"`
		UpdatedAt       *ftapi.Time      `json:"updated_at,omitempty"`
		Exam            bool             `json:"exam,omitempty"`
		Cursus          []Cursus         `json:"cursus,omitempty"`
		Campus          []Campus         `json:"campus,omitempty"`
		Skills          []Skill          `json:"skills,omitempty"`
		Tags            []Tag            `json:"tags,omitempty"`
		ProjectSessions []ProjectSession `json:"project_sessions,omitempty"`
	}
	Projects struct {
		req        ftapi.RequestData
		Collection []Project
	}
)

func (proj *Project) Get(ctx context.Context) ftapi.CachedRequest {
	proj.req.Endpoint = ftapi.GetEndpoint("projects/"+strconv.Itoa(proj.ID), nil)
	proj.req.ExecuteMethod = func() {
		proj.req.Get(ctx, proj)
	}
	return &proj.req
}

func (projs *Projects) GetAll(ctx context.Context) ftapi.CollectionRequest {
	projs.req.Endpoint = ftapi.GetEndpoint("projects", nil)
	projs.req.ExecuteMethod = func() {
		projs.req.GetAll(ctx, &projs.Collection)
	}
	return &projs.req
}
