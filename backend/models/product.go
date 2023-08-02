package models

import "time"

type Product struct {
	ID         uint       `json:"id" gorm:"primary_key"`
	Name       string     `json:"name"`
	Price      int        `json:"price"`
	SupplierID int        `json:"supplier_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}
