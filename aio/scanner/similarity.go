//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/embeddings
//

package scanner

import (
	"github.com/chewxy/math32"
)

// High Similarity is cosine distance [0, 0.2].
// Use this range when you need very close matches (e.g., finding duplicate documents).
func HighSimilarity(a, b []float32) bool {
	x := cosine(a, b)
	return 0.0 <= x && x <= 0.2
}

// Medium Similarity is cosine distance (0.2, 0.5].
// Useful when you want to find items that are related but not identical.
func MediumSimilarity(a, b []float32) bool {
	x := cosine(a, b)
	return 0.2 < x && x <= 0.5
}

// Weak Similarity is cosine distance (0.5, 0.8].
// This range could be used for exploratory results where you want to include
// some diversity.
func WeakSimilarity(a, b []float32) bool {
	x := cosine(a, b)
	return 0.5 < x && x <= 0.8
}

// Dissimilar is cosine distance (0.8, 1.0].
// Typically, these items are unrelated, and you might filter them out unless
// dissimilarity is desirable (e.g., in anomaly detection).
func Dissimilar(a, b []float32) bool {
	x := cosine(a, b)
	return 0.8 < x && x <= 1.0
}

// Similarity on custom cosine distance [lo, hi].
// Use this range when you need custom interval.
func RangeSimilarity(lo, hi float32) func(a, b []float32) bool {
	return func(a, b []float32) bool {
		x := cosine(a, b)
		return lo <= x && x <= hi
	}
}

// Similarity with custom assert of cosine distance
func CosineSimilarity(f func(float32) bool) func(a, b []float32) bool {
	return func(a, b []float32) bool {
		return f(cosine(a, b))
	}
}

func cosine(a, b []float32) (d float32) {
	if len(a) != len(b) {
		panic("vectors must have equal lengths")
	}

	if len(a)%4 != 0 {
		panic("vector length must be multiple of 4")
	}

	ab := float32(0.0)
	aa := float32(0.0)
	bb := float32(0.0)

	for i := 0; i < len(a); i += 4 {
		asl := a[i : i+4 : i+4]
		bsl := b[i : i+4 : i+4]

		ax0, ax1, ax2, ax3 := asl[0], asl[1], asl[2], asl[3]
		bx0, bx1, bx2, bx3 := bsl[0], bsl[1], bsl[2], bsl[3]

		ab0 := ax0 * bx0
		ab1 := ax1 * bx1
		ab2 := ax2 * bx2
		ab3 := ax3 * bx3
		ab += ab0 + ab1 + ab2 + ab3

		aa0 := ax0 * ax0
		aa1 := ax1 * ax1
		aa2 := ax2 * ax2
		aa3 := ax3 * ax3
		aa += aa0 + aa1 + aa2 + aa3

		bb0 := bx0 * bx0
		bb1 := bx1 * bx1
		bb2 := bx2 * bx2
		bb3 := bx3 * bx3
		bb += bb0 + bb1 + bb2 + bb3
	}

	s := math32.Sqrt(aa) * math32.Sqrt(bb)

	// Note: two proportional vectors have a cosine similarity of 1 |d|=0
	//       two orthogonal vectors have a similarity of 0          |d|=0.5
	//       and two opposite vectors have a similarity of -1.      |d|=1.0
	d = (1 - ab/s) / 2

	return
}
