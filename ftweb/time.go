package ftweb

import (
	"fmt"
	"strings"
	"time"
)

type Time struct {
	time.Time
}

const WebTimeFormat = "2006-01-02 15:04:05 MST"

func NewWebTime(t time.Time) *Time {
	return &Time{Time: t}
}

func (wt *Time) MarshalJSON() ([]byte, error) {
	if wt.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("%q", wt.Format(WebTimeFormat))), nil
}

func (wt *Time) UnmarshalJSON(data []byte) error {
	raw := strings.Trim(string(data), "\"")
	if raw == "null" {
		wt.Time = time.Time{}
		return nil
	}
	// Default to local time, as Intra webhooks will try to return locally formatted time
	date, err := time.ParseInLocation(WebTimeFormat, raw, time.Local)
	if err == nil {
		wt.Time = date
	}
	return err
}
