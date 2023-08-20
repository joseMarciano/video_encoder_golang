package repository

import (
	"encoder/domain"
	"fmt"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type VideoRepository interface {
	Insert(video *domain.Video) (*domain.Video, error)
	Find(id string) (*domain.Video, error)
}

type VideoRepositoryDb struct {
	Db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *VideoRepositoryDb {
	return &VideoRepositoryDb{
		Db: db,
	}
}

func (this *VideoRepositoryDb) Insert(video *domain.Video) (*domain.Video, error) {

	if video.ID == "" {
		video.ID = uuid.NewV4().String()
	}

	err := this.Db.Create(video).Error

	if err != nil {
		return nil, err
	}

	return video, nil
}

func (this *VideoRepositoryDb) Find(id string) (*domain.Video, error) {
	var video domain.Video

	if id == "" {
		return nil, fmt.Errorf("id must not be null")
	}

	this.Db.Preload("Jobs").First(&video, "id = ?", id)

	if video.ID != "" {
		return &video, nil
	}

	return nil, fmt.Errorf("video with identifier %v not found", id)

}
