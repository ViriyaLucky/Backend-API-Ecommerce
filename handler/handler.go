package handler

import (
	"ecommerce-api/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	gubrak "github.com/novalagung/gubrak/v2"
)

var APPLICATION_NAME = "Backend API"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("the secret of kalimdor")

type M map[string]interface{}

func UserHandler(c echo.Context) error {
	// get the user id
	Id := getUserIdFromJwt(c)
	user := models.GetUserById(Id)
	if user.Username == "" {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "user does not exist",
		})
	}
	res := map[string]string{
		"username": user.Username,
		"email":    user.Email,
	}
	return c.JSON(http.StatusOK, res)
}

func Loginhandler(c echo.Context) error {
	if c.Request().Method != "POST" {
		return echo.ErrUnauthorized
	}

	username := c.FormValue("username")
	password := c.FormValue("password")
	if username == "" && password == "" {
		return c.JSON(http.StatusOK, echo.Map{
			"error": "Wrong username or password",
		})
	}

	ok, userInfo := models.AuthenticateUser(username, password)
	if !ok {
		return c.JSON(http.StatusOK, echo.Map{
			"error": "Wrong username or password",
		})
	}

	claims := models.MyClaims{
		Username: userInfo.Username,
		Email:    userInfo.Email,
		Id:       userInfo.Id,
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": signedToken,
	})
}

func ProductHandler(c echo.Context) error {

	// Add Product
	if c.Request().Method == "POST" {
		// Declare a new Person struct.
		var p models.Product

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(c.Request().Body).Decode(&p)
		if err != nil {
			return err
		}
		errModel := models.AddProduct(&p)
		if errModel != nil {
			return err
		} else {
			return c.JSON(http.StatusOK, p)
		}
	}

	// get the product id
	if c.Request().URL.Query().Has("id") {
		id := c.Request().URL.Query().Get("id")
		result := models.GetProductById(id)
		return c.JSON(http.StatusOK, result)
	} else {
		result := models.GetProducts()
		return c.JSON(http.StatusOK, result)
	}
}
func OrderHandler(c echo.Context) error {
	userId := getUserIdFromJwt(c)

	result := models.GetOrders(userId)
	return c.JSON(http.StatusOK, result)
}

func CartHandler(c echo.Context) error {
	userId := getUserIdFromJwt(c)

	result := models.GetCart(userId)
	return c.JSON(http.StatusOK, result)
}

func CheckoutHandler(c echo.Context) error {
	userId := getUserIdFromJwt(c)
	var products []int
	err := json.NewDecoder(c.Request().Body).Decode(&products)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	result, err := models.Checkout(products, userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if len(result) == 0 {
		return c.JSON(http.StatusBadGateway, err)
	}
	return c.JSON(http.StatusOK, result)
}

func UpdateCartHandler(c echo.Context) error {
	userId := getUserIdFromJwt(c)
	method := c.Request().Method
	itemIdParam := c.Param("itemid")
	itemId, errConv := strconv.Atoi(itemIdParam)

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	if method == "POST" {
		res, err := models.AddToCart(itemId, userId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Bad Request")
		}
		return c.JSON(http.StatusOK, res)
	} else {
		res, err := models.RemoveFromCart(itemId, userId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Bad Request")
		}
		return c.JSON(http.StatusOK, res)
	}
}

func ErrorHandler(c echo.Context) error {
	return echo.ErrUnauthorized
}

func registerUser(username, password string) (bool, M) {
	// hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if errHash != nil {
	// 	return false, nil
	// }
	basePath, _ := os.Getwd()
	dbPath := filepath.Join(basePath, "users.json")
	buf, _ := ioutil.ReadFile(dbPath)

	data := make([]M, 0)
	err := json.Unmarshal(buf, &data)
	if err != nil {
		return false, nil
	}

	res := gubrak.From(data).Find(func(each M) bool {
		return each["username"] == username && each["password"] == password
	}).Result()

	if res != nil {
		resM := res.(M)
		delete(resM, "password")
		return true, resM
	}

	return false, nil
}

// Helper

func getUserIdFromJwt(c echo.Context) int {
	// get the user id
	userClaim := c.Get("user").(*jwt.Token)
	claims := userClaim.Claims.(*models.MyClaims)
	Id := claims.Id
	return Id
}
