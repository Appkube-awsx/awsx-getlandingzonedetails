package CDN

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/spf13/cobra"
	"log"
)

var AwsxCdnCmd = &cobra.Command{
	Use:   "getCdnList",
	Short: "getCdnList command gets list of cdn functions of an aws account",
	Long:  `getCdnList command gets list of cdn functions of an aws account`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getCdnList command")
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
			resp, err := CdnDistributionConfigWithTagList(clientAuth, nil)
			if err != nil {
				log.Println("error getting getCdnList: ", err)
				return
			}
			fmt.Println(resp)
		}
	},
}

func CloudFrontDistributionList(clientAuth *model.Auth, client *cloudfront.CloudFront) (*cloudfront.ListDistributionsOutput, error) {
	log.Println("getting aws (cdn) cloudfront distribution list")
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.CLOUD_FRONT_CLIENT).(*cloudfront.CloudFront)
	}
	input := &cloudfront.ListDistributionsInput{}
	response, err := client.ListDistributions(input)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	return response, err
}

type Cdn struct {
	Distribution       interface{} `json:"distribution"`
	DistributionConfig interface{} `json:"distribution_config"`
	Tags               interface{} `json:"tags"`
}

func CdnDistributionConfigWithTagList(clientAuth *model.Auth, client *cloudfront.CloudFront) ([]Cdn, error) {
	log.Println("getting aws (cdn) cloudfront distribution list with tags")
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.CLOUD_FRONT_CLIENT).(*cloudfront.CloudFront)
	}
	distributionList, err := CloudFrontDistributionList(clientAuth, client)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}

	cdnList := []Cdn{}
	for _, distributionItem := range distributionList.DistributionList.Items {
		configInput := &cloudfront.GetDistributionConfigInput{
			Id: distributionItem.Id,
		}
		distributionConfigOutput, err := client.GetDistributionConfig(configInput)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		tagInput := &cloudfront.ListTagsForResourceInput{
			Resource: distributionItem.ARN,
		}
		tagOutput, err := client.ListTagsForResource(tagInput)
		if err != nil {
			fmt.Println("error in getting cdn tags:", err)
			continue
		}
		cdn := Cdn{
			Distribution:       distributionItem,
			DistributionConfig: distributionConfigOutput,
			Tags:               tagOutput,
		}
		cdnList = append(cdnList, cdn)

	}
	return cdnList, err
}

func init() {
	AwsxCdnCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxCdnCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxCdnCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxCdnCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxCdnCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxCdnCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxCdnCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxCdnCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxCdnCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxCdnCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxCdnCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxCdnCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxCdnCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxCdnCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxCdnCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxCdnCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxCdnCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxCdnCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxCdnCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxCdnCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxCdnCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxCdnCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxCdnCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxCdnCmd.PersistentFlags().String("functionName", "", "lambda function name")
	AwsxCdnCmd.PersistentFlags().String("bucketName", "", "s3 bucket name")
	AwsxCdnCmd.PersistentFlags().String("arn", "", "arn")
	AwsxCdnCmd.PersistentFlags().String("query", "", "query")
	AwsxCdnCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxCdnCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxCdnCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
