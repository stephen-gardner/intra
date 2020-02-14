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
	CloseSelf                  = -1
)

func (userClose *UserClose) CloseUser(
	ctx context.Context,
	bypassCache bool,
	userID, closerID int,
	kind, reason string,
) (int, []byte, error) {
	endpoint := GetEndpoint(fmt.Sprintf("users/%d/closes", userID), nil)
	params := url.Values{}
	if closerID == CloseSelf {
		params.Set("close[closer_id]", strconv.Itoa(userID))
	} else {
		params.Set("close[closer_id]", strconv.Itoa(closerID))
	}
	params.Set("close[kind]", kind)
	params.Set("close[reason]", reason)
	status, respData, err := RunRequest(GetClient(ctx, "public", "tig"), http.MethodPost, endpoint, params)
	if err == nil {
		if err = json.Unmarshal(respData, userClose); err == nil && !bypassCache {
			endpoint := GetEndpoint("closes/"+strconv.Itoa(userClose.ID), nil)
			intraCache.put(endpoint, *userClose)
		}
	}
	return status, respData, err
}

func (userClose *UserClose) Unclose(ctx context.Context, bypassCache bool) (int, []byte, error) {
	client := GetClient(ctx, "public", "tig")
	endpoint := GetEndpoint(fmt.Sprintf("closes/%d/unclose", userClose.ID), nil)
	status, respData, err := RunRequest(client, http.MethodPatch, endpoint, nil)
	if err == nil {
		userClose.State = "unclose"
		if !bypassCache {
			endpoint := GetEndpoint("closes/"+strconv.Itoa(userClose.ID), nil)
			intraCache.put(endpoint, *userClose)
		}
	}
	return status, respData, err
}

func (userClose *UserClose) GetClose(ctx context.Context, bypassCache bool, ID int) error {
	IDStr := strconv.Itoa(ID)
	endpoint := GetEndpoint("closes/"+IDStr, nil)
	if !bypassCache {
		if res, present := intraCache.get(endpoint); present {
			*userClose = res.(UserClose)
			return nil
		}
	}
	userCloses := &UserCloses{}
	if err := userCloses.GetAllCloses(ctx, bypassCache, getSingleParams(IDStr)); err != nil {
		return err
	}
	if len(*userCloses) == 0 {
		return fmt.Errorf("close %d does not exist", ID)
	}
	*userClose = (*userCloses)[0]
	return nil
}

func (userCloses *UserCloses) GetAllCloses(ctx context.Context, bypassCache bool, params url.Values) error {
	if err := GetAll(GetClient(ctx, "public", "tig"), "closes", params, userCloses); err != nil {
		return err
	}
	if !bypassCache {
		for _, userClose := range *userCloses {
			endpoint := GetEndpoint("closes/"+strconv.Itoa(userClose.ID), nil)
			intraCache.put(endpoint, userClose)
		}
	}
	return nil
}
