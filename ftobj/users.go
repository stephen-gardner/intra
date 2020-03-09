package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	User struct {
		req             ftapi.RequestData
		ID              int              `json:"id"`
		Email           string           `json:"email"`
		Login           string           `json:"login"`
		FirstName       string           `json:"first_name"`
		LastName        string           `json:"last_name"`
		URL             string           `json:"url"`
		Phone           string           `json:"phone"`
		DisplayName     string           `json:"displayname"`
		ImageURL        string           `json:"image_url"`
		Staff           bool             `json:"staff?"`
		CorrectionPoint int              `json:"correction_point"`
		PoolMonth       string           `json:"pool_month"`
		PoolYear        string           `json:"pool_year"`
		Location        string           `json:"location"`
		Wallet          int              `json:"wallet"`
		Groups          []Group          `json:"groups"`
		CursusUsers     []CursusUser     `json:"cursus_users"`
		ProjectsUsers   []ProjectsUser   `json:"projects_users"`
		LanguagesUsers  []LanguagesUser  `json:"languages_users"`
		Achievements    []Achievement    `json:"achievements"`
		Titles          []Title          `json:"titles"`
		TitlesUsers     []TitlesUser     `json:"titles_users"`
		Partnerships    []Partnership    `json:"partnerships"`
		Patroned        []Patronage      `json:"patroned"`
		Patroning       []Patronage      `json:"patroning"`
		ExpertisesUsers []ExpertisesUser `json:"expertises_users"`
		Campus          []Campus         `json:"campus"`
		CampusUsers     []CampusUser     `json:"campus_users"`
	}
	Users struct {
		req        ftapi.RequestData
		Collection []User
	}
	UserCUParams struct {
		Login           string `json:"login,omitempty"`
		Email           string `json:"email,omitempty"`
		FirstName       string `json:"first_name,omitempty"`
		LastName        string `json:"last_name,omitempty"`
		Password        string `json:"password,omitempty"`
		PoolYear        string `json:"pool_year,omitempty"`
		PoolMonth       string `json:"pool_month,omitempty"`
		Kind            string `json:"kind,omitempty"`
		Status          string `json:"status,omitempty"`
		Image           string `json:"image,omitempty"`
		CampusID        string `json:"campus_id,omitempty"`
		EmailStop       bool   `json:"email_stop,omitempty"`
		SkipWelcomeMail string `json:"skip_welcome_mail,omitempty"`
	}
)

const (
	UserKindAdmin    = "admin"
	UserKindExternal = "external"
	UserKindStudent  = "student"
)

func (user *User) Create(ctx context.Context, params UserCUParams) ftapi.CachedRequest {
	user.req.Endpoint = ftapi.GetEndpoint("users", nil)
	user.req.ExecuteMethod = func() {
		user.req.Create(ctx, user, ftapi.EncapsulatedMarshal("user", params))
	}
	return &user.req
}

func (user *User) Patch(ctx context.Context, params UserCUParams) ftapi.Request {
	user.req.Endpoint = ftapi.GetEndpoint("users/"+strconv.Itoa(user.ID), nil)
	user.req.ExecuteMethod = func() {
		user.req.Patch(ctx, ftapi.EncapsulatedMarshal("user", params))
	}
	return &user.req
}

func (user *User) Get(ctx context.Context) ftapi.CachedRequest {
	var ID string
	if user.ID != 0 {
		ID = strconv.Itoa(user.ID)
	} else {
		ID = user.Login
	}
	user.req.Endpoint = ftapi.GetEndpoint("users/"+ID, nil)
	user.req.ExecuteMethod = func() {
		user.req.Get(ctx, user)
	}
	return &user.req
}

func (users *Users) GetAll(ctx context.Context) ftapi.CollectionRequest {
	users.req.Endpoint = ftapi.GetEndpoint("users", nil)
	users.req.ExecuteMethod = func() {
		users.req.GetAll(ctx, &users.Collection)
	}
	return &users.req
}
