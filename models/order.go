package models

import (
	"ecommerce-api/database"
)

type Orders struct {
	Id           string  `json:"Id"`
	ProductId    string  `json:"ProductId"`
	Product      Product `json:"Product"`
	UserId       string  `json:"Description"`
	Quantity     string  `json:"Quantity"`
	Price        int     `json:"Price"`
	Created_Date string  `json:"Created_Date" validate:"required"`
}

func GetOrders(userId string) []Orders {
	defer database.Connector.Close()

	var result []Orders
	//check if product exist in user cart
	database.Connector.Table("tbl_orders").Where("user_id = ?", userId).Find(&result)
	for i, _ := range result {
		result[i].Product = GetProductById(result[i].ProductId)
	}
	return result
}
