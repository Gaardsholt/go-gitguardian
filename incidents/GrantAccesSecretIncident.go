package incidents

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type GrantAccessSecretIncidentOptions struct {
	Email    string `json:"email"`
	MemberID int64  `json:"member_id"`
}

type GrantAccessSecretIncidentResponse struct {
	ShareURL           string `json:"share_url"`
	IncidentID         int64  `json:"incident_id"`
	FeedbackCollection bool   `json:"feedback_collection"`
	AutoHealing        bool   `json:"auto_healing"`
	Token              string `json:"token"`
	ExpireAt           string `json:"expire_at"`
}

type GrantAccessSecretIncidentResult struct {
	Result GrantAccessSecretIncidentResponse `json:"result"`
	Error  *Error                            `json:"error"`
}

func (c *IncidentsClient) GrantAccessSecretIncident(IncidentId int, lo GrantAccessSecretIncidentOptions) (bool, error) {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(lo)
	if err != nil {
		return false, err
	}

	r, err := c.client.NewRequest("POST", fmt.Sprintf("/v1/incidents/secrets/%d/share", IncidentId), b)
	if err != nil {
		return false, err
	}

	return r.Response.StatusCode == http.StatusOK, nil

}
