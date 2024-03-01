package DYNAMODB

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/spf13/cobra"
)

var AwsxDynamodDbConfigCmd = &cobra.Command{
	Use:   "getDynamoDbConfig",
	Short: "getDynamoDbConfig command gets dynamo-db configuration",
	Long:  `getDynamoDbConfig command gets dynamo-db configuration`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getDynamoDbConfig command")
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
			tableName, _ := cmd.Flags().GetString("tableName")
			if tableName == "" {
				log.Printf("dynamo-db table name missing")
				err := cmd.Help()
				if err != nil {
					return
				}
				return
			}
			instances, err := GetDynamoDbInstanceByTableName(tableName, clientAuth, nil)
			if err != nil {
				log.Println("error getting getDynamoDbConfig by table name: ", err)
				return
			}
			fmt.Println(instances)
		}
	},
}

func GetDynamoDbInstanceByTableName(tableName string, clientAuth *model.Auth, client *dynamodb.DynamoDB) (*dynamodb.DescribeTableOutput, error) {
	log.Println("getting aws dynamo-db instance of given table name: ", tableName)
	if tableName == "" {
		log.Println("dynamo-db table name missing")
		return nil, fmt.Errorf("dynamo-db table name missing")
	}
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.DYNAMODB_CLIENT).(*dynamodb.DynamoDB)
	}

	input := &dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	}
	tableData, err := client.DescribeTable(input)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	return tableData, nil
}

func init() {
	AwsxDynamodDbConfigCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("functionName", "", "lambda function name")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("bucketName", "", "s3 bucket name")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("arn", "", "arn")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("query", "", "query")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxDynamodDbConfigCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
