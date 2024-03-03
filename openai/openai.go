package openai

import (
	"context"
	"errors"

	"github.com/fogfish/gurl/v2/http"
	ƒ "github.com/fogfish/gurl/v2/http/recv"
	ø "github.com/fogfish/gurl/v2/http/send"
)

// Creates OpenAI embeddings client.
//
// By default OpenAI reads access token from `~/.netrc`,
// supply custom secret `WithSecret(secret string)` if needed.
//
// The client is configurable using
//
//	WithSecret(secret string)
//	WithNetRC(host string)
//	WithModelText3Small()
//	WithModelText3Large()
//	WithModelTextAda002()
//	WithHTTP(opts ...http.Config)
func New(opts ...Option) (*Client, error) {
	api := &Client{
		host: ø.Authority("api.openai.com"),
	}

	defs := []Option{
		WithModelText3Small(),
		WithNetRC(string(api.host)),
	}

	for _, opt := range defs {
		opt(api)
	}

	for _, opt := range opts {
		opt(api)
	}

	if api.Stack == nil {
		api.Stack = http.New()
	}

	return api, nil
}

// Number of tokens consumed within the session
func (c *Client) ConsumedTokens() int { return c.consumedTokens }

// Calculates embedding vector
func (c *Client) Embedding(ctx context.Context, text string) ([]float32, error) {
	bag, err := http.IO[embeddings](c.WithContext(ctx),
		http.POST(
			ø.URI("https://%s/v1/embeddings", c.host),
			ø.Accept.JSON,
			ø.Authorization.Set(c.secret),
			ø.ContentType.JSON,
			ø.Send(request{
				Model: c.model,
				Text:  text,
			}),

			ƒ.Status.OK,
			ƒ.ContentType.JSON,
		),
	)
	if err != nil {
		return nil, err
	}

	if len(bag.Vectors) != 1 {
		return nil, errors.New("invalid response")
	}

	c.consumedTokens += bag.Usage.UsedTokens

	return bag.Vectors[0].Vector, nil
}
