# scanner

Semantic Scanner implements [semantic chunking](https://github.com/FullStackRetrieval-com/RetrievalTutorials/blob/main/tutorials/LevelsOfTextSplitting/5_Levels_Of_Text_Splitting.ipynb).

## Quick Example

```go
package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"

  "github.com/kshard/embeddings/bedrock"
  "github.com/kshard/embeddings/scanner"
)

func main() {
  // Semantic scanner requires text embedding model
  embeddings, err := bedrock.New(
    bedrock.WithModel(bedrock.TITAN_EMBED_TEXT_V2),
    bedrock.WithDimension(256),
  )
  if err != nil {
    panic(err)
  }

  fd, err := os.Open("path to your files")
  if err != nil {
    panic(err)
  }

  // create and config scanner instance
  s := scanner.New(embeddings, bufio.NewScanner(fd))
  s.Similarity(scanner.HighSimilarity)
  s.Window(96)

  // scan through text
  for s.Scan() {
    text := s.Text()
    fmt.Printf("%s\n", strings.Join(text, " "))
  }

  if err := s.Err(); err != nil {
    panic(err)
  }
```
