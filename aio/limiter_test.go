//
// Copyright (C) 2025 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/chatter
//

package aio_test

import (
	"context"
	"testing"
	"time"

	"github.com/fogfish/it/v2"
	"github.com/kshard/embeddings"
	"github.com/kshard/embeddings/aio"
)

func mockTokensUsage(n int) mock {
	return mock{embeddings.Embedding{Vector: nil, UsedTokens: n}}
}

func TestLimiter(t *testing.T) {
	n := 8

	t.Run("RequestPerMinute", func(t *testing.T) {
		rpm := n
		tpm := 100000
		api := aio.NewLimiter(rpm, tpm, mockTokensUsage(1000))

		prompt := func() error {
			ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
			defer cancel()

			_, err := api.Embedding(ctx, "text")
			return err
		}

		for range rpm {
			err := prompt()
			it.Then(t).Should(it.Nil(err))
		}

		err := prompt()
		it.Then(t).ShouldNot(it.Nil(err))
	})

	t.Run("TokensPerMinute", func(t *testing.T) {
		rpm := 100000
		tpm := n * 1000
		api := aio.NewLimiter(rpm, tpm, mockTokensUsage(1000))

		prompt := func() error {
			ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
			defer cancel()

			_, err := api.Embedding(ctx, "text")
			return err
		}

		// +1 due to debt model
		for range n + 1 {
			err := prompt()
			it.Then(t).Should(it.Nil(err))
		}

		err := prompt()
		it.Then(t).ShouldNot(it.Nil(err))
	})
}
