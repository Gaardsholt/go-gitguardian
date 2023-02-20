package incidents

import (
	"fmt"
	"net/http"

	"github.com/Gaardsholt/go-gitguardian/types"
)

type UnshareSecretIncidentResponse struct {
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

type UnshareSecretIncidentResult struct {
	Result UnshareSecretIncidentResponse `json:"result"`
	Error  *Error                        `json:"error"`
}

func (c *IncidentsClient) UnshareSecretIncident(IncidentId int) (bool, error) {
	ep := types.Endpoints["UnshareSecretIncident"]

	r, err := c.client.NewRequest(ep.Operation, fmt.Sprintf(ep.Path, IncidentId), nil)
	if err != nil {
		return false, err
	}

	return r.Response.StatusCode == http.StatusOK, nil

}
