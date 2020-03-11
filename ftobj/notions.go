package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	Notion struct {
		req        ftapi.RequestData
		ID         int         `json:"id,omitempty"`
		Name       string      `json:"name,omitempty"`
		Slug       string      `json:"slug,omitempty"`
		CreatedAt  *ftapi.Time `json:"created_at,omitempty"`
		Subnotions []Subnotion `json:"subnotions,omitempty"`
		Tags       []Tag       `json:"tags,omitempty"`
		Cursus     []Cursus    `json:"cursus,omitempty"`
	}
	Notions struct {
		req        ftapi.RequestData
		Collection []Notion
	}
	NotionCUParams struct {
		Name                 string `json:"name,omitempty"`
		TagIDs               []int  `json:"tag_ids,omitempty"`
		CursusIDs            []int  `json:"cursus_ids,omitempty"`
		SubnotionsAttributes []struct {
			ID                    int                `json:"id,omitempty"`
			Name                  string             `json:"name,omitempty"`
			AttachmentsAttributes AttachmentCUParams `json:"attachments_attributes,omitempty"`
		} `json:"subnotions_attributes"`
	}
)

func (notion *Notion) Create(ctx context.Context, params NotionCUParams) ftapi.CachedRequest {
	notion.req.Endpoint = ftapi.GetEndpoint("notions", nil)
	notion.req.ExecuteMethod = func() {
		notion.req.Create(ctx, notion, ftapi.EncapsulatedMarshal("notion", params))
	}
	return &notion.req
}

func (notion *Notion) Delete(ctx context.Context) ftapi.Request {
	notion.req.Endpoint = ftapi.GetEndpoint("notions/"+strconv.Itoa(notion.ID), nil)
	notion.req.ExecuteMethod = func() {
		notion.req.Delete(ctx, notion)
	}
	return &notion.req
}

func (notion *Notion) Patch(ctx context.Context, params NotionCUParams) ftapi.Request {
	notion.req.Endpoint = ftapi.GetEndpoint("notions/"+strconv.Itoa(notion.ID), nil)
	notion.req.ExecuteMethod = func() {
		notion.req.Patch(ctx, ftapi.EncapsulatedMarshal("notion", params))
	}
	return &notion.req
}

func (notion *Notion) Get(ctx context.Context) ftapi.CachedRequest {
	notion.req.Endpoint = ftapi.GetEndpoint("notions/"+strconv.Itoa(notion.ID), nil)
	notion.req.ExecuteMethod = func() {
		notion.req.Get(ctx, notion)
	}
	return &notion.req
}

func (notions *Notions) GetAll(ctx context.Context) ftapi.CollectionRequest {
	notions.req.Endpoint = ftapi.GetEndpoint("notions", nil)
	notions.req.ExecuteMethod = func() {
		notions.req.GetAll(ctx, &notions.Collection)
	}
	return &notions.req
}
