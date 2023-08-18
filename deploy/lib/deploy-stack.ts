import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as iam from 'aws-cdk-lib/aws-iam';
import * as go from '@aws-cdk/aws-lambda-go-alpha';
import * as lambda from 'aws-cdk-lib/aws-lambda';
import { CfnOutput, Duration, Tags } from 'aws-cdk-lib';
// import * as sqs from 'aws-cdk-lib/aws-sqs';

export class DeployStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    // The code that defines your stack goes here

    // example resource
    // const queue = new sqs.Queue(this, 'DeployQueue', {
    //   visibilityTimeout: cdk.Duration.seconds(300)
    // });
    const uploadServiceLambdaRole = new iam.Role(this, "execution-role", {
      assumedBy: new iam.ServicePrincipal("lambda.amazonaws.com"),
      managedPolicies: [
        iam.ManagedPolicy.fromAwsManagedPolicyName("service-role/AWSLambdaBasicExecutionRole")
      ],
    });
    
    const s3BucketAccess = new iam.PolicyStatement({
      effect: iam.Effect.ALLOW,
      resources: ["arn:aws:s3:::x-stream-videos"],
      actions: ["s3:*"]
    });

    uploadServiceLambdaRole.addToPolicy(s3BucketAccess);

    const uploadServiceLambda = new go.GoFunction(this, 'UploadServiceLambda', {
      entry: '../cmd/upload-service/main.go',
      memorySize: 256,
      tracing: lambda.Tracing.ACTIVE,
      architecture: lambda.Architecture.ARM_64,
      bundling: {
        goBuildFlags: ['-ldflags "-s -w"'],
      },
      role: uploadServiceLambdaRole
    });

    const lambdaCfn = uploadServiceLambda.node.defaultChild as lambda.CfnFunction;
    lambdaCfn.overrideLogicalId('UploadServiceLambda');

    new CfnOutput(this, 'UploadServiceLambdaArn', {
      value: uploadServiceLambda.functionArn,
      exportName: 'UploadServiceLambdaArn',
    });
  }
}
