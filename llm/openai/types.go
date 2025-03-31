//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/embeddings
//

package openai

import (
	"fmt"
	"os/user"
	"path/filepath"

	"github.com/fogfish/gurl/v2/http"
	ø "github.com/fogfish/gurl/v2/http/send"
	"github.com/fogfish/opts"
	"github.com/jdxcode/netrc"
	"github.com/kshard/embeddings"
)

type LLM string

const (
	TEXT_EMBEDDING_3_SMALL = LLM("text-embedding-3-small")
	TEXT_EMBEDDING_3_LARGE = LLM("text-embedding-3-large")
	TEXT_ADA_002           = LLM("text-embedding-ada-002")
)

type Option = opts.Option[Client]

func (c *Client) checkRequired() error {
	return opts.Required(c,
		WithLLM(""),
		WithHTTP(nil),
	)
}

var (
	// Set OpenAI LLM
	//
	// This option is required.
	WithLLM = opts.ForType[Client, LLM]()

	// Config HTTP stack
	WithHTTP = opts.Use[Client](http.NewStack)

	// Config the host, api.openai.com is default
	WithHost = opts.ForType[Client, ø.Authority]()

	// Config API secret key
	WithSecret = opts.ForName[Client, string]("secret")

	// Set api secret from ~/.netrc file
	WithNetRC = opts.FMap(withNetRC)
)

func withNetRC(h *Client, host string) error {
	if h.secret != "" {
		return nil
	}

	usr, err := user.Current()
	if err != nil {
		return err
	}

	n, err := netrc.Parse(filepath.Join(usr.HomeDir, ".netrc"))
	if err != nil {
		return err
	}

	machine := n.Machine(host)
	if machine == nil {
		return fmt.Errorf("undefined secret for host <%s> at ~/.netrc", host)
	}

	h.secret = machine.Get("password")
	return nil
}

type Client struct {
	http.Stack
	host       ø.Authority
	secret     string
	model      LLM
	usedTokens int
}

var _ embeddings.Embedder = (*Client)(nil)

type request struct {
	Model LLM    `json:"model"`
	Text  string `json:"input"`
}

type embedding struct {
	Object  string   `json:"object"`
	Vectors []vector `json:"data"`
	Model   string   `json:"model"`
	Usage   usage    `json:"usage"`
}

type vector struct {
	Object string    `json:"object"`
	Index  int       `json:"index"`
	Vector []float32 `json:"embedding"`
}

type usage struct {
	PromptTokens int `json:"prompt_tokens"`
	UsedTokens   int `json:"total_tokens"`
}
