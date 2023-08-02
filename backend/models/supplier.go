package models

import "time"

type Supplier struct {
	ID           uint       `json:"id" gorm:"primary_key"`
	Name         string     `json:"name"`
	MobileNumber *string    `json:"mobile_number"`
	Address      *string    `json:"address"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}
