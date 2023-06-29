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
`DescribeVpcssInput`
```bash
{
  # Check whether you have the required permissions for the action
  DryRun *bool => true or false (default false),
  # There are many FilterTypes (cidr, cidr-block-association.cidr-block, state, owner-id, is-default, vpc-id, ...etc)
  Filters []types.Filter,
  # The maximum number of items to return
  MaxResults *int32,
  # The token returned from a previous paginated req
  NextToken *string,
  # One or more VPC ids (default all VPCs)
  VpcIds []string
}
```
`DescribeVpcsOutput`
```bash
{
  NextToken *string,
  # Information about one or more vpc
  Vpcs []types.Vpc,
  # Metadata about operation's result (req ID, ...etc)
  ResultMetadata middleware.Metadata
}
```


### GetVPCSubnets
`DescribeSubnetsInput`
```bash
{
  #Check whether you have the required permissions for the action
  DryRun *bool => true or false (default false),
  # There are many FilterTypes (availability-zone, availability-zone-id, available-ip-address-count,cidr-block,subnet-arn,subnet-id,state, ...etc)
  Filters []types.Filter,
  # The maximum number of items to return
  MaxResults *int32,
  # The token returned from a previous paginated req
  NextToken *string,
  # One or more subnet ids (default all subnets)
  SubnetIds []string
}
```
`DescribeSubnetsOutput`
```bash
{
  NextToken *string,
  # Information about one or more subnets
  Subnets []types.Subnet,
  # Metadata about operation's result (req ID, ...etc)
  ResultMetadata middleware.Metadata
}
```

### GetEC2InstanceTypes
`DescribeInstanceTypesInput`
```bash
{
  # Check whether you have the required permissions for the action
  DryRun *bool => true or false (default false),
  # There are many FilterTypes (bare-metal, free-tier-eligible, instance-type(using '*'), ebs-info.*, instance-storage-info.*, network-info.* ...etc)
  Filters []types.Filter,
  # The maximum number of items to return
  MaxResults *int32,
  # The token returned from a previous paginated req
  NextToken *string,
  # One or more instance types (default all instance types)
  InstanceTypes []types.InstanceType
}
```
`DescribeInstanceTypesOutput`
```bash
{
  NextToken *string,
  # The instance Type
  InstanceTypes []types.InstanceTypeInfo,
  # Metadata about operation's result (req ID, ...etc)
  ResultMetadata middleware.Metadata
}
```
