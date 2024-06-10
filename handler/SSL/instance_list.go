package SSL

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/service/acm"
	"github.com/spf13/cobra"
	"log"
)

var AwsxSslListCmd = &cobra.Command{
	Use:   "getSslList",
	Short: "getSslList command gets list of ssl instances of an aws account",
	Long:  `getSslList command gets list of ssl instances of an aws account`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getSslList command")
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
			resp, err := ListSslInstances(clientAuth, nil)
			if err != nil {
				log.Println("error getting getSslList: ", err)
				return
			}
			fmt.Println(resp)
		}
	},
}

func ListSslInstances(clientAuth *model.Auth, client *acm.ACM) (*acm.ListCertificatesOutput, error) {
	log.Println("getting ssl list")

	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.ACM_CLIENT).(*acm.ACM)
	}

	inputRequest := acm.ListCertificatesInput{}
	response, err := client.ListCertificates(&inputRequest)
	if err != nil {
		log.Println("error getting ssl list", err)
		return nil, err
	}
	return response, err
}

func init() {
	AwsxSslListCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxSslListCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxSslListCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxSslListCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxSslListCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxSslListCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxSslListCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxSslListCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxSslListCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxSslListCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxSslListCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxSslListCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxSslListCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxSslListCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxSslListCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxSslListCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxSslListCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxSslListCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxSslListCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxSslListCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxSslListCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxSslListCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxSslListCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxSslListCmd.PersistentFlags().String("functionName", "", "lambda function name")
	AwsxSslListCmd.PersistentFlags().String("bucketName", "", "s3 bucket name")
	AwsxSslListCmd.PersistentFlags().String("arn", "", "arn")
	AwsxSslListCmd.PersistentFlags().String("query", "", "query")
	AwsxSslListCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxSslListCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxSslListCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
