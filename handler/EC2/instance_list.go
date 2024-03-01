package EC2

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
	"log"
)

var AwsxEc2ListCmd = &cobra.Command{
	Use:   "getEc2List",
	Short: "getEc2List command gets list of ec2 instances of an aws account",
	Long:  `getEc2List command gets list of ec2 instances of an aws account`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getEc2List command")
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
			resp, err := ListEc2Instances(clientAuth, nil)
			if err != nil {
				log.Println("error getting getEc2List: ", err)
				return
			}
			fmt.Println(resp)
		}
	},
}

func ListEc2Instances(clientAuth *model.Auth, client *ec2.EC2) (*ec2.DescribeInstancesOutput, error) {
	log.Println("getting aws ec2 instance list")
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.EC2_CLIENT).(*ec2.EC2)
	}

	ec2Request := &ec2.DescribeInstancesInput{}
	ec2Response, err := client.DescribeInstances(ec2Request)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	for _, reservation := range ec2Response.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Println("ec2 instance id: ", *instance.InstanceId)
		}
	}
	return ec2Response, err
}

func init() {
	AwsxEc2ListCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxEc2ListCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxEc2ListCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxEc2ListCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxEc2ListCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxEc2ListCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxEc2ListCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxEc2ListCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxEc2ListCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxEc2ListCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxEc2ListCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxEc2ListCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxEc2ListCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxEc2ListCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxEc2ListCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxEc2ListCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxEc2ListCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxEc2ListCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxEc2ListCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxEc2ListCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxEc2ListCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxEc2ListCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxEc2ListCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxEc2ListCmd.PersistentFlags().String("functionName", "", "lambda function name")
	AwsxEc2ListCmd.PersistentFlags().String("bucketName", "", "s3 bucket name")
	AwsxEc2ListCmd.PersistentFlags().String("arn", "", "arn")
	AwsxEc2ListCmd.PersistentFlags().String("query", "", "query")
	AwsxEc2ListCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxEc2ListCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxEc2ListCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
