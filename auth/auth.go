package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var sessions = map[string]time.Time{}

func NewSession() string {
	sessionToken := uuid.NewString()
	sessions[sessionToken] = time.Now()
	return sessionToken
}

func RedirectToLogin(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/public/register.html")
}

func Authenticate(c *gin.Context) {
	println("thseria")
	sessionCookie, err := c.Request.Cookie("session_token")
	if err != nil {
		RedirectToLogin(c)
		return
	}

	val, ok := sessions[sessionCookie.Value]
	if !ok {
		RedirectToLogin(c)
		return
	}

	expires := val.Add(2 * time.Minute)
	if expires.Before(time.Now()) {
		RedirectToLogin(c)
		return
	}
}
