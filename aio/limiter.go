//
// Copyright (C) 2024 - 2025 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/embeddings
//

package aio

import (
	"context"
	"log/slog"

	"github.com/kshard/embeddings"
	"golang.org/x/time/rate"
)

// Rate limit startegy for LLMs I/O
type Limiter struct {
	embeddings.Embedder
	debt int
	rps  *rate.Limiter
	tps  *rate.Limiter
}

var _ embeddings.Embedder = (*Limiter)(nil)

// Create rate limit strategy for LLMs.
// It defines per minute policy for requests and tokens.
func NewLimiter(requestPerMin int, tokensPerMin int, embedder embeddings.Embedder) *Limiter {
	return &Limiter{
		Embedder: embedder,
		debt:     0,
		rps:      rate.NewLimiter(rate.Limit(requestPerMin)/60, requestPerMin),
		tps:      rate.NewLimiter(rate.Limit(tokensPerMin)/60, tokensPerMin),
	}
}

func (c *Limiter) Embedding(ctx context.Context, text string) (embeddings.Embedding, error) {
	if err := c.rps.Wait(ctx); err != nil {
		return embeddings.Embedding{}, err
	}

	if err := c.tps.WaitN(ctx, c.debt); err != nil {
		return embeddings.Embedding{}, err
	}

	reply, err := c.Embedder.Embedding(ctx, text)
	if err != nil {
		return embeddings.Embedding{}, err
	}

	c.debt = reply.UsedTokens

	slog.Debug("LLM is prompted",
		slog.Float64("budget", c.tps.Tokens()),
		slog.Int("debt", c.debt),
		slog.Int("sessionTokens", c.Embedder.UsedTokens()),
		slog.Int("replyTokens", reply.UsedTokens),
	)

	return reply, nil
}
