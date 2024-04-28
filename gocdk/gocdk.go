package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

const (
	USER_TABLE_NAME = "userTable"
)

type GocdkStackProps struct {
	awscdk.StackProps
}

func NewGocdkStack(scope constructs.Construct, id string, props *GocdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}

	stack := awscdk.NewStack(scope, &id, &sprops)

	// create the dynamodb table
	table := awsdynamodb.NewTable(stack, jsii.String(USER_TABLE_NAME), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("username"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		TableName:     jsii.String(USER_TABLE_NAME),
		BillingMode:   awsdynamodb.BillingMode_PAY_PER_REQUEST,
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	// Grant the Lambda function permissions to read and write to the DynamoDB table
	myLambda := awslambda.NewFunction(stack, jsii.String("myFirstGoLambdaFunction"), &awslambda.FunctionProps{
		Timeout:    awscdk.Duration_Seconds(jsii.Number(3)),
		Runtime:    awslambda.Runtime_PROVIDED_AL2023(),
		Handler:    jsii.String("main"),
		MemorySize: jsii.Number(512),
		Environment: &map[string]*string{
			"FOO": jsii.String("bar"),
		},
		Code: awslambda.Code_FromAsset(jsii.String("lambdas/lambda/function.zip"), nil),
	})

	table.GrantReadWriteData(myLambda)

	// create the api gateway with CloudWatch logging enabled
	api := awsapigateway.NewRestApi(stack, jsii.String("myFirstGoApi"), &awsapigateway.RestApiProps{
		EndpointTypes: &[]awsapigateway.EndpointType{awsapigateway.EndpointType_REGIONAL},
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowHeaders: jsii.Strings("Content-Type", "Authorization"),
			AllowMethods: jsii.Strings("GET", "POST", "PUT", "DELETE", "OPTIONS"),
			AllowOrigins: awsapigateway.Cors_ALL_ORIGINS(),
		},
		DeployOptions: &awsapigateway.StageOptions{
			LoggingLevel: awsapigateway.MethodLoggingLevel_INFO,
		},
		CloudWatchRole: jsii.Bool(true),
	})

	// Define the Lambda integration
	lambdaIntegration := awsapigateway.NewLambdaIntegration(myLambda, nil)
	registerResource := api.Root().AddResource(jsii.String("register"), nil)
	registerResource.AddMethod(jsii.String("POST"), lambdaIntegration, nil)
	registerResource.DefaultCorsPreflightOptions().AllowOrigins = awsapigateway.Cors_ALL_ORIGINS()
	registerResource.DefaultCorsPreflightOptions().AllowMethods = jsii.Strings("OPTIONS", "POST", "GET", "PUT", "DELETE", "PATCH", "HEAD")
	registerResource.DefaultCorsPreflightOptions().AllowHeaders = jsii.Strings("Content-Type", "Authorization", "X-API-Key", "Content-Length", "X-Amz-Date", "X-Amz-Security-Token", "X-Amz-User-Agent")

	// Login resource
	loginResource := api.Root().AddResource(jsii.String("login"), nil)
	loginResource.AddMethod(jsii.String("POST"), lambdaIntegration, nil)
	loginResource.DefaultCorsPreflightOptions().AllowOrigins = awsapigateway.Cors_ALL_ORIGINS()
	loginResource.DefaultCorsPreflightOptions().AllowMethods = jsii.Strings("OPTIONS", "POST", "GET", "PUT", "DELETE", "PATCH", "HEAD")
	loginResource.DefaultCorsPreflightOptions().AllowHeaders = jsii.Strings("Content-Type", "Authorization", "X-API-Key", "Content-Length", "X-Amz-Date", "X-Amz-Security-Token", "X-Amz-User-Agent")

	// User resource
	userResource := api.Root().AddResource(jsii.String("user"), nil)
	userResource.AddMethod(jsii.String("GET"), lambdaIntegration, nil)
	userResource.DefaultCorsPreflightOptions().AllowOrigins = awsapigateway.Cors_ALL_ORIGINS()
	userResource.DefaultCorsPreflightOptions().AllowMethods = jsii.Strings("OPTIONS", "POST", "GET", "PUT", "DELETE", "PATCH", "HEAD")
	userResource.DefaultCorsPreflightOptions().AllowHeaders = jsii.Strings("Content-Type", "Authorization", "X-API-Key", "Content-Length", "X-Amz-Date", "X-Amz-Security-Token", "X-Amz-User-Agent")

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewGocdkStack(app, "GocdkStack", &GocdkStackProps{
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
	//---------------------------------------------------------------------------
	// return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	return &awscdk.Environment{
		Account: jsii.String("444212982395"),
		Region:  jsii.String("eu-west-1"),
	}

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}

// &awss3assets.AssetOptions{
// 	// Bundling: &awscdk.BundlingOptions{
// 	// 	Image:   awscdk.DockerImage_FromRegistry(jsii.String("golang:1.22-alpine")),
// 	// 	Command: jsii.Strings("go", "build", "-o", "/asset-output/main", "-ldflags", "-s -w", "."),
// 	// },
// }),
