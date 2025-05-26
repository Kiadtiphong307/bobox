package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    Username  string         `gorm:"unique;not null" json:"username"`
    Email     string         `gorm:"unique;not null" json:"email"`
    Password  string         `gorm:"not null" json:"-"`
    Role      string         `gorm:"type:enum('user','admin');default:'user'" json:"role"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}


