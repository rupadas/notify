package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Provider struct {
	ID          uint        `gorm:"primary_key"`
	Name        string      `gorm:"not null;unique"`
	AccessKey   string      `gorm:"not null"`
	AccessToken string      `gorm:"not null"`
	Environment Environment `gorm:"type:ENUM('PRODUCTION', 'STAGING', 'QA', 'DEVELOPMENT');not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type ChannelProviderRule struct {
	Channel     Channel
	ChannelId   uint
	Provider    Provider
	ProviderId  uint
	Country     string      `gorm:"not null"`
	Environment Environment `gorm:"type:ENUM('PRODUCTION', 'STAGING', 'QA', 'DEVELOPMENT');not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type ProviderSetting struct {
	Provider   Provider
	ProviderId uint
	Settings   json.RawMessage
}
