package models

import "time"

type Order struct {
    ID           int       `json:"id"`
    CustomerName string    `json:"customer_name"`
    TotalAmount  float64   `json:"total_amount"`
    CreatedAt    time.Time `json:"created_at"`
}