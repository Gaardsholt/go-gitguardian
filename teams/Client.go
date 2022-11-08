package teams

import (
	"github.com/Gaardsholt/go-gitguardian/client"
)

type TeamsClient struct {
	client *client.Client
}

func NewClient(opts ...client.ClientOption) (*TeamsClient, error) {
	client, err := client.New(opts...)
	if err != nil {
		return nil, err
	}

	return &TeamsClient{
		client: client,
	}, nil
}
