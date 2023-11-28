package routes
import (
	"net/http"
	"github.com/gin-gonic/gin"
)
func setupWishlistRoutes(router *gin.Engine) {
	router.GET("/wishlist", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Wishlist Route",
		})
	})
}