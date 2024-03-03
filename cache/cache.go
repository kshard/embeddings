//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/embeddings
//

package cache

import (
	"context"
	"encoding/binary"
	"log/slog"
	"math"

	"github.com/kshard/atom"
	"github.com/kshard/embeddings"
)

// Creates caching layer for embeddings client.
//
// Use github.com/akrylysov/pogreb to cache embedding on local file systems:
//
//	cli, err := /* create embeddings client */
//	db, err := pogreb.Open("embeddings.cache", nil)
//	text := cache.New(db, cli)
func New(cache Cache, embed embeddings.Embeddings) *Client {
	return &Client{
		atoms: atom.New(atom.NewPermanentMap(cache)),
		embed: embed,
		cache: cache,
	}
}

func (c *Client) ConsumedTokens() int { return c.embed.ConsumedTokens() }

// Calculates embedding vector
func (c *Client) Embedding(ctx context.Context, text string) ([]float32, error) {
	code, err := c.atoms.Atom(text)
	if err != nil {
		return nil, err
	}

	var bkey [5]byte
	bkey[0] = 'f'
	binary.LittleEndian.PutUint32(bkey[1:], code)

	val, err := c.cache.Get(bkey[:])
	if err != nil {
		return nil, err
	}

	if len(val) != 0 {
		return decodeFVec(val), nil
	}

	vec, err := c.embed.Embedding(ctx, text)
	if err != nil {
		return nil, err
	}

	err = c.cache.Put(bkey[:], encodeFVec(vec))
	if err != nil {
		slog.Warn("failed to cache vector")
	}

	return vec, nil
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
