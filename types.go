package embeddings

import "context"

type Embeddings interface {
	ConsumedTokens() int
	Embedding(ctx context.Context, text string) ([]float32, error)
}
