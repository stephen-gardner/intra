package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	TeamsUpload struct {
		req       ftapi.RequestData
		ID        int         `json:"id,omitempty"`
		FinalMark int         `json:"final_mark,omitempty"`
		Comment   string      `json:"comment,omitempty"`
		CreatedAt *ftapi.Time `json:"created_at,omitempty"`
		UploadID  int         `json:"upload_id,omitempty"`
		Upload    struct {
			ID           int         `json:"id,omitempty"`
			EvaluationID int         `json:"evaluation_id,omitempty"`
			Name         string      `json:"name,omitempty"`
			Description  string      `json:"description,omitempty"`
			CreatedAt    *ftapi.Time `json:"created_at,omitempty"`
			UpdatedAt    *ftapi.Time `json:"updated_at,omitempty"`
		} `json:"upload"`
	}
	TeamsUploads struct {
		req        ftapi.RequestData
		Collection []TeamsUpload
	}
	TeamsUploadCUParams struct {
		TeamID    int    `json:"team_id,omitempty"`
		UploadID  int    `json:"upload_id,omitempty"`
		FinalMark int    `json:"final_mark,omitempty"`
		Comment   string `json:"comment,omitempty"`
	}
)

func (tu *TeamsUpload) Create(ctx context.Context, params TeamsUploadCUParams) ftapi.CachedRequest {
	tu.req.Endpoint = ftapi.GetEndpoint("teams_uploads", nil)
	tu.req.ExecuteMethod = func() {
		tu.req.Create(ctx, tu, ftapi.EncapsulatedMarshal("teams_upload", params))
	}
	return &tu.req
}

func (tu *TeamsUpload) Delete(ctx context.Context) ftapi.Request {
	tu.req.Endpoint = ftapi.GetEndpoint("teams_uploads/"+strconv.Itoa(tu.ID), nil)
	tu.req.ExecuteMethod = func() {
		tu.req.Delete(ctx, tu)
	}
	return &tu.req
}

func (tu *TeamsUpload) Patch(ctx context.Context, params TeamsUploadCUParams) ftapi.Request {
	tu.req.Endpoint = ftapi.GetEndpoint("teams_uploads/"+strconv.Itoa(tu.ID), nil)
	tu.req.ExecuteMethod = func() {
		tu.req.Patch(ctx, ftapi.EncapsulatedMarshal("teams_upload", params))
	}
	return &tu.req
}

func (tu *TeamsUpload) Get(ctx context.Context) ftapi.CachedRequest {
	tu.req.Endpoint = ftapi.GetEndpoint("teams_uploads/"+strconv.Itoa(tu.ID), nil)
	tu.req.ExecuteMethod = func() {
		tu.req.Get(ctx, tu)
	}
	return &tu.req
}

func (tus *TeamsUploads) GetAll(ctx context.Context) ftapi.CollectionRequest {
	tus.req.Endpoint = ftapi.GetEndpoint("teams_uploads", nil)
	tus.req.ExecuteMethod = func() {
		tus.req.GetAll(ctx, &tus.Collection)
	}
	return &tus.req
}
