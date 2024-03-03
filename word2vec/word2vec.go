package word2vec

import (
	"context"
	"errors"

	"github.com/fogfish/word2vec"
)

// Creates word2vec embeddings client.
//
// # It requires configuration of the path to trained model
//
// The client is configurable using
//
//	WithModel(path string)
//	WithVectorSize(size int)
func New(opts ...Option) (*Client, error) {
	api := &Client{}

	defs := []Option{
		WithVectorSize(300),
	}

	for _, opt := range defs {
		opt(api)
	}

	for _, opt := range opts {
		opt(api)
	}

	if api.model == "" {
		return nil, errors.New("model is not defined")
	}

	w2v, err := word2vec.Load(api.model, api.vectorSize)
	if err != nil {
		return nil, err
	}
	api.w2v = w2v

	return api, nil
}

// Number of tokens consumed within the session
func (c *Client) ConsumedTokens() int { return c.consumedTokens }

// Calculates embedding vector
func (c *Client) Embedding(ctx context.Context, text string) ([]float32, error) {
	vec := make([]float32, c.vectorSize)
	err := c.w2v.Embedding(text, vec)
	if err != nil {
		return nil, err
	}
	return vec, nil
}
