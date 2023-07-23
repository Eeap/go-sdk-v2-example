package ec2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"log"
)

type EC2DataTypes struct {
	instanceTypes []types.InstanceTypeInfo
}

func GetEC2InstanceTypes() []types.InstanceTypeInfo {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		log.Fatal(err)
	}
	client := ec2.NewFromConfig(cfg)
	ec2InstanceTypes := &EC2DataTypes{}
	pagingInstanceTypes(client, nil, ec2InstanceTypes)
	return ec2InstanceTypes.instanceTypes
	//for _, object := range resp.InstanceTypes {
	//	objectData := make(map[string]any)
	//	objectData["InstanceType"] = string(object.InstanceType)
	//	objectData["InstanceStorageInfo"] = object.InstanceStorageInfo
	//	objectData["EbsInfo"] = object.EbsInfo
	//	objectData["ProcessorInfo"] = object.ProcessorInfo
	//	objectData["GpuInfo"] = object.GpuInfo
	//	objectData["NetworkInfo"] = object.NetworkInfo
	//	objectData["MemoryInfo"] = object.MemoryInfo
	//
	//	result, _ := json.MarshalIndent(objectData, "", "	")
	//	fmt.Println(string(result))
	//}
	//return nil
}

func pagingInstanceTypes(client *ec2.Client, nextToken *string, ec2InstanceTypes *EC2DataTypes) {
	//filterType := "instance-type"
	//instanceName := "t1.micro"
	resp, err := client.DescribeInstanceTypes(context.TODO(), &ec2.DescribeInstanceTypesInput{
		NextToken: nextToken,
		//Filters: []types.Filter{
		//	{
		//		Name:   &filterType,
		//		Values: []string{instanceName},
		//	},
		//},
	})
	if err != nil {
		return
	}
	if resp.NextToken != nil {
		pagingInstanceTypes(client, resp.NextToken, ec2InstanceTypes)
	}
	ec2InstanceTypes.instanceTypes = append(ec2InstanceTypes.instanceTypes, resp.InstanceTypes...)
}
