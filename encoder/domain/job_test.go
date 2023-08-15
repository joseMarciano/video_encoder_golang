package domain_test

import (
	"encoder/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewJob(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "filePath"
	video.CreatedAt = time.Now()
	job, err := domain.NewJob("path", "CONVERTED", video)

	require.NotNil(t, job)
	require.Nil(t, err)
}
