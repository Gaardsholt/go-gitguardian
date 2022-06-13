package incidents

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type IgnoreReason string

const (
	LowRisk IgnoreReason = "low_risk"
	FalsePositive IgnoreReason = "false_positive"
	test_credential IgnoreReason = "test_credential"
)

type IgnoreOptions struct {
	IgnoreReason IgnoreReason `json:"ignore_reason"`
}

type OccurrencesKind string
const (
	Realtime OccurrencesKind = "Realtime"
	Historical OccurrencesKind = "Historical"
)

type OccurrencesPresence string
const (
	Present OccurrencesPresence = "Present"
	Removed OccurrencesPresence = "Removed"
)

type SourceHealth string
const (
	Safe SourceHealth = "safe"
	AtRisk SourceHealth = "at_risk"
	Unknown SourceHealth = "unknown"
)

type IncidentIgnoreResult struct {
	Result IncidentIgnoreResponse `json:"result"`
	Error  *Error                 `json:"error"`
}

type IncidentIgnoreResponse struct {
	ID       int       `json:"id"`
	Date     string `json:"date"`
	Detector struct {
		Name                     string `json:"name"`
		DisplayName              string `json:"display_name"`
		Nature                   string `json:"nature"`
		Family                   string `json:"family,omitempty"`
		DetectorGroupName        string `json:"detector_group_name,omitempty"`
		DetectorGroupDisplayName string `json:"detector_group_display_name,omitempty"`
	} `json:"detector"`
	SecretHash      string      `json:"secret_hash,omitempty"`
	GitguardianURL  string      `json:"gitguardian_url,omitempty"`
	Regression      bool        `json:"regression"`
	Status          string      `json:"status"`
	AssigneeEmail   string      `json:"assignee_email,omitempty"`
	OccurrenceCount int         `json:"occurrence_count,omitempty"`
	Occurrences     struct {
		ID int `json:"id,omitempty"`
		IncidentID int `json:"incident_id,omitempty"`
		Kind *OccurrencesKind `json:"kind"`
		Sha string `json:"sha"`
		Source struct {
			ID int `json:"id"`
			Url string `json:"url"`
			Type string `json:"type"`
			Fullname string `json:"full_name,omitempty"`
			Health *SourceHealth `json:"health,omitempty"`
			OpenIncidentsCount int `json:"open_incidents_count,omitempty"`
			ClosedIncidentsCount int `json:"closed_incidents_count,omitempty"`
			Visibility string `json:"visibility,omitempty"`
			LastScan string `json:"last_scan,omitempty"`
		} `json:"source"`
		AuthorName string `json:"author_name"`
		AuthorInfo string `json:"author_info"`
		Date string `json:"date"`
		Presence *OccurrencesPresence `json:"presence,omitempty"`
		Url string `json:"url"`
		Matches []struct {
			Name string `json:"name"`
			IndiceStart int `json:"indice_start"`
			IndiceEnd int `json:"indice_end"`
			PreLineStart int `json:"pre_line_start"`
			PreLineEnd int `json:"pre_line_end"`
			PostLineStart int `json:"post_line_start"`
			PostLineEnd int `json:"post_line_end"` 
		}
		Filepath string `json:"filepath"`
	} `json:"occurrences,omitempty"`
	IgnoreReason    string      `json:"ignore_reason,omitempty"`
	IgnoredAt       string   `json:"ignored_at,omitempty"`
	SecretRevoked   bool        `json:"secret_revoked,omitempty"`
	Severity        string      `json:"severity,omitempty"`
	Validity        string      `json:"validity,omitempty"`
	ResolvedAt      string `json:"resolved_at,omitempty"`
	ShareURL        string      `json:"share_url,omitempty"`
}

type IncidentIgnoreError struct {
	Detail string `json:"detail"`
}

func (c *IncidentsClient) Ignore(IncidentId int, lo IgnoreOptions) (*IncidentIgnoreResult, error) {
	b, err := json.Marshal(lo)
	if err != nil {
		return nil, err
	}

	bbuffer := bytes.NewBuffer(b)

	req, err := c.client.NewRequest("POST", fmt.Sprintf("/v1/incidents/secrets/%d/ignore", IncidentId), bbuffer)
	if err != nil {
		return nil, err
	}

	r, err := c.client.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		var target Error
		decode := json.NewDecoder(r.Body)
		err = decode.Decode(&target)
		if err != nil {
			return nil, err
		}
		return &IncidentIgnoreResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	var target IncidentIgnoreResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &IncidentIgnoreResult{Result: target}, nil
}