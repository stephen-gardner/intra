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
)

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
