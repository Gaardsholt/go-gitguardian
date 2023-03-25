package incidents

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gaardsholt/go-gitguardian/types"
)

type ListSecretTeamsOptions struct {
	Cursor             string              `json:"-" url:"cursor"`              // Pagination cursor.
	TeamId             *int                `json:"-" url:"team_id"`             // Filter on a specific invitation id.
	IncidentPermission *IncidentPermission `json:"-" url:"incident_permission"` // Filter accesses with a specific permission.
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
