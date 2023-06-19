package main

import (
	ec2Test "main/ec2"
	vpcTest "main/vpc"
)

func main() {
	errVpc := vpcTest.GetVpc()
	if errVpc != nil {
		panic(errVpc)
	}
	errEC2 := ec2Test.GetEC2InstanceTypes()
	if errEC2 != nil {
		panic(errEC2)
	}
}
