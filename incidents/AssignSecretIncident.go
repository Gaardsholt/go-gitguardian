package incidents

import (
	"encoding/json"
	"fmt"

	"github.com/Gaardsholt/go-gitguardian/types"
)

type AssignSecretIncidentOptions struct {
	Email    string `json:"email"`
	MemberID int64  `json:"member_id"`
}

type AssignSecretIncidentResponse struct {
	ID              int64    `json:"id"`
	Date            string   `json:"date"`
	Detector        Detector `json:"detector"`
	SecretHash      string   `json:"secret_hash"`
	GitguardianURL  string   `json:"gitguardian_url"`
	Regression      bool     `json:"regression"`
	Status          string   `json:"status"`
	AssigneeEmail   string   `json:"assignee_email"`
	OccurrenceCount int64    `json:"occurrence_count"`
	IgnoreReason    string   `json:"ignore_reason"`
	IgnoredAt       string   `json:"ignored_at"`
	SecretRevoked   bool     `json:"secret_revoked"`
	Severity        string   `json:"severity"`
	Validity        string   `json:"validity"`
	ShareURL        string   `json:"share_url"`
}

type AssignSecretIncidentResult struct {
	Result AssignSecretIncidentResponse `json:"result"`
	Error  *Error                       `json:"error"`
}

func (c *IncidentsClient) AssignSecretIncident(IncidentId int, lo AssignSecretIncidentOptions) (*AssignSecretIncidentResult, error) {
	ep := types.Endpoints["AssignSecretIncident"]

	r, err := c.client.NewRequest(ep.Operation, fmt.Sprintf(ep.Path, IncidentId), lo)
	if err != nil {
		return nil, err
	}

	var target AssignSecretIncidentResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &AssignSecretIncidentResult{Result: target}, nil

}
