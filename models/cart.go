package models

import (
	"ecommerce-api/database"
	"time"
)

type Cart struct {
	Id           string    `json:"Id"`
	ProductId    string    `json:"ProductId"`
	UserId       User      `json:"Description"`
	Quantity     string    `json:"Quantity"`
	Created_Date time.Time `json:"Created_Date" validate:"required"`
}

func AddToCart(productId string, userId string) string {
	defer database.Connector.Close()

	var result interface{}
	//check if product exist in user cart
	database.Connector.Table("tbl_wishlist").Where("user_id = ? && product_id = ", userId, productId).Find(&result)

	return ""
	// db.Transaction(func(tx *gorm.DB) error {
	// 	// do some database operations in the transaction (use 'tx' from this point, not 'db')
	// 	if err := tx.Create(&Animal{Name: "Giraffe"}).Error; err != nil {
	// 		// return any error will rollback
	// 		return err
	// 	}

	// 	if err := tx.Create(&Animal{Name: "Lion"}).Error; err != nil {
	// 		return err
	// 	}

	// 	// return nil will commit the whole transaction
	// 	return nil
	// })
}
