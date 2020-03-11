package ftobj

import (
	"context"
	"encoding/json"
	"fmt"
	"intra/ftapi"
	"strconv"
	"strings"
)

type (
	// Polymorphic object
	Attachment struct {
		req      ftapi.RequestData
		ID       int       `json:"attachment_id,omitempty"`
		BaseID   int       `json:"base_id,omitempty"`
		Language *Language `json:"language,omitempty"`
		Type     string    `json:"type,omitempty"`
		Content  interface{}
	}
	Attachments struct {
		req        ftapi.RequestData
		Collection []Attachment
	}
	AttachmentCUParams struct {
		Kind                 string `json:"kind,omitempty"`
		LanguageID           int    `json:"language_id,omitempty"`
		Untranslatable       bool   `json:"untranslatable,omitempty"`
		AttachableAttributes []struct {
			Name  string `json:"name,omitempty"`
			Type  string `json:"attachable_type,omitempty"`
			PDF   string `json:"pdf,omitempty"`
			Video string `json:"video,omitempty"`
		} `json:"attachable_attributes"`
	}
	AttachmentCode struct {
		Code string `json:"code,omitempty"`
	}
	AttachmentDocument struct {
		URL  string `json:"url,omitempty"`
		Name string `json:"name,omitempty"`
	}
	AttachmentLink struct {
		URL   string `json:"url,omitempty"`
		Title string `json:"title,omitempty"`
	}
	AttachmentPDF struct {
		Name string `json:"name,omitempty"`
		PDF  struct {
			PDF struct {
				URL   string `json:"url,omitempty"`
				Thumb struct {
					URL string `json:"url,omitempty"`
				} `json:"thumb"`
			} `json:"pdf"`
		} `json:"pdf"`
		PageCount     int         `json:"page_count,omitempty"`
		CreatedAt     *ftapi.Time `json:"created_at,omitempty"`
		PDFProcessing bool        `json:"pdf_processing,omitempty"`
		Slug          string      `json:"slug,omitempty"`
		URL           string      `json:"url,omitempty"`
		ThumbURL      string      `json:"thumb_url,omitempty"`
	}
	AttachmentVideo struct {
		Name            string      `json:"name,omitempty"`
		Duration        int         `json:"duration,omitempty"`
		VideoProcessing bool        `json:"video_processing,omitempty"`
		CreatedAt       *ftapi.Time `json:"created_at,omitempty"`
		Slug            string      `json:"slug,omitempty"`
		URLs            struct {
			URL    string   `json:"url,omitempty"`
			LowD   string   `json:"low_d,omitempty"`
			HighD  string   `json:"high_d,omitempty"`
			Thumbs []string `json:"thumbs,omitempty"`
		} `json:"urls"`
	}
)

const (
	AttachmentKindCode     = "code"
	AttachmentKindDocument = "document"
	AttachmentKindLink     = "link"
	AttachmentKindPDF      = "pdf"
	AttachmentKindVideo    = "video"
)

func (attach *Attachment) UnmarshalJSON(data []byte) error {
	mapped := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &mapped); err != nil {
		return err
	}
	_ = json.Unmarshal(mapped["id"], &attach.ID)
	_ = json.Unmarshal(mapped["base_id"], &attach.BaseID)
	_ = json.Unmarshal(mapped["type"], &attach.Type)
	if raw, present := mapped["language"]; present {
		attach.Language = &Language{}
		_ = json.Unmarshal(raw, &attach.Language)
	}
	switch strings.ToLower(attach.Type) {
	case "code":
		attach.Content = &AttachmentCode{}
	case "document":
		attach.Content = &AttachmentDocument{}
	case "link":
		attach.Content = &AttachmentLink{}
	case "pdf":
		attach.Content = &AttachmentPDF{}
	case "video":
		attach.Content = &AttachmentVideo{}
	default:
		return fmt.Errorf("unknown attachment type: %s", attach.Type)
	}
	return json.Unmarshal(data, attach.Content)
}

func (attach *Attachment) Create(ctx context.Context, params AttachmentCUParams) ftapi.CachedRequest {
	attach.req.Endpoint = ftapi.GetEndpoint("attachments", nil)
	attach.req.ExecuteMethod = func() {
		attach.req.Create(ctx, attach, ftapi.EncapsulatedMarshal("attachment", params))
	}
	return &attach.req
}

func (attach *Attachment) Delete(ctx context.Context) ftapi.Request {
	attach.req.Endpoint = ftapi.GetEndpoint("attachments/"+strconv.Itoa(attach.ID), nil)
	attach.req.ExecuteMethod = func() {
		attach.req.Delete(ctx, attach)
	}
	return &attach.req
}

func (attach *Attachment) Patch(ctx context.Context, params AttachmentCUParams) ftapi.Request {
	attach.req.Endpoint = ftapi.GetEndpoint("attachments/"+strconv.Itoa(attach.ID), nil)
	attach.req.ExecuteMethod = func() {
		attach.req.Patch(ctx, ftapi.EncapsulatedMarshal("attachment", params))
	}
	return &attach.req
}

func (attach *Attachment) Get(ctx context.Context) ftapi.CachedRequest {
	attach.req.Endpoint = ftapi.GetEndpoint("attachments/"+strconv.Itoa(attach.ID), nil)
	attach.req.ExecuteMethod = func() {
		attach.req.Get(ctx, attach)
	}
	return &attach.req
}

func (attachments *Attachments) GetAll(ctx context.Context) ftapi.CollectionRequest {
	attachments.req.Endpoint = ftapi.GetEndpoint("attachments", nil)
	attachments.req.ExecuteMethod = func() {
		attachments.req.GetAll(ctx, &attachments.Collection)
	}
	return &attachments.req
}
