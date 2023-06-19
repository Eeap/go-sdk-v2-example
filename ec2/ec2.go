package ec2

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"log"
	"strings"
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
	typeData := []string{}
	for _, object := range resp.InstanceTypes {
		typeData = append(typeData, string(object.InstanceType))
	}
	fmt.Printf("[%s]", strings.Join(typeData, ", "))
	return nil
}
