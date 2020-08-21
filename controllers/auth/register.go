package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/mihett05/auth-service-go/core/forms"
	"github.com/mihett05/auth-service-go/core/models"
	"github.com/mihett05/auth-service-go/libs"
	"github.com/mihett05/auth-service-go/midlewares"
	"github.com/mihett05/auth-service-go/services"
	"gorm.io/gorm"
	"net/http"
)

func Register(c *gin.Context)  {
	var query forms.RegisterForm

	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		db := services.GetDB()

		user := new(models.User)
		if err := db.Where("username = ? OR email = ?", query.Username, query.Email).First(user).Error
			err == gorm.ErrRecordNotFound {
			user.Username = query.Username
			user.Email = query.Email
			user.Salt, user.Password = libs.GenerateHash(query.Password)

			db.Save(user).Commit()

			token, expire, err := midlewares.AuthMiddleware().TokenGenerator(user)

			if err != nil {
				panic(err)
			}

			midlewares.AuthMiddleware().LoginResponse(c, http.StatusOK, token, expire)

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "login is already in use",
			})
		}
	}
}
