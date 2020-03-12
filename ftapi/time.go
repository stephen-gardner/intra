package ftapi

import (
	"fmt"
	"strings"
	"time"
)

type Time struct {
	time.Time
}

const TimeFormat = "2006-01-02T15:04:05.000Z"

func NewTime(t time.Time) *Time {
	return &Time{Time: t}
}

func (it *Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", it.UTC().Format(TimeFormat))), nil
}

func (it *Time) UnmarshalJSON(data []byte) error {
	raw := strings.Trim(string(data), "\"")
	if raw == "null" {
		it.Time = time.Time{}
		return nil
	}
	date, err := time.Parse(TimeFormat, raw)
	if err == nil {
		it.Time = date
	}
	return err
}
