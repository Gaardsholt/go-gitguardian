package incidents

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type IncidentsListStatus string

const (
	Ignored   IncidentsListStatus = "IGNORED"
	Triggered IncidentsListStatus = "TRIGGERED"
	Assigned  IncidentsListStatus = "ASSIGNED"
	Viewer    IncidentsListStatus = "RESOLVED"
)

type IncidentsListSeverity string

const (
	Critical        IncidentsListSeverity = "critical"
	High            IncidentsListSeverity = "high"
	Medium          IncidentsListSeverity = "medium"
	Low             IncidentsListSeverity = "low"
	Info            IncidentsListSeverity = "info"
	UnknownSeverity IncidentsListSeverity = "unknown"
)

type IncidentsListValidity string

const (
	Valid           IncidentsListValidity = "valid"
	Invalid         IncidentsListValidity = "invalid"
	UnknownValidity IncidentsListValidity = "unknown"
	Cannot_Check    IncidentsListValidity = "cannot_check"
)

type ListOptions struct {
	Page          *int                   `json:"page"`           // Page number.
	PerPage       *int                   `json:"per_page"`       // Number of items to list per page.	[ 1 .. 100 ]
	DateBefore    *time.Time             `json:"date_before"`    // Entries found before this date.
	DateAfter     *time.Time             `json:"date_after"`     // Entries found after this date.
	Status        *IncidentsListStatus   `json:"status"`         // Incidents with the following status.
	AssigneeEmail string                 `json:"assignee_email"` // Incidents assigned to this email.
	Severity      *IncidentsListSeverity `json:"severity"`       // Filter incidents by severity.
	Validity      *IncidentsListValidity `json:"validity"`       // Secrets with the following validity.
}

func (c *IncidentsClient) List(lo ListOptions) (*IncidentListResult, error) {
	req, err := c.client.NewRequest("GET", "/v1/incidents/secrets", nil)
	if err != nil {
		return nil, err
	}

	// Add query parameters
	q := req.URL.Query()

	if lo.PerPage != nil {
		if !(*lo.PerPage >= 1 && *lo.PerPage <= 100) {
			return nil, fmt.Errorf("PerPage must be between 1 and 100")
		}
		q.Add("per_page", strconv.Itoa(*lo.PerPage))
	}

	if lo.Page != nil {
		if !(*lo.Page >= 1) {
			return nil, fmt.Errorf("Page must be geater than zero")
		}
		q.Add("page", strconv.Itoa(*lo.Page))
	}

	if lo.DateBefore != nil {
		q.Add("date_before", lo.DateBefore.Format(time.RFC3339Nano))
	}

	if lo.DateAfter != nil {
		q.Add("date_after", lo.DateAfter.Format(time.RFC3339Nano))
	}

	if lo.Status != nil {
		q.Add("status", string(*lo.Status))
	}

	if lo.AssigneeEmail != "" {
		q.Add("assignee_email", string(lo.AssigneeEmail))
	}

	if lo.Severity != nil {
		q.Add("severity", string(*lo.Severity))
	}

	if lo.Validity != nil {
		q.Add("validity", string(*lo.Validity))
	}

	req.URL.RawQuery = q.Encode()

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
		return &IncidentListResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	var target []IncidentListResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &IncidentListResult{Result: target}, nil
}
