package database

import (
	"errors"

	"github.com/HublastX/HubLast-Hub/internal/models"
	"gorm.io/gorm"
)

func SeedTechnologies(db *gorm.DB) {
	techs := []models.Technology{
		{Name: "React", Type: "frontend"},
		{Name: "Angular", Type: "frontend"},
		{Name: "Vue", Type: "frontend"},
		{Name: "Next.js", Type: "frontend"},
		{Name: "Tailwind", Type: "frontend"},
		{Name: "Node.js", Type: "backend"},
		{Name: "MySQL", Type: "backend"},
	}

	for _, tech := range techs {
		var existing models.Technology
		err := db.Where("name = ? AND type = ?", tech.Name, tech.Type).First(&existing).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&tech)
		}
	}
}
