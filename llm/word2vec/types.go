//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/embeddings
//

package word2vec

import (
	"github.com/fogfish/opts"
	"github.com/fogfish/word2vec"
	"github.com/kshard/embeddings"
)

type Option = opts.Option[Client]

func (c *Client) checkRequired() error {
	return opts.Required(c,
		WithLLM(""),
		WithEmbeddingSize(0),
	)
}

var (
	// Set path to trained word2vec model
	//
	// This option is required.
	WithLLM = opts.ForName[Client, string]("model")

	// Set the dimension of embeddings vector
	//
	// This option is required.
	WithEmbeddingSize = opts.ForName[Client, int]("embeddingSize")
)

type Client struct {
	model          string
	embeddingSize  int
	w2v            word2vec.Model
	consumedTokens int
}

var _ embeddings.Embedder = (*Client)(nil)
