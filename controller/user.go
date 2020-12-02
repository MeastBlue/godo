package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/meastblue/godo/model"
	"github.com/meastblue/godo/service"
	"github.com/meastblue/godo/util"
)

// GetUsers controller function
func GetUsers(c *gin.Context) {
	users, err := service.GetUsers()
	if err != nil {
		util.SendJsonError(c, err)
		return
	}

	util.SendJsonOK(c, &users)
}

// GetUser controller function
func GetUser(c *gin.Context) {
	user, err := service.GetUser(c.Param("id"))
	if err != nil {
		util.SendJsonError(c, err.Error())
		return
	}

	util.SendJsonOK(c, &user)
}

// AddUser controller function
func AddUser(c *gin.Context) {
	u := model.User{}
	if err := c.ShouldBindJSON(&u); err != nil {
		util.SendJsonError(c, err.Error())
		return
	}

	id, err := service.AddUser(&u)
	if err != nil {
		util.SendJsonError(c, err.Error())
		return
	}

	token, err := util.CreateToken(id)
	if err != nil {
		util.SendJsonError(c, err.Error())
		return
	}

	err = util.CreateAuth(id, token)
	if err != nil {
		util.SendJsonError(c, err.Error())
		return
	}

	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}

	util.SendJsonCreated(c, tokens)
}

// UpdateUser controller function
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

// DeleteUser controller function
func DeleteUser(c *gin.Context) {
	err := service.DeleteUser(c.Param("id"))
	if err != nil {
		util.SendJsonError(c, err.Error())
		return
	}

	util.SendJsonNoCotent(c, nil)
}
