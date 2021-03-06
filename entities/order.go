package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID          uint
	InvoiceID   string `gorm:"default:NULL" json:"invoice_id" form:"invoice_id"`
	PaymentLink string `gorm:"default:NULL" json:"payment_id" form:"payment_id"`
	UserID      uint   `json:"user_id" form:"user_id"`
	Total       int    `json:"total" form:"total"`
	Status      string `json:"status" form:"status"`
}
