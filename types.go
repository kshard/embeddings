//
// Copyright (C) 2024 - 2025 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/embeddings
//

package embeddings

import "context"

type Embedder interface {
	UsedTokens() int
	Embedding(ctx context.Context, text string) (Embedding, error)
}

// Embeddings
type Embedding struct {
	Text       string
	Vector     []float32
	UsedTokens int
}
