package models

import "time"

type User struct {
	ID           uint       `json:"id" gorm:"primary_key"`
	Name         *string    `json:"name"`
	MobileNumber *string    `json:"mobile_number"`
	Email        string     `json:"email" gorm:"unique"`
	Password     string     `json:"password"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"-"`
}
