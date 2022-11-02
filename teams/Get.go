package teams

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *TeamsClient) Get(TeamId int) (*TeamGetResult, error) {
	req, err := c.client.NewRequest("GET", fmt.Sprintf("/v1/teams/%d", TeamId), nil)
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
