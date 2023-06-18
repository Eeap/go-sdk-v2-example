package vpc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func GetVpc() error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		return err
	}
	svc := ec2.NewFromConfig(cfg)
	resp, err := svc.DescribeVpcs(context.TODO(), &ec2.DescribeVpcsInput{})

	for _, object := range resp.Vpcs {
		obj, _ := json.MarshalIndent(object, "", "\t")
		fmt.Println(string(obj))
	}
	return nil
}
