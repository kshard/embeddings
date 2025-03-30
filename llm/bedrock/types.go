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

// See https://docs.aws.amazon.com/bedrock/latest/userguide/model-ids.html
const (
	TITAN_EMBED_TEXT_V1 = ModelID("amazon.titan-embed-text-v1")
	TITAN_EMBED_TEXT_V2 = ModelID("amazon.titan-embed-text-v2:0")
)

const (
	EMBEDDING_SIZE_256  = 256
	EMBEDDING_SIZE_512  = 512
	EMBEDDING_SIZE_1024 = 1024
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

func WithTitanV1() Option { return WithModel(TITAN_EMBED_TEXT_V1) }
func WithTitanV2() Option { return WithModel(TITAN_EMBED_TEXT_V2) }

func WithDimension(d int) Option {
	return func(c *Client) {
		c.dimensions = d
	}
}

type Client struct {
	api            *bedrockruntime.Client
	model          ModelID
	dimensions     int
	consumedTokens int
}

type request struct {
	Text       string `json:"inputText"`
	Dimensions int    `json:"dimensions,omitempty"`
}

type embeddings struct {
	Vector         []float32 `json:"embedding"`
	UsedTextTokens int       `json:"inputTextTokenCount"`
}
