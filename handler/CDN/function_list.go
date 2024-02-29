package CDN

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/spf13/cobra"
	"log"
)

var AwsxCdnFunctionListCmd = &cobra.Command{
	Use:   "getCdnFunctionList",
	Short: "getCdnFunctionList command gets list of cdn functions of an aws account",
	Long:  `getCdnFunctionList command gets list of cdn functions of an aws account`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getCdnFunctionList command")
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
			resp, err := ListCdnFunctionInstances(clientAuth, nil)
			if err != nil {
				log.Println("error getting getCdnFunctionList: ", err)
				return
			}
			fmt.Println(resp)
		}
	},
}

func ListCdnFunctionInstances(clientAuth *model.Auth, client *cloudfront.CloudFront) (*cloudfront.ListFunctionsOutput, error) {
	log.Println("getting aws (cdn) cloudfront functions list")
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.CLOUD_FRONT_CLIENT).(*cloudfront.CloudFront)
	}
	input := &cloudfront.ListFunctionsInput{}
	response, err := client.ListFunctions(input)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	return response, err
}

func init() {
	AwsxCdnFunctionListCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxCdnFunctionListCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxCdnFunctionListCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxCdnFunctionListCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxCdnFunctionListCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxCdnFunctionListCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxCdnFunctionListCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxCdnFunctionListCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxCdnFunctionListCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxCdnFunctionListCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxCdnFunctionListCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxCdnFunctionListCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxCdnFunctionListCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxCdnFunctionListCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxCdnFunctionListCmd.PersistentFlags().String("ServiceName", "", "Service Name")
	AwsxCdnFunctionListCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxCdnFunctionListCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxCdnFunctionListCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxCdnFunctionListCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxCdnFunctionListCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxCdnFunctionListCmd.PersistentFlags().String("query", "", "query")
	AwsxCdnFunctionListCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxCdnFunctionListCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxCdnFunctionListCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
