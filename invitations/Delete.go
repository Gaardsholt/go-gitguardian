package invitations

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gaardsholt/go-gitguardian/types"
)

type InvitationsDeleteResult struct {
	Error *Error `json:"error"`
}

func (c *InvitationsClient) Delete(InvitationId int) (*InvitationsDeleteResult, error) {
	ep := types.Endpoints["InvitationsDelete"]

	req, err := c.client.NewRequest(ep.Operation, fmt.Sprintf(ep.Path, InvitationId), nil)
	if err != nil {
		return nil, err
	}

	r, err := c.client.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusNoContent {
		var target Error
		decode := json.NewDecoder(r.Body)
		err = decode.Decode(&target)
		if err != nil {
			return nil, err
		}
		return &InvitationsDeleteResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	return nil, nil
}
