package LAMBDA

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/aws"
	"log"

	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/spf13/cobra"
)

var AwsxLambdaConfigCmd = &cobra.Command{
	Use:   "getLambdaConfig",
	Short: "getLambdaConfig command gets lambda configuration",
	Long:  `getLambdaConfig command gets lambda configuration`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getLambdaConfig command")
		var authFlag, clientAuth, err = authenticate.AuthenticateCommand(cmd)
		if err != nil {
			log.Printf("error during authentication: %v\n", err)
			err := cmd.Help()
			if err != nil {
				return
			}
			return
		}
		if authFlag {
			functionName, _ := cmd.Flags().GetString("functionName")
			if functionName == "" {
				log.Printf("lambda function name missing")
				err := cmd.Help()
				if err != nil {
					return
				}
				return
			}

			instances, err := GetLambdaFunctionByFunctionName(functionName, clientAuth, nil)
			if err != nil {
				log.Println("error getting getLambdaConfig by function name: ", err)
				return
			}
			fmt.Println(instances)

		}
	},
}

func GetLambdaFunctionByFunctionName(functionName string, clientAuth *model.Auth, client *lambda.Lambda) (*lambda.GetFunctionOutput, error) {
	log.Println("getting aws lambda function of given function name: ", functionName)
	if functionName == "" {
		log.Println("function name missing")
		return nil, fmt.Errorf("function name missing")
	}
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.LAMBDA_CLIENT).(*lambda.Lambda)
	}
	input := &lambda.GetFunctionInput{
		FunctionName: aws.String(functionName),
	}
	lambdaData, err := client.GetFunction(input)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	return lambdaData, nil
}

func init() {
	AwsxLambdaConfigCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxLambdaConfigCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxLambdaConfigCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxLambdaConfigCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxLambdaConfigCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxLambdaConfigCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxLambdaConfigCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxLambdaConfigCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxLambdaConfigCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxLambdaConfigCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxLambdaConfigCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxLambdaConfigCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxLambdaConfigCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxLambdaConfigCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxLambdaConfigCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxLambdaConfigCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxLambdaConfigCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxLambdaConfigCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxLambdaConfigCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxLambdaConfigCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxLambdaConfigCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxLambdaConfigCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxLambdaConfigCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxLambdaConfigCmd.PersistentFlags().String("functionName", "", "lambda function name")
	AwsxLambdaConfigCmd.PersistentFlags().String("bucketName", "", "s3 bucket name")
	AwsxLambdaConfigCmd.PersistentFlags().String("arn", "", "arn")
	AwsxLambdaConfigCmd.PersistentFlags().String("query", "", "query")
	AwsxLambdaConfigCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxLambdaConfigCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxLambdaConfigCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
