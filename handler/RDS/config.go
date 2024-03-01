package RDS

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/aws"
	"log"

	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/spf13/cobra"
)

var AwsxRdsConfigCmd = &cobra.Command{
	Use:   "getRdsConfig",
	Short: "getRdsConfig command gets rds configuration",
	Long:  `getRdsConfig command gets rds configuration`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getRdsConfig command")
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
			arn, _ := cmd.Flags().GetString("arn")
			if arn == "" {
				log.Printf("rds arn missing")
				err := cmd.Help()
				if err != nil {
					return
				}
				return
			}

			instances, err := GetRdsInstanceByArn(arn, clientAuth, nil)
			if err != nil {
				log.Println("error getting getRdsConfig by function name: ", err)
				return
			}
			fmt.Println(instances)

		}
	},
}

func GetRdsInstanceByArn(arn string, clientAuth *model.Auth, client *rds.RDS) (*rds.DescribeDBInstancesOutput, error) {
	log.Println("getting aws rds instance of given arn: ", arn)
	if arn == "" {
		log.Println("arn missing")
		return nil, fmt.Errorf("arn missing")
	}
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.RDS_CLIENT).(*rds.RDS)
	}
	input := &rds.DescribeDBInstancesInput{
		DBInstanceIdentifier: aws.String(arn),
	}
	response, err := client.DescribeDBInstances(input)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	return response, nil
}

func init() {
	AwsxRdsConfigCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxRdsConfigCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxRdsConfigCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxRdsConfigCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxRdsConfigCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxRdsConfigCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxRdsConfigCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxRdsConfigCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxRdsConfigCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxRdsConfigCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxRdsConfigCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxRdsConfigCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxRdsConfigCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxRdsConfigCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxRdsConfigCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxRdsConfigCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxRdsConfigCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxRdsConfigCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxRdsConfigCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxRdsConfigCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxRdsConfigCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxRdsConfigCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxRdsConfigCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxRdsConfigCmd.PersistentFlags().String("functionName", "", "lambda function name")
	AwsxRdsConfigCmd.PersistentFlags().String("bucketName", "", "s3 bucket name")
	AwsxRdsConfigCmd.PersistentFlags().String("arn", "", "arn")
	AwsxRdsConfigCmd.PersistentFlags().String("query", "", "query")
	AwsxRdsConfigCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxRdsConfigCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxRdsConfigCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
