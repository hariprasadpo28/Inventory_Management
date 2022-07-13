package Controller

import (
	"Inventory_Management/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddProduct - Add a new product to the database
func AddProduct(c *gin.Context) {
	var product Models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please add valid product details"})
		return
		//c.AbortWithStatus(http.StatusNotFound)
	} else {
		if err := Models.AddProduct(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
			//c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, product)
		}
	}
}

//UpdateProduct - Update existing product details
func UpdateProduct(c *gin.Context) {
	var product Models.Product
	uid := c.Params.ByName("uid")
	err := Models.GetProductByID(&product, uid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found!"})
		return
		//c.AbortWithStatus(http.StatusNotFound)
	}
	c.BindJSON(&product)
	err = Models.UpdateProduct(&product, uid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

//PlaceOrder - Place order function takes product ID, quantity, username to place a new order
func PlaceOrder(c *gin.Context) {
	var order Models.Order
	c.BindJSON(&order)
	//c.Request.Header
	//c.Request.Context()
	//c.GetHeader("username")

	err := Models.PlaceOrder(&order)
	if err != nil {
		//fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		//c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

//GetProducts - Get all products available in the database
func GetProducts(c *gin.Context) {
	var products []Models.Product
	err := Models.GetProducts(&products)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, products)
	}
}

//CreateUser - Creating user takes username, name, email data (all are required)
func CreateUser(c *gin.Context) {
	var user Models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Please provide valid details"})
		c.AbortWithStatus(http.StatusNotFound)
	}

	if err := Models.CreateUser(&user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Try again with different username"})
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}

}

//GetUser- Get all users from the database (should be accessed by admin only; need to implement role based authentication)
func GetAllUsers(c *gin.Context) {
	var users []Models.User
	err := Models.GetAllUsers(&users)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

// GetUserOrders - Get all orders of the user (username need to be taken from the logged in user; will do that after implementing the authentication part)
func GetUserOrders(c *gin.Context) {
	var orders []Models.Order
	username := c.Params.ByName("username")
	err := Models.GetUserOrders(&orders, username)
	if err != nil {
		//fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, orders)
	}
}

//GetAllOrders - Get all orders placed
func GetAllOrders(c *gin.Context) {
	var orders []Models.Order
	err := Models.GetAllOrders(&orders)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, orders)
	}
}
