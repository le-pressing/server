package models

import "time"

type Customer struct {
	ID             int       `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	BuildingNumber int       `json:"building_number"`
	StreetName     string    `json:"street_name"`
	City           string    `json:"city"`
	State          string    `json:"State"`
	ZipCode        int       `json:"zip_code"`
	Orders         []Order   `json:"orders"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Order struct {
	ID         int         `json:"id"`
	CustomerID int         `json:"customer_id"`
	Items      []OrderItem `json:"items"`
	Status     string      `json:"status"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

type OrderItem struct {
	Type     string  `json:"type"`
	Price    float64 `json:"price"`
	Qty      int     `json:"qty"`
	UpCharge float64 `json:"up_charge"`
}
