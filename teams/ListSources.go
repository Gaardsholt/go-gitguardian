package teams

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gaardsholt/go-gitguardian/client"
	"github.com/Gaardsholt/go-gitguardian/types"
)

type LastScanStatus string

const (
	Pending  LastScanStatus = "pending"
	Running  LastScanStatus = "running"
	Canceled LastScanStatus = "canceled"
	Failed   LastScanStatus = "failed"
	TooLarge LastScanStatus = "too_large"
	Timeout  LastScanStatus = "timeout"
	Finished LastScanStatus = "finished"
)

type Health string

const (
	Safe    Health = "safe"
	Unknown Health = "unknown"
	AtRisk  Health = "at_risk"
)

type Type string

const (
	Bitbucket   Type = "bitbucket"
	Github      Type = "github"
	Gitlab      Type = "gitlab"
	AzureDevops Type = "azure_devops"
)

type Ordering string

const (
	ASC  Ordering = "last_scan_date"
	DESC Ordering = "-last_scan_date"
)

type Visibility string

const (
	Public   Visibility = "public"
	Private  Visibility = "private"
	Internal Visibility = "internal"
)

type AuditLogsListOptions struct {
	Cursor         string         `json:"-" url:"cursor"`           // Pagination cursor.
	PerPage        *int           `json:"-" url:"per_page"`         // Number of items to list per page.	[ 1 .. 100 ]
	Search         string         `json:"-" url:"search"`           // Sources matching this search.
	LastScanStatus LastScanStatus `json:"-" url:"last_scan_status"` // Filter sources based on the status of their latest historical scan.
	Health         Health         `json:"-" url:"health"`           // Filter sources based on their health status.
	Type           Type           `json:"-" url:"type"`             // Filter by integration type.
	Ordering       Ordering       `json:"-" url:"ordering"`         // Sort the results by their field value. The default sort is ASC, DESC if the field is preceded by a '-'.
	Visibility     Visibility     `json:"-" url:"visibility"`       // Filter by visibility status.
	ExternalId     string         `json:"-" url:"external_id"`      // Filter by specific external id.
}

func (c *TeamsClient) ListSources(TeamId int) (*TeamGetResult, *client.PaginationMeta, error) {
	ep := types.Endpoints["TeamsListSources"]

	req, err := c.client.NewRequest(ep.Operation, fmt.Sprintf(ep.Path, TeamId), nil)
	if err != nil {
		return nil, nil, err
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
		return &TeamGetResult{Error: &target}, nil, fmt.Errorf("%s", target.Detail)
	}

	var target TeamsResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, nil, err
	}

	pagination, err := client.GetPaginationMeta(r)
	if err != nil {
		return nil, nil, err
	}

	return &TeamGetResult{Result: target}, pagination, nil
}
