package repository

import (
	"encoder/domain"
	"fmt"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type JobRepository interface {
	Insert(video *domain.Job) (*domain.Job, error)
	Find(id string) (*domain.Job, error)
	Update(job *domain.Job) (*domain.Job, error)
}

type JobRepositoryDb struct {
	Db *gorm.DB
}

func NewJobRepository(db *gorm.DB) *JobRepositoryDb {
	return &JobRepositoryDb{
		Db: db,
	}
}

func (this *JobRepositoryDb) Insert(job *domain.Job) (*domain.Job, error) {

	if job.ID == "" {
		job.ID = uuid.NewV4().String()
	}

	err := this.Db.Create(job).Error

	if err != nil {
		return nil, err
	}

	return job, nil
}

func (this *JobRepositoryDb) Find(id string) (*domain.Job, error) {
	var job domain.Job

	if id == "" {
		return nil, fmt.Errorf("id must not be null")
	}

	this.Db.Preload("Video").First(&job, "id = ?", id)

	if job.ID != "" {
		return &job, nil
	}

	return nil, fmt.Errorf("job with identifier %v not found", id)

}

func (this *JobRepositoryDb) Update(job *domain.Job) (*domain.Job, error) {

	err := this.Db.Save(&job).Error

	if err != nil {
		return nil, err
	}

	return job, nil
}
