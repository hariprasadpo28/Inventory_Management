package Controller

import (
	"Inventory_Management/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PlaceOrder(c *gin.Context) {
	//var user Models.User
	var order Models.Order
	//var prod Models.Product
	id := c.Params.ByName("id")
	prodID := c.Params.ByName("prod_id")
	//if err := Models.GetProductByID(&prod, prodID); err != nil{
	//	c.AbortWithStatus(http.StatusNotFound)
	//}
	//err := Models.GetUserByID(&user, id)
	//if err!= nil{
	//	c.JSON(http.StatusNotFound, user)
	//	c.AbortWithStatus(http.StatusNotFound)
	//}else {
	c.BindJSON(&order)
	err := Models.PlaceOrder(&order, id, prodID)
	if err != nil {
		//fmt.Println(err.Error())
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
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		if err := Models.AddProduct(&product); err != nil {
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
