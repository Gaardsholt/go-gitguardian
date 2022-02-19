package sources

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *SourcesClient) Get(sourceId int) (*SourcesGetResult, error) {
	req, err := c.client.NewRequest("GET", fmt.Sprintf("/v1/sources/%d", sourceId), nil)
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
		return &SourcesGetResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	var target SourcesResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &SourcesGetResult{Result: target}, nil
}
