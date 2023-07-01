package ec2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"log"
)

func GetEC2InstanceTypes() error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		log.Fatal(err)
	}
	svc := ec2.NewFromConfig(cfg)
	resp, err := svc.DescribeInstanceTypes(context.TODO(), &ec2.DescribeInstanceTypesInput{})
	if err != nil {
		log.Fatal(err)
	}
	for _, object := range resp.InstanceTypes {
		objectData := make(map[string]any)
		objectData["InstanceType"] = string(object.InstanceType)
		objectData["InstanceStorageInfo"] = object.InstanceStorageInfo
		objectData["EbsInfo"] = object.EbsInfo
		objectData["ProcessorInfo"] = object.ProcessorInfo
		objectData["GpuInfo"] = object.GpuInfo
		objectData["NetworkInfo"] = object.NetworkInfo
		objectData["MemoryInfo"] = object.MemoryInfo

		result, _ := json.MarshalIndent(objectData, "", "	")
		fmt.Println(string(result))
	}
	return nil
}
