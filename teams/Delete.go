package teams

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gaardsholt/go-gitguardian/types"
)

type TeamsDeleteResult struct {
	Error *Error `json:"error"`
}

func (c *TeamsClient) Delete(TeamId int) (*TeamsDeleteResult, error) {
	ep := types.Endpoints["TeamsDelete"]

	req, err := c.client.NewRequest(ep.Operation, fmt.Sprintf(ep.Path, TeamId), nil)
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
		return &TeamsDeleteResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	return nil, nil
}
