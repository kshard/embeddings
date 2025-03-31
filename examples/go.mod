module github.com/kshard/embeddings/examples

go 1.23.0

replace github.com/kshard/embeddings/llm/bedrock => ../llm/bedrock

replace github.com/kshard/embeddings/llm/openai => ../llm/openai

// replace github.com/kshard/embeddings/llm/word2vec => ../llm/word2vec
// replace github.com/fogfish/word2vec => ./word2vec

require (
	github.com/kshard/embeddings v0.2.0
	github.com/kshard/embeddings/llm/bedrock v0.0.0
	github.com/kshard/embeddings/llm/openai v0.0.0
// github.com/kshard/embeddings/llm/word2vec v0.0.0
)

require (
	github.com/Masterminds/semver/v3 v3.2.1 // indirect
	github.com/ajg/form v1.5.2-0.20200323032839-9aeb3cf462e1 // indirect
	github.com/aws/aws-cdk-go/awscdk/v2 v2.150.0 // indirect
	github.com/aws/aws-sdk-go-v2 v1.25.2 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.1 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.27.4 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.17.4 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.15.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/bedrockruntime v1.7.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.11.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.11.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.20.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.23.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.28.1 // indirect
	github.com/aws/constructs-go/constructs/v10 v10.3.0 // indirect
	github.com/aws/jsii-runtime-go v1.101.0 // indirect
	github.com/aws/smithy-go v1.20.1 // indirect
	github.com/cdklabs/awscdk-asset-awscli-go/awscliv1/v2 v2.2.202 // indirect
	github.com/cdklabs/awscdk-asset-kubectl-go/kubectlv20/v2 v2.1.2 // indirect
	github.com/cdklabs/awscdk-asset-node-proxy-agent-go/nodeproxyagentv6/v2 v2.0.3 // indirect
	github.com/fatih/color v1.17.0 // indirect
	github.com/fogfish/golem/hseq v1.3.0 // indirect
	github.com/fogfish/golem/optics v0.14.0 // indirect
	github.com/fogfish/gurl/v2 v2.10.0 // indirect
	github.com/fogfish/opts v0.0.5 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/jdxcode/netrc v1.0.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/yuin/goldmark v1.4.13 // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/mod v0.18.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/tools v0.22.0 // indirect
)
