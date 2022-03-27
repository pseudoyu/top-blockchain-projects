package models

import "time"

// HeadCommit describes a head commit of default branch
type HeadCommit struct {
	Sha    string `json:"sha"`
	Commit struct {
		Committer struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"committer"`
	} `json:"commit"`
}
