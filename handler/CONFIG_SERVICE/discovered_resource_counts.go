package CONFIG_SERVICE

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/service/configservice"
	"github.com/spf13/cobra"
	"log"
)

var AwsxDiscoveredResourceCountsCmd = &cobra.Command{
	Use:   "getDiscoveredResourceCounts",
	Short: "getDiscoveredResourceCounts command gets resource counts",
	Long:  `getDiscoveredResourceCounts command gets resource counts`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing command getDiscoveredResourceCounts")
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
			resp, err := DiscoveredResourceCounts(clientAuth, nil)
			if err != nil {
				log.Println("error getting getDiscoveredResourceCounts: ", err)
				return
			}
			fmt.Println(resp)
		} else {
			cmd.Help()
			return
		}
	},
}

func DiscoveredResourceCounts(clientAuth *model.Auth, client *configservice.ConfigService) (*configservice.GetDiscoveredResourceCountsOutput, error) {
	log.Println("Getting aws discovered resource counts")
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.CONFIG_SERVICE_CLIENT).(*configservice.ConfigService)
	}

	configResourceRequest := &configservice.GetDiscoveredResourceCountsInput{}
	configResourceResponse, err := client.GetDiscoveredResourceCounts(configResourceRequest)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	log.Println(configResourceResponse)
	return configResourceResponse, nil
}

func init() {
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("functionName", "", "lambda function name")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("bucketName", "", "s3 bucket name")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("arn", "", "arn")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("query", "", "query")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxDiscoveredResourceCountsCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
