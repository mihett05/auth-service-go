package midlewares

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/mihett05/auth-service-go/core/forms"
	"github.com/mihett05/auth-service-go/core/models"
	"github.com/mihett05/auth-service-go/libs"
	"github.com/mihett05/auth-service-go/services"
)

var _authMiddleware *jwt.GinJWTMiddleware = nil

func AuthMiddleware() *jwt.GinJWTMiddleware {
	if _authMiddleware == nil {
		authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
			Key:         []byte(libs.EnvDefault("KEY")),
			IdentityKey: "username",
			PayloadFunc: func(data interface{}) jwt.MapClaims {
				if v, ok := data.(*models.User); ok {
					return jwt.MapClaims{
						"username": v.Username,
					}
				}
				return jwt.MapClaims{}
			},
			IdentityHandler: func(c *gin.Context) interface{} {
				db := services.GetDB()
				claims := jwt.ExtractClaims(c)
				user := new(models.User)
				db.Where("username = ?", claims["username"]).First(user)
				return user
			},
			Authenticator: func(c *gin.Context) (interface{}, error) {
				var query forms.LoginForm

				if err := c.ShouldBindJSON(&query); err != nil {
					return nil, jwt.ErrMissingLoginValues
				}
				db := services.GetDB()

				user := new(models.User)
				if err := db.Where("username = ?", query.Username).First(user).Error; err != nil {
					return nil, jwt.ErrFailedAuthentication
				}

				if user.ValidPassword(query.Password) {
					return user, nil
				}

				return nil, jwt.ErrFailedAuthentication
			},
			Unauthorized: func(c *gin.Context, code int, msg string) {
				c.JSON(code, gin.H{
					"error": msg,
				})
			},
		})
		if err != nil {
			panic(err)
		}
		_authMiddleware = authMiddleware
	}

	return _authMiddleware
}
