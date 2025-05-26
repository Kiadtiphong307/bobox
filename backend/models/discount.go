package models

import "time"

type Discount struct {
	ID         uint      `gorm:"primaryKey"`
	Code       string    `gorm:"uniqueIndex"`
	Percentage int       `gorm:"not null"`  // ลด 10% = 10
	UsageLimit int       `gorm:"not null"`  // จำกัดจำนวนครั้ง
	UsedCount  int       `gorm:"default:0"` // ใช้ไปแล้ว
	ExpiresAt  time.Time `gorm:"not null"`  // วันหมดอายุ
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
