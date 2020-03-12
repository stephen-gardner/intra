package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	Exam struct {
		req            ftapi.RequestData
		ID             int         `json:"id,omitempty"`
		IPRange        string      `json:"ip_range,omitempty"`
		BeginAt        *ftapi.Time `json:"begin_at,omitempty"`
		EndAt          *ftapi.Time `json:"end_at,omitempty"`
		Location       string      `json:"location,omitempty"`
		MaxPeople      int         `json:"max_people,omitempty"`
		NbrSubscribers int         `json:"nbr_subscribers,omitempty"`
		Name           string      `json:"name,omitempty"`
		Campus         *Campus     `json:"campus,omitempty"`
		Cursus         []Cursus    `json:"cursus,omitempty"`
		Projects       []Project   `json:"projects,omitempty"`
	}
	Exams struct {
		req        ftapi.RequestData
		Collection []Exam
	}
	ExamCUParams struct {
		Name       string      `json:"name,omitempty"`
		BeginAt    *ftapi.Time `json:"begin_at,omitempty"`
		EndAt      *ftapi.Time `json:"end_at,omitempty"`
		Location   string      `json:"location,omitempty"`
		IPRange    string      `json:"ip_range,omitempty"`
		Visible    bool        `json:"visible,omitempty"`
		MaxPeople  int         `json:"max_people,omitempty"`
		CampusID   int         `json:"campus_id,omitempty"`
		ProjectIDs []int       `json:"project_ids,omitempty"`
	}
)

func (exam *Exam) Create(ctx context.Context, params ExamCUParams) ftapi.CachedRequest {
	exam.req.Endpoint = ftapi.GetEndpoint("exams", nil)
	exam.req.ExecuteMethod = func() {
		exam.req.Create(ctx, exam, ftapi.EncapsulatedMarshal("exam", params))
	}
	return &exam.req
}

func (exam *Exam) Delete(ctx context.Context) ftapi.Request {
	exam.req.Endpoint = ftapi.GetEndpoint("exams/"+strconv.Itoa(exam.ID), nil)
	exam.req.ExecuteMethod = func() {
		exam.req.Delete(ctx, exam)
	}
	return &exam.req
}

func (exam *Exam) Patch(ctx context.Context, params ExamCUParams) ftapi.Request {
	exam.req.Endpoint = ftapi.GetEndpoint("exams/"+strconv.Itoa(exam.ID), nil)
	exam.req.ExecuteMethod = func() {
		exam.req.Patch(ctx, ftapi.EncapsulatedMarshal("exam", params))
	}
	return &exam.req
}

func (exam *Exam) Get(ctx context.Context) ftapi.CachedRequest {
	exam.req.Endpoint = ftapi.GetEndpoint("exams/"+strconv.Itoa(exam.ID), nil)
	exam.req.ExecuteMethod = func() {
		exam.req.Get(ctx, exam)
	}
	return &exam.req
}

func (exams *Exams) GetAll(ctx context.Context) ftapi.CollectionRequest {
	exams.req.Endpoint = ftapi.GetEndpoint("exams", nil)
	exams.req.ExecuteMethod = func() {
		exams.req.GetAll(ctx, &exams.Collection)
	}
	return &exams.req
}
