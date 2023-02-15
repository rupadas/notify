package models

import (
	"time"
)

type status string

const (
	ACTIVE   status = "ACTIVE"
	INACTIVE status = "INACTIVE"
)

type Event struct {
	ID          uint        `gorm:"primary_key"`
	Name        string      `gorm:"not null;unique"`
	Status      status      `gorm:"type:ENUM('ACTIVE', 'INACTIVE');not null;default:'ACTIVE'"`
	Environment Environment `gorm:"type:ENUM('PRODUCTION', 'STAGING', 'QA', 'DEVELOPMENT');not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Channels    []Channel `gorm:"many2many:event_channels;"`
}

type EventChannel struct {
	EventId     uint
	ChannelId   uint
	Environment Environment `gorm:"type:ENUM('PRODUCTION', 'STAGING', 'QA', 'DEVELOPMENT');not null"`
}
