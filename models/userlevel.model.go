package models

import (
	"time"

	"gorm.io/gorm"
)

type UserLevel struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	LevelName string `gorm:"type:varchar(20)"`

	User []User
}
