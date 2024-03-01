# awsx-getlandingzonedetails
AWS plugin for get landing zone details

cli commands

    app-config
1. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getDiscoveredResourceCounts"
  

    ec2
1. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getEc2List"  

2. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getEc2ConfigByTag" --tagName="###########"  

3. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getEc2ConfigById" --instanceId="i-###########"  


    cdn 
1. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getCdnList"  

2. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getCdnFunctionList"  


    ApiGw
1. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getApiGwConfig" --apiKey="##########"  

2. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getApiGwList"  


    LoadBalancer
1. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getLbList"  
  

    DynamoDB
1. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getDynamoDbConfig" --tableName="##########"  
2. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getDynamoDbList"  


    ecs
1. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getEcsConfig" --clusterName="##########"
2. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getEcsList"  


    eks
1. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getEksConfig" --clusterName="##########"
2. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getEksList"  


    kinesis
1. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getKinesisConfig" --streamName="##########"
2. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getKinesisList"  
3. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getKinesisRecordList"  


    kms
1. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getKmsConfig" --keyId="##########"
2. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getKmsList"


    lambda
1. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getLambdaConfig" --functionName="##########"
2. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getLambdaList"	


    rds
1. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getRdsConfig" --arn="##########"
2. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getRdsList"	


    s3
1. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getS3Config" --bucketName="##########"
2. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getS3List"


    vpc
1. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getVpcConfig" --instanceId="##########"
2. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getVpcList"		


    waf
1. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getWafConfig" --instanceId="##########"
2. go run awsx-getLandingZoneDetails.go --vaultUrl=vault.synectiks.net --landingZoneId="1" --query="getWafList"			


   