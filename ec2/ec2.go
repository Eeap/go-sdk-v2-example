package ec2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"log"
)

func GetEC2InstanceTypes() []types.InstanceTypeInfo {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		log.Fatal(err)
	}
	svc := ec2.NewFromConfig(cfg)
	filterType := "instance-type"
	instanceName := "t1.micro"
	resp, err := svc.DescribeInstanceTypes(context.TODO(), &ec2.DescribeInstanceTypesInput{
		Filters: []types.Filter{
			{
				Name:   &filterType,
				Values: []string{instanceName},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	return resp.InstanceTypes
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
