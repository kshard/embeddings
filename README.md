<p align="center">
  <img src="./doc/golem.svg" height="240" />
  <h3 align="center">Embeddings</h3>
  <p align="center"><strong>adapter over various popular vector embeddings interfaces: AWS BedRock, OpenAI, word2vec</strong></p>

  <p align="center">
    <!-- Build Status  -->
    <a href="https://github.com/kshard/embeddings/actions/">
      <img src="https://github.com/kshard/embeddings/workflows/build/badge.svg" />
    </a>
    <!-- GitHub -->
    <a href="https://github.com/kshard/embeddings">
      <img src="https://img.shields.io/github/last-commit/kshard/embeddings.svg" />
    </a>
    <!-- Coverage -->
    <a href="https://coveralls.io/github/kshard/embeddings?branch=main">
      <img src="https://coveralls.io/repos/github/kshard/embeddings/badge.svg?branch=main" />
    </a>
    <!-- Go Card -->
    <a href="https://goreportcard.com/report/github.com/kshard/embeddings">
      <img src="https://goreportcard.com/badge/github.com/kshard/embeddings" />
    </a>
  </p>

  <table align="center">
    <thead><tr><th>sub-module</th><th>doc</th><th>about</th></tr></thead>
    <tbody>
    <!-- Module bedrock -->
    <tr><td><a href="./bedrock/">
      <img src="https://img.shields.io/github/v/tag/kshard/embeddings?label=version&filter=bedrock/*"/>
    </a></td>
    <td><a href="https://pkg.go.dev/github.com/kshard/embeddings/bedrock">
      <img src="https://img.shields.io/badge/doc-bedrock-007d9c?logo=go&logoColor=white&style=flat-square" />
    </a></td>
    <td>
    AWS Bedrock embeddings models
    </td></tr>
		<!-- Module cache -->
    <tr><td><a href="./cache/">
      <img src="https://img.shields.io/github/v/tag/kshard/embeddings?label=version&filter=cache/*"/>
    </a></td>
    <td><a href="https://pkg.go.dev/github.com/kshard/embeddings/cache">
      <img src="https://img.shields.io/badge/doc-cache-007d9c?logo=go&logoColor=white&style=flat-square" />
    </a></td>
    <td>
    Caching vector embeddings
    </td></tr>
		<!-- Module openai -->
    <tr><td><a href="./openai/">
      <img src="https://img.shields.io/github/v/tag/kshard/embeddings?label=version&filter=openai/*"/>
    </a></td>
    <td><a href="https://pkg.go.dev/github.com/kshard/embeddings/openai">
      <img src="https://img.shields.io/badge/doc-openai-007d9c?logo=go&logoColor=white&style=flat-square" />
    </a></td>
    <td>
    OpenAI embeddings models
    </td></tr>
		<!-- Module scanner -->
    <tr><td><a href="./scanner/">
      <img src="https://img.shields.io/github/v/tag/kshard/embeddings?label=version&filter=scanner/*"/>
    </a></td>
    <td><a href="https://pkg.go.dev/github.com/kshard/embeddings/scanner">
      <img src="https://img.shields.io/badge/doc-scanner-007d9c?logo=go&logoColor=white&style=flat-square" />
    </a></td>
    <td>
    Semantic chunking utility
    </td></tr>
		<!-- Module word2vec -->
    <tr><td><a href="./scanner/">
      <img src="https://img.shields.io/github/v/tag/kshard/embeddings?label=version&filter=word2vec/*"/>
    </a></td>
    <td><a href="https://pkg.go.dev/github.com/kshard/embeddings/scanner">
      <img src="https://img.shields.io/badge/doc-word2vec-007d9c?logo=go&logoColor=white&style=flat-square" />
    </a></td>
    <td>
    Word2Vec embeddings model
    </td></tr>
	</table>
</p>

---

## Inspiration

The library implements generic trait to transform text into vector embeddings.

```go
type Embeddings interface {
	Embedding(ctx context.Context, text string) ([]float32, error)
}
```

The library is structured from submodules, each implements the defined interface towards vendor. 
* [AWS BedRock embeddings](https://docs.aws.amazon.com/bedrock/latest/userguide/titan-embedding-models.html)
* [OpenAI Embeddings](https://platform.openai.com/docs/guides/embeddings) 
* [word2vec model](https://github.com/fogfish/word2vec)


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

