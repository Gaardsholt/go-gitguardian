package incidents

import (
	"encoding/json"
	"fmt"
)

type ReopenSecretIncidentResponse struct {
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

type ReopenSecretIncidentResult struct {
	Result ReopenSecretIncidentResponse `json:"result"`
	Error  *Error                       `json:"error"`
}

func (c *IncidentsClient) ReopenSecretIncident(IncidentId int) (*ReopenSecretIncidentResult, error) {
	r, err := c.client.NewRequest("POST", fmt.Sprintf("/v1/incidents/secrets/%d/reopen", IncidentId), nil)
	if err != nil {
		return nil, err
	}

	var target ReopenSecretIncidentResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &ReopenSecretIncidentResult{Result: target}, nil

}
