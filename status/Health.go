package status

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gaardsholt/go-gitguardian/types"
)

func (c *StatusClient) Health() (*HealthResult, error) {
	ep := types.Endpoints["Health"]

	req, err := c.client.NewRequest(ep.Operation, ep.Path, nil)
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
		return &HealthResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	var target HealthResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &HealthResult{Result: target}, nil
}
