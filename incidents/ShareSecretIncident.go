package incidents

import (
	"encoding/json"
	"fmt"

	"github.com/Gaardsholt/go-gitguardian/types"
)

type ShareSecretIncidentOptions struct {
	AutoHealing        bool  `json:"auto_healing"`
	FeedbackCollection bool  `json:"feedback_collection"`
	Lifespan           int64 `json:"lifespan"`
}

type ShareSecretIncidentResponse struct {
	ShareURL           string `json:"share_url"`
	IncidentID         int64  `json:"incident_id"`
	FeedbackCollection bool   `json:"feedback_collection"`
	AutoHealing        bool   `json:"auto_healing"`
	Token              string `json:"token"`
	ExpireAt           string `json:"expire_at"`
}

type ShareSecretIncidentResult struct {
	Result ShareSecretIncidentResponse `json:"result"`
	Error  *Error                      `json:"error"`
}

func (c *IncidentsClient) ShareSecretIncident(IncidentId int, lo ShareSecretIncidentOptions) (*ShareSecretIncidentResult, error) {
	ep := types.Endpoints["ShareSecretIncident"]

	r, err := c.client.NewRequest(ep.Operation, fmt.Sprintf(ep.Path, IncidentId), lo)
	if err != nil {
		return nil, err
	}

	var target ShareSecretIncidentResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &ShareSecretIncidentResult{Result: target}, nil

}
