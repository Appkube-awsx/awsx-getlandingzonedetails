package VPC

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"

	"github.com/spf13/cobra"
)

var AwsxVpcConfigCmd = &cobra.Command{
	Use:   "getVpcConfig",
	Short: "getVpcConfig command gets vpc configuration",
	Long:  `getVpcConfig command gets vpc configuration`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getVpcConfig command")
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
				log.Printf("vpc id missing")
				err := cmd.Help()
				if err != nil {
					return
				}
				return
			}

			instances, err := GetVpcInstanceById(instanceId, clientAuth, nil)
			if err != nil {
				log.Println("error getting getVpcConfig by bucket name: ", err)
				return
			}
			fmt.Println(instances)

		}
	},
}

func GetVpcInstanceById(instanceId string, clientAuth *model.Auth, client *ec2.EC2) (*ec2.DescribeVpcsOutput, error) {
	log.Println("getting aws vpc instance of given instanceId: ", instanceId)
	if instanceId == "" {
		log.Println("vpd id missing")
		return nil, fmt.Errorf("vpd id missing")
	}
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.EC2_CLIENT).(*ec2.EC2)
	}
	input := &ec2.DescribeVpcsInput{
		VpcIds: []*string{aws.String(instanceId)},
	}
	result, err := client.DescribeVpcs(input)
	if err != nil {
		log.Println("error in getting vpc detail ", err)
		return nil, err
	}
	return result, nil
}

func init() {
	AwsxVpcConfigCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxVpcConfigCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxVpcConfigCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxVpcConfigCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxVpcConfigCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxVpcConfigCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxVpcConfigCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxVpcConfigCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxVpcConfigCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxVpcConfigCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxVpcConfigCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxVpcConfigCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxVpcConfigCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxVpcConfigCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxVpcConfigCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxVpcConfigCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxVpcConfigCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxVpcConfigCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxVpcConfigCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxVpcConfigCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxVpcConfigCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxVpcConfigCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxVpcConfigCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxVpcConfigCmd.PersistentFlags().String("functionName", "", "lambda function name")
	AwsxVpcConfigCmd.PersistentFlags().String("bucketName", "", "s3 bucket name")
	AwsxVpcConfigCmd.PersistentFlags().String("arn", "", "arn")
	AwsxVpcConfigCmd.PersistentFlags().String("query", "", "query")
	AwsxVpcConfigCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxVpcConfigCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxVpcConfigCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
