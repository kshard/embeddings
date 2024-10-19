//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/embeddings
//

package scanner

import (
	"context"
	"fmt"

	"github.com/fogfish/golem/optics"
	"github.com/fogfish/golem/trait/seq"
	"github.com/kshard/embeddings"
)

// Sorter provides a convenient solution for semantic sorting.
//
// Successive calls to the Sorter.Sort method will step through the context
// windows of a slice and grouping 'sentences' semantically. The context window
// is defined either by number sentences, use Window method to change
// default 32 sentences value.
//
// The input slice is assumed to be split into sentences already.
//
// The sorter uses embeddings to determine similarity. Use Similarity method
// to change the default high cosine similarity to own implementation.
// The module provides high, medium, weak and dissimilarity functions based on
// cosine distance.
type Sorter[T any] struct {
	embed             embeddings.Embeddings
	similarity        func([]float32, []float32) bool
	windowInSentences int
	scanner           seq.Seq[T]
	lens              optics.Lens[T, string]
	err               error
	eof               bool
	window            []typed[T]
	cursor            []T
}

type typed[T any] struct {
	object T
	vector []float32
}

// Creates new instance of Sorter to read from seq.Seq[T] and using embedding.
func NewSorter[T any](embed embeddings.Embeddings, lens optics.Lens[T, string], seq seq.Seq[T]) *Sorter[T] {
	return &Sorter[T]{
		embed:             embed,
		similarity:        HighSimilarity,
		windowInSentences: 32,
		scanner:           seq,
		lens:              lens,
		window:            make([]typed[T], 0),
	}
}

// Similarity sets the similarity function for the Scanner.
// The default is HighSimilarity.
func (s *Sorter[T]) Similarity(f func([]float32, []float32) bool) {
	s.similarity = f
}

// Widow defines the context window for similarity detection.
// The default value is 32 sentences.
func (s *Sorter[T]) Window(n int) {
	s.windowInSentences = n
}

func (s *Sorter[T]) Err() error { return s.err }
func (s *Sorter[T]) Value() []T { return s.cursor }

// Next advances the Sorter through context window, sequences will be available
// through [Scanner.Text]. It returns false if there was I/O error or EOF is reached.
func (s *Sorter[T]) Next() bool {
	if s.err != nil {
		return false
	}

	if !s.eof {
		s.eof, s.err = s.fill()
		if s.err != nil {
			return false
		}
	}

	s.cursor = s.peek()

	return !(s.eof && len(s.cursor) == 0)
}

// fill the window
func (s *Sorter[T]) fill() (bool, error) {
	wn := s.windowInSentences - len(s.window)

	has := s.scanner != nil
	for ; wn > 0 && has; has = s.scanner.Next() {
		// for wn > 0 && s.scanner.Next() {
		obj := s.scanner.Value()
		txt := s.lens.Get(&obj)
		v32, err := s.embed.Embedding(context.Background(), txt)
		if err != nil {
			return false, fmt.Errorf("embedding has failed: %w, for {%s}", err, txt)
		}

		s.window = append(s.window, typed[T]{object: obj, vector: v32})
		wn--
	}

	return !has || wn != 0, nil
}

// peek similar from the window
func (s *Sorter[T]) peek() []T {
	if len(s.window) == 0 {
		return nil
	}

	a, b := make([]typed[T], 0), make([]typed[T], 0)
	a = append(a, s.window[0])

	for i := 1; i < len(s.window); i++ {
		tail := a[len(a)-1]
		if s.similarity(tail.vector, s.window[i].vector) {
			a = append(a, s.window[i])
		} else {
			b = append(b, s.window[i])
		}
	}

	s.window = b

	seq := make([]T, len(a))
	for i, x := range a {
		seq[i] = x.object
	}
	return seq
}
