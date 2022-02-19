package scan

import (
	"github.com/Gaardsholt/go-gitguardian/client"
)

type ScanClient struct {
	client *client.Client
}

func NewClient(opts ...client.ClientOption) (*ScanClient, error) {
	client, err := client.New(opts...)
	if err != nil {
		return nil, err
	}

	return &ScanClient{
		client: client,
	}, nil
}
