package members

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gaardsholt/go-gitguardian/types"
)

func (c *MembersClient) Delete(MemberId int) error {
	ep := types.Endpoints["MembersDelete"]

	req, err := c.client.NewRequest(ep.Operation, fmt.Sprintf(ep.Path, MemberId), nil)
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
