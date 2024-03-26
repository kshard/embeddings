//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/embeddings
//

package bedrock

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

type Option func(*Client)

type ModelID string

const (
	TITAN_EMBED_TEXT_V1 = ModelID("amazon.titan-embed-text-v1")
)

func WithConfig(cfg aws.Config) Option {
	return func(c *Client) {
		c.api = bedrockruntime.NewFromConfig(cfg)
	}
}

func WithModel(id ModelID) Option {
	return func(c *Client) {
		c.model = id
	}
}

type Client struct {
	api            *bedrockruntime.Client
	model          ModelID
	consumedTokens int
}

type request struct {
	Text string `json:"inputText"`
}

type embeddings struct {
	Vector         []float32 `json:"embedding"`
	UsedTextTokens int       `json:"inputTextTokenCount"`
}
