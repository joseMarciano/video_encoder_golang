package services

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
)

const ENCODER_BUCKET_NAME = "admin-catalog"

func GetS3Session() *session.Session {
	get, _ := credentials.NewEnvCredentials().Get()
	fmt.Printf("Credentilas %s - %s ", get.AccessKeyID, get.SecretAccessKey)
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewEnvCredentials(),
	})

	if err != nil {
		log.Fatalf("error on get S3 session %s", err.Error())
	}

	return sess
}
