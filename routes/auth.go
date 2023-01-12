package routes

import (
	"go-authentication-system/controllers"
	"go-authentication-system/utils"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {

	router.POST("/login", controllers.Login())
	router.POST("/signup", controllers.Signup())
}

func HomeRoute(router *gin.Engine) {
	router.Use(utils.Authenticate())
	router.GET("/", controllers.Home())
}