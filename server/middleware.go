package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"github.com/enkhalifapro/pgen/services"
)

// Middleware collection struct.
type Middleware struct {
	SessionService *services.SessionService `inject:""`
}

// ReqAuthUser returns middleware which requires authenticated user for request.
func (s *Middleware) ReqAuthUser(c *gin.Context) {
	authToken := c.Request.Header.Get("Authorization")
	authToken = strings.Replace(authToken, "Bearer ", "", -1)

	if s.SessionService.Valid(authToken) == false {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthoirzed user"})
		c.Abort()
		return
	}
	c.Next()
}
