package S3

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/aws"
	"log"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

var AwsxS3ConfigCmd = &cobra.Command{
	Use:   "getS3Config",
	Short: "getS3Config command gets s3 configuration",
	Long:  `getS3Config command gets s3 configuration`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getS3Config command")
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
			bucketName, _ := cmd.Flags().GetString("bucketName")
			if bucketName == "" {
				log.Printf("s3 bucket name missing")
				err := cmd.Help()
				if err != nil {
					return
				}
				return
			}

			instances, err := GetS3InstanceByBucketName(bucketName, clientAuth, nil)
			if err != nil {
				log.Println("error getting getS3Config by bucket name: ", err)
				return
			}
			fmt.Println(instances)

		}
	},
}

type S3Bucket struct {
	Bucket interface{} `json:"bucket"`
	Tags   interface{} `json:"tags"`
}

func GetS3InstanceByBucketName(bucketName string, clientAuth *model.Auth, client *s3.S3) (*S3Bucket, error) {
	log.Println("getting aws s3 instance of given bucketName: ", bucketName)
	if bucketName == "" {
		log.Println("bucket name missing")
		return nil, fmt.Errorf("bucket name missing")
	}
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.S3_CLIENT).(*s3.S3)
	}
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	}
	response, err := client.ListObjectsV2(input)
	if err != nil {
		log.Println("error in getting bucket detail ", err)
		return nil, err
	}
	s3Bucket := S3Bucket{
		Bucket: response,
	}
	tagInput := &s3.GetBucketTaggingInput{
		Bucket: aws.String(bucketName),
	}
	tagOutput, err := client.GetBucketTagging(tagInput)
	if err != nil {
		log.Println("error in getting bucket tags ", err)
		return &s3Bucket, nil
	}
	s3Bucket.Tags = tagOutput

	return &s3Bucket, nil
}

func init() {
	AwsxS3ConfigCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxS3ConfigCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxS3ConfigCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxS3ConfigCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxS3ConfigCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxS3ConfigCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxS3ConfigCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxS3ConfigCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxS3ConfigCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxS3ConfigCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxS3ConfigCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxS3ConfigCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxS3ConfigCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxS3ConfigCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxS3ConfigCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxS3ConfigCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxS3ConfigCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxS3ConfigCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxS3ConfigCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxS3ConfigCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxS3ConfigCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxS3ConfigCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxS3ConfigCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxS3ConfigCmd.PersistentFlags().String("functionName", "", "lambda function name")
	AwsxS3ConfigCmd.PersistentFlags().String("bucketName", "", "s3 bucket name")
	AwsxS3ConfigCmd.PersistentFlags().String("arn", "", "arn")
	AwsxS3ConfigCmd.PersistentFlags().String("query", "", "query")
	AwsxS3ConfigCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxS3ConfigCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxS3ConfigCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
