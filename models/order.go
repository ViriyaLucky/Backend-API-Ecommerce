package models

import (
	"ecommerce-api/database"
)

type Orders struct {
	Id           int     `json:"Id" gorm:"primary_key"`
	ProductId    int     `json:"ProductId"`
	Product      Product `json:"Product" gorm:"foreignKey:ProductId"`
	UserId       int     `json:"Description"`
	Quantity     int     `json:"Quantity"`
	Price        int     `json:"Price"`
	Created_Date string  `json:"Created_Date" validate:"required"`
}

func (Orders) TableName() string {
	return "tbl_orders"
}
func GetOrders(userId int) []Orders {
	var result []Orders
	//check if product exist in user cart
	database.Connector.Preload("Product").Where("user_id = ?", userId).Find(&result)
	return result
}
