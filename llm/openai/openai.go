//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/embeddings
//

package openai

import (
	"context"
	"errors"

	"github.com/fogfish/gurl/v2/http"
	ƒ "github.com/fogfish/gurl/v2/http/recv"
	ø "github.com/fogfish/gurl/v2/http/send"
	"github.com/fogfish/opts"
	"github.com/kshard/embeddings"
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
//	WithModel(...)
//	WithHTTP(opts ...http.Config)
func New(opt ...Option) (*Client, error) {
	api := &Client{
		host: ø.Authority("https://api.openai.com"),
	}

	if err := opts.Apply(api, opt); err != nil {
		return nil, err
	}

	if api.Stack == nil {
		api.Stack = http.New()
	}

	return api, api.checkRequired()
}

// Number of tokens consumed within the session
func (c *Client) UsedTokens() int { return c.usedTokens }

// Calculates embedding vector
func (c *Client) Embedding(ctx context.Context, text string) (embeddings.Embedding, error) {
	bag, err := http.IO[embedding](c.WithContext(ctx),
		http.POST(
			ø.URI("%s/v1/embeddings", c.host),
			ø.Accept.JSON,
			ø.Authorization.Set("Bearer "+c.secret),
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
		return embeddings.Embedding{}, err
	}

	if len(bag.Vectors) != 1 {
		return embeddings.Embedding{}, errors.New("invalid response")
	}

	c.usedTokens += bag.Usage.UsedTokens

	return embeddings.Embedding{
		Text:       text,
		Vector:     bag.Vectors[0].Vector,
		UsedTokens: bag.Usage.UsedTokens,
	}, nil
}
