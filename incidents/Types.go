package incidents

import (
	"time"
)

type Error struct {
	Detail string `json:"detail"`
}

type IncidentListResult struct {
	Result []IncidentListResponse `json:"result"`
	Error  *Error                 `json:"error"`
}

type IncidentListResponse struct {
	ID               int64      `json:"id"`
	Date             time.Time  `json:"date"`
	Detector         Detector   `json:"detector"`
	SecretHash       string     `json:"secret_hash"`
	GitguardianURL   string     `json:"gitguardian_url"`
	Regression       bool       `json:"regression"`
	Status           string     `json:"status"`
	AssigneeEmail    string     `json:"assignee_email"`
	OccurrencesCount int64      `json:"occurrences_count"`
	IgnoreReason     string     `json:"ignore_reason"`
	IgnoredAt        *time.Time `json:"ignored_at"`
	SecretRevoked    bool       `json:"secret_revoked"`
	Severity         string     `json:"severity"`
	Validity         string     `json:"validity"`
	ResolvedAt       *time.Time `json:"resolved_at"`
	ShareURL         string     `json:"share_url"`
	Tags             []string   `json:"tags"`
}

type Detector struct {
	Name                     string `json:"name"`
	DisplayName              string `json:"display_name"`
	Nature                   string `json:"nature"`
	Family                   string `json:"family"`
	DetectorGroupName        string `json:"detector_group_name"`
	DetectorGroupDisplayName string `json:"detector_group_display_name"`
}

type IncidentGetResult struct {
	Result IncidentGetResponse `json:"result"`
	Error  *Error              `json:"error"`
}
type IncidentGetResponse struct {
	ID               int64        `json:"id"`
	Date             time.Time    `json:"date"`
	Detector         Detector     `json:"detector"`
	SecretHash       string       `json:"secret_hash"`
	GitguardianURL   string       `json:"gitguardian_url"`
	Regression       bool         `json:"regression"`
	Status           string       `json:"status"`
	AssigneeEmail    string       `json:"assignee_email"`
	OccurrencesCount int64        `json:"occurrences_count"`
	Occurrences      []Occurrence `json:"occurrences"`
	IgnoreReason     string       `json:"ignore_reason"`
	Severity         string       `json:"severity"`
	Validity         string       `json:"validity"`
	IgnoredAt        *time.Time   `json:"ignored_at"`
	SecretRevoked    bool         `json:"secret_revoked"`
	ResolvedAt       *time.Time   `json:"resolved_at"`
	ShareURL         string       `json:"share_url"`
	Tags             []string     `json:"tags"`
}

type Occurrence struct {
	ID         int64   `json:"id"`
	IncidentID int64   `json:"incident_id"`
	Kind       string  `json:"kind"`
	SHA        string  `json:"sha"`
	Source     Source  `json:"source"`
	AuthorName string  `json:"author_name"`
	AuthorInfo string  `json:"author_info"`
	Date       string  `json:"date"`
	Presence   string  `json:"presence"`
	URL        string  `json:"url"`
	Matches    []Match `json:"matches"`
	Filepath   string  `json:"filepath"`
}

type Match struct {
	Name          string `json:"name"`
	IndiceStart   int64  `json:"indice_start"`
	IndiceEnd     int64  `json:"indice_end"`
	PostLineStart int64  `json:"post_line_start"`
	PostLineEnd   int64  `json:"post_line_end"`
}

type Source struct {
	ID                   int64    `json:"id"`
	URL                  string   `json:"url"`
	Type                 string   `json:"type"`
	FullName             string   `json:"full_name"`
	Health               string   `json:"health"`
	OpenIncidentsCount   int64    `json:"open_incidents_count"`
	ClosedIncidentsCount int64    `json:"closed_incidents_count"`
	Visibility           string   `json:"visibility"`
	LastScan             LastScan `json:"last_scan"`
}

type LastScan struct {
	Date   string `json:"date"`
	Status string `json:"status"`
}
