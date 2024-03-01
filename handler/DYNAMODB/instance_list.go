package DYNAMODB

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/spf13/cobra"
	"log"
)

var AwsxDynamoDbListCmd = &cobra.Command{
	Use:   "getDynamoDbList",
	Short: "getDynamoDbList command gets list of dynamo-db instances of an aws account",
	Long:  `getDynamoDbList command gets list of dynamo-db instances of an aws account`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getDynamoDbList command")
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
			instances, err := ListDynamoDbInstance(clientAuth, nil)
			if err != nil {
				log.Println("error getting getDynamoDbList: ", err)
				return
			}
			fmt.Println(instances)
		}

	},
}

func ListDynamoDbInstanceNames(clientAuth *model.Auth, client *dynamodb.DynamoDB) (*dynamodb.ListTablesOutput, error) {
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.DYNAMODB_CLIENT).(*dynamodb.DynamoDB)
	}

	input := &dynamodb.ListTablesInput{}
	tableList, err := client.ListTables(input)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	return tableList, nil
}

type DynamoDb struct {
	Table interface{} `json:"table"`
	Tags  interface{} `json:"tags"`
}

func ListDynamoDbInstance(clientAuth *model.Auth, client *dynamodb.DynamoDB) ([]DynamoDb, error) {
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.DYNAMODB_CLIENT).(*dynamodb.DynamoDB)
	}
	log.Println("Getting detail of each dynamodb table")
	tableNamesOutput, err := ListDynamoDbInstanceNames(clientAuth, client)
	if err != nil {
		log.Fatalln("error in getting dynamodb table list", err)
		return nil, err
	}
	allTableDetais := []DynamoDb{}

	for _, table := range tableNamesOutput.TableNames {
		tableDetail, err := GetDynamoDbInstanceByTableName(*table, clientAuth, client)
		tagInput := &dynamodb.ListTagsOfResourceInput{
			ResourceArn: tableDetail.Table.TableArn,
		}
		tagOutput, err := client.ListTagsOfResource(tagInput)
		dynamodbObj := DynamoDb{
			Table: tableDetail,
			Tags:  tagOutput,
		}
		if err != nil {
			log.Println("Error: in getting dynamodb table detail", err)
			continue
		} else {
			allTableDetais = append(allTableDetais, dynamodbObj)
		}

	}
	return allTableDetais, err
}

func init() {
	AwsxDynamoDbListCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxDynamoDbListCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxDynamoDbListCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxDynamoDbListCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxDynamoDbListCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxDynamoDbListCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxDynamoDbListCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxDynamoDbListCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxDynamoDbListCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxDynamoDbListCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxDynamoDbListCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxDynamoDbListCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxDynamoDbListCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxDynamoDbListCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxDynamoDbListCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxDynamoDbListCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxDynamoDbListCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxDynamoDbListCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxDynamoDbListCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxDynamoDbListCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxDynamoDbListCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxDynamoDbListCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxDynamoDbListCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxDynamoDbListCmd.PersistentFlags().String("functionName", "", "lambda function name")
	AwsxDynamoDbListCmd.PersistentFlags().String("bucketName", "", "s3 bucket name")
	AwsxDynamoDbListCmd.PersistentFlags().String("arn", "", "arn")
	AwsxDynamoDbListCmd.PersistentFlags().String("query", "", "query")
	AwsxDynamoDbListCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxDynamoDbListCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxDynamoDbListCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
