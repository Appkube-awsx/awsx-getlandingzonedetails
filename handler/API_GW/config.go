package API_GW

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/spf13/cobra"
	"log"
)

var AwsxApiGwConfigCmd = &cobra.Command{
	Use:   "getApiGwConfig",
	Short: "getApiGwConfig command gets apigw configuration",
	Long:  `getApiGwConfig command gets apigw configuration`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getApiGwRestApiConfig command")
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
			apiKey, _ := cmd.Flags().GetString("apiKey")
			if apiKey == "" {
				log.Printf("api key missing")
				err := cmd.Help()
				if err != nil {
					return
				}
				return
			}

			instances, err := GetApiGwById(apiKey, clientAuth, nil)
			if err != nil {
				log.Println("error getting getApiGwConfig by api key: ", err)
				return
			}
			fmt.Println(instances)
		}
	},
}

func GetApiGwById(apiKey string, clientAuth *model.Auth, client *apigateway.APIGateway) (*apigateway.RestApi, error) {
	log.Println("getting aws apigw. api key: ", apiKey)
	if apiKey == "" {
		log.Println("api key missing")
		return nil, fmt.Errorf("api key missing")
	}
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.APIGATEWAY_CLIENT).(*apigateway.APIGateway)
	}

	apiGwRequest := &apigateway.GetRestApiInput{
		RestApiId: &apiKey,
	}
	apiGwResponse, err := client.GetRestApi(apiGwRequest)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	return apiGwResponse, nil
}

func init() {
	AwsxApiGwConfigCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxApiGwConfigCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxApiGwConfigCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxApiGwConfigCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxApiGwConfigCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxApiGwConfigCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxApiGwConfigCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxApiGwConfigCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxApiGwConfigCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxApiGwConfigCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxApiGwConfigCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxApiGwConfigCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxApiGwConfigCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxApiGwConfigCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxApiGwConfigCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxApiGwConfigCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxApiGwConfigCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxApiGwConfigCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxApiGwConfigCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxApiGwConfigCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxApiGwConfigCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxApiGwConfigCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxApiGwConfigCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxApiGwConfigCmd.PersistentFlags().String("functionName", "", "lambda function name")
	AwsxApiGwConfigCmd.PersistentFlags().String("bucketName", "", "s3 bucket name")
	AwsxApiGwConfigCmd.PersistentFlags().String("arn", "", "arn")
	AwsxApiGwConfigCmd.PersistentFlags().String("query", "", "query")
	AwsxApiGwConfigCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxApiGwConfigCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxApiGwConfigCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
