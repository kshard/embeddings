# embeddings

The library is adapter over various popular vector embedding interfaces: AWS BedRock, OpenAI, word2vec.

[![Version](https://img.shields.io/github/v/tag/kshard/embeddings?label=version&filter=bedrock/*)](https://github.com/kshard/embeddings/releases)
[![Version](https://img.shields.io/github/v/tag/kshard/embeddings?label=version&filter=openai/*)](https://github.com/kshard/embeddings/releases)
[![Version](https://img.shields.io/github/v/tag/kshard/embeddings?label=version&filter=word2vec/*)](https://github.com/kshard/embeddings/releases)
[![Version](https://img.shields.io/github/v/tag/kshard/embeddings?label=version&filter=cache/*)](https://github.com/kshard/embeddings/releases)
[![Documentation](https://pkg.go.dev/badge/github.com/kshard/embeddings)](https://pkg.go.dev/github.com/kshard/embeddings)
[![Build Status](https://github.com/kshard/embeddings/workflows/build/badge.svg)](https://github.com/kshard/embeddings/actions/)
[![Git Hub](https://img.shields.io/github/last-commit/kshard/embeddings.svg)](https://github.com/kshard/embeddings)
[![Coverage Status](https://coveralls.io/repos/github/kshard/embeddings/badge.svg?branch=main)](https://coveralls.io/github/kshard/embeddings?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/kshard/embeddings)](https://goreportcard.com/report/github.com/kshard/embeddings)

## Inspiration

The library implements generic trait to transform text into vector embeddings.

```go
type Embeddings interface {
	Embedding(ctx context.Context, text string) ([]float32, error)
}
```

The library is structured from submodules, each implements the defined interface towards vendor. 
* [github.com/kshard/embeddings/bedrock](./bedrock/) adapts [AWS BedRock embeddings](https://docs.aws.amazon.com/bedrock/latest/userguide/titan-embedding-models.html)
* [github.com/kshard/embeddings/openai](./openai/) adapts [OpenAI Embeddings](https://platform.openai.com/docs/guides/embeddings) 
* [github.com/kshard/embeddings/word2vec](./word2vec/) adapts [word2vec model](https://github.com/fogfish/word2vec)
* [github.com/kshard/embeddings/cache](./cache/) caching embeddings either in-memory or persistently.


## Getting started

The latest version of the library is available at `main` branch of this repository. All development, including new features and bug fixes, take place on the `main` branch using forking and pull requests as described in contribution guidelines. The stable version is available via Golang modules.

```go
import embeddings "github.com/kshard/embeddings/{your-model}"

text, err := embeddings.New(/* config options */)

// Calculate embeddings
vector, err := text.Embedding(context.Background(), "text embeddings")

// Checks number of tokens consumed by active sessions
text.ConsumedTokens()
```

## How To Contribute

The library is [MIT](LICENSE) licensed and accepts contributions via GitHub pull requests:

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Added some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

The build and testing process requires [Go](https://golang.org) version 1.13 or later.

**build** and **test** library.

```bash
git clone https://github.com/kshard/embeddings
cd embeddings
go test ./...
```

### commit message

The commit message helps us to write a good release note, speed-up review process. The message should address two question what changed and why. The project follows the template defined by chapter [Contributing to a Project](http://git-scm.com/book/ch5-2.html) of Git book.

### bugs

If you experience any issues with the library, please let us know via [GitHub issues](https://github.com/kshard/embeddings/issue). We appreciate detailed and accurate reports that help us to identity and replicate the issue. 


## License

[![See LICENSE](https://img.shields.io/github/license/kshard/embeddings.svg?style=for-the-badge)](LICENSE)

