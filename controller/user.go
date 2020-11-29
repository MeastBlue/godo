package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/meastblue/godo/model"
	"github.com/meastblue/godo/service"
	"github.com/meastblue/godo/util"
)

func GetUsers(c *gin.Context) {
	users, err := service.GetUsers()
	if err != nil {
		util.SendJsonError(c, err)
		return
	}

	util.SendJsonOK(c, &users)
}

func GetUser(c *gin.Context) {
	user, err := service.GetUser(c.Param("id"))
	if err != nil {
		util.SendJsonError(c, err.Error())
		return
	}

	util.SendJsonOK(c, &user)
}

func AddUser(c *gin.Context) {
	u := model.User{}
	if err := c.ShouldBindJSON(&u); err != nil {
		util.SendJsonError(c, err.Error())
		return
	}

	err := service.AddUser(&u)
	if err != nil {
		util.SendJsonError(c, err.Error())
		return
	}

	util.SendJsonCreated(c, "created")
}

func UpdateUser(c *gin.Context) {
	u := model.User{}
	if err := c.ShouldBindJSON(&u); err != nil {
		util.SendJsonError(c, err.Error())
		return
	}

	err := service.UpdateUser(&u)
	if err != nil {
		util.SendJsonError(c, err.Error())
		return
	}

	util.SendJsonNoCotent(c, nil)
}

func DeleteUser(c *gin.Context) {
	err := service.DeleteUser(c.Param("id"))
	if err != nil {
		util.SendJsonError(c, err.Error())
		return
	}

	util.SendJsonNoCotent(c, nil)
}
