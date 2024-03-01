package KMS

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"log"

	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/spf13/cobra"
)

var AwsxKmsConfigCmd = &cobra.Command{
	Use:   "getKmsConfig",
	Short: "getKmsConfig command gets kms configuration",
	Long:  `getKmsConfig command gets kms configuration`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getKmsConfig command")
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
			keyId, _ := cmd.Flags().GetString("keyId")
			if keyId == "" {
				log.Printf("kms key id missing")
				err := cmd.Help()
				if err != nil {
					return
				}
				return
			}

			instances, err := GetKmsInstanceByKeyId(keyId, clientAuth, nil)
			if err != nil {
				log.Println("error getting getKmsConfig by key id: ", err)
				return
			}
			fmt.Println(instances)

		}
	},
}

func GetKmsInstanceByKeyId(keyId string, clientAuth *model.Auth, client *kms.KMS) (*kms.DescribeKeyOutput, error) {
	log.Println("getting aws kms instance of given key id: ", keyId)
	if keyId == "" {
		log.Println("key id missing")
		return nil, fmt.Errorf("key id missing")
	}
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.KMS_CLIENT).(*kms.KMS)
	}

	kmsRequest := &kms.DescribeKeyInput{
		KeyId: &keyId,
	}
	kmsResponse, err := client.DescribeKey(kmsRequest)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	return kmsResponse, nil
}

func init() {
	AwsxKmsConfigCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxKmsConfigCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxKmsConfigCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxKmsConfigCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxKmsConfigCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxKmsConfigCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxKmsConfigCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxKmsConfigCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxKmsConfigCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxKmsConfigCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxKmsConfigCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxKmsConfigCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxKmsConfigCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxKmsConfigCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxKmsConfigCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxKmsConfigCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxKmsConfigCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxKmsConfigCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxKmsConfigCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxKmsConfigCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxKmsConfigCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxKmsConfigCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxKmsConfigCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxKmsConfigCmd.PersistentFlags().String("functionName", "", "lambda function name")
	AwsxKmsConfigCmd.PersistentFlags().String("bucketName", "", "s3 bucket name")
	AwsxKmsConfigCmd.PersistentFlags().String("arn", "", "arn")
	AwsxKmsConfigCmd.PersistentFlags().String("query", "", "query")
	AwsxKmsConfigCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxKmsConfigCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxKmsConfigCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
