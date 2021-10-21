package models

import (
	"ecommerce-api/database"
	"errors"
	"time"
)

type Product struct {
	Id            int     `json:"Id" gorm:"primaryKey"`
	ProductName   string  `json:"ProductName"`
	Description   string  `json:"Description"`
	Price         float32 `json:"Price"`
	Image_Path    string  `json:"Image_Path"`
	Quantity      int     `json:"Quantity"`
	Category      string  `json:"Category"`
	Created_Date  string  `json:"Created_Date" validate:"required"`
	Modified_Date string  `json:"Modified_Date" validate:"required"`
	IsActive      int     `json:"isActive" gorm:"column:isActive"`
}

func (Product) TableName() string {
	return "tbl_products"
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
func AddProduct(product *Product) error {
	product.Created_Date = time.Now().String()
	product.Modified_Date = time.Now().String()
	product.IsActive = 1
	res := database.Connector.Create(&product)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}

func DecreaseProductQty(productId int, size int) error {
	var product Product
	err := database.Connector.Where("id = ?", productId).Find(&product).Error
	if err != nil {
		return err
	}

	res := (product.Quantity - size)
	if res < 0 {
		return errors.New("minus")
	}

	product.Quantity = res

	err = database.Connector.Where("id = ?", productId).Update(&product).Error
	if err != nil {
		return err
	}
	return nil
}
