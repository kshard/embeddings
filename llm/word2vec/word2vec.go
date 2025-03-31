//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/embeddings
//

package word2vec

import (
	"context"

	"github.com/fogfish/opts"
	"github.com/fogfish/word2vec"
	"github.com/kshard/embeddings"
)

// Creates word2vec embeddings client.
func New(opt ...Option) (*Client, error) {
	api := &Client{}

	if err := opts.Apply(api, opt); err != nil {
		return nil, err
	}

	if err := api.checkRequired(); err != nil {
		return nil, err
	}

	w2v, err := word2vec.Load(api.model, api.embeddingSize)
	if err != nil {
		return nil, err
	}
	api.w2v = w2v

	return api, nil
}

// Number of tokens consumed within the session
func (c *Client) UsedTokens() int { return c.consumedTokens }

// Calculates embedding vector
func (c *Client) Embedding(ctx context.Context, text string) (embeddings.Embedding, error) {
	vec := make([]float32, c.embeddingSize)
	err := c.w2v.Embedding(text, vec)
	if err != nil {
		return embeddings.Embedding{}, err
	}
	return embeddings.Embedding{
		Text:   text,
		Vector: vec,
	}, nil
}
