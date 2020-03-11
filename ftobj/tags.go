package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	Tag struct {
		req        ftapi.RequestData
		ID         int         `json:"id,omitempty"`
		Name       string      `json:"name,omitempty"`
		Kind       string      `json:"kind,omitempty"`
		Users      []User      `json:"users,omitempty"`
		Subnotions []Subnotion `json:"subnotions,omitempty"`
	}
	Tags struct {
		req        ftapi.RequestData
		Collection []Tag
	}
	TagCUParams struct {
		Name                string `json:"name,omitempty"`
		TagsUsersAttributes []struct {
			ID     int `json:"id,omitempty"`
			UserID int `json:"user_id,omitempty"`
		} `json:"tags_users_attributes"`
		CursusIDs []int `json:"cursus_ids,omitempty"`
	}
)

func (tag *Tag) Create(ctx context.Context, params TagCUParams) ftapi.CachedRequest {
	tag.req.Endpoint = ftapi.GetEndpoint("tags", nil)
	tag.req.ExecuteMethod = func() {
		tag.req.Create(ctx, tag, ftapi.EncapsulatedMarshal("tag", params))
	}
	return &tag.req
}

func (tag *Tag) Delete(ctx context.Context) ftapi.Request {
	tag.req.Endpoint = ftapi.GetEndpoint("tags/"+strconv.Itoa(tag.ID), nil)
	tag.req.ExecuteMethod = func() {
		tag.req.Delete(ctx, tag)
	}
	return &tag.req
}

func (tag *Tag) Patch(ctx context.Context, params TagCUParams) ftapi.Request {
	tag.req.Endpoint = ftapi.GetEndpoint("tags/"+strconv.Itoa(tag.ID), nil)
	tag.req.ExecuteMethod = func() {
		tag.req.Patch(ctx, ftapi.EncapsulatedMarshal("tag", params))
	}
	return &tag.req
}

func (tag *Tag) Get(ctx context.Context) ftapi.CachedRequest {
	tag.req.Endpoint = ftapi.GetEndpoint("tags/"+strconv.Itoa(tag.ID), nil)
	tag.req.ExecuteMethod = func() {
		tag.req.Get(ctx, tag)
	}
	return &tag.req
}

func (tags *Tags) GetAll(ctx context.Context) ftapi.CollectionRequest {
	tags.req.Endpoint = ftapi.GetEndpoint("tags", nil)
	tags.req.ExecuteMethod = func() {
		tags.req.GetAll(ctx, &tags.Collection)
	}
	return &tags.req
}
