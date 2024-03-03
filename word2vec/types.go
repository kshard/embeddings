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
