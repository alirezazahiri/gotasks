package taskservice

import "github.com/alirezazahiri/gotasks/internal/entity"

type TaskService struct {
	repo TaskRepository
}

func New(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task *entity.Task) error {
	_, err := s.repo.CreateTask(task)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) GetTask(id string) (*entity.Task, error) {
	task, err := s.repo.GetTask(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) UpdateTask(task *entity.Task) error {
	_, err := s.repo.UpdateTask(task)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) DeleteTask(id string) error {
	err := s.repo.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}
