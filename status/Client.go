package status

import (
	"github.com/Gaardsholt/go-gitguardian/client"
)

type StatusClient struct {
	client *client.Client
}

func NewClient(opts ...client.ClientOption) (*StatusClient, error) {
	client, err := client.New(opts...)
	if err != nil {
		return nil, err
	}

	return &StatusClient{
		client: client,
	}, nil
}
