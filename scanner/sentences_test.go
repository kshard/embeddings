//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/embeddings
//

package scanner_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/fogfish/it/v2"
	"github.com/kshard/embeddings/scanner"
)

func TestSplitSentences(t *testing.T) {
	for input, expected := range map[string][]string{
		"Hello World!":      {"Hello World!"},
		"Hello! World.":     {"Hello!", "World."},
		"Hello!\nWorld.":    {"Hello!", "World."},
		`Hello!\xWorld.`:    {`Hello!\xWorld.`},
		"Hello 3.14 World!": {"Hello 3.14 World!"},
		"Hello! World 3.14": {"Hello!", "World 3.14"},
	} {
		s := bufio.NewScanner(strings.NewReader(input))
		s.Split(scanner.ScanSentence)

		seq := make([]string, 0)
		for s.Scan() {
			seq = append(seq, s.Text())
		}

		it.Then(t).Should(
			it.Seq(seq).Equal(expected...),
		)
	}
}
