# go-sdk-v2-example
This is go-sdk-v2 sample

## Usage
You should check go version and install sdk for go v2(It is for 1.15 or later)
```bash
go mod init main

go get github.com/aws/aws-sdk-go-v2
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/service/ec2
```
Also, you should get your aws keys(access,secret). There are many ways to store keys. I stored `~/.aws/credentials` like this.
```bash
[default]
aws_access_key_id = your_access_key
aws_secret_access_key = your_secret_key
```
`context.TODO()` will find your access key in your credentials directory.

### GetVpc
TBA

### GetVPCSubnets
TBA

### GetEC2InstanceTypes
TBA