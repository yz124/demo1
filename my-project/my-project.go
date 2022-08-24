package main

import (
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapprunner"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/jsii-runtime-go"
	"os"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	// "github.com/aws/jsii-runtime-go"
)

var account = os.Getenv("CDK_DEFAULT_ACCOUNT")                 // account:number
var region = os.Getenv("CDK_DEFAULT_REGION")                   // region
var accessRoleArn = os.Getenv("CDK_DEFAULT_ACCESSROLEARN")     // access ECR role
var imageIdentifier = os.Getenv("CDK_DEFAULT_IMAGEIDENTIFIER") // ERC url
var serviceName = os.Getenv("CDK_DEFAULT_SERVICENAME")         // service name

type MyProjectStackProps struct {
	awscdk.StackProps
}

func NewMyProjectStack(scope constructs.Construct, id string, props *MyProjectStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	// example resource
	// queue := awssqs.NewQueue(stack, jsii.String("MyProjectQueue"), &awssqs.QueueProps{
	// 	VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	// })
	// create S3 storage
	var bucket = awss3.NewBucket(stack, jsii.String("MyBucket"), &awss3.BucketProps{
		Versioned: jsii.Bool(true),
	})
	fmt.Println("NewBucket BucketWebsiteUrl: ", *bucket.BucketWebsiteUrl())
	// deploy apprunner service
	cnfService := awsapprunner.NewCfnService(stack, jsii.String("MyCfnService"), &awsapprunner.CfnServiceProps{
		SourceConfiguration: awsapprunner.CfnService_SourceConfigurationProperty{
			AuthenticationConfiguration: awsapprunner.CfnService_AuthenticationConfigurationProperty{
				AccessRoleArn: jsii.String(accessRoleArn),
				//ConnectionArn: jsii.String(connectionArn),
			},
			AutoDeploymentsEnabled: jsii.Bool(false),
			CodeRepository:         nil,
			ImageRepository: awsapprunner.CfnService_ImageRepositoryProperty{
				ImageIdentifier:     jsii.String(imageIdentifier),
				ImageRepositoryType: jsii.String("ECR"),
				ImageConfiguration: awsapprunner.CfnService_ImageConfigurationProperty{
					Port:                        jsii.String("80"),
					RuntimeEnvironmentVariables: nil,
					StartCommand:                jsii.String("/start.sh"),
				},
			},
		},
		AutoScalingConfigurationArn: nil,
		EncryptionConfiguration:     nil,
		HealthCheckConfiguration:    nil,
		InstanceConfiguration: awsapprunner.CfnService_InstanceConfigurationProperty{
			Cpu: jsii.String("1 vCPU"),
			//InstanceRoleArn: jsii.String(instanceRoleArn),
			Memory: jsii.String("2 GB"),
		},
		NetworkConfiguration:       nil,
		ObservabilityConfiguration: nil,
		ServiceName:                jsii.String("yidemo2"),
		Tags:                       nil,
	})
	fmt.Println("NewCfnService ServiceName=", *cnfService.ServiceName())

	return stack
}

func main() {
	fmt.Println("NewMyProjectStack...")
	app := awscdk.NewApp(nil)

	NewMyProjectStack(app, "MyProjectStack", &MyProjectStackProps{
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
	//return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	return &awscdk.Environment{
		Account: jsii.String(account),
		Region:  jsii.String(region),
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
