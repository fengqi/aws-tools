package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var (
	region string
)

func init() {
	flag.StringVar(&region, "r", "", "The Region of the AWS")
	flag.Parse()
}

func main() {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))
	client := ec2.NewFromConfig(cfg)

	input := &ec2.DescribeInstancesInput{}
	output, err := client.DescribeInstances(context.Background(), input)
	if err != nil {
		panic(err)
	}

	for _, item := range output.Reservations {
		for _, instance := range item.Instances {
			fmt.Printf("instance: %s, ec2-public-ip: %s\n",
				*instance.InstanceId,
				*instance.NetworkInterfaces[0].Association.PublicIp,
			)
		}
	}
}
