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

func TestNewVideoRepositoryDBInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "resource_id"
	video.FilePath = "file_path"
	video.CreatedAt = time.Now()

	repo := repository.VideoRepositoryDb{db}
	repo.Insert(video)

	v, err := repo.Find(video.ID)

	require.NotEmpty(t, v.ID)
	require.Nil(t, err)
	require.Equal(t, v.ID, video.ID)

}
