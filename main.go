package main

import (
	"api/auth"
	"api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// api
	api := router.Group("/api")
	routes.UserRoute(api)

	// public
	router.Static("/public", "./public")

	// auth
	router.Use(auth.Authenticate)

	// private
	router.Static("/private", "./private")

	router.Run("localhost:8080")
}
