package EKS

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/spf13/cobra"
)

var AwsxEksConfigCmd = &cobra.Command{
	Use:   "getEksConfig",
	Short: "getEksConfig command gets eks configuration",
	Long:  `getEksConfig command gets eks configuration`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getEksConfig command")
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
			clusterName, _ := cmd.Flags().GetString("clusterName")
			if clusterName == "" {
				log.Printf("eks cluster name missing")
				err := cmd.Help()
				if err != nil {
					return
				}
				return
			}
			instances, err := GetEksInstanceByClusterName(clusterName, clientAuth, nil)
			if err != nil {
				log.Println("error getting getEksConfig by instance id: ", err)
				return
			}
			fmt.Println(instances)
		}
	},
}

func GetEksInstanceByClusterName(clusterName string, clientAuth *model.Auth, client *eks.EKS) (*eks.DescribeClusterOutput, error) {
	log.Println("getting aws eks instance of given cluster name: ", clusterName)
	if clusterName == "" {
		log.Println("cluster name missing")
		return nil, fmt.Errorf("cluster name missing")
	}
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.EKS_CLIENT).(*eks.EKS)
	}

	input := &eks.DescribeClusterInput{
		Name: aws.String(clusterName),
	}
	clusterDetailsResponse, err := client.DescribeCluster(input)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	return clusterDetailsResponse, nil
}

func init() {
	AwsxEksConfigCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxEksConfigCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxEksConfigCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxEksConfigCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxEksConfigCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxEksConfigCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxEksConfigCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxEksConfigCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxEksConfigCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxEksConfigCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxEksConfigCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxEksConfigCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxEksConfigCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxEksConfigCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxEksConfigCmd.PersistentFlags().String("ServiceName", "", "Service Name")
	AwsxEksConfigCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxEksConfigCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxEksConfigCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxEksConfigCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxEksConfigCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxEksConfigCmd.PersistentFlags().String("query", "", "query")
	AwsxEksConfigCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxEksConfigCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxEksConfigCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
