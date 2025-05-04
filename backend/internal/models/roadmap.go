package models

import (
	"gorm.io/gorm"
)

type RoadmapArea string

const (
	Frontend RoadmapArea = "frontend"
	Backend  RoadmapArea = "backend"
)

type DifficultyLevel string

const (
	Beginner     DifficultyLevel = "beginner"
	Intermediate DifficultyLevel = "intermediate"
	Advanced     DifficultyLevel = "advanced"
)

type Roadmap struct {
	gorm.Model
	Title       string          `gorm:"not null"`
	Area        RoadmapArea     `gorm:"type:varchar(20);not null"`
	Difficulty  DifficultyLevel `gorm:"type:varchar(20);not null"`
	Content     string          `gorm:"type:text;not null"`
	CourseLinks string          `gorm:"type:text"`
}
