package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/meastblue/godo/model"
	"github.com/meastblue/godo/service"
	"github.com/meastblue/godo/util"
)

func Login(c *gin.Context) {
	auth := model.Auth{}
	if err := c.ShouldBindJSON(&auth); err != nil {
		util.SendJsonError(c, err.Error())
		return
	}

	err := service.Login(&auth)
	if err != nil {
		util.SendJsonError(c, err.Error())
		return
	}

	util.SendJsonOK(c, nil)
}

func Logout(c *gin.Context) {
	au, err := util.ExtractTokenMetadata(c.Request)
	if err != nil {
		util.SendJsonUnauthorized(c, err.Error())
		return
	}
	deleted, delErr := util.DeleteAuth(au.AccessUuid)
	if delErr != nil || deleted == 0 { //if any goes wrong
		util.SendJsonUnauthorized(c, err.Error())
		return
	}
	util.SendJsonOK(c, nil)
}
