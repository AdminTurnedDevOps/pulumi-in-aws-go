package main

import (
	"fmt"
	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"log"
)

func main() {
	bucketName := "mjlpulumitest92"
	createS3Bucket(bucketName)
}

func createS3Bucket(bucketName string) {
	if bucketName == "" {
		log.Panic("No bucket name specific. Please try again...")
	}

	pulumi.Run(func(ctx *pulumi.Context) error {
		fmt.Println("Creating S3 Bucket...")

		bucket, err := s3.NewBucket(ctx, bucketName, nil)
		if err != nil {
			fmt.Printf("S3 bucket not created. An error occurred")
			log.Println(err)
		} else {
			fmt.Printf("New S3 Bucket Created: %s", bucket)
		}

		return nil
	})

	projectInfo := pulumi.RunInfo{}

	projectInfo.Project = "aws-go-s3"
	projectInfo.Stack = "Development"
}
