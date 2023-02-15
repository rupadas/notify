package models

import (
	"time"

	"gorm.io/gorm"
)

type Channel struct {
	ID          uint        `gorm:"primary_key"`
	Name        string      `gorm:"not null;unique"`
	Environment Environment `gorm:"type:ENUM('PRODUCTION', 'STAGING', 'QA', 'DEVELOPMENT');not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Events      []Event        `gorm:"many2many:event_channels;"`
}
