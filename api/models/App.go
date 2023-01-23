package models

import (
	"time"

	"gorm.io/gorm"
)

type Environment string

const (
	PRODUCTION  Environment = "PRODUCTION"
	STAGING     Environment = "STAGING"
	QA          Environment = "QA"
	DEVELOPMENT Environment = "DEVELOPMENT"
)

type App struct {
	ID          uint        `gorm:"primary_key"`
	Environment Environment `gorm:"type:ENUM('PRODUCTION', 'STAGING', 'QA', 'DEVELOPMENT');not null"`
	AccessKey   string      `gorm:"not null"`
	AccessToken string      `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
