package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	Campus struct {
		req      ftapi.RequestData
		ID       int    `json:"id"`
		Name     string `json:"name"`
		TimeZone string `json:"time_zone"`
		Language struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			Identifier string `json:"identifier"`
		} `json:"language"`
		UsersCount  int    `json:"users_count"`
		VogsphereID int    `json:"vogsphere_id"`
		Country     string `json:"country"`
		Address     string `json:"address"`
		Zip         string `json:"zip"`
		City        string `json:"city"`
		Website     string `json:"website"`
		Facebook    string `json:"facebook"`
		Twitter     string `json:"twitter"`
		Endpoint    struct {
			ID          int        `json:"id"`
			URL         string     `json:"url"`
			Description string     `json:"description"`
			CreatedAt   ftapi.Time `json:"created_at"`
			UpdatedAt   ftapi.Time `json:"updated_at"`
		} `json:"endpoint"`
	}
	Campuses struct {
		req        ftapi.RequestData
		Collection []Campus
	}
	CampusCUParams struct {
		Name               string `json:"name,omitempty"`
		Slug               string `json:"slug,omitempty"`
		DisplayName        string `json:"display_name,omitempty"`
		TimeZone           string `json:"time_zone,omitempty"`
		LanguageID         int    `json:"language_id,omitempty"`
		EmailExtension     string `json:"email_extension,omitempty"`
		MainEmail          string `json:"main_email,omitempty"`
		EndpointID         int    `json:"endpoint_id,omitempty"`
		VogsphereID        int    `json:"vogsphere_id,omitempty"`
		ContentEmail       string `json:"content_email,omitempty"`
		LaunchDate         string `json:"time_of_community_service_started,omitempty"`
		CompaniesMail      string `json:"companies_mail,omitempty"`
		Address            string `json:"address,omitempty"`
		Zip                string `json:"zip,omitempty"`
		City               string `json:"city,omitempty"`
		Country            string `json:"country,omitempty"`
		ProNeedsValidation bool   `json:"pro_needs_validation,omitempty"`
		Logo               string `json:"logo,omitempty"`
		Website            string `json:"website,omitempty"`
		Facebook           string `json:"facebook,omitempty"`
		Twitter            string `json:"twitter,omitempty"`
		HelpURL            string `json:"help_url,omitempty"`
	}
)

func (campus *Campus) Create(ctx context.Context, params CampusCUParams) ftapi.CachedRequest {
	campus.req.Endpoint = ftapi.GetEndpoint("campus", nil)
	campus.req.ExecuteMethod = func() {
		campus.req.Create(ctx, campus, ftapi.EncapsulatedMarshal("campus", params))
	}
	return &campus.req
}

func (campus *Campus) Patch(ctx context.Context, params CampusCUParams) ftapi.Request {
	campus.req.Endpoint = ftapi.GetEndpoint("campus/"+strconv.Itoa(campus.ID), nil)
	campus.req.ExecuteMethod = func() {
		data := ftapi.EncapsulatedMarshal("campus", params)
		campus.req.Patch(ctx, data)
	}
	return &campus.req
}

func (campus *Campus) Get(ctx context.Context) ftapi.CachedRequest {
	campus.req.Endpoint = ftapi.GetEndpoint("campus/"+strconv.Itoa(campus.ID), nil)
	campus.req.ExecuteMethod = func() {
		campus.req.Get(ctx, campus)
	}
	return &campus.req
}

func (campuses *Campuses) GetAll(ctx context.Context) ftapi.CollectionRequest {
	campuses.req.Endpoint = ftapi.GetEndpoint("campus", nil)
	campuses.req.ExecuteMethod = func() {
		campuses.req.GetAll(ctx, &campuses.Collection)
	}
	return &campuses.req
}
