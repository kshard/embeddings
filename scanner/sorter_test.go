//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/embeddings
//

package scanner_test

import (
	"testing"

	"github.com/fogfish/golem/optics"
	"github.com/fogfish/golem/trait/seq"
	"github.com/fogfish/it/v2"
	"github.com/kshard/embeddings/scanner"
)

type obj struct {
	V string
}

func TestSorter(t *testing.T) {
	text := []obj{{"a."}, {"bb."}, {"c."}, {"ddd."}, {"ff."}}

	s := scanner.NewSorter(embed{},
		optics.ForProduct1[obj, string](),
		seq.FromSlice(text),
	)
	s.Similarity(similar)
	s.Window(3)

	it.Then(t).Should(
		it.True(s.Next()),
		it.Seq(s.Value()).Equal(obj{"a."}, obj{"c."}),
		it.True(s.Next()),
		it.Seq(s.Value()).Equal(obj{"bb."}, obj{"ff."}),
		it.True(s.Next()),
		it.Seq(s.Value()).Equal(obj{"ddd."}),
	)

	it.Then(t).ShouldNot(
		it.True(s.Next()),
	)
}
