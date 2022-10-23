package members

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *MembersClient) Get(MemberId int) (*MemberGetResult, error) {
	req, err := c.client.NewRequest("GET", fmt.Sprintf("/v1/members/%d", MemberId), nil)
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
		return &MemberGetResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	var target MembersResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &MemberGetResult{Result: target}, nil
}
