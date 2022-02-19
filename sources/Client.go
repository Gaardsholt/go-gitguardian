package sources

import (
	"github.com/Gaardsholt/go-gitguardian/client"
)

type SourcesClient struct {
	client *client.Client
}

func NewClient(opts ...client.ClientOption) (*SourcesClient, error) {
	client, err := client.New(opts...)
	if err != nil {
		return nil, err
	}

	return &SourcesClient{
		client: client,
	}, nil
}
