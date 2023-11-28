package routes
import (
	"net/http"
	"github.com/gin-gonic/gin"
)
func setupCartRoutes(router *gin.Engine) {
	router.GET("/cart", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Cart Route",
		})
	})
}