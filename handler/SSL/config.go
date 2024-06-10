package SSL

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/acm"
	"log"

	"github.com/spf13/cobra"
)

var AwsxSslConfigCmd = &cobra.Command{
	Use:   "getSslConfig",
	Short: "getSslConfig command gets ssl configuration",
	Long:  `getSslConfig command gets ssl configuration`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getSslConfig command")
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
				log.Printf("ssl arn missing")
				err := cmd.Help()
				if err != nil {
					return
				}
				return
			}

			instances, err := GetSslInstanceByArn(arn, clientAuth, nil)
			if err != nil {
				log.Println("error getting getSslConfig by function name: ", err)
				return
			}
			fmt.Println(instances)

		}
	},
}

func GetSslInstanceByArn(arn string, clientAuth *model.Auth, client *acm.ACM) (*acm.DescribeCertificateOutput, error) {
	log.Println("getting aws ssl instance of given arn: ", arn)
	if arn == "" {
		log.Println("arn missing")
		return nil, fmt.Errorf("arn missing")
	}
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.ACM_CLIENT).(*acm.ACM)
	}
	response, err := client.DescribeCertificate(&acm.DescribeCertificateInput{
		CertificateArn: aws.String(arn),
	})
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	return response, nil
}

func init() {
	AwsxSslConfigCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxSslConfigCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxSslConfigCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxSslConfigCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxSslConfigCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxSslConfigCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxSslConfigCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxSslConfigCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxSslConfigCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxSslConfigCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxSslConfigCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxSslConfigCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxSslConfigCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxSslConfigCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxSslConfigCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxSslConfigCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxSslConfigCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxSslConfigCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxSslConfigCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxSslConfigCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxSslConfigCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxSslConfigCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxSslConfigCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxSslConfigCmd.PersistentFlags().String("functionName", "", "lambda function name")
	AwsxSslConfigCmd.PersistentFlags().String("bucketName", "", "s3 bucket name")
	AwsxSslConfigCmd.PersistentFlags().String("arn", "", "arn")
	AwsxSslConfigCmd.PersistentFlags().String("query", "", "query")
	AwsxSslConfigCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxSslConfigCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxSslConfigCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
