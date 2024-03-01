package EKS

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/spf13/cobra"
	"log"
)

var AwsxEksListCmd = &cobra.Command{
	Use:   "getEksList",
	Short: "getEksList command gets list of eks instances of an aws account",
	Long:  `getEksList command gets list of eks instances of an aws account`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getEksList command")
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
			resp, err := ListEksInstances(clientAuth, nil)
			if err != nil {
				log.Println("error getting getEksList: ", err)
				return
			}
			fmt.Println(resp)
		}
	},
}

func ListEksInstances(clientAuth *model.Auth, client *eks.EKS) ([]*eks.DescribeClusterOutput, error) {
	log.Println("getting eks cluster list")

	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.EKS_CLIENT).(*eks.EKS)
	}

	request := &eks.ListClustersInput{}
	response, err := client.ListClusters(request)
	if err != nil {
		log.Println("error getting eks cluster list", err)
		return nil, err
	}
	allClusters := []*eks.DescribeClusterOutput{}

	for _, clusterName := range response.Clusters {
		clusterDetail, err := GetEksInstanceByClusterName(*clusterName, clientAuth, client)
		if err != nil {
			log.Println("error in getting eks instance: ", err)
			continue
		}
		allClusters = append(allClusters, clusterDetail)
	}
	return allClusters, err
}

func init() {
	AwsxEksListCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxEksListCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxEksListCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxEksListCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxEksListCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxEksListCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxEksListCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxEksListCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxEksListCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxEksListCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxEksListCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxEksListCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxEksListCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxEksListCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxEksListCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxEksListCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxEksListCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxEksListCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxEksListCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxEksListCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxEksListCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxEksListCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxEksListCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxEksListCmd.PersistentFlags().String("functionName", "", "lambda function name")
	AwsxEksListCmd.PersistentFlags().String("bucketName", "", "s3 bucket name")
	AwsxEksListCmd.PersistentFlags().String("arn", "", "arn")
	AwsxEksListCmd.PersistentFlags().String("query", "", "query")
	AwsxEksListCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxEksListCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxEksListCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
