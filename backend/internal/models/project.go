package models

import (
	"time"

	"gorm.io/gorm"
)

type ProjectStatus string

const (
	Pending  ProjectStatus = "pending"
	Approved ProjectStatus = "approved"
	Rejected ProjectStatus = "rejected"
)

type Project struct {
	gorm.Model
	Title             string        `gorm:"not null"`
	Description       string        `gorm:"not null"`
	FrontendTech      string        `gorm:"not null"`
	BackendTech       string        `gorm:"not null"`
	EstimatedTime     int           `gorm:"not null"` // In days
	DeliveryDate      time.Time     `gorm:"not null"`
	Status            ProjectStatus `gorm:"type:varchar(20);default:'pending'"`
	ResponsibleUserID uint
	ResponsibleUser   User   `gorm:"foreignKey:ResponsibleUserID"`
	Users             []User `gorm:"many2many:user_projects;"`
}
