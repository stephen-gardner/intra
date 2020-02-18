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
	Experience struct {
		ID                int       `json:"id"`
		UserID            int       `json:"user_id"`
		SkillID           int       `json:"skill_id"`
		ExperiancableID   int       `json:"experiancable_id"`
		ExperiancableType string    `json:"experiancable_type"`
		Amount            int       `json:"experience"`
		CreatedAt         time.Time `json:"created_at"`
		CursusID          int       `json:"cursus_id"`
		User              struct {
			ID    int    `json:"id"`
			Login string `json:"login"`
			URL   string `json:"url"`
		} `json:"user"`
		Skill struct {
			ID        int       `json:"id"`
			Name      string    `json:"name"`
			CreatedAt time.Time `json:"created_at"`
		} `json:"skill"`
		Cursus struct {
			ID        int       `json:"id"`
			CreatedAt time.Time `json:"created_at"`
			Name      string    `json:"name"`
			Slug      string    `json:"slug"`
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
			CursusIds   []int     `json:"cursus_ids"`
			MarkedAt    time.Time `json:"marked_at"`
			Marked      bool      `json:"marked"`
			RetriableAt time.Time `json:"retriable_at"`
		} `json:"experiancable"`
	}
	Experiences []Experience
)

func (exp *Experience) Create(ctx context.Context, bypassCache bool) error {
	endpoint := GetEndpoint("experiences", nil)
	params := url.Values{}
	params.Set("experience[user_id]", strconv.Itoa(exp.UserID))
	params.Set("experience[skill_id]", strconv.Itoa(exp.SkillID))
	params.Set("experience[experiancable_id]", strconv.Itoa(exp.ExperiancableID))
	params.Set("experience[experiancable_type]", exp.ExperiancableType)
	params.Set("experience[experience]", strconv.Itoa(exp.Amount))
	if !exp.CreatedAt.IsZero() {
		params.Set("experience[created_at]", exp.CreatedAt.Format(intraTimeFormat))
	}
	params.Set("experience[cursus_id]", strconv.Itoa(exp.CursusID))
	_, respData, err := RunRequest(GetClient(ctx, "public"), http.MethodPost, endpoint, params)
	if err == nil {
		if err = json.Unmarshal(respData, exp); err == nil && !bypassCache {
			cached := *exp
			intraCache.put(catExperiences, cached.ID, &cached)
		}
	}
	return err
}

func (exp *Experience) Delete(ctx context.Context) error {
	endpoint := GetEndpoint("experiences/"+strconv.Itoa(exp.ID), nil)
	_, _, err := RunRequest(GetClient(ctx, "public"), http.MethodDelete, endpoint, nil)
	if err == nil {
		intraCache.delete(catExperiences, exp.ID)
	}
	return err
}

func (exp *Experience) Get(ctx context.Context, bypassCache bool) error {
	if !bypassCache {
		if xp, present := intraCache.get(catExperiences, exp.ID); present {
			*exp = *xp.(*Experience)
			return nil
		}
	}
	experiences := &Experiences{}
	if err := experiences.GetAll(ctx, bypassCache, getSingleParams(exp.ID)); err != nil {
		return err
	}
	if len(*experiences) == 0 {
		return fmt.Errorf("experience %d does not exist", exp.ID)
	}
	*exp = (*experiences)[0]
	return nil
}

func (experiences *Experiences) GetForProjectsUser(
	ctx context.Context,
	bypassCache bool,
	projectsUserID int,
	params url.Values,
) error {
	endpoint := fmt.Sprintf("projects_users/%d/experiences", projectsUserID)
	return experiences.getAll(ctx, bypassCache, endpoint, params)
}

func (experiences *Experiences) GetAll(ctx context.Context, bypassCache bool, params url.Values) error {
	return experiences.getAll(ctx, bypassCache, "experiences", params)
}

func (experiences *Experiences) getAll(
	ctx context.Context,
	bypassCache bool,
	endpoint string,
	params url.Values,
) error {
	if err := GetAll(GetClient(ctx, "public"), endpoint, params, experiences); err != nil {
		return err
	}
	if !bypassCache {
		for _, exp := range *experiences {
			cached := exp
			intraCache.put(catExperiences, cached.ID, &cached)
		}
	}
	return nil
}
