package LB

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/spf13/cobra"
	"log"
)

var AwsxLbListCmd = &cobra.Command{
	Use:   "getLbList",
	Short: "getLbList command gets list of load-balancers",
	Long:  `getLbList command gets list of load-balancers`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing getLbList command")
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
			instances, err := ListLbInstances(clientAuth, nil)
			if err != nil {
				log.Println("error getting getLbList: ", err)
				return
			}
			fmt.Println(instances)
		}
	},
}

func ListLbInstances(clientAuth *model.Auth, client *elbv2.ELBV2) (*elbv2.DescribeLoadBalancersOutput, error) {
	if client == nil {
		client = awsclient.GetClient(*clientAuth, awsclient.ELBV2_CLIENT).(*elbv2.ELBV2)
	}
	input := &elbv2.DescribeLoadBalancersInput{}
	result, err := client.DescribeLoadBalancers(input)
	if err != nil {
		log.Println("error in getting all load-balancer instances: ", err)
		return nil, err
	}
	return result, nil
}

func init() {
	AwsxLbListCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxLbListCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxLbListCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxLbListCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxLbListCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxLbListCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxLbListCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxLbListCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxLbListCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxLbListCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxLbListCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxLbListCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxLbListCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxLbListCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxLbListCmd.PersistentFlags().String("ServiceName", "", "Service Name")
	AwsxLbListCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxLbListCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxLbListCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxLbListCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxLbListCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxLbListCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxLbListCmd.PersistentFlags().String("query", "", "query")
	AwsxLbListCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxLbListCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxLbListCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
