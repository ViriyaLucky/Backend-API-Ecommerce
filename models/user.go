package models

import (
	"ecommerce-api/database"

	"github.com/novalagung/gubrak"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id            string `json:"Id" validate:"required"`
	Username      string `json:"Name"  validate:"required"`
	Email         string `json:"Email" validate:"required,email"`
	Password      string `json:"Password" validate:"required"`
	Created_Date  string `json:"Created_Date" validate:"required"`
	Modified_Date string `json:"Modified_Date" validate:"required"`
}

func AuthenticateUser(username string, password string) (bool, User) {
	defer database.Connector.Close()

	var user User
	database.Connector.Table("tbl_user").Where("username = ?", username).Find(&user)
	if gubrak.IsNil(user) {
		return false, user
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, user
	}

	return true, user
}

func GetUserById(id string) User {
	var user User
	database.Connector.Table("tbl_user").Where("id = ?", id).Find(&user)
	if gubrak.IsNil(user) {
		return user
	}

	return user
}
