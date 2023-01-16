package models

import (
	"time"

	"gorm.io/gorm"
)

type ProviderModel struct {
	ID          uint        `gorm:"primary_key"`
	Name        string      `gorm:"not null;unique"`
	Key         string      `gorm:"not null"`
	Token       string      `gorm:"not null"`
	Environment environment `gorm:"type:ENUM('PRODUCTION', 'STAGING', 'QA', 'DEVELOPMENT');not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type ChannelProviderRuleModel struct {
	Channel     ChannelModel
	ChannelId   uint
	Provider    ProviderModel
	ProviderId  uint
	Country     string      `gorm:"not null"`
	Environment environment `gorm:"type:ENUM('PRODUCTION', 'STAGING', 'QA', 'DEVELOPMENT');not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
