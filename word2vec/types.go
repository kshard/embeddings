//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/embeddings
//

package word2vec

import (
	"github.com/fogfish/word2vec"
)

type Option func(*Client)

func WithModel(path string) Option {
	return func(c *Client) {
		c.model = path
	}
}

func WithVectorSize(size int) Option {
	return func(c *Client) {
		c.vectorSize = size
	}
}

type Client struct {
	model          string
	vectorSize     int
	w2v            word2vec.Model
	consumedTokens int
}
