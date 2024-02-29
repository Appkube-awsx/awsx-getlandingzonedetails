package KINESIS

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/spf13/cobra"
)

var AwsxKinesisConfigCmd = &cobra.Command{
	Use:   "getKinesisConfig",
	Short: "getKinesisConfig command gets kinesis configuration",
	Long:  `getKinesisConfig command gets kinesis configuration`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getKinesisConfig command")
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
			streamName, _ := cmd.Flags().GetString("streamName")
			if streamName == "" {
				log.Printf("kinesis stream name missing")
				err := cmd.Help()
				if err != nil {
					return
				}
				return
			}
			if streamName != "" {
				instances, err := GetKinesisInstanceByStreamName(streamName, clientAuth, nil)
				if err != nil {
					log.Println("error getting getKinesisConfig by instance id: ", err)
					return
				}
				fmt.Println(instances)
			}
		}
	},
}

func GetKinesisInstanceByStreamName(streamName string, clientAuth *model.Auth, client *kinesis.Kinesis) (*kinesis.DescribeStreamOutput, error) {
	log.Println("getting aws kinesis instance of given stream name: ", streamName)
	if streamName == "" {
		log.Println("stream name missing")
		return nil, fmt.Errorf("stream name missing")
	}
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.KINESIS_CLIENT).(*kinesis.Kinesis)
	}

	input := &kinesis.DescribeStreamInput{
		StreamName: aws.String(streamName),
	}
	kinesisData, err := client.DescribeStream(input)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	return kinesisData, nil
}

func GetKinesisStreamList(clientAuth *model.Auth, client *kinesis.Kinesis) (*kinesis.ListStreamsOutput, error) {
	log.Println("getting kinesis streams")
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.KINESIS_CLIENT).(*kinesis.Kinesis)
	}
	input := &kinesis.ListStreamsInput{}
	streamList, err := client.ListStreams(input)
	if err != nil {
		log.Println("Error: in getting kinesis streams", err)
		return nil, err
	}
	log.Println(streamList)
	return streamList, err
}

func init() {
	AwsxKinesisConfigCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxKinesisConfigCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxKinesisConfigCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxKinesisConfigCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxKinesisConfigCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxKinesisConfigCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxKinesisConfigCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxKinesisConfigCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxKinesisConfigCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxKinesisConfigCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxKinesisConfigCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxKinesisConfigCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxKinesisConfigCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxKinesisConfigCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxKinesisConfigCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxKinesisConfigCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxKinesisConfigCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxKinesisConfigCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxKinesisConfigCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxKinesisConfigCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxKinesisConfigCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxKinesisConfigCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxKinesisConfigCmd.PersistentFlags().String("query", "", "query")
	AwsxKinesisConfigCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxKinesisConfigCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxKinesisConfigCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
