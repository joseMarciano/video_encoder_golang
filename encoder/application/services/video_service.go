package services

import (
	"encoder/application/repository"
	"encoder/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"os"
	"path"
)

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

	content, err := io.ReadAll(object.Body)

	if err != nil {
		return err
	}

	file, err := os.Create(path.Base("") + "/.local-path-enconder" + this.Video.ID + ".mp4")
	if err != nil {
		return err
	}

	_, err = file.Write(content)
	if err != nil {
		return err
	}

	defer file.Close()
	return nil

}
