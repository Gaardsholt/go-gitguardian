package incidents

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Gaardsholt/go-gitguardian/client"
	"github.com/Gaardsholt/go-gitguardian/types"
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

type IncidentsListTag string

const (
	FromHistoricalScan IncidentsListTag = "FROM_HISTORICAL_SCAN"
	IgnoredInCheckRun  IncidentsListTag = "IGNORED_IN_CHECK_RUN"
	Public             IncidentsListTag = "PUBLIC"
	Regression         IncidentsListTag = "REGRESSION"
	SensitiveFile      IncidentsListTag = "SENSITIVE_FILE"
	TestFile           IncidentsListTag = "TEST_FILE"
	None               IncidentsListTag = "NONE"
)

type ListOptions struct {
	Cursor        string                 `json:"-" url:"cursor"`         // The pagination cursor
	PerPage       *int                   `json:"-" url:"per_page"`       // Number of items to list per page.	[ 1 .. 100 ]
	DateBefore    *time.Time             `json:"-" url:"date_before"`    // Entries found before this date.
	DateAfter     *time.Time             `json:"-" url:"date_after"`     // Entries found after this date.
	Status        *IncidentsListStatus   `json:"-" url:"status"`         // Incidents with the following status.
	AssigneeEmail string                 `json:"-" url:"assignee_email"` // Incidents assigned to this email.
	Severity      *IncidentsListSeverity `json:"-" url:"severity"`       // Filter incidents by severity.
	Validity      *IncidentsListValidity `json:"-" url:"validity"`       // Secrets with the following validity.
	Tags          []IncidentsListTag     `json:"-" url:"tags"`           // Secrets with the following tags.
}

func (c *IncidentsClient) List(lo ListOptions) (*IncidentListResult, *client.PaginationMeta, error) {
	ep := types.Endpoints["ListSecretIncidents"]

	req, err := c.client.NewRequest(ep.Operation, ep.Path, lo)
	if err != nil {
		return nil, nil, err
	}

	// Validate query parameters
	if lo.PerPage != nil {
		if !(*lo.PerPage >= 1 && *lo.PerPage <= 100) {
			return nil, nil, fmt.Errorf("PerPage must be between 1 and 100")
		}
	}

	r, err := c.client.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		var target Error
		decode := json.NewDecoder(r.Body)
		err = decode.Decode(&target)
		if err != nil {
			return nil, nil, err
		}
		return &IncidentListResult{Error: &target}, nil, fmt.Errorf("%s", target.Detail)
	}

	var target []IncidentListResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, nil, err
	}
	pagination, err := client.GetPaginationMeta(r)
	if err != nil {
		return nil, nil, err
	}

	return &IncidentListResult{Result: target}, pagination, nil
}

func joinTags(tags []IncidentsListTag) string {
	var result string
	for _, tag := range tags {
		result = result + string(tag)
	}
	return result
}
