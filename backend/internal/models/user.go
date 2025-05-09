package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username      string     `gorm:"uniqueIndex;not null"`
	Email         string     `gorm:"uniqueIndex;not null"`
	Password      string     `gorm:"not null"`
	Role          Role       `gorm:"type:varchar(10);default:'user'"`
	Projects      []Project  `gorm:"many2many:user_projects;"`
	OwnedProjects []Project  `gorm:"foreignKey:ResponsibleUserID"`
	Level         Level      `gorm:"type:varchar(10);default:'junior'"`
	Experience    Experience `gorm:"type:varchar(15);default:'basic'"`
	Employment    Employment `gorm:"type:varchar(15);default:'backend'"`
}
