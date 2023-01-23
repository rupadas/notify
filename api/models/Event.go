package models

import (
	"time"

	"gorm.io/gorm"
)

type status string

const (
	ACTIVE   status = "ACTIVE"
	INACTIVE status = "INACTIVE"
)

type Event struct {
	gorm.Model
	ID          uint        `gorm:"primary_key"`
	Name        string      `gorm:"not null;unique"`
	Status      status      `gorm:"type:ENUM('ACTIVE', 'INACTIVE');not null;default:'ACTIVE'"`
	Environment Environment `gorm:"type:ENUM('PRODUCTION', 'STAGING', 'QA', 'DEVELOPMENT');not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type EventChannel struct {
	Channel     Channel
	ChannelId   uint
	Event       Event
	EventId     uint
	Environment Environment `gorm:"type:ENUM('PRODUCTION', 'STAGING', 'QA', 'DEVELOPMENT');not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
