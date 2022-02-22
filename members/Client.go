package members

import (
	"github.com/Gaardsholt/go-gitguardian/client"
)

type MembersClient struct {
	client *client.Client
}

func NewClient(opts ...client.ClientOption) (*MembersClient, error) {
	client, err := client.New(opts...)
	if err != nil {
		return nil, err
	}

	return &MembersClient{
		client: client,
	}, nil
}
