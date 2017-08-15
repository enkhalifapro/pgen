package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"github.com/enkhalifapro/pgen/services"
)

type CategoryController struct {
	UserService     *services.UserService           `inject:""`
	CategoryService *services.CategoryService   `inject:""`
}

func (c *CategoryController) GetAll(ctx *gin.Context) {
	pageSize, err := strconv.Atoi(ctx.Query("pageSize"))
	if err != nil {
		pageSize = 100
	}

	pageNumber, err := strconv.Atoi(ctx.Query("pageNumber"))
	if err != nil {
		pageNumber = 1
	}

	categories, count := c.CategoryService.QueryByPage(&bson.M{}, pageSize, pageNumber)
	ctx.JSON(http.StatusOK, gin.H{"data": gin.H{"categories": categories, "count": count}})
}

func (c *PortalController) GetBySlug(ctx *gin.Context) {
	name := ctx.Param("name")
	portal, err := c.PortalService.FindOne(&bson.M{"slug": name})
	if err != nil {
		Error(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, portal)
}

func (c *PortalController) GetByOwnerName(ctx *gin.Context) {
	ownerSlug := ctx.Param("ownerName")
	owner, err := c.UserService.FindOne(&bson.M{"slug": ownerSlug})
	if err != nil {
		Error(ctx, http.StatusBadRequest, fmt.Errorf("%v", "owner not found"))
		return
	}

	portals, err := c.PortalService.FindAll(&bson.M{"ownerid": owner.Id.Hex()})
	response := gin.H{
		"portals": portals,
		"owner": gin.H{
			"Id":        owner.Id.Hex(),
			"name":      owner.UserName,
			"email":     owner.Email,
			"slug":      owner.Slug,
			"firstName": owner.FirstName,
			"lastName":  owner.LastName,
			"image":     owner.Image,
		},
	}
	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

func (c *PortalController) CreatePortal(ctx *gin.Context) {
	var portal models.Portal
	if err := ctx.Bind(&portal); err != nil {
		Error(ctx, http.StatusBadRequest, err)
		return
	}

	if err := portal.IsValid(); err != nil {
		Error(ctx, http.StatusBadRequest, err)
		return
	}

	portal.Name = strings.ToLower(portal.Name)

	if _, err := c.PortalService.FindOne(&bson.M{"name": portal.Name}); err == nil {
		Error(ctx, http.StatusBadRequest, fmt.Errorf("portal name '%v' is already exist", portal.Name))
		return
	}

	user, err := GetSessUser(ctx, c.UserService)
	if err != nil {
		Error(ctx, http.StatusUnauthorized, err)
		return
	}

	portal.OwnerId = user.Id.Hex()
	if err := c.PortalService.Insert(portal.OwnerId, &portal); err != nil {
		Error(ctx, http.StatusInternalServerError, err)
		return
	}

	pageDefCnt, err := c.SitePageService.GetDefaultContent()
	if err != nil {
		Error(ctx, http.StatusInternalServerError, err)
		return
	}

	sitePage := &models.SitePage{Title: "landing", PortalId: portal.Id.Hex(), Html: pageDefCnt}
	if err := c.SitePageService.Insert(portal.OwnerId, sitePage); err != nil {
		Error(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": gin.H{"portal": portal}})
}

func (c *PortalController) IsAvailableName(ctx *gin.Context) {
	name := ctx.Param("name")
	// name exist
	if _, err := c.PortalService.FindOne(&bson.M{"name": name}); err == nil {
		ctx.JSON(http.StatusOK, gin.H{"data": gin.H{"available": false}})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": gin.H{"available": true}})
}

func (c *PortalController) UpdateBySlug(ctx *gin.Context) {
	slug := ctx.Param("name")
	var portal models.Portal
	err := ctx.Bind(&portal)
	if err != nil {
		Error(ctx, http.StatusBadRequest, err)
		return
	}

	err = portal.IsValid()
	if err != nil {
		Error(ctx, http.StatusBadRequest, err)
		return
	}

	user, err := GetSessUser(ctx, c.UserService)
	if err != nil {
		Error(ctx, http.StatusUnauthorized, err)
		return
	}

	err = c.PortalService.UpdateBySlug(user.Id.Hex(), slug, &portal)
	if err != nil {
		Error(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": gin.H{"portal": portal}})
}

// DeleteBySlug portal and all related site pages.
func (c *PortalController) DeleteBySlug(ctx *gin.Context) {
	user, err := GetSessUser(ctx, c.UserService)
	if err != nil {
		Error(ctx, http.StatusUnauthorized, err)
		return
	}

	slug := ctx.Param("name")
	portal, err := c.PortalService.FindOne(&bson.M{"slug": slug})
	if err != nil {
		Error(ctx, http.StatusNotFound, err)
		return
	}

	if err = c.PortalService.DeleteBySlug(user.Id.Hex(), slug); err != nil {
		Error(ctx, http.StatusInternalServerError, err)
		return
	}

	if err = c.SitePageService.DeleteByPortalID(portal.Id.Hex()); err != nil {
		Error(ctx, http.StatusInternalServerError, fmt.Errorf("delete site pages: %v", err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": gin.H{"portal": slug}})
}

func (c *PortalController) GetAllPortalCategories(ctx *gin.Context) {
	name := ctx.Param("name")
	portal, err := c.PortalService.FindOne(&bson.M{"slug": name})
	if err != nil {
		Error(ctx, http.StatusBadRequest, err)
		return
	}

	portalCategories, err := c.PortalCategoryService.Find(&bson.M{"portalid": portal.Id.Hex()})
	if err != nil {
		Error(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": portalCategories})
}
