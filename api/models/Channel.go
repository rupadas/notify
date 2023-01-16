package models

import (
	"time"

	"gorm.io/gorm"
)

type ChannelModel struct {
	ID          uint        `gorm:"primary_key"`
	Name        string      `gorm:"not null;unique"`
	Environment environment `gorm:"type:ENUM('PRODUCTION', 'STAGING', 'QA', 'DEVELOPMENT');not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
