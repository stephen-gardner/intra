package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	User struct {
		req             ftapi.RequestData
		ID              int              `json:"id,omitempty"`
		Email           string           `json:"email,omitempty"`
		Login           string           `json:"login,omitempty"`
		FirstName       string           `json:"first_name,omitempty"`
		LastName        string           `json:"last_name,omitempty"`
		URL             string           `json:"url,omitempty"`
		Phone           string           `json:"phone,omitempty"`
		DisplayName     string           `json:"displayname,omitempty"`
		ImageURL        string           `json:"image_url,omitempty"`
		Staff           bool             `json:"staff?,omitempty"`
		CorrectionPoint int              `json:"correction_point,omitempty"`
		PoolMonth       string           `json:"pool_month,omitempty"`
		PoolYear        string           `json:"pool_year,omitempty"`
		Location        string           `json:"location,omitempty"`
		Wallet          int              `json:"wallet,omitempty"`
		Groups          []Group          `json:"groups,omitempty"`
		CursusUsers     []CursusUser     `json:"cursus_users,omitempty"`
		ProjectsUsers   []ProjectsUser   `json:"projects_users,omitempty"`
		LanguagesUsers  []LanguagesUser  `json:"languages_users,omitempty"`
		Achievements    []Achievement    `json:"achievements,omitempty"`
		Titles          []Title          `json:"titles,omitempty"`
		TitlesUsers     []TitlesUser     `json:"titles_users,omitempty"`
		Partnerships    []Partnership    `json:"partnerships,omitempty"`
		Patroned        []Patronage      `json:"patroned,omitempty"`
		Patroning       []Patronage      `json:"patroning,omitempty"`
		ExpertisesUsers []ExpertisesUser `json:"expertises_users,omitempty"`
		Campus          []Campus         `json:"campus,omitempty"`
		CampusUsers     []CampusUser     `json:"campus_users,omitempty"`
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
