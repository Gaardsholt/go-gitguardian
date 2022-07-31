package auditlogs

import (
	"github.com/Gaardsholt/go-gitguardian/client"
)

type AuditLogsClient struct {
	client *client.Client
}

func NewClient(opts ...client.ClientOption) (*AuditLogsClient, error) {
	client, err := client.New(opts...)
	if err != nil {
		return nil, err
	}

	return &AuditLogsClient{
		client: client,
	}, nil
}
