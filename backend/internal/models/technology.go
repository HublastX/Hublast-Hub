package models

import "gorm.io/gorm"

type Technology struct {
	gorm.Model
	Name     string    `gorm:"not null;uniqueIndex"`
	Type     TechType  `gorm:"type:varchar(10);not null"`
	Projects []Project `gorm:"many2many:project_technologies;"`
}
