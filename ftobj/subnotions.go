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
	SubnotionCUParams struct {
		Name                  string               `json:"name,omitempty"`
		NotionID              int                  `json:"notion_id,omitempty"`
		AttachmentsAttributes []AttachmentCUParams `json:"attachments_attributes,omitempty"`
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

func (sn *Subnotion) Create(ctx context.Context, params SubnotionCUParams) ftapi.CachedRequest {
	sn.req.Endpoint = ftapi.GetEndpoint("subnotions", nil)
	sn.req.ExecuteMethod = func() {
		sn.req.Create(ctx, sn, ftapi.EncapsulatedMarshal("subnotion", params))
	}
	return &sn.req
}

func (sn *Subnotion) Delete(ctx context.Context) ftapi.Request {
	sn.req.Endpoint = ftapi.GetEndpoint("subnotions/"+strconv.Itoa(sn.ID), nil)
	sn.req.ExecuteMethod = func() {
		sn.req.Delete(ctx, sn)
	}
	return &sn.req
}

func (sn *Subnotion) Patch(ctx context.Context, params SubnotionCUParams) ftapi.Request {
	sn.req.Endpoint = ftapi.GetEndpoint("subnotions/"+strconv.Itoa(sn.ID), nil)
	sn.req.ExecuteMethod = func() {
		sn.req.Patch(ctx, ftapi.EncapsulatedMarshal("subnotion", params))
	}
	return &sn.req
}

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
