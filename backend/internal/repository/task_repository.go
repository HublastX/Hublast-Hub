package repository

import (
	"github.com/HublastX/HubLast-Hub/internal/models"
	"github.com/jinzhu/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *TaskRepository) Find(id uint) (*models.Task, error) {
	var task models.Task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) Update(task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *TaskRepository) Delete(id uint) error {
	return r.db.Delete(&models.Task{}, id).Error
}

func (r *TaskRepository) FindAll() ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
