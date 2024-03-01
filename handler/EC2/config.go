package EC2

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
)

var AwsxEc2ConfigCmd = &cobra.Command{
	Use:   "getEc2Config",
	Short: "getEc2Config command gets ec2 configuration",
	Long:  `getEc2Config command gets ec2 configuration`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getEc2Config command")
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
			tagName, _ := cmd.Flags().GetString("tagName")
			instanceId, _ := cmd.Flags().GetString("instanceId")
			if tagName == "" && instanceId == "" {
				log.Printf("ec2 instance-id or tag missing")
				err := cmd.Help()
				if err != nil {
					return
				}
				return
			}
			if tagName != "" {
				instances, err := GetEc2InstanceByTagName(tagName, clientAuth, nil)
				if err != nil {
					log.Println("error getting getEc2Config by tag name: ", err)
					return
				}
				fmt.Println(instances)
			}
			if instanceId != "" {
				instances, err := GetEc2InstanceById(instanceId, clientAuth, nil)
				if err != nil {
					log.Println("error getting getEc2Config by instance id: ", err)
					return
				}
				fmt.Println(instances)
			}
		}
	},
}

func GetEc2InstanceByTagName(tagName string, clientAuth *model.Auth, ec2Client *ec2.EC2) (*ec2.DescribeInstancesOutput, error) {
	log.Println("getting aws ec2 instance of given tag: ", tagName)
	if tagName == "" {
		log.Println("tag missing")
		return nil, fmt.Errorf("tag missing")
	}
	if ec2Client == nil {
		ec2Client = awsclient.GetClient(*clientAuth, awsclient.EC2_CLIENT).(*ec2.EC2)
	}
	filters := []*ec2.Filter{
		{
			Name:   aws.String("tag:Name"),
			Values: []*string{aws.String(tagName)},
		},
	}
	ec2Request := &ec2.DescribeInstancesInput{
		Filters: filters,
	}
	ec2Response, err := ec2Client.DescribeInstances(ec2Request)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	return ec2Response, nil
}

func GetEc2InstanceById(instanceId string, clientAuth *model.Auth, client *ec2.EC2) (*ec2.DescribeInstancesOutput, error) {
	log.Println("getting aws ec2 instance of given instance id: ", instanceId)
	if instanceId == "" {
		log.Println("instance-id missing")
		return nil, fmt.Errorf("instance-id missing")
	}
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.EC2_CLIENT).(*ec2.EC2)
	}

	ec2Request := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{aws.String(instanceId)},
	}
	ec2Response, err := client.DescribeInstances(ec2Request)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	return ec2Response, nil
}

func init() {
	AwsxEc2ConfigCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxEc2ConfigCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxEc2ConfigCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxEc2ConfigCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxEc2ConfigCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxEc2ConfigCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxEc2ConfigCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxEc2ConfigCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxEc2ConfigCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxEc2ConfigCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxEc2ConfigCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxEc2ConfigCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxEc2ConfigCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxEc2ConfigCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxEc2ConfigCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxEc2ConfigCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxEc2ConfigCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxEc2ConfigCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxEc2ConfigCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxEc2ConfigCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxEc2ConfigCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxEc2ConfigCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxEc2ConfigCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxEc2ConfigCmd.PersistentFlags().String("functionName", "", "lambda function name")
	AwsxEc2ConfigCmd.PersistentFlags().String("bucketName", "", "s3 bucket name")
	AwsxEc2ConfigCmd.PersistentFlags().String("arn", "", "arn")
	AwsxEc2ConfigCmd.PersistentFlags().String("query", "", "query")
	AwsxEc2ConfigCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxEc2ConfigCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxEc2ConfigCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
