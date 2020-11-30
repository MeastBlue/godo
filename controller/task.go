package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/meastblue/godo/service"
	"github.com/meastblue/godo/util"
)

func GetTasks(c *gin.Context) {
	tasks, err := service.GetTasks()
	if err != nil {
		util.SendJsonError(c, err)
		return
	}

	util.SendJsonOK(c, &tasks)
}

func GetTask(c *gin.Context) {
	task, err := service.GetTask(c.Param("id"))
	if err != nil {
		util.SendJsonError(c, err.Error())
		return
	}

	util.SendJsonOK(c, &task)
}

func AddTask(c *gin.Context) {
}

func UpdateTask(c *gin.Context) {

}

func DeleteTask(c *gin.Context) {

}
