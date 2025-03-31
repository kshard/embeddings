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
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/fogfish/opts"
	"github.com/kshard/embeddings"
)

// Creates AWS BedRock embeddings client.
//
// By default `us-west-2` region is used, supply custom `aws.Config`
// to alter behavior.
//
// The client is configurable using
//
//	WithConfig(cfg aws.Config)
func New(opt ...Option) (*Client, error) {
	c := &Client{}

	if err := opts.Apply(c, opt); err != nil {
		return nil, err
	}

	if c.api == nil {
		if err := optsFromRegion(c, defaultRegion); err != nil {
			return nil, err
		}
	}

	return c, c.checkRequired()
}

// Number of tokens consumed within the session
func (c *Client) UsedTokens() int { return c.usedTokens }

// Calculates embedding vector
func (c *Client) Embedding(ctx context.Context, text string) (embeddings.Embedding, error) {
	body, err := json.Marshal(
		request{
			Text:       text,
			Dimensions: c.embeddingSize,
		},
	)
	if err != nil {
		return embeddings.Embedding{}, err
	}

	req := &bedrockruntime.InvokeModelInput{
		ModelId:     aws.String(string(c.model)),
		ContentType: aws.String("application/json"),
		Body:        body,
	}

	result, err := c.api.InvokeModel(ctx, req)
	if err != nil {
		return embeddings.Embedding{}, err
	}

	var embedding embedding
	if err := json.Unmarshal(result.Body, &embedding); err != nil {
		return embeddings.Embedding{}, err
	}

	c.usedTokens += embedding.UsedTextTokens

	return embeddings.Embedding{
		Text:       text,
		Vector:     embedding.Vector,
		UsedTokens: embedding.UsedTextTokens,
	}, nil
}
