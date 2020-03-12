package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	Partnership struct {
		req                  ftapi.RequestData
		ID                   int                 `json:"id,omitempty"`
		Name                 string              `json:"name,omitempty"`
		Slug                 string              `json:"slug,omitempty"`
		Difficulty           int                 `json:"difficulty,omitempty"`
		URL                  string              `json:"url,omitempty"`
		PartnershipsUsersURL string              `json:"partnerships_users_url,omitempty"`
		PartnershipsSkills   []PartnershipsSkill `json:"partnerships_skills,omitempty"`
	}
	Partnerships struct {
		req        ftapi.RequestData
		Collection []Partnership
	}
	PartnershipCUParams struct {
		Name                         string              `json:"name,omitempty"`
		Description                  string              `json:"description,omitempty"`
		Difficulty                   int                 `json:"difficulty,omitempty"`
		File                         string              `json:"file,omitempty"`
		CursusID                     int                 `json:"cursus_id,omitempty"`
		UserIDs                      []int               `json:"user_ids,omitempty"`
		PartnershipsSkillsAttributes []PartnershipsSkill `json:"partnerships_skills_attributes,omitempty"`
	}
	PartnershipsSkill struct {
		ID            int         `json:"id,omitempty"`
		PartnershipID int         `json:"partnership_id,omitempty"`
		SkillID       int         `json:"skill_id,omitempty"`
		Value         float64     `json:"value,omitempty"`
		CreatedAt     *ftapi.Time `json:"created_at,omitempty"`
		UpdatedAt     *ftapi.Time `json:"updated_at,omitempty"`
	}
)

func (p *Partnership) Create(ctx context.Context, params PartnershipCUParams) ftapi.CachedRequest {
	p.req.Endpoint = ftapi.GetEndpoint("partnerships", nil)
	p.req.ExecuteMethod = func() {
		p.req.Create(ctx, p, ftapi.EncapsulatedMarshal("partnership", params))
	}
	return &p.req
}

func (p *Partnership) Delete(ctx context.Context) ftapi.Request {
	p.req.Endpoint = ftapi.GetEndpoint("partnerships/"+strconv.Itoa(p.ID), nil)
	p.req.ExecuteMethod = func() {
		p.req.Delete(ctx, p)
	}
	return &p.req
}

func (p *Partnership) Patch(ctx context.Context, params PartnershipCUParams) ftapi.Request {
	p.req.Endpoint = ftapi.GetEndpoint("partnerships/"+strconv.Itoa(p.ID), nil)
	p.req.ExecuteMethod = func() {
		p.req.Patch(ctx, ftapi.EncapsulatedMarshal("partnership", params))
	}
	return &p.req
}

func (p *Partnership) Get(ctx context.Context) ftapi.CachedRequest {
	p.req.Endpoint = ftapi.GetEndpoint("partnerships/"+strconv.Itoa(p.ID), nil)
	p.req.ExecuteMethod = func() {
		p.req.Get(ctx, p)
	}
	return &p.req
}

func (ps *Partnerships) GetAll(ctx context.Context) ftapi.CollectionRequest {
	ps.req.Endpoint = ftapi.GetEndpoint("partnerships", nil)
	ps.req.ExecuteMethod = func() {
		ps.req.GetAll(ctx, &ps.Collection)
	}
	return &ps.req
}
