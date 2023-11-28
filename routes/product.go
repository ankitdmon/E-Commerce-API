package routes
import (
	"net/http"
	"github.com/gin-gonic/gin"
)
func setupProductsRoutes(router *gin.Engine) {
	router.GET("/product", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Products Route",
		})
	})
}