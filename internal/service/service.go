package service

import "todo/internal/repository"

type TaskService struct {
	repo *repository.InMemoryRepo
}

func NewTaskService(repo *repository.InMemoryRepo) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(t *repository.Task) *repository.Task {
	return s.repo.Create(t)
}

func (s *TaskService) GetTask(id int) (*repository.Task, bool) {
	return s.repo.Get(id)
}

func (s *TaskService) ListTasks() []*repository.Task {
	return s.repo.List()
}

func (s *TaskService) UpdateTask(id int, title *string, status *string) (*repository.Task, bool) {
	return s.repo.Update(id, title, status)
}

func (s *TaskService) DeleteTask(id int) bool {
	return s.repo.Delete(id)
}
