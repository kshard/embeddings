package cache

import (
	"github.com/kshard/atom"
	"github.com/kshard/embeddings"
)

// Getter interface abstract storage
type Getter interface{ Get([]byte) ([]byte, error) }

// Setter interface abstract storage
type Putter interface{ Put([]byte, []byte) error }

// Cache interface
type Cache interface {
	Getter
	Putter
}

type Client struct {
	atoms *atom.Pool
	embed embeddings.Embeddings
	cache Cache
}
