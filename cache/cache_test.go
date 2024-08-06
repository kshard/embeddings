//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/embeddings
//

package cache_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/kshard/embeddings/cache"
)

func TestCache(t *testing.T) {
	kv := keyval{}
	c := cache.New(kv, embed{})

	if c.ConsumedTokens() != 10 {
		t.Errorf("unexpected ConsumedTokens output")
	}

	text := "hello world"
	c.Embedding(context.Background(), text)
	c.Embedding(context.Background(), text)

	for k := range kv {
		if !bytes.Equal([]byte(k), c.HashKey(text)) {
			t.Errorf("unexpected key")
		}
	}
}

// mock embedding client
type embed struct{}

func (embed) ConsumedTokens() int { return 10 }

func (embed) Embedding(ctx context.Context, text string) ([]float32, error) {
	return []float32{0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0}, nil
}

// mock key-value
type keyval map[string][]byte

func (kv keyval) Get(key []byte) ([]byte, error) {
	if val, has := kv[string(key)]; has {
		return val, nil
	}

	return nil, nil
}

// Setter interface abstract storage
func (kv keyval) Put(key []byte, val []byte) error {
	kv[string(key)] = val
	return nil
}
