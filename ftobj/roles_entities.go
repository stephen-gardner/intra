package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	RolesEntity struct {
		req        ftapi.RequestData
		ID         int        `json:"id"`
		EntityID   int        `json:"entity_id"`
		EntityType string     `json:"entity_type"`
		CreatedAt  ftapi.Time `json:"created_at"`
		ExpiresAt  ftapi.Time `json:"expires_at"`
		Entity     struct {
			ID          int        `json:"id"`
			Name        string     `json:"name"`
			Description string     `json:"description"`
			Image       string     `json:"image"`
			Website     string     `json:"website"`
			Public      bool       `json:"public"`
			Scopes      []string   `json:"scopes"`
			CreatedAt   ftapi.Time `json:"created_at"`
			UpdatedAt   ftapi.Time `json:"updated_at"`
			Owner       struct {
				ID    int    `json:"id"`
				Login string `json:"login"`
				URL   string `json:"url"`
			} `json:"owner"`
			RateLimit int `json:"rate_limit"`
		} `json:"entity"`
		Role struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"role"`
	}
	RolesEntities struct {
		req        ftapi.RequestData
		Collection []RolesEntity
	}
	RolesEntityCUParams struct {
		RoleID     int        `json:"role_id,omitempty"`
		EntityID   int        `json:"entity_id,omitempty"`
		ExpiresAt  ftapi.Time `json:"expires_at,omitempty"`
		EntityType string     `json:"entity_type,omitempty"`
	}
)

const (
	RolesEntityTypeApplication = "Doorkeeper::Application"
	RolesEntityTypeUser        = "User"
)

func (entity *RolesEntity) Create(ctx context.Context, params RolesEntityCUParams) ftapi.CachedRequest {
	entity.req.Endpoint = ftapi.GetEndpoint("roles_entities", nil)
	entity.req.ExecuteMethod = func() {
		data := ftapi.EncapsulatedMarshal("roles_entity", params)
		entity.req.Create(ftapi.GetClient(ctx, "public"), entity, data)
	}
	return &entity.req
}

func (entity *RolesEntity) Delete(ctx context.Context) ftapi.Request {
	entity.req.Endpoint = ftapi.GetEndpoint("roles_entities/"+strconv.Itoa(entity.ID), nil)
	entity.req.ExecuteMethod = func() {
		entity.req.Delete(ftapi.GetClient(ctx, "public"), entity)
	}
	return &entity.req
}

func (entity *RolesEntity) Patch(ctx context.Context, params RolesEntityCUParams) ftapi.Request {
	entity.req.Endpoint = ftapi.GetEndpoint("roles_entities/"+strconv.Itoa(entity.ID), nil)
	entity.req.ExecuteMethod = func() {
		data := ftapi.EncapsulatedMarshal("roles_entity", params)
		entity.req.Patch(ftapi.GetClient(ctx, "public"), entity, data)
	}
	return &entity.req
}

func (entity *RolesEntity) Get(ctx context.Context) ftapi.CachedRequest {
	entity.req.Endpoint = ftapi.GetEndpoint("roles_entities/"+strconv.Itoa(entity.ID), nil)
	entity.req.ExecuteMethod = func() {
		entity.req.Get(ftapi.GetClient(ctx, "public"), entity)
	}
	return &entity.req
}

func (entities *RolesEntities) GetAll(ctx context.Context) ftapi.CollectionRequest {
	entities.req.Endpoint = ftapi.GetEndpoint("roles_entities", nil)
	entities.req.ExecuteMethod = func() {
		entities.req.GetAll(ftapi.GetClient(ctx, "public"), &entities.Collection)
	}
	return &entities.req
}
