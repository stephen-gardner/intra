package ftobj

import (
	"context"
	"intra/ftapi"
	"strconv"
)

type (
	RolesEntity struct {
		req        ftapi.RequestData
		ID         int        `json:"id,omitempty"`
		EntityID   int        `json:"entity_id,omitempty"`
		EntityType string     `json:"entity_type,omitempty"`
		CreatedAt  ftapi.Time `json:"created_at,omitempty"`
		ExpiresAt  ftapi.Time `json:"expires_at,omitempty"`
		Entity     struct {
			ID          int        `json:"id,omitempty"`
			Name        string     `json:"name,omitempty"`
			Description string     `json:"description,omitempty"`
			Image       string     `json:"image,omitempty"`
			Website     string     `json:"website,omitempty"`
			Public      bool       `json:"public,omitempty"`
			Scopes      []string   `json:"scopes,omitempty"`
			CreatedAt   ftapi.Time `json:"created_at,omitempty"`
			UpdatedAt   ftapi.Time `json:"updated_at,omitempty"`
			Owner       struct {
				ID    int    `json:"id,omitempty"`
				Login string `json:"login,omitempty"`
				URL   string `json:"url,omitempty"`
			} `json:"owner,omitempty"`
			RateLimit int `json:"rate_limit,omitempty"`
		} `json:"entity,omitempty"`
		Role struct {
			ID          int    `json:"id,omitempty"`
			Name        string `json:"name,omitempty"`
			Description string `json:"description,omitempty"`
		} `json:"role,omitempty"`
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
		entity.req.Create(ctx, entity, ftapi.EncapsulatedMarshal("roles_entity", params))
	}
	return &entity.req
}

func (entity *RolesEntity) Delete(ctx context.Context) ftapi.Request {
	entity.req.Endpoint = ftapi.GetEndpoint("roles_entities/"+strconv.Itoa(entity.ID), nil)
	entity.req.ExecuteMethod = func() {
		entity.req.Delete(ctx, entity)
	}
	return &entity.req
}

func (entity *RolesEntity) Patch(ctx context.Context, params RolesEntityCUParams) ftapi.Request {
	entity.req.Endpoint = ftapi.GetEndpoint("roles_entities/"+strconv.Itoa(entity.ID), nil)
	entity.req.ExecuteMethod = func() {
		entity.req.Patch(ctx, ftapi.EncapsulatedMarshal("roles_entity", params))
	}
	return &entity.req
}

func (entity *RolesEntity) Get(ctx context.Context) ftapi.CachedRequest {
	entity.req.Endpoint = ftapi.GetEndpoint("roles_entities/"+strconv.Itoa(entity.ID), nil)
	entity.req.ExecuteMethod = func() {
		entity.req.Get(ctx, entity)
	}
	return &entity.req
}

func (entities *RolesEntities) GetAll(ctx context.Context) ftapi.CollectionRequest {
	entities.req.Endpoint = ftapi.GetEndpoint("roles_entities", nil)
	entities.req.ExecuteMethod = func() {
		entities.req.GetAll(ctx, &entities.Collection)
	}
	return &entities.req
}
