# Welcome to your CDK Go project!

This is a blank project for CDK development with Go.

The `cdk.json` file tells the CDK toolkit how to execute your app.

## environment variable
export CDK_DEFAULT_ACCOUNT=12345678

export CDK_DEFAULT_REGION=us-east-1

export CDK_DEFAULT_ACCESSROLEARN=arn:aws:iam::12345678:role/service-role/AppRunnerECRAccessRole

export CDK_DEFAULT_IMAGEIDENTIFIER=12345678.dkr.ecr.us-east-1.amazonaws.com/yidemo:latest

export CDK_DEFAULT_SERVICENAME=yidemo

## Useful commands

 * `cdk synth`       emits the synthesized CloudFormation template
 * `cdk diff`        compare deployed stack with current state
 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `go test`         run unit tests
