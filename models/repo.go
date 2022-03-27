package models

import "time"

// Repo describes a Github repository with additional field, last commit date
type Repo struct {
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	DefaultBranch  string    `json:"default_branch"`
	Stars          int       `json:"stargazers_count"`
	Forks          int       `json:"forks_count"`
	Issues         int       `json:"open_issues_count"`
	Created        time.Time `json:"created_at"`
	Updated        time.Time `json:"updated_at"`
	URL            string    `json:"html_url"`
	LastCommitDate time.Time `json:"-"`
}
