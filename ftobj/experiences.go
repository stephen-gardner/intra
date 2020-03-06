package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
	"time"
)

type (
	Experience struct {
		req               ftapi.RequestData
		ID                int        `json:"id"`
		UserID            int        `json:"user_id"`
		SkillID           int        `json:"skill_id"`
		ExperiancableID   int        `json:"experiancable_id"`
		ExperiancableType string     `json:"experiancable_type"`
		Amount            int        `json:"experience"`
		CreatedAt         ftapi.Time `json:"created_at"`
		CursusID          int        `json:"cursus_id"`
		User              struct {
			ID    int    `json:"id"`
			Login string `json:"login"`
			URL   string `json:"url"`
		} `json:"user"`
		Skill struct {
			ID        int        `json:"id"`
			Name      string     `json:"name"`
			CreatedAt ftapi.Time `json:"created_at"`
		} `json:"skill"`
		Cursus struct {
			ID        int        `json:"id"`
			CreatedAt ftapi.Time `json:"created_at"`
			Name      string     `json:"name"`
			Slug      string     `json:"slug"`
		} `json:"cursus"`
		Experiancable struct {
			ID            int    `json:"id"`
			Occurrence    int    `json:"occurrence"`
			FinalMark     int    `json:"final_mark"`
			Status        string `json:"status"`
			Validated     bool   `json:"validated?"`
			CurrentTeamID int    `json:"current_team_id"`
			Project       struct {
				ID       int         `json:"id"`
				Name     string      `json:"name"`
				Slug     string      `json:"slug"`
				ParentID interface{} `json:"parent_id"`
			} `json:"project"`
			CursusIds   []int      `json:"cursus_ids"`
			MarkedAt    ftapi.Time `json:"marked_at"`
			Marked      bool       `json:"marked"`
			RetriableAt ftapi.Time `json:"retriable_at"`
		} `json:"experiancable"`
	}
	Experiences struct {
		req        ftapi.RequestData
		Collection []Experience
	}
	ExperienceCUParams struct {
		UserID            int        `json:"user_id,omitempty"`
		SkillID           int        `json:"skill_id,omitempty"`
		ExperiancableID   int        `json:"experiancable_id,omitempty"`
		ExperiancableType string     `json:"experiancable_type,omitempty"`
		Experience        int        `json:"experience,omitempty"`
		CreatedAt         ftapi.Time `json:"created_at,omitempty"`
		CursusID          int        `json:"cursus_id,omitempty"`
	}
)

var expLevels = []int{
	0,
	113,
	241,
	386,
	551,
	738,
	950,
	1190,
	1462,
	1770,
	2119,
	2515,
	2963,
	3471,
	4046,
	4698,
	5437,
	6274,
	7223,
	8298,
	9516,
	10896,
	-1,
}

func (exp *Experience) Create(ctx context.Context, params ExperienceCUParams) ftapi.CachedRequest {
	exp.req.Endpoint = ftapi.GetEndpoint("experiences", nil)
	exp.req.ExecuteMethod = func() {
		exp.req.Create(ctx, exp, ftapi.EncapsulatedMarshal("experience", params))
	}
	return &exp.req
}

func (exp *Experience) Delete(ctx context.Context) ftapi.Request {
	exp.req.Endpoint = ftapi.GetEndpoint("experiences/"+strconv.Itoa(exp.ID), nil)
	exp.req.ExecuteMethod = func() {
		exp.req.Delete(ctx, exp)
	}
	return &exp.req
}

func (exp *Experience) Patch(ctx context.Context, params ExperienceCUParams) ftapi.Request {
	exp.req.Endpoint = ftapi.GetEndpoint("experiences/"+strconv.Itoa(exp.ID), nil)
	exp.req.ExecuteMethod = func() {
		exp.req.Patch(ctx, ftapi.EncapsulatedMarshal("experience", params))
	}
	return &exp.req
}

func (exp *Experience) Get(ctx context.Context) ftapi.CachedRequest {
	exp.req.Endpoint = ftapi.GetEndpoint("experiences/"+strconv.Itoa(exp.ID), nil)
	exp.req.ExecuteMethod = func() {
		exp.req.Get(ctx, exp)
	}
	return &exp.req
}

func (exps *Experiences) GetAll(ctx context.Context) ftapi.CollectionRequest {
	exps.req.Endpoint = ftapi.GetEndpoint("experiences", nil)
	exps.req.ExecuteMethod = func() {
		exps.req.GetAll(ctx, &exps.Collection)
	}
	return &exps.req
}

func GetLevel(exp int) int {
	for level, expReq := range expLevels {
		if exp < expReq || expReq == -1.0 {
			return level - 1
		}
	}
	return 0
}

func (exps *Experiences) LevelAt(cursusID int, when time.Time) int {
	totalExp := 0
	for _, exp := range exps.Collection {
		if exp.CursusID != cursusID || exp.CreatedAt.After(when) {
			continue
		}
		totalExp += exp.Amount
	}
	return GetLevel(totalExp)
}
