package routes
import (
	"net/http"
	"github.com/gin-gonic/gin"
)
func setupPublicRoutes(router *gin.Engine) {
	router.GET("/public", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Cart Route",
		})
	})
}