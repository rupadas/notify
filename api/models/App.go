package models

import (
	"time"

	"gorm.io/gorm"
)

type environment string

const (
	PRODUCTION  environment = "PRODUCTION"
	STAGING     environment = "STAGING"
	QA          environment = "QA"
	DEVELOPMENT environment = "DEVELOPMENT"
)

type AppModel struct {
	ID          uint        `gorm:"primary_key"`
	Environment environment `gorm:"type:ENUM('PRODUCTION', 'STAGING', 'QA', 'DEVELOPMENT');not null"`
	Key         string      `gorm:"not null"`
	Token       string      `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
