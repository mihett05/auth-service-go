package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mihett05/auth-service-go/controllers/auth"
	"github.com/mihett05/auth-service-go/midlewares"
)

func InitAuthRoutes(group *gin.RouterGroup)  {
	group.POST("/login/", midlewares.AuthMiddleware().LoginHandler)
	group.POST("/register/", auth.Register)
}
