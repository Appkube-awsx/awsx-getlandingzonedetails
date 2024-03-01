package WAF

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/waf"
	"log"

	"github.com/spf13/cobra"
)

var AwsxWafConfigCmd = &cobra.Command{
	Use:   "getWafConfig",
	Short: "getWafConfig command gets waf configuration",
	Long:  `getWafConfig command gets waf configuration`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getWafConfig command")
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
			instanceId, _ := cmd.Flags().GetString("instanceId")
			if instanceId == "" {
				log.Printf("waf id missing")
				err := cmd.Help()
				if err != nil {
					return
				}
				return
			}

			instances, err := GetWafInstanceById(instanceId, clientAuth, nil)
			if err != nil {
				log.Println("error getting getWafConfig by bucket name: ", err)
				return
			}
			fmt.Println(instances)

		}
	},
}

func GetWafInstanceById(instanceId string, clientAuth *model.Auth, client *waf.WAF) (*waf.GetWebACLOutput, error) {
	log.Println("getting aws waf instance of given instanceId: ", instanceId)
	if instanceId == "" {
		log.Println("waf id missing")
		return nil, fmt.Errorf("waf id missing")
	}
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.WAF_CLIENT).(*waf.WAF)
	}
	input := &waf.GetWebACLInput{
		WebACLId: aws.String(instanceId),
	}

	result, err := client.GetWebACL(input)
	if err != nil {
		log.Println("error in getting waf detail ", err)
		return nil, err
	}
	return result, nil
}

func init() {
	AwsxWafConfigCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxWafConfigCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxWafConfigCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxWafConfigCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxWafConfigCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxWafConfigCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxWafConfigCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxWafConfigCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxWafConfigCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxWafConfigCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxWafConfigCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxWafConfigCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxWafConfigCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxWafConfigCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxWafConfigCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxWafConfigCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxWafConfigCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxWafConfigCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxWafConfigCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxWafConfigCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxWafConfigCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxWafConfigCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxWafConfigCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxWafConfigCmd.PersistentFlags().String("functionName", "", "lambda function name")
	AwsxWafConfigCmd.PersistentFlags().String("bucketName", "", "s3 bucket name")
	AwsxWafConfigCmd.PersistentFlags().String("arn", "", "arn")
	AwsxWafConfigCmd.PersistentFlags().String("query", "", "query")
	AwsxWafConfigCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxWafConfigCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxWafConfigCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
