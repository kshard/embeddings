//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/embeddings
//

package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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

func (c *FoundationModel) GrantAccessIn(grantee awsiam.IGrantable, region *string) {
	arn := awscdk.Stack_Of(c.Construct).FormatArn(
		&awscdk.ArnComponents{
			ArnFormat:    awscdk.ArnFormat_SLASH_RESOURCE_NAME,
			Service:      jsii.String("bedrock"),
			Account:      jsii.String(""),
			Region:       region,
			Resource:     jsii.String("foundation-model"),
			ResourceName: c.embeddings.ModelId(),
		},
	)

	awsiam.Grant_AddToPrincipal(
		&awsiam.GrantOnPrincipalOptions{
			Grantee:      grantee,
			Actions:      jsii.Strings("bedrock:InvokeModel"),
			ResourceArns: jsii.Strings(*arn),
		},
	)
}

func NewTitanTextEmbeddingsV1(scope constructs.Construct) *FoundationModel {
	return NewFoundationModel(scope, jsii.String("TitanTextEmbeddingsV1"),
		awsbedrock.FoundationModelIdentifier_AMAZON_TITAN_EMBEDDINGS_G1_TEXT_V1(),
	)
}

func NewTitanTextEmbeddingsV2(scope constructs.Construct) *FoundationModel {
	return NewFoundationModel(scope, jsii.String("TitanTextEmbeddingsV2"),
		awsbedrock.FoundationModelIdentifier_AMAZON_TITAN_EMBED_TEXT_V2_0(),
	)
}
