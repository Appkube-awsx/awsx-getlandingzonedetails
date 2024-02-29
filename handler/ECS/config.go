package ECS

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/spf13/cobra"
)

var AwsxEcsConfigCmd = &cobra.Command{
	Use:   "getEcsConfig",
	Short: "getEcsConfig command gets ecs configuration",
	Long:  `getEcsConfig command gets ecs configuration`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getEcsConfig command")
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
				log.Printf("ecs cluster name missing")
				err := cmd.Help()
				if err != nil {
					return
				}
				return
			}
			instances, err := GetEcsInstanceByClusterName(clusterName, clientAuth, nil)
			if err != nil {
				log.Println("error getting getEcsConfig by instance id: ", err)
				return
			}
			fmt.Println(instances)
		}
	},
}

func GetEcsInstanceByClusterName(clusterName string, clientAuth *model.Auth, client *ecs.ECS) (*ecs.DescribeClustersOutput, error) {
	log.Println("getting aws ecs instance of given cluster name: ", clusterName)
	if clusterName == "" {
		log.Println("cluster name missing")
		return nil, fmt.Errorf("cluster name missing")
	}
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.ECS_CLIENT).(*ecs.ECS)
	}

	input := &ecs.DescribeClustersInput{
		Clusters: []*string{aws.String(clusterName)},
	}
	clusterDetailsResponse, err := client.DescribeClusters(input)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	return clusterDetailsResponse, nil
}

func init() {
	AwsxEcsConfigCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxEcsConfigCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxEcsConfigCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxEcsConfigCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxEcsConfigCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxEcsConfigCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxEcsConfigCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxEcsConfigCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxEcsConfigCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxEcsConfigCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxEcsConfigCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxEcsConfigCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxEcsConfigCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxEcsConfigCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxEcsConfigCmd.PersistentFlags().String("ServiceName", "", "Service Name")
	AwsxEcsConfigCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxEcsConfigCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxEcsConfigCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxEcsConfigCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxEcsConfigCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxEcsConfigCmd.PersistentFlags().String("query", "", "query")
	AwsxEcsConfigCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxEcsConfigCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxEcsConfigCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
