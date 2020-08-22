package auth

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/mihett05/auth-service-go/midlewares"
	"net/http"
)

func Verify(c *gin.Context)  {
	token, err := midlewares.AuthMiddleware().ParseToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "token not found",
		})
	} else {
		claims := jwt.ExtractClaimsFromToken(token)
		c.JSON(http.StatusOK, gin.H{
			"username": claims["username"],
		})
	}
}
