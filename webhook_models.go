package intra

import (
	"strings"
	"time"
)

type (
	WebTime struct {
		time.Time
	}
	WebTeam struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Project struct {
			ID   int    `json:"id"`
			Slug string `json:"slug"`
		} `json:"project"`
		Leader struct {
			ID    int    `json:"id"`
			Login string `json:"login"`
			URL   string `json:"url"`
		} `json:"leader"`
		Users []struct {
			UsualFullName string                          `json:"usual_full_name"`
			Location      string                          `json:"location"`
			Cursus        []map[string]map[string]float64 `json:"cursus"`
			Login         string                          `json:"login"`
			Email         string                          `json:"email"`
			FullName      string                          `json:"full_name"`
			ImageURL      string                          `json:"image_url"`
		} `json:"users"`
		RepoURL         string  `json:"repo_url"`
		RepoUUID        string  `json:"repo_uuid"`
		FinalMark       int     `json:"final_mark"`
		PrimaryCampusID int     `json:"primary_campus_id"`
		DeadlineAt      WebTime `json:"deadline_at"`
		TerminatingAt   WebTime `json:"terminating_at"`
	}
)

const webTimeFormat = "2006-01-02 15:04:05 MST"

func (wt *WebTime) UnmarshalJSON(data []byte) error {
	raw := strings.Trim(string(data), "\"")
	if raw == "null" {
		wt.Time = time.Time{}
		return nil
	}
	date, err := time.ParseInLocation(webTimeFormat, raw, time.Local)
	if err == nil {
		wt.Time = date
	}
	return err
}
