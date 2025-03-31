//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/embeddings
//

package bedrock

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/fogfish/opts"
	"github.com/kshard/embeddings"
)

type LLM string

// See https://docs.aws.amazon.com/bedrock/latest/userguide/model-ids.html
const (
	TITAN_EMBED_TEXT_V1 = LLM("amazon.titan-embed-text-v1")
	TITAN_EMBED_TEXT_V2 = LLM("amazon.titan-embed-text-v2:0")
)

const (
	EMBEDDING_SIZE_256  = 256
	EMBEDDING_SIZE_512  = 512
	EMBEDDING_SIZE_1024 = 1024
)

type Option = opts.Option[Client]

func (c *Client) checkRequired() error {
	return opts.Required(c,
		WithLLM(""),
		WithBedrock(nil),
	)
}

const defaultRegion = "us-west-2"

var (
	// Set AWS Bedrock Foundational LLM
	//
	// This option is required.
	WithLLM     = opts.ForType[Client, LLM]()
	WithTitanV1 = WithLLM(TITAN_EMBED_TEXT_V1)
	WithTitanV2 = WithLLM(TITAN_EMBED_TEXT_V2)

	// Set the dimension of embeddings vector
	WithEmbeddingSize     = opts.ForName[Client, int]("embeddingSize")
	WithEmbeddingSize256  = WithEmbeddingSize(EMBEDDING_SIZE_256)
	WithEmbeddingSize512  = WithEmbeddingSize(EMBEDDING_SIZE_512)
	WithEmbeddingSize1024 = WithEmbeddingSize(EMBEDDING_SIZE_1024)

	// Use aws.Config to config the client
	WithConfig = opts.FMap(optsFromConfig)

	// Use region for aws.Config
	WithRegion = opts.FMap(optsFromRegion)

	// Set us-west-2 as default region
	WithDefaultRegion = WithRegion(defaultRegion)

	// Set AWS Bedrock Runtime
	WithBedrock = opts.ForType[Client, Bedrock]()
)

func optsFromRegion(c *Client, region string) error {
	cfg, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(region),
	)
	if err != nil {
		return err
	}

	return optsFromConfig(c, cfg)
}

func optsFromConfig(c *Client, cfg aws.Config) (err error) {
	if c.api == nil {
		c.api = bedrockruntime.NewFromConfig(cfg)
	}

	return
}

type Bedrock interface {
	InvokeModel(ctx context.Context, params *bedrockruntime.InvokeModelInput, optFns ...func(*bedrockruntime.Options)) (*bedrockruntime.InvokeModelOutput, error)
}

type Client struct {
	api           Bedrock
	model         LLM
	embeddingSize int
	usedTokens    int
}

var _ embeddings.Embedder = (*Client)(nil)

type request struct {
	Text       string `json:"inputText"`
	Dimensions int    `json:"dimensions,omitempty"`
}

type embedding struct {
	Vector         []float32 `json:"embedding"`
	UsedTextTokens int       `json:"inputTextTokenCount"`
}
