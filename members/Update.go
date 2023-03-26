package members

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gaardsholt/go-gitguardian/types"
)

type MemberUpdateOptions struct {
	Role string `json:"role" url:"-"`
}

func (c *MembersClient) Update(MemberId int, lo MemberUpdateOptions) (*MemberGetResult, error) {
	ep := types.Endpoints["MembersUpdate"]

	req, err := c.client.NewRequest(ep.Operation, fmt.Sprintf(ep.Path, MemberId), lo)
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
