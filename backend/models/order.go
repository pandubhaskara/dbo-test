package models

import "time"

type Order struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	UserID    int        `json:"user_id"`
	ProductID int        `json:"product_id"`
	Quantity  int        `json:"quantity"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
