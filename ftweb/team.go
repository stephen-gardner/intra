package ftweb

import (
	"intra/ftobj"
)

type Team struct {
	ID              int            `json:"id,omitempty"`
	Name            string         `json:"name,omitempty"`
	Project         *ftobj.Project `json:"project,omitempty"`
	Leader          *ftobj.User    `json:"leader,omitempty"`
	Users           []ftobj.User   `json:"users,omitempty"`
	RepoURL         string         `json:"repo_url,omitempty"`
	RepoUUID        string         `json:"repo_uuid,omitempty"`
	FinalMark       int            `json:"final_mark,omitempty"`
	PrimaryCampusID int            `json:"primary_campus_id,omitempty"`
	DeadlineAt      *Time          `json:"deadline_at,omitempty"`
	TerminatingAt   *Time          `json:"terminating_at,omitempty"`
}
