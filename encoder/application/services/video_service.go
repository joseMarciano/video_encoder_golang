package services

import (
	"encoder/application/repository"
	"encoder/domain"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

const LOCAL_STORAGE_PATH = "local-path-enconder"
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
	directory := LOCAL_STORAGE_PATH + "/" + this.Video.ID
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		log.Fatalf("error on Mkdir %s", err.Error())
		return err
	}

	sess := GetS3Session()

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
	directory := LOCAL_STORAGE_PATH + "/" + this.Video.ID + "/" + this.Video.ID
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

func (this *VideoService) Encode() error {
	cmdArgs := []string{}
	directory := LOCAL_STORAGE_PATH + "/" + this.Video.ID
	cmdArgs = append(cmdArgs, directory+"/"+this.Video.ID+FRAG)
	cmdArgs = append(cmdArgs, "--use-segment-timeline")
	cmdArgs = append(cmdArgs, "-o")
	cmdArgs = append(cmdArgs, directory+"/out")
	cmdArgs = append(cmdArgs, "-f")
	cmdArgs = append(cmdArgs, "--exec-dir")
	cmdArgs = append(cmdArgs, "/opt/bento4/bin/")
	cmd := exec.Command("mp4dash", cmdArgs...)

	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("error on Encode | CombinedOutput %s", err.Error())
		return err
	}

	printOutput(output)

	return nil
}

func (this *VideoService) Finish() error {
	directory := LOCAL_STORAGE_PATH + "/" + this.Video.ID + "/" + this.Video.ID
	mp4 := directory + MP4
	frag := directory + FRAG

	err := os.Remove(mp4)
	if err != nil {
		log.Println("Error on remove ", this.Video.ID, ".mp4")
		return err
	}

	err = os.Remove(frag)
	if err != nil {
		log.Println("Error on remove ", this.Video.ID, ".frag")
		return err
	}

	log.Println("Files has been removed ", this.Video.ID)

	return nil
}

func printOutput(out []byte) {
	if len(out) > 0 {
		log.Printf("===> Output %s\n", string(out))
	}
}
