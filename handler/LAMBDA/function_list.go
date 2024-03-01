package LAMBDA

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/spf13/cobra"
	"log"
)

var AwsxLambdaListCmd = &cobra.Command{
	Use:   "getLambdaList",
	Short: "getLambdaList command gets list of lambda functions of an aws account",
	Long:  `getLambdaList command gets list of lambda functions of an aws account`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getLambdaList command")
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
			resp, err := ListLambdaInstances(clientAuth, nil)
			if err != nil {
				log.Println("error getting getLambdaList: ", err)
				return
			}
			fmt.Println(resp)
		}
	},
}

func ListLambdaInstances(clientAuth *model.Auth, client *lambda.Lambda) (*lambda.ListFunctionsOutput, error) {
	log.Println("getting lambda list")
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.LAMBDA_CLIENT).(*lambda.Lambda)
	}
	input := &lambda.ListFunctionsInput{}
	functionList, err := client.ListFunctions(input)
	if err != nil {
		log.Println("error getting lambda list", err)
		return nil, err
	}
	return functionList, nil
}

func init() {
	AwsxLambdaListCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxLambdaListCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxLambdaListCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxLambdaListCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxLambdaListCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxLambdaListCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxLambdaListCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxLambdaListCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxLambdaListCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxLambdaListCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxLambdaListCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxLambdaListCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxLambdaListCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxLambdaListCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxLambdaListCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxLambdaListCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxLambdaListCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxLambdaListCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxLambdaListCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxLambdaListCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxLambdaListCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxLambdaListCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxLambdaListCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxLambdaListCmd.PersistentFlags().String("functionName", "", "lambda function name")
	AwsxLambdaListCmd.PersistentFlags().String("bucketName", "", "s3 bucket name")
	AwsxLambdaListCmd.PersistentFlags().String("arn", "", "arn")
	AwsxLambdaListCmd.PersistentFlags().String("query", "", "query")
	AwsxLambdaListCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxLambdaListCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxLambdaListCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
