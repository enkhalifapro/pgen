package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"github.com/enkhalifapro/pgen/utilities"
	"github.com/enkhalifapro/pgen/services"
)

type UserController struct {
	CryptUtil           *utilities.CryptUtil          `inject:""`
	UserService         *services.UserService         `inject:""`
}

func (c *UserController) Root(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Service is working")
}

func (c *UserController) GetAll(ctx *gin.Context) {
	pageSize, err := utilities.ToInt(ctx.Query("pageSize"))
	if err != nil {
		pageSize = 100
	}

	pageNumber, err := utilities.ToInt(ctx.Query("pageNumber"))
	if err != nil {
		pageNumber = 1
	}

	users, count, err := c.UserService.QueryByPage(&bson.M{}, pageSize, pageNumber)
	if err != nil {
		Error(ctx, http.StatusInternalServerError, fmt.Errorf("query users: %v", err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": gin.H{"users": users, "count": count}})
}