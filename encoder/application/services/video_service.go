package services

import (
	"encoder/application/repository"
	"encoder/domain"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

const localStoragePath = "local-path-enconder"
const MP4 = ".mp4"
const FRAG = ".frag"

type VideoService struct {
	Video           *domain.Video
	VideoRepository repository.VideoRepository
}

func NewVideoService(video *domain.Video) VideoService {
	return VideoService{
		Video: video,
	}
}

func (this *VideoService) Download(bucketName string) error {
	directory := localStoragePath + "/" + this.Video.ID
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		log.Fatalf("error on Mkdir %s", err.Error())
		return err
	}

	get, _ := credentials.NewEnvCredentials().Get()
	fmt.Printf("Credentilas %s - %s ", get.AccessKeyID, get.SecretAccessKey)
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewEnvCredentials(),
	})

	if err != nil {
		return err
	}

	// Create S3 service client
	svc := s3.New(sess)
	object, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &this.Video.FilePath,
	})

	if err != nil {
		return err
	}

	content, err := ioutil.ReadAll(object.Body)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(directory+"/"+this.Video.ID+".mp4", content, os.ModePerm)
	if err != nil {
		return err
	}

	return nil

}

func (this *VideoService) Fragment() error {
	directory := localStoragePath + "/" + this.Video.ID + "/" + this.Video.ID
	source := directory + MP4
	target := directory + FRAG

	cmd := exec.Command("mp4fragment", source, target)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("error on CombinedOutput %s", err.Error())
		return err
	}

	printOutput(output)

	return nil
}

func printOutput(out []byte) {
	if len(out) > 0 {
		log.Printf("===> Output %s\n", string(out))
	}
}
