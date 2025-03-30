//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/embeddings
//

package aio_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/kshard/embeddings/aio"
)

func TestCache(t *testing.T) {
	kv := keyval{}
	c := aio.NewCache(kv, mockVector())

	if c.UsedTokens() != 10 {
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
