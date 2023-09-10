package services

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type VideoUpload struct {
	Paths     []string
	VideoPath string
	//OutputBucket string
	Errors []string
}

func NewVideoUpload() *VideoUpload {
	return &VideoUpload{}
}

func (this *VideoUpload) UploadObject(
	objectPath string,
	session *session.Session,
	ctx context.Context,
) error {
	path := strings.Split(objectPath, LOCAL_STORAGE_PATH+"/")

	f, err := os.Open(objectPath)
	if err != nil {
		return err
	}
	defer f.Close()

	uploader := s3manager.NewUploader(session)

	_, err = uploader.UploadWithContext(ctx,
		&s3manager.UploadInput{
			Bucket: aws.String(ENCODER_BUCKET_NAME),
			//ACL:    aws.String("public-read"),
			Key:  aws.String("videoId-" + path[1]),
			Body: f})

	if err != nil {
		log.Println("Error on UploadObject ", err.Error())
		return err
	}

	log.Println("UploadObject successfully", objectPath)
	return nil
}

func (this *VideoUpload) LoadPaths() error {

	err := filepath.Walk(this.VideoPath, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() {
			this.Paths = append(this.Paths, path)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
