package bedrock

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

type Option func(*Client)

func WithConfig(cfg aws.Config) Option {
	return func(e *Client) {
		e.api = bedrockruntime.NewFromConfig(cfg)
	}
}

type Client struct {
	api            *bedrockruntime.Client
	model          string
	consumedTokens int
}

type request struct {
	Text string `json:"inputText"`
}

type embeddings struct {
	Vector         []float32 `json:"embedding"`
	UsedTextTokens int       `json:"inputTextTokenCount"`
}
