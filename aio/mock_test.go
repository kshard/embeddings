//
// Copyright (C) 2024 - 2025 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/embeddings
//

package aio_test

import (
	"context"

	"github.com/kshard/embeddings"
)

// mock embedding client
type mock struct {
	reply embeddings.Embedding
}

func mockVector() mock {
	return mock{
		embeddings.Embedding{
			UsedTokens: 10,
			Vector:     []float32{0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0},
		},
	}
}

func (mock mock) UsedTokens() int { return mock.reply.UsedTokens }

func (mock mock) Embedding(ctx context.Context, text string) (embeddings.Embedding, error) {
	return mock.reply, nil
}
