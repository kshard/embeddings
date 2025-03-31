//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/embeddings
//

package main

//
// Command line utility to build embeddings for text file
//
// txt2vec {model} {text file}
//
// It outputs embeddings as text to each file
//

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/kshard/embeddings"
	"github.com/kshard/embeddings/llm/bedrock"
	"github.com/kshard/embeddings/llm/openai"
	// "github.com/kshard/embeddings/llm/word2vec"
)

func main() {
	lpd := os.Args[1]
	ifl := os.Args[2]
	if lpd == "" || ifl == "" {
		panic(fmt.Errorf("bad input\ntxt2vec {model} {text file}"))
	}

	llm := setup(lpd)
	fd, err := os.Open(ifl)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		word := scanner.Text()
		v32, err := llm.Embedding(context.Background(), word)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s ", word)
		for _, f := range v32.Vector {
			fmt.Printf("%f ", f)
		}
		fmt.Println()
	}
}

func setup(lpd string) embeddings.Embedder {
	switch lpd {
	case "aws":
		cli, err := bedrock.New(
			bedrock.WithTitanV2,
			bedrock.WithEmbeddingSize256,
			bedrock.WithRegion("us-east-1"),
		)
		if err != nil {
			panic(err)
		}
		return cli
	case "oai":
		cli, err := openai.New(
			openai.WithLLM(openai.TEXT_ADA_002),
			openai.WithNetRC("api.openai.com"),
		)
		if err != nil {
			panic(err)
		}
		return cli
	// case "w2v":
	// 	cli, err := word2vec.New(
	// 		word2vec.WithLLM("~/devel/datasets/word2vec/t3bln-v300w5e5s1h05-en.bin"),
	// 		word2vec.WithEmbeddingSize(300),
	// 	)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	return cli
	default:
		panic(fmt.Errorf("unknown LLM provider, required one of aws, oai."))
	}
}
