package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/kshard/embeddings/bedrock"
)

func main() {
	fd, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	embed, err := bedrock.New(
		bedrock.WithModel(bedrock.TITAN_EMBED_TEXT_V2),
		bedrock.WithDimension(bedrock.EMBEDDING_SIZE_256),
	)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		word := scanner.Text()
		v32, err := embed.Embedding(context.Background(), word)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s ", word)
		for _, f := range v32 {
			fmt.Printf("%f ", f)
		}
		fmt.Println()
	}
}
