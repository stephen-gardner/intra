package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	Skill struct {
		req       ftapi.RequestData
		ID        int         `json:"id,omitempty"`
		Name      string      `json:"name,omitempty"`
		CreatedAt *ftapi.Time `json:"created_at,omitempty"`
	}
	Skills struct {
		req        ftapi.RequestData
		Collection []Skill
	}
	SkillCUParams struct {
		Name string `json:"name,omitempty"`
	}
)

func (skill *Skill) Create(ctx context.Context, params SkillCUParams) ftapi.CachedRequest {
	skill.req.Endpoint = ftapi.GetEndpoint("skills", nil)
	skill.req.ExecuteMethod = func() {
		skill.req.Create(ctx, skill, ftapi.EncapsulatedMarshal("skill", params))
	}
	return &skill.req
}

func (skill *Skill) Delete(ctx context.Context) ftapi.Request {
	skill.req.Endpoint = ftapi.GetEndpoint("skills/"+strconv.Itoa(skill.ID), nil)
	skill.req.ExecuteMethod = func() {
		skill.req.Delete(ctx, skill)
	}
	return &skill.req
}

func (skill *Skill) Patch(ctx context.Context, params SkillCUParams) ftapi.Request {
	skill.req.Endpoint = ftapi.GetEndpoint("skills/"+strconv.Itoa(skill.ID), nil)
	skill.req.ExecuteMethod = func() {
		skill.req.Patch(ctx, ftapi.EncapsulatedMarshal("skill", params))
	}
	return &skill.req
}

func (skill *Skill) Get(ctx context.Context) ftapi.CachedRequest {
	skill.req.Endpoint = ftapi.GetEndpoint("skills/"+strconv.Itoa(skill.ID), nil)
	skill.req.ExecuteMethod = func() {
		skill.req.Get(ctx, skill)
	}
	return &skill.req
}

func (skills *Skills) GetAll(ctx context.Context) ftapi.CollectionRequest {
	skills.req.Endpoint = ftapi.GetEndpoint("skills", nil)
	skills.req.ExecuteMethod = func() {
		skills.req.GetAll(ctx, &skills.Collection)
	}
	return &skills.req
}
