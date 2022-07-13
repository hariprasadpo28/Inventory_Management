package Routes

import (
	"Inventory_Management/Controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/product")
	{
		grp1.GET("", Controller.GetProducts)
		grp1.POST("", Controller.AddProduct)
		grp1.POST("/:uid", Controller.UpdateProduct)
	}
	grp2 := r.Group("/order")
	{
		grp2.POST("/", Controller.PlaceOrder)
		grp2.GET("/:username", Controller.GetUserOrders)
	}
	grp3 := r.Group("/user")
	{
		grp3.POST("", Controller.CreateUser)
		grp3.GET("", Controller.GetAllUsers)
	}
	grp4 := r.Group("/retailer")
	{
		grp4.POST("", Controller.CreateRetailer)
		grp4.GET("/orders/:retailerid", Controller.GetRetailerOrders)
	}

	return r
}
