package sources

import "time"

type Error struct {
	Detail string `json:"detail"`
}

type SourcesListResult struct {
	Result []SourcesResponse `json:"result"`
	Error  *Error            `json:"error"`
}

type SourcesGetResult struct {
	Result SourcesResponse `json:"result"`
	Error  *Error          `json:"error"`
}

type LastScan struct {
	Date   *time.Time `json:"date"`
	Status string     `json:"status"`
}

type SourcesResponse struct {
	ID                   int64    `json:"id"`
	URL                  string   `json:"url"`
	Type                 string   `json:"type"`
	FullName             string   `json:"full_name"`
	Visibility           string   `json:"visibility"`
	Health               string   `json:"health"`
	OpenIncidentsCount   int      `json:"open_incidents_count"`
	ClosedIncidentsCount int      `json:"closed_incidents_count"`
	LastScan             LastScan `json:"last_scan"`
}
