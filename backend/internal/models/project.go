package models

import (
	"time"

	"gorm.io/gorm"
)

var (
	FrontendTechs = []string{"React", "Angular", "Vue", "Next.js", "Svelte", "HTML/CSS", "Bootstrap", "Tailwind"}

	BackendTechs = []string{"Go", "Python", "Node.js", "Java", "C#", "PHP", "Ruby", "Django", "Express", "Spring", "Laravel"}
)

type Project struct {
	gorm.Model
	Title             string        `gorm:"not null"`
	Description       string        `gorm:"not null"`
	Technologies      []Technology  `gorm:"many2many:project_technologies;"`
	EstimatedTime     int           `gorm:"not null"`
	DeliveryDate      time.Time     `gorm:"not null"`
	Status            ProjectStatus `gorm:"type:varchar(20);default:'pending'"`
	ResponsibleUserID uint
	ResponsibleUser   User         `gorm:"foreignKey:ResponsibleUserID"`
	Users             []User       `gorm:"many2many:user_projects;"`
	QuantyMaxUsers    int          `gorm:"not null"`
	Level             ProjectLevel `gorm:"type:varchar(10);default:'low'"`
}
