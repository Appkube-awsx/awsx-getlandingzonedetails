package KMS

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/spf13/cobra"
	"log"
)

var AwsxKmsListCmd = &cobra.Command{
	Use:   "getKmsList",
	Short: "getKmsList command gets list of kms instances of an aws account",
	Long:  `getKmsList command gets list of kms instances of an aws account`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getKmsList command")
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
			resp, err := ListKmsInstances(clientAuth, nil)
			if err != nil {
				log.Println("error getting getKmsList: ", err)
				return
			}
			fmt.Println(resp)
		}
	},
}

func ListKmsInstances(clientAuth *model.Auth, client *kms.KMS) ([]*kms.DescribeKeyOutput, error) {
	log.Println("getting kms list")

	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.KMS_CLIENT).(*kms.KMS)
	}

	input := &kms.ListKeysInput{}
	keyList, err := client.ListKeys(input)
	if err != nil {
		log.Println("error getting kms list", err)
		return nil, err
	}
	allRecords := []*kms.DescribeKeyOutput{}
	for _, keyEntry := range keyList.Keys {
		instance, err := GetKmsInstanceByKeyId(*keyEntry.KeyId, clientAuth, client)
		if err != nil {
			log.Println("error in getting kms detail", err)
			continue
		}
		allRecords = append(allRecords, instance)
	}
	return allRecords, err
}

func init() {
	AwsxKmsListCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxKmsListCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxKmsListCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxKmsListCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxKmsListCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxKmsListCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxKmsListCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxKmsListCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxKmsListCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxKmsListCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxKmsListCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxKmsListCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxKmsListCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxKmsListCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxKmsListCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxKmsListCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxKmsListCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxKmsListCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxKmsListCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxKmsListCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxKmsListCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxKmsListCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxKmsListCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxKmsListCmd.PersistentFlags().String("query", "", "query")
	AwsxKmsListCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxKmsListCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxKmsListCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
