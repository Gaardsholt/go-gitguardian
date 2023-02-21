package incidents

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Gaardsholt/go-gitguardian/types"
)

type ListSecretTeamsOptions struct {
	Cursor             string              `json:"cursor"`              // Pagination cursor.
	TeamId             *int                `json:"team_id"`             // Filter on a specific invitation id.
	IncidentPermission *IncidentPermission `json:"incident_permission"` // Filter accesses with a specific permission.
}

type ListSecretTeamsResult struct {
	Result []ListSecretTeamsResponse `json:"result"`
	Error  *Error                    `json:"error"`
}

type ListSecretTeamsResponse struct {
	TeamId             int64  `json:"team_id"`
	IncidentId         int64  `json:"incident_id"`
	IncidentPermission string `json:"incident_permission"`
}

func (c *IncidentsClient) ListSecretTeams(IncidentId int, lo ListSecretTeamsOptions) (*ListSecretTeamsResult, error) {
	ep := types.Endpoints["ListSecretTeams"]

	req, err := c.client.NewRequest(ep.Operation, fmt.Sprintf(ep.Path, IncidentId), nil)
	if err != nil {
		return nil, err
	}

	// Add query parameters
	q := req.URL.Query()

	if lo.TeamId != nil {
		q.Add("team_id", strconv.Itoa(*lo.TeamId))
	}
	if lo.IncidentPermission != nil {
		q.Add("incident_permission", string(*lo.IncidentPermission))
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
		return &ListSecretTeamsResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	var target []ListSecretTeamsResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &ListSecretTeamsResult{Result: target}, nil
}
