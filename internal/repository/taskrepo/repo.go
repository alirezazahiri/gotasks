package taskrepo

import (
	"github.com/alirezazahiri/gotasks/internal/entity"
	"github.com/alirezazahiri/gotasks/internal/repository/postgresql"
)

type Repository struct {
	repo *postgresql.Repository
}

func New(repo *postgresql.Repository) *Repository {
	return &Repository{repo: repo}
}

func (r *Repository) CreateTask(task *entity.Task) (*entity.Task, error) {
	err := r.repo.DB.Create(task).Scan(&task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *Repository) GetTask(id string) (*entity.Task, error) {
	var task entity.Task
	err := r.repo.DB.Where("id = ?", id).First(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *Repository) UpdateTask(task *entity.Task) (*entity.Task, error) {
	err := r.repo.DB.Save(task).Scan(&task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *Repository) DeleteTask(id string) error {
	err := r.repo.DB.Where("id = ?", id).Delete(&entity.Task{}).Error
	if err != nil {
		return err
	}
	return nil
}
