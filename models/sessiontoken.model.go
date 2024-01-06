package models

import (
	"time"
)

type SessionToken struct {
	Id        int       `gorm:"primaryKey"`
	Token     string    `gorm:"varchar(300)"`
	ExpiresAt time.Time `gorm:"timestamp"`
	UserID    int       `gorm:"index"`
	User      User      `gorm:"foreignKey:UserID"`
}
