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

	"github.com/kshard/embeddings"
)

// Scanner provides a convenient solution for semantic chunking.
// Successive calls to the Scanner.Scan method will step through the context
// windows of a file and grouping sentences semantically. The context window
// is defined either by the length of text or number sentences, use Window
// method to change default 4K and 32 sentences value. The specification of
// a sentence is defined by a split function of type SplitFunc; the default
// split function breaks the input into sentences using punctuation runes.
// Use Split function to define own algorithms.
//
// The scanner uses embeddings to determine similarity. Use Similarity method
// to change the default high cosine similarity to own implementation.
// The module provides high, medium, weak and dissimilarity functions based on
// cosine distance.
//
// Scanning stops unrecoverably at EOF or the first I/O error.
type Scanner struct {
	embed             embeddings.Embeddings
	similarity        func([]float32, []float32) bool
	windowInBytes     int
	windowInSentences int
	scanner           Reader
	err               error
	cursor            int
	text              [][]string
}

// Reader is an interface similar to [bufio.Scanner].
// It defines core functionality used by semantic chunking.
type Reader interface {
	Scan() bool
	Text() string
	Err() error
}

// Creates new instance of Scanner to read from io.Reader and using embedding.
func New(embed embeddings.Embeddings, r Reader) *Scanner {
	return &Scanner{
		embed:             embed,
		similarity:        HighSimilarity,
		windowInBytes:     4096,
		windowInSentences: 32,
		scanner:           r,
		cursor:            -1,
	}
}

// Similarity sets the similarity function for the Scanner.
// The default is HighSimilarity.
func (s *Scanner) Similarity(f func([]float32, []float32) bool) {
	s.similarity = f
}

// Widow defines the context window for similarity detection.
// The default value is either 4K bytes or 32 sentences.
func (s *Scanner) Window(n int, size int) {
	s.windowInBytes = size
	s.windowInSentences = n
}

func (s *Scanner) Err() error     { return s.err }
func (s *Scanner) Text() []string { return s.text[s.cursor] }

// Scan advances the Scanner through context window, sequences will be available
// through [Scanner.Text]. It returns false if there was I/O error or EOF is reached.
func (s *Scanner) Scan() bool {
	if s.text == nil || s.cursor == len(s.text)-1 {
		s.text, s.err = s.read()
		if s.err != nil {
			return false
		}
		s.cursor = -1
	}
	s.cursor++

	return s.cursor < len(s.text)
}

func (s *Scanner) read() ([][]string, error) {
	var seq []string
	var vec [][]float32

	wb := s.windowInBytes
	wn := s.windowInSentences
	for s.scanner.Scan() && wb > 0 && wn > 0 {
		txt := s.scanner.Text()
		v32, err := s.embed.Embedding(context.Background(), txt)
		if err != nil {
			return nil, err
		}

		seq = append(seq, txt)
		vec = append(vec, v32)
		wb -= len(txt)
		wn--
	}

	if err := s.scanner.Err(); err != nil {
		return nil, err
	}

	return s.groupBy(seq, vec), nil
}

func (s *Scanner) groupBy(seq []string, vec [][]float32) [][]string {
	var groups [][]string
	visited := make([]bool, len(vec))

	for i, p1 := range seq {
		if visited[i] {
			continue
		}
		// Start a new group with point p1
		group := []string{p1}
		visited[i] = true

		for j, p2 := range seq {
			if !visited[j] && s.similarity(vec[i], vec[j]) {
				group = append(group, p2)
				visited[j] = true
			}
		}
		groups = append(groups, group)
	}

	return groups
}
