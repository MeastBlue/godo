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
}
