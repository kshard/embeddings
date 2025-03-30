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
	"crypto/sha1"
	"encoding/binary"
	"log/slog"
	"math"

	"github.com/kshard/embeddings"
)

// Getter interface abstract storage
type Getter interface{ Get([]byte) ([]byte, error) }

// Setter interface abstract storage
type Putter interface{ Put([]byte, []byte) error }

// KeyVal interface
type KeyVal interface {
	Getter
	Putter
}

type Cache struct {
	embeddings.Embedder
	cache KeyVal
}

var _ embeddings.Embedder = (*Cache)(nil)

// Creates caching layer for embeddings client.
//
// Use github.com/akrylysov/pogreb to cache embedding on local file systems:
//
//	cli, err := /* create embeddings client */
//	db, err := pogreb.Open("embeddings.cache", nil)
//	text := cache.NewCache(db, cli)
func NewCache(cache KeyVal, embedder embeddings.Embedder) *Cache {
	return &Cache{
		Embedder: embedder,
		cache:    cache,
	}
}

func (c *Cache) HashKey(text string) []byte {
	hash := sha1.New()
	hash.Write([]byte(text))
	return hash.Sum(nil)
}

// Calculates embedding vector
func (c *Cache) Embedding(ctx context.Context, text string) (embeddings.Embedding, error) {
	hkey := c.HashKey(text)

	val, err := c.cache.Get(hkey)
	if err != nil {
		return embeddings.Embedding{}, err
	}

	if len(val) != 0 {
		return embeddings.Embedding{Vector: decodeFVec(val)}, nil
	}

	reply, err := c.Embedder.Embedding(ctx, text)
	if err != nil {
		return embeddings.Embedding{}, err
	}

	err = c.cache.Put(hkey, encodeFVec(reply.Vector))
	if err != nil {
		slog.Warn("failed to cache vector", "error", err)
	}

	return reply, nil
}

func encodeFVec(v []float32) []byte {
	b := make([]byte, len(v)*4)

	p := 0
	for i := 0; i < len(v); i++ {
		u := math.Float32bits(v[i])
		binary.LittleEndian.PutUint32(b[p:p+4], u)

		p += 4
	}

	return b
}

func decodeFVec(b []byte) []float32 {
	v := make([]float32, len(b)/4)

	p := 0
	for i := 0; i < len(b); i += 4 {
		v[p] = math.Float32frombits(binary.LittleEndian.Uint32(b[i : i+4]))
		p++
	}

	return v
}
