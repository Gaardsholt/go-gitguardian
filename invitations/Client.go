package invitations

import (
	"github.com/Gaardsholt/go-gitguardian/client"
)

type InvitationsClient struct {
	client *client.Client
}

func NewClient(opts ...client.ClientOption) (*InvitationsClient, error) {
	client, err := client.New(opts...)
	if err != nil {
		return nil, err
	}

	return &InvitationsClient{
		client: client,
	}, nil
}
