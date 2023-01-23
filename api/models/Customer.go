package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Customer struct {
	ID          uint        `gorm:"primary_key"`
	Environment Environment `gorm:"type:ENUM('PRODUCTION', 'STAGING', 'QA', 'DEVELOPMENT');not null"`
	Email       string      `gorm:"column:email;not null;unique"`
	Mobile      string      `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type CustomerMeta struct {
	ID          uint     `gorm:"primary_key"`
	Customer    Customer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CustomerId  uint
	Value       datatypes.JSON
	Environment Environment `gorm:"type:ENUM('PRODUCTION', 'STAGING', 'QA', 'DEVELOPMENT');not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
