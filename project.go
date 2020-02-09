package intra

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
)

type (
	Project struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Slug        string `json:"slug"`
		Description string `json:"description"`
		Exam        bool   `json:"exam"`
	}
	Projects []Project
)

func (project *Project) GetProject(ctx context.Context, bypassCache bool, ID int) error {
	IDStr := strconv.Itoa(ID)
	endpoint := GetEndpoint("projects/"+IDStr, nil)
	if !bypassCache {
		if proj, present := intraCache.get(endpoint); present {
			*project = proj.(Project)
			return nil
		}
	}
	params := url.Values{}
	params.Set("filter[id]", IDStr)
	params.Set("page[number]", "1")
	projects := &Projects{}
	err := projects.GetAllProjects(ctx, bypassCache, params)
	if err == nil && len(*projects) > 0 {
		*project = (*projects)[0]
	}
	return err
}

func (projects *Projects) GetAllProjects(ctx context.Context, bypassCache bool, params url.Values) error {
	data, err := GetAll(GetClient(ctx, "public"), "projects", params)
	if err != nil {
		return err
	}
	for _, dataPage := range data {
		var page Projects
		if err := json.Unmarshal(dataPage, &page); err != nil {
			return err
		}
		if !bypassCache {
			for _, proj := range page {
				endpoint := GetEndpoint("projects/"+strconv.Itoa(proj.ID), nil)
				intraCache.put(endpoint, proj)
			}
		}
		*projects = append(*projects, page...)
	}
	return nil
}
