package controllers

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/enkhalifapro/pgen/services"
	"github.com/enkhalifapro/pgen/models"
)

// GetSessUser extract session token and return owner user.
func GetSessUser(ctx *gin.Context, userService *services.UserService) (*models.User, error) {
	authToken := ctx.Request.Header.Get("Authorization")
	authToken = strings.Replace(authToken, "Bearer ", "", -1)
	user, err := userService.CurrentUser(authToken)
	if err != nil {
		return nil, fmt.Errorf("user from session: %v", err)
	}
	return user, nil
}

// Error returns to client JSON object with 'error' field and statusCode.
func Error(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{"error": err.Error()})
}
