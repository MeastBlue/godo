package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/meastblue/godo/model"
	"github.com/meastblue/godo/service"
	"github.com/meastblue/godo/util"
)

// GetTasks controller function
func GetTasks(c *gin.Context) {
	userID, err := util.GetUserIDFromJwt(c.Request)
	if err != nil {
		util.SendJsonUnauthorized(c, err.Error())
		return
	}

	tasks, err := service.GetTasks(userID)
	if err != nil {
		util.SendJsonError(c, err)
		return
	}

	util.SendJsonOK(c, &tasks)
}

// GetTask controller function
func GetTask(c *gin.Context) {
	task, err := service.GetTask(c.Param("id"))
	if err != nil {
		util.SendJsonError(c, err.Error())
		return
	}

	util.SendJsonOK(c, &task)
}

// AddTask controller function
func AddTask(c *gin.Context) {
	t := model.Task{}
	if err := c.ShouldBindJSON(&t); err != nil {
		util.SendJsonError(c, err.Error())
		return
	}

	userID, err := util.GetUserIDFromJwt(c.Request)
	if err != nil {
		util.SendJsonUnauthorized(c, err.Error())
		return
	}

	t.UserID = userID
	id, err := service.AddTask(&t)
	if err != nil {
		util.SendJsonError(c, err.Error())
		return
	}

	util.SendJsonCreated(c, id)
}

// UpdateTask controller function
func UpdateTask(c *gin.Context) {
	t := model.Task{}
	if err := c.ShouldBindJSON(&t); err != nil {
		util.SendJsonError(c, err.Error())
		return
	}

	err := service.UpdateTask(&t)
	if err != nil {
		util.SendJsonError(c, err.Error())
		return
	}

	util.SendJsonNoCotent(c, nil)

}

// DeleteTask controller function
func DeleteTask(c *gin.Context) {
	err := service.DeleteTask(c.Param("id"))
	if err != nil {
		util.SendJsonError(c, err.Error())
		return
	}

	util.SendJsonNoCotent(c, nil)

}
