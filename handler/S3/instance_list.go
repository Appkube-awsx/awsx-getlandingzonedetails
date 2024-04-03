package S3

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
	"log"
)

var AwsxS3ListCmd = &cobra.Command{
	Use:   "getS3List",
	Short: "getS3List command gets list of s3 instances of an aws account",
	Long:  `getS3List command gets list of s3 instances of an aws account`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getS3List command")
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
			resp, err := ListS3Instances(clientAuth, nil)
			if err != nil {
				log.Println("error getting getS3List: ", err)
				return
			}
			fmt.Println(resp)
		}
	},
}

func ListS3Instances(clientAuth *model.Auth, client *s3.S3) ([]S3Bucket, error) {
	log.Println("getting s3 list")
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.S3_CLIENT).(*s3.S3)
	}
	request := &s3.ListBucketsInput{}
	response, err := client.ListBuckets(request)
	if err != nil {
		log.Println("error getting s3 list", err)
		return nil, err
	}
	allBuckets := []S3Bucket{}
	for _, bucket := range response.Buckets {
		//s3Bucket, err := GetS3InstanceByBucketName(*bucket.Name, clientAuth, client)
		//if err != nil {
		//	continue
		//}
		s3b := S3Bucket{
			Bucket: bucket,
		}
		allBuckets = append(allBuckets, s3b)
	}
	return allBuckets, err
}

func init() {
	AwsxS3ListCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxS3ListCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxS3ListCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxS3ListCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxS3ListCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxS3ListCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxS3ListCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxS3ListCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxS3ListCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxS3ListCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxS3ListCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxS3ListCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxS3ListCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxS3ListCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxS3ListCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxS3ListCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxS3ListCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxS3ListCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxS3ListCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxS3ListCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxS3ListCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxS3ListCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxS3ListCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxS3ListCmd.PersistentFlags().String("functionName", "", "lambda function name")
	AwsxS3ListCmd.PersistentFlags().String("bucketName", "", "s3 bucket name")
	AwsxS3ListCmd.PersistentFlags().String("arn", "", "arn")
	AwsxS3ListCmd.PersistentFlags().String("query", "", "query")
	AwsxS3ListCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxS3ListCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxS3ListCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
