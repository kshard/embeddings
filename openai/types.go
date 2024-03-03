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

func WithModelText3Small() Option {
	return func(c *Client) {
		c.model = "text-embedding-3-small"
	}
}

func WithModelText3Large() Option {
	return func(c *Client) {
		c.model = "text-embedding-3-large"
	}
}

func WithModelTextAda002() Option {
	return func(c *Client) {
		c.model = "text-embedding-ada-002"
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
	model          string
	consumedTokens int
}

type request struct {
	Model string `json:"model"`
	Text  string `json:"input"`
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
