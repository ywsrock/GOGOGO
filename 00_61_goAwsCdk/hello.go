package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type HelloStackProps struct {
	awscdk.StackProps
}

func NewCfdStack(scope constructs.Construct, id string, props *HelloStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// Lambda関数
	handler := NewAfLambdaHandler(stack)
	// ApiGateway
	NewAfApiGatewap(stack, handler)

	return stack
}

func NewAfApiGatewap(stack awscdk.Stack, handler awscdklambdagoalpha.GoFunction) {
	api := awsapigateway.NewRestApi(stack, jsii.String("AFDemoCdkApi"), &awsapigateway.RestApiProps{
		RestApiName:    jsii.String("AFDemoAPI"),
		CloudWatchRole: jsii.Bool(false),
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowOrigins: awsapigateway.Cors_ALL_ORIGINS(),
			AllowMethods: awsapigateway.Cors_ALL_METHODS(),
			AllowHeaders: awsapigateway.Cors_DEFAULT_HEADERS(),
		},
	})

	api.Root().AddMethod(
		jsii.String("POST"),
		awsapigateway.NewLambdaIntegration(handler, &awsapigateway.LambdaIntegrationOptions{}),
		api.Root().DefaultMethodOptions(),
	)

	api.Root().AddResource(jsii.String("history"), &awsapigateway.ResourceOptions{}).AddMethod(
		jsii.String("GET"),
		awsapigateway.NewLambdaIntegration(handler, &awsapigateway.LambdaIntegrationOptions{}),
		api.Root().DefaultMethodOptions(),
	)

	api.Root().AddResource(jsii.String("historys"), &awsapigateway.ResourceOptions{}).AddMethod(
		jsii.String("GET"),
		awsapigateway.NewLambdaIntegration(handler, &awsapigateway.LambdaIntegrationOptions{}),
		api.Root().DefaultMethodOptions(),
	)
}

func NewAfLambdaHandler(stack awscdk.Stack) awscdklambdagoalpha.GoFunction {
	handler := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("AFDemoCdkLambda"), &awscdklambdagoalpha.GoFunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Entry:   jsii.String("./lambda"),
		Bundling: &awscdklambdagoalpha.BundlingOptions{
			GoBuildFlags: jsii.Strings(`-ldflags "-s -w"`),
		},
	})
	return handler
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewCfdStack(app, "CfdDemoStack", &HelloStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	// ---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	// ---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	// ---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
