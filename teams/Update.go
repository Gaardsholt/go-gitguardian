package teams

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Update struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c *TeamsClient) Update(TeamId int, lo Update) (*TeamGetResult, error) {
	payload := new(bytes.Buffer)
	err := json.NewEncoder(payload).Encode(lo)
	if err != nil {
		return nil, err
	}

	req, err := c.client.NewRequest("PATCH", fmt.Sprintf("/v1/teams/%d", TeamId), payload)
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
		return &TeamGetResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	var target TeamsResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &TeamGetResult{Result: target}, nil
}
