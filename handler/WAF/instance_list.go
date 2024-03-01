package WAF

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/spf13/cobra"
	"log"
)

var AwsxWafListCmd = &cobra.Command{
	Use:   "getWafList",
	Short: "getWafList command gets list of waf instances of an aws account",
	Long:  `getWafList command gets list of waf instances of an aws account`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getWafList command")
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
			resp, err := ListWafInstances(clientAuth, nil)
			if err != nil {
				log.Println("error getting getWafList: ", err)
				return
			}
			fmt.Println(resp)
		}
	},
}

func ListWafInstances(clientAuth *model.Auth, client *waf.WAF) (*waf.ListWebACLsOutput, error) {
	log.Println("getting waf list")
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.WAF_CLIENT).(*waf.WAF)
	}
	input := &waf.ListWebACLsInput{}
	result, err := client.ListWebACLs(input)
	if err != nil {
		log.Println("error getting waf list", err)
		return nil, err
	}
	return result, err
}

func init() {
	AwsxWafListCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxWafListCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxWafListCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxWafListCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxWafListCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxWafListCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxWafListCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxWafListCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxWafListCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxWafListCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxWafListCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxWafListCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxWafListCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxWafListCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxWafListCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxWafListCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxWafListCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxWafListCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxWafListCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxWafListCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxWafListCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxWafListCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxWafListCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxWafListCmd.PersistentFlags().String("functionName", "", "lambda function name")
	AwsxWafListCmd.PersistentFlags().String("bucketName", "", "s3 bucket name")
	AwsxWafListCmd.PersistentFlags().String("arn", "", "arn")
	AwsxWafListCmd.PersistentFlags().String("query", "", "query")
	AwsxWafListCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxWafListCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxWafListCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
