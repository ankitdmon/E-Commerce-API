package routes
import (
	"net/http"
	"github.com/gin-gonic/gin"
)
func SetupRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PONG",
		})
	})
	setupCartRoutes(router)
	setupOrderRoutes(router)
	setupUserRoutes(router)
	setupWishlistRoutes(router)
	setupProductsRoutes(router)
	setupPublicRoutes(router)
}