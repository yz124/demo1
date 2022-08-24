package main

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	assertions "github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

//example tests. To run these tests, uncomment this file along with the
//example resource in my-project_test.go
func TestMyProjectStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := NewMyProjectStack(app, "MyProjectStack", nil)

	// THEN
	template := assertions.Template_FromStack(stack)
	template.HasResourceProperties(jsii.String("AWS::S3::Bucket"), map[string]interface{}{
	})
	template.HasResourceProperties(jsii.String("AWS::AppRunner::Service"), map[string]interface{}{
		"ServiceName": "yidemo2",
	})
	//template.HasResourceProperties(jsii.String("AWS::SQS::Queue"), map[string]interface{}{
	//	"VisibilityTimeout": 300,
	//})
}
