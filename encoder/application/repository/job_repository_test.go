package repository_test

import (
	"encoder/application/repository"
	"encoder/domain"
	"encoder/framework/database"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewJobRepositoryDBInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "resource_id"
	video.FilePath = "file_path"
	video.CreatedAt = time.Now()

	videoRepository := repository.VideoRepositoryDb{db}
	videoRepository.Insert(video)

	job, err := domain.NewJob(
		"output_path",
		"Pending",
		video,
	)

	require.Nil(t, err)

	jobRepository := repository.JobRepositoryDb{Db: db}
	jobRepository.Insert(job)

	j, err := jobRepository.Find(job.ID)

	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, j.Video.ID, video.ID)

}

func TestNewJobRepositoryDBUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "resource_id"
	video.FilePath = "file_path"
	video.CreatedAt = time.Now()

	videoRepository := repository.VideoRepositoryDb{db}
	videoRepository.Insert(video)

	job, err := domain.NewJob(
		"output_path",
		"Pending",
		video,
	)

	require.Nil(t, err)

	job.Status = "Completed"

	jobRepository := repository.JobRepositoryDb{Db: db}
	jobRepository.Update(job)

	j, err := jobRepository.Find(job.ID)

	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, j.Status, job.Status)
	require.Equal(t, j.Video.ID, video.ID)

}
