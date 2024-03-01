package RDS

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/spf13/cobra"
	"log"
)

var AwsxRdsListCmd = &cobra.Command{
	Use:   "getRdsList",
	Short: "getRdsList command gets list of rds instances of an aws account",
	Long:  `getRdsList command gets list of rds instances of an aws account`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getRdsList command")
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
			resp, err := ListRdsInstances(clientAuth, nil)
			if err != nil {
				log.Println("error getting getRdsList: ", err)
				return
			}
			fmt.Println(resp)
		}
	},
}

func ListRdsInstances(clientAuth *model.Auth, client *rds.RDS) ([]*rds.DescribeDBInstancesOutput, error) {
	log.Println("getting rds list")

	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.RDS_CLIENT).(*rds.RDS)
	}

	dbRequest := rds.DescribeDBInstancesInput{}
	dbResponse, err := client.DescribeDBInstances(&dbRequest)
	if err != nil {
		log.Println("error getting rds list", err)
		return nil, err
	}

	allInstances := []*rds.DescribeDBInstancesOutput{}
	for _, dbInstanceIdentifier := range dbResponse.DBInstances {
		instances, err := GetRdsInstanceByArn(*dbInstanceIdentifier.DBInstanceArn, clientAuth, client)
		if err != nil {
			log.Println("error in getting rds detail", err)
			continue
		}
		allInstances = append(allInstances, instances)
	}
	return allInstances, err
}

func init() {
	AwsxRdsListCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxRdsListCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxRdsListCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxRdsListCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxRdsListCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxRdsListCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxRdsListCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxRdsListCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxRdsListCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxRdsListCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxRdsListCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxRdsListCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxRdsListCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxRdsListCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxRdsListCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxRdsListCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxRdsListCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxRdsListCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxRdsListCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxRdsListCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxRdsListCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxRdsListCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxRdsListCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxRdsListCmd.PersistentFlags().String("functionName", "", "lambda function name")
	AwsxRdsListCmd.PersistentFlags().String("bucketName", "", "s3 bucket name")
	AwsxRdsListCmd.PersistentFlags().String("arn", "", "arn")
	AwsxRdsListCmd.PersistentFlags().String("query", "", "query")
	AwsxRdsListCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxRdsListCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxRdsListCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
