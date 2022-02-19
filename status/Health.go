package status

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *StatusClient) Health() (*HealthResult, error) {
	req, err := c.client.NewRequest("GET", "/v1/health", nil)
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
