package models

import "time"

type Product struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	Name         string `json:"string"`
	SerialNumber string `json:"serial_number"`
}
