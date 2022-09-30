SHELL := /bin/bash

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BUILD_FLAG=CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags "-w -s"

.PHONY: all ec2-public-ip ec2-power
default: all
all: ec2-public-ip ec2-power

ec2-power:
	$(BUILD_FLAG) -o aws-ec2-power ec2-power/main.go

ec2-public-ip:
	$(BUILD_FLAG) -o aws-ec2-public-ip ec2-public-ip/main.go