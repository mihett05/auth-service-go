package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mihett05/auth-service-go/controllers"
	"github.com/mihett05/auth-service-go/core/models"
	"github.com/mihett05/auth-service-go/libs"
	"github.com/mihett05/auth-service-go/services"
	"net/http"
)

func main() {
	app := gin.Default()

	app.Use(cors.Default())

	db := services.GetDB()
	_ = db.AutoMigrate(&models.User{})

	controllers.InitAuthRoutes(app.Group("/auth"))

	http.ListenAndServe(":" + libs.EnvDefault("PORT"), app)
}
