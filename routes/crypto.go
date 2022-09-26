package routes

import (
	"github.com/caiomp87/crypto-votes-api/controllers"
	"github.com/gin-gonic/gin"
)

func AddCryptoRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		crypto := v1.Group("/crypto")
		{
			crypto.GET("/", controllers.ListCryptos)
			crypto.GET("/:id", controllers.GetCrypto)
			crypto.POST("/", controllers.CreateCrypto)
			crypto.PATCH("/:id", controllers.UpdateCrypto)
			crypto.DELETE("/:id", controllers.DeleteCrypto)
		}
	}
}
