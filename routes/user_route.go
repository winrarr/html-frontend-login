package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

// /api
func UserRoute(router *gin.RouterGroup) {
	router.POST("/register", controllers.CreateUser)
	router.POST("/login", controllers.Login)
	router.GET("/users", controllers.GetAllUsers)
}
