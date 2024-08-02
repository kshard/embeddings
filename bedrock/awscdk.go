package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

// Foundation Model L3 construct simplify grant access
type FoundationModel struct {
	constructs.Construct
	embeddings awsbedrock.FoundationModel
}

func NewFoundationModel(scope constructs.Construct, id *string, foundationModelId awsbedrock.FoundationModelIdentifier) *FoundationModel {
	c := &FoundationModel{Construct: constructs.NewConstruct(scope, id)}
	c.embeddings = awsbedrock.FoundationModel_FromFoundationModelId(
		c.Construct,
		jsii.String("Embeddings"),
		foundationModelId,
	)

	return c
}

func (c *FoundationModel) GrantAccess(grantee awsiam.IGrantable) {
	awsiam.Grant_AddToPrincipal(
		&awsiam.GrantOnPrincipalOptions{
			Grantee:      grantee,
			Actions:      jsii.Strings("bedrock:InvokeModel"),
			ResourceArns: jsii.Strings(*c.embeddings.ModelArn()),
		},
	)
}

func NewTitanTextEmbeddingsV1(scope constructs.Construct, id *string) *FoundationModel {
	return NewFoundationModel(scope, id, awsbedrock.FoundationModelIdentifier_AMAZON_TITAN_EMBEDDINGS_G1_TEXT_V1())
}

func NewTitanTextEmbeddingsV2(scope constructs.Construct, id *string) *FoundationModel {
	return NewFoundationModel(scope, id, awsbedrock.FoundationModelIdentifier_AMAZON_TITAN_EMBED_TEXT_V2_0())
}
