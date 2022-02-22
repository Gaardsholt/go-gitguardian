package incidents

import (
	"github.com/Gaardsholt/go-gitguardian/client"
)

type IncidentsClient struct {
	client *client.Client
}

func NewClient(opts ...client.ClientOption) (*IncidentsClient, error) {
	client, err := client.New(opts...)
	if err != nil {
		return nil, err
	}

	return &IncidentsClient{
		client: client,
	}, nil
}
