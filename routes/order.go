package routes
import (
	"net/http"
	"github.com/gin-gonic/gin"
)
func setupOrderRoutes(router *gin.Engine) {
	router.GET("/order", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Cart Route",
		})
	})
}