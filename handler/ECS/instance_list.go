package ECS

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/spf13/cobra"
	"log"
)

var AwsxEcsListCmd = &cobra.Command{
	Use:   "getEcsList",
	Short: "getEcsList command gets list of ecs instances of an aws account",
	Long:  `getEcsList command gets list of ecs instances of an aws account`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getEcsList command")
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
			resp, err := ListEcsInstances(clientAuth, nil)
			if err != nil {
				log.Println("error getting getEcsList: ", err)
				return
			}
			fmt.Println(resp)
		}
	},
}

func ListEcsInstances(clientAuth *model.Auth, client *ecs.ECS) ([]*ecs.DescribeClustersOutput, error) {
	log.Println("getting ecs cluster list")

	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.ECS_CLIENT).(*ecs.ECS)
	}

	request := &ecs.ListClustersInput{}
	response, err := client.ListClusters(request)
	if err != nil {
		log.Println("error getting ecs cluster list", err)
		return nil, err
	}
	allClusters := []*ecs.DescribeClustersOutput{}

	for _, clusterArn := range response.ClusterArns {
		clusterDetail, err := getEcsInstanceByClusterArn(*clusterArn, clientAuth, client)
		if err != nil {
			log.Println("error getting ecs cluster by cluster arn", err)
			continue
		}
		allClusters = append(allClusters, clusterDetail)

		for _, cluster := range clusterDetail.Clusters {
			input := &ecs.ListTagsForResourceInput{
				ResourceArn: cluster.ClusterArn,
			}
			tagOutput, err := client.ListTagsForResource(input)
			if err != nil {
				log.Println("error in getting ecs tag: ", err)
				continue
			}
			cluster.SetTags(tagOutput.Tags)
			ecsInstance, err := GetEcsInstanceByClusterName(*cluster.ClusterName, clientAuth, client)
			if err != nil {
				log.Println("error in getting ecs cluster by cluster name: ", err)
				continue
			}
			if len(ecsInstance.Clusters) == 0 {
				log.Println("ecs cluster not found")
				continue
			}

		}

	}
	return allClusters, err
}

func getEcsInstanceByClusterArn(clusterArn string, clientAuth *model.Auth, client *ecs.ECS) (*ecs.DescribeClustersOutput, error) {
	log.Println("getting aws ecs instance of given cluster arn: ", clusterArn)
	if clusterArn == "" {
		log.Println("cluster arn missing")
		return nil, fmt.Errorf("cluster arn missing")
	}
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.ECS_CLIENT).(*ecs.ECS)
	}
	input := &ecs.DescribeClustersInput{
		Clusters: []*string{aws.String(clusterArn)},
	}
	clusterDetailsResponse, err := client.DescribeClusters(input)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}
	return clusterDetailsResponse, nil
}

func init() {
	AwsxEcsListCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxEcsListCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxEcsListCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxEcsListCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxEcsListCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxEcsListCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxEcsListCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxEcsListCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxEcsListCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxEcsListCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxEcsListCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxEcsListCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxEcsListCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxEcsListCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxEcsListCmd.PersistentFlags().String("ServiceName", "", "Service Name")
	AwsxEcsListCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxEcsListCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxEcsListCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxEcsListCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxEcsListCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxEcsListCmd.PersistentFlags().String("query", "", "query")
	AwsxEcsListCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxEcsListCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxEcsListCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
