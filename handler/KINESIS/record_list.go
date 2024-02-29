package KINESIS

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/spf13/cobra"
	"log"
)

var AwsxKinesisRecordListCmd = &cobra.Command{
	Use:   "getKinesisRecordList",
	Short: "getKinesisRecordList command gets list of kinesis instances of an aws account",
	Long:  `getKinesisRecordList command gets list of kinesis instances of an aws account`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getKinesisRecordList command")
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
			resp, err := ListKinesisRecordInstances(clientAuth, nil)
			if err != nil {
				log.Println("error getting getKinesisRecordList: ", err)
				return
			}
			fmt.Println(resp)
		}
	},
}

func ListKinesisRecordInstances(clientAuth *model.Auth, client *kinesis.Kinesis) ([]*kinesis.Record, error) {
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.KINESIS_CLIENT).(*kinesis.Kinesis)
	}
	streamList, err := GetKinesisStreamList(clientAuth, client)
	if err != nil {
		log.Println("error in getting kinesis streams", err)
		return nil, err
	}
	allRecords := []*kinesis.Record{}
	shardIteratorType := "LATEST"

	for _, name := range streamList.StreamNames {
		kinesisDetail, err := GetKinesisInstanceByStreamName(*name, clientAuth, client)
		if err != nil {
			log.Println("error in getting kinesis detail", err)
			continue
		}
		shards := kinesisDetail.StreamDescription.Shards
		for _, shard := range shards {
			shardIteratorInput := &kinesis.GetShardIteratorInput{
				ShardId:           shard.ShardId,
				ShardIteratorType: aws.String(shardIteratorType),
				StreamName:        aws.String(*name),
			}
			shardIteratorOutput, err := client.GetShardIterator(shardIteratorInput)
			if err != nil {
				fmt.Println("error in getting kinesis shard iterator:", err)
				continue
			}
			shardIterator := shardIteratorOutput.ShardIterator

			recordsInput := &kinesis.GetRecordsInput{
				ShardIterator: shardIterator,
			}
			recordsOutput, err := client.GetRecords(recordsInput)
			if err != nil {
				fmt.Println("error getting kinesis shard records', err")
				return nil, err
			}

			for _, record := range recordsOutput.Records {
				allRecords = append(allRecords, record)
			}
		}
	}

	return allRecords, nil
}

func init() {
	AwsxKinesisRecordListCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxKinesisRecordListCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxKinesisRecordListCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxKinesisRecordListCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxKinesisRecordListCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxKinesisRecordListCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxKinesisRecordListCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxKinesisRecordListCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxKinesisRecordListCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxKinesisRecordListCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxKinesisRecordListCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxKinesisRecordListCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxKinesisRecordListCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxKinesisRecordListCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxKinesisRecordListCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxKinesisRecordListCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxKinesisRecordListCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxKinesisRecordListCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxKinesisRecordListCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxKinesisRecordListCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxKinesisRecordListCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxKinesisRecordListCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxKinesisRecordListCmd.PersistentFlags().String("query", "", "query")
	AwsxKinesisRecordListCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxKinesisRecordListCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxKinesisRecordListCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
