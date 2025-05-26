package models

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"unique"`
	Category    string
	Description string
	Price       float64
	Stock       int
	CreatedAt   time.Time
}

