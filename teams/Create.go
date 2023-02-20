package teams

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TeamsCreateResult struct {
	Result TeamsResponse `json:"result"`
	Error  *Error        `json:"error"`
}

type TeamsCreateOptions struct {
	Name string `json:"name"`
}

func (c *TeamsClient) Create(lo TeamsCreateOptions) (*TeamsCreateResult, error) {
	req, err := c.client.NewRequest("POST", "/v1/teams", lo)
	if err != nil {
		return nil, err
	}

	r, err := c.client.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusCreated {
		var target Error
		decode := json.NewDecoder(r.Body)
		err = decode.Decode(&target)
		if err != nil {
			return nil, err
		}
		return &TeamsCreateResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	var target TeamsResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &TeamsCreateResult{Result: target}, nil
}
