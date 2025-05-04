package models

import (
	"gorm.io/gorm"
)

type Task struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ProjectID   uint   `json:"project_id"`
	Status      string `json:"status"`
}

func (t *Task) Create(db *gorm.DB) error {
	return db.Create(t).Error
}

func (t *Task) Update(db *gorm.DB) error {
	return db.Save(t).Error
}

func (t *Task) Delete(db *gorm.DB) error {
	return db.Delete(t).Error
}

func GetTaskByID(db *gorm.DB, id uint) (*Task, error) {
	var task Task
	err := db.First(&task, id).Error
	return &task, err
}

func GetAllTasks(db *gorm.DB) ([]Task, error) {
	var tasks []Task
	err := db.Find(&tasks).Error
	return tasks, err
}
