package Controller

import (
	"Inventory_Management/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PlaceOrder(c *gin.Context) {
	var order Models.Order
	c.BindJSON(&order)

	err := Models.PlaceOrder(&order)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		//user.Orders = append(user.Orders, order)
		//Config.DB.Model(&order).Update(Models.Order{UserId: user, TotalAmount: c.Params.ByName("quantity")*c.Params.ByName("amount")})
		//Config.DB.Model(&user).Update("Orders", user.Orders)
		c.JSON(http.StatusOK, order)
	}
}

func AddProduct(c *gin.Context) {
	var product Models.Product
	if err1 := c.ShouldBindJSON(&product); err1 != nil {
		fmt.Println(err1)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		if err := Models.AddProduct(&product); err != nil {
			fmt.Println(err1)
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, product)
		}
	}
}

func GetProducts(c *gin.Context) {
	var products []Models.Product
	err := Models.GetProducts(&products)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, products)
	}
}

func CreateUser(c *gin.Context) {
	var user Models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if err := Models.CreateUser(&user); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}

}

func GetUser(c *gin.Context) {
	var users []Models.User
	err := Models.GetUser(&users)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, users)
	}

}

func UpdateProduct(c *gin.Context) {
	var product Models.Product
	uid := c.Params.ByName("uid")
	err := Models.GetProductByID(&product, uid)
	if err != nil {
		c.JSON(http.StatusNotFound, product)
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.BindJSON(&product)
	err = Models.UpdateProduct(&product, uid)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}
