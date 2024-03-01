package KINESIS

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/spf13/cobra"
	"log"
)

var AwsxKinesisListCmd = &cobra.Command{
	Use:   "getKinesisList",
	Short: "getKinesisList command gets list of kinesis instances of an aws account",
	Long:  `getKinesisList command gets list of kinesis instances of an aws account`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getKinesisList command")
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
			resp, err := ListKinesisInstances(clientAuth, nil)
			if err != nil {
				log.Println("error getting getKinesisList: ", err)
				return
			}
			fmt.Println(resp)
		}
	},
}

type KinesysObj struct {
	Stream interface{} `json:"stream"`
	Tags   interface{} `json:"tags"`
}

func ListKinesisInstances(clientAuth *model.Auth, client *kinesis.Kinesis) ([]KinesysObj, error) {
	log.Println("getting kinesis streams with tags")
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.KINESIS_CLIENT).(*kinesis.Kinesis)
	}
	streamList, err := GetKinesisStreamList(clientAuth, client)
	if err != nil {
		log.Println("error in getting kinesis streams", err)
		return nil, err
	}

	allKinesysDetailWithTag := []KinesysObj{}
	for _, name := range streamList.StreamNames {
		kinesysDetail, err := GetKinesisInstanceByStreamName(*name, clientAuth, client)
		if err != nil {
			log.Println("error in getting kinesis detail", err)
			continue
		}

		tagInput := &kinesis.ListTagsForStreamInput{
			StreamName: name,
		}
		tagOutput, err := client.ListTagsForStream(tagInput)
		if err != nil {
			log.Println("error in getting kinesis tag", err)
			continue
		}
		kinesysObj := KinesysObj{
			Stream: kinesysDetail,
			Tags:   tagOutput,
		}
		allKinesysDetailWithTag = append(allKinesysDetailWithTag, kinesysObj)
	}
	return allKinesysDetailWithTag, err
}

func init() {
	AwsxKinesisListCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxKinesisListCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxKinesisListCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxKinesisListCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxKinesisListCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxKinesisListCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxKinesisListCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxKinesisListCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxKinesisListCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxKinesisListCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxKinesisListCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxKinesisListCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxKinesisListCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxKinesisListCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxKinesisListCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxKinesisListCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxKinesisListCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxKinesisListCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxKinesisListCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxKinesisListCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxKinesisListCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxKinesisListCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxKinesisListCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxKinesisListCmd.PersistentFlags().String("functionName", "", "lambda function name")
	AwsxKinesisListCmd.PersistentFlags().String("bucketName", "", "s3 bucket name")
	AwsxKinesisListCmd.PersistentFlags().String("arn", "", "arn")
	AwsxKinesisListCmd.PersistentFlags().String("query", "", "query")
	AwsxKinesisListCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxKinesisListCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxKinesisListCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
