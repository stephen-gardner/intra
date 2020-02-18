package intra

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type (
	CursusUser struct {
		ID           int       `json:"id"`
		CursusID     int       `json:"cursus_id"`
		BeginAt      time.Time `json:"begin_at"`
		EndAt        time.Time `json:"end_at"`
		Grade        string    `json:"grade"`
		Level        float64   `json:"level"`
		HasCoalition bool      `json:"has_coalition"`
		Skills       []struct {
			ID    int     `json:"id"`
			Name  string  `json:"name"`
			Level float64 `json:"level"`
		} `json:"skills"`
		User struct {
			ID    int    `json:"id"`
			Login string `json:"login"`
			URL   string `json:"url"`
		} `json:"user"`
		Cursus struct {
			ID        int       `json:"id"`
			CreatedAt time.Time `json:"created_at"`
			Name      string    `json:"name"`
			Slug      string    `json:"slug"`
		} `json:"cursus"`
	}
	CursusUsers []CursusUser
)

func (cursusUser *CursusUser) Delete(ctx context.Context) error {
	endpoint := GetEndpoint("cursus_users/"+strconv.Itoa(cursusUser.ID), nil)
	_, _, err := RunRequest(GetClient(ctx, "public"), http.MethodDelete, endpoint, nil)
	if err == nil {
		intraCache.delete(catCursusUsers, cursusUser.ID)
	}
	return err
}

func (cursusUser *CursusUser) Get(ctx context.Context, bypassCache bool) error {
	if !bypassCache {
		if cu, present := intraCache.get(catCursusUsers, cursusUser.ID); present {
			*cursusUser = *cu.(*CursusUser)
			return nil
		}
	}
	cursusUsers := &CursusUsers{}
	if err := cursusUsers.GetAll(ctx, bypassCache, getSingleParams(cursusUser.ID)); err != nil {
		return err
	}
	if len(*cursusUsers) == 0 {
		return fmt.Errorf("cursus user %d does not exist", cursusUser.ID)
	}
	*cursusUser = (*cursusUsers)[0]
	return nil
}

func (cursusUsers *CursusUsers) GetAll(ctx context.Context, bypassCache bool, params url.Values) error {
	if err := GetAll(GetClient(ctx, "public"), "cursus_users", params, cursusUsers); err != nil {
		return err
	}
	if !bypassCache {
		for _, cu := range *cursusUsers {
			cached := cu
			intraCache.put(catCursusUsers, cached.ID, &cached)
		}
	}
	return nil
}
