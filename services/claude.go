package services

import (
	"github.com/madebywelch/anthropic-go/v2/pkg/anthropic"
)

type Claude struct {
	client *anthropic.Client
}

func NewClaude(apiKey string) *Claude {
	client, err := anthropic.NewClient(apiKey)
	if err != nil {
		panic(err)
	}

	return &Claude{
		client: client,
	}
}
