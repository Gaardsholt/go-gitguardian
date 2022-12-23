package teams

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TeamsAddMemberOptions struct {
	MemberId           int64              `json:"member_id"`
	TeamPermission     TeamPermission     `json:"team_permission"`
	IncidentPermission IncidentPermission `json:"incident_permission"`
}

func (c *TeamsClient) AddMember(TeamId int, lo TeamsAddMemberOptions) (*ListMembershipsResult, error) {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(lo)
	if err != nil {
		return nil, err
	}

	req, err := c.client.NewRequest("POST", fmt.Sprintf("/v1/teams/%d/team_memberships", TeamId), b)
	if err != nil {
		return nil, err
	}

	r, err := c.client.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		var target Error
		decode := json.NewDecoder(r.Body)

		err = decode.Decode(&target)
		if err != nil {
			return nil, err
		}

		if target.Detail == "" {
			target.Detail = string(bodyBytes)
		}

		return &ListMembershipsResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	var target ListMembershipsResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &ListMembershipsResult{Result: target}, nil
}
