package domain_test

import (
	"encoder/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestIdIsNotAnUuid(t *testing.T) {
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.ResourceID = "resource_id"
	video.FilePath = "file_path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Nil(t, err)
}
