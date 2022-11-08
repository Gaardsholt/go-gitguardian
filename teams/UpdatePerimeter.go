package teams

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type UpdatePerimeterOptions struct {
	SourcesToAdd    []int64 `json:"sources_to_add"`
	SourcesToRemove []int64 `json:"sources_to_remove"`
}

func (c *TeamsClient) UpdatePerimeter(TeamId int, lo UpdatePerimeterOptions) error {

	// defaults to an empty array, the GitGuardian API wants an empty array, otherwise it will fail.
	if len(lo.SourcesToAdd) == 0 {
		lo.SourcesToAdd = []int64{}
	}
	if len(lo.SourcesToRemove) == 0 {
		lo.SourcesToRemove = []int64{}
	}

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(lo)
	if err != nil {
		return err
	}

	req, err := c.client.NewRequest("POST", fmt.Sprintf("/v1/teams/%d/sources", TeamId), b)
	if err != nil {
		return err
	}

	r, err := c.client.Client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusNoContent {
		var target Error
		decode := json.NewDecoder(r.Body)
		err = decode.Decode(&target)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s", target.Detail)
	}

	return nil
}
