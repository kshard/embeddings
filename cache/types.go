//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/embeddings
//

package cache

import (
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
	embed embeddings.Embeddings
	cache Cache
}
