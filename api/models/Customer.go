package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type CustomerModel struct {
	ID          uint        `gorm:"primary_key"`
	Environment environment `gorm:"type:ENUM('PRODUCTION', 'STAGING', 'QA', 'DEVELOPMENT');not null"`
	Email       string      `gorm:"column:email;not null;unique"`
	Mobile      string      `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type CustomerMetaModel struct {
	ID          uint          `gorm:"primary_key"`
	Customer    CustomerModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CustomerId  uint
	Value       datatypes.JSON
	Environment environment `gorm:"type:ENUM('PRODUCTION', 'STAGING', 'QA', 'DEVELOPMENT');not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
