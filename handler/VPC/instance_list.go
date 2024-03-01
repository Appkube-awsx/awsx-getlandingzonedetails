package VPC

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
	"log"
)

var AwsxVpcListCmd = &cobra.Command{
	Use:   "getVpcList",
	Short: "getVpcList command gets list of vpc instances of an aws account",
	Long:  `getVpcList command gets list of vpc instances of an aws account`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getVpcList command")
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
			resp, err := ListVpcInstances(clientAuth, nil)
			if err != nil {
				log.Println("error getting getVpcList: ", err)
				return
			}
			fmt.Println(resp)
		}
	},
}

func ListVpcInstances(clientAuth *model.Auth, client *ec2.EC2) (*ec2.DescribeVpcsOutput, error) {
	log.Println("getting vpc list")
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.EC2_CLIENT).(*ec2.EC2)
	}
	result, err := client.DescribeVpcs(nil)
	if err != nil {
		log.Println("error getting vpc list", err)
		return nil, err
	}
	return result, err
}

func init() {
	AwsxVpcListCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxVpcListCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxVpcListCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxVpcListCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxVpcListCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxVpcListCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxVpcListCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxVpcListCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxVpcListCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxVpcListCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxVpcListCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxVpcListCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxVpcListCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxVpcListCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxVpcListCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxVpcListCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxVpcListCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxVpcListCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxVpcListCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxVpcListCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxVpcListCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxVpcListCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxVpcListCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxVpcListCmd.PersistentFlags().String("functionName", "", "lambda function name")
	AwsxVpcListCmd.PersistentFlags().String("bucketName", "", "s3 bucket name")
	AwsxVpcListCmd.PersistentFlags().String("arn", "", "arn")
	AwsxVpcListCmd.PersistentFlags().String("query", "", "query")
	AwsxVpcListCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxVpcListCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxVpcListCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
