package routes

import (
	"ecommerce-cart/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.SignUp())
	incomingRoutes.POST("users/login", controllers.Login())
	incomingRoutes.POST("admin/addproduct", controllers.ProductViewAdmin())
	incomingRoutes.POST("users/productview", controllers.ViewProduct())
	incomingRoutes.POST("users/search", controllers.SearchProductByQuery())
}
