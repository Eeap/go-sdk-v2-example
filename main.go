package main

import vpcTest "main/vpc"

func main() {
	err := vpcTest.GetVpc()
	if err != nil {
		panic(err)
	}
}
