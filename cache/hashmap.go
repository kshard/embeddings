//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/embeddings
//

package cache

import (
	"encoding/binary"
	"unsafe"

	"github.com/kshard/atom"
)

type hashmap struct {
	cache Cache
}

func atomHashMap(cache Cache) atom.HashMap {
	return &hashmap{
		cache: cache,
	}
}

func (m *hashmap) Get(key atom.Atom) (string, error) {
	var bkey [5]byte
	bkey[0] = 'a'
	binary.LittleEndian.PutUint32(bkey[1:], key)

	val, err := m.cache.Get(bkey[:])
	if err != nil {
		return "", err
	}

	// This is copied from runtime. It relies on the string
	// header being a prefix of the slice header!
	s := *(*string)(unsafe.Pointer(&val))

	return s, nil
}

func (m *hashmap) Put(key atom.Atom, val string) error {
	var bkey [5]byte
	bkey[0] = 'a'
	binary.LittleEndian.PutUint32(bkey[1:], key)

	// This is copied from runtime. It relies on the string
	// header being a prefix of the slice header!
	bval := *(*[]byte)(unsafe.Pointer(&val))

	err := m.cache.Put(bkey[:], bval)
	if err != nil {
		return err
	}

	return nil
}
