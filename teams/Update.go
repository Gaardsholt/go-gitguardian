package teams

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gaardsholt/go-gitguardian/types"
)

type Update struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c *TeamsClient) Update(TeamId int, lo Update) (*TeamGetResult, error) {
	ep := types.Endpoints["TeamsUpdate"]

	req, err := c.client.NewRequest(ep.Operation, fmt.Sprintf(ep.Path, TeamId), lo)
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
