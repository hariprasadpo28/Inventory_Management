package Models

import (
	"time"
)

type Product struct {
	Id          int     `json:"id"`
	UniqueId    string  `json:"unique_id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Quantity    int     `json:"quantity"`
	Description string  `json:"description"`
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	UserName string `json:"user_name" binding:"required" gorm:"unique"`
	Email    string `json:"email" binding:"required"`
	//Orders   []Order `json:"orders"`
}

type Order struct {
	Id          int     `json:"id"`
	ProductId   string  `json:"product_id" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	TotalAmount float32 `json:"total_amount"`
	UserName    string  `json:"user_name"`
	Status      string  `json:"status"`
	OrderTime time.Time `json:"order_time" sql:"DEFAULT:CURRENT_TIMESTAMP"`
}

func (b *User) TableName() string {
	return "user"
}

func (b *Product) TableName() string {
	return "product"
}

func (b *Order) TableName() string {
	return "order"
}
