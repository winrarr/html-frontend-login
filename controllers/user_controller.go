package controllers

import (
	"api/auth"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var users = []models.User{}

func AddUser(user models.User) {
	users = append(users, user)
}

func AuthoriseUser(user models.User) bool {
	for _, u := range users {
		if u == user {
			return true
		}
	}
	return false
}

// POST /register
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	AddUser(user)
	c.Redirect(http.StatusMovedPermanently, "/public/login.html")
}

// POST /login
func Login(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if !AuthoriseUser(user) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.SetCookie("session_token", auth.NewSession(), 2, "/", "localhost", false, false)
	c.Redirect(http.StatusMovedPermanently, "/private/secret.html")
}

// GET /users
func GetAllUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
