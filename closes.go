package intra

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type (
	UserClose struct {
		ID                int       `json:"id"`
		Reason            string    `json:"reason"`
		State             string    `json:"state"`
		PrimaryCampusID   int       `json:"primary_campus_id"`
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
		CommunityServices []struct {
			ID         int       `json:"id"`
			Duration   int       `json:"duration"`
			ScheduleAt time.Time `json:"schedule_at"`
			Occupation string    `json:"occupation"`
			State      string    `json:"state"`
			CreatedAt  time.Time `json:"created_at"`
			UpdatedAt  time.Time `json:"updated_at"`
		} `json:"community_services"`
		User struct {
			ID    int    `json:"id"`
			Login string `json:"login"`
			URL   string `json:"url"`
		} `json:"user"`
		Closer struct {
			ID    int    `json:"id"`
			Login string `json:"login"`
			URL   string `json:"url"`
		} `json:"closer"`
	}
	UserCloses []UserClose
)

const (
	CloseKindBlackHole         = "black_hole"
	CloseKindDeserter          = "deserter"
	CloseKindNonAdmitted       = "non_admitted"
	CloseKindOther             = "other"
	CloseKindSeriousMisconduct = "serious_misconduct"
)

func (userClose *UserClose) Create(ctx context.Context, bypassCache bool, kind string) (int, error) {
	endpoint := GetEndpoint(fmt.Sprintf("users/%d/closes", userClose.User.ID), nil)
	params := url.Values{}
	params.Set("close[closer_id]", strconv.Itoa(userClose.Closer.ID))
	params.Set("close[kind]", kind)
	params.Set("close[reason]", userClose.Reason)
	status, respData, err := RunRequest(GetClient(ctx, "public", "tig"), http.MethodPost, endpoint, params)
	if err == nil {
		if err = json.Unmarshal(respData, userClose); err == nil && !bypassCache {
			cached := *userClose
			intraCache.put(catCloses, cached.ID, &cached)
		}
	}
	return status, err
}

func (userClose *UserClose) Delete(ctx context.Context) error {
	endpoint := GetEndpoint("closes/"+strconv.Itoa(userClose.ID), nil)
	_, _, err := RunRequest(GetClient(ctx, "public", "tig"), http.MethodDelete, endpoint, nil)
	if err == nil {
		intraCache.delete(catCloses, userClose.ID)
	}
	return err
}

func (userClose *UserClose) Unclose(ctx context.Context, bypassCache bool) (int, error) {
	client := GetClient(ctx, "public", "tig")
	endpoint := GetEndpoint(fmt.Sprintf("closes/%d/unclose", userClose.ID), nil)
	status, _, err := RunRequest(client, http.MethodPatch, endpoint, nil)
	if err == nil {
		userClose.State = "unclose"
		if !bypassCache {
			cached := *userClose
			intraCache.put(catCloses, cached.ID, &cached)
		}
	}
	return status, err
}

func (userClose *UserClose) Get(ctx context.Context, bypassCache bool) error {
	if !bypassCache {
		if res, present := intraCache.get(catCloses, userClose.ID); present {
			*userClose = *res.(*UserClose)
			return nil
		}
	}
	userCloses := &UserCloses{}
	if err := userCloses.GetAll(ctx, bypassCache, getSingleParams(userClose.ID)); err != nil {
		return err
	}
	if len(*userCloses) == 0 {
		return fmt.Errorf("close %d does not exist", userClose.ID)
	}
	*userClose = (*userCloses)[0]
	return nil
}

func (userCloses *UserCloses) GetAll(ctx context.Context, bypassCache bool, params url.Values) error {
	if err := GetAll(GetClient(ctx, "public", "tig"), "closes", params, userCloses); err != nil {
		return err
	}
	if !bypassCache {
		for _, userClose := range *userCloses {
			cached := userClose
			intraCache.put(catCloses, cached.ID, &cached)
		}
	}
	return nil
}
