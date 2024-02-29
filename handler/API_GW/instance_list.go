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

var AwsxApiGwListCmd = &cobra.Command{
	Use:   "getApiGwList",
	Short: "getApiGwList command gets list of apigw rest api instances of an aws account",
	Long:  `getApiGwList command gets list of apigw rest api instances of an aws account`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getApiGwList command")
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
			instances, err := ListApiGwInstances(clientAuth, nil)
			if err != nil {
				log.Println("error getting getApiGwList: ", err)
				return
			}
			fmt.Println(instances)
		}

	},
}

func ListApiGwInstances(clientAuth *model.Auth, client *apigateway.APIGateway) (*apigateway.GetRestApisOutput, error) {
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.APIGATEWAY_CLIENT).(*apigateway.APIGateway)
	}

	apiGwResponse, err := client.GetRestApis(nil)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	return apiGwResponse, nil
}

func init() {
	AwsxApiGwListCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxApiGwListCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxApiGwListCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxApiGwListCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxApiGwListCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxApiGwListCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxApiGwListCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxApiGwListCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxApiGwListCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxApiGwListCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxApiGwListCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxApiGwListCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxApiGwListCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxApiGwListCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxApiGwListCmd.PersistentFlags().String("ServiceName", "", "Service Name")
	AwsxApiGwListCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxApiGwListCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxApiGwListCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxApiGwListCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxApiGwListCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxApiGwListCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxApiGwListCmd.PersistentFlags().String("query", "", "query")
	AwsxApiGwListCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxApiGwListCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxApiGwListCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
