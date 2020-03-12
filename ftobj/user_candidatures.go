package ftobj

import (
	"context"
	"fmt"
	"intra/ftapi"
	"strconv"
)

type (
	UserCandidature struct {
		req                ftapi.RequestData
		ID                 int         `json:"id,omitempty"`
		UserID             int         `json:"user_id,omitempty"`
		BirthDate          string      `json:"birth_date,omitempty"`
		Gender             string      `json:"gender,omitempty"`
		ZipCode            string      `json:"zip_code,omitempty"`
		Country            string      `json:"country,omitempty"`
		BirthCity          string      `json:"birth_city,omitempty"`
		BirthCountry       string      `json:"birth_country,omitempty"`
		PostalStreet       string      `json:"postal_street,omitempty"`
		PostalComplement   string      `json:"postal_complement,omitempty"`
		PostalCity         string      `json:"postal_city,omitempty"`
		PostalZipCode      string      `json:"postal_zip_code,omitempty"`
		PostalCountry      string      `json:"postal_country,omitempty"`
		ContactAffiliation string      `json:"contact_affiliation,omitempty"`
		ContactLastName    string      `json:"contact_last_name,omitempty"`
		ContactFirstName   string      `json:"contact_first_name,omitempty"`
		ContactPhone1      string      `json:"contact_phone1,omitempty"`
		ContactPhone2      string      `json:"contact_phone2,omitempty"`
		MaxLevelMemory     int         `json:"max_level_memory,omitempty"`
		MaxLevelLogic      int         `json:"max_level_logic,omitempty"`
		OtherInformation   string      `json:"other_information,omitempty"`
		Language           string      `json:"language,omitempty"`
		MeetingDate        *ftapi.Time `json:"meeting_date,omitempty"`
		PiscineDate        string      `json:"piscine_date,omitempty"`
		CreatedAt          *ftapi.Time `json:"created_at,omitempty"`
		UpdatedAt          *ftapi.Time `json:"updated_at,omitempty"`
		Phone              string      `json:"phone,omitempty"`
		Email              string      `json:"email,omitempty"`
		Pin                string      `json:"pin,omitempty"`
		PhoneCountryCode   string      `json:"phone_country_code,omitempty"`
		HiddenPhone        bool        `json:"hidden_phone,omitempty"`
	}
	UserCandidatures struct {
		req        ftapi.RequestData
		Collection []UserCandidature
	}
	UserCandidatureCUParams struct {
		UserID             int         `json:"user_id,omitempty"`
		BirthDate          string      `json:"birth_date,omitempty"`
		Gender             string      `json:"gender,omitempty"`
		ZipCode            string      `json:"zip_code,omitempty"`
		Country            string      `json:"country,omitempty"`
		PhoneCountryCode   string      `json:"phone_country_code,omitempty"`
		BirthCity          string      `json:"birth_city,omitempty"`
		BirthCountry       string      `json:"birth_country,omitempty"`
		PostalStreet       string      `json:"postal_street,omitempty"`
		PostalComplement   string      `json:"postal_complement,omitempty"`
		PostalCity         string      `json:"postal_city,omitempty"`
		PostalZipCode      string      `json:"postal_zip_code,omitempty"`
		PostalCountry      string      `json:"postal_country,omitempty"`
		ContactAffiliation string      `json:"contact_affiliation,omitempty"`
		ContactLastName    string      `json:"contact_last_name,omitempty"`
		ContactFirstName   string      `json:"contact_first_name,omitempty"`
		ContactPhone1      string      `json:"contact_phone1,omitempty"`
		ContactPhone2      string      `json:"contact_phone2,omitempty"`
		MaxLevelMemory     int         `json:"max_level_memory,omitempty"`
		MaxLevelLogic      int         `json:"max_level_logic,omitempty"`
		OtherInformation   string      `json:"other_information,omitempty"`
		Language           string      `json:"language,omitempty"`
		MeetingDate        *ftapi.Time `json:"meeting_date,omitempty"`
		PiscineDate        string      `json:"piscine_date,omitempty"`
		Email              string      `json:"email,omitempty"`
		Pin                string      `json:"pin,omitempty"`
		Phone              string      `json:"phone,omitempty"`
	}
)

func (uc *UserCandidature) Create(ctx context.Context, params UserCandidatureCUParams) ftapi.CachedRequest {
	uc.req.Endpoint = ftapi.GetEndpoint("user_candidatures", nil)
	uc.req.ExecuteMethod = func() {
		uc.req.Create(ctx, uc, ftapi.EncapsulatedMarshal("user_candidature", params))
	}
	return &uc.req
}

func (uc *UserCandidature) Delete(ctx context.Context) ftapi.Request {
	uc.req.Endpoint = ftapi.GetEndpoint("user_candidatures/"+strconv.Itoa(uc.ID), nil)
	uc.req.ExecuteMethod = func() {
		uc.req.Delete(ctx, uc)
	}
	return &uc.req
}

func (uc *UserCandidature) Patch(ctx context.Context, params UserCandidatureCUParams) ftapi.Request {
	uc.req.Endpoint = ftapi.GetEndpoint("user_candidatures/"+strconv.Itoa(uc.ID), nil)
	uc.req.ExecuteMethod = func() {
		uc.req.Patch(ctx, ftapi.EncapsulatedMarshal("user_candidature", params))
	}
	return &uc.req
}

func (uc *UserCandidature) Get(ctx context.Context) ftapi.CachedRequest {
	uc.req.Endpoint = ftapi.GetEndpoint("user_candidatures/"+strconv.Itoa(uc.ID), nil)
	uc.req.ExecuteMethod = func() {
		uc.req.Get(ctx, uc)
	}
	return &uc.req
}

func (ucs *UserCandidatures) GetAll(ctx context.Context) ftapi.CollectionRequest {
	ucs.req.Endpoint = ftapi.GetEndpoint("user_candidatures", nil)
	ucs.req.ExecuteMethod = func() {
		ucs.req.GetAll(ctx, &ucs.Collection)
	}
	return &ucs.req
}

func (uc *UserCandidature) GetForUser(ctx context.Context, userID int) ftapi.CachedRequest {
	uc.req.Endpoint = ftapi.GetEndpoint(fmt.Sprintf("users/%d/user_candidature", userID), nil)
	uc.req.ExecuteMethod = func() {
		uc.req.Get(ctx, uc)
	}
	return &uc.req
}
