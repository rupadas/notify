package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Provider struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"not null;unique"`
	Type        string
	Environment Environment `gorm:"type:ENUM('PRODUCTION', 'STAGING', 'QA', 'DEVELOPMENT');not null"`
	App         App         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AppId       uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type ChannelProvider struct {
	Channel     Channel
	ChannelId   uint
	Provider    Provider
	ProviderId  uint
	Environment Environment `gorm:"type:ENUM('PRODUCTION', 'STAGING', 'QA', 'DEVELOPMENT');not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type ProviderSetting struct {
	ID          uint `gorm:"primary_key"`
	Provider    Provider
	ProviderId  uint
	Country     string      `gorm:"not null"`
	Environment Environment `gorm:"type:ENUM('PRODUCTION', 'STAGING', 'QA', 'DEVELOPMENT');not null"`
	Settings    json.RawMessage
}
