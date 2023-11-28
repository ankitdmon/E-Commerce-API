package routes
import (
	"net/http"
	"github.com/gin-gonic/gin"
)
func setupUserRoutes(router *gin.Engine) {
	router.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "User Route",
		})
	})
}