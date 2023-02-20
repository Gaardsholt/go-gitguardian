package incidents

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gaardsholt/go-gitguardian/types"
)

type IgnoreReason string

const (
	LowRisk         IgnoreReason = "low_risk"
	FalsePositive   IgnoreReason = "false_positive"
	test_credential IgnoreReason = "test_credential"
)

type IgnoreOptions struct {
	IgnoreReason IgnoreReason `json:"ignore_reason"`
}

type OccurrencesKind string

const (
	Realtime   OccurrencesKind = "Realtime"
	Historical OccurrencesKind = "Historical"
)

type Presence string

const (
	Present Presence = "Present"
	Removed Presence = "Removed"
)

type SourceHealth string

const (
	Safe    SourceHealth = "safe"
	AtRisk  SourceHealth = "at_risk"
	Unknown SourceHealth = "unknown"
)

type IncidentIgnoreResult struct {
	Result IncidentIgnoreResponse `json:"result"`
	Error  *Error                 `json:"error"`
}

type IncidentIgnoreResponse struct {
	ID              int        `json:"id"`
	Date            string     `json:"date"`
	Detector        Detector   `json:"detector"`
	SecretHash      string     `json:"secret_hash,omitempty"`
	GitguardianURL  string     `json:"gitguardian_url,omitempty"`
	Regression      bool       `json:"regression"`
	Status          string     `json:"status"`
	AssigneeEmail   string     `json:"assignee_email,omitempty"`
	OccurrenceCount int        `json:"occurrence_count,omitempty"`
	Occurrences     Occurrence `json:"occurrences,omitempty"`
	IgnoreReason    string     `json:"ignore_reason,omitempty"`
	IgnoredAt       string     `json:"ignored_at,omitempty"`
	SecretRevoked   bool       `json:"secret_revoked,omitempty"`
	Severity        string     `json:"severity,omitempty"`
	Validity        string     `json:"validity,omitempty"`
	ResolvedAt      string     `json:"resolved_at,omitempty"`
	ShareURL        string     `json:"share_url,omitempty"`
}

type IncidentIgnoreError struct {
	Detail string `json:"detail"`
}

func (c *IncidentsClient) Ignore(IncidentId int, lo IgnoreOptions) (*IncidentIgnoreResult, error) {
	ep := types.Endpoints["IgnoreSecretIncident"]

	req, err := c.client.NewRequest(ep.Operation, fmt.Sprintf(ep.Path, IncidentId), lo)
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
