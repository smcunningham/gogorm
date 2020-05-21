package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// Order represents the model for an order
// Default table will be 'orders'
type Order struct {
	// gorm.Model
	OrderID uint `json:"orderId" gorm:"primary_key"`
	CustomerName string `json:"customerName"`
	OrderedAt time.Time `json:"orderedAt"`
	Items []Item `json:"items" gorm:"foreignkey:OrderID"`
}

type Item struct {
	// gorm.Model
	LineItemID uint `json:"lineItemId" gorm:"primary_key"`
	ItemCode string `json:"itemCode"`
	Description string `json:"description"`
	Quantity uint `json:"quantity"`
	OrderID uint `json:"-"`
}