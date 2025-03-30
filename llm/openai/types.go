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
	"github.com/jdxcode/netrc"
)

type Option func(*Client)

type ModelID string

const (
	TEXT_EMBEDDING_3_SMALL = ModelID("text-embedding-3-small")
	TEXT_EMBEDDING_3_LARGE = ModelID("text-embedding-3-large")
	TEXT_ADA_002           = ModelID("text-embedding-ada-002")
)

func WithModel(id ModelID) Option {
	return func(c *Client) {
		c.model = id
	}
}

func WithHTTP(opts ...http.Config) Option {
	return func(c *Client) {
		c.Stack = http.New(opts...)
	}
}

func WithSecret(secret string) Option {
	return func(c *Client) {
		c.secret = "Bearer " + secret
	}
}

func WithNetRC(host string) Option {
	return func(c *Client) {
		if c.secret != "" {
			return
		}

		usr, err := user.Current()
		if err != nil {
			panic(err)
		}

		n, err := netrc.Parse(filepath.Join(usr.HomeDir, ".netrc"))
		if err != nil {
			panic(err)
		}

		machine := n.Machine(host)
		if machine == nil {
			panic(fmt.Errorf("undefined secret for host <%s> at ~/.netrc", host))
		}

		c.secret = "Bearer " + machine.Get("password")
	}
}

type Client struct {
	http.Stack
	host           ø.Authority
	secret         string
	model          ModelID
	consumedTokens int
}

type request struct {
	Model ModelID `json:"model"`
	Text  string  `json:"input"`
}

type embeddings struct {
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
