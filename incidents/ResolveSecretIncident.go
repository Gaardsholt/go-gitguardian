package incidents

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ResolveSecretIncidentOptions struct {
	SecretRevoked bool `json:"secret_revoked"`
}

type ResolveSecretIncidentResponse struct {
	ID              int64        `json:"id"`
	Date            string       `json:"date"`
	Detector        Detector     `json:"detector"`
	SecretHash      string       `json:"secret_hash"`
	GitguardianURL  string       `json:"gitguardian_url"`
	Regression      bool         `json:"regression"`
	Status          string       `json:"status"`
	AssigneeEmail   string       `json:"assignee_email"`
	OccurrenceCount int64        `json:"occurrence_count"`
	Occurrences     []Occurrence `json:"occurrences"`
	IgnoreReason    string       `json:"ignore_reason"`
	IgnoredAt       string       `json:"ignored_at"`
	SecretRevoked   bool         `json:"secret_revoked"`
	Severity        string       `json:"severity"`
	Validity        string       `json:"validity"`
	ShareURL        string       `json:"share_url"`
}

type ResolveSecretIncidentResult struct {
	Result ResolveSecretIncidentResponse `json:"result"`
	Error  *Error                        `json:"error"`
}

func (c *IncidentsClient) ResolveSecretIncident(IncidentId int, lo ResolveSecretIncidentOptions) (*ResolveSecretIncidentResult, error) {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(lo)
	if err != nil {
		return nil, err
	}

	r, err := c.client.NewRequest("POST", fmt.Sprintf("/v1/incidents/secrets/%d/resolve", IncidentId), b)
	if err != nil {
		return nil, err
	}

	var target ResolveSecretIncidentResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &ResolveSecretIncidentResult{Result: target}, nil

}
