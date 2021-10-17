package models

import (
	"ecommerce-api/database"
)

type Product struct {
	Id            string `json:"Id" gorm:"foreignKey:ProductId"`
	ProductName   string `json:"ProductName"`
	Description   string `json:"Description"`
	Price         string `json:"Price"`
	Image_Path    string `json:"Image_Path"`
	Quantity      string `json:"Quantity"`
	Category      string `json:"Category"`
	Created_Date  string `json:"Created_Date" validate:"required"`
	Modified_Date string `json:"Modified_Date" validate:"required"`
	IsActive      bool   `json:"isActive"`
}

func GetProducts() []Product {
	var result []Product
	database.Connector.Table("tbl_products").Find(&result)
	return result
}

func GetProductById(productId string) Product {
	var result Product
	database.Connector.Table("tbl_products").Where("id = ? ", productId).First(&result)
	return result
}
