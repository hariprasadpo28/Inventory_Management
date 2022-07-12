package Routes

import (
	"Inventory_Management/Controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/api")
	{
		grp1.GET("product", Controller.GetProducts)
		grp1.POST("product", Controller.AddProduct)
		grp1.POST("product/:uid", Controller.UpdateProduct)
		grp1.POST("order/", Controller.PlaceOrder)
		grp1.POST("user", Controller.CreateUser)
		grp1.GET("user", Controller.GetUser)

		//grp1.PUT("student/:id", Controller.UpdateUser)
		//grp1.DELETE("student/:id", Controller.DeleteUser)
		//grp1.PATCH("student/:id", Controller.PatchUser)
	}
	return r
}
