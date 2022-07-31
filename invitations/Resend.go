package invitations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type InvitationsResendResult struct {
	Error *Error `json:"error"`
}

func (c *InvitationsClient) Resend(InvitationId int) (*InvitationsResendResult, error) {
	req, err := c.client.NewRequest("POST", fmt.Sprintf("/v1/invitations/%d/resend", InvitationId), nil)
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
		return &InvitationsResendResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	return nil, nil
}
