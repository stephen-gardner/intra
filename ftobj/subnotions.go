package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
	"time"
)

type (
	Subnotion struct {
		req         ftapi.RequestData
		ID          int          `json:"id,omitempty"`
		Name        string       `json:"name,omitempty"`
		Slug        string       `json:"slug,omitempty"`
		CreatedAt   *ftapi.Time  `json:"created_at,omitempty"`
		Notepad     *Notepad     `json:"notepad,omitempty"`
		Attachments []Attachment `json:"attachments,omitempty"`
		Notion      *Notion      `json:"notion,omitempty"`
	}
	Subnotions struct {
		req        ftapi.RequestData
		Collection []Subnotion
	}
	Notepad struct {
		ID          int        `json:"id,omitempty"`
		UserID      int        `json:"user_id,omitempty"`
		Content     string     `json:"content,omitempty"`
		CreatedAt   *time.Time `json:"created_at,omitempty"`
		UpdatedAt   *time.Time `json:"updated_at,omitempty"`
		SubnotionID int        `json:"subnotion_id,omitempty"`
	}
)

func (sn *Subnotion) Get(ctx context.Context) ftapi.CachedRequest {
	sn.req.Endpoint = ftapi.GetEndpoint("subnotions/"+strconv.Itoa(sn.ID), nil)
	sn.req.ExecuteMethod = func() {
		sn.req.Get(ctx, sn)
	}
	return &sn.req
}

func (sns *Subnotions) GetAll(ctx context.Context) ftapi.CollectionRequest {
	sns.req.Endpoint = ftapi.GetEndpoint("subnotions", nil)
	sns.req.ExecuteMethod = func() {
		sns.req.GetAll(ctx, &sns.Collection)
	}
	return &sns.req
}
