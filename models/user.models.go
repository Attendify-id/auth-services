package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id           int             `gorm:"primaryKey"`
	Username     string          `gorm:"type:varchar(16)"`
	Email        string          `gorm:"type:varchar(50)"`
	Password     string          `gorm:"type:varchar(255)" json:"-"`
	Fullname     string          `gorm:"type:varchar(40)"`
	Registered   *time.Time      `gorm:"timestamp" json:"-"`
	LastLogin    *time.Time      `gorm:"timestamp" json:"-"`
	IpAddress    *string         `gorm:"type:varchar(15)" json:"-"`
	Browser      *string         `gorm:"type:varchar(255)" json:"-"`
	CreatedAt    time.Time       `gorm:"timestamp" json:"-"`
	UserLevel    UserLevel       `gorm:"foreignKey:UserLevelID"`
	UpdatedAt    time.Time       `gorm:"timestamp" json:"-"`
	DeletedAt    *gorm.DeletedAt `gorm:"index" json:"-"`
	UserLevelID  int             ` json:"-"`
	SessionToken []SessionToken  ` json:"-"`
}
