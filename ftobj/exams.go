package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	Exam struct {
		req            ftapi.RequestData
		ID             int        `json:"id"`
		IPRange        string     `json:"ip_range"`
		BeginAt        ftapi.Time `json:"begin_at"`
		EndAt          ftapi.Time `json:"end_at"`
		Location       string     `json:"location"`
		MaxPeople      int        `json:"max_people"`
		NbrSubscribers int        `json:"nbr_subscribers"`
		Name           string     `json:"name"`
		Campus         struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			TimeZone string `json:"time_zone"`
			Language struct {
				ID         int        `json:"id"`
				Name       string     `json:"name"`
				Identifier string     `json:"identifier"`
				CreatedAt  ftapi.Time `json:"created_at"`
				UpdatedAt  ftapi.Time `json:"updated_at"`
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
		} `json:"campus"`
		Cursus []struct {
			ID        int        `json:"id"`
			CreatedAt ftapi.Time `json:"created_at"`
			Name      string     `json:"name"`
			Slug      string     `json:"slug"`
		} `json:"cursus"`
		Projects []struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Slug   string `json:"slug"`
			Parent struct {
				Name string `json:"name"`
				ID   int    `json:"id"`
				Slug string `json:"slug"`
				URL  string `json:"url"`
			} `json:"parent"`
			Children []struct {
				Name string `json:"name"`
				ID   int    `json:"id"`
				Slug string `json:"slug"`
				URL  string `json:"url"`
			} `json:"children"`
			Objectives []string   `json:"objectives"`
			CreatedAt  ftapi.Time `json:"created_at"`
			UpdatedAt  ftapi.Time `json:"updated_at"`
			Exam       bool       `json:"exam"`
		} `json:"projects"`
	}
	Exams struct {
		req        ftapi.RequestData
		Collection []Exam
	}
	ExamCUParams struct {
		Name       string     `json:"name,omitempty"`
		BeginAt    ftapi.Time `json:"begin_at,omitempty"`
		EndAt      ftapi.Time `json:"end_at,omitempty"`
		Location   string     `json:"location,omitempty"`
		IPRange    string     `json:"ip_range,omitempty"`
		Visible    bool       `json:"visible,omitempty"`
		MaxPeople  int        `json:"max_people,omitempty"`
		CampusID   int        `json:"campus_id,omitempty"`
		ProjectIDs []int      `json:"project_ids,omitempty"`
	}
)

func (exam *Exam) Create(ctx context.Context, params ExamCUParams) ftapi.CachedRequest {
	exam.req.Endpoint = ftapi.GetEndpoint("exams", nil)
	exam.req.ExecuteMethod = func() {
		data := ftapi.EncapsulatedMarshal("exam", params)
		exam.req.Create(ftapi.GetClient(ctx, "public"), exam, data)
	}
	return &exam.req
}

func (exam *Exam) Delete(ctx context.Context) ftapi.Request {
	exam.req.Endpoint = ftapi.GetEndpoint("exams/"+strconv.Itoa(exam.ID), nil)
	exam.req.ExecuteMethod = func() {
		exam.req.Delete(ftapi.GetClient(ctx, "public"), exam)
	}
	return &exam.req
}

func (exam *Exam) Patch(ctx context.Context, params ExamCUParams) ftapi.Request {
	exam.req.Endpoint = ftapi.GetEndpoint("exams/"+strconv.Itoa(exam.ID), nil)
	exam.req.ExecuteMethod = func() {
		data := ftapi.EncapsulatedMarshal("exam", params)
		exam.req.Patch(ftapi.GetClient(ctx, "public"), exam, data)
	}
	return &exam.req
}

func (exam *Exam) Get(ctx context.Context) ftapi.CachedRequest {
	exam.req.Endpoint = ftapi.GetEndpoint("exams/"+strconv.Itoa(exam.ID), nil)
	exam.req.ExecuteMethod = func() {
		exam.req.Get(ftapi.GetClient(ctx, "public"), exam)
	}
	return &exam.req
}

func (exams *Exams) GetAll(ctx context.Context) ftapi.CollectionRequest {
	exams.req.Endpoint = ftapi.GetEndpoint("exams", nil)
	exams.req.ExecuteMethod = func() {
		exams.req.GetAll(ftapi.GetClient(ctx, "public"), &exams.Collection)
	}
	return &exams.req
}
