package taskservice

import "github.com/alirezazahiri/gotasks/internal/entity"

type TaskRepository interface {
	CreateTask(task *entity.Task) (*entity.Task, error)
	GetTask(id string) (*entity.Task, error)
	UpdateTask(task *entity.Task) (*entity.Task, error)
	DeleteTask(id string) error
}
