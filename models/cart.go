package models

import (
	"ecommerce-api/database"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type CartItem struct {
	Id           string  `json:"Id"`
	ProductId    int     `json:"ProductId"`
	Product      Product `json:"Product" gorm:"foreignKey:ProductId"`
	UserId       int     `json:"Description"`
	Quantity     int     `json:"Quantity"`
	Created_Date string  `json:"Created_Date" validate:"required"`
}

func (CartItem) TableName() string {
	return "tbl_wishlist"
}

func GetCart(userId int) []CartItem {
	var result []CartItem
	//check if product exist in user cart
	database.Connector.Preload("Product").Where("user_id = ?", userId).Find(&result)
	return result
}
func GetCartItemById(productId int, userId int) CartItem {
	var result CartItem
	//check if product exist in user cart
	database.Connector.Preload("Product").Where("user_id = ? AND product_id = ?", userId, productId).Find(&result)
	return result
}
func AddToCart(productId int, userId int) ([]CartItem, error) {
	var result CartItem
	//check if product exist in user cart
	database.Connector.Where("user_id = ? AND product_id = ?", userId, productId).First(&result)
	if result.Id != "" {
		//increase quantity
		result.Quantity++
		err := database.Connector.Table("tbl_wishlist").Where("user_id = ? AND product_id = ?", userId, productId).Update(&result).Error
		if err != nil {
			return nil, err
		}
		var res []CartItem
		err = database.Connector.Preload("Product").Where("user_id = ?", userId).Find(&res).Error
		if err != nil {
			return nil, err
		}
		return res, nil
	} else {
		//increase quantity
		result.ProductId = productId
		result.Quantity = 1
		result.UserId = userId
		result.Created_Date = time.Now().String()
		err := database.Connector.Table("tbl_wishlist").Where("user_id = ? AND product_id = ?", userId, productId).Create(&result).Error
		if err != nil {
			return nil, err
		}
		var res []CartItem
		database.Connector.Preload("Product").Where("user_id = ?", userId).Find(&res)
		return res, nil
	}
}
func RemoveFromCart(productId int, userId int) ([]CartItem, error) {
	var result CartItem
	//check if product exist in user cart
	database.Connector.Where("user_id = ? AND product_id = ?", userId, productId).First(&result)
	if result.Id != "" {
		if result.Quantity == 1 {
			//remove item
			err := database.Connector.Where("user_id = ? AND product_id = ?", userId, productId).Delete(&result).Error
			if err != nil {
				return nil, err
			}
		} else {
			//decrease quantity
			result.Quantity--
			err := database.Connector.Table("tbl_wishlist").Where("user_id = ? AND product_id = ?", userId, productId).Update(&result).Error
			if err != nil {
				return nil, err
			}
		}

		var res []CartItem
		database.Connector.Preload("Product").Where("user_id = ?", userId).Find(&res)
		return res, nil
	} else {
		var res []CartItem
		database.Connector.Preload("Product").Where("user_id = ?", userId).Find(&res)
		return res, nil
	}
}

func Checkout(products []int, userId int) ([]Orders, error) {
	var orders []Orders
	err := database.Connector.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		for _, v := range products {
			var order Orders

			var ct CartItem = GetCartItemById(v, userId)
			if ct.Id == "" {
				return errors.New("No Such Item in Cart")
			}
			order.ProductId = v
			order.UserId = userId
			order.Quantity = ct.Quantity
			order.Price = int(ct.Product.Price)
			order.Created_Date = time.Now().String()
			if err := tx.Create(&order).Error; err != nil {
				return err
			}

			// Decrease Product Qty
			var product Product
			if err := tx.Where("id = ?", v).Find(&product).Error; err != nil {
				return err
			}

			res := (product.Quantity - ct.Quantity)
			if res < 0 {
				return errors.New("minus")
			}

			product.Quantity = res
			//update product qty
			if err := tx.Table("tbl_products").Where("id = ?", v).Update(&product).Error; err != nil {
				return err
			}

			orders = append(orders, order)
		}

		for _, v := range orders {
			err := database.Connector.Where("user_id = ? AND product_id = ?", userId, v.ProductId).Delete(CartItem{}).Error
			if err != nil {
				return err
			}
		}
		// return nil will commit the whole transaction
		return nil
	})

	return orders, err
}
