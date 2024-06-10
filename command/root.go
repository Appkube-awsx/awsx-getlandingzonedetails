package command

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-getlandingzonedetails/handler/API_GW"
	"github.com/Appkube-awsx/awsx-getlandingzonedetails/handler/CDN"
	"github.com/Appkube-awsx/awsx-getlandingzonedetails/handler/CONFIG_SERVICE"
	"github.com/Appkube-awsx/awsx-getlandingzonedetails/handler/DYNAMODB"
	"github.com/Appkube-awsx/awsx-getlandingzonedetails/handler/EC2"
	"github.com/Appkube-awsx/awsx-getlandingzonedetails/handler/ECS"
	"github.com/Appkube-awsx/awsx-getlandingzonedetails/handler/EKS"
	"github.com/Appkube-awsx/awsx-getlandingzonedetails/handler/KINESIS"
	"github.com/Appkube-awsx/awsx-getlandingzonedetails/handler/KMS"
	"github.com/Appkube-awsx/awsx-getlandingzonedetails/handler/LAMBDA"
	"github.com/Appkube-awsx/awsx-getlandingzonedetails/handler/LB"
	"github.com/Appkube-awsx/awsx-getlandingzonedetails/handler/RDS"
	"github.com/Appkube-awsx/awsx-getlandingzonedetails/handler/S3"
	"github.com/Appkube-awsx/awsx-getlandingzonedetails/handler/SSL"
	"github.com/Appkube-awsx/awsx-getlandingzonedetails/handler/VPC"
	"github.com/Appkube-awsx/awsx-getlandingzonedetails/handler/WAF"
	"github.com/spf13/cobra"
	"log"
)

var AwsxLandingZoneDetailsCmd = &cobra.Command{
	Use:   "getLandingZoneDetails",
	Short: "getLandingZoneDetails command gets aws resource details of a landing-zone",
	Long:  `getLandingZoneDetails command gets aws resource details of a landing-zone`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing command getLandingZoneDetails")
		var authFlag, clientAuth, err = authenticate.AuthenticateCommand(cmd)
		if err != nil {
			cmd.Help()
			return
		}
		if authFlag {
			queryName, _ := cmd.PersistentFlags().GetString("query")
			//elementType, _ := cmd.PersistentFlags().GetString("elementType")
			if queryName == "getSslConfig" {
				arn, _ := cmd.Flags().GetString("arn")
				resp, err := SSL.GetSslInstanceByArn(arn, clientAuth, nil)
				if err != nil {
					log.Println("error while getting ssl instance: ", err)
					cmd.Help()
					return
				}
				fmt.Println(resp)
			} else if queryName == "getSslList" {
				resp, err := SSL.ListSslInstances(clientAuth, nil)
				if err != nil {
					log.Println("error while getting ssl list: ", err)
					cmd.Help()
					return
				}
				fmt.Println(resp)
			} else if queryName == "getWafConfig" {
				instanceId, _ := cmd.Flags().GetString("instanceId")
				resp, err := WAF.GetWafInstanceById(instanceId, clientAuth, nil)
				if err != nil {
					log.Println("error while getting waf instance: ", err)
					cmd.Help()
					return
				}
				fmt.Println(resp)
			} else if queryName == "getWafList" {
				resp, err := WAF.ListWafInstances(clientAuth, nil)
				if err != nil {
					log.Println("error while getting waf list: ", err)
					cmd.Help()
					return
				}
				fmt.Println(resp)
			} else if queryName == "getVpcConfig" {
				instanceId, _ := cmd.Flags().GetString("instanceId")
				resp, err := VPC.GetVpcInstanceById(instanceId, clientAuth, nil)
				if err != nil {
					log.Println("error while getting vpc instance: ", err)
					cmd.Help()
					return
				}
				fmt.Println(resp)
			} else if queryName == "getVpcList" {
				resp, err := VPC.ListVpcInstances(clientAuth, nil)
				if err != nil {
					log.Println("error while getting vpc list: ", err)
					cmd.Help()
					return
				}
				fmt.Println(resp)
			} else if queryName == "getS3Config" {
				bucketName, _ := cmd.Flags().GetString("bucketName")
				resp, err := S3.GetS3InstanceByBucketName(bucketName, clientAuth, nil)
				if err != nil {
					log.Println("error while getting s3 instance: ", err)
					cmd.Help()
					return
				}
				fmt.Println(resp)
			} else if queryName == "getS3List" {
				resp, err := S3.ListS3Instances(clientAuth, nil)
				if err != nil {
					log.Println("error while getting s3 list: ", err)
					cmd.Help()
					return
				}
				fmt.Println(resp)
			} else if queryName == "getRdsConfig" {
				arn, _ := cmd.Flags().GetString("arn")
				resp, err := RDS.GetRdsInstanceByArn(arn, clientAuth, nil)
				if err != nil {
					log.Println("error while getting rds instance: ", err)
					cmd.Help()
					return
				}
				fmt.Println(resp)
			} else if queryName == "getRdsList" {
				resp, err := RDS.ListRdsInstances(clientAuth, nil)
				if err != nil {
					log.Println("error while getting rds list: ", err)
					cmd.Help()
					return
				}
				fmt.Println(resp)
			} else if queryName == "getLambdaConfig" {
				functionName, _ := cmd.Flags().GetString("functionName")
				resp, err := LAMBDA.GetLambdaFunctionByFunctionName(functionName, clientAuth, nil)
				if err != nil {
					log.Println("error while getting lambda function: ", err)
					cmd.Help()
					return
				}
				fmt.Println(resp)
			} else if queryName == "getLambdaList" {
				resp, err := LAMBDA.ListLambdaInstances(clientAuth, nil)
				if err != nil {
					log.Println("error while getting lambda list: ", err)
					cmd.Help()
					return
				}
				fmt.Println(resp)
			} else if queryName == "getKmsConfig" {
				keyId, _ := cmd.Flags().GetString("keyId")
				resp, err := KMS.GetKmsInstanceByKeyId(keyId, clientAuth, nil)
				if err != nil {
					log.Println("error while getting kms instance: ", err)
					cmd.Help()
					return
				}
				fmt.Println(resp)
			} else if queryName == "getKmsList" {
				resp, err := KMS.ListKmsInstances(clientAuth, nil)
				if err != nil {
					log.Println("error while getting kms list: ", err)
					cmd.Help()
					return
				}
				fmt.Println(resp)
			} else if queryName == "getKinesisRecordList" {
				resp, err := KINESIS.ListKinesisRecordInstances(clientAuth, nil)
				if err != nil {
					log.Println("error while getting kinesis records list: ", err)
					cmd.Help()
					return
				}
				fmt.Println(resp)
			} else if queryName == "getKinesisList" {
				resp, err := KINESIS.ListKinesisInstances(clientAuth, nil)
				if err != nil {
					log.Println("error while getting kinesis instances list: ", err)
					cmd.Help()
					return
				}
				fmt.Println(resp)
			} else if queryName == "getKinesisConfig" {
				streamName, _ := cmd.Flags().GetString("streamName")
				resp, err := KINESIS.GetKinesisInstanceByStreamName(streamName, clientAuth, nil)
				if err != nil {
					log.Println("error while getting kinesis instance: ", err)
					cmd.Help()
					return
				}
				fmt.Println(resp)
			} else if queryName == "getEksList" {
				resp, err := EKS.ListEksInstances(clientAuth, nil)
				if err != nil {
					log.Println("error while getting eks instances list: ", err)
					cmd.Help()
					return
				}
				fmt.Println(resp)
			} else if queryName == "getEksConfig" {
				clusterName, _ := cmd.Flags().GetString("clusterName")
				resp, err := EKS.GetEksInstanceByClusterName(clusterName, clientAuth, nil)
				if err != nil {
					log.Println("error while getting eks instance: ", err)
					cmd.Help()
					return
				}
				fmt.Println(resp)
			} else if queryName == "getEcsList" {
				resp, err := ECS.ListEcsInstances(clientAuth, nil)
				if err != nil {
					log.Println("error while getting ecs instances list: ", err)
					cmd.Help()
					return
				}
				fmt.Println(resp)
			} else if queryName == "getEcsConfig" {
				clusterName, _ := cmd.Flags().GetString("clusterName")
				resp, err := ECS.GetEcsInstanceByClusterName(clusterName, clientAuth, nil)
				if err != nil {
					log.Println("error while getting ecs instance: ", err)
					cmd.Help()
					return
				}
				fmt.Println(resp)
			} else if queryName == "getDynamoDbConfig" {
				tableName, _ := cmd.Flags().GetString("tableName")
				resp, err := DYNAMODB.GetDynamoDbInstanceByTableName(tableName, clientAuth, nil)
				if err != nil {
					log.Println("error while getting dynamo-db instance: ", err)
					cmd.Help()
					return
				}
				fmt.Println(resp)
			} else if queryName == "getDynamoDbList" {
				resp, err := DYNAMODB.ListDynamoDbInstance(clientAuth, nil)
				if err != nil {
					log.Println("error while getting dynamo-db list: ", err)
					return
				}
				fmt.Println(resp)
			} else if queryName == "getLbList" {
				resp, err := LB.ListLbInstances(clientAuth, nil)
				if err != nil {
					log.Println("error while getting load-balancer list: ", err)
					return
				}
				fmt.Println(resp)
			} else if queryName == "getApiGwList" {
				resp, err := API_GW.ListApiGwInstances(clientAuth, nil)
				if err != nil {
					log.Println("error while getting apigw list: ", err)
					return
				}
				fmt.Println(resp)
			} else if queryName == "getApiGwConfig" {
				apiKey, _ := cmd.Flags().GetString("apiKey")
				resp, err := API_GW.GetApiGwById(apiKey, clientAuth, nil)
				if err != nil {
					log.Println("error while getting apigw instance: ", err)
					cmd.Help()
					return
				}
				fmt.Println(resp)
			} else if queryName == "getCdnFunctionList" {
				resp, err := CDN.ListCdnFunctionInstances(clientAuth, nil)
				if err != nil {
					log.Println("error while getting cdn function names list: ", err)
					return
				}
				fmt.Println(resp)
			} else if queryName == "getCdnList" {
				resp, err := CDN.CdnDistributionConfigWithTagList(clientAuth, nil)
				if err != nil {
					log.Println("error while getting cdn distribution list: ", err)
					return
				}
				fmt.Println(resp)
			} else if queryName == "getDiscoveredResourceCounts" {
				counts, err := CONFIG_SERVICE.DiscoveredResourceCounts(clientAuth, nil)
				if err != nil {
					log.Println("error while getting discovered resource counts: ", err)
					return
				}
				fmt.Println(counts)
			} else if queryName == "getEc2List" {
				instances, err := EC2.ListEc2Instances(clientAuth, nil)
				if err != nil {
					log.Println("error while getting ec2 instances list: ", err)
					return
				}
				fmt.Println(instances)
			} else if queryName == "getEc2ConfigByTag" {
				tagName, _ := cmd.Flags().GetString("tagName")
				instances, err := EC2.GetEc2InstanceByTagName(tagName, clientAuth, nil)
				if err != nil {
					log.Println("error while getting ec2 instance by tag name: ", err)
					cmd.Help()
					return
				}
				fmt.Println(instances)
			} else if queryName == "getEc2ConfigById" {
				instanceId, _ := cmd.Flags().GetString("instanceId")
				instances, err := EC2.GetEc2InstanceById(instanceId, clientAuth, nil)
				if err != nil {
					log.Println("error while getting ec2 instance by instance id: ", err)
					cmd.Help()
					return
				}
				fmt.Println(instances)
			} else {
				fmt.Println("query not found")
			}

		} else {
			cmd.Help()
			return
		}
	},
}

func Execute() {
	if err := AwsxLandingZoneDetailsCmd.Execute(); err != nil {
		log.Printf("error executing getLandingZoneDetails command: %v\n", err)
	}
}

func init() {
	AwsxLandingZoneDetailsCmd.AddCommand(CONFIG_SERVICE.AwsxDiscoveredResourceCountsCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(EC2.AwsxEc2ListCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(EC2.AwsxEc2ConfigCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(CDN.AwsxCdnCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(CDN.AwsxCdnFunctionListCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(API_GW.AwsxApiGwConfigCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(API_GW.AwsxApiGwListCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(LB.AwsxLbListCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(DYNAMODB.AwsxDynamoDbListCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(DYNAMODB.AwsxDynamodDbConfigCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(ECS.AwsxEcsConfigCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(ECS.AwsxEcsListCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(EKS.AwsxEksConfigCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(EKS.AwsxEksListCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(KINESIS.AwsxKinesisConfigCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(KINESIS.AwsxKinesisListCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(KINESIS.AwsxKinesisRecordListCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(KMS.AwsxKmsListCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(KMS.AwsxKmsConfigCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(LAMBDA.AwsxLambdaListCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(LAMBDA.AwsxLambdaConfigCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(RDS.AwsxRdsConfigCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(RDS.AwsxRdsListCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(S3.AwsxS3ListCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(S3.AwsxS3ConfigCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(VPC.AwsxVpcListCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(VPC.AwsxVpcConfigCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(WAF.AwsxWafListCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(WAF.AwsxWafConfigCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(SSL.AwsxSslConfigCmd)
	AwsxLandingZoneDetailsCmd.AddCommand(SSL.AwsxSslListCmd)

	AwsxLandingZoneDetailsCmd.PersistentFlags().String("rootVolumeId", "", "root volume id")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("ebsVolume1Id", "", "ebs volume 1 id")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("ebsVolume2Id", "", "ebs volume 2 id")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("landingZoneId", "", "aws landingZoneId")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("serviceName", "", "service name")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("tagName", "", "tag name")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("apiKey", "", "api gateway key/id")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("tableName", "", "dynamo-db table name")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("streamName", "", "kinesis stream name")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("keyId", "", "kms key id")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("functionName", "", "lambda function name")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("bucketName", "", "s3 bucket name")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("arn", "", "arn")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("query", "", "query")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxLandingZoneDetailsCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
