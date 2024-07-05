package CLOUDWATCH

import (
	"encoding/json"
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/spf13/cobra"
	"log"
)

type FilteredLogEvent struct {
	Timestamp int64
	Message   string
	// Add other fields as needed
}

var AwsxCwAlarmListCmd = &cobra.Command{
	Use:   "getCwAlarmList",
	Short: "getCwAlarmList command gets list of cloudwatch alarm rest api instances of an aws account",
	Long:  `getCwAlarmList command gets list of cloudwatch alarm rest api instances of an aws account`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getCwAlarmList command")
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
				log.Printf("instance-id missing")
				err := cmd.Help()
				if err != nil {
					return
				}
				return
			}
			if instanceId != "" {
				instances, err := ListCwAlarms(instanceId, clientAuth, nil)
				if err != nil {
					log.Println("error getting getCwAlarmList: ", err)
					return
				}
				fmt.Println(instances)
			}

		}

	},
}

func ListCwAlarms(instanceId string, clientAuth *model.Auth, client *cloudwatch.CloudWatch) ([]*cloudwatch.MetricAlarm, error) {
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.CLOUDWATCH).(*cloudwatch.CloudWatch)
	}
	allRecords := []*cloudwatch.MetricAlarm{}
	var nextToken *string
	for {
		// Describe alarms with pagination
		input := &cloudwatch.DescribeAlarmsInput{
			NextToken: nextToken,
		}
		result, err := client.DescribeAlarms(input)
		if err != nil {
			log.Println("failed to describe alarms: %v", err)
			return nil, err
		}
		// Print the alarms related to the specific EC2 instance
		for _, alarm := range result.MetricAlarms {
			for _, dimension := range alarm.Dimensions {
				if aws.StringValue(dimension.Value) == instanceId {
					allRecords = append(allRecords, alarm)
				}
			}

		}

		// Check if there are more alarms to retrieve
		if result.NextToken == nil {
			break
		}
		nextToken = result.NextToken
	}
	return allRecords, nil
}
func ListCwLogsStream(logGroupName string, clientAuth *model.Auth, client *cloudwatchlogs.CloudWatchLogs) ([]*string, error) {
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.CLOUDWATCH_LOG).(*cloudwatchlogs.CloudWatchLogs)
	}

	var logStreamNames []*string

	input := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: aws.String(logGroupName),
	}
	err := client.DescribeLogStreamsPages(input,
		func(page *cloudwatchlogs.DescribeLogStreamsOutput, lastPage bool) bool {
			for _, logStream := range page.LogStreams {
				logStreamNames = append(logStreamNames, logStream.LogStreamName)
				//fmt.Printf("Instance ID: %s\n", aws.StringValue(logStream.LogStreamName))
			}
			return !lastPage
		})
	if err != nil {
		log.Printf("failed to describe log streams for log group %s: %v", logGroupName, err)
		return nil, err
	}
	return logStreamNames, nil
}
func ListCwLogsGorup(clientAuth *model.Auth, client *cloudwatchlogs.CloudWatchLogs) ([]string, error) {
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.CLOUDWATCH_LOG).(*cloudwatchlogs.CloudWatchLogs)
	}

	var logGroupNames []string
	input := &cloudwatchlogs.DescribeLogGroupsInput{}

	err := client.DescribeLogGroupsPages(input, func(page *cloudwatchlogs.DescribeLogGroupsOutput, lastPage bool) bool {
		for _, logGroup := range page.LogGroups {
			logGroupNames = append(logGroupNames, *logGroup.LogGroupName)
			fmt.Println(*logGroup.LogGroupName)
		}
		return !lastPage
	})

	if err != nil {
		log.Fatalf("Failed to describe log groups: %v", err)
		return nil, err
	}

	return logGroupNames, nil
}
func ListCwEvent(instanceId string, clientAuth *model.Auth, client *cloudtrail.CloudTrail) (string, error) {
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.CLOUDTRAIL_CLIENT).(*cloudtrail.CloudTrail)
	}

	input := &cloudtrail.LookupEventsInput{
		LookupAttributes: []*cloudtrail.LookupAttribute{
			{
				AttributeKey:   aws.String("ResourceName"),
				AttributeValue: aws.String(instanceId),
			},
		},
	}

	result, err := client.LookupEvents(input)
	if err != nil {
		return "", fmt.Errorf("error looking up events: %v", err)
	}

	jsonResp, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error encoding result to JSON: %v", err)
	}

	return string(jsonResp), nil
}
func init() {
	AwsxCwAlarmListCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxCwAlarmListCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxCwAlarmListCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxCwAlarmListCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxCwAlarmListCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxCwAlarmListCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxCwAlarmListCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxCwAlarmListCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxCwAlarmListCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxCwAlarmListCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxCwAlarmListCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxCwAlarmListCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxCwAlarmListCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxCwAlarmListCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxCwAlarmListCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxCwAlarmListCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxCwAlarmListCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxCwAlarmListCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxCwAlarmListCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxCwAlarmListCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxCwAlarmListCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxCwAlarmListCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxCwAlarmListCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxCwAlarmListCmd.PersistentFlags().String("functionName", "", "lambda function name")
	AwsxCwAlarmListCmd.PersistentFlags().String("bucketName", "", "s3 bucket name")
	AwsxCwAlarmListCmd.PersistentFlags().String("arn", "", "arn")
	AwsxCwAlarmListCmd.PersistentFlags().String("query", "", "query")
	AwsxCwAlarmListCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxCwAlarmListCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxCwAlarmListCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
