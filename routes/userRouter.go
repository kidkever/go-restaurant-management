package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/kidkever/go-restautrant-management/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.GET("/users/signup", controller.SignUp())
	incomingRoutes.GET("/users/login", controller.Login())
}