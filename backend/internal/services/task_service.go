package services

import (
	"errors"

	"github.com/HublastX/HubLast-Hub/internal/models"
	"github.com/HublastX/HubLast-Hub/internal/repository"
)

type TaskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(taskRepo repository.TaskRepository) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) CreateTask(task *models.Task) error {
	if task.Title == "" {
		return errors.New("task title cannot be empty")
	}
	return s.taskRepo.Create(task)
}

func (s *TaskService) GetTask(id uint) (*models.Task, error) {
	return s.taskRepo.Find(id)
}

func (s *TaskService) UpdateTask(task *models.Task) error {
	if task.ID == 0 {
		return errors.New("task ID must be provided")
	}
	return s.taskRepo.Update(task)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.taskRepo.Delete(id)
}
