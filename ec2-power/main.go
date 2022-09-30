package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/smithy-go"
	"log"
)

var (
	region     string
	instanceId string
	action     string
)

func init() {
	flag.StringVar(&region, "r", "", "The Region of the AWS")
	flag.StringVar(&instanceId, "i", "", "The ID of the instance to stop")
	flag.StringVar(&action, "a", "", "The Action of the instance, start or stop")
	flag.Parse()
}

func main() {
	if instanceId == "" {
		fmt.Println("You must supply an instance ID (-i INSTANCE-ID")
		return
	}

	if action == "" {
		fmt.Println("You must supply an action (-a ACTION")
		return
	}

	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))
	client := ec2.NewFromConfig(cfg)

	switch action {
	case "start":
		input := &ec2.StartInstancesInput{
			InstanceIds: []string{instanceId},
			DryRun:      aws.Bool(true),
		}
		_, err = start(context.Background(), client, input)
	case "stop":
		input := &ec2.StopInstancesInput{
			InstanceIds: []string{instanceId},
			DryRun:      aws.Bool(true),
		}
		_, err = stop(context.Background(), client, input)
	default:
		log.Fatalf("unknow action: %s\r", action)
	}

	if err != nil {
		panic(err)
	}
}

func start(c context.Context, client *ec2.Client, input *ec2.StartInstancesInput) (*ec2.StartInstancesOutput, error) {
	resp, err := client.StartInstances(c, input)

	var apiErr smithy.APIError
	if errors.As(err, &apiErr) && apiErr.ErrorCode() == "DryRunOperation" {
		fmt.Println("User has permission to start an instance.")
		input.DryRun = aws.Bool(false)
		return client.StartInstances(c, input)
	}

	return resp, err
}

func stop(c context.Context, client *ec2.Client, input *ec2.StopInstancesInput) (*ec2.StopInstancesOutput, error) {
	resp, err := client.StopInstances(c, input)

	var apiErr smithy.APIError
	if errors.As(err, &apiErr) && apiErr.ErrorCode() == "DryRunOperation" {
		fmt.Println("User has permission to stop instances.")
		input.DryRun = aws.Bool(false)
		return client.StopInstances(c, input)
	}

	return resp, err
}
